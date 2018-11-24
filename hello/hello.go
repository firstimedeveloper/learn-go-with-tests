package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		return "Hello, World"
	}

	return greetingPrefix(lang) + name

}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		return spanishPrefix
	case french:
		return frenchPrefix
	default:
		return helloPrefix
	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}

//https://splice.com/blog/contributing-open-source-git-repositories-go/
