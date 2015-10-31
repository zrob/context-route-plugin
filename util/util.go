package util

import (
	"fmt"
	"os"
)

func FreakOut(err error) {
	if err != nil {
		fmt.Println("An unexpected error occurred: ", err.Error())
		os.Exit(1)
	}

}
