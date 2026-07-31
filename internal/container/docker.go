package container

import (
	"log"

	"github.com/moby/moby/client"
)

func container() { // not permantly func(for now atleast im just experimenting)
	cli, err := client.New(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	if cli == nil { //temp because i CANT STAND error underlines

	}

}
