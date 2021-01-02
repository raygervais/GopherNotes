package main

import (
	"fmt"
	"os"

	"github.com/raygervais/gophernotes/pkg/app"
)

func main() {
	code, message := app.Application()
	fmt.Println(message)
	os.Exit(code)
}
