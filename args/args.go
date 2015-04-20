package args

import (
	"fmt"
	"github.com/alexandrecormier/setwp/pref"
	"github.com/alexandrecormier/setwp/pref/event"
	"github.com/alexandrecormier/setwp/pref/position"
	"github.com/docopt/docopt-go"
	"os"
	"strconv"
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

	version = "%s version 1.0"
)

type arg struct {
	flagPrefs  []pref.Pref
	valuePrefs []pref.KeyType
	value      func(interface{}) (interface{}, error)
}

var (
	defaultPrefs = []pref.Pref{
		pref.Pref{pref.Position, position.Fill},
	}

	argMap = map[string]arg{
		"--fit": arg{
			flagPrefs:  []pref.Pref{pref.Pref{pref.Position, position.Fit}},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--stretch": arg{
			flagPrefs:  []pref.Pref{pref.Pref{pref.Position, position.Stretch}},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--center": arg{
			flagPrefs:  []pref.Pref{pref.Pref{pref.Position, position.Center}},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--tile": arg{
			flagPrefs:  []pref.Pref{pref.Pref{pref.Position, position.Tile}},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--interval": arg{
			flagPrefs:  []pref.Pref{pref.Pref{pref.ChangeEvent, event.Interval}},
			valuePrefs: []pref.KeyType{pref.Interval},
			value: func(value interface{}) (interface{}, error) {
				return strconv.ParseUint(value.(string), 10, 0)
			},
		},
		"--login": arg{
			flagPrefs:  []pref.Pref{pref.Pref{pref.ChangeEvent, event.Login}},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--wake": arg{
			flagPrefs:  []pref.Pref{pref.Pref{pref.ChangeEvent, event.Wake}},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"--random": arg{
			flagPrefs:  []pref.Pref{pref.Pref{pref.Random, true}},
			valuePrefs: []pref.KeyType{},
			value:      func(value interface{}) (interface{}, error) { return value, nil },
		},
		"<wallpaper>": arg{
			flagPrefs:  []pref.Pref{},
			valuePrefs: []pref.KeyType{pref.Wallpaper},
			value: func(value interface{}) (interface{}, error) {
				info, err := os.Stat(value.(string))
				if err != nil {
					return value, err
				}
				if info.IsDir() {
					return value, fmt.Errorf("invalid wallpaper; %s is a directory", value)
				}
				return value, nil
			},
		},
		"<directory>": arg{
			flagPrefs:  []pref.Pref{},
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

func Parse() ([]pref.Pref, error) {
	parsedArgs := defaultPrefs
	opts, err := docopt.Parse(fmt.Sprintf(usage, programName), nil, true, fmt.Sprintf(version, programName), false)
	if err != nil {
		return []pref.Pref{}, fmt.Errorf("cannot parse arguments (%s)", err)
	}
	for optKey, optValue := range opts {
		if b, ok := optValue.(bool); !ok && optValue != nil || b {
			if a, ok := argMap[optKey]; ok {
				for _, p := range a.flagPrefs {
					parsedArgs = append(parsedArgs, p)
				}
				for _, pKey := range a.valuePrefs {
					value, err := a.value(optValue)
					if err != nil {
						return []pref.Pref{}, err
					}
					parsedArgs = append(parsedArgs, pref.Pref{pKey, value})
				}
			}
		}
	}
	return parsedArgs, nil
}
