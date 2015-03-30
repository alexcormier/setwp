package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

const dbPathRelative = "Library/Application Support/Dock/desktoppicture.db"

var dbStatements = [...]string{
	"delete from data;",
	"delete from preferences;",
	"insert into data values (?);",
	"insert into preferences select 1, 1, ROWID from pictures;",
}

func main() {
	log.SetFlags(0)

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <wallpaper>", os.Args[0])
	}

	wallpaperPath := os.Args[1]

	home, err := homeDir()
	if err != nil {
		log.Fatalf("cannot open database", err)
	}
	dbPath := filepath.Join(home, dbPathRelative)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("cannot open database (%s)", err)
	}
	defer db.Close()

	for _, statement := range dbStatements {
		if _, err := db.Exec(statement, wallpaperPath); err != nil {
			log.Fatalf("error updating database (%s)", err)
		}
	}

	if err := exec.Command("killall", "Dock").Run(); err != nil {
		log.Println("error killing Dock, wallpaper will be applied on your next login")
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
