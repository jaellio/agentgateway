# HBONE_WAYPOINT Demo

This demo shows waypoint-style HBONE handling with `agentgateway-waypoint`, and a service-scoped policy route (`parentRefs.kind: Service`) so `/header-demo` is handled by the explicit policy route instead of `_waypoint-default`.

## Mode characteristics

1. GatewayClass: `agentgateway-waypoint`
2. Tunnel protocol in dump: `hboneWaypoint`
3. No inner HTTP listener is required for service-bound routing.
4. Route attaches to `kind: Service` (`backend`) via `parentRefs`.

## Deploy (one command)

```bash
kubectl apply -f waypointdemo/hbone-waypoint-manifest.yaml
kubectl -n agw-waypoint-mesh-e2e wait --for=condition=Programmed gateway/agw-waypoint --timeout=2m
kubectl -n agw-waypoint-mesh-e2e rollout status deploy/client --timeout=2m
```

## Validate

Route parent attachment:

```bash
kubectl -n agw-waypoint-mesh-e2e get httproute backend-waypoint-policy-header -o jsonpath='{.spec.parentRefs[*].kind}{"\n"}{.spec.parentRefs[*].name}{"\n"}{.status.parents[0].conditions[?(@.type=="Accepted")].status}{"\n"}'
```

Expected shape:

```text
Service
backend
True
```

Control path:

```bash
kubectl exec -n agw-waypoint-mesh-e2e deploy/client -- \
  sh -c 'curl -sS -i http://backend.agw-waypoint-mesh-e2e.svc.cluster.local/ | sed -n "1,12p"'
```

Policy path:

```bash
kubectl exec -n agw-waypoint-mesh-e2e deploy/client -- \
  sh -c 'curl -sS -i http://backend.agw-waypoint-mesh-e2e.svc.cluster.local/header-demo | sed -n "1,24p"'
```

Notes:

1. The route is service-scoped (`parentRefs.kind: Service`, name `backend`).
2. `/header-demo` should include `x-waypoint-policy: applied-at-waypoint`.

Optional log check:

```bash
WP=$(kubectl get pod -n agw-waypoint-mesh-e2e -l gateway.networking.k8s.io/gateway-name=agw-waypoint -o jsonpath='{.items[0].metadata.name}')
kubectl logs -n agw-waypoint-mesh-e2e "$WP" --since=3m | grep -E 'backend-waypoint-policy-header|header-demo|_waypoint-default'
```

For `/header-demo`, expect route `backend-waypoint-policy-header` in logs rather than `_waypoint-default`.

Troubleshooting:

1. If `/header-demo` falls through `_waypoint-default`, ensure your controller includes the service-key FQDN translation fix in `controller/pkg/agentgateway/translator/route_collections.go` and restart the waypoint pod after rolling out the controller.

## Capture config dump

```bash
POD=$(kubectl -n agw-waypoint-mesh-e2e get pod -l gateway.networking.k8s.io/gateway-name=agw-waypoint -o jsonpath='{.items[0].metadata.name}')
kubectl -n agw-waypoint-mesh-e2e port-forward pod/$POD 25600:15000
```

In another terminal:

```bash
curl -sS http://127.0.0.1:25600/config_dump | jq -S . > waypointdemo/HBONE_WAYPOINT_CONFIG_DUMP_FULL.json
```

## Cleanup

```bash
kubectl -n agw-waypoint-mesh-e2e delete httproute backend-waypoint-policy-header --ignore-not-found
kubectl -n agw-waypoint-mesh-e2e delete service backend-direct --ignore-not-found
kubectl -n agw-waypoint-mesh-e2e delete service backend --ignore-not-found
kubectl -n agw-waypoint-mesh-e2e delete deploy backend --ignore-not-found
kubectl -n agw-waypoint-mesh-e2e delete deploy client --ignore-not-found
kubectl -n agw-waypoint-mesh-e2e delete gateway agw-waypoint --ignore-not-found
kubectl -n agw-waypoint-mesh-e2e delete agentgatewayparameters agw-waypoint-params --ignore-not-found
kubectl delete namespace agw-waypoint-mesh-e2e --ignore-not-found
```
