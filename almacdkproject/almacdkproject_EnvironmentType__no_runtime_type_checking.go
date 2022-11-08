//go:build no_runtime_type_checking

// Opinionated CDK Project “Framework”
package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func validateEnvironmentType_GetParameters(scope constructs.Construct, allowedEnvironments *[]*string) error {
	return nil
}

func validateEnvironmentType_SetParameters(scope constructs.Construct, environmentType *string) error {
	return nil
}

func validateEnvironmentType_TryGetParameters(scope constructs.Construct) error {
	return nil
}

