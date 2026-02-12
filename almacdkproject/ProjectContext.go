package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type ProjectContext interface {
}

// The jsii proxy struct for ProjectContext
type jsiiProxy_ProjectContext struct {
	_ byte // padding
}

func NewProjectContext() ProjectContext {
	_init_.Initialize()

	j := jsiiProxy_ProjectContext{}

	_jsii_.Create(
		"@alma-cdk/project.ProjectContext",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewProjectContext_Override(p ProjectContext) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.ProjectContext",
		nil, // no parameters
		p,
	)
}

func ProjectContext_GetAccountConfig(scope constructs.Construct, key *string, defaultValue interface{}) interface{} {
	_init_.Initialize()

	if err := validateProjectContext_GetAccountConfigParameters(scope, key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getAccountConfig",
		[]interface{}{scope, key, defaultValue},
		&returns,
	)

	return returns
}

func ProjectContext_GetAccountId(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateProjectContext_GetAccountIdParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getAccountId",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Returns the account type given in runtime/CLI context.
func ProjectContext_GetAccountType(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateProjectContext_GetAccountTypeParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getAccountType",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func ProjectContext_GetAllowedEnvironments(scope constructs.Construct) *[]*string {
	_init_.Initialize()

	if err := validateProjectContext_GetAllowedEnvironmentsParameters(scope); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getAllowedEnvironments",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func ProjectContext_GetAuthorEmail(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateProjectContext_GetAuthorEmailParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getAuthorEmail",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func ProjectContext_GetAuthorName(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateProjectContext_GetAuthorNameParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getAuthorName",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func ProjectContext_GetAuthorOrganization(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateProjectContext_GetAuthorOrganizationParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getAuthorOrganization",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func ProjectContext_GetDefaultRegion(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateProjectContext_GetDefaultRegionParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getDefaultRegion",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func ProjectContext_GetEnvironment(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateProjectContext_GetEnvironmentParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getEnvironment",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func ProjectContext_GetName(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateProjectContext_GetNameParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"getName",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func ProjectContext_TryGetEnvironment(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateProjectContext_TryGetEnvironmentParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.ProjectContext",
		"tryGetEnvironment",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

