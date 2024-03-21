package almacdkproject


// Available Environment Labels.
//
// Labels are useful since Environment Name can be complex,
// such as `feature/foo-bar` or `qa3`,
// but we need to be able to “label” all `feature/*` and `qaN` environments
// as either `feature` or `qa`.
// Experimental.
type EnvironmentLabel string

const (
	// Experimental.
	EnvironmentLabel_MOCK EnvironmentLabel = "MOCK"
	// Experimental.
	EnvironmentLabel_DEVELOPMENT EnvironmentLabel = "DEVELOPMENT"
	// Experimental.
	EnvironmentLabel_FEATURE EnvironmentLabel = "FEATURE"
	// Experimental.
	EnvironmentLabel_TEST EnvironmentLabel = "TEST"
	// Experimental.
	EnvironmentLabel_STAGING EnvironmentLabel = "STAGING"
	// Experimental.
	EnvironmentLabel_QA EnvironmentLabel = "QA"
	// Experimental.
	EnvironmentLabel_PREPRODUCTION EnvironmentLabel = "PREPRODUCTION"
	// Experimental.
	EnvironmentLabel_PRODUCTION EnvironmentLabel = "PRODUCTION"
)

