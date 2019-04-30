package main

import "fmt"

func main() {
	fmt.Println(fortuneCookie())
	sexyFortuneCookie(fortuneCookie)
}

func fortuneCookie() string {
	return "you will find happiness"
}

func sexyFortuneCookie(original func() string) {
	fmt.Println(original(), "in bed.")
}
