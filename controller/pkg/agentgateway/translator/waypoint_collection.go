package translator

import "k8s.io/apimachinery/pkg/types"

// WaypointServiceBinding maps a fronted service to the AGW waypoint Gateway that fronts it.
type WaypointServiceBinding struct {
	// ServiceKey is the NamespacedName of the fronted service
	ServiceKey types.NamespacedName
	// WaypointGateway is the NamespacedName of the waypoint Gateway
	WaypointGateway types.NamespacedName
}

func (w WaypointServiceBinding) ResourceName() string {
	return w.ServiceKey.String()
}

func (w WaypointServiceBinding) Equals(other WaypointServiceBinding) bool {
	return w.ServiceKey == other.ServiceKey && w.WaypointGateway == other.WaypointGateway
}