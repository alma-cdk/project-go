package almacdkproject


// Interface for acknowledging warnings.
// Experimental.
type Acknowledgeable struct {
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Experimental.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

