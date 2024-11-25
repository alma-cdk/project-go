package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type PathName interface {
	UrlName
}

// The jsii proxy struct for PathName
type jsiiProxy_PathName struct {
	jsiiProxy_UrlName
}

func NewPathName_Override(p PathName) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.PathName",
		nil, // no parameters
		p,
	)
}

// PascalCase naming with global prefixes (org, project…).
func PathName_Globally(scope constructs.Construct, baseName *string, props *NameProps) *string {
	_init_.Initialize()

	if err := validatePathName_GloballyParameters(scope, baseName, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.PathName",
		"globally",
		[]interface{}{scope, baseName, props},
		&returns,
	)

	return returns
}

func PathName_It(scope constructs.Construct, baseName *string, props *NameProps) *string {
	_init_.Initialize()

	if err := validatePathName_ItParameters(scope, baseName, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.PathName",
		"it",
		[]interface{}{scope, baseName, props},
		&returns,
	)

	return returns
}

func PathName_WithProject(scope constructs.Construct, baseName *string, props *NameProps) *string {
	_init_.Initialize()

	if err := validatePathName_WithProjectParameters(scope, baseName, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.PathName",
		"withProject",
		[]interface{}{scope, baseName, props},
		&returns,
	)

	return returns
}

