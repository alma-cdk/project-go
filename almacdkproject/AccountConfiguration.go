package almacdkproject


// Interface for a single account type configuration.
type AccountConfiguration struct {
	Id *string `field:"required" json:"id" yaml:"id"`
	Config *map[string]interface{} `field:"optional" json:"config" yaml:"config"`
}

