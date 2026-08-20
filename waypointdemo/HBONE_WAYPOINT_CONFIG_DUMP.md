# HBONE_WAYPOINT Full Config Dump

Source: live Gateway admin /config_dump snapshot

```json
{
  "workloads": [
    {
      "workloadIps": [
        "10.244.0.45"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/agentgateway-egress/agw-gateway-68867df46c-bck7z",
      "name": "agw-gateway-68867df46c-bck7z",
      "namespace": "agentgateway-egress",
      "trustDomain": "cluster.local",
      "serviceAccount": "agw-gateway",
      "workloadName": "agw-gateway",
      "workloadType": "pod",
      "canonicalName": "agw-gateway",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "agentgateway-egress/agw-gateway.agentgateway-egress.svc.cluster.local": {
          "15008": 15008
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.47"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/agentgateway-system/agentgateway-764f786975-84drj",
      "name": "agentgateway-764f786975-84drj",
      "namespace": "agentgateway-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "agentgateway",
      "workloadName": "agentgateway",
      "workloadType": "pod",
      "canonicalName": "agentgateway",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "agentgateway-system/agentgateway.agentgateway-system.svc.cluster.local": {
          "9978": 9978,
          "9093": 9093,
          "9092": 9092
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.43"
      ],
      "protocol": "HBONE",
      "networkMode": "Standard",
      "uid": "//Pod/agents/curl-588f749d75-s8lg5",
      "name": "curl-588f749d75-s8lg5",
      "namespace": "agents",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "curl",
      "workloadType": "pod",
      "canonicalName": "curl",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.24"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/agw-waypoint-e2e/agw-waypoint-74d49d4bd6-vz4ht",
      "name": "agw-waypoint-74d49d4bd6-vz4ht",
      "namespace": "agw-waypoint-e2e",
      "trustDomain": "cluster.local",
      "serviceAccount": "agw-waypoint",
      "workloadName": "agw-waypoint",
      "workloadType": "pod",
      "canonicalName": "agw-waypoint",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "agw-waypoint-e2e/agw-waypoint.agw-waypoint-e2e.svc.cluster.local": {
          "15008": 15008
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.37"
      ],
      "protocol": "HBONE",
      "networkMode": "Standard",
      "uid": "//Pod/agw-waypoint-e2e/backend-69cf876bf4-6bgtc",
      "name": "backend-69cf876bf4-6bgtc",
      "namespace": "agw-waypoint-e2e",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "backend",
      "workloadType": "pod",
      "canonicalName": "backend",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "agw-waypoint-e2e/backend.agw-waypoint-e2e.svc.cluster.local": {
          "80": 8080
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.38"
      ],
      "protocol": "HBONE",
      "networkMode": "Standard",
      "uid": "//Pod/agw-waypoint-e2e/backend-b-7c54f989fc-tqkp9",
      "name": "backend-b-7c54f989fc-tqkp9",
      "namespace": "agw-waypoint-e2e",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "backend-b",
      "workloadType": "pod",
      "canonicalName": "backend-b",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "agw-waypoint-e2e/backend-b.agw-waypoint-e2e.svc.cluster.local": {
          "80": 8080
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.48"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/agw-waypoint-mesh-e2e/agw-waypoint-64db89d659-qnqk7",
      "name": "agw-waypoint-64db89d659-qnqk7",
      "namespace": "agw-waypoint-mesh-e2e",
      "trustDomain": "cluster.local",
      "serviceAccount": "agw-waypoint",
      "workloadName": "agw-waypoint",
      "workloadType": "pod",
      "canonicalName": "agw-waypoint",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "agw-waypoint-mesh-e2e/agw-waypoint.agw-waypoint-mesh-e2e.svc.cluster.local": {
          "15008": 15008
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.39"
      ],
      "protocol": "HBONE",
      "networkMode": "Standard",
      "uid": "//Pod/agw-waypoint-mesh-e2e/backend-c548b79c4-hxrdp",
      "name": "backend-c548b79c4-hxrdp",
      "namespace": "agw-waypoint-mesh-e2e",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "backend",
      "workloadType": "pod",
      "canonicalName": "backend",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "agw-waypoint-mesh-e2e/backend-direct.agw-waypoint-mesh-e2e.svc.cluster.local": {
          "80": 8080
        },
        "agw-waypoint-mesh-e2e/backend.agw-waypoint-mesh-e2e.svc.cluster.local": {
          "80": 8080
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.46"
      ],
      "protocol": "HBONE",
      "networkMode": "Standard",
      "uid": "//Pod/agw-waypoint-mesh-e2e/client-7d7b9db46-ljnn5",
      "name": "client-7d7b9db46-ljnn5",
      "namespace": "agw-waypoint-mesh-e2e",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "client",
      "workloadType": "pod",
      "canonicalName": "client",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.10"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/istio-system/istio-cni-node-xmkpn",
      "name": "istio-cni-node-xmkpn",
      "namespace": "istio-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "istio-cni",
      "workloadName": "istio-cni-node",
      "workloadType": "pod",
      "canonicalName": "istio-cni",
      "canonicalRevision": "1.0.0",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.36"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/istio-system/istiod-7fffbcb657-jxgvr",
      "name": "istiod-7fffbcb657-jxgvr",
      "namespace": "istio-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "istiod",
      "workloadName": "istiod",
      "workloadType": "pod",
      "canonicalName": "istiod",
      "canonicalRevision": "1.0.0",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "istio-system/istiod.istio-system.svc.cluster.local": {
          "443": 15017,
          "15010": 15010,
          "15012": 15012,
          "15014": 15014
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.25"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/istio-system/ztunnel-wk8rh",
      "name": "ztunnel-wk8rh",
      "namespace": "istio-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "ztunnel",
      "workloadName": "ztunnel",
      "workloadType": "pod",
      "canonicalName": "ztunnel",
      "canonicalRevision": "1.0.0",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.22"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/kube-system/coredns-589f44dc88-4x7pw",
      "name": "coredns-589f44dc88-4x7pw",
      "namespace": "kube-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "coredns",
      "workloadName": "coredns",
      "workloadType": "pod",
      "canonicalName": "coredns",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "kube-system/kube-dns.kube-system.svc.cluster.local": {
          "53": 53,
          "9153": 9153
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.23"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/kube-system/coredns-589f44dc88-fk6hx",
      "name": "coredns-589f44dc88-fk6hx",
      "namespace": "kube-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "coredns",
      "workloadName": "coredns",
      "workloadType": "pod",
      "canonicalName": "coredns",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "kube-system/kube-dns.kube-system.svc.cluster.local": {
          "9153": 9153,
          "53": 53
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "172.18.0.2"
      ],
      "protocol": "TCP",
      "networkMode": "HostNetwork",
      "uid": "//Pod/kube-system/etcd-kind-control-plane",
      "name": "etcd-kind-control-plane",
      "namespace": "kube-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "etcd-kind-control-plane",
      "workloadType": "pod",
      "canonicalName": "etcd-kind-control-plane",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "172.18.0.2"
      ],
      "protocol": "TCP",
      "networkMode": "HostNetwork",
      "uid": "//Pod/kube-system/kindnet-pzswh",
      "name": "kindnet-pzswh",
      "namespace": "kube-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "kindnet",
      "workloadName": "kindnet",
      "workloadType": "pod",
      "canonicalName": "kindnet",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "172.18.0.2"
      ],
      "protocol": "TCP",
      "networkMode": "HostNetwork",
      "uid": "//Pod/kube-system/kube-apiserver-kind-control-plane",
      "name": "kube-apiserver-kind-control-plane",
      "namespace": "kube-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "kube-apiserver-kind-control-plane",
      "workloadType": "pod",
      "canonicalName": "kube-apiserver-kind-control-plane",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "172.18.0.2"
      ],
      "protocol": "TCP",
      "networkMode": "HostNetwork",
      "uid": "//Pod/kube-system/kube-controller-manager-kind-control-plane",
      "name": "kube-controller-manager-kind-control-plane",
      "namespace": "kube-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "kube-controller-manager-kind-control-plane",
      "workloadType": "pod",
      "canonicalName": "kube-controller-manager-kind-control-plane",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "172.18.0.2"
      ],
      "protocol": "TCP",
      "networkMode": "HostNetwork",
      "uid": "//Pod/kube-system/kube-proxy-wj9zm",
      "name": "kube-proxy-wj9zm",
      "namespace": "kube-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "kube-proxy",
      "workloadName": "kube-proxy",
      "workloadType": "pod",
      "canonicalName": "kube-proxy",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "172.18.0.2"
      ],
      "protocol": "TCP",
      "networkMode": "HostNetwork",
      "uid": "//Pod/kube-system/kube-scheduler-kind-control-plane",
      "name": "kube-scheduler-kind-control-plane",
      "namespace": "kube-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "kube-scheduler-kind-control-plane",
      "workloadType": "pod",
      "canonicalName": "kube-scheduler-kind-control-plane",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.33"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/local-path-storage/local-path-provisioner-855c7b7774-ps47h",
      "name": "local-path-provisioner-855c7b7774-ps47h",
      "namespace": "local-path-storage",
      "trustDomain": "cluster.local",
      "serviceAccount": "local-path-provisioner-service-account",
      "workloadName": "local-path-provisioner",
      "workloadType": "pod",
      "canonicalName": "local-path-provisioner",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "10.244.0.28"
      ],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "//Pod/metallb-system/controller-77f7b7f69d-hbhkt",
      "name": "controller-77f7b7f69d-hbhkt",
      "namespace": "metallb-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "controller",
      "workloadName": "controller",
      "workloadType": "pod",
      "canonicalName": "metallb",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "services": {
        "metallb-system/webhook-service.metallb-system.svc.cluster.local": {
          "443": 9443
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "172.18.0.2"
      ],
      "protocol": "TCP",
      "networkMode": "HostNetwork",
      "uid": "//Pod/metallb-system/speaker-nn578",
      "name": "speaker-nn578",
      "namespace": "metallb-system",
      "trustDomain": "cluster.local",
      "serviceAccount": "speaker",
      "workloadName": "speaker",
      "workloadType": "pod",
      "canonicalName": "metallb",
      "canonicalRevision": "latest",
      "node": "kind-control-plane",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "locality": {
        "region": "region",
        "zone": "zone",
        "subzone": ""
      },
      "capacity": 1
    },
    {
      "workloadIps": [
        "172.18.0.2"
      ],
      "protocol": "TCP",
      "networkMode": "HostNetwork",
      "uid": "/discovery.k8s.io/EndpointSlice/default/kubernetes/172.18.0.2",
      "name": "kubernetes",
      "namespace": "default",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadType": "deployment",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "services": {
        "default/kubernetes.default.svc.cluster.local": {
          "443": 6443
        }
      },
      "capacity": 1
    },
    {
      "workloadIps": [],
      "protocol": "TCP",
      "networkMode": "Standard",
      "uid": "/networking.istio.io/ServiceEntry/agents/httpbin-external/httpbin.org",
      "name": "httpbin-external",
      "namespace": "agents",
      "trustDomain": "cluster.local",
      "serviceAccount": "default",
      "workloadName": "httpbin-external",
      "workloadType": "pod",
      "canonicalName": "httpbin-external",
      "canonicalRevision": "latest",
      "hostname": "httpbin.org",
      "status": "Healthy",
      "clusterId": "Kubernetes",
      "services": {
        "agents/httpbin.org": {
          "80": 80
        }
      },
      "capacity": 1
    }
  ],
  "services": [
    {
      "name": "agentgateway",
      "namespace": "agentgateway-system",
      "hostname": "agentgateway.agentgateway-system.svc.cluster.local",
      "vips": [
        "/10.96.153.206"
      ],
      "ports": {
        "9093": 0,
        "9978": 0,
        "9092": 0
      },
      "appProtocols": {
        "9978": "Grpc"
      },
      "endpoints": [
        {
          "active": {
            "//Pod/agentgateway-system/agentgateway-764f786975-84drj": {
              "endpoint": {
                "workloadUid": "//Pod/agentgateway-system/agentgateway-764f786975-84drj",
                "port": {
                  "9978": 9978,
                  "9093": 9093,
                  "9092": 9092
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "agw-gateway",
      "namespace": "agentgateway-egress",
      "hostname": "agw-gateway.agentgateway-egress.svc.cluster.local",
      "vips": [
        "/10.96.165.84"
      ],
      "ports": {
        "15008": 15008
      },
      "appProtocols": {},
      "endpoints": [
        {
          "active": {
            "//Pod/agentgateway-egress/agw-gateway-68867df46c-bck7z": {
              "endpoint": {
                "workloadUid": "//Pod/agentgateway-egress/agw-gateway-68867df46c-bck7z",
                "port": {
                  "15008": 15008
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "agw-waypoint",
      "namespace": "agw-waypoint-e2e",
      "hostname": "agw-waypoint.agw-waypoint-e2e.svc.cluster.local",
      "vips": [
        "/10.96.41.215"
      ],
      "ports": {
        "15008": 15008
      },
      "appProtocols": {},
      "endpoints": [
        {
          "active": {
            "//Pod/agw-waypoint-e2e/agw-waypoint-74d49d4bd6-vz4ht": {
              "endpoint": {
                "workloadUid": "//Pod/agw-waypoint-e2e/agw-waypoint-74d49d4bd6-vz4ht",
                "port": {
                  "15008": 15008
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "agw-waypoint",
      "namespace": "agw-waypoint-mesh-e2e",
      "hostname": "agw-waypoint.agw-waypoint-mesh-e2e.svc.cluster.local",
      "vips": [
        "/10.96.33.186"
      ],
      "ports": {
        "15008": 15008
      },
      "appProtocols": {},
      "endpoints": [
        {
          "active": {
            "//Pod/agw-waypoint-mesh-e2e/agw-waypoint-64db89d659-qnqk7": {
              "endpoint": {
                "workloadUid": "//Pod/agw-waypoint-mesh-e2e/agw-waypoint-64db89d659-qnqk7",
                "port": {
                  "15008": 15008
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "backend-b",
      "namespace": "agw-waypoint-e2e",
      "hostname": "backend-b.agw-waypoint-e2e.svc.cluster.local",
      "vips": [
        "/10.96.158.14"
      ],
      "ports": {
        "80": 8080
      },
      "appProtocols": {
        "80": "Http11"
      },
      "endpoints": [
        {
          "active": {
            "//Pod/agw-waypoint-e2e/backend-b-7c54f989fc-tqkp9": {
              "endpoint": {
                "workloadUid": "//Pod/agw-waypoint-e2e/backend-b-7c54f989fc-tqkp9",
                "port": {
                  "80": 8080
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "backend-direct",
      "namespace": "agw-waypoint-mesh-e2e",
      "hostname": "backend-direct.agw-waypoint-mesh-e2e.svc.cluster.local",
      "vips": [
        "/10.96.147.45"
      ],
      "ports": {
        "80": 8080
      },
      "appProtocols": {
        "80": "Http11"
      },
      "endpoints": [
        {
          "active": {
            "//Pod/agw-waypoint-mesh-e2e/backend-c548b79c4-hxrdp": {
              "endpoint": {
                "workloadUid": "//Pod/agw-waypoint-mesh-e2e/backend-c548b79c4-hxrdp",
                "port": {
                  "80": 8080
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.004199425362,
                "pendingRequests": 0,
                "totalRequests": 4,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "backend",
      "namespace": "agw-waypoint-e2e",
      "hostname": "backend.agw-waypoint-e2e.svc.cluster.local",
      "vips": [
        "/10.96.161.234"
      ],
      "ports": {
        "80": 8080
      },
      "appProtocols": {
        "80": "Http11"
      },
      "endpoints": [
        {
          "active": {
            "//Pod/agw-waypoint-e2e/backend-69cf876bf4-6bgtc": {
              "endpoint": {
                "workloadUid": "//Pod/agw-waypoint-e2e/backend-69cf876bf4-6bgtc",
                "port": {
                  "80": 8080
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "backend",
      "namespace": "agw-waypoint-mesh-e2e",
      "hostname": "backend.agw-waypoint-mesh-e2e.svc.cluster.local",
      "vips": [
        "/10.96.3.57"
      ],
      "ports": {
        "80": 8080
      },
      "appProtocols": {
        "80": "Http11"
      },
      "endpoints": [
        {
          "active": {
            "//Pod/agw-waypoint-mesh-e2e/backend-c548b79c4-hxrdp": {
              "endpoint": {
                "workloadUid": "//Pod/agw-waypoint-mesh-e2e/backend-c548b79c4-hxrdp",
                "port": {
                  "80": 8080
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "waypoint": {
        "destination": "/10.96.33.186",
        "hboneMtlsPort": 15008
      },
      "ipFamilies": "IPv4"
    },
    {
      "name": "frontend",
      "namespace": "agw-waypoint-e2e",
      "hostname": "frontend.agw-waypoint-e2e.svc.cluster.local",
      "vips": [
        "/10.96.43.84"
      ],
      "ports": {
        "80": 8080
      },
      "appProtocols": {
        "80": "Http11"
      },
      "endpoints": [
        {
          "active": {},
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "waypoint": {
        "destination": "/10.96.41.215",
        "hboneMtlsPort": 15008
      },
      "ipFamilies": "IPv4",
      "ingressUseWaypoint": true
    },
    {
      "name": "httpbin-external",
      "namespace": "agents",
      "hostname": "httpbin.org",
      "vips": [
        "/240.240.0.1",
        "/2001:2::1"
      ],
      "ports": {
        "80": 80
      },
      "appProtocols": {
        "80": "Http11"
      },
      "endpoints": [
        {
          "active": {
            "/networking.istio.io/ServiceEntry/agents/httpbin-external/httpbin.org": {
              "endpoint": {
                "workloadUid": "/networking.istio.io/ServiceEntry/agents/httpbin-external/httpbin.org",
                "port": {
                  "80": 80
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "waypoint": {
        "destination": "/10.96.165.84",
        "hboneMtlsPort": 15008
      }
    },
    {
      "name": "istiod",
      "namespace": "istio-system",
      "hostname": "istiod.istio-system.svc.cluster.local",
      "vips": [
        "/10.96.157.91"
      ],
      "ports": {
        "15014": 15014,
        "15012": 15012,
        "15010": 15010,
        "443": 15017
      },
      "appProtocols": {
        "15012": "Tls",
        "15014": "Http11",
        "15010": "Grpc",
        "443": "Tls"
      },
      "endpoints": [
        {
          "active": {
            "//Pod/istio-system/istiod-7fffbcb657-jxgvr": {
              "endpoint": {
                "workloadUid": "//Pod/istio-system/istiod-7fffbcb657-jxgvr",
                "port": {
                  "443": 15017,
                  "15010": 15010,
                  "15012": 15012,
                  "15014": 15014
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "kube-dns",
      "namespace": "kube-system",
      "hostname": "kube-dns.kube-system.svc.cluster.local",
      "vips": [
        "/10.96.0.10"
      ],
      "ports": {
        "53": 53,
        "9153": 9153
      },
      "appProtocols": {
        "53": "Tcp"
      },
      "endpoints": [
        {
          "active": {
            "//Pod/kube-system/coredns-589f44dc88-fk6hx": {
              "endpoint": {
                "workloadUid": "//Pod/kube-system/coredns-589f44dc88-fk6hx",
                "port": {
                  "9153": 9153,
                  "53": 53
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            },
            "//Pod/kube-system/coredns-589f44dc88-4x7pw": {
              "endpoint": {
                "workloadUid": "//Pod/kube-system/coredns-589f44dc88-4x7pw",
                "port": {
                  "53": 53,
                  "9153": 9153
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "kubernetes",
      "namespace": "default",
      "hostname": "kubernetes.default.svc.cluster.local",
      "vips": [
        "/10.96.0.1"
      ],
      "ports": {
        "443": 6443
      },
      "appProtocols": {
        "443": "Tls"
      },
      "endpoints": [
        {
          "active": {
            "/discovery.k8s.io/EndpointSlice/default/kubernetes/172.18.0.2": {
              "endpoint": {
                "workloadUid": "/discovery.k8s.io/EndpointSlice/default/kubernetes/172.18.0.2",
                "port": {
                  "443": 6443
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    },
    {
      "name": "webhook-service",
      "namespace": "metallb-system",
      "hostname": "webhook-service.metallb-system.svc.cluster.local",
      "vips": [
        "/10.96.224.34"
      ],
      "ports": {
        "443": 9443
      },
      "appProtocols": {},
      "endpoints": [
        {
          "active": {
            "//Pod/metallb-system/controller-77f7b7f69d-hbhkt": {
              "endpoint": {
                "workloadUid": "//Pod/metallb-system/controller-77f7b7f69d-hbhkt",
                "port": {
                  "443": 9443
                },
                "status": "Healthy"
              },
              "info": {
                "health": 1.0,
                "requestLatency": 0.0,
                "pendingRequests": 0,
                "totalRequests": 0,
                "consecutiveFailures": 0,
                "timesEjected": 0,
                "evictedUntil": null
              },
              "capacity": 1
            }
          },
          "rejected": {}
        }
      ],
      "subjectAltNames": [],
      "ipFamilies": "IPv4"
    }
  ],
  "binds": [
    {
      "key": "15008/agw-waypoint-mesh-e2e/agw-waypoint",
      "address": "[::]:15008",
      "protocol": "http",
      "tunnelProtocol": "hboneWaypoint",
      "mode": "standard",
      "listeners": {
        "agw-waypoint-mesh-e2e/agw-waypoint.hbone": {
          "key": "agw-waypoint-mesh-e2e/agw-waypoint.hbone",
          "gatewayName": "agw-waypoint",
          "gatewayNamespace": "agw-waypoint-mesh-e2e",
          "listenerName": "hbone",
          "hostname": "",
          "protocol": "HBONE"
        }
      }
    }
  ],
  "routes": {
    "httpMesh": {
      "agw-waypoint-mesh-e2e/backend.agw-waypoint-mesh-e2e.svc.cluster.local": {
        "agw-waypoint-mesh-e2e/backend-waypoint-policy-header.00.svc.agw-waypoint-mesh-e2e.backend": {
          "key": "agw-waypoint-mesh-e2e/backend-waypoint-policy-header.00.svc.agw-waypoint-mesh-e2e.backend",
          "serviceKey": "agw-waypoint-mesh-e2e/backend.agw-waypoint-mesh-e2e.svc.cluster.local",
          "name": "backend-waypoint-policy-header",
          "namespace": "agw-waypoint-mesh-e2e",
          "kind": "HTTPRoute",
          "matches": [
            {
              "path": {
                "pathPrefix": "/header-demo"
              }
            }
          ],
          "backends": [
            {
              "weight": 1,
              "service": {
                "name": "agw-waypoint-mesh-e2e/backend-direct.agw-waypoint-mesh-e2e.svc.cluster.local",
                "port": 80
              }
            }
          ],
          "inlinePolicies": [
            {
              "responseHeaderModifier": {
                "add": {
                  "x-waypoint-policy": "applied-at-waypoint"
                }
              }
            }
          ]
        }
      }
    },
    "tcpMesh": {},
    "routeGroups": {}
  },
  "policies": [],
  "backends": [],
  "version": {
    "version": "89f3f08f22",
    "git_revision": "89f3f08f22d93ce9af419e84cff7b412497fec80",
    "rust_version": "1.97.0",
    "build_profile": "release",
    "build_target": "x86_64-unknown-linux-gnu"
  },
  "config": {
    "ipv6Enabled": true,
    "network": "",
    "selfIdentity": {
      "wds": {
        "name": "agw-waypoint-64db89d659-qnqk7",
        "namespace": "agw-waypoint-mesh-e2e",
        "cluster_id": "Kubernetes"
      }
    },
    "terminationMaxDeadline": "55s",
    "terminationMinDeadline": "10s",
    "numWorkerThreads": 16,
    "adminAddr": {
      "Localhost": [
        true,
        15000
      ]
    },
    "statsAddr": {
      "SocketAddr": "[::]:15020"
    },
    "readinessAddr": {
      "SocketAddr": "[::]:15021"
    },
    "selfAddr": {
      "gateway": "agw-waypoint",
      "namespace": "agw-waypoint-mesh-e2e"
    },
    "hbone": {
      "windowSize": 4194304,
      "connectionWindowSize": 16777216,
      "frameSize": 1048576,
      "poolMaxStreamsPerConn": 100,
      "poolUnusedReleaseTimeout": {
        "secs": 300,
        "nanos": 0
      }
    },
    "xds": {
      "address": "https://agentgateway.agentgateway-system.svc.cluster.local:9978",
      "auth": {
        "token": [
          "./var/run/secrets/xds-tokens/xds-token",
          "Kubernetes"
        ]
      },
      "caCert": {
        "File": "/etc/xds-tls/ca.crt"
      },
      "namespace": "agw-waypoint-mesh-e2e",
      "gateway": "agw-waypoint",
      "localConfig": "/config/config.yaml"
    },
    "ca": {
      "address": "https://istiod.istio-system.svc:15012",
      "secretTtl": "24h0m0s",
      "identity": "spiffe://cluster.local/ns/agw-waypoint-mesh-e2e/sa/agw-waypoint",
      "auth": {
        "token": [
          "./var/run/secrets/tokens/istio-token",
          "Kubernetes"
        ]
      },
      "caCert": {
        "File": "./var/run/secrets/istio/root-cert.pem"
      },
      "caHeaders": [],
      "allowedTrustDomains": [
        "cluster.local"
      ],
      "skipValidateTrustDomain": false
    },
    "tracing": null,
    "metrics": {
      "metric_fields": {
        "add": {}
      },
      "excluded_metrics": []
    },
    "logging": {
      "filter": null,
      "fields": {
        "remove": [],
        "add": {}
      },
      "database_fields": {
        "remove": [],
        "add": {}
      },
      "level": "",
      "format": "text",
      "database": null
    },
    "database": null,
    "storage": {
      "mode": "file"
    },
    "dns": {
      "resolverCfg": {
        "domain": null,
        "search": [
          "agw-waypoint-mesh-e2e.svc.cluster.local",
          "svc.cluster.local",
          "cluster.local"
        ],
        "name_servers": [
          {
            "ip": "10.96.0.10",
            "trust_negative_responses": true,
            "connections": [
              {
                "port": 53,
                "protocol": {
                  "type": "udp"
                },
                "bind_addr": null
              },
              {
                "port": 53,
                "protocol": {
                  "type": "tcp"
                },
                "bind_addr": null
              }
            ]
          }
        ]
      },
      "resolverOpts": {
        "ndots": 5,
        "timeout": 5,
        "attempts": 2,
        "edns0": false,
        "ip_strategy": "Ipv4thenIpv6",
        "cache_size": 8192,
        "use_hosts_file": "Auto",
        "positive_min_ttl": null,
        "negative_min_ttl": null,
        "positive_max_ttl": null,
        "negative_max_ttl": null,
        "num_concurrent_reqs": 2,
        "max_active_requests": 32,
        "preserve_intermediates": true,
        "try_tcp_on_error": false,
        "server_ordering_strategy": "QueryStatistics",
        "recursion_desired": true,
        "avoid_local_udp_ports": [],
        "os_port_selection": false,
        "case_randomization": false,
        "trust_anchor": null,
        "allow_answers": [],
        "deny_answers": [],
        "edns_payload_len": 1232
      }
    },
    "proxyMetadata": {
      "instanceIp": "10.244.0.48",
      "podName": "agw-waypoint-64db89d659-qnqk7",
      "podNamespace": "agw-waypoint-mesh-e2e",
      "nodeName": "kind-control-plane",
      "role": "agw-waypoint-mesh-e2e~agw-waypoint",
      "nodeId": "agentgateway~10.244.0.48~agw-waypoint-64db89d659-qnqk7.agw-waypoint-mesh-e2e~agw-waypoint-mesh-e2e.svc.cluster.local"
    },
    "threadingMode": "multithreaded",
    "sessionEncoder": "aes",
    "oidcCookieEncoder": null,
    "backend": {
      "keepalives": {
        "enabled": true,
        "time": "3m0s",
        "interval": "3m0s",
        "retries": 9
      },
      "connectTimeout": "10s",
      "poolIdleTimeout": "1m30s",
      "poolMaxSize": null
    },
    "mcp": {
      "sessionTtl": "30m0s"
    },
    "dynamicCaCertCache": {
      "ttl": "5m0s",
      "capacity": 256
    },
    "modelCatalog": {
      "sources": []
    }
  }
}
```
