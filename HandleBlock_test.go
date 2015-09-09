
package go_aes

import (
	"testing"
)

/** Test against a known result
 */
func TestEncryptBlockKnown(t *testing.T) {
	inblock := make([][]byte, 4)
	inblock[0] = []byte{0x32, 0x43, 0xf6, 0xa8}
	inblock[1] = []byte{0x88, 0x5a, 0x30, 0x8d}
	inblock[2] = []byte{0x31, 0x31, 0x98, 0xa2}
	inblock[3] = []byte{0xe0, 0x37, 0x07, 0x34}
	inkey := make([][]byte, 4)
	inkey[0] = []byte{0x2b, 0x7e, 0x15, 0x16}
	inkey[1] = []byte{0x28, 0xae, 0xd2, 0xa6}
	inkey[2] = []byte{0xab, 0xf7, 0x15, 0x88}
	inkey[3] = []byte{0x09, 0xcf, 0x4f, 0x3c}
	
	encrypted := make([][]byte, 4)
	encrypted[0] = []byte{0x39, 0x25, 0x84, 0x1d}
	encrypted[1] = []byte{0x02, 0xdc, 0x09, 0xfb}
	encrypted[2] = []byte{0xdc, 0x11, 0x85, 0x97}
	encrypted[3] = []byte{0x19, 0x6a, 0x0b, 0x32}
	
	result := EncryptBlock(inblock, inkey)
	
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] != encrypted[i][j] {
				t.Errorf("at %d %d found %x when expected %x", i, j, result[i][j], encrypted[i][j])
			}
		}
	}
}
