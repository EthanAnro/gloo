package deployer

import (
	extcorev1 "github.com/solo-io/gloo/projects/gateway2/pkg/api/external/kubernetes/api/core/v1"
	v1alpha1kube "github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1/kube"
)

// The top-level helm values used by the deployer.
type helmConfig struct {
	Gateway *helmGateway `json:"gateway,omitempty"`
}

type helmGateway struct {
	// naming
	Name             *string `json:"name,omitempty"`
	GatewayName      *string `json:"gatewayName,omitempty"`
	GatewayNamespace *string `json:"gatewayNamespace,omitempty"`
	NameOverride     *string `json:"nameOverride,omitempty"`
	FullnameOverride *string `json:"fullnameOverride,omitempty"`

	// deployment/service values
	ReplicaCount *uint32          `json:"replicaCount,omitempty"`
	Autoscaling  *helmAutoscaling `json:"autoscaling,omitempty"`
	Ports        []helmPort       `json:"ports,omitempty"`
	// TODO: This is unused
	ReadinessPort *uint16      `json:"readinessPort,omitempty"`
	Service       *helmService `json:"service,omitempty"`

	// pod template values
	ExtraPodAnnotations map[string]string                 `json:"extraPodAnnotations,omitempty"`
	ExtraPodLabels      map[string]string                 `json:"extraPodLabels,omitempty"`
	ImagePullSecrets    []*extcorev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	PodSecurityContext  *extcorev1.PodSecurityContext     `json:"podSecurityContext,omitempty"`
	NodeSelector        map[string]string                 `json:"nodeSelector,omitempty"`
	Affinity            *extcorev1.Affinity               `json:"affinity,omitempty"`
	Tolerations         []*extcorev1.Toleration           `json:"tolerations,omitempty"`

	// sds values
	Sds *helmSds `json:"sds,omitempty"`

	// envoy container values
	LogLevel          *string                            `json:"logLevel,omitempty"`
	ComponentLogLevel *string                            `json:"componentLogLevel,omitempty"`
	Image             *helmImage                         `json:"image,omitempty"`
	Resources         *v1alpha1kube.ResourceRequirements `json:"resources,omitempty"`
	SecurityContext   *extcorev1.SecurityContext         `json:"securityContext,omitempty"`

	// xds values
	Xds *helmXds `json:"xds,omitempty"`

	// serviceaccount values
	ServiceAccount *helmServiceAccount `json:"serviceAccount,omitempty"`

	//TODO(npolshak) Remove this once default GatewayParameters are supported: https://github.com/solo-io/solo-projects/issues/6107
	IstioSDS *istioSDS `json:"istioSDS,omitempty"`
}

// helmPort represents a Gateway Listener port
type helmPort struct {
	Port       *uint16 `json:"port,omitempty"`
	Protocol   *string `json:"protocol,omitempty"`
	Name       *string `json:"name,omitempty"`
	TargetPort *uint16 `json:"targetPort,omitempty"`
}

type helmImage struct {
	Registry   *string `json:"registry,omitempty"`
	Repository *string `json:"repository,omitempty"`
	Tag        *string `json:"tag,omitempty"`
	Digest     *string `json:"digest,omitempty"`
	PullPolicy *string `json:"pullPolicy,omitempty"`
}

type helmService struct {
	Type             *string           `json:"type,omitempty"`
	ClusterIP        *string           `json:"clusterIP,omitempty"`
	ExtraAnnotations map[string]string `json:"extraAnnotations,omitempty"`
	ExtraLabels      map[string]string `json:"extraLabels,omitempty"`
}

// helmXds represents the xds host and port to which envoy will connect
// to receive xds config updates
type helmXds struct {
	Host *string `json:"host,omitempty"`
	Port *int32  `json:"port,omitempty"`
}

type helmAutoscaling struct {
	Enabled                           *bool   `json:"enabled,omitempty"`
	MinReplicas                       *uint32 `json:"minReplicas,omitempty"`
	MaxReplicas                       *uint32 `json:"maxReplicas,omitempty"`
	TargetCPUUtilizationPercentage    *uint32 `json:"targetCPUUtilizationPercentage,omitempty"`
	TargetMemoryUtilizationPercentage *uint32 `json:"targetMemoryUtilizationPercentage,omitempty"`
}

type helmSds struct {
	Image           *helmImage                         `json:"image,omitempty"`
	Resources       *v1alpha1kube.ResourceRequirements `json:"resources,omitempty"`
	SecurityContext *extcorev1.SecurityContext         `json:"securityContext,omitempty"`
	SdsBootstrap    *sdsBootstrap                      `json:"sdsBootstrap,omitempty"`

	Istio *helmIstioSds `json:"istioIntegration,omitempty"`
}

type sdsBootstrap struct {
	LogLevel *string `json:"logLevel,omitempty"`
}

// TODO: Remove this once default GatewayParameters are supported: https://github.com/solo-io/solo-projects/issues/6107
type istioSDS struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type helmIstioSds struct {
	Image    *helmImage `json:"image,omitempty"`
	LogLevel *string    `json:"logLevel,omitempty"`
	// Note: This is set by envoySidecarResources in helm chart
	Resources       *v1alpha1kube.ResourceRequirements `json:"resources,omitempty"`
	SecurityContext *extcorev1.SecurityContext         `json:"securityContext,omitempty"`

	IstioDiscoveryAddress *string `json:"istioDiscoveryAddress,omitempty"`
	IstioMetaMeshId       *string `json:"istioMetaMeshId,omitempty"`
	IstioMetaClusterId    *string `json:"istioMetaClusterId,omitempty"`
}

type helmServiceAccount struct {
	Create      *bool             `json:"create,omitempty"`
	Name        *string           `json:"name,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}
