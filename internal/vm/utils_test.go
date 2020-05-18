package vm_test

import (
	"testing"

	"github.com/mhoertnagl/epic-evm/internal/vm"
)

func TestSext0(t *testing.T) {
	i := uint32(0x00001000)
	a := vm.Sext(i, 16)
	e := uint32(0xFFFFF000)
	if a != e {
		t.Errorf("Expecting %v but got %v.", e, a)
	}
}
