package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mholt/archiver"
)

const inisDirectory string = "INIs"

func main() {
	err := os.MkdirAll("staging/Data/NVSE/Plugins", 0o744)
	if err != nil {
		panic("could not create staging directory:" + err.Error())
	}

	files, err := os.ReadDir(inisDirectory)
	if err != nil {
		panic("could not read INIs directory:" + err.Error())
	}

	for _, f := range files {
		sourceFilePath := fmt.Sprintf("%s%c%s", inisDirectory, os.PathSeparator, f.Name())
		destinationFilePath := fmt.Sprintf("staging/Data/NVSE/Plugins/%s", f.Name())

		src, err := os.Open(sourceFilePath)
		if err != nil {
			panic(fmt.Sprintf("could not read file at \"%s\": %s", sourceFilePath, err.Error()))
		}
		defer src.Close()

		dst, err := os.Create(destinationFilePath)
		if err != nil {
			panic(fmt.Sprintf("could not create file at \"%s\":%s", destinationFilePath, err.Error()))
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			panic(fmt.Sprintf("could not copy file data from \"%s\" to \"%s\": %s", sourceFilePath, destinationFilePath, err.Error()))
		}

		err = dst.Sync()
		if err != nil {
			panic(fmt.Sprintf("could not sync file to disk: %s", err.Error()))
		}
	}

	// f, err := os.Create("Jolly's INI Tweaks.zip")
	// if err != nil {
	// 	panic(fmt.Sprintf("could not create zip file: %s", err.Error()))
	// }

	err = archiver.Archive([]string{"staging/Data"}, "Jolly's INI Tweaks.zip")
	if err != nil {
		panic(fmt.Sprintf("could not create zip file: %s", err.Error()))
	}
}
