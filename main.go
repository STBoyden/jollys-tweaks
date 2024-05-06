package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/mholt/archiver"
)

//go:embed INIs/*
var inisDirectoryFS embed.FS

const inisDirectory string = "INIs"

func main() {
	var outDirectoryPath string
	flag.StringVar(&outDirectoryPath, "out", "out", "")
	flag.Parse()

	err := os.MkdirAll("staging/Data/NVSE/Plugins", 0o744)
	if err != nil {
		panic("could not create staging directory:" + err.Error())
	}

	err = os.MkdirAll(outDirectoryPath, 0744)
	if err != nil {
		panic("could not create out directory:" + err.Error())
	}

	files, err := inisDirectoryFS.ReadDir(inisDirectory)
	if err != nil {
		panic("could not read INIs directory:" + err.Error())
	}

	for _, f := range files {
		sourceFilePath := inisDirectory + string(os.PathSeparator) + f.Name()
		destinationFilePath := fmt.Sprintf("staging/Data/NVSE/Plugins/%s", f.Name())

		src, err := inisDirectoryFS.Open(sourceFilePath)
		if err != nil {
			panic(fmt.Sprintf("could not read file at \"%s\": %s", sourceFilePath, err.Error()))
		}

		dst, err := os.Create(destinationFilePath)
		if err != nil {
			panic(fmt.Sprintf("could not create file at \"%s\":%s", destinationFilePath, err.Error()))
		}

		_, err = io.Copy(dst, src)
		if err != nil {
			panic(fmt.Sprintf("could not copy file data from \"%s\" to \"%s\": %s", sourceFilePath, destinationFilePath, err.Error()))
		}

		err = dst.Sync()
		if err != nil {
			panic(fmt.Sprintf("could not sync file to disk: %s", err.Error()))
		}

		err = src.Close()
		if err != nil {
			panic(fmt.Sprintf("could not close file at \"%s\":%s", sourceFilePath, err.Error()))
		}

		err = dst.Close()
		if err != nil {
			panic(fmt.Sprintf("could not close file at \"%s\":%s", destinationFilePath, err.Error()))
		}
	}

	err = archiver.Archive([]string{"staging/Data"}, outDirectoryPath+string(os.PathSeparator)+"Jolly's INI Tweaks.zip")
	if err != nil {
		panic(fmt.Sprintf("could not create zip file: %s", err.Error()))
	}
}
