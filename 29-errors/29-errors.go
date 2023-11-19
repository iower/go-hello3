package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

func (g *Grid) Set(row, col int, digit int8) error {
	if !inBounds(row, col) {
		return ErrBounds
	}
	var digitErr error
	if !validDigit(digit) {
		return errors.New("Wrong digit")
	}
	if digitErr != nil {
		return digitErr
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

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

// implement error

type error2 interface {
	Error() string
}

// few errors
type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func (g *Grid) SetImproved(row, col int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, col) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][col] = digit
	return nil
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
		switch gridErr {
		case ErrBounds, ErrDigit:
			fmt.Printf("Known error: %v\n", gridErr)
		default:
			fmt.Println(gridErr)
		}
	}

	// type assertion
	err3 := g.SetImproved(10, 0, 15)
	if err3 != nil {
		if errs, ok := err3.(SudokuError); ok {
			fmt.Printf("%d errors(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}
}
