package compiler

// Config contains the compiler configuration.
type Config struct {
	// DisableComments does not output any comments.
	DisableComments bool
}

func (c Config) validate() error {
	return nil
}
