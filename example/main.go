package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/xtdlib/ignore"
)

func main() {
	fmt.Print("wer")
	rules, err := ignore.ParseFile(".ignore")
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(".",
		func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if rules.Ignore(path, fi) {
			log.Println("ignore", path)
		} else {
			log.Println("no ignore", path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	// ignore.New()
}
