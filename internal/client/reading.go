package client

import (
  "encoding/binary"
  "bytes"
  "errors"
  "math"
  "github.com/spin-org/thermomatic/internal/common"
)

var (
	ErrDecodeLength  = errors.New("client: decode length is not 40 bytes")
)

// Reading is the set of device readings.
type Reading struct {
	// Temperature denotes the temperature reading of the message.
	Temperature float64

	// Altitude denotes the altitude reading of the message.
	Altitude float64

	// Latitude denotes the latitude reading of the message.
	Latitude float64

	// Longitude denotes the longitude reading of the message.
	Longitude float64

	// BatteryLevel denotes the battery level reading of the message.
	BatteryLevel float64
}

// Decode decodes the reading message payload in the given b into r.
//
// If any of the fields are outside their valid min/max ranges ok will be unset.
//
// Decode does NOT allocate under any condition. Additionally, it panics if b
// isn't at least 40 bytes long.
func (r *Reading) Decode(b []byte) (ok bool) {
  // First thing is to check the length to ensure it's 40 bytes long
  if len(b) != 40 {
    common.Err(ErrDecodeLength)
    panic(ErrDecodeLength)
  }

  // First convert the values
  r.Temperature = math.Float64frombits(binary.LittleEndian.Uint64(b[0:8]))
  r.Altitude = math.Float64frombits(binary.LittleEndian.Uint64(b[8:16]))
  r.Latitude = math.Float64frombits(binary.LittleEndian.Uint64(b[16:24]))
  r.Longitude = math.Float64frombits(binary.LittleEndian.Uint64(b[24:32]))
  r.BatteryLevel = math.Float64frombits(binary.LittleEndian.Uint64(b[32:40]))

  // Next check the values against the limits and unset is outside

  return true
}

// This function will take a Reading struct and return the byte
// array from it that will be sent through the socket.  This function
// it ONLY used in the client for testing and is not part of the server.
func (r *Reading) Encode() (b []byte) {
  total := append(toBytes(r.Temperature), toBytes(r.Altitude)...)
  total = append(total, toBytes(r.Latitude)...)
  total = append(total, toBytes(r.Longitude)...)
  total = append(total, toBytes(r.BatteryLevel)...)

  return total
}

// This function will turn a float64 into a byte array
// LittleEndian order
func toBytes(input float64) (b []byte) {
  buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, input)
	if err != nil {
		common.Err(err)
	}
	return buf.Bytes()
}
