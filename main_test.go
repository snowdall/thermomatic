package main

import (
	"testing"
	"github.com/spin-org/thermomatic/internal/client"
)

// This set of tests assumes the server is already running
// on the local machine.  With more time, a more self-contained
// testing framework would be setup to run the server and test it.

// This tests the nominal case running a single client against
// the server
func TestClientConnectionNominal(t *testing.T) {

  // It will send 10 data packets and then return 0.
  result := client.ClientConnect(client.TestImei, 200, 200, 10)

  if result != 0 {
    t.Error("Error in connection with the server")
  }
}

// This tests if the IMEI isn't sent within the first second
func TestClientConnectionImeiTimeout(t *testing.T) {

  // Set the IMEI timeout to 2 seconds to force the error
  result := client.ClientConnect(client.TestImei, 2000, 200, 10)

  // The result should be equal to 1 since the timeout isn't detected
  // client side until the timeout on the data reading write
  if result != 1 {
    t.Error("Error in connection with the server")
  }
}

// This tests if the Reading isn't sent within two seconds
func TestClientConnectionReadingTimeout(t *testing.T) {

  // Set the Reading timeout to 3 seconds to force the error
  result := client.ClientConnect(client.TestImei, 200, 3000, 10)

  // The result should be equal to 1 since the timeout isn't detected
  // client side until the timeout on the data reading write
  if result != 1 {
    t.Error("Error in connection with the server")
  }
}

