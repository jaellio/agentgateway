package syncer

import (
	"github.com/agentgateway/agentgateway/controller/pkg/agentgateway/translator"
	"github.com/agentgateway/agentgateway/controller/pkg/agentgateway/utils"
	"github.com/agentgateway/agentgateway/controller/pkg/pluginsdk/krtutil"
	"github.com/agentgateway/agentgateway/controller/pkg/wellknown"
	"istio.io/istio/pilot/pkg/serviceregistry/ambient"
	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/ptr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// BuildWaypointServiceBindings creates a collection mapping services to their AGW waypoint gateways.
// For each k8s Service with a use-waypoint label (or inheriting from namespace) pointing to an
// agentgateway-waypoint class Gateway, a WaypointServiceBinding is created.
func BuildWaypointServiceBindings(
	services krt.Collection[*corev1.Service],
	namespaces krt.Collection[*corev1.Namespace],
	gateways krt.Collection[*translator.GatewayListener],
	krtopts krtutil.KrtOptions,
) krt.Collection[translator.WaypointServiceBinding] {
	// Index gateway listeners by their parent Gateway. GatewayListener is a
	// per-listener resource whose ResourceName is ns/name/listener, so we cannot
	// look it up by the gateway-level ns/name key directly.
	gwIndex := krt.NewIndex(gateways, "parentGateway", func(o *translator.GatewayListener) []utils.TypedNamespacedName {
		return []utils.TypedNamespacedName{o.ParentObject}
	})
	return krt.NewCollection(services, func(ctx krt.HandlerContext, svc *corev1.Service) *translator.WaypointServiceBinding {
		// check if the service or its namespace has the use-waypoint label
		wpRef := resolveUseWaypoint(ctx, svc.ObjectMeta, namespaces)
		if wpRef == nil {
			return nil
		}

		// Find a listener belonging to the referenced waypoint Gateway.
		wpKey := utils.TypedNamespacedName{
			Kind: wellknown.GatewayGVK.Kind,
			NamespacedName: types.NamespacedName{
				Name:      wpRef.Name,
				Namespace: wpRef.Namespace,
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
			ServiceKey:      types.NamespacedName{Namespace: svc.Namespace, Name: svc.Name},
			WaypointGateway: types.NamespacedName{Namespace: wpRef.Namespace, Name: wpRef.Name},
		}
	}, krtopts.ToOptions("WaypointServiceBindings")...)
}

// resolveUseWaypoint looks up the use-waypoint label on a service or its namespace
// and returns the referenced waypoint gateway, if any.
func resolveUseWaypoint(
	ctx krt.HandlerContext,
	meta metav1.ObjectMeta,
	namespaces krt.Collection[*corev1.Namespace],
) *krt.Named {
	// Check object labels first
	// These labels take precedence over namespace labels
	wp, isNone := ambient.GetUseWaypoint(meta, meta.Namespace)
	if isNone {
		return nil
	}
	if wp != nil {
		return wp
	}

	// Fall back to namespace labels
	ns := ptr.Flatten(krt.FetchOne(ctx, namespaces, krt.FilterKey(meta.Namespace)))
	if ns == nil {
		return nil
	}
	wp, _ = ambient.GetUseWaypoint(ns.ObjectMeta, meta.Namespace)
	return wp
}