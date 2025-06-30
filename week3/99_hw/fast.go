package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

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

	fmt.Fprintln(out, "found users:")
	//users := make([]map[string]interface{}, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		//for _, line := range lines {
		user := make(map[string]interface{})
		// fmt.Printf("%v %v\n", err, line)
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			panic(err)
		}

		i++

		browsers, ok := user["browsers"].([]interface{})
		if !ok {
			continue
		}

		isAndroid := false
		isMSIE := false

		for _, browserRaw := range browsers {

			browser, ok := browserRaw.(string)
			if !ok {
				continue
			}
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

		email := strings.Replace(user["email"].(string), "@", " [at] ", -1)

		fmt.Fprintf(out, "[%d] %s <%s>\n", i, user["name"], email)
	}

	fmt.Fprintln(out, "")

	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
