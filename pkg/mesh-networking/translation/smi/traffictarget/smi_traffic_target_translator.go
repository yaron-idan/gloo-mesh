package traffictarget

import (
	"context"

	discoveryv1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/smi"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/smi/traffictarget/access"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/smi/traffictarget/split"
)

//go:generate mockgen -source ./smi_traffic_target_translator.go -destination mocks/smi_traffic_target_translator.go

// the VirtualService translator translates a TrafficTarget into a VirtualService.
type Translator interface {
	// Translate translates the appropriate VirtualService and DestinationRule for the given TrafficTarget.
	// returns nil if no VirtualService or DestinationRule is required for the TrafficTarget (i.e. if no VirtualService/DestinationRule features are required, such as subsets).
	// Output resources will be added to the smi
	// Errors caused by invalid user config will be reported using the Reporter.
	Translate(
		ctx context.Context,
		in input.LocalSnapshot,
		trafficTarget *discoveryv1alpha2.TrafficTarget,
		outputs smi.Builder,
		reporter reporting.Reporter,
	)
}

type translator struct {
	trafficSplit  split.Translator
	trafficTarget access.Translator
}

func NewTranslator(tsTranslator split.Translator, ttTranslator access.Translator) Translator {
	return &translator{
		trafficSplit:  tsTranslator,
		trafficTarget: ttTranslator,
	}
}

// translate the appropriate resources for the given TrafficTarget.
func (t *translator) Translate(
	ctx context.Context,
	in input.LocalSnapshot,
	trafficTarget *discoveryv1alpha2.TrafficTarget,
	outputs smi.Builder,
	reporter reporting.Reporter,
) {
	// Translate TrafficSplit for TrafficTarget, can be nil if non-kube service or no applied traffic policy
	trafficSplit := t.trafficSplit.Translate(ctx, in, trafficTarget, reporter)

	// Translate output TrafficTargets and HttpRouteGroups for discovered TrafficTarget
	trafficTargets, httpRouteGroups := t.trafficTarget.Translate(ctx, in, trafficTarget, reporter)

	outputs.AddTrafficSplits(trafficSplit)
	outputs.AddTrafficTargets(trafficTargets...)
	outputs.AddHTTPRouteGroups(httpRouteGroups...)
}
