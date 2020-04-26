package client

import (
  "testing"
  "bytes"
)

// This tests the Encode function that will be used by the client to
// test the server portion.
func TestReadingEncode(t *testing.T) {
  // Declare a Reading struc and populate the attributes
  var reading Reading
  reading.Temperature = 24.3
  reading.Altitude = 1324.34
  reading.Latitude = 37.773972
  reading.Longitude = -122.431297
  reading.BatteryLevel = 82.4

  // Encode to get the byte array
  result := reading.Encode()

  // Declare statics for comparison
  tmp_bytes := []uint8{0xcd, 0xcc, 0xcc, 0xcc, 0xcc, 0x4c, 0x38, 0x40}
  alt_bytes := []uint8{0x8f, 0xc2, 0xf5, 0x28, 0x5c, 0xb1, 0x94, 0x40}
  lat_bytes := []uint8{0x86, 0x02, 0xb6, 0x83, 0x11, 0xe3, 0x42, 0x40}
  lon_bytes := []uint8{0x3a, 0x77, 0xbb, 0x5e, 0x9a, 0x9b, 0x5e, 0xc0}
  bat_bytes := []uint8{0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0x54, 0x40}

  // Compare each individual slice of the array corresponding to the
  // particular attribute in the original struc
  if !bytes.Equal([]byte(tmp_bytes), []byte(result[0:8])) {
    t.Errorf("Error in encoding Temperature")
  }
  if !bytes.Equal([]byte(alt_bytes), []byte(result[8:16])) {
    t.Errorf("Error in encoding Altitude")
  }
  if !bytes.Equal([]byte(lat_bytes), []byte(result[16:24])) {
    t.Errorf("Error in encoding Latitude")
  }
  if !bytes.Equal([]byte(lon_bytes), []byte(result[24:32])) {
    t.Errorf("Error in encoding Longitude")
  }
  if !bytes.Equal([]byte(bat_bytes), []byte(result[32:40])) {
    t.Errorf("Error in encoding BatteryLevel")
  }
}
