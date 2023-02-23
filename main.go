package main

import (
	"github.com/mati/latencia/api"
	"github.com/mati/latencia/cmd"
)

func main() {
	cmd.StartTasks()
	api.Router()
}
