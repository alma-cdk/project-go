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

