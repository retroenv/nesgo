package compiler

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

var errNoModules = errors.New("no valid go.mod found")

// currentPackage gets the current package based on the content of a go.mod
// file in the given directory or any of its parent directories.
// It returns the package name and the directory containing the go.mod file.
func currentPackage() (pack string, directory string, err error) {
	parent, err := os.Getwd()
	if err != nil {
		// A nonexistent working directory can't be in a module.
		return "", "", fmt.Errorf("getting working directory: %w", err)
	}

	var info os.FileInfo
	for {
		info, err = os.Stat(filepath.Join(parent, "go.mod"))
		if err == nil && !info.IsDir() {
			break
		}
		d := filepath.Dir(parent)
		if len(d) >= len(parent) {
			return "", "", errNoModules // reached top of file system, no go.mod
		}
		parent = d
	}

	full := path.Join(parent, info.Name())
	data, err := os.ReadFile(full)
	if err != nil {
		return "", "", fmt.Errorf("reading file '%s': %w", full, err)
	}

	reader := bytes.NewReader(data)
	buf := bufio.NewReader(reader)

	moduleText := []byte("module ")
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", "", err
		}
		if !bytes.HasPrefix(line, moduleText) {
			continue
		}
		line = bytes.TrimPrefix(line, moduleText)
		list := bytes.Split(line, []byte("\n"))
		if len(list) > 0 {
			return string(list[0]), parent, nil
		}
	}
	return "", "", errNoModules
}
