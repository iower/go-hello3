package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func out(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Out string 1")
	if err != nil {
		f.Close()
		return err
	}

	_, err = fmt.Fprintln(f, "Out string 2")
	f.Close()
	return err
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.Size(), file.ModTime(), file.Mode())
	}

	// files2, err2 := ioutil.ReadDir("/etc/hosts")
	// fmt.Println(files2, err2)

	outErr := out("out.txt")
	if outErr != nil {
		fmt.Println("Out err:", outErr)
		os.Exit(1)
	}
}
