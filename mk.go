// Package mk provides type-inference wrappers around make().
//
// Sometimes when calling make(T), T has already been declared
// elsewhere. For example, if initializing a map in a struct, you
// might wind up with a situation like
//
//	struct Example {
//	  Values map[string]any
//	}
//
//	func NewExample() *Example {
//	  return &Example{Values: make(map[string]any)}
//	}
//
// Note how the type of the map is repeated. It is this exact
// situation that this package is for. Instead of repeating the type,
// you can use this package to eliminate the repetition:
//
//	func NewExample() *Example {
//	  var ex Example
//	  mk.Map(&ex.Values, 0)
//	  return &ex
//	}
package mk

// Slice makes a slice of the given length and capacity and stores it
// in s.
func Slice[S ~[]E, E any](s *S, len, cap int) {
	*s = make(S, len, cap)
}

// Map makes a map of the given capacity and stores it in m.
func Map[M ~map[K]V, K comparable, V any](m *M, cap int) {
	*m = make(M, cap)
}

// Cap makes a channel of the given capacity and stores it in c.
func Chan[C ~chan E, E any](c *C, cap int) {
	*c = make(C, cap)
}
