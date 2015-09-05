
package aes

import (
	"testing"
)

/** Test against a known result
 */
func TestMixKnown(t *testing.T) {
	var example_state State = make([][]byte, 4)
	// set the columns
	example_state[0] = []byte{0xdb, 0x13, 0x53, 0x45}
	example_state[1] = []byte{0xf2, 0x0a, 0x22, 0x5c}
	example_state[2] = []byte{0xc6, 0xc6, 0xc6, 0xc6}
	example_state[3] = []byte{0xd4, 0xd4, 0xd4, 0xd5}
	
	// mix columns
	for i := 0; i < len(example_state); i++ {
		example_state[i] = MixColumns(example_state[i])
	}
	
	// test result
	var expected_state State = make([][]byte, 4)
	expected_state[0] = []byte{0x8e, 0x4d, 0xa1, 0xbc}
	expected_state[1] = []byte{0x9f, 0xdc, 0x58, 0x9d}
	expected_state[2] = []byte{0xc6, 0xc6, 0xc6, 0xc6}
	expected_state[3] = []byte{0xd5, 0xd5, 0xd7, 0xd6}
	
	for i, v := range example_state {
		for j, byte := range v {
			if byte != expected_state[i][j] {
				t.Errorf("%v does not match expected %v", v, expected_state[i])
			}
		}
	}
}
