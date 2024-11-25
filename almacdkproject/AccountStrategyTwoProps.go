package almacdkproject


// Props `AccountStrategy.two`.
type AccountStrategyTwoProps struct {
	Dev *AccountConfiguration `field:"required" json:"dev" yaml:"dev"`
	Prod *AccountConfiguration `field:"required" json:"prod" yaml:"prod"`
	Mock *AccountConfiguration `field:"optional" json:"mock" yaml:"mock"`
}

