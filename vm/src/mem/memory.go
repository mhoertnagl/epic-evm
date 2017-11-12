package mem

import (
	"mem/exceptions"
)

/*
import (
  "fmt"
  "bufio"
  "encoding/binary"
  "io"
  "os"
)
*/
type IMemory interface {
	
	Read(address uint32) (int32, error)

	Write(address uint32, value int32) error
}

type Memory struct {
	// origin uint32
	data []int32
}

func (mem *Memory) Read(address uint32) (int32, error) {
	limit := uint32(len(mem.data))
	if address >= limit {
		return 0, exceptions.OutOfBoundsError{address, limit}
	}
	return mem.data[address], nil
}

func (mem *Memory) Write(address uint32, value int32) error {
	limit := uint32(len(mem.data))
	if address >= limit {
		return exceptions.OutOfBoundsError{address, limit}
	}
	mem.data[address] = value
	return nil
}

func NewMemory(path string, size uint16) (*Memory, error) {

	return nil, nil
}
