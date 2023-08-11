package cpt1

import (
	"fmt"
)

func EchoParams(params []string) {

	if len(params) == 0 {
		fmt.Println("No param found")
	} else {
		output, sep := "", ""

		for _, param := range params {
			output += sep + param
			sep = " "
		}

		fmt.Println(output)
	}

}
