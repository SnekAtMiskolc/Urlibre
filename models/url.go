package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/teris-io/shortid"
)

// Used to represent a url!
type URL struct {
	ID      string `json:"id" bson:"id"`
	Url     string `json:"url" bson:"url"`
	Expires int64  `json:"exp" bson:"exp"`
}

// Used to represent the creation of a url!
type NewURL struct {
	Url     string `json:"url"`
	Expires int    `json:"exp"`
}

func (nu *NewURL) IntoURL() *URL {
	return &URL{
		ID:      shortid.MustGenerate(),
		Url:     nu.Url,
		Expires: addDays(nu.Expires),
	}
}

func addDays(days int) int64 {
	return time.Now().AddDate(0, 0, days).Unix()
}

func FilterByList(url string) (bool, error) {
	// Attempt to read filter.list
	readFile, err := os.Open("filter.list")
	if err != nil {
		return false, err
	}

	// Create a new Scanner
	fScanner := bufio.NewScanner(readFile)

	// Scan the lines
	fScanner.Split(bufio.ScanLines)

	// Iterate trough the lines
	for fScanner.Scan() {
		// If url contains the pattern on the line then return false, nil
		if strings.Contains(url, fScanner.Text()) && !strings.HasPrefix(fScanner.Text(), "#") {
			fmt.Println(fScanner.Text())
			return false, nil
		}
	}

	// If all is fine then return true, nil
	return true, nil

}
