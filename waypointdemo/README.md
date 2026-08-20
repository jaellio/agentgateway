# Waypoint Demo Index

This folder now has two separate demos and two separate READMEs.

## Guides

1. HBONE_GATEWAY guide: `waypointdemo/README_HBONE_GATEWAY.md`
2. HBONE_WAYPOINT guide: `waypointdemo/README_HBONE_WAYPOINT.md`

## Core differences

| Topic | HBONE_GATEWAY | HBONE_WAYPOINT |
|---|---|---|
| GatewayClass | `agentgateway` | `agentgateway-waypoint` |
| Bind tunnel protocol | `hboneGateway` | `hboneWaypoint` |
| Route attachment model | Gateway listener attachment on internal `inner-http` | Service parentRef attachment (`backend`) |
| Inner listener required | Yes (for this demo: `inner-http` on internal port `80`) | No |
| Typical use | Egress/ingress gateway-style re-entry into local bind/listener pipeline | Ambient waypoint preserving original destination semantics |

## Manifests

1. HBONE_GATEWAY manifest: `waypointdemo/waypoint-demo-manifest.yaml`
2. HBONE_WAYPOINT manifest: `waypointdemo/hbone-waypoint-manifest.yaml`

## Dumps

1. HBONE_GATEWAY full dump: `waypointdemo/HBONE_GATEWAY_CONFIG_DUMP_FULL.json`
2. HBONE_WAYPOINT full dump: `waypointdemo/HBONE_WAYPOINT_CONFIG_DUMP_FULL.json`
