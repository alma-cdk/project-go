// Opinionated CDK Project “Framework”
package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Internal class to handle set/get operations for Account Type.
// Experimental.
type AccountType interface {
}

// The jsii proxy struct for AccountType
type jsiiProxy_AccountType struct {
	_ byte // padding
}

// Experimental.
func NewAccountType() AccountType {
	_init_.Initialize()

	j := jsiiProxy_AccountType{}

	_jsii_.Create(
		"@alma-cdk/project.AccountType",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAccountType_Override(a AccountType) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.AccountType",
		nil, // no parameters
		a,
	)
}

// Experimental.
func AccountType_Get(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateAccountType_GetParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountType",
		"get",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Experimental.
func AccountType_MatchFromEnvironment(scope constructs.Construct, accounts *map[string]*Account, environmentType *string) *string {
	_init_.Initialize()

	if err := validateAccountType_MatchFromEnvironmentParameters(scope, accounts, environmentType); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountType",
		"matchFromEnvironment",
		[]interface{}{scope, accounts, environmentType},
		&returns,
	)

	return returns
}

// Experimental.
func AccountType_Set(scope constructs.Construct, accountType *string) {
	_init_.Initialize()

	if err := validateAccountType_SetParameters(scope, accountType); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@alma-cdk/project.AccountType",
		"set",
		[]interface{}{scope, accountType},
	)
}

