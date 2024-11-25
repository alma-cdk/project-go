package almacdkproject


// Available Environment Labels.
//
// Labels are useful since Environment Name can be complex,
// such as `feature/foo-bar` or `qa3`,
// but we need to be able to “label” all `feature/*` and `qaN` environments
// as either `feature` or `qa`.
type EnvironmentLabel string

const (
	EnvironmentLabel_MOCK EnvironmentLabel = "MOCK"
	EnvironmentLabel_DEVELOPMENT EnvironmentLabel = "DEVELOPMENT"
	EnvironmentLabel_FEATURE EnvironmentLabel = "FEATURE"
	EnvironmentLabel_TEST EnvironmentLabel = "TEST"
	EnvironmentLabel_STAGING EnvironmentLabel = "STAGING"
	EnvironmentLabel_QA EnvironmentLabel = "QA"
	EnvironmentLabel_PREPRODUCTION EnvironmentLabel = "PREPRODUCTION"
	EnvironmentLabel_PRODUCTION EnvironmentLabel = "PRODUCTION"
)

