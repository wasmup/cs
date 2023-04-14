package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	// args := []string{""}

	if len(args) != 4 {
		fmt.Println("Usage: go run main.go hashfile1 hashfile2 output_dir")
		os.Exit(1)
	}
	outputDir := args[3]

	left := args[1]
	a, err := parseFile(left)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	right := args[2]
	b, err := parseFile(right)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	matching, nonMatching := findMatchingAndNonMatching(a, b)
	matchingFile := filepath.Join(outputDir, "matching.html")
	nonMatchingFile := filepath.Join(outputDir, "not_matching.html")
	cpFile := filepath.Join(outputDir, "run.sh")

	if err := writeHTML(matchingFile, matching, "Matching Rows", left, right); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := writeHTML(nonMatchingFile, nonMatching, "Non-Matching Rows", left, right); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := writeCopyRightToLeft(cpFile, nonMatching, left, right); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writeCopyRightToLeft(outputFile string, rows map[string][2][]string, left, right string) error {
	left = filepath.Base(left)
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, v := range rows {
		if len(v[0]) == 0 && len(v[1]) == 1 {
			s := filepath.Join(left, v[1][0])
			f.WriteString("cp " + s + " " + v[1][0])
		}
	}
	return nil
}

func findMatchingAndNonMatching(a, b map[string][]string) (map[string][2][]string, map[string][2][]string) {
	matching := make(map[string][2][]string)
	nonMatching := make(map[string][2][]string)

	for k, va := range a {
		if vb, ok := b[k]; ok {
			if !equal(va, vb) {
				matching[k] = [2][]string{va, vb}
			}
		} else {
			nonMatching[k] = [2][]string{va, {}}
		}
	}

	for k, vb := range b {
		if _, ok := a[k]; !ok {
			nonMatching[k] = [2][]string{{}, vb}
		}
	}

	return matching, nonMatching
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	values := make(map[string]int)

	for _, value := range a {
		values[value]++
	}

	for _, value := range b {
		if values[value] == 0 {
			return false
		}
		values[value]--
	}

	return true
}

func findMatching(a, b map[string][]string) map[string][2][]string {
	matching := make(map[string][2][]string)

	for keyA, valueA := range a {
		if valueB, ok := b[keyA]; ok {
			matching[keyA] = [2][]string{valueA, valueB}
		}
	}

	for key, values := range matching {
		if len(values[0]) == 1 && len(values[1]) == 1 && values[0][0] == values[1][0] {
			delete(matching, key)
		}
	}

	return matching
}

func writeHTML(outputFile string, rows map[string][2][]string, title, left, right string) error {
	html := genHTML(rows, left, right)

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("<html><head><title>" + title + "</title>")
	f.WriteString("<style>table { border-collapse: collapse; width: 100%; }")
	f.WriteString("td, th { border: 1px solid #ddd; padding: 8px; text-align: left; }")
	f.WriteString("tr:nth-child(even) { background-color: #84b582; }</style>")
	f.WriteString("</head><body>")
	f.WriteString(html)
	f.WriteString("</body></html>")

	return nil
}

func genHTML(matching map[string][2][]string, left, right string) string {
	var buf bytes.Buffer
	buf.WriteString("<table>\n")
	buf.WriteString(fmt.Sprintf("<tr><th>%q</th><th>%q</th></tr>\n", left, right))

	for _, values := range matching {
		buf.WriteString("<tr>\n")
		// buf.WriteString("<td><h3>" + key + "</h3>")
		buf.WriteString("<td>")
		buf.WriteString("<ul>")
		for _, value := range values[0] {
			buf.WriteString("<li>" + value + "</li>")
		}
		buf.WriteString("</ul></td>\n")

		// buf.WriteString("<td><h3>" + key + "</h3>")
		buf.WriteString("<td>")
		buf.WriteString("<ul>")
		for _, value := range values[1] {
			buf.WriteString("<li>" + value + "</li>")
		}
		buf.WriteString("</ul></td>\n")

		buf.WriteString("</tr>\n")
	}

	buf.WriteString("</table>\n")

	return buf.String()
}

func parseFile(filename string) (map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %s", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make(map[string][]string)
	var hash string
	var values []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result[hash] = values
			hash = ""
			values = nil
		} else if hash == "" {
			hash = line
			values = make([]string, 0)
		} else {
			values = append(values, line)
		}
	}

	if hash != "" {
		result[hash] = values
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %s", filename, err)
	}

	return result, nil
}
