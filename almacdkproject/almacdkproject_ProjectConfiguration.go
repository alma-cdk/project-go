// Opinionated CDK Project “Framework”
package almacdkproject


// Experimental.
type ProjectConfiguration struct {
	// Dictionary of AWS account specific configuration.
	//
	// The key value can be anything (such as AWS Account alias), but it's recommended to keep it short such as `dev` or `prod`.
	//
	// Example:
	//   accounts: {
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
	Accounts *map[string]*Account `field:"required" json:"accounts" yaml:"accounts"`
	// Author information.
	//
	// I.e. who owns/develops this project/service.
	// Experimental.
	Author *Author `field:"required" json:"author" yaml:"author"`
	// Name of your project/service.
	//
	// Prefer `hyphen-case`.
	//
	// Example:
	//   'my-cool-project'
	//
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify default region you wish to use.
	//
	// If left empty will default to one of the following in order:
	// 1. `$CDK_DEFAULT_REGION`
	// 2. `$AWS_REGION`
	// 3. 'us-east-1'
	// Experimental.
	DefaultRegion *string `field:"optional" json:"defaultRegion" yaml:"defaultRegion"`
}

