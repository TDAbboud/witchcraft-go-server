// This file was generated by Conjure and should not be manually edited.

package health

import (
	"encoding/json"
)

type HealthStatus struct {
	Checks map[CheckType]HealthCheckResult `json:"checks" yaml:"checks,omitempty"`
}

func (o HealthStatus) MarshalJSON() ([]byte, error) {
	if o.Checks == nil {
		o.Checks = make(map[CheckType]HealthCheckResult, 0)
	}
	type HealthStatusAlias HealthStatus
	return json.Marshal(HealthStatusAlias(o))
}

func (o *HealthStatus) UnmarshalJSON(data []byte) error {
	type HealthStatusAlias HealthStatus
	var rawHealthStatus HealthStatusAlias
	if err := json.Unmarshal(data, &rawHealthStatus); err != nil {
		return err
	}
	if rawHealthStatus.Checks == nil {
		rawHealthStatus.Checks = make(map[CheckType]HealthCheckResult, 0)
	}
	*o = HealthStatus(rawHealthStatus)
	return nil
}

func (o HealthStatus) MarshalYAML() (interface{}, error) {
	if o.Checks == nil {
		o.Checks = make(map[CheckType]HealthCheckResult, 0)
	}
	type HealthStatusAlias HealthStatus
	return HealthStatusAlias(o), nil
}

func (o *HealthStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type HealthStatusAlias HealthStatus
	var rawHealthStatus HealthStatusAlias
	if err := unmarshal(&rawHealthStatus); err != nil {
		return err
	}
	if rawHealthStatus.Checks == nil {
		rawHealthStatus.Checks = make(map[CheckType]HealthCheckResult, 0)
	}
	*o = HealthStatus(rawHealthStatus)
	return nil
}

// Metadata describing the status of a service.
type HealthCheckResult struct {
	// A constant representing the type of health check. Values should be uppercase, underscore delimited, ascii letters with no spaces, ([A-Z_]).
	Type CheckType `json:"type" yaml:"type,omitempty" conjure-docs:"A constant representing the type of health check. Values should be uppercase, underscore delimited, ascii letters with no spaces, ([A-Z_]).\n"`
	// Health state of the check.
	State HealthState `json:"state" yaml:"state,omitempty" conjure-docs:"Health state of the check.\n"`
	// Text describing the state of the check which should provide enough information for the check to be actionable when included in an alert.
	Message *string `json:"message" yaml:"message,omitempty" conjure-docs:"Text describing the state of the check which should provide enough information for the check to be actionable when included in an alert.\n"`
	// Additional redacted information on the nature of the health check.
	Params map[string]interface{} `json:"params" yaml:"params,omitempty" conjure-docs:"Additional redacted information on the nature of the health check.\n"`
}

func (o HealthCheckResult) MarshalJSON() ([]byte, error) {
	if o.Params == nil {
		o.Params = make(map[string]interface{}, 0)
	}
	type HealthCheckResultAlias HealthCheckResult
	return json.Marshal(HealthCheckResultAlias(o))
}

func (o *HealthCheckResult) UnmarshalJSON(data []byte) error {
	type HealthCheckResultAlias HealthCheckResult
	var rawHealthCheckResult HealthCheckResultAlias
	if err := json.Unmarshal(data, &rawHealthCheckResult); err != nil {
		return err
	}
	if rawHealthCheckResult.Params == nil {
		rawHealthCheckResult.Params = make(map[string]interface{}, 0)
	}
	*o = HealthCheckResult(rawHealthCheckResult)
	return nil
}

func (o HealthCheckResult) MarshalYAML() (interface{}, error) {
	if o.Params == nil {
		o.Params = make(map[string]interface{}, 0)
	}
	type HealthCheckResultAlias HealthCheckResult
	return HealthCheckResultAlias(o), nil
}

func (o *HealthCheckResult) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type HealthCheckResultAlias HealthCheckResult
	var rawHealthCheckResult HealthCheckResultAlias
	if err := unmarshal(&rawHealthCheckResult); err != nil {
		return err
	}
	if rawHealthCheckResult.Params == nil {
		rawHealthCheckResult.Params = make(map[string]interface{}, 0)
	}
	*o = HealthCheckResult(rawHealthCheckResult)
	return nil
}