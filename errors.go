package main

import (
	"log"
)

func handleArgumentError(err error) int {
	log.Printf("%s", err)
	return 1
}

func handleDbReadError(err error) int {
	log.Printf("error opening database (%s)", err)
	return 2
}

func handleDbWriteError(err error) int {
	log.Printf("error updating database (%s)", err)
	return 3
}

func handleDbRollbackError(err error) int {
	log.Printf("error aborting database changes (%s)", err)
	return 4
}
