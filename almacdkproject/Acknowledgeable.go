package almacdkproject


// Interface for acknowledging warnings.
type Acknowledgeable struct {
	Id *string `field:"required" json:"id" yaml:"id"`
	Message *string `field:"optional" json:"message" yaml:"message"`
}

