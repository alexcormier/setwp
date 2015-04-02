# setwp
A command line utility to set the same wallpaper for every space and desktop on OS X Yosemite.

## Installation
There are 3 ways to install.

1. As any other Go package : `go get github.com/alexandrecormier/setwp`.

2. With [Homebrew](http://brew.sh/) : `brew install https://raw.githubusercontent.com/alexandrecormier/setwp/master/homebrew/setwp.rb`.

3. Download the binary for [32](https://github.com/alexandrecormier/setwp/releases/download/v0.1.1/setwp-i386-v0.1.1.tar.gz) or [64](https://github.com/alexandrecormier/setwp/releases/download/v0.1.1/setwp-amd64-v0.1.1.tar.gz) bit.

## Usage
Running `setwp <wallpaper>` will set \<wallpaper\> to fill the screen. For more options, see help :

~~~
Usage:
  setwp [-f | -s | -c | -t] <wallpaper>
  setwp -h | -v

Options:
  -f --fit      Fit wallpaper to screen.
  -s --stretch  Stretch wallpaper to fill screen.
  -c --center   Center wallpaper.
  -t --tile     Tile wallpaper.
  -h --help     Show this help message.
  -v --version  Show version information.
~~~

## Todo
- [x] Option to set picture position (fill, fit, strech, center, tile)
- [ ] Image validation
- [x] bash/zsh completion
- [ ] Set whole folder as wallpaper ?
