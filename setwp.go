package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/alexandrecormier/setwp/args"
	"github.com/alexandrecormier/setwp/pref"
	_ "github.com/mattn/go-sqlite3"
)

const (
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

func main() {
	// call another function to properly run defers
	os.Exit(run())
}

func run() int {
	log.SetFlags(0)

	prefs, err := args.Parse()
	if err != nil {
		return handleArgumentError(err)
	}

	home, err := homeDir()
	if err != nil {
		return handleDbReadError(err)
	}
	dbPath := filepath.Join(home, dbRelativePath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return handleDbReadError(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return handleDbWriteError(err)
	}
	success := false
	defer closeTx(tx, &success)

	if err := clearDB(tx); err != nil {
		return handleDbWriteError(err)
	}
	for key, value := range prefs {
		if err := setPref(tx, pref.Pref{Key: key, Value: value}); err != nil {
			return handleDbWriteError(err)
		}
	}

	success = true

	if err := exec.Command("killall", "Dock").Run(); err != nil {
		log.Println("error applying wallpaper, it will be applied on your next login")
	}

	return 0
}

// Clears the wallpaper preferences database.
func clearDB(tx *sql.Tx) error {
	_, err := tx.Exec(clearDBStatement)
	return err
}

// Sets a preference in the database.
func setPref(tx *sql.Tx, p pref.Pref) error {
	_, err := tx.Exec(setPrefDBStatement, p.Value, p.Value, p.Key, p.Value)
	return err
}

// Commits or rollbacks the transaction depending on success.
func closeTx(tx *sql.Tx, success *bool) {
	if *success {
		if err := tx.Commit(); err != nil {
			os.Exit(handleDbWriteError(err))
		}
	} else {
		if err := tx.Rollback(); err != nil {
			os.Exit(handleDbRollbackError(err))
		}
	}
}

// Gets the current user's home directory.
func homeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", errors.New("unknown current user")
	}
	if currentUser.HomeDir == "" {
		return "", errors.New("unknown home directory")
	}
	return currentUser.HomeDir, nil
}
