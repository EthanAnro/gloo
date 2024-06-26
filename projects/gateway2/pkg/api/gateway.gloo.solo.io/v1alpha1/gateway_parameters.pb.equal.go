// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/gateway_parameters.proto

package v1alpha1

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
func (m *GatewayParametersSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayParametersSpec)
	if !ok {
		that2, ok := that.(GatewayParametersSpec)
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

	switch m.EnvironmentType.(type) {

	case *GatewayParametersSpec_Kube:
		if _, ok := target.EnvironmentType.(*GatewayParametersSpec_Kube); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKube()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKube()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKube(), target.GetKube()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.EnvironmentType != target.EnvironmentType {
			return false
		}
	}

	return true
}

// Equal function
func (m *KubernetesProxyConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*KubernetesProxyConfig)
	if !ok {
		that2, ok := that.(KubernetesProxyConfig)
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

	if h, ok := interface{}(m.GetEnvoyContainer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEnvoyContainer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEnvoyContainer(), target.GetEnvoyContainer()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPodTemplate()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPodTemplate()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPodTemplate(), target.GetPodTemplate()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetService(), target.GetService()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAutoscaling()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAutoscaling()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAutoscaling(), target.GetAutoscaling()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSds()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSds()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSds(), target.GetSds()) {
			return false
		}
	}

	switch m.WorkloadType.(type) {

	case *KubernetesProxyConfig_Deployment:
		if _, ok := target.WorkloadType.(*KubernetesProxyConfig_Deployment); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDeployment()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDeployment()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDeployment(), target.GetDeployment()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.WorkloadType != target.WorkloadType {
			return false
		}
	}

	return true
}

// Equal function
func (m *ProxyDeployment) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ProxyDeployment)
	if !ok {
		that2, ok := that.(ProxyDeployment)
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

	if h, ok := interface{}(m.GetReplicas()).(equality.Equalizer); ok {
		if !h.Equal(target.GetReplicas()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetReplicas(), target.GetReplicas()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *EnvoyContainer) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*EnvoyContainer)
	if !ok {
		that2, ok := that.(EnvoyContainer)
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

	if h, ok := interface{}(m.GetBootstrap()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBootstrap()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBootstrap(), target.GetBootstrap()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetImage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetImage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetImage(), target.GetImage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecurityContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecurityContext(), target.GetSecurityContext()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResources()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResources()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResources(), target.GetResources()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *EnvoyBootstrap) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*EnvoyBootstrap)
	if !ok {
		that2, ok := that.(EnvoyBootstrap)
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

	if strings.Compare(m.GetLogLevel(), target.GetLogLevel()) != 0 {
		return false
	}

	if len(m.GetComponentLogLevels()) != len(target.GetComponentLogLevels()) {
		return false
	}
	for k, v := range m.GetComponentLogLevels() {

		if strings.Compare(v, target.GetComponentLogLevels()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *SdsIntegration) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SdsIntegration)
	if !ok {
		that2, ok := that.(SdsIntegration)
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

	if h, ok := interface{}(m.GetSdsContainer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSdsContainer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSdsContainer(), target.GetSdsContainer()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIstioIntegration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstioIntegration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstioIntegration(), target.GetIstioIntegration()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *IstioIntegration) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioIntegration)
	if !ok {
		that2, ok := that.(IstioIntegration)
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

	if h, ok := interface{}(m.GetIstioContainer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstioContainer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstioContainer(), target.GetIstioContainer()) {
			return false
		}
	}

	if strings.Compare(m.GetIstioDiscoveryAddress(), target.GetIstioDiscoveryAddress()) != 0 {
		return false
	}

	if strings.Compare(m.GetIstioMetaMeshId(), target.GetIstioMetaMeshId()) != 0 {
		return false
	}

	if strings.Compare(m.GetIstioMetaClusterId(), target.GetIstioMetaClusterId()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *SdsContainer) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SdsContainer)
	if !ok {
		that2, ok := that.(SdsContainer)
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

	if h, ok := interface{}(m.GetImage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetImage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetImage(), target.GetImage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecurityContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecurityContext(), target.GetSecurityContext()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResources()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResources()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResources(), target.GetResources()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetBootstrap()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBootstrap()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBootstrap(), target.GetBootstrap()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *SdsBootstrap) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SdsBootstrap)
	if !ok {
		that2, ok := that.(SdsBootstrap)
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

	if strings.Compare(m.GetLogLevel(), target.GetLogLevel()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *IstioContainer) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioContainer)
	if !ok {
		that2, ok := that.(IstioContainer)
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

	if h, ok := interface{}(m.GetImage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetImage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetImage(), target.GetImage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecurityContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecurityContext(), target.GetSecurityContext()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResources()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResources()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResources(), target.GetResources()) {
			return false
		}
	}

	if strings.Compare(m.GetLogLevel(), target.GetLogLevel()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GatewayParametersStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayParametersStatus)
	if !ok {
		that2, ok := that.(GatewayParametersStatus)
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

	return true
}
