//go:build no_runtime_type_checking

// Opinionated CDK Project “Framework”
package almacdkproject

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SmartStack) validateAddDependencyParameters(target awscdk.Stack) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateAddTransformParameters(transform *string) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateAllocateLogicalIdParameters(cfnElement awscdk.CfnElement) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateExportValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateFormatArnParameters(components *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateGetLogicalIdParameters(element awscdk.CfnElement) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateRegionalFactParameters(factName *string) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateResolveParameters(obj interface{}) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateSplitArnParameters(arn *string, arnFormat awscdk.ArnFormat) error {
	return nil
}

func (s *jsiiProxy_SmartStack) validateToJsonStringParameters(obj interface{}) error {
	return nil
}

func validateSmartStack_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSmartStack_IsStackParameters(x interface{}) error {
	return nil
}

func validateSmartStack_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSmartStackParameters(scope constructs.Construct, id *string, props *awscdk.StackProps) error {
	return nil
}

