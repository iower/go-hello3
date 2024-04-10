package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "@ab\tAaa\tBbb\t")
	fmt.Fprintln(w, "@cd\tCcc\tDdd\t")
	w.Flush()
}
