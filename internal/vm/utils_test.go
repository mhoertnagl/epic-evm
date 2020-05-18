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
