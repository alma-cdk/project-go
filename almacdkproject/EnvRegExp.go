package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type EnvRegExp interface {
	// Experimental.
	Test(value *string) *bool
}

// The jsii proxy struct for EnvRegExp
type jsiiProxy_EnvRegExp struct {
	_ byte // padding
}

// Experimental.
func NewEnvRegExp(base *string) EnvRegExp {
	_init_.Initialize()

	if err := validateNewEnvRegExpParameters(base); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnvRegExp{}

	_jsii_.Create(
		"@alma-cdk/project.EnvRegExp",
		[]interface{}{base},
		&j,
	)

	return &j
}

// Experimental.
func NewEnvRegExp_Override(e EnvRegExp, base *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.EnvRegExp",
		[]interface{}{base},
		e,
	)
}

func (e *jsiiProxy_EnvRegExp) Test(value *string) *bool {
	if err := e.validateTestParameters(value); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		e,
		"test",
		[]interface{}{value},
		&returns,
	)

	return returns
}

