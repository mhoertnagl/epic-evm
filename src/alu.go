package main

/*
type aluOpFun func(uint32, uint32) uint32

var add = func(left uint32, right uint32) uint32 { return left + right }
var sub = func(left uint32, right uint32) uint32 { return left - right }
var mul = func(left uint32, right uint32) uint32 { return left * right }
var div = func(left uint32, right uint32) uint32 { return left / right }
var and = func(left uint32, right uint32) uint32 { return left & right }
var oor = func(left uint32, right uint32) uint32 { return left | right }
var nor = func(left uint32, right uint32) uint32 { return ^(left | right) }
var xor = func(left uint32, right uint32) uint32 { return left ^ right }
var sll = func(left uint32, right uint32) uint32 { return left << right }
var srl = func(left uint32, right uint32) uint32 { return left >> right }
var sra = func(left uint32, right uint32) uint32 { return uint32(int(left) >> right) }
var ror = func(left uint32, right uint32) uint32 { return (left << (32 - right)) | (left >> right) }

var auluOps = [...]aluOpFun{add, sub, mul, div, and, oor, nor, xor, sll, srl, sra, ror}
*/

func compute(aluOp uint32, left uint32, right uint32) uint32 {
	switch aluOp {
	case AluOpAdd:
		return left + right
	case AluOpSub:
		return left - right
	case AluOpMul:
		return left * right
	case AluOpDiv:
		return left / right
	case AluOpAnd:
		return left & right
	case AluOpOor:
		return left | right
	case AluOpNor:
		return ^(left | right)
	case AluOpXor:
		return left ^ right
	case AluOpSll:
		return left << right
	case AluOpSrl:
		return left >> right
	case AluOpSra:
		return uint32(int(left) >> right)
	case AluOpRor:
		return (left << (32 - right)) | (left >> right)
	}
	return 0
}
