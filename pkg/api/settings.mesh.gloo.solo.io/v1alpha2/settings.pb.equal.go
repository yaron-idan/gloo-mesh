// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/settings/v1alpha2/settings.proto

package v1alpha2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *SettingsSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SettingsSpec)
	if !ok {
		that2, ok := that.(SettingsSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMtls()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMtls()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMtls(), target.GetMtls()) {
			return false
		}
	}

	if len(m.GetNetworkingExtensionServers()) != len(target.GetNetworkingExtensionServers()) {
		return false
	}
	for idx, v := range m.GetNetworkingExtensionServers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetNetworkingExtensionServers()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetNetworkingExtensionServers()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetIstio()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstio()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstio(), target.GetIstio()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *NetworkingExtensionsServer) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*NetworkingExtensionsServer)
	if !ok {
		that2, ok := that.(NetworkingExtensionsServer)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetAddress(), target.GetAddress()) != 0 {
		return false
	}

	if m.GetInsecure() != target.GetInsecure() {
		return false
	}

	if m.GetReconnectOnNetworkFailures() != target.GetReconnectOnNetworkFailures() {
		return false
	}

	return true
}

// Equal function
func (m *SettingsStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SettingsStatus)
	if !ok {
		that2, ok := that.(SettingsStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *SettingsSpec_Istio) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SettingsSpec_Istio)
	if !ok {
		that2, ok := that.(SettingsSpec_Istio)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetIngressGatewayDetectors()) != len(target.GetIngressGatewayDetectors()) {
		return false
	}
	for k, v := range m.GetIngressGatewayDetectors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetIngressGatewayDetectors()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetIngressGatewayDetectors()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *SettingsSpec_Istio_IngressGatewayDetector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SettingsSpec_Istio_IngressGatewayDetector)
	if !ok {
		that2, ok := that.(SettingsSpec_Istio_IngressGatewayDetector)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetGatewayWorkloadLabels()) != len(target.GetGatewayWorkloadLabels()) {
		return false
	}
	for k, v := range m.GetGatewayWorkloadLabels() {

		if strings.Compare(v, target.GetGatewayWorkloadLabels()[k]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetGatewayTlsPortName(), target.GetGatewayTlsPortName()) != 0 {
		return false
	}

	return true
}
