package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type User struct {
	Browsers []string `json:"browsers"`
	Email    string   `json:"email"`
	Name     string   `json:"name"`
}

func FastSearch(out io.Writer) {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	defer file.Close()

	seenBrowsers := map[string]bool{}
	uniqueBrowsers := 0

	i := -1
	//user := &User{}
	user := new(User)
	fmt.Fprintln(out, "found users:")

	lineBytes := make([]byte, 0, 1024)

	for fileScanner.Scan() {

		lineBytes = fileScanner.Bytes()

		err := json.Unmarshal(lineBytes, user)
		if err != nil {
			panic(err)
		}

		i++

		isAndroid := false
		isMSIE := false

		for _, browserRaw := range user.Browsers {

			browser := browserRaw

			if strings.Contains(browser, "Android") {
				isAndroid = true

			} else if strings.Contains(browser, "MSIE") {
				isMSIE = true

			} else {
				continue
			}

			_, seen := seenBrowsers[browser]
			if !seen {
				seenBrowsers[browser] = true
				uniqueBrowsers++
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		fmt.Fprintf(out, "[%d] %s <%s>\n", i, user.Name, strings.Replace(user.Email, "@", " [at] ", -1))
	}

	fmt.Fprintln(out, "")

	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
