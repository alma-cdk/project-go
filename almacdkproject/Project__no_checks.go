//go:build no_runtime_type_checking

package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Project) validateSynthParameters(options *awscdk.StageSynthesisOptions) error {
	return nil
}

func validateProject_GetAccountParameters(scope constructs.Construct, accountType *string) error {
	return nil
}

func validateProject_GetConfigurationParameters(scope constructs.Construct) error {
	return nil
}

func validateProject_IsAppParameters(obj interface{}) error {
	return nil
}

func validateProject_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProject_IsStageParameters(x interface{}) error {
	return nil
}

func validateProject_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewProjectParameters(props *ProjectProps) error {
	return nil
}

