//go:build no_runtime_type_checking

// Opinionated CDK Project “Framework”
package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func validateAccountType_GetParameters(scope constructs.Construct) error {
	return nil
}

func validateAccountType_MatchFromEnvironmentParameters(scope constructs.Construct, accounts *map[string]*Account, environmentType *string) error {
	return nil
}

func validateAccountType_SetParameters(scope constructs.Construct, accountType *string) error {
	return nil
}

