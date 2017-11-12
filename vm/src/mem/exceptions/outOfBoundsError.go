package exceptions

import "fmt"

type OutOfBoundsError struct {
  Address uint32
  Limit uint32
}

func (err OutOfBoundsError) Error() string {
  return fmt.Sprintf("Address [%x] out of bounds. Memory size [%x].", err.Address, err.Limit)
}
