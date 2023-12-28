package interfaces

// swap takes in two values and returns them swapped
// because every type imeplements any
// so we do not have to specify the type - will go further into this later.
func swap[T any](a, b T) (T, T) {
	return b, a
}
