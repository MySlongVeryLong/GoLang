package crack

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func CrackMD5(targetHash string, wordlistPath string) error {
	file, err := os.Open(wordlistPath)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		password := strings.TrimSpace(line)
		if password != "" {
			hashBytes := md5.Sum([]byte(password))
			checksum := hex.EncodeToString(hashBytes[:])

			if checksum == targetHash {
				fmt.Printf("PASSWORD FOUND (MD5): %s\n", password)
				return nil
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading file: %w", err)
		}
	}

	fmt.Println("Password not found in wordlist.")
	return nil
}
