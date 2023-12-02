package main

import "fmt"

const (
	// languages
	klingon = "Klingon"
	romulan = "Romulan"

	// prefixes
	defaultPrefix = "Hello, "
	klingonPrefix = "nuqneH, "
	romulanPrefix = "jolan'tru, "

	// suffix
	defaultSuffix = "!!!"
)

func Hello(name, language string) string {
	// if name == "" {
	if len(name) == 0 {
		name = "world"
	}

	return greetingPrefix(language) + name + defaultSuffix
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case klingon:
		prefix = klingonPrefix
	case romulan:
		prefix = romulanPrefix
	default:
		prefix = defaultPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Eve", ""))
}
