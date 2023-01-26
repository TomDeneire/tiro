package util

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

// General functions

func ReadFile(path string) (string, error) {
	result := ""
	file, _ := os.Open(path)
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return result, err
	}
	result = string(b)
	return result, nil
}

// Test functions

// Compare result and expected for tests
func Check(result string, expected string, t *testing.T) {
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}

// XML functions

// CSV functions

func ParseCSV(input io.Reader) ([][]string, error) {
	var data [][]string

	csvReader := csv.NewReader(input)
	data, err := csvReader.ReadAll()
	if err != nil {
		return data, fmt.Errorf("cannot read csv: %v", err)
	}
	return data, nil
}

func Csvconvert(data *[][]string, mode string) ([]byte, error) {

	var convertedData []byte

	headers := (*data)[0]
	rows := (*data)[1:]

	switch mode {

	case "json":
		result := make([]map[string]string, len(rows))
		for i, row := range rows {
			result[i] = map[string]string{}
			for j, cell := range row {
				info := result[i]
				info[headers[j]] = cell
				result[i] = info
			}
		}
		// encoding is done here
		out, err := json.Marshal(result)
		if err != nil {
			return convertedData, fmt.Errorf("cannot marshal json: %v", err)
		}
		convertedData = out

	case "xml":
		result := `<?xml version="1.0" encoding="UTF-8"?><data>`
		xmlHeaders := make([]string, len(headers))
		for i, header := range headers {
			var b bytes.Buffer
			// explicit encoding needed
			err := xml.EscapeText(&b, []byte(header))
			if err != nil {
				return convertedData, fmt.Errorf("cannot escape xml: %v", err)
			}
			xmlHeaders[i] = b.String()
		}
		for i, row := range rows {
			result = result + `<record nr="` + strconv.Itoa(i+1) + `">`
			for j, cell := range row {
				var b bytes.Buffer
				err := xml.EscapeText(&b, []byte(cell))
				if err != nil {
					return convertedData, fmt.Errorf("cannot escape xml: %v", err)
				}
				text := b.String()
				result = result + `<column label="` + xmlHeaders[j] + `">` + text + `</column>`
			}
			result = result + "</record>"
		}
		result = result + "</data>"
		convertedData = []byte(result)
	}

	return convertedData, nil
}

// Helper functions

func Save2File(data *[]byte, filepath string) error {
	err := ioutil.WriteFile(filepath, *data, 0664)
	if err != nil {
		return fmt.Errorf("cannot write to file: %v", err)
	}
	return nil
}

func HandleResult(data *[]byte, verbose bool, target string) error {
	if verbose {
		fmt.Println(string(*data))
	}
	if target != "" {
		err := Save2File(data, target)
		if err != nil {
			return fmt.Errorf("cannot save to file: %v", err)
		}
	}

	return nil

}
