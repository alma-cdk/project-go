package almacdkproject

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@alma-cdk/project.Account",
		reflect.TypeOf((*Account)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/project.AccountConfiguration",
		reflect.TypeOf((*AccountConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.AccountContext",
		reflect.TypeOf((*AccountContext)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AccountContext{}
		},
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.AccountStrategy",
		reflect.TypeOf((*AccountStrategy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AccountStrategy{}
		},
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/project.AccountStrategyOneProps",
		reflect.TypeOf((*AccountStrategyOneProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/project.AccountStrategyThreeProps",
		reflect.TypeOf((*AccountStrategyThreeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/project.AccountStrategyTwoProps",
		reflect.TypeOf((*AccountStrategyTwoProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.AccountType",
		reflect.TypeOf((*AccountType)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AccountType{}
		},
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.AccountWrapper",
		reflect.TypeOf((*AccountWrapper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AccountWrapper{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/project.Author",
		reflect.TypeOf((*Author)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.EnvRegExp",
		reflect.TypeOf((*EnvRegExp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "test", GoMethod: "Test"},
		},
		func() interface{} {
			return &jsiiProxy_EnvRegExp{}
		},
	)
	_jsii_.RegisterEnum(
		"@alma-cdk/project.EnvironmentCategory",
		reflect.TypeOf((*EnvironmentCategory)(nil)).Elem(),
		map[string]interface{}{
			"MOCK": EnvironmentCategory_MOCK,
			"DEVELOPMENT": EnvironmentCategory_DEVELOPMENT,
			"FEATURE": EnvironmentCategory_FEATURE,
			"VERIFICATION": EnvironmentCategory_VERIFICATION,
			"STABLE": EnvironmentCategory_STABLE,
		},
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.EnvironmentContext",
		reflect.TypeOf((*EnvironmentContext)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EnvironmentContext{}
		},
	)
	_jsii_.RegisterEnum(
		"@alma-cdk/project.EnvironmentLabel",
		reflect.TypeOf((*EnvironmentLabel)(nil)).Elem(),
		map[string]interface{}{
			"MOCK": EnvironmentLabel_MOCK,
			"DEVELOPMENT": EnvironmentLabel_DEVELOPMENT,
			"FEATURE": EnvironmentLabel_FEATURE,
			"TEST": EnvironmentLabel_TEST,
			"STAGING": EnvironmentLabel_STAGING,
			"QA": EnvironmentLabel_QA,
			"PREPRODUCTION": EnvironmentLabel_PREPRODUCTION,
			"PRODUCTION": EnvironmentLabel_PRODUCTION,
		},
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.EnvironmentType",
		reflect.TypeOf((*EnvironmentType)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EnvironmentType{}
		},
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.EnvironmentWrapper",
		reflect.TypeOf((*EnvironmentWrapper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EnvironmentWrapper{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.Name",
		reflect.TypeOf((*Name)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Name{}
		},
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/project.NameProps",
		reflect.TypeOf((*NameProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.PathName",
		reflect.TypeOf((*PathName)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			j := jsiiProxy_PathName{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_UrlName)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.Project",
		reflect.TypeOf((*Project)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "assetOutdir", GoGetter: "AssetOutdir"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "parentStage", GoGetter: "ParentStage"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Project{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkApp)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/project.ProjectConfiguration",
		reflect.TypeOf((*ProjectConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.ProjectContext",
		reflect.TypeOf((*ProjectContext)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ProjectContext{}
		},
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/project.ProjectProps",
		reflect.TypeOf((*ProjectProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.SmartStack",
		reflect.TypeOf((*SmartStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addTransform", GoMethod: "AddTransform"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "bundlingRequired", GoGetter: "BundlingRequired"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "exportValue", GoMethod: "ExportValue"},
			_jsii_.MemberMethod{JsiiMethod: "formatArn", GoMethod: "FormatArn"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "nested", GoGetter: "Nested"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackParent", GoGetter: "NestedStackParent"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackResource", GoGetter: "NestedStackResource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "regionalFact", GoMethod: "RegionalFact"},
			_jsii_.MemberMethod{JsiiMethod: "renameLogicalId", GoMethod: "RenameLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContextKey", GoMethod: "ReportMissingContextKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "splitArn", GoMethod: "SplitArn"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateFile", GoGetter: "TemplateFile"},
			_jsii_.MemberProperty{JsiiProperty: "templateOptions", GoGetter: "TemplateOptions"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_SmartStack{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkStack)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@alma-cdk/project.UrlName",
		reflect.TypeOf((*UrlName)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			j := jsiiProxy_UrlName{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Name)
			return &j
		},
	)
}
