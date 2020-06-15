package api

import (
	"fmt"
	"github.com/mccoyst/ogg"
	"os"
)

func OpusReader(fileName string, directory string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Errorf(err)
	}

}