package main

import (
	"fmt"
	"hello/countdown"
	"os"
)

const spanish = "Spanish"
const french = "French"
const romanian = "Romanian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const romanianHelloPrefix = "Salut, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%v%v", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case romanian:
		prefix = romanianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	countdown.Countdown(os.Stdout, &countdown.DefaultSleeper{})
}
