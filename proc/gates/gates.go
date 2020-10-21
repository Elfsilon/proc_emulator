package gates

// NAND gate is implementation of the NOT AND logical operator
func NAND(x, y bool) bool {
	return !(x && y)
}

// NOT gate is implementation of the NOT logical operator
func NOT(x bool) bool {
	return !x
}

// AND gate is implementation of the AND logical operator
func AND(x, y bool) bool {
	return x && y
}

// OR gate is implementation of the OR logical operator
func OR(x, y bool) bool {
	return x || y
}

// XOR gate is implementation of the XOR logical operator
func XOR(x, y bool) bool {
	return x && !y || !x && y
}

// AND3 gate is implementation of the AND logical operator
func AND3(x, y, z bool) bool {
	return x && y && z
}

// OR3 gate is implementation of the OR logical operator
func OR3(x, y, z bool) bool {
	return x || y || z
}

// OR4 gate is implementation of the OR logical operator
func OR4(x, y, z, k bool) bool {
	return x || y || z || k
}
