/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kratos

import (
	"encoding/json"
)

// RecoveryFlowMethod struct for RecoveryFlowMethod
type RecoveryFlowMethod struct {
	Config RecoveryFlowMethodConfig `json:"config"`
	// Method contains the request credentials type.
	Method string `json:"method"`
}

// NewRecoveryFlowMethod instantiates a new RecoveryFlowMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryFlowMethod(config RecoveryFlowMethodConfig, method string) *RecoveryFlowMethod {
	this := RecoveryFlowMethod{}
	this.Config = config
	this.Method = method
	return &this
}

// NewRecoveryFlowMethodWithDefaults instantiates a new RecoveryFlowMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryFlowMethodWithDefaults() *RecoveryFlowMethod {
	this := RecoveryFlowMethod{}
	return &this
}

// GetConfig returns the Config field value
func (o *RecoveryFlowMethod) GetConfig() RecoveryFlowMethodConfig {
	if o == nil {
		var ret RecoveryFlowMethodConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlowMethod) GetConfigOk() (*RecoveryFlowMethodConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *RecoveryFlowMethod) SetConfig(v RecoveryFlowMethodConfig) {
	o.Config = v
}

// GetMethod returns the Method field value
func (o *RecoveryFlowMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlowMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *RecoveryFlowMethod) SetMethod(v string) {
	o.Method = v
}

func (o RecoveryFlowMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryFlowMethod struct {
	value *RecoveryFlowMethod
	isSet bool
}

func (v NullableRecoveryFlowMethod) Get() *RecoveryFlowMethod {
	return v.value
}

func (v *NullableRecoveryFlowMethod) Set(val *RecoveryFlowMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryFlowMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryFlowMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryFlowMethod(val *RecoveryFlowMethod) *NullableRecoveryFlowMethod {
	return &NullableRecoveryFlowMethod{value: val, isSet: true}
}

func (v NullableRecoveryFlowMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryFlowMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
