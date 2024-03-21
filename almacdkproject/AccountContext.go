package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type AccountContext interface {
}

// The jsii proxy struct for AccountContext
type jsiiProxy_AccountContext struct {
	_ byte // padding
}

// Experimental.
func NewAccountContext() AccountContext {
	_init_.Initialize()

	j := jsiiProxy_AccountContext{}

	_jsii_.Create(
		"@alma-cdk/project.AccountContext",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAccountContext_Override(a AccountContext) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.AccountContext",
		nil, // no parameters
		a,
	)
}

// Experimental.
func AccountContext_GetAccountConfig(scope constructs.Construct, key *string) interface{} {
	_init_.Initialize()

	if err := validateAccountContext_GetAccountConfigParameters(scope, key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountContext",
		"getAccountConfig",
		[]interface{}{scope, key},
		&returns,
	)

	return returns
}

// Experimental.
func AccountContext_GetAccountId(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateAccountContext_GetAccountIdParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountContext",
		"getAccountId",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Experimental.
func AccountContext_GetAccountType(scope constructs.Construct) *string {
	_init_.Initialize()

	if err := validateAccountContext_GetAccountTypeParameters(scope); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountContext",
		"getAccountType",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Experimental.
func AccountContext_IsDev(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateAccountContext_IsDevParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountContext",
		"isDev",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Experimental.
func AccountContext_IsMock(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateAccountContext_IsMockParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountContext",
		"isMock",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Experimental.
func AccountContext_IsPreProd(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateAccountContext_IsPreProdParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountContext",
		"isPreProd",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Experimental.
func AccountContext_IsProd(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateAccountContext_IsProdParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountContext",
		"isProd",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Experimental.
func AccountContext_IsShared(scope constructs.Construct) *bool {
	_init_.Initialize()

	if err := validateAccountContext_IsSharedParameters(scope); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountContext",
		"isShared",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

