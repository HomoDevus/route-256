package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"encoding/json"
)

type Category struct {
	name string `json:"name"`
	id int `json:"id"`
	parent *int `json:"parent,omitempty"`
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var dataSetsAmount int

	fmt.Fscan(in, &dataSetsAmount)

	for dataSetIndex := 0; dataSetIndex < dataSetsAmount; dataSetIndex++ {
		var jsonLinesAmount int
		var categories []Category

		fmt.Fscan(in, &jsonLinesAmount)

		jsonData := scanJson(jsonLinesAmount)

		json.Unmarshal([]byte(jsonData), &categories)

		printJson(jsonData, out)
	}

	defer out.Flush()
}

func scanJson(linesAmount int) string {
	var jsonText string
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < linesAmount; i++ {
		scanner.Scan()
		jsonText += scanner.Text()
	}

	return jsonText
}

func printJson(data interface{}, out io.Writer) {
	// json, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprintf(out, "%s\n", data)
}