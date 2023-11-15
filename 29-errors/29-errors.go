package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
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

//

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func outImproved(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "Out string 1")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Out string 2")
	return err
}

func writeStringsToFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	for i := 0; i < 20; i++ {
		sw.writeln("s" + strconv.Itoa(i))
	}
	return err
}

const rows, cols = 9, 9

type Grid [rows][cols]int8

func (g *Grid) Set(row, col int, digit int8) error {
	if !inBounds(row, col) {
		return errors.New("Out of bounds")
	}
	g[row][col] = digit
	return nil
}

func inBounds(row, col int) bool {
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return false
	}
	return true
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

	writeStringsToFile("out2.txt")

	var g Grid
	gridErr := g.Set(10, 0, 5)
	if gridErr != nil {
		fmt.Printf("Error: %v\n", gridErr)
	}
}
