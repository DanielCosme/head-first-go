package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	defer reportPanic()

	c := 0
	err := scanDirectory(os.Args[1], &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files", c)
}

func scanDirectory(path string, count *int) error {
	fmt.Println(path)
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			//fmt.Println("Directory: " + file.Name())
			defer scanDirectory(filePath, count)
		} else {
			fmt.Println(" ----- File: " + file.Name())
			//fmt.Println(filePath)
			*count++
		}
	}
	return err
}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}
