package imei

import (
	"testing"
)

func TestDecode(t *testing.T) {
  // example valid IMEI number 490154203237518
  var imei = [...]byte { 4, 9, 0, 1, 5, 4, 2, 0, 3, 2, 3, 7, 5, 1, 8}

  // Ask for the decode and check the err
  result, err := Decode(imei[:])

  if err != nil {
    t.Errorf("Error in calling decode")
  }

  // Should be the uint64 vcalue of the IMEI
  if result != 490154203237518 {
    t.Errorf("Did not receive the proper decoded IMEI")
  }

}

func TestDecodeInvalidPanic(t *testing.T) {
  // example invalid IMEI number 49015420323751
  var imei = [...]byte { 4, 9, 0, 1, 5, 4, 2, 0, 3, 2, 3, 7, 5, 1}

  defer func() {
    if r := recover(); r == nil {
      t.Errorf("The code did not produce the appropriate panic")
    }
  }()

  // Ask for the decode and check for panic
  Decode(imei[:])
}

func TestDecodeInvalidChecksumPanic(t *testing.T) {
  // example invalid checksum IMEI number 490154203237519
  var imei = [...]byte { 4, 9, 0, 1, 5, 4, 2, 0, 3, 2, 3, 7, 5, 1, 9}

  defer func() {
    if r := recover(); r == nil {
      t.Errorf("The code did not produce the appropriate panic")
    }
  }()

  // Ask for the decode and check for panic
  Decode(imei[:])
}

func BenchmarkDecode(b *testing.B) {
	b.ReportAllocs()

  var imei = [...]byte { 4, 9, 0, 1, 5, 4, 2, 0, 3, 2, 3, 7, 5, 1, 8}

	for i := 0; i < b.N; i++ {
		result, _ := Decode(imei[:])
    if result != 490154203237518 {
			b.Error("Did not receive the proper decoded IMEI")
		}
	}
}
