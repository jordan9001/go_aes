
package aes

import (
	"os"
	"fmt"
)

/** Creates an encrypted file from a key and a file.
 * @arg infile A string path to a file to be encrypted.
 * @arg outfile A string path for a new file to be created, with the encrypted data
 * @arg key A 128, 192, or 256 bit key as a string
 */
func EncryptFile(infile string, outfile string, key string) (success bool) {
	// create the outfile
	// open the infile
	// encrypt the file by blocks
	
	// we are going to pad our blocks, no matter the length. If there is only one byte of padding (the minimum) then the last byte will just be 01
	// the last byte will always say how many bytes have been padded on, we only deal with files that have full bytes
	// this could be suceptible to a padding oracle attack
}
