package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type EnvironmentContext interface {
}

// The jsii proxy struct for EnvironmentContext
type jsiiProxy_EnvironmentContext struct {
	_ byte // padding
}

func NewEnvironmentContext() EnvironmentContext {
	_init_.Initialize()

	j := jsiiProxy_EnvironmentContext{}

	_jsii_.Create(
		"@alma-cdk/project.EnvironmentContext",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewEnvironmentContext_Override(e EnvironmentContext) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.EnvironmentContext",
		nil, // no parameters
		e,
	)
}

// Get Environment Category.
//
// Categories are useful grouping to make distinction between `stable`
// environments (`staging` & `production`) from `feature` or `verification`
// environments (such as `test` or `preproduction`).
//
// Returns: Environment Category.
//
// Example:
//   'mock'
//   'development'
//   'feature'
//   'verification'
//   'stable'
//
func EnvironmentContext_GetCategory(scope constructs.Construct) EnvironmentCategory {
	_init_.Initialize()

	if err := validateEnvironmentContext_GetCategoryParameters(scope); err != nil {
		panic(err)
	}
	var returns EnvironmentCategory

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"getCategory",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Get Feature Info.
//
// If environment belongs to `feature` category,
// this will return a string describing the feature (sting after `feature/`-prefix).
//
// If environment is not a feature environment, will return an empty string.
//
// Returns: string indicating the feature this environment relates to, if not feature environment returns an empty string.
func EnvironmentContext_GetFeatureInfo(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateEnvironmentContext_GetFeatureInfoParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"getFeatureInfo",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Get Environment Label.
//
// Labels are useful since Environment Name can be complex,
// such as `feature/foo-bar` or `qa3`,
// but we need to be able to “label” all `feature/*` and `qaN` environments
// as either `feature` or `qa`.
//
// Returns: Environment Label.
//
// Example:
//   'mock'
//   'development'
//   'feature'
//   'test'
//   'staging'
//   'qa'
//   'preproduction'
//   'production'
//
func EnvironmentContext_GetLabel(scope constructs.Construct) EnvironmentLabel {
	_init_.Initialize()

	if err := validateEnvironmentContext_GetLabelParameters(scope); err != nil {
		panic(err)
	}
	var returns EnvironmentLabel

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"getLabel",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Get Environment Name.
//
// Returns: Environment Name (as given via `--context environment`).
//
// Example:
//   'mock1'
//   'mock2'
//   'mock3'
//   'development'
//   'feature/foo-bar'
//   'feature/ABC-123/new-stuff'
//   'test'
//   'staging'
//   'qa1'
//   'qa2'
//   'qa3'
//   'preproduction'
//   'production'
//
func EnvironmentContext_GetName(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateEnvironmentContext_GetNameParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"getName",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Get Environment URL/DNS Compatible Name.
//
// Returns: Environment URL/DNS Compatible Name (as given via `--context environment` but `param-cased`).
//
// Example:
//   'mock1'
//   'mock2'
//   'mock3'
//   'development'
//   'feature-foo-bar'
//   'feature-abc-123-new-stuff'
//   'test'
//   'staging'
//   'qa1'
//   'qa2'
//   'qa3'
//   'preproduction'
//   'production'
//
func EnvironmentContext_GetUrlName(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateEnvironmentContext_GetUrlNameParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"getUrlName",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Check if Environment is part of `development` category.
//
// Returns true for `development`, otherwise `false`.
//
// Returns: boolean indicating does Environment belong to `development` category.
func EnvironmentContext_IsDevelopment(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateEnvironmentContext_IsDevelopmentParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"isDevelopment",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Check if Environment is part of `feature` category.
//
// Returns `true` for environments with name beginning with `feature/`-prefix, otherwise `false`.
//
// Returns: boolean indicating does Environment belong to `feature` category.
func EnvironmentContext_IsFeature(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateEnvironmentContext_IsFeatureParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"isFeature",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Check if Environment is part of `mock` category.
//
// Returns: boolean indicating does Environment belong to `mock` category.
func EnvironmentContext_IsMock(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateEnvironmentContext_IsMockParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"isMock",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Check if Environment is part of `stable` category.
//
// Returns `true` for `staging` & `production`, otherwise `false`.
//
// Returns: boolean indicating does Environment belong to `stable` category.
func EnvironmentContext_IsStable(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateEnvironmentContext_IsStableParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"isStable",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Check if Environment is part of `verification` category.
//
// Returns `true` for `test` & `preproduction`, otherwise `false`.
//
// Returns: boolean indicating does Environment belong to `verification` category.
func EnvironmentContext_IsVerification(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateEnvironmentContext_IsVerificationParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentContext",
		"isVerification",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

