package dotenv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Run checks the current environment (i.e., dev|staging|production).
// If not (staging|production), Run looks for the passed filename and sets
// the environment variables appropriately, if the file exists.
func Run(filename string) error {
	isProductionOrStaging, err := regexp.MatchString("(production|staging)", os.Getenv("ENVIRONMENT"))
	if err != nil {
		return fmt.Errorf("error: identifying environment: %s", err)
	}
	if isProductionOrStaging {
		return nil
	}

	// Open file.
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error: opening file %s: %s", filename, err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Printf("error: closing file %s in dotenv.Run: %s", filename, err)
		}
	}()

	// Read and process file.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			keyValuePair := strings.SplitN(line, "=", 2)
			err = os.Setenv(keyValuePair[0], keyValuePair[1])
			if err != nil {
				return fmt.Errorf("error: setting env var pair k=%s, v=%s: %s", keyValuePair[0], keyValuePair[1], err)
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error: processing file %s: %s", filename, err)
	}

	return nil
}
