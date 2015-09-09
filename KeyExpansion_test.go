
package go_aes

import (
	"testing"
)

func TestKey128Known(t *testing.T) {
	inkey := make([][]byte, 4)
	inkey[0] = []byte{0x2b, 0x7e, 0x15, 0x16}
	inkey[1] = []byte{0x28, 0xae, 0xd2, 0xa6}
	inkey[2] = []byte{0xab, 0xf7, 0x15, 0x88}
	inkey[3] = []byte{0x09, 0xcf, 0x4f, 0x3c}
	//check pieces
	expanded := KeyExpansion(inkey)
	if expanded[0][0] != 0x2b {
		t.Errorf("Expected 0x2b at 0,0; got %x\n", expanded[0][0])
	}
	if expanded[4][3] != 0x17 {
		t.Errorf("Expected 0x17 at 4,3; got %x\n", expanded[4][3])
	}
	if expanded[30][2] != 0x4f {
		t.Errorf("Expected 0x4f at 30,2; got %x\n", expanded[30][2])
	}
	if expanded[43][3] != 0xa6 {
		t.Errorf("Expected 0xa6 at 43,3; got %x\n", expanded[43][3])
	}
}

func TestKey192Known(t *testing.T) {
	inkey := make([][]byte, 6)
	inkey[0] = []byte{0x8e, 0x73, 0xb0, 0xf7}
	inkey[1] = []byte{0xda, 0x0e, 0x64, 0x52} 
	inkey[2] = []byte{0xc8, 0x10, 0xf3, 0x2b}
	inkey[3] = []byte{0x80, 0x90, 0x79, 0xe5}
	inkey[4] = []byte{0x62, 0xf8, 0xea, 0xd2}
	inkey[5] = []byte{0x52, 0x2c, 0x6b, 0x7b}
	//check pieces
	expanded := KeyExpansion(inkey)
	if expanded[0][0] != 0x8e {
		t.Errorf("Expected 0x8e at 0,0; got", expanded[0][0])
	}
	if expanded[6][3] != 0xf7 {
		t.Errorf("Expected 0xf7 at 6,3; got %x\n", expanded[6][3])
	}
	if expanded[33][1] != 0xcb {
		t.Errorf("Expected 0xcb at 33,1; got %x\n", expanded[33][1])
	}
	if expanded[51][3] != 0x02 {
		t.Errorf("Expected 0x02 at 51,3; got %x\n", expanded[51][3])
	}
}

func TestKey256Known(t *testing.T) {
	inkey := make([][]byte, 8)
	inkey[0] = []byte{0x60, 0x3d, 0xeb, 0x10}
	inkey[1] = []byte{0x15, 0xca, 0x71, 0xbe}
	inkey[2] = []byte{0x2b, 0x73, 0xae, 0xf0}
	inkey[3] = []byte{0x85, 0x7d, 0x77, 0x81}
	inkey[4] = []byte{0x1f, 0x35, 0x2c, 0x07}
	inkey[5] = []byte{0x3b, 0x61, 0x08, 0xd7}
	inkey[6] = []byte{0x2d, 0x98, 0x10, 0xa3}
	inkey[7] = []byte{0x09, 0x14, 0xdf, 0xf4}
	//check pieces
	expanded := KeyExpansion(inkey)
	if expanded[0][0] != 0x60 {
		t.Errorf("Expected 0x60 at 0,0; got %x\n", expanded[0][0])
	}
	if expanded[8][3] != 0x11 {
		t.Errorf("Expected 0x11 at 8,3; got %x\n", expanded[8][3])
	}
	if expanded[36][3] != 0x04 {
		t.Errorf("Expected 0x04 at 36,3; got %x\n", expanded[36][3])
	}
	if expanded[59][3] != 0x1e {
		t.Errorf("Expected 0x1e at 59,3; got %x\n", expanded[59][3])
	}
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
