package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	inDir := flag.String("in", "/Users/mmoln/Desktop/sql", "Directory to read files from")
	outDir := flag.String("out", filepath.Join(*inDir, "out"), "Directory to read files from")
	flag.Parse()
	entries, err := os.ReadDir(*inDir)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.Mkdir(*outDir, os.ModeDir)
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}
	outFile, err := os.Create(filepath.Join(*outDir, "full.sql"))
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}
	defer outFile.Close()
	for i := range entries {
		if !entries[i].IsDir() {
			fileContent, err := os.ReadFile(filepath.Join(*inDir, entries[i].Name()))
			if err != nil {
				fmt.Printf("COUND NOT READ FILE %s: %v \n", entries[i].Name(), err)
			}

			fmt.Printf("Appending %s \n", entries[i].Name())
			_, err = outFile.Write(fileContent)
			if err != nil {
				fmt.Printf("COUND NOT WRITE TO FILE %v \n", err)
			}
		}
	}
}
