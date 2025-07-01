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

// вам надо написать более быструю оптимальную этой функции
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
	user := &User{}
	fmt.Fprintln(out, "found users:")

	for fileScanner.Scan() {

		lineBytes := fileScanner.Bytes()

		err := json.Unmarshal(lineBytes, user)
		if err != nil {
			panic(err)
		}

		i++

		browsers := user.Browsers

		isAndroid := false
		isMSIE := false

		for _, browserRaw := range browsers {

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
