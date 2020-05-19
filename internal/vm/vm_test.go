package vm_test

import (
	"testing"

	"github.com/mhoertnagl/epic-evm/internal/vm"
)

func TestShiftSLL(t *testing.T) {
	testShiftOps(t, 0x00000000, 0x00000000, vm.OpSLL, 0)
	testShiftOps(t, 0x00000002, 0x00000001, vm.OpSLL, 1)
	testShiftOps(t, 0x80000000, 0x00000001, vm.OpSLL, 31)
	testShiftOps(t, 0x00001010, 0x00000101, vm.OpSLL, 4)
}

func TestShiftROL(t *testing.T) {
	testShiftOps(t, 0x00000000, 0x00000000, vm.OpROL, 0)
	testShiftOps(t, 0x00000001, 0x80000000, vm.OpROL, 1)
	testShiftOps(t, 0x80000001, 0xC0000000, vm.OpROL, 1)
}

func TestShiftSRL(t *testing.T) {
	testShiftOps(t, 0x00000000, 0x00000000, vm.OpSRL, 0)
	testShiftOps(t, 0x00000001, 0x00000002, vm.OpSRL, 1)
	testShiftOps(t, 0x40000000, 0x80000000, vm.OpSRL, 1)
	testShiftOps(t, 0x00000001, 0x80000000, vm.OpSRL, 31)
	testShiftOps(t, 0x01010000, 0x10100000, vm.OpSRL, 4)
}

func TestShiftSRA(t *testing.T) {
	testShiftOps(t, 0x00000000, 0x00000000, vm.OpSRA, 0)
	testShiftOps(t, 0x00000001, 0x00000002, vm.OpSRA, 1)
	testShiftOps(t, 0x38000000, 0x70000000, vm.OpSRA, 1)
	testShiftOps(t, 0xC0000000, 0x80000000, vm.OpSRA, 1)
	testShiftOps(t, 0xFFFFFFFF, 0x80000000, vm.OpSRA, 31)
}

func TestAluADD(t *testing.T) {
	testAluOps(t, 0x0000000000000002, vm.OpADD, 0x00000001, 0x00000001)
	testAluOps(t, 0x0000000000000000, vm.OpADD, 0x00000001, 0xFFFFFFFF)
	testAluOps(t, 0xFFFFFFFFFE000002, vm.OpADD, 0xFF000001, 0xFF000001)
	testAluOps(t, 0x0000000000000000, vm.OpADD, 0xFFFFFFFF, 0x00000001)
	testAluOps(t, 0xFFFFFFFFFFFFFFFE, vm.OpADD, 0xFFFFFFFF, 0xFFFFFFFF)
	testAluOps(t, 0xFFFFFFFF00000000, vm.OpADD, 0x80000000, 0x80000000)
}

func TestAluADU(t *testing.T) {
	testAluOps(t, 0x0000000000000002, vm.OpADU, 0x00000001, 0x00000001)
	testAluOps(t, 0x0000000100000000, vm.OpADU, 0x00000001, 0xFFFFFFFF)
	testAluOps(t, 0x00000001FE000002, vm.OpADU, 0xFF000001, 0xFF000001)
	testAluOps(t, 0x0000000100000000, vm.OpADU, 0xFFFFFFFF, 0x00000001)
	testAluOps(t, 0x00000001FFFFFFFE, vm.OpADU, 0xFFFFFFFF, 0xFFFFFFFF)
	testAluOps(t, 0x0000000100000000, vm.OpADU, 0x80000000, 0x80000000)
}

func testShiftOps(t *testing.T, e uint32, vb uint32, op vm.SOp, shamt uint32) {
	t.Helper()
	a := vm.Shift(vb, op, shamt)
	if a != e {
		t.Errorf(
			"VA = 0x%x (%v), Expecting 0x%x (%v) but got 0x%x (%v).",
			vb,
			int32(vb),
			e,
			int32(e),
			a,
			int32(a),
		)
	}
}

func testAluOps(t *testing.T, e uint64, op vm.AOp, va uint32, vb uint32) {
	t.Helper()
	a := vm.Alu(op, va, vb)
	if a != e {
		t.Errorf(
			"VA = 0x%x (%v), VB = 0x%x (%v), Expecting 0x%x (%v) but got 0x%x (%v).",
			va,
			int32(va),
			vb,
			int32(vb),
			e,
			int32(e),
			a,
			int32(a),
		)
	}
}
