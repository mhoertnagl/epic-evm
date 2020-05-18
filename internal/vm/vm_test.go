package vm_test

import (
	"testing"

	"github.com/mhoertnagl/epic-evm/internal/vm"
)

func TestAluADD(t *testing.T) {
	testAluOps(t, 0x00000002, vm.OpADD, 0x00000001, 0x00000001)
	testAluOps(t, 0x00000000, vm.OpADD, 0x00000001, 0xFFFFFFFF)
	testAluOps(t, 0xFE000002, vm.OpADD, 0xFF000001, 0xFF000001)
	testAluOps(t, 0x00000000, vm.OpADD, 0xFFFFFFFF, 0x00000001)
	testAluOps(t, 0xFFFFFFFE, vm.OpADD, 0xFFFFFFFF, 0xFFFFFFFF)
	testAluOps(t, 0xFFFFFFFE, vm.OpADD, 0x80000000, 0x80000000)
}

func TestAluADU(t *testing.T) {
	testAluOps(t, 0x00000002, vm.OpADU, 0x00000001, 0x00000001)
	testAluOps(t, 0x00000000, vm.OpADU, 0x00000001, 0xFFFFFFFF)
	testAluOps(t, 0xFE000002, vm.OpADD, 0xFF000001, 0xFF000001)
	testAluOps(t, 0x00000000, vm.OpADU, 0xFFFFFFFF, 0x00000001)
	testAluOps(t, 0xFFFFFFFE, vm.OpADU, 0xFFFFFFFF, 0xFFFFFFFF)
	testAluOps(t, 0xFFFFFFFE, vm.OpADU, 0x80000000, 0x80000000)
}

func testAluOps(t *testing.T, e uint32, op vm.AOp, va uint32, vb uint32) {
	t.Helper()
	a := vm.Alu(op, va, vb)
	if a != e {
		t.Errorf(
			"VA = %x (%v), VB = %x (%v), Expecting %x (%v) but got %x (%v).",
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
