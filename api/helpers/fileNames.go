package helpers

import (
	"flag"
	"github.com/kardianos/osext"
	"log"
	"os"
	"path/filepath"
)

var (
	// Initialization of the working directory. Needed to load asset files.
	filePath = determineWorkingDirectory()
	// File names for the HTTPS certificate
	certFilename = filepath.Join(filePath, "cert.pem")
	keyFilename  = filepath.Join(filePath, "key.pem")
)

func determineWorkingDirectory() string {
	var customPath string
	// Check if a custom path has been provided by the user.
	flag.StringVar(&customPath, "custom-path", "", "Specify a custom path to the asset files. This needs to be an absolute path.")
	flag.Parse()
	// Get the absolute path this executable is located in.
	executablePath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal("Error: Couldn't determine working directory: " + err.Error())
	}
	// Set the working directory to the path the executable is located in.
	os.Chdir(executablePath)
	// Return the user-specified path. Empty string if no path was provided.
	return customPath
}
