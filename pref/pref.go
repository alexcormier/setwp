// Package pref contains constants for the preferences to set in the database.
package pref

type KeyType uint8
type ValueType interface{}
type Prefs map[KeyType]ValueType

const (
	// Path to the wallpaper
	Wallpaper KeyType = 1

	// Position of the wallpaper
	// Value should be one of the constants in the position package
	Position KeyType = 2

	// Event on which to change wallpaper
	// Value should be one of the constants in the event package
	ChangeEvent KeyType = 9

	// Directory of changing wallpapers
	Directory KeyType = 10

	// Interval at which to change wallpaper in seconds
	// Only has an effect if 'Directory' is set
	Interval KeyType = 11

	// Whether or not to select wallpaper at random when changing
	// Only has an effect if 'Directory' is set
	Random KeyType = 12

	// Current wallpaper if 'Directory' is set
	Current KeyType = 16
)

// Type Pref represents a preference in the database.
type Pref struct {
	Key   KeyType
	Value ValueType
}
