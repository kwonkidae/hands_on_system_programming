package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

var cmdFunc func(w io.Writer, args ...string) (exit bool)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	fmt.Fprint(w, "Some welcome message\n")
	for {
		s.Scan()
		args := strings.Split(string(s.Bytes()), " ")
		cmd := args[0]
		args = args[1:]

		switch cmd {
		case "exit":
			cmdFunc = exitCmd
		case "shuffle":
			cmdFunc = shuffle
		case "print":
			cmdFunc = print
		}

		if cmdFunc == nil {
			fmt.Fprintf(w, "%q not found\n", cmd)
		}
		if cmdFunc(w, args...) {
			return
		}
	}
}

func exitCmd(w io.Writer, args ...string) bool {
	fmt.Fprintf(w, "Goodbye! :)")
	return true
}

func shuffle(w io.Writer, args ...string) bool {
	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
	})

	for i := range args {
		if i > 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprintf(w, "%s", args[i])
	}
	fmt.Fprintln(w)
	return false
}

func print(w io.Writer, args ...string) bool {
	if len(args) != 1 {
		fmt.Fprintln(w, "Please specify one file!")
		return false
	}
	f, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(w, "Cannot open %s: %s\n", args[0], err)
	}
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		fmt.Fprintf(w, "Cannot print %s: %s\n", args[0], err)
	}
	fmt.Fprintln(w)
	return false
}
