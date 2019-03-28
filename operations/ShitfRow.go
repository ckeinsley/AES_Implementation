package operations

// ShiftRow shifts the rows
func ShiftRow(block [][]byte) {
	temp := block[1][0]
	block[1][0] = block[1][1]
	block[1][1] = block[1][2]
	block[1][2] = block[1][3]
	block[1][3] = temp

	temp = block[2][0]
	temp2 := block[2][3]
	block[2][0] = block[2][2]
	block[2][2] = temp
	block[2][3] = block[2][1]
	block[2][1] = temp2

	temp = block[3][3]
	block[3][3] = block[3][2]
	block[3][2] = block[3][1]
	block[3][1] = block[3][0]
	block[3][0] = temp
}

// ShiftCol shifts a column by 1
func ShiftCol(key []byte) {
	temp := key[0]
	key[0] = key[1]
	key[1] = key[2]
	key[2] = key[3]
	key[3] = temp
}
