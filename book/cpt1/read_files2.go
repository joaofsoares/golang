package cpt1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFilesOnly(filenames []string) {
	counts := make(map[string]int)

	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ReadFilesOnly: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
