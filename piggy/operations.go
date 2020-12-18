package piggy

// Operations is just a slice of operations.
// This type exists because []Operation is not a valid
// receiver type in Go.
type Operations []Operation

// Where returns all operations of the current slice
// that match a given predicate.
//
// Example:
//  matching := operations.Where(func(op Operation) bool { return op.ID == 42 })
func (ops Operations) Where(predicate func(Operation) bool) Operations {
	operations := Operations{}
	for _, op := range ops {
		if predicate(op) {
			operations = append(operations, op)
		}
	}
	return operations
}
