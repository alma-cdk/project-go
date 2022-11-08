// Opinionated CDK Project “Framework”
package almacdkproject


// Props `AccountStrategy.two`.
// Experimental.
type AccountStrategyTwoProps struct {
	// Experimental.
	Dev *AccountConfiguration `field:"required" json:"dev" yaml:"dev"`
	// Experimental.
	Prod *AccountConfiguration `field:"required" json:"prod" yaml:"prod"`
	// Experimental.
	Mock *AccountConfiguration `field:"optional" json:"mock" yaml:"mock"`
}

