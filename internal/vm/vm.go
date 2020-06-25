package vm

import "github.com/mhoertnagl/epic-evm/internal/dsp"

type VM struct {
	dsp *dsp.Dsp
}

func NewVM() *VM {

	dsp := dsp.NewDsp()

	return &VM{dsp: dsp}
}

func (vm *VM) SetContent(x int, y int, c rune) {
	vm.SetContent(x, y, c)
}

// func (vm *VM) Continue() bool {
// 	return vm.dsp.Continue()
// }

func (vm *VM) Finalize() {
	vm.dsp.Finalize()
}
