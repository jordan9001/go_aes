
package go_aes

import (
	"testing"
)

/** Test against a known result
 */
func TestShiftKnown(t *testing.T) {
	var example_state State = make([][]byte, 4)
	// set the columns
	example_state[0] = []byte{0x00, 0x01, 0x02, 0x03}
	example_state[1] = []byte{0x04, 0x05, 0x06, 0x07}
	example_state[2] = []byte{0x08, 0x09, 0x0a, 0x0b}
	example_state[3] = []byte{0x0c, 0x0d, 0x0e, 0x0f}
	
	// shift rows
	example_state = ShiftRows(example_state)
	
	// test result
	var expected_state State = make([][]byte, 4)
	expected_state[0] = []byte{0x00, 0x01, 0x02, 0x03}
	expected_state[1] = []byte{0x05, 0x06, 0x07, 0x04}
	expected_state[2] = []byte{0x0a, 0x0b, 0x08, 0x09}
	expected_state[3] = []byte{0x0f, 0x0c, 0x0d, 0x0e}
	
	for i, v := range example_state {
		for j, byte := range v {
			if byte != expected_state[i][j] {
				t.Errorf("%v does not match expected %v", v, expected_state[i])
			}
		}
	}
}
