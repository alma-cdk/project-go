package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alma-cdk/project-go/almacdkproject/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Wrapper for environmental stacks.
type EnvironmentWrapper interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EnvironmentWrapper
type jsiiProxy_EnvironmentWrapper struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EnvironmentWrapper) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEnvironmentWrapper(scope constructs.Construct) EnvironmentWrapper {
	_init_.Initialize()

	if err := validateNewEnvironmentWrapperParameters(scope); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnvironmentWrapper{}

	_jsii_.Create(
		"@alma-cdk/project.EnvironmentWrapper",
		[]interface{}{scope},
		&j,
	)

	return &j
}

func NewEnvironmentWrapper_Override(e EnvironmentWrapper, scope constructs.Construct) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.EnvironmentWrapper",
		[]interface{}{scope},
		e,
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
func EnvironmentWrapper_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnvironmentWrapper_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/project.EnvironmentWrapper",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentWrapper) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

