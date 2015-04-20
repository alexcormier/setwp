# setwp
A command line utility to set the same wallpaper for every space and desktop on OS X Yosemite.

## Installation
There are 3 ways to install.

1. As any other Go package : `go get github.com/alexandrecormier/setwp`.

2. With [Homebrew](http://brew.sh/) : `brew install https://raw.githubusercontent.com/alexandrecormier/setwp/master/homebrew/setwp.rb`.

3. Download the binary for [32](https://github.com/alexandrecormier/setwp/releases/download/v0.1.1-1/setwp-i386-v0.1.1-1.tar.gz) or [64](https://github.com/alexandrecormier/setwp/releases/download/v0.1.1-1/setwp-amd64-v0.1.1-1.tar.gz) bit.

### Completion
If you want command completion, make sure you have installed the bash-completion or zsh-completions package, depending on your shell.

If you installed with Homebrew, bash and zsh completion scripts are automatically installed; else you'll have to install them manually.

#### bash
Drop the `setwp-completion.bash` file in `~/.bash_completion`, `/etc/bash_completion.d`, or any other directory that gets sourced automatically.

#### zsh
Put the `setwp-completion.zsh` file in a folder that's in your fpath. To check what folders are in your fpath, run `echo $fpath`. To add a directory to your fpath, add `fpath+="<directory>"` to `~/.zshrc`.

## Usage
Running `setwp <wallpaper>` will set \<wallpaper\> to fill the screen. For more options, see help :
~~~
Usage:
  setwp [--fit | --stretch | --center | --tile] <wallpaper>
  setwp --help | --version

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
