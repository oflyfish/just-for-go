package main

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
)

func main() {
	fmt.Println("just for go")

	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig("https://github.com/oflyfish/just-for-go", config)
}
