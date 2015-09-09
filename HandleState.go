
package go_aes

import (
	"fmt"
)

type State [][]byte

func (s State) String() string {
	var output string
	for i := 0; i < len(s); i++ {
		output += fmt.Sprint("[")
		for j := 0; j < len(s[i]); j++ {
			output += fmt.Sprintf(" 0x%x", s[i][j])
		}
		output += fmt.Sprint(" ]\n")
	}
	return output
}

