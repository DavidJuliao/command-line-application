package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search ips and names of servers in web"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search ips of web hosts",
			Flags: flags,
			Action: func(c *cli.Context) {
				host := c.String("host")
				ips, erro := net.LookupIP(host)
				if erro != nil {
					log.Fatal(erro)
				}

				for _, ip := range ips {
					fmt.Println(ip)
				}
			},
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome do servidor",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, occursError := net.LookupNS(host)
	if occursError != nil {
		log.Fatal(occursError)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
