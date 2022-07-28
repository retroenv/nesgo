package bus

// Mapper represents a mapper memory access interface.
type Mapper interface {
	BasicMemory

	Name() string
}
