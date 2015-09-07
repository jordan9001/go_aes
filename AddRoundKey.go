
package aes

import (
)

func AddRoundKey(instate State, roundkey State) (outstate State){
	outstate = make([][]byte, 4)
	for col := 0; col < len(outstate); col++ {
		outstate[col] = make([]byte, 4)
	}
	for col := 0; col < len(instate); col++ {
		for row := 0; row < len(instate[col]); row++ {
			// Here we just do a bitwise XOR with the subkey on each byte
			outstate[col][row] = instate[col][row] ^ roundkey[col][row]
		}
	}
	return outstate
}
