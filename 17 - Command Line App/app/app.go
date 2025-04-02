package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line App"
	app.Usage = "A simple command line app to get public IP"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Get Ip",
			Flags:  flags,
			Action: getIPByHost,
		},
		{
			Name:   "servidores",
			Usage:  "Get Server Name",
			Flags:  flags,
			Action: getServerNameByHost,
		},
	}

	return app
}

func getIPByHost(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServerNameByHost(c *cli.Context) {
	host := c.String("host")
	names, errors := net.LookupNS(host)
	if errors != nil {
		log.Fatal(errors)
	}

	for _, server := range names {
		fmt.Println(server.Host)
	}
}
