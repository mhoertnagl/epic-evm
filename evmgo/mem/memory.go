package mem

import (
	"epic-evm/evm/mem/ex"
)

// type IMemory interface {
// 	Read(address uint32) (int32, error)
//
// 	Write(address uint32, value int32) error
// }

// Memory models a contiguous block of editable 32bit words.
type Memory struct {
	// origin uint32
	data []int32
}

func (mem *Memory) Read(address uint32) (int32, error) {
	limit := uint32(len(mem.data))
	if address >= limit {
		return 0, ex.OutOfBoundsError{Address: address, Limit: limit}
	}
	return mem.data[address], nil
}

func (mem *Memory) Write(address uint32, value int32) error {
	limit := uint32(len(mem.data))
	if address >= limit {
		return ex.OutOfBoundsError{Address: address, Limit: limit}
	}
	mem.data[address] = value
	return nil
}

// NewMemory creates a new memory block from a binary file.
func NewMemory(path string, size uint16) (*Memory, error) {

	return nil, nil
}
