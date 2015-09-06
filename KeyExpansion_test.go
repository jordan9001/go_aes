
package aes

import (
	"testing"
)

func TestKeyKnown(t *testing.T) {
	inkey := make([][]byte, 4)
	inkey[0] = []byte{0x00, 0x01, 0x02, 0x03}
	inkey[1] = []byte{0x04, 0x05, 0x06, 0x07}
	inkey[2] = []byte{0x08, 0x09, 0x0a, 0x0b}
	inkey[3] = []byte{0x0c, 0x0d, 0x0e, 0x0f}
	t.Log(KeyExpansion(inkey))
}

func TestXOR(t *testing.T) {
	col1 := []byte{0x00, 0x01, 0x02, 0x03}
	col2 := []byte{0x04, 0x05, 0x06, 0x07}
	
	result := XorCol(col1, col2)
	for i := 0; i< len(result); i++ {
		if result[i] != 0x04 {
			t.Fail()
		}
	}
}
