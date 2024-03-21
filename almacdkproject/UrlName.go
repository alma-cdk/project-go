package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type UrlName interface {
	Name
}

// The jsii proxy struct for UrlName
type jsiiProxy_UrlName struct {
	jsiiProxy_Name
}

// Experimental.
func NewUrlName_Override(u UrlName) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.UrlName",
		nil, // no parameters
		u,
	)
}

// PascalCase naming with global prefixes (org, project…).
// Experimental.
func UrlName_Globally(scope constructs.Construct, baseName *string, props *NameProps) *string {
	_init_.Initialize()

	if err := validateUrlName_GloballyParameters(scope, baseName, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.UrlName",
		"globally",
		[]interface{}{scope, baseName, props},
		&returns,
	)

	return returns
}

// Experimental.
func UrlName_It(scope constructs.Construct, baseName *string, props *NameProps) *string {
	_init_.Initialize()

	if err := validateUrlName_ItParameters(scope, baseName, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.UrlName",
		"it",
		[]interface{}{scope, baseName, props},
		&returns,
	)

	return returns
}

// Experimental.
func UrlName_WithProject(scope constructs.Construct, baseName *string, props *NameProps) *string {
	_init_.Initialize()

	if err := validateUrlName_WithProjectParameters(scope, baseName, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.UrlName",
		"withProject",
		[]interface{}{scope, baseName, props},
		&returns,
	)

	return returns
}

