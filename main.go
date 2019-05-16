package main

import (
	"os"
	"fmt"
	"log"
	"regexp"
	"path/filepath"
)

func main() {
	err := filepath.Walk(".",
    	func(path string, info os.FileInfo, err error) error {
    		if err != nil {
				return err
			}
			if info.IsDir() {
				fmt.Println("Dir: ", info.Name())
			} else if r, err := regexp.MatchString(".go", info.Name()); err == nil && r {
				fmt.Println("Go file: "+ path)
			} else {
				fmt.Println("Unknown: "+ path)
			}
    		return nil
		})
	if err != nil {
    	log.Println(err)
	}
}
