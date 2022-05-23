package reader

import (
	"bufio"
	"strings"
)

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func ParseStatusCode(s string) string {
	parsed := strings.Split(s, ":")
	return parsed[1]
}

func CheckValid(url, status string) bool {
	return status == "200" && !strings.HasSuffix(url, ".GIF")
}
