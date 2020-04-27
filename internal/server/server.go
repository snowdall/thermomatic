// This package handles the socket server.  Both setting up the
// socket listen but also handling the incoming data packets
package server

import (
  "net"
  "io"
  "fmt"
  "errors"
  "time"
  "strconv"
	"github.com/spin-org/thermomatic/internal/common"
	"github.com/spin-org/thermomatic/internal/imei"
)

var (
	ErrImeiTimeout  = errors.New("server: imei send timeout")
	ErrReadingTimeout = errors.New("server: data reading timeout")
)

// This is the handler that is called when a client connects
// Its takes the connection as a parameter
func handleConnection(con net.Conn) {
  common.Out("Have a valid connection")

  // Recover, close connection, and return in the event of a panic
  defer func() {
    recover()
    con.Close()
    return
  }()

	// Temporary storage for reading from socket. Nothing
  // should be more that 64 bytes nominally
	readBuf := make([]byte, 64)

  // Set the timeout for the initial IMEI call first
  con.SetReadDeadline(time.Now().Add(2 * time.Second))

  // This first read will bring in the IMEI.
  dataLen, err := con.Read(readBuf)
	if err != nil {
    common.Err(err)
    con.Close()
    return
  }

  // Check the IMEI for vailidity and return the string
  imei, err := imei.Decode(readBuf[0:15])
	if err != nil {
    common.Err(err)
    con.Close()
    return
  }

  // Check the length, if it's 0 then we've timed out and
  // need to close and break.
  if dataLen == 0 {
    common.Err(ErrImeiTimeout)
    con.Close()
    return
  }

  for {
    // Set the timeout for the data reading to 1 second
    con.SetReadDeadline(time.Now().Add(1 * time.Second))

		// This first read will bring in the IMEI.
		dataLen, err := con.Read(readBuf)

    // Check the length, if it's 0 then we've timed out and
    // need to close and break.
    if dataLen == 0 {
      common.Err(ErrReadingTimeout)
      con.Close()
      return
    }

		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client!")
				return
			}
		}

    fmt.Println(dataLen)
    fmt.Println(readBuf)
	}
  // First get the IMEI from the login message

  // Next start looping to read data, making sure that they
  // are coming every 2 seconds
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

