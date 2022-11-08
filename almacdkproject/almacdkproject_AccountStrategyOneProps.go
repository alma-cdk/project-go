// Opinionated CDK Project “Framework”
package almacdkproject


// Props `AccountStrategy.one`.
// Experimental.
type AccountStrategyOneProps struct {
	// Experimental.
	Shared *AccountConfiguration `field:"required" json:"shared" yaml:"shared"`
	// Experimental.
	Mock *AccountConfiguration `field:"optional" json:"mock" yaml:"mock"`
}

