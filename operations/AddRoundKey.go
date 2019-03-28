package operations

var roundKeys = []byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1b, 0x36}

// ConvertTo2D converts a 16 byte array into a column order 2d byte array
// firstSize must be greater than secondSize
func ConvertTo2D(arr []byte, firstSize int, secondSize int) [][]byte {
	arr2D := make([][]byte, firstSize)
	for i := 0; i < firstSize; i++ {
		arr2D[i] = make([]byte, 0, secondSize)
		vector := make([]byte, secondSize)
		for j := 0; j < secondSize; j++ {
			arr2D[i] = append(arr2D[i], vector[j])
		}
	}
	arr2D[0][0] = arr[0]
	arr2D[1][0] = arr[1]
	arr2D[2][0] = arr[2]
	arr2D[3][0] = arr[3]
	arr2D[0][1] = arr[4]
	arr2D[1][1] = arr[5]
	arr2D[2][1] = arr[6]
	arr2D[3][1] = arr[7]
	arr2D[0][2] = arr[8]
	arr2D[1][2] = arr[9]
	arr2D[2][2] = arr[10]
	arr2D[3][2] = arr[11]
	arr2D[0][3] = arr[12]
	arr2D[1][3] = arr[13]
	arr2D[2][3] = arr[14]
	arr2D[3][3] = arr[15]
	return arr2D
}

// ConvertTo1D converts 2D byte array into single dimension byte array
func ConvertTo1D(arr2D [][]byte) []byte {
	arr := make([]byte, 16)
	arr[0] = arr2D[0][0]
	arr[1] = arr2D[1][0]
	arr[2] = arr2D[2][0]
	arr[3] = arr2D[3][0]
	arr[4] = arr2D[0][1]
	arr[5] = arr2D[1][1]
	arr[6] = arr2D[2][1]
	arr[7] = arr2D[3][1]
	arr[8] = arr2D[0][2]
	arr[9] = arr2D[1][2]
	arr[10] = arr2D[2][2]
	arr[11] = arr2D[3][2]
	arr[12] = arr2D[0][3]
	arr[13] = arr2D[1][3]
	arr[14] = arr2D[2][3]
	arr[15] = arr2D[3][3]
	return arr
}

func specialRow(baseExtendedKey [][]byte, row int) {
	keySlice := make([]byte, 4)

	keySlice[0] = baseExtendedKey[0][row-1]
	keySlice[1] = baseExtendedKey[1][row-1]
	keySlice[2] = baseExtendedKey[2][row-1]
	keySlice[3] = baseExtendedKey[3][row-1]

	ShiftCol(keySlice)
	ByteSubWholeByte(keySlice)
	keySlice[0] ^= roundKeys[0]

	baseExtendedKey[0][row] = keySlice[0]
	baseExtendedKey[1][row] = keySlice[1]
	baseExtendedKey[2][row] = keySlice[2]
	baseExtendedKey[3][row] = keySlice[3]

	baseExtendedKey[0][row] ^= baseExtendedKey[0][row-4]
	baseExtendedKey[0][row] ^= baseExtendedKey[0][row-4]
	baseExtendedKey[0][row] ^= baseExtendedKey[0][row-4]
	baseExtendedKey[0][row] ^= baseExtendedKey[0][row-4]
}

func normalRow(baseExtendedKey [][]byte, row int) {
	baseExtendedKey[0][row] = baseExtendedKey[0][row-1] ^ baseExtendedKey[0][row-4]
	baseExtendedKey[1][row] = baseExtendedKey[1][row-1] ^ baseExtendedKey[1][row-4]
	baseExtendedKey[2][row] = baseExtendedKey[2][row-1] ^ baseExtendedKey[2][row-4]
	baseExtendedKey[3][row] = baseExtendedKey[3][row-1] ^ baseExtendedKey[3][row-4]
}

// ExtendKey generates the full round key
func ExtendKey(key []byte) [][]byte {
	baseExtendedKey := ConvertTo2D(key, 4, 44)
	for i := 4; i < 44; i++ {
		if i%4 == 0 {
			specialRow(baseExtendedKey, i)
		} else {
			normalRow(baseExtendedKey, i)
		}
	}
	return baseExtendedKey
}

// AddRoundKey XOR the block with the roundkey
func AddRoundKey(block [][]byte, fullKey [][]byte, round int) {
	i1 := round * 4
	i2 := i1 + 1
	i3 := i2 + 1
	i4 := i3 + 1

	block[0][0] ^= fullKey[0][i1]
	block[0][1] ^= fullKey[1][i1]
	block[0][2] ^= fullKey[2][i1]
	block[0][3] ^= fullKey[3][i1]

	block[1][0] ^= fullKey[0][i2]
	block[1][1] ^= fullKey[1][i2]
	block[1][2] ^= fullKey[2][i2]
	block[1][3] ^= fullKey[3][i2]

	block[2][0] ^= fullKey[0][i3]
	block[2][1] ^= fullKey[1][i3]
	block[2][2] ^= fullKey[2][i3]
	block[2][3] ^= fullKey[3][i3]

	block[3][0] ^= fullKey[0][i4]
	block[3][1] ^= fullKey[1][i4]
	block[3][2] ^= fullKey[2][i4]
	block[3][3] ^= fullKey[3][i4]
}
