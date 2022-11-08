//go:build no_runtime_type_checking

// Opinionated CDK Project “Framework”
package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func validateEnvironmentWrapper_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEnvironmentWrapperParameters(scope constructs.Construct) error {
	return nil
}

