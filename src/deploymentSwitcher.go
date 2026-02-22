package deploy

// DeploymentSwitcher — manages blue-green deployments.
//
// YOU MUST IMPLEMENT the methods marked with TODO.
// HealthGate is working — use it to validate before switching.

import (
	"fmt"
	"time"
)

type Environment string

const (
	Blue  Environment = "blue"
	Green Environment = "green"
)

type DeploymentState struct {
	ActiveEnv    Environment
	InactiveEnv  Environment
	Version      string
	DeployedAt   time.Time
	SwitchedAt   *time.Time
}

type DeploymentRecord struct {
	FromEnv    Environment
	ToEnv      Environment
	Version    string
	Success    bool
	Timestamp  time.Time
	Message    string
}

type DeploymentSwitcher struct {
	healthGate  *HealthGate
	activeEnv   Environment
	versions    map[Environment]string
	history     []DeploymentRecord
}

func NewDeploymentSwitcher(gate *HealthGate) *DeploymentSwitcher {
	return &DeploymentSwitcher{
		healthGate: gate,
		activeEnv:  Blue,
		versions:   map[Environment]string{Blue: "v1.0.0", Green: ""},
		history:    make([]DeploymentRecord, 0),
	}
}

// Deploy deploys a new version to the INACTIVE environment.
//
// 1. Determine which environment is inactive (opposite of activeEnv)
// 2. Set the version for the inactive environment
// 3. Record a DeploymentRecord with Success=true
// 4. Return the inactive environment name and nil error
// 5. Return error if version string is empty
func (ds *DeploymentSwitcher) Deploy(version string) (Environment, error) {
	return "", fmt.Errorf("not implemented")
}

// Switch routes traffic to the inactive environment after health checks pass.
//
// 1. Get the inactive environment
// 2. Check that it has a deployed version (not empty)
// 3. Run health checks using healthGate.RunChecks(string(inactiveEnv))
// 4. If health checks fail, return error with details
// 5. If health checks pass, swap activeEnv to the inactive one
// 6. Record the switch in history
func (ds *DeploymentSwitcher) Switch() error {
	return fmt.Errorf("not implemented")
}

// Rollback reverts to the previous active environment.
//
// 1. Swap activeEnv back to the other environment
// 2. Record the rollback in history
// 3. Return error if there's no previous deployment to roll back to
func (ds *DeploymentSwitcher) Rollback() error {
	return fmt.Errorf("not implemented")
}

// GetActiveEnvironment returns the currently active environment.
func (ds *DeploymentSwitcher) GetActiveEnvironment() Environment {
	return ds.activeEnv
}

// GetVersion returns the deployed version for an environment.
func (ds *DeploymentSwitcher) GetVersion(env Environment) string {
	return ds.versions[env]
}

// GetHistory returns the deployment history.
func (ds *DeploymentSwitcher) GetHistory() []DeploymentRecord {
	return ds.history
}
