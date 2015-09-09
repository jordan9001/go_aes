
package go_aes

import (
)

/** Takes a 128 bit block, and a (128, 192, or 256) bit key, and returns the encrypted block.
 * Expects a full block, and a key of one of those sizes.
 */
func EncryptBlock(block State, key [][]byte) State {
	key_len := len(key) // number of 32 bit columns in key 4(128 bit), 6(192 bit), or 8(256 bit)
	rounds := key_len + 6
	
	// expand key
	key = KeyExpansion(key)
	
	// initial round
	_ = "breakpoint"
	block = AddRoundKey(block, key[:4])
	key = key[4:]
	
	// rounds
	for i := 0; i < (rounds - 1); i++ {
		_ = "breakpoint"
		block = SubState(block) 
		block = ShiftRows(block)
		block = MixState(block)
		block = AddRoundKey(block, key[:4])
		key = key[4:]
	}

	// final round
	block = SubState(block) 
	block = ShiftRows(block)
	block = AddRoundKey(block, key[:4])
	
	return block
}

/** Takes a 128 bit encrypted block, and a (128, 192, or 256) bit key, and returns the decrypted block.
 * Expects a full block, and a key of one of those sizes.
 */
func DecryptBlock(block State, key [][]byte) State {
	key_len := len(key) // number of 32 bit columns in key 4(128 bit), 6(192 bit), or 8(256 bit)
	rounds := key_len + 6
	
	// expand key
	key = KeyExpansion(key)
	
	// initial round
	block = AddRoundKey(block, key[len(key)-4:])
	key = key[:len(key)-4]
	
	// rounds
	for i := 0; i < (rounds - 1); i++ {
		block = InvShiftRows(block)
		block = InvSubState(block) 
		block = AddRoundKey(block, key[len(key)-4:])
		key = key[:len(key)-4]
		block = InvMixState(block)
	}

	// final round
	block = InvShiftRows(block)
	block = InvSubState(block)
	block = AddRoundKey(block, key[len(key)-4:])
	
	return block
}
