package client

import (
  "testing"
  "bytes"
)

// This tests the Decode function that is used by the server to
// convert a byte array into float64 values in the Reading struc
func TestReadingDecode(t *testing.T) {
  // Declare a Reading struc and populate the attributes
  var reading Reading

  // Prepare a byte array to be decoded into the Reading struc
  input_bytes := []uint8{
    0x40, 0x38, 0x4c, 0xcc, 0xcc, 0xcc, 0xcc, 0xcd,
    0x40, 0x94, 0xb1, 0x5c, 0x28, 0xf5, 0xc2, 0x8f,
    0x40, 0x42, 0xe3, 0x11, 0x83, 0xb6, 0x02, 0x86,
    0xc0, 0x5e, 0x9b, 0x9a, 0x5e, 0xbb, 0x77, 0x3a,
    0x40, 0x54, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a,
  }

  result := reading.Decode(input_bytes)

  if !result {
    t.Errorf("Error in decoding bytes in Reading struct")
  }

  if reading.Temperature != 24.3 {
    t.Error("Error decoding Temperature")
  }
  if reading.Altitude != 1324.34 {
    t.Error("Error decoding Altitude")
  }
  if reading.Latitude != 37.773972 {
    t.Error("Error decoding Latitude")
  }
  if reading.Longitude != -122.431297 {
    t.Error("Error decoding Longitude")
  }
  if reading.BatteryLevel != 82.4 {
    t.Error("Error decoding BatteryLevel")
  }
}

// This tests to see if values are unset when out of bounds
// This tests the Decode function that is used by the server to
// convert a byte array into float64 values in the Reading struc
func TestReadingDecodeLimits(t *testing.T) {
  // Declare a Reading struc and populate the attributes
  var reading Reading

  // Prepare a byte array to be decoded into the Reading struc
  // The Temperature attribute is set to 321.12 to test limit checks
  input_bytes := []uint8{
    0x40, 0x74, 0x11, 0xeb, 0x85, 0x1e, 0xb8, 0x52,
    0x40, 0x94, 0xb1, 0x5c, 0x28, 0xf5, 0xc2, 0x8f,
    0x40, 0x42, 0xe3, 0x11, 0x83, 0xb6, 0x02, 0x86,
    0xc0, 0x5e, 0x9b, 0x9a, 0x5e, 0xbb, 0x77, 0x3a,
    0x40, 0x54, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a,
  }

  result := reading.Decode(input_bytes)

  if result {
    t.Errorf("Error in decoding bytes in Reading struct")
  }
}

// This tests to ensure there is a panic when there are more or less
// than 40 bytes sent to the Decode function
func TestReadingDecodePanic(t *testing.T) {
  // Declare a Reading struc and populate the attributes
  var reading Reading

  // Prepare a byte array that is one byte too short
  input_bytes := []uint8{
    0x40, 0x38, 0x4c, 0xcc, 0xcc, 0xcc, 0xcc, 0xcd,
    0x40, 0x94, 0xb1, 0x5c, 0x28, 0xf5, 0xc2, 0x8f,
    0x40, 0x42, 0xe3, 0x11, 0x83, 0xb6, 0x02, 0x86,
    0xc0, 0x5e, 0x9b, 0x9a, 0x5e, 0xbb, 0x77, 0x3a,
    0x40, 0x54, 0x99, 0x99, 0x99, 0x99, 0x99,
  }

  defer func() {
    if r := recover(); r == nil {
      t.Errorf("The code did not produce the appropriate panic")
    }
  }()

  result := reading.Decode(input_bytes)

  if !result {
    t.Errorf("Error in decoding bytes in Reading struct")
  }
}

func BenchmarkDecode(b *testing.B) {
	b.ReportAllocs()

  // Declare a Reading struc and populate the attributes
  var reading Reading

  // Prepare a byte array to be decoded into the Reading struc
  input_bytes := []uint8{
    0x40, 0x38, 0x4c, 0xcc, 0xcc, 0xcc, 0xcc, 0xcd,
    0x40, 0x94, 0xb1, 0x5c, 0x28, 0xf5, 0xc2, 0x8f,
    0x40, 0x42, 0xe3, 0x11, 0x83, 0xb6, 0x02, 0x86,
    0xc0, 0x5e, 0x9b, 0x9a, 0x5e, 0xbb, 0x77, 0x3a,
    0x40, 0x54, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a,
  }

	for i := 0; i < b.N; i++ {
		result := reading.Decode(input_bytes)
    if !result  {
			b.Error("Did not receive the proper decoded values")
		}
	}
}

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
  tmp_bytes := []uint8{0x40, 0x38, 0x4c, 0xcc, 0xcc, 0xcc, 0xcc, 0xcd}
  alt_bytes := []uint8{0x40, 0x94, 0xb1, 0x5c, 0x28, 0xf5, 0xc2, 0x8f}
  lat_bytes := []uint8{0x40, 0x42, 0xe3, 0x11, 0x83, 0xb6, 0x02, 0x86}
  lon_bytes := []uint8{0xc0, 0x5e, 0x9b, 0x9a, 0x5e, 0xbb, 0x77, 0x3a}
  bat_bytes := []uint8{0x40, 0x54, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a}

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
