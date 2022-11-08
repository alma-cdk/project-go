// Opinionated CDK Project “Framework”
package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alma-cdk/project-go/almacdkproject/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Wrapper for account-level stacks.
// Experimental.
type AccountWrapper interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountWrapper
type jsiiProxy_AccountWrapper struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AccountWrapper) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAccountWrapper(scope constructs.Construct) AccountWrapper {
	_init_.Initialize()

	if err := validateNewAccountWrapperParameters(scope); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountWrapper{}

	_jsii_.Create(
		"@alma-cdk/project.AccountWrapper",
		[]interface{}{scope},
		&j,
	)

	return &j
}

// Experimental.
func NewAccountWrapper_Override(a AccountWrapper, scope constructs.Construct) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.AccountWrapper",
		[]interface{}{scope},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AccountWrapper_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccountWrapper_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountWrapper",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountWrapper) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

