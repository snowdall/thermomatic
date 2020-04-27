package main

import (
	"github.com/spin-org/thermomatic/internal/server"
	"github.com/spin-org/thermomatic/internal/common"
)

const PORT = 1337

func main() {
  common.Out("Starting thermomatic service")
  server.StartServer(PORT)
}

