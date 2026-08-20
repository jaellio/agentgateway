package syncer

import (
	"strings"

	"github.com/agentgateway/agentgateway/controller/pkg/agentgateway/translator"
	"github.com/agentgateway/agentgateway/controller/pkg/agentgateway/utils"
	"github.com/agentgateway/agentgateway/controller/pkg/pluginsdk/krtutil"
	"github.com/agentgateway/agentgateway/controller/pkg/wellknown"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/config/schema/kind"
	"istio.io/istio/pkg/kube/krt"
	"k8s.io/apimachinery/pkg/types"
)

// BuildWaypointServiceBindings creates a collection mapping services to their AGW waypoint gateways.
// For each ambient-derived k8s Service bound to a waypoint Gateway, a
// WaypointServiceBinding is created.
func BuildWaypointServiceBindings(
	services krt.Collection[model.ServiceInfo],
	gateways krt.Collection[*translator.GatewayListener],
	krtopts krtutil.KrtOptions,
) krt.Collection[translator.WaypointServiceBinding] {
	// Index gateway listeners by their parent Gateway. GatewayListener is a
	// per-listener resource whose ResourceName is ns/name/listener, so we cannot
	// look it up by the gateway-level ns/name key directly.
	gwIndex := krt.NewIndex(gateways, "parentGateway", func(o *translator.GatewayListener) []utils.TypedNamespacedName {
		return []utils.TypedNamespacedName{o.ParentObject}
	})
	return krt.NewCollection(services, func(ctx krt.HandlerContext, svc model.ServiceInfo) *translator.WaypointServiceBinding {
		if svc.Source.Kind != kind.Service {
			return nil
		}

		waypointResource := svc.Waypoint.ResourceName
		if waypointResource == "" {
			return nil
		}
		wpNamespace, wpName, found := strings.Cut(waypointResource, "/")
		if !found || wpNamespace == "" || wpName == "" {
			return nil
		}

		// Find a listener belonging to the referenced waypoint Gateway.
		wpKey := utils.TypedNamespacedName{
			Kind: wellknown.GatewayGVK.Kind,
			NamespacedName: types.NamespacedName{
				Name:      wpName,
				Namespace: wpNamespace,
			},
		}
		listeners := krt.Fetch(ctx, gateways, krt.FilterIndex(gwIndex, wpKey))
		if len(listeners) == 0 {
			return nil
		}

		// Check that the gateway is a waypoint
		if !listeners[0].ParentInfo.Waypoint {
			return nil
		}

		return &translator.WaypointServiceBinding{
			ServiceKey:      svc.Source.NamespacedName,
			WaypointGateway: types.NamespacedName{Namespace: wpNamespace, Name: wpName},
		}
	}, krtopts.ToOptions("WaypointServiceBindings")...)
}
