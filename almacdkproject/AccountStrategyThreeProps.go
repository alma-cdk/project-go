package almacdkproject


// Props `AccountStrategy.three`.
// Experimental.
type AccountStrategyThreeProps struct {
	// Experimental.
	Dev *AccountConfiguration `field:"required" json:"dev" yaml:"dev"`
	// Experimental.
	Preprod *AccountConfiguration `field:"required" json:"preprod" yaml:"preprod"`
	// Experimental.
	Prod *AccountConfiguration `field:"required" json:"prod" yaml:"prod"`
	// Experimental.
	Mock *AccountConfiguration `field:"optional" json:"mock" yaml:"mock"`
}

