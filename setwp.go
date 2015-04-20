package main

import (
	"database/sql"
	"fmt"
	"github.com/alexandrecormier/setwp/args"
	"github.com/alexandrecormier/setwp/pref"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os/exec"
	"os/user"
	"path/filepath"
)

type Args struct {
}

const (
	programName = "setwp"

	dbRelativePath = "Library/Application Support/Dock/desktoppicture.db"

	clearDBStatement = `
		delete from data;
		delete from preferences;
	`

	setPrefDBStatement = `
		insert into data
		select ?
		where not exists (
			select value
			from data
			where value = ?);
	 	insert into preferences
	 	select ?, data.ROWID, pictures.ROWID
	 	from pictures
	 	inner join data
	 	on data.value = ?;
	`
)

var (
	success = false
)

func main() {
	log.SetFlags(0)

	prefs, err := args.Parse()
	if err != nil {
		log.Fatalf("%s", err)
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

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("error updating database (%s)", err)
	}
	defer closeTx(tx)

	clearDB(tx)
	for _, p := range prefs {
		setPref(tx, p)
	}

	success = true

	if err := exec.Command("killall", "Dock").Run(); err != nil {
		log.Println("error applying wallpaper, it will be applied on your next login")
	}
}

func clearDB(tx *sql.Tx) {
	if _, err := tx.Exec("delete from data; delete from preferences;"); err != nil {
		log.Fatalf("error updating database (%s)", err)
	}
}

func setPref(tx *sql.Tx, p pref.Pref) {
	if _, err := tx.Exec(setPrefDBStatement, p.Value, p.Value, p.Key, p.Value); err != nil {
		log.Fatalf("error updating database (%s)", err)
	}
}

func closeTx(tx *sql.Tx) {
	if success {
		if err := tx.Commit(); err != nil {
			log.Fatalf("error updating database (%s)", err)
		}
	} else {
		if err := tx.Rollback(); err != nil {
			log.Fatalf("error aborting database changes (%s)", err)
		}
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
