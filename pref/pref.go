package pref

type KeyType uint8

const (
	Wallpaper   KeyType = 1
	Position    KeyType = 2
	ChangeEvent KeyType = 9
	Directory   KeyType = 10
	Interval    KeyType = 11
	Random      KeyType = 12
	Current     KeyType = 16
)

type Pref struct {
	Key   KeyType
	Value interface{}
}
