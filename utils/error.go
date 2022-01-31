package utils

import (
	"fmt"
)

func Assert(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
