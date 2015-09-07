
package aes

import (
)


var fixed_matrix = [][]byte {
	[]byte{0x02, 0x03, 0x01, 0x01},
	[]byte{0x01, 0x02, 0x03, 0x01},
	[]byte{0x01, 0x01, 0x02, 0x03},
	[]byte{0x03, 0x01, 0x01, 0x02},
	}

var irr_poly byte = 0x1b

/** Takes 1 column and "mixes it" by multiplying against the fixed mixing matrix
 */
func MixColumns(inbytes []byte) (outbytes []byte) {
	outbytes = make([]byte, 4)
	for out_index := 0; out_index < len(fixed_matrix); out_index++ {
		for i := 0; i < len(fixed_matrix[out_index]); i++ {
			outbytes[out_index] = outbytes[out_index] ^ galois_mult(inbytes[i], fixed_matrix[out_index][i])
		}
	}
	return outbytes
}

/** Takes 2 bytes in the Finite Field, and multiplies them modulo our irregular_poly
 */
func galois_mult(in_b, fixed_b byte) byte {
	var result byte = 0x00
	for fixed_b != 0 {
		// if there is a bit in this order of the polynomial, we add to the result
		if fixed_b & 1 != 0 {
			result = result ^ in_b	
		}
		// multiply by 1 (bitshift 1), and mod to stay within the Field (XOR with irreducable poly)
		toolarge := in_b & 0x80
		in_b = in_b << 1
		if toolarge != 0 {
			in_b = in_b ^ irr_poly
		}
		//move to the next bit 
		fixed_b = fixed_b >> 1
	}
	return result
}

func MixState(instate State) (outstate State) {	
	outstate = make([][]byte, 4)
	for i := 0; i < len(instate); i++ {
		outstate[i] = MixColumns(instate[i])
	}
	return outstate
}
