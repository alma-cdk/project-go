// Opinionated CDK Project “Framework”
package almacdkproject


// AWS account configuration.
// Experimental.
type Account struct {
	// AWS Account ID.
	//
	// Example:
	//   '123456789012'
	//
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// AWS account specific configuration.
	//
	// For example VPC IDs (for existing VPCs), Direct Connect Gateway IDs, apex domain names (for Route53 Zone lookups), etc. Basically configuration for resources that are defined outside of this CDK application.
	//
	// Example:
	//   {
	//     dev: {
	//       id: '111111111111',
	//       config: {
	//         baseDomain: 'example.net',
	//       },
	//     },
	//     prod: {
	//       id: '222222222222',
	//       config: {
	//         baseDomain: 'example.com',
	//       },
	//     },
	//   },
	//
	// Experimental.
	Config *map[string]interface{} `field:"optional" json:"config" yaml:"config"`
	// List of accepted environments for the given account.
	//
	// List of strings or strings representing regexp initialization (passed onto `new Regexp("^"+environment+"$", "i")`).
	//
	// Example:
	//   ["development", "feature/.*"]
	//
	// Experimental.
	Environments *[]*string `field:"optional" json:"environments" yaml:"environments"`
}

