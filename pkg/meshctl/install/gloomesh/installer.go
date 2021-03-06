package gloomesh

import (
	"context"
	"fmt"

	"github.com/solo-io/gloo-mesh/pkg/common/version"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/install/helm"
)

const (
	GlooMeshChartUriTemplate           = "https://storage.googleapis.com/gloo-mesh/gloo-mesh/gloo-mesh-%s.tgz"
	GlooMeshEnterpriseChartUriTemplate = "https://storage.googleapis.com/gloo-mesh-enterprise/gloo-mesh-enterprise/gloo-mesh-enterprise-%s.tgz"
	AgentCrdsChartUriTemplate          = "https://storage.googleapis.com/gloo-mesh/agent-crds/agent-crds-%s.tgz"
	CertAgentChartUriTemplate          = "https://storage.googleapis.com/gloo-mesh/cert-agent/cert-agent-%s.tgz"
	WasmAgentChartUriTemplate          = "https://storage.googleapis.com/gloo-mesh-enterprise/wasm-agent/wasm-agent-%s.tgz"
	GlooMeshReleaseName                = "gloo-mesh"
	GlooMeshEnterpriseReleaseName      = "gloo-mesh-enterprise"
	agentCrdsReleaseName               = "agent-crds"
	certAgentReleaseName               = "cert-agent"
	wasmAgentReleaseName               = "wasm-agent"
)

type Installer struct {
	HelmChartPath  string
	HelmValuesPath string
	KubeConfig     string
	KubeContext    string
	Namespace      string
	ReleaseName    string
	Values         map[string]string
	Verbose        bool
	DryRun         bool
}

func (i Installer) InstallGlooMesh(
	ctx context.Context,
) error {
	return i.install(ctx, GlooMeshChartUriTemplate, GlooMeshReleaseName)
}

func (i Installer) InstallGlooMeshEnterprise(
	ctx context.Context,
) error {
	return i.install(ctx, GlooMeshEnterpriseChartUriTemplate, GlooMeshEnterpriseReleaseName)
}

func (i Installer) InstallAgentCrds(
	ctx context.Context,
) error {
	return i.install(ctx, AgentCrdsChartUriTemplate, agentCrdsReleaseName)
}

func (i Installer) InstallCertAgent(
	ctx context.Context,
) error {
	return i.install(ctx, CertAgentChartUriTemplate, certAgentReleaseName)
}

func (i Installer) InstallWasmAgent(
	ctx context.Context,
) error {
	return i.install(ctx, WasmAgentChartUriTemplate, wasmAgentReleaseName)
}

func (i Installer) install(
	ctx context.Context,
	chartUriTemplate string,
	releaseName string,
) error {
	helmChartOverride := i.HelmChartPath

	if i.ReleaseName != "" {
		releaseName = i.ReleaseName
	}

	if helmChartOverride == "" {
		helmChartOverride = fmt.Sprintf(chartUriTemplate, version.Version)
	}

	return helm.Installer{
		KubeConfig:  i.KubeConfig,
		KubeContext: i.KubeContext,
		ChartUri:    helmChartOverride,
		Namespace:   i.Namespace,
		ReleaseName: releaseName,
		ValuesFile:  i.HelmValuesPath,
		Values:      i.Values,
		Verbose:     i.Verbose,
		DryRun:      i.DryRun,
	}.InstallChart(ctx)
}
