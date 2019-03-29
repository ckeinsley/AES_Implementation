package operations

func multiplyByXPlus1(b byte) byte {
	return multiplyByX(b) ^ b
}

func multiplyByX(b byte) byte {
	if b&0x80 > 0 {
		return (b << 1) ^ 27
	}
	return b << 1
}

// MixColumn multiplies the block within GF(2^8)
func MixColumn(block [][]byte) {
	temp0 := multiplyByX(block[0][0]) ^ multiplyByXPlus1(block[1][0]) ^ block[2][0] ^ block[3][0]
	temp1 := multiplyByX(block[0][1]) ^ multiplyByXPlus1(block[1][1]) ^ block[2][1] ^ block[3][1]
	temp2 := multiplyByX(block[0][2]) ^ multiplyByXPlus1(block[1][2]) ^ block[2][2] ^ block[3][2]
	temp3 := multiplyByX(block[0][3]) ^ multiplyByXPlus1(block[1][3]) ^ block[2][3] ^ block[3][3]
	temp4 := multiplyByX(block[1][0]) ^ multiplyByXPlus1(block[2][0]) ^ block[3][0] ^ block[0][0]
	temp5 := multiplyByX(block[1][1]) ^ multiplyByXPlus1(block[2][1]) ^ block[3][1] ^ block[0][1]
	temp6 := multiplyByX(block[1][2]) ^ multiplyByXPlus1(block[2][2]) ^ block[3][2] ^ block[0][2]
	temp7 := multiplyByX(block[1][3]) ^ multiplyByXPlus1(block[2][3]) ^ block[3][3] ^ block[0][3]
	temp8 := multiplyByX(block[2][0]) ^ multiplyByXPlus1(block[3][0]) ^ block[0][0] ^ block[1][0]
	temp9 := multiplyByX(block[2][1]) ^ multiplyByXPlus1(block[3][1]) ^ block[0][1] ^ block[1][1]
	temp10 := multiplyByX(block[2][2]) ^ multiplyByXPlus1(block[3][2]) ^ block[0][2] ^ block[1][2]
	temp11 := multiplyByX(block[2][3]) ^ multiplyByXPlus1(block[3][3]) ^ block[0][3] ^ block[1][3]
	temp12 := multiplyByX(block[3][0]) ^ multiplyByXPlus1(block[0][0]) ^ block[1][0] ^ block[2][0]
	temp13 := multiplyByX(block[3][1]) ^ multiplyByXPlus1(block[0][1]) ^ block[1][1] ^ block[2][1]
	temp14 := multiplyByX(block[3][2]) ^ multiplyByXPlus1(block[0][2]) ^ block[1][2] ^ block[2][2]
	temp15 := multiplyByX(block[3][3]) ^ multiplyByXPlus1(block[0][3]) ^ block[1][3] ^ block[2][3]
	block[0][0] = temp0
	block[0][1] = temp1
	block[0][2] = temp2
	block[0][3] = temp3
	block[1][0] = temp4
	block[1][1] = temp5
	block[1][2] = temp6
	block[1][3] = temp7
	block[2][0] = temp8
	block[2][1] = temp9
	block[2][2] = temp10
	block[2][3] = temp11
	block[3][0] = temp12
	block[3][1] = temp13
	block[3][2] = temp14
	block[3][3] = temp15
}
