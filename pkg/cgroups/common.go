package cgroups

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

// ParseFileKeyValue parses a space-separated "name value" kind of cgroup file
// parameter and returns its key as a string, and its value as uint64. For example,
// "io_service_bytes 1234" will be returned as "io_service_bytes"(key), 1234(value).
func ParseFileKeyValue(dir, file string) (map[string]uint64, error) {
	p := path.Join(dir, file)
	if p == "" {
		return nil, fmt.Errorf("no path specified for %s", p)
	}
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var maps = map[string]uint64{}
	content := bufio.NewReader(f)
	for {
		line, _, err := content.ReadLine()
		if err == io.EOF {
			break
		}
		key, value, err := ParseLineKeyValue(string(line))
		if err != nil {
			return nil, err
		}
		maps[key] = value

	}
	return maps, nil
}

// ParseFileKeyValue parses a space-separated "name value" kind of cgroup
// parameter and returns its key as a string, and its value as uint64.
func ParseLineKeyValue(line string) (string, uint64, error) {
	items := strings.Split(string(line), " ")
	if len(items) != 2 {
		return "", 0, fmt.Errorf("line %q is not in key value format", line)

	}
	value, err := strconv.ParseUint(items[1], 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("")
	}
	return items[0], value, nil
}
