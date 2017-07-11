// Package args is the bridge between command line arguments and wallpaper preferences.
package args

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/alexcormier/setwp/pref"
	"github.com/alexcormier/setwp/pref/event"
	"github.com/alexcormier/setwp/pref/position"
	docopt "github.com/docopt/docopt-go"
)

const (
	programName = "setwp"

	usage = `Sets wallpaper to <wallpaper>, a <directory> of wallpapers or a solid color.
Fills the screen by default.

Usage:
  %[1]s [--fill | --fit | --stretch | --center | --tile] [--color=<hex>] <wallpaper>
  %[1]s (--interval=<s> | --login | --wake) [--random] [--fill | --fit | --stretch | --center | --tile] [--color=<hex>] <directory>
  %[1]s --color=<hex>
  %[1]s --help | --version

Options:
  -C --color=<hex>  Color to fill the screen with, as an RGB hex code.
  -h --help         Show this help message.
  -v --version      Show version information.

Fit options:
  -F --fill         Scale wallpaper to fill screen [default].
  -f --fit          Fit wallpaper to screen.
  -s --stretch      Stretch wallpaper to fill screen.
  -c --center       Center wallpaper, scaling it down if it is too large.
  -t --tile         Tile wallpaper.

Directory options:
  -i --interval=<s>  Interval at which to change wallpaper, in seconds.
  -l --login         Change wallpaper when logging in.
  -w --wake          Change wallpaper when waking from sleep.
  -r --random        Randomize wallpaper selection.

`

	version = "%s version 1.1.1"
)

type argument struct {
	prefs func(value interface{}) (pref.Prefs, error)
}

// Type arg represents the preferences set by an argument.
type argPrefs struct {
	// Preferences to set when this argument is specified.
	flagPrefs pref.Prefs

	// Preferences to set to the value of this argument.
	valuePrefs []pref.KeyType

	// Function to validate and extract the preference's value from the argument's value.
	value func(interface{}) (pref.ValueType, error)
}

var (
	defaultPrefs = pref.Prefs{
		pref.Position:   position.Fill,
		pref.SolidColor: true,
	}

	argMap = map[string]argument{
		"--fit": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				return pref.Prefs{pref.Position: position.Fit}, nil
			},
		},
		"--stretch": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				return pref.Prefs{pref.Position: position.Stretch}, nil
			},
		},
		"--center": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				return pref.Prefs{pref.Position: position.Center}, nil
			},
		},
		"--tile": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				return pref.Prefs{pref.Position: position.Tile}, nil
			},
		},
		"--color": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				colorString := value.(string)
				if len(colorString) == 6 {
					color, err := strconv.ParseUint(colorString, 16, 32)
					if err == nil {
						prefs := pref.Prefs{
							pref.Red:   float64((color>>16)&0xFF) / 255,
							pref.Green: float64((color>>8)&0xFF) / 255,
							pref.Blue:  float64(color&0xFF) / 255,
						}
						return prefs, nil
					}
				}
				return nil, errors.New("invalid color")
			},
		},
		"--interval": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				interval, err := strconv.ParseUint(value.(string), 10, 0)
				if err != nil {
					return nil, errors.New("invalid interval")
				}
				prefs := pref.Prefs{
					pref.ChangeEvent: event.Interval,
					pref.Interval:    interval,
				}
				return prefs, nil
			},
		},
		"--login": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				return pref.Prefs{pref.ChangeEvent: event.Login}, nil
			},
		},
		"--wake": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				return pref.Prefs{pref.ChangeEvent: event.Wake}, nil
			},
		},
		"--random": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				return pref.Prefs{pref.Random: true}, nil
			},
		},
		"<wallpaper>": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				path, err := filepath.Abs(value.(string))
				if err != nil {
					return nil, fmt.Errorf("invalid wallpaper '%s': %v", value, err)
				}
				info, err := os.Stat(path)
				if err != nil {
					return nil, fmt.Errorf("invalid wallpaper '%s': %v", value, err.(*os.PathError).Err)
				}
				if info.IsDir() {
					return nil, fmt.Errorf("invalid wallpaper: %s is a directory", value)
				}
				prefs := pref.Prefs{
					pref.Wallpaper:  path,
					pref.SolidColor: false,
				}
				return prefs, nil
			},
		},
		"<directory>": argument{
			prefs: func(value interface{}) (pref.Prefs, error) {
				path, err := filepath.Abs(value.(string))
				if err != nil {
					return nil, fmt.Errorf("invalid wallpaper '%s': %v", value, err)
				}
				info, err := os.Stat(path)
				if err != nil {
					return nil, fmt.Errorf("invalid wallpaper '%s': %v", value, err.(*os.PathError).Err)
				}
				if !info.IsDir() {
					return nil, fmt.Errorf("%s is not a directory", path)
				}

				prefs := pref.Prefs{
					pref.Directory:  path,
					pref.Current:    path, // partial fix for recent macOS versions (see issue #3)
					pref.SolidColor: false,
				}
				return prefs, nil
			},
		},
	}
)

// Parse command line arguments and returns the preferences to apply or an error if there is any.
// If the help or version flag is passed, the corresponding message is printed and the program exits.
// If the arguments don't match one of the usage patterns, the usage message is printed and the program exits.
func Parse() (pref.Prefs, error) {
	parsedArgs := defaultPrefs

	opts, err := docopt.Parse(fmt.Sprintf(usage, programName), nil, true, fmt.Sprintf(version, programName), true)
	if err != nil {
		return parsedArgs, fmt.Errorf("cannot parse arguments (%s)", err)
	}

	for optKey, optValue := range opts {
		if b, ok := optValue.(bool); !ok && optValue != nil || b {
			// this option has a value or is a flag and was specified
			if argument, ok := argMap[optKey]; ok {
				// specifying this option has an effect that's not default so we process it

				prefs, err := argument.prefs(optValue)
				if err != nil {
					return parsedArgs, err
				}

				for key, value := range prefs {
					parsedArgs[key] = value
				}
			} else {
				// this option was not specified
			}
		}
	}
	return parsedArgs, nil
}
