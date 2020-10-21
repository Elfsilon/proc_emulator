package core

import (
	"github.com/Elfsilon/proc_emulator/proc/gates"
	"github.com/Elfsilon/proc_emulator/proc/util"
)

// half is implemetation of the half adder
func half(x, y bool) (bool, bool) {
	sum := gates.XOR(x, y)
	carry := gates.AND(x, y)
	return sum, carry
}

// full is implemetation of the full adder
func full(x, y, carryIn bool) (bool, bool) {
	sum1, carry1 := half(x, y)
	sumOut, carry2 := half(sum1, carryIn)
	carryOut := gates.OR(carry1, carry2)
	return sumOut, carryOut
}

// RCA is implementation of the Ripple Carry Adder
// Return couple (sum array, output carry)
func RCA(xs, ys []bool) ([]bool, bool) {
	var sums []bool
	var curCarry bool = false
	for i := len(xs) - 1; i >= 0; i-- {
		sum, carry := full(xs[i], ys[i], curCarry)
		sums = append(sums, sum)
		curCarry = carry
	}
	return util.ReverseBoolArray(sums), curCarry
}
