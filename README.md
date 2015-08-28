# setwp
A command line utility to set the same wallpaper settings for every space and desktop on OS X Yosemite.

## Installation
There are 3 ways to install.

1. As any other Go package : `go get github.com/alexandrecormier/setwp`.

2. With [Homebrew](http://brew.sh/) : `brew install alexandrecormier/personal/setwp`.

3. Download [the binary](https://github.com/alexandrecormier/setwp/releases/download/v1.0.1/setwp-v1.0.1.tar.gz).

### Completion
If you want command completion, make sure you have installed the bash-completion or zsh-completions package, depending on your shell.

If you installed with Homebrew, bash and zsh completion scripts are automatically installed; else you'll have to install them manually.

#### bash
Drop the `setwp-completion.bash` file in `~/.bash_completion`, `/etc/bash_completion.d`, or any other directory that gets sourced automatically.

#### zsh
Put the `setwp-completion.zsh` file in a folder that's in your fpath. To check what folders are in your fpath, run `echo $fpath`. To add a directory to your fpath, add `fpath+="<directory>"` to `~/.zshrc`.

## Usage
~~~
Sets wallpaper to <wallpaper> or a <directory> of wallpapers.
Fills the screen by default.

Usage:
  setwp [--fit | --stretch | --center | --tile] <wallpaper>
  setwp (--interval=<s> | --login | --wake) [--random] [--fit | --stretch | --center | --tile] <directory>
  setwp --help | --version

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
~~~

## Limitations
When setting wallpaper to a directory, spaces and desktops are not synced. Having them in sync is outside the scope of this project, as setwp changes wallpaper settings for all spaces and lets the OS change the actual wallpaper.

## Todo
- [x] Option to set picture position (fill, fit, stretch, center, tile)
- [ ] Image validation
- [x] bash/zsh completion
- [x] Set whole folder as wallpaper ?
