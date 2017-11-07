package leetcode

import (
	"fmt"
	"strings"
)

// path = "/home/", => "/home"
// path = "/a/./b/../../c/", => "/c"
// "/home//foo/"=> "/home/foo"

type pathStack []string

func (s pathStack) Push(v string) pathStack {
	return append(s, v)
}

func (s pathStack) Pop() (pathStack, string) {
	l := len(s)

	if l == 0 {
		return s, ""
	}

	v := s[l-1]
	s = s[:l-1]

	return s, v
}

func SimplifyPath(path string) string {
	path = strings.Replace(path, "//", "/", -1)
	pathParts := strings.Split(path, "/")
	// var absPathParts []string
	ps := pathStack{}

	for _, v := range pathParts {
		if v == "." || v == "" {
			continue
		} else if v == ".." && ps != nil {
			// absPathParts = absPathParts[:len(absPathParts)-1]
			ps, _ = ps.Pop()
			continue
		} else {
			// absPathParts = append(absPathParts, v)
			ps = ps.Push(v)
		}
	}

	fmt.Printf("%+v\n", ps)

	if ps == nil {
		return "/"
	}

	var absPath string = ""
	for {
		nps, p := ps.Pop()
		if p == "" {
			break
		}

		absPath += "/" + p
		ps = nps
	}

	// absPath = getPath(absPath)

	return absPath
}

func getPath(input string) string {
	if input == "/" || input == "" {
		return "/"
	}

	len := len(input)

	if len <= 1 {
		return input
	}

	if input[len-1:] == "/" {
		return input[:len-1]
	}

	return input
}
