package imei

import (
	"testing"
	//"github.com/spin-org/thermomatic/internal/common"
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

/*
func TestDecodeAllocations(t *testing.T) {
	panic(common.ErrNotImplemented)
}

func TestDecodePanics(t *testing.T) {
	panic(common.ErrNotImplemented)
}

func BenchmarkDecode(b *testing.B) {
	panic(common.ErrNotImplemented)
}
*/
