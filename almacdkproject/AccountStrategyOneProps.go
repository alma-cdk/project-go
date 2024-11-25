package almacdkproject


// Props `AccountStrategy.one`.
type AccountStrategyOneProps struct {
	Shared *AccountConfiguration `field:"required" json:"shared" yaml:"shared"`
	Mock *AccountConfiguration `field:"optional" json:"mock" yaml:"mock"`
}

