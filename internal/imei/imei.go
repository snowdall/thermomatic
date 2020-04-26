// Package imei implements an IMEI decoder.
package imei

// NOTE: for more information about IMEI codes and their structure you may
// consult with:
//
// https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity.

import (
	"errors"

	"github.com/spin-org/thermomatic/internal/common"
)

var (
	ErrInvalid  = errors.New("imei: invalid ")
	ErrChecksum = errors.New("imei: invalid checksum")
)

// Decode returns the IMEI code contained in the first 15 bytes of b.
//
// In case b isn't strictly composed of digits, the returned error will be
// ErrInvalid.
//
// In case b's checksum is wrong, the returned error will be ErrChecksum.
//
// Decode does NOT allocate under any condition. Additionally, it panics if b
// isn't at least 15 bytes long.
func Decode(b []byte) (code uint64, err error) {

  // Declare a couple place holders
  var result uint64 = 0
  var sum uint8 = 0

  // Check that the incoming slice is exactly 15 bytes
  // NOTE: Since the challenge instructions mentioned "IMEI" and
  // not specifically "IMEISV" I made the express decision to assume
  // that the optionally longer 17 byte IMEISV was not being used.
  // This of course was a conscious choice to make an interrpretation of
  // the challenge instructions ... my the coding gods have mercy upon me.
  if len(b) != 15 {
    common.Err(ErrInvalid)
    panic(ErrInvalid)
  }

  // Cycle through the byte slice
  for i := 0; i < 15; i++ {
    // If the index is even (meaning odd digit) just add to summ
    if i % 2 == 0 {
      sum = sum + uint8(b[i])

    // Otherwise double the byte
    } else {
      double := 2 * uint8(b[i])

      // Check to see if the double is over 9.  If so then take 1
      // plus the modulo of 10 (assumes largest double would be 18
      // therefore add the "1")
      if double > 9 {
        sum = sum + 1 + (double % 10)
      } else {
        sum = sum + double
      }
    }
  }

  // If the sum modulo 10 is 0 then it's valid.  At which point
  // cycle through the bytes and build it into the final uint64
  // for return.  Otherwise print the err and panic.
  if sum % 10 == 0 {
    for i := 0; i < 15; i++ {
      result = result * uint64(10) + uint64(b[i])
    }
  } else {
    common.Err(ErrChecksum)
    panic(ErrChecksum)
  }

  return result, nil
}
