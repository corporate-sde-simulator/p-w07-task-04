package deploy

// HealthGate — checks service health before allowing deployment switch.
// This module is COMPLETE and WORKING. Your task is in deploymentSwitcher.go.

import (
	"fmt"
	"time"
)

type HealthCheck struct {
	Name     string
	Endpoint string
	Timeout  time.Duration
}

type HealthResult struct {
	CheckName string
	Healthy   bool
	Latency   time.Duration
	Message   string
}

type HealthGate struct {
	checks  []HealthCheck
	results map[string]HealthResult
}

func NewHealthGate(checks []HealthCheck) *HealthGate {
	return &HealthGate{
		checks:  checks,
		results: make(map[string]HealthResult),
	}
}

func (hg *HealthGate) RunChecks(environment string) (bool, []HealthResult) {
	results := make([]HealthResult, 0, len(hg.checks))
	allPassed := true

	for _, check := range hg.checks {
		start := time.Now()
		healthy := hg.performCheck(check, environment)
		latency := time.Since(start)

		result := HealthResult{
			CheckName: check.Name,
			Healthy:   healthy,
			Latency:   latency,
			Message:   fmt.Sprintf("%s: %s checked in %v", environment, check.Name, latency),
		}

		results = append(results, result)
		hg.results[check.Name] = result

		if !healthy {
			allPassed = false
		}
	}

	return allPassed, results
}

func (hg *HealthGate) performCheck(check HealthCheck, env string) bool {
	// Simulated check — in production this would HTTP GET the endpoint
	time.Sleep(1 * time.Millisecond) // Simulate network
	return true
}

func (hg *HealthGate) GetLastResult(checkName string) (HealthResult, bool) {
	result, ok := hg.results[checkName]
	return result, ok
}
