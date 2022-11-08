// Opinionated CDK Project “Framework”
package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type Name interface {
}

// The jsii proxy struct for Name
type jsiiProxy_Name struct {
	_ byte // padding
}

// Experimental.
func NewName_Override(n Name) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.Name",
		nil, // no parameters
		n,
	)
}

// PascalCase naming with global prefixes (org, project…).
// Experimental.
func Name_Globally(scope constructs.Construct, baseName *string, props *NameProps) *string {
	_init_.Initialize()

	if err := validateName_GloballyParameters(scope, baseName, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.Name",
		"globally",
		[]interface{}{scope, baseName, props},
		&returns,
	)

	return returns
}

// Experimental.
func Name_It(scope constructs.Construct, baseName *string, props *NameProps) *string {
	_init_.Initialize()

	if err := validateName_ItParameters(scope, baseName, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.Name",
		"it",
		[]interface{}{scope, baseName, props},
		&returns,
	)

	return returns
}

// Experimental.
func Name_WithProject(scope constructs.Construct, baseName *string, props *NameProps) *string {
	_init_.Initialize()

	if err := validateName_WithProjectParameters(scope, baseName, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.Name",
		"withProject",
		[]interface{}{scope, baseName, props},
		&returns,
	)

	return returns
}

