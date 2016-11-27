package folder

import "strings"

var replaces = []struct{ a, b string }{
	{":", "-COLON-"},
	{"/", "-SLASH-"},
	{"@", "-AT-"},
}

// FromURL converts the given URL to a folder name
func FromURL(url string) string {
	result := url
	for _, replace := range replaces {
		result = strings.Replace(result, replace.a, replace.b, -1)
	}
	return result
}

// ToURL converts the given folder to an URL
func ToURL(folder string) string {
	result := folder
	for _, replace := range replaces {
		result = strings.Replace(result, replace.b, replace.a, -1)
	}
	return result
}
