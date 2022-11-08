//go:build no_runtime_type_checking

// Opinionated CDK Project “Framework”
package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func validateAccountContext_GetAccountConfigParameters(scope constructs.Construct, key *string) error {
	return nil
}

func validateAccountContext_GetAccountIdParameters(scope constructs.Construct) error {
	return nil
}

func validateAccountContext_GetAccountTypeParameters(scope constructs.Construct) error {
	return nil
}

func validateAccountContext_IsDevParameters(scope constructs.Construct) error {
	return nil
}

func validateAccountContext_IsMockParameters(scope constructs.Construct) error {
	return nil
}

func validateAccountContext_IsPreProdParameters(scope constructs.Construct) error {
	return nil
}

func validateAccountContext_IsProdParameters(scope constructs.Construct) error {
	return nil
}

func validateAccountContext_IsSharedParameters(scope constructs.Construct) error {
	return nil
}

