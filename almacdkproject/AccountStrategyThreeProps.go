package almacdkproject


// Props `AccountStrategy.three`.
type AccountStrategyThreeProps struct {
	Dev *AccountConfiguration `field:"required" json:"dev" yaml:"dev"`
	Preprod *AccountConfiguration `field:"required" json:"preprod" yaml:"preprod"`
	Prod *AccountConfiguration `field:"required" json:"prod" yaml:"prod"`
	Mock *AccountConfiguration `field:"optional" json:"mock" yaml:"mock"`
}

