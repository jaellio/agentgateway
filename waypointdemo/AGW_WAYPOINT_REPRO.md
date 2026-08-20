# AGW Waypoint Mesh Client -> Use-Waypoint Backend

Date: 2026-08-20

## Goal

Validate the waypoint flow where a meshed client calls a meshed backend Service that is configured with use-waypoint.

## Why This Manifest Was Run

- Namespace sets ambient mode for meshed workloads.
- AgentgatewayParameters enables Istio integration for the AGW waypoint and configures the waypoint Service as ClusterIP.
- Gateway creates the AGW-managed waypoint dataplane.
- Backend Service sets istio.io/use-waypoint so traffic to that Service goes through the waypoint.
- Client pod provides in-cluster test requests.

## Inline Manifest (Applied To Cluster)

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: agw-waypoint-mesh-e2e
  labels:
    istio.io/dataplane-mode: ambient
---
apiVersion: agentgateway.dev/v1alpha1
kind: AgentgatewayParameters
metadata:
  name: agw-waypoint-params
  namespace: agw-waypoint-mesh-e2e
spec:
  istio:
    enabled: true
  service:
    spec:
      type: ClusterIP
  deployment:
    spec:
      template:
        metadata:
          labels:
            istio.io/dataplane-mode: none
---
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: agw-waypoint
  namespace: agw-waypoint-mesh-e2e
spec:
  gatewayClassName: agentgateway-waypoint
  infrastructure:
    parametersRef:
      group: agentgateway.dev
      kind: AgentgatewayParameters
      name: agw-waypoint-params
  listeners:
  - name: hbone
    port: 15008
    protocol: HBONE
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
  namespace: agw-waypoint-mesh-e2e
spec:
  replicas: 1
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
      - name: backend
        image: hashicorp/http-echo:1.0
        args: ["-text=backend-waypoint-ok", "-listen=:8080"]
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: backend
  namespace: agw-waypoint-mesh-e2e
  labels:
    istio.io/use-waypoint: agw-waypoint
spec:
  selector:
    app: backend
  ports:
  - name: http
    port: 80
    targetPort: 8080
---
apiVersion: v1
kind: Pod
metadata:
  name: client
  namespace: agw-waypoint-mesh-e2e
spec:
  containers:
  - name: curl
    image: curlimages/curl:8.12.1
    command: ["sh", "-c", "sleep 3600"]
  restartPolicy: Never
```

## Visual HTTPRoute Policy Demo At Waypoint

This defines a service-scoped route policy intended to inject a response header for one path. In the current build, the request still falls through to _waypoint-default, so the header is not present yet.

### Policy Manifest

```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend-direct
  namespace: agw-waypoint-mesh-e2e
spec:
  selector:
    app: backend
  ports:
  - name: http
    port: 80
    targetPort: 8080
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: backend-waypoint-policy-header
  namespace: agw-waypoint-mesh-e2e
spec:
  hostnames:
  - backend.agw-waypoint-mesh-e2e.svc.cluster.local
  parentRefs:
  - group: ""
    kind: Service
    name: backend
    port: 80
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /header-demo
    filters:
    - type: ResponseHeaderModifier
      responseHeaderModifier:
        add:
        - name: x-waypoint-policy
          value: applied-at-waypoint
    backendRefs:
    - group: ""
      kind: Service
      name: backend-direct
      port: 80
```

### Proof Commands

```bash
# Control request (normal backend path)
kubectl exec -n agw-waypoint-mesh-e2e client -- \
  sh -c 'curl -sS -i http://backend.agw-waypoint-mesh-e2e.svc.cluster.local/ | sed -n "1,12p"'

# Policy request (current behavior: x-waypoint-policy header is NOT present)
kubectl exec -n agw-waypoint-mesh-e2e client -- \
  sh -c 'curl -sS -i http://backend.agw-waypoint-mesh-e2e.svc.cluster.local/header-demo | sed -n "1,24p"'

# Waypoint log proof
WP=$(kubectl get pod -n agw-waypoint-mesh-e2e -l gateway.networking.k8s.io/gateway-name=agw-waypoint -o jsonpath='{.items[0].metadata.name}')
kubectl logs -n agw-waypoint-mesh-e2e "$WP" --since=3m | \
  grep -E 'backend-waypoint-policy-header|header-demo|_waypoint-default'
```

### Expected Visual Result

- Request to / returns HTTP 200 with backend-waypoint-ok.
- Request to /header-demo returns HTTP 200.
- Current observed behavior in this build: request still matches _waypoint-default, so x-waypoint-policy is not yet present in the response.