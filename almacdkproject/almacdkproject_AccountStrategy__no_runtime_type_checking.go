//go:build no_runtime_type_checking

// Opinionated CDK Project “Framework”
package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func validateAccountStrategy_OneParameters(props *AccountStrategyOneProps) error {
	return nil
}

func validateAccountStrategy_ThreeParameters(props *AccountStrategyThreeProps) error {
	return nil
}

func validateAccountStrategy_TwoParameters(props *AccountStrategyTwoProps) error {
	return nil
}

