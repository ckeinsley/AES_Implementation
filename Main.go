package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"

	"strings"

	"github.com/ckeinsley/operations"
	"github.com/hashicorp/vault/helper/xor"
)

func errorCheck(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}

func extractInputs(data string) (int, int, []byte, []byte) {
	lines := strings.Split(data, "\n")
	iterations, err := strconv.ParseInt(lines[0], 10, 64)
	errorCheck(err, "Unable to parse number of iterations")

	rounds, err := strconv.ParseInt(lines[1], 10, 64)
	errorCheck(err, "Unable to parse number of rounds")

	key, err := hex.DecodeString(lines[2])
	errorCheck(err, "Unable to parse key")

	plaintext, err := hex.DecodeString(lines[3])
	errorCheck(err, "Unable to parse plaintext")

	return int(iterations), int(rounds), key, plaintext
}

func runIteration(rounds int, extendedkey []byte, plaintext []byte) []byte {
	return plaintext
}

func main() {
	data, err := ioutil.ReadFile("./aesinput.txt")
	errorCheck(err, "Unable to read input file ./aesinput.txt")
	iterations, rounds, key, plaintext := extractInputs(string(data))

	var output []byte
	extendedkey := operations.ExtendKey(key)
	for i := 0; i < iterations; i++ {
		output = runIteration(rounds, key, plaintext)
		if i != iterations-1 {
			output, err = xor.XORBytes(output, plaintext)
		}
	}
	outputString := hex.EncodeToString(output)
	fmt.Println(outputString)

}
