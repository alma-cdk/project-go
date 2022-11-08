//go:build no_runtime_type_checking

// Opinionated CDK Project “Framework”
package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EnvRegExp) validateTestParameters(value *string) error {
	return nil
}

func validateNewEnvRegExpParameters(base *string) error {
	return nil
}

