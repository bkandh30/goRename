package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "Real or dry run")
	flag.Parse()

	walkDir := "sample"
	toRename := make(map[string][]string)
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		curDir := filepath.Dir(path)

		if m, err := match(info.Name()); err == nil {
			key := filepath.Join(curDir, fmt.Sprintf("%s.%s", m.base, m.ext))
			toRename[key] = append(toRename[key], info.Name())
		}

		return nil
	})

	for key, files := range toRename {
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)

		for i, filename := range files {
			res, _ := match(filename)
			newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, (i + 1), n, res.ext)
			oldPath := filepath.Join(dir, filename)
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("mv %s => %s\n", oldPath, newPath)

			if !dry {
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println("Error renaming:", oldPath, newPath, err.Error())
				}
			}
		}
	}
}

type matchResult struct {
	base  string
	index int
	ext   string
}

func match(fileName string) (*matchResult, error) {
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return nil, fmt.Errorf("%s did not match our pattern", fileName)
	}

	caser := cases.Title(language.English)
	return &matchResult{caser.String(name), number, ext}, nil
}
