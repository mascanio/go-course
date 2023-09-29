package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"unicode"
	"unicode/utf8"
)

type RegexNamed struct {
	named_map map[string]int
	*regexp.Regexp
}

const error_string = string(unicode.ReplacementChar)

func parse_bytes(in []byte, until rune) ([]string, error) {
	re_named_match := regexp.MustCompile(`^\?P\<(.*?)\>`)
	re_no_capture_match := regexp.MustCompile(`^\?:`)
	if len(in) == 0 {
		return nil, nil
	}
	nextrun, runlen := utf8.DecodeRune(in)
	in = in[runlen:]
	if nextrun == until {
		return parse_bytes(in, 0)
	}
	var prepend_return []string
	var until_recursive rune
	switch nextrun {
	case '\\':
		if len(in) == 0 {
			return nil, errors.New("malformed - trailing \\")
		}
		// Skip scaped rune
		if nextrun, runlen := utf8.DecodeRune(in); nextrun != utf8.RuneError {
			in = in[runlen:]
		} else {
			return nil, errors.New("incorrect rune")
		}
		return parse_bytes(in, until_recursive)
	case '(':
		until_recursive = ')'
		if m := re_named_match.FindSubmatchIndex(in); m != nil {
			// Named pattern ?P<name>, return name
			prepend_return = []string{string(in[m[2]:m[3]])}
			in = in[m[3]+1:]
		} else if m := re_no_capture_match.FindSubmatchIndex(in); m != nil {
			// no capture, skip ?: part
			in = in[m[1]+1:]
		} else {
			prepend_return = []string{error_string}
		}
		if rec, err := parse_bytes(in, until_recursive); err == nil {
			return append(prepend_return, rec...), nil
		} else {
			return nil, err
		}
	default:
		return parse_bytes(in, until_recursive)
	}
}

func parse(in string) ([]string, error) {
	return parse_bytes([]byte(in), 0)
}

func buildMap(namedMatches []string) map[string]int {
	r := make(map[string]int)
	for i, name := range namedMatches {
		if name != error_string {
			r[name] = i + 1
		}
	}
	return r
}

func New(re *regexp.Regexp) RegexNamed {
	if parsed, err := parse(re.String()); err == nil {
		return RegexNamed{buildMap(parsed), re}
	} else {
		log.Fatalf("Invalid regex %v", err)
	}
	return RegexNamed{nil, nil}
}

func (re *RegexNamed) FindStringNamed(s, name string) string {
	match := re.FindStringSubmatch(s)
	return match[re.named_map[name]]
}

func main() {
	re := New(regexp.MustCompile(`^(?P<pathname>/?(?:[a-zA-Z0-9\.]*/)*)(?P<filename>.+)$`))
	files := []string{
		"/home/clinares/latex/go.tex", "tmp/regexp.go", "../org/index.org",
	}
	for _, file := range files {
		fmt.Println(re.FindStringNamed(file, "pathname"))
		fmt.Println(re.FindStringNamed(file, "filename"))
	}
}
