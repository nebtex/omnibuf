package corev1

//ResourceIDReader ...
type ResourceIDReader interface {
	//Application ...
	Application() string

	//Kind ...
	Kind() uint16

	//ID ...
	ID() string

	//Parent ...
	Parent() (ResourceIDReader, error)
}

//VectorResourceIDReader ...
type VectorResourceIDReader interface {
	// Returns the current size of this vector
	Len() int

	//Get the item in the position i, if i < Len(),
	//if item does not exist should return the default value for the underlying data type
	//when i > Len() should return an VectorInvalidIndexError
	Get(i int) (item ResourceIDReader, err error)
}