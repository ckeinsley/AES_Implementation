package operations

import "fmt"

func convertKeyTo2D(key []byte) [][]byte {
	var fullKey [4][44]byte

	fmt.Println(fullKey)
	return nil
}

// ExtendKey generates the full round key
func ExtendKey(key []byte) [][]byte {
	convertKeyTo2D(key)
	return nil
}
