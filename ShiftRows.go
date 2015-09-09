
package go_aes

import (
)

func ShiftRows(instate State) (outstate State){
	outstate = make([][]byte, 4)
	for col := 0; col < len(outstate); col++ {
		outstate[col] = make([]byte, 4)
	}
	for col := 0; col < len(instate); col++ {
		for row := 0; row < len(instate[col]); row++ {
			// for our shift function we fill each column with the appropriately shifted neighbor
			// conviently the column number is the amount we need to offset, mod to wrap around
			outstate[col][row] = instate[(col+row) % len(instate)][row]
		}
	}
	return outstate
}

func InvShiftRows(instate State) (outstate State){
	outstate = make([][]byte, 4)
	for col := 0; col < len(outstate); col++ {
		outstate[col] = make([]byte, 4)
	}
	for col := 0; col < len(instate); col++ {
		for row := 0; row < len(instate[col]); row++ {
			// This is pretty simple, just put where it would have been taken from
			outstate[(col+row) % len(instate)][row] = instate[col][row]
		}
	}
	return outstate
}
