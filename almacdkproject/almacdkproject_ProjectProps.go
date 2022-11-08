// Opinionated CDK Project “Framework”
package almacdkproject


// Props given to `Project`.
//
// I.e. custom props for this construct and the usual props given to `cdk.App`.
// Experimental.
type ProjectProps struct {
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
	// Include runtime versioning information in the Stacks of this app.
	// Experimental.
	AnalyticsReporting *bool `field:"optional" json:"analyticsReporting" yaml:"analyticsReporting"`
	// Automatically call `synth()` before the program exits.
	//
	// If you set this, you don't have to call `synth()` explicitly. Note that
	// this feature is only available for certain programming languages, and
	// calling `synth()` is still recommended.
	// Experimental.
	AutoSynth *bool `field:"optional" json:"autoSynth" yaml:"autoSynth"`
	// Additional context values for the application.
	//
	// Context set by the CLI or the `context` key in `cdk.json` has precedence.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// The output directory into which to emit synthesized artifacts.
	//
	// You should never need to set this value. By default, the value you pass to
	// the CLI's `--output` flag will be used, and if you change it to a different
	// directory the CLI will fail to pick up the generated Cloud Assembly.
	//
	// This property is intended for internal and testing use.
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Include construct creation stack trace in the `aws:cdk:trace` metadata key of all constructs.
	// Experimental.
	StackTraces *bool `field:"optional" json:"stackTraces" yaml:"stackTraces"`
	// Include construct tree metadata as part of the Cloud Assembly.
	// Experimental.
	TreeMetadata *bool `field:"optional" json:"treeMetadata" yaml:"treeMetadata"`
}

