package download

// Data struct for download options.
type Options struct {
	Destination string
}

// Set the destination
func (options Options) SetDestination(destination string) {
	options.Destination = destination
}