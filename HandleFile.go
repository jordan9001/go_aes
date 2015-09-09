
package go_aes

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
	outf, err := os.Create(outfile)
	check(err)
	// open the infile
	inf, err := os.Open(infile)
	check(err)
	
	// for demonstration purposes, expect a file of a perfect 16 bytes, and use getPureBlock
	block := getPureBlock(inf)
	fmt.Printf("Plaintext Block :\n%v\n",block)
	
	keyblock := parseKey(key)
	fmt.Printf("Key             :\n%v\n\n",keyblock)
	
	encrypted := EncryptBlock(block, keyblock)
	fmt.Printf("Encrypted Block :\n%v\n", encrypted)
	
	for _, v := range encrypted {
		_, err = outf.Write(v)
		check(err)
	}
	
	return true;
}

func DecryptFile(infile string, outfile string, key string) (success bool) {
	// create the outfile
	outf, err := os.Create(outfile)
	check(err)
	// open the infile
	inf, err := os.Open(infile)
	check(err)
	
	// for demonstration purposes, expect a file of a perfect 16 bytes, and use getPureBlock
	block := getPureBlock(inf)
	fmt.Printf("Encrypted Block :\n%v\n",block)
	
	keyblock := parseKey(key)
	fmt.Printf("Key             :\n%v\n\n",keyblock)
	
	encrypted := DecryptBlock(block, keyblock)
	fmt.Printf("Decrypted Block :\n%v\n", encrypted)
	
	for _, v := range encrypted {
		_, err = outf.Write(v)
		check(err)
	}
	
	return true;
}

//This turns the file with the key into our nice 2d array form
func parseKey(key string) (keyblock [][]byte) {
	keyf, err := os.Open(key)
	check(err)
	for {
		line := make([]byte, 4)
		_, err := keyf.Read(line)
		if err != nil {
			break
		} else {
			keyblock = append(keyblock, line)
		}
	}
	return keyblock
}

//This is a function that will crash if the file doesn't have the proper length
func getPureBlock(inf *os.File) (s State) {
	s = make([][]byte, 4)
	for i := 0; i < 4; i++ {
		s[i] = make([]byte, 4)
		_, err := inf.Read(s[i])
		check(err)
	}
	return s
}

//This is a function that adds padding, which should be used when processing files of unknown lengths
func getNextState(inf *os.File) (nextstate State) {
	nextstate = make([][]byte, 4)
	padding := 0
	var padchar byte = 0x07 //I put it as the bell character, because then it will ring if you cat the file out, haha
	for i := 0; i < 3; i++ {
		nextstate[i] = make([]byte, 4)
		if padding == 0 {
			nb, _ := inf.Read(nextstate[i])
			for j := nb; j < 4; j++ {
				nextstate[i][j] = padchar
				padding++
			}
		} else {
			for j := 0; j < 4; j++ {
				nextstate[i][j] = padchar
				padding++
			}
		}
	}
	nextstate[3] = make([]byte, 3)
	if padding > 0 {
		for j := 0; j < 3; j++ {
			nextstate[3][j] = padchar
			padding++
		}
	} else {
		nb, _ := inf.Read(nextstate[3])
		for j := nb; j < 3; j++ {
			nextstate[3][j] = padchar
			padding++
		}
	}
	nextstate[3] = append(nextstate[3], byte(padding+1))
	return nextstate
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

