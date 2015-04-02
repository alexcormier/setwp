package main

import (
	"database/sql"
	"fmt"
	"github.com/docopt/docopt-go"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os/exec"
	"os/user"
	"path/filepath"
)

const (
	programName = "setwp"

	dbRelativePath = "Library/Application Support/Dock/desktoppicture.db"
	dbStatements   = `
		delete from data;
		delete from preferences;
		insert into data values (?);
		insert into data values (?);
		insert into preferences select 1, 1, ROWID from pictures;
		insert into preferences select 2, 2, ROWID from pictures;
	`
	usage = `Sets wallpaper to <wallpaper>. Fills the screen by default.

Usage:
  %[1]s [--fit | --stretch | --center | --tile] <wallpaper>
  %[1]s --help | --version

Options:
  -f --fit      Fit wallpaper to screen.
  -s --stretch  Stretch wallpaper to fill screen.
  -c --center   Center wallpaper.
  -t --tile     Tile wallpaper.
  -h --help     Show this help message.
  -v --version  Show version information.

`
	version = "%s version 0.1.1-1"
)

var positions = [...]string{"--fit", "--stretch", "--center", "--tile"}

func main() {
	log.SetFlags(0)

	args, err := docopt.Parse(fmt.Sprintf(usage, programName), nil, true, fmt.Sprintf(version, programName), false)
	if err != nil {
		log.Fatalf("cannot parse arguments (%s)", err)
	}

	wallpaper := args["<wallpaper>"]

	position := 5
	for _, p := range positions {
		if args[p].(bool) {
			break
		}
		position--
	}

	home, err := homeDir()
	if err != nil {
		log.Fatalf("cannot open database", err)
	}
	dbPath := filepath.Join(home, dbRelativePath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("cannot open database (%s)", err)
	}
	defer db.Close()

	if _, err := db.Exec(dbStatements, wallpaper, position); err != nil {
		log.Fatalf("error updating database (%s)", err)
	}

	if err := exec.Command("killall", "Dock").Run(); err != nil {
		log.Println("error applying wallpaper, it will be applied on your next login")
	}
}

func homeDir() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("unknown current user")
	}
	if user.HomeDir == "" {
		return "", fmt.Errorf("unknown home directory")
	}
	return user.HomeDir, nil
}
