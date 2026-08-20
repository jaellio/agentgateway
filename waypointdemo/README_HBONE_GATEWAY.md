# HBONE_GATEWAY Demo

This demo shows gateway-style HBONE handling, where inbound HBONE CONNECT traffic is terminated and re-enters the local bind/listener pipeline.

## Mode characteristics

1. GatewayClass: `agentgateway`
2. Tunnel protocol in dump: `hboneGateway`
3. Uses an internal HTTP listener (`inner-http` on port `80`) for route attachment.
4. Route is attached via `HTTPRoute.spec.parentRefs -> Gateway sectionName: inner-http`.

## Deploy (one command)

```bash
kubectl apply -f waypointdemo/waypoint-demo-manifest.yaml
kubectl -n agents rollout status deploy/curl
kubectl -n agentgateway-egress wait --for=condition=Programmed gateway/agw-gateway --timeout=2m
kubectl -n agentgateway-egress wait --for=jsonpath='{.status.parents[0].conditions[?(@.type=="Accepted")].status}'=True httproute/httpbin-via-agw --timeout=2m
```

## Validate

Listener attachment:

```bash
kubectl -n agentgateway-egress get gateway agw-gateway -o jsonpath='{.status.listeners[*].name}{"\n"}{.status.listeners[*].attachedRoutes}{"\n"}'
```

Expected shape:

```text
mesh inner-http
0 1
```

Traffic test:

```bash
kubectl -n agents exec deploy/curl -- curl -sSI http://httpbin.org/get
```

Expected header:

```text
x-agw-waypoint: true
```

## Capture config dump

```bash
POD=$(kubectl -n agentgateway-egress get pod -l gateway.networking.k8s.io/gateway-name=agw-gateway -o jsonpath='{.items[0].metadata.name}')
kubectl -n agentgateway-egress port-forward pod/$POD 25500:15000
```

In another terminal:

```bash
curl -sS http://127.0.0.1:25500/config_dump | jq -S . > waypointdemo/HBONE_GATEWAY_CONFIG_DUMP_FULL.json
```

## Cleanup

```bash
kubectl -n agentgateway-egress delete httproute httpbin-via-agw --ignore-not-found
kubectl -n agentgateway-egress delete agentgatewaybackend httpbin-backend --ignore-not-found
kubectl -n agents delete serviceentry httpbin-external --ignore-not-found
kubectl -n agentgateway-egress delete gateway agw-gateway --ignore-not-found
kubectl -n agentgateway-egress delete agentgatewayparameters agw-gateway-params --ignore-not-found
kubectl -n agents delete deploy curl --ignore-not-found
kubectl delete namespace agents agentgateway-egress --ignore-not-found
```
