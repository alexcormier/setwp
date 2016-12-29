# setwp
A command line utility to set the same wallpaper settings for every space and desktop on OS X Yosemite and up.

## Installation
There are 3 ways to install.

1. As any other Go package : `go get github.com/alexcormier/setwp`.

2. With [Homebrew](http://brew.sh/) : `brew install alexcormier/personal/setwp`.

3. Download [the binary](https://github.com/alexcormier/setwp/releases/download/v1.0.3/setwp-v1.0.3.tar.gz).

### Completion
If you want command completion, make sure you have installed the bash-completion or zsh-completions package, depending on your shell.

If you installed with Homebrew, bash and zsh completion scripts are automatically installed; else you'll have to install them manually.

#### bash
Drop the `setwp-completion.bash` file in `~/.bash_completion`, `/etc/bash_completion.d`, or any other directory that gets sourced automatically.

#### zsh
Put the `setwp-completion.zsh` file in a folder that's in your fpath. To check what folders are in your fpath, run `echo $fpath`. To add a directory to your fpath, add `fpath+="<directory>"` to `~/.zshrc`.

## Usage
~~~
Sets wallpaper to <wallpaper>, a <directory> of wallpapers or a solid color.
Fills the screen by default.

Usage:
  setwp [--fit | --stretch | --center | --tile] [--color=<hex>] <wallpaper>
  setwp (--interval=<s> | --login | --wake) [--random] [--fit | --stretch | --center | --tile] [--color=<hex>] <directory>
  setwp --color=<hex>
  setwp --help | --version

Options:
  -f --fit          Fit wallpaper to screen.
  -s --stretch      Stretch wallpaper to fill screen.
  -c --center       Center wallpaper, scaling it down if it is too large.
  -t --tile         Tile wallpaper.
  -C --color=<hex>  Color to fill the screen with, as an RGB hex code.
  -h --help         Show this help message.
  -v --version      Show version information.

Directory options:
  -i --interval=<s>  Interval at which to change wallpaper, in seconds.
  -l --login         Change wallpaper when logging in.
  -w --wake          Change wallpaper when waking from sleep.
  -r --random        Randomize wallpaper selection.
~~~

## Limitations
- When setting wallpaper to a directory, spaces and desktops are not synced. Having them in sync is outside the scope of this project, as setwp changes wallpaper settings for all spaces and lets the OS change the actual wallpaper.
- When setting wallpaper to a directory, the search for images is not recursive. Only the images directly in the given directory will be used, not those in subdirectories.
