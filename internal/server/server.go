// This package handles the socket server.  Both setting up the
// socket listen but also handling the incoming data packets
package server

import (
  "net"
  "strconv"
	"github.com/spin-org/thermomatic/internal/common"
)

// This is the handler that is called when a client connects
// Its takes the connection as a parameter
func handleConnection(con net.Conn) {
  common.Out("Have a valid connection")
}

// This function is called in the main to start the socket server.
// It takes a network port number as a paramter.
func StartServer(port int) {
  common.Out("StartServer ... ")

  // Set up the socket and start listening on it.
  link, err := net.Listen("tcp", ":" + strconv.Itoa(port))

  // If there is a problem with the socket open, the record it and return
  if err != nil {
    common.Err(err)
    return
  }

  // Make sure the socket is closed in the event things shutdown
  defer link.Close()

  // Wait for client connections here.  If a client does connect
  // spin the connection off into a handler routine to be processed.
  for {
    con, err := link.Accept()
    if err != nil {
      common.Err(err)
      return
    }
    go handleConnection(con)
  }
}

