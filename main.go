package main

import (
	"command-line-application/command-line-application/app"
	"log"
	"os"
)

func main() {
	application := app.Generate()
	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}

}
