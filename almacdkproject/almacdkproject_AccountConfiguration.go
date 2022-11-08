// Opinionated CDK Project “Framework”
package almacdkproject


// Interface for a single account type configuration.
// Experimental.
type AccountConfiguration struct {
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Experimental.
	Config *map[string]interface{} `field:"optional" json:"config" yaml:"config"`
}

