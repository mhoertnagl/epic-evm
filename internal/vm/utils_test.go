package vm_test

import (
	"testing"

	"github.com/mhoertnagl/epic-evm/internal/vm"
)

func TestSextPositive12(t *testing.T) {
	i := uint32(0x00000701)
	a := vm.Sext(i, 12)
	e := uint32(0x00000701)
	if a != e {
		t.Errorf("Expecting %v but got %v.", e, a)
	}
}

func TestSextNegative12(t *testing.T) {
	i := uint32(0x00000F23)
	a := vm.Sext(i, 12)
	e := uint32(0xFFFFFF23)
	if a != e {
		t.Errorf("Expecting %v but got %v.", e, a)
	}
}

func TestSextPositive16(t *testing.T) {
	i := uint32(0x00007123)
	a := vm.Sext(i, 16)
	e := uint32(0x00007123)
	if a != e {
		t.Errorf("Expecting %v but got %v.", e, a)
	}
}

func TestSextNegative16(t *testing.T) {
	i := uint32(0x0000F123)
	a := vm.Sext(i, 16)
	e := uint32(0xFFFFF123)
	if a != e {
		t.Errorf("Expecting %v but got %v.", e, a)
	}
}

func TestSextPositive32(t *testing.T) {
	i := uint32(0x71234567)
	a := vm.Sext64(i)
	e := uint64(0x0000000071234567)
	if a != e {
		t.Errorf("Expecting %v but got %v.", e, a)
	}
}

func TestSextNegative32(t *testing.T) {
	i := uint32(0xF1234567)
	a := vm.Sext64(i)
	e := uint64(0xFFFFFFFFF1234567)
	if a != e {
		t.Errorf("Expecting %v but got %v.", e, a)
	}
}

func TestBit1(t *testing.T) {
	i := uint32(0x01000000)
	a := vm.Bit(i, 24)
	if a != 1 {
		t.Errorf("Bit %v of 0x%x is not set.", 24, a)
	}
}

func TestBit2(t *testing.T) {
	i := uint32(0x00000001)
	a := vm.Bit(i, 0)
	if a != 1 {
		t.Errorf("Bit %v of 0x%x is not set.", 0, a)
	}
}

func TestBit3(t *testing.T) {
	i := uint32(0x80000000)
	a := vm.Bit(i, 31)
	if a != 1 {
		t.Errorf("Bit %v of 0x%x is not set.", 31, a)
	}
}

func TestBit64_1(t *testing.T) {
	i := uint64(0x01000000)
	a := vm.Bit64(i, 24)
	if a != 1 {
		t.Errorf("Bit %v of 0x%x is not set.", 24, a)
	}
}

func TestBit64_2(t *testing.T) {
	i := uint64(0x00000001)
	a := vm.Bit64(i, 0)
	if a != 1 {
		t.Errorf("Bit %v of 0x%x is not set.", 0, a)
	}
}

func TestBit64_3(t *testing.T) {
	i := uint64(0x80000000)
	a := vm.Bit64(i, 31)
	if a != 1 {
		t.Errorf("Bit %v of 0x%x is not set.", 31, a)
	}
}

func TestSet1(t *testing.T) {
	i := uint32(0x00000000)
	a := vm.Set(i, 24, 1)
	e := uint32(0x01000000)
	if a != e {
		t.Errorf("Expecting 0x%x but got 0x%x.", e, a)
	}
}

func TestSet2(t *testing.T) {
	i := uint32(0x00000000)
	a := vm.Set(i, 0, 1)
	e := uint32(0x00000001)
	if a != e {
		t.Errorf("Expecting 0x%x but got 0x%x.", e, a)
	}
}

func TestSet3(t *testing.T) {
	i := uint32(0x00000000)
	a := vm.Set(i, 31, 1)
	e := uint32(0x80000000)
	if a != e {
		t.Errorf("Expecting 0x%x but got 0x%x.", e, a)
	}
}

func TestSet4(t *testing.T) {
	i := uint32(0x80000000)
	a := vm.Set(i, 31, 1)
	e := uint32(0x80000000)
	if a != e {
		t.Errorf("Expecting 0x%x but got 0x%x.", e, a)
	}
}

func TestSet5(t *testing.T) {
	i := uint32(0x80000000)
	a := vm.Set(i, 31, 0)
	e := uint32(0x00000000)
	if a != e {
		t.Errorf("Expecting 0x%x but got 0x%x.", e, a)
	}
}

func TestSetBool1(t *testing.T) {
	i := uint32(0x00000000)
	a := vm.SetBool(i, 24, true)
	e := uint32(0x01000000)
	if a != e {
		t.Errorf("Expecting 0x%x but got 0x%x.", e, a)
	}
}

func TestSetBool2(t *testing.T) {
	i := uint32(0x01000000)
	a := vm.SetBool(i, 24, false)
	e := uint32(0x00000000)
	if a != e {
		t.Errorf("Expecting 0x%x but got 0x%x.", e, a)
	}
}
