package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	cmd "tomdeneire.github.io/tiro/cli/cmd"
)

var buildTime string
var goVersion string
var buildHost string

func main() {
	if len(os.Args) > 2 && os.Args[1] == "arg" {
		data := make([]byte, 0)
		mode := os.Args[2]
		var err error
		switch mode {
		case "stdin":
			data, err = io.ReadAll(os.Stdin)
			if err != nil {
				data = nil
				break
			}
		case "file":
			if len(os.Args) < 4 {
				data = nil
				break
			}
			fname := os.Args[3]
			var err error
			file, err := os.Open(fname)
			if err != nil {
				data = nil
				break
			}
			defer file.Close()
			data, err = io.ReadAll(file)
			if err != nil {
				data = nil
				break
			}
		case "json":
			if len(os.Args) < 4 {
				data = nil
				break
			}
			data = []byte(os.Args[3])

		case "url":
			if len(os.Args) < 4 {
				data = nil
				break
			}
			jarg := os.Args[3]
			resp, err := http.Get(jarg)
			if err != nil {
				break
			}
			defer resp.Body.Close()
			data, err = io.ReadAll(resp.Body)
			if err != nil {
				data = nil
				break
			}
		}

		args := make([]string, 0)
		args = append(args, os.Args[0])
		if data != nil {
			data = bytes.TrimSpace(data)
			sdata := string(data)
			if !strings.HasPrefix(sdata, "[") {
				lines := strings.SplitN(sdata, "\n", -1)
				ok := false
				for _, line := range lines {
					line = strings.TrimSpace(line)
					if line == "" {
						continue
					}
					if !ok && line != "tiro" {
						args = nil
						break
					}
					if !ok {
						ok = true
						continue
					}
					args = append(args, line)
				}
			} else {
				argums := make([]string, 0)
				err := json.Unmarshal(data, &argums)
				if err != nil {
					args = nil
				} else {
					if len(argums) == 0 || argums[0] != "tiro" {
						args = nil
					} else {
						args = append(args, argums[1:]...)
					}
				}
			}

		}
		if args != nil {
			os.Args = args
		}
	}
	cmd.Execute(buildTime, goVersion, buildHost, os.Args)
}
