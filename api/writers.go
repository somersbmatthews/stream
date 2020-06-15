package api

import (
	"fmt"
	"github.com/somersbmatthews/stream/ogg"
	"os"
)

func OpusReader(fileName string, directory string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Errorf(err)
	}

}