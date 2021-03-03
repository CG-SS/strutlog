package strutlog

// Fieldable represents a struct that can be broken down into Fields, hence "fieldable".
// TODO: come up with a better name.
type Fieldable interface {
	Fields() []Field
}
