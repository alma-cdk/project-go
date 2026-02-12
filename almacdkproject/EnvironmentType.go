package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Internal class to handle set/get operations for Environment Type.
type EnvironmentType interface {
}

// The jsii proxy struct for EnvironmentType
type jsiiProxy_EnvironmentType struct {
	_ byte // padding
}

func NewEnvironmentType() EnvironmentType {
	_init_.Initialize()

	j := jsiiProxy_EnvironmentType{}

	_jsii_.Create(
		"@alma-cdk/project.EnvironmentType",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewEnvironmentType_Override(e EnvironmentType) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.EnvironmentType",
		nil, // no parameters
		e,
	)
}

func EnvironmentType_Get(scope constructs.Construct, allowedEnvironments *[]*string) *string {
	_init_.Initialize()

	if err := validateEnvironmentType_GetParameters(scope, allowedEnvironments); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentType",
		"get",
		[]interface{}{scope, allowedEnvironments},
		&returns,
	)

	return returns
}

func EnvironmentType_Set(scope constructs.Construct, environmentType *string) {
	_init_.Initialize()

	if err := validateEnvironmentType_SetParameters(scope, environmentType); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@alma-cdk/project.EnvironmentType",
		"set",
		[]interface{}{scope, environmentType},
	)
}

func EnvironmentType_TryGet(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateEnvironmentType_TryGetParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentType",
		"tryGet",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

