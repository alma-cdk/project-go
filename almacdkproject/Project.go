package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alma-cdk/project-go/almacdkproject/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/constructs-go/constructs/v10"
)

// High-level wrapper for `cdk.App` with specific requirements for props.
//
// Use it like you would `cdk.App` and assign stacks into it.
//
// Example:
//   // new Project instead of new App
//   const project = new Project({
//     name: 'my-cool-project',
//     author: {
//       organization: 'Acme Corp',
//       name: 'Mad Scientists',
//       email: 'mad.scientists@acme.example.com',
//   },
//   defaultRegion: 'eu-west-1', // defaults to one of: $CDK_DEFAULT_REGION, $AWS_REGION or us-east-1
//   accounts: {
//   dev: {
//    id: '111111111111',
//    environments: ['development', 'feature/.*', 'staging'],
//    config: {
//      baseDomain: 'example.net',
//    },
//   },
//   prod: {
//    id: '222222222222',
//    environments: ['production'],
//    config: {
//      baseDomain: 'example.com',
//    },
//   },
//   },
//   })
//
// Experimental.
type Project interface {
	awscdk.App
	// The default account for all resources defined within this stage.
	// Experimental.
	Account() *string
	// Artifact ID of the assembly if it is a nested stage. The root stage (app) will return an empty string.
	//
	// Derived from the construct path.
	// Experimental.
	ArtifactId() *string
	// The cloud assembly asset output directory.
	// Experimental.
	AssetOutdir() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The cloud assembly output directory.
	// Experimental.
	Outdir() *string
	// The parent stage or `undefined` if this is the app.
	//
	// *.
	// Experimental.
	ParentStage() awscdk.Stage
	// Validation plugins to run during synthesis.
	//
	// If any plugin reports any violation,
	// synthesis will be interrupted and the report displayed to the user.
	// Default: - no validation plugins are used.
	//
	// Experimental.
	PolicyValidationBeta1() *[]awscdk.IPolicyValidationPluginBeta1
	// The default region for all resources defined within this stage.
	// Experimental.
	Region() *string
	// The name of the stage.
	//
	// Based on names of the parent stages separated by
	// hypens.
	// Experimental.
	StageName() *string
	// Synthesize this stage into a cloud assembly.
	//
	// Once an assembly has been synthesized, it cannot be modified. Subsequent
	// calls will return the same assembly.
	// Experimental.
	Synth(options *awscdk.StageSynthesisOptions) cxapi.CloudAssembly
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Project
type jsiiProxy_Project struct {
	internal.Type__awscdkApp
}

func (j *jsiiProxy_Project) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) AssetOutdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetOutdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ParentStage() awscdk.Stage {
	var returns awscdk.Stage
	_jsii_.Get(
		j,
		"parentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PolicyValidationBeta1() *[]awscdk.IPolicyValidationPluginBeta1 {
	var returns *[]awscdk.IPolicyValidationPluginBeta1
	_jsii_.Get(
		j,
		"policyValidationBeta1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}


// Initializes a new Project (which can be used in place of cdk.App).
// Experimental.
func NewProject(props *ProjectProps) Project {
	_init_.Initialize()

	if err := validateNewProjectParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Project{}

	_jsii_.Create(
		"@alma-cdk/project.Project",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Initializes a new Project (which can be used in place of cdk.App).
// Experimental.
func NewProject_Override(p Project, props *ProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.Project",
		[]interface{}{props},
		p,
	)
}

// Return account configuration.
// Experimental.
func Project_GetAccount(scope constructs.Construct, accountType *string) *Account {
	_init_.Initialize()

	if err := validateProject_GetAccountParameters(scope, accountType); err != nil {
		panic(err)
	}
	var returns *Account

	_jsii_.StaticInvoke(
		"@alma-cdk/project.Project",
		"getAccount",
		[]interface{}{scope, accountType},
		&returns,
	)

	return returns
}

// Return the project configuration as given in ProjectProps.
// Experimental.
func Project_GetConfiguration(scope constructs.Construct) *ProjectConfiguration {
	_init_.Initialize()

	if err := validateProject_GetConfigurationParameters(scope); err != nil {
		panic(err)
	}
	var returns *ProjectConfiguration

	_jsii_.StaticInvoke(
		"@alma-cdk/project.Project",
		"getConfiguration",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Checks if an object is an instance of the `App` class.
//
// Returns: `true` if `obj` is an `App`.
// Experimental.
func Project_IsApp(obj interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsAppParameters(obj); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.Project",
		"isApp",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func Project_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.Project",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given construct is a stage.
// Experimental.
func Project_IsStage(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsStageParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.Project",
		"isStage",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return the stage this construct is contained with, if available.
//
// If called
// on a nested stage, returns its parent.
// Experimental.
func Project_Of(construct constructs.IConstruct) awscdk.Stage {
	_init_.Initialize()

	if err := validateProject_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns awscdk.Stage

	_jsii_.StaticInvoke(
		"@alma-cdk/project.Project",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Project_CONTEXT_SCOPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@alma-cdk/project.Project",
		"CONTEXT_SCOPE",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Project) Synth(options *awscdk.StageSynthesisOptions) cxapi.CloudAssembly {
	if err := p.validateSynthParameters(options); err != nil {
		panic(err)
	}
	var returns cxapi.CloudAssembly

	_jsii_.Invoke(
		p,
		"synth",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

