/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsStepDetailWarnings TODO.
type SyntheticsStepDetailWarnings struct {
	// TODO.
	Message string                `json:"message"`
	Type    SyntheticsWarningType `json:"type"`
}

// NewSyntheticsStepDetailWarnings instantiates a new SyntheticsStepDetailWarnings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsStepDetailWarnings(message string, type_ SyntheticsWarningType) *SyntheticsStepDetailWarnings {
	this := SyntheticsStepDetailWarnings{}
	this.Message = message
	this.Type = type_
	return &this
}

// NewSyntheticsStepDetailWarningsWithDefaults instantiates a new SyntheticsStepDetailWarnings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsStepDetailWarningsWithDefaults() *SyntheticsStepDetailWarnings {
	this := SyntheticsStepDetailWarnings{}
	return &this
}

// GetMessage returns the Message field value
func (o *SyntheticsStepDetailWarnings) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetailWarnings) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SyntheticsStepDetailWarnings) SetMessage(v string) {
	o.Message = v
}

// GetType returns the Type field value
func (o *SyntheticsStepDetailWarnings) GetType() SyntheticsWarningType {
	if o == nil {
		var ret SyntheticsWarningType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsStepDetailWarnings) GetTypeOk() (*SyntheticsWarningType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntheticsStepDetailWarnings) SetType(v SyntheticsWarningType) {
	o.Type = v
}

func (o SyntheticsStepDetailWarnings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsStepDetailWarnings struct {
	value *SyntheticsStepDetailWarnings
	isSet bool
}

func (v NullableSyntheticsStepDetailWarnings) Get() *SyntheticsStepDetailWarnings {
	return v.value
}

func (v *NullableSyntheticsStepDetailWarnings) Set(val *SyntheticsStepDetailWarnings) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsStepDetailWarnings) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsStepDetailWarnings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsStepDetailWarnings(val *SyntheticsStepDetailWarnings) *NullableSyntheticsStepDetailWarnings {
	return &NullableSyntheticsStepDetailWarnings{value: val, isSet: true}
}

func (v NullableSyntheticsStepDetailWarnings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsStepDetailWarnings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
