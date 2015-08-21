// Package args is the bridge between command line arguments and wallpaper preferences.
package args

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alexandrecormier/setwp/pref"
	"github.com/alexandrecormier/setwp/pref/event"
	"github.com/alexandrecormier/setwp/pref/position"
	"github.com/docopt/docopt-go"
)

const (
	programName = "setwp"

	usage = `Sets wallpaper to <wallpaper>. Fills the screen by default.

Usage:
  %[1]s [--fit | --stretch | --center | --tile] <wallpaper>
  %[1]s (--interval=<s> | --login | --wake) [--random] [--fit | --stretch | --center | --tile] <directory>
  %[1]s --help | --version

Options:
  -f --fit      Fit wallpaper to screen.
  -s --stretch  Stretch wallpaper to fill screen.
  -c --center   Center wallpaper.
  -t --tile     Tile wallpaper.
  -h --help     Show this help message.
  -v --version  Show version information.

Directory options:
  -i --interval=<s>  Interval at which to change wallpaper in seconds.
  -l --login         Change wallpaper when logging in.
  -w --wake          Change wallpaper when waking from sleep.
  -r --random        Randomize wallpaper selection.

`

	version = "%s version 1.0.1"
)

// Type arg represents the preferences set by an argument.
type argPrefs struct {
	// Preferences to set when this argument is specified.
	flagPrefs pref.Prefs

	// Preferences to set to the value of this argument.
	valuePrefs []pref.KeyType

	// Function to validate and extract the preference's value from the argument's value.
	value func(interface{}) (interface{}, error)
}

var (
	defaultPrefs = pref.Prefs{
		pref.Position: position.Fill,
	}

	argMap = map[string]argPrefs{
		"--fit": argPrefs{
			flagPrefs:  pref.Prefs{pref.Position: position.Fit},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--stretch": argPrefs{
			flagPrefs:  pref.Prefs{pref.Position: position.Stretch},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--center": argPrefs{
			flagPrefs:  pref.Prefs{pref.Position: position.Center},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--tile": argPrefs{
			flagPrefs:  pref.Prefs{pref.Position: position.Tile},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--interval": argPrefs{
			flagPrefs:  pref.Prefs{pref.ChangeEvent: event.Interval},
			valuePrefs: []pref.KeyType{pref.Interval},
			value: func(value interface{}) (interface{}, error) {
				return strconv.ParseUint(value.(string), 10, 0)
			},
		},
		"--login": argPrefs{
			flagPrefs:  pref.Prefs{pref.ChangeEvent: event.Login},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--wake": argPrefs{
			flagPrefs:  pref.Prefs{pref.ChangeEvent: event.Wake},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--random": argPrefs{
			flagPrefs:  pref.Prefs{pref.Random: true},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"<wallpaper>": argPrefs{
			flagPrefs:  pref.Prefs{},
			valuePrefs: []pref.KeyType{pref.Wallpaper},
			value: func(value interface{}) (interface{}, error) {
				info, err := os.Stat(value.(string))
				if err != nil {
					return value, err
				}
				if info.IsDir() {
					return value, fmt.Errorf("invalid wallpaper: %s is a directory", value)
				}
				return value, nil
			},
		},
		"<directory>": argPrefs{
			flagPrefs:  pref.Prefs{},
			valuePrefs: []pref.KeyType{pref.Directory},
			value: func(value interface{}) (interface{}, error) {
				info, err := os.Stat(value.(string))
				if err != nil {
					return value, err
				}
				if !info.IsDir() {
					return value, fmt.Errorf("%s is not a directory", value)
				}
				return value, nil
			},
		},
	}
)

// Parses command line arguments and returns the preferences to apply or an error if there is any.
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
			if argPref, ok := argMap[optKey]; ok {
				// specifying this option has an effect that's not default so we process it
				prefValue, err := argPref.value(optValue)
				if err != nil {
					return parsedArgs, err
				}

				for key, value := range argPref.flagPrefs {
					parsedArgs[key] = value
				}
				for _, key := range argPref.valuePrefs {
					parsedArgs[key] = prefValue
				}
			} else {

			}
		}
	}
	return parsedArgs, nil
}
