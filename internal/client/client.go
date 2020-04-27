// TODO: document package.
package client

import (
  "time"
  "math/rand"
  "fmt"
  "net"
  "strconv"
)

const PORT = 1337
var TestImei = []byte { 4, 9, 0, 1, 5, 4, 2, 0, 3, 2, 3, 7, 5, 1, 8}

func ClientConnect(imei []byte, imei_timeout uint64, reading_timeout uint64) {

  CONNECT := "localhost:" + strconv.Itoa(PORT)
  con, err := net.Dial("tcp", CONNECT)
  if err != nil {
    fmt.Println(err)
    return
  }

  // For testing, put in a delay here for the IMEI sending.  This will
  // be used to test the server response.
  time.Sleep(time.Duration(imei_timeout) * time.Millisecond)

  // Set the timeout for all operations to 3 second
  con.SetWriteDeadline(time.Now().Add(3 * time.Second))

  dataLen, err := con.Write(imei)
  if err != nil {
    fmt.Println(err)
    return
  }

  // This is the loop for sending data to the server
  for {

    // Set the timeout for all operations to 3 second
    con.SetWriteDeadline(time.Now().Add(3 * time.Second))

    dataLen, err = con.Write(generateRandomReading())
    if err != nil {
      fmt.Println(err)
      return
    }

    // For testing, put in a delay between each data reading sent.  This
    // weill test the server response.
    time.Sleep(time.Duration(reading_timeout) * time.Millisecond)
  }
}

// Used for testing, this function creates a Reading struct that has
// randomly assigned attributes
func generateRandomReading() (b []byte) {

  // Generate a Reading and populate with random data
  var reading Reading
  reading.Temperature = (rand.Float64()*600)-300
  reading.Altitude = (rand.Float64()*40000)-20000
  reading.Latitude = (rand.Float64()*180)-90
  reading.Longitude = (rand.Float64()*360)-180
  reading.BatteryLevel = (rand.Float64()*100)

  return reading.Encode()
}
