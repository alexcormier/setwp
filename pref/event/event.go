// Package event contains constants for the events used to change wallpaper.
package event

const (
	// Change wallpaper at a certain interval
	Interval uint8 = 1

	// Change wallpaper when logging in
	Login uint8 = 2

	// Change wallpaper when waking from sleep
	Wake uint8 = 3
)
