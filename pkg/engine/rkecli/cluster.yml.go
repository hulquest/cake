package rkecli

var rawClusterYML = `# If you intended to deploy Kubernetes in an air-gapped environment,
# please consult the documentation on how to configure custom RKE images.
nodes:
- address: ""
  port: "22"
  internal_address: ""
  role:
  - controlplane
  - worker
  - etcd
  hostname_override: ""
  user: rke
  docker_socket: /var/run/docker.sock
  ssh_key: ""
  ssh_key_path: /root/.ssh/id_rsa
  ssh_cert: ""
  ssh_cert_path: ""
  labels: {}
  taints: []
services:
  etcd:
    image: ""
    extra_args: {}
    extra_binds: []
    extra_env: []
    external_urls: []
    ca_cert: ""
    cert: ""
    key: ""
    path: ""
    uid: 0
    gid: 0
    snapshot: null
    retention: ""
    creation: ""
    backup_config: null
  kube-api:
    image: ""
    extra_args: {}
    extra_binds: []
    extra_env: []
    service_cluster_ip_range: 10.43.0.0/16
    service_node_port_range: ""
    pod_security_policy: false
    always_pull_images: false
    secrets_encryption_config: null
    audit_log: null
    admission_configuration: null
    event_rate_limit: null
  kube-controller:
    image: ""
    extra_args: {}
    extra_binds: []
    extra_env: []
    cluster_cidr: 10.42.0.0/16
    service_cluster_ip_range: 10.43.0.0/16
  scheduler:
    image: ""
    extra_args: {}
    extra_binds: []
    extra_env: []
  kubelet:
    image: ""
    extra_args: {}
    extra_binds: []
    extra_env: []
    cluster_domain: cluster.local
    infra_container_image: ""
    cluster_dns_server: 10.43.0.10
    fail_swap_on: false
    generate_serving_certificate: false
  kubeproxy:
    image: ""
    extra_args: {}
    extra_binds: []
    extra_env: []
network:
  plugin: canal
  options: {}
  mtu: 0
  node_selector: {}
  update_strategy: null
authentication:
  strategy: x509
  sans: []
  webhook: null
addons: ""
addons_include: []
system_images:
  etcd: rancher/coreos-etcd:v3.4.3-rancher1
  alpine: rancher/rke-tools:v0.1.56
  nginx_proxy: rancher/rke-tools:v0.1.56
  cert_downloader: rancher/rke-tools:v0.1.56
  kubernetes_services_sidecar: rancher/rke-tools:v0.1.56
  kubedns: rancher/k8s-dns-kube-dns:1.15.0
  dnsmasq: rancher/k8s-dns-dnsmasq-nanny:1.15.0
  kubedns_sidecar: rancher/k8s-dns-sidecar:1.15.0
  kubedns_autoscaler: rancher/cluster-proportional-autoscaler:1.7.1
  coredns: rancher/coredns-coredns:1.6.5
  coredns_autoscaler: rancher/cluster-proportional-autoscaler:1.7.1
  nodelocal: rancher/k8s-dns-node-cache:1.15.7
  kubernetes: rancher/hyperkube:v1.17.4-rancher1
  flannel: rancher/coreos-flannel:v0.11.0-rancher1
  flannel_cni: rancher/flannel-cni:v0.3.0-rancher5
  calico_node: rancher/calico-node:v3.13.0
  calico_cni: rancher/calico-cni:v3.13.0
  calico_controllers: rancher/calico-kube-controllers:v3.13.0
  calico_ctl: rancher/calico-ctl:v2.0.0
  calico_flexvol: rancher/calico-pod2daemon-flexvol:v3.13.0
  canal_node: rancher/calico-node:v3.13.0
  canal_cni: rancher/calico-cni:v3.13.0
  canal_flannel: rancher/coreos-flannel:v0.11.0
  canal_flexvol: rancher/calico-pod2daemon-flexvol:v3.13.0
  weave_node: weaveworks/weave-kube:2.5.2
  weave_cni: weaveworks/weave-npc:2.5.2
  pod_infra_container: rancher/pause:3.1
  ingress: rancher/nginx-ingress-controller:nginx-0.25.1-rancher1
  ingress_backend: rancher/nginx-ingress-controller-defaultbackend:1.5-rancher1
  metrics_server: rancher/metrics-server:v0.3.6
  windows_pod_infra_container: rancher/kubelet-pause:v0.1.3
ssh_key_path: /root/.ssh/id_rsa
ssh_cert_path: ""
ssh_agent_auth: false
authorization:
  mode: none
  options: {}
ignore_docker_version: false
kubernetes_version: ""
private_registries: []
ingress:
  provider: "nginx"
  options: {}
  node_selector: {}
  extra_args: {}
  dns_policy: ""
  extra_envs: []
  extra_volumes: []
  extra_volume_mounts: []
  update_strategy: null
cluster_name: ""
cloud_provider:
  name: ""
prefix_path: ""
addon_job_timeout: 60
bastion_host:
  address: ""
  port: ""
  user: ""
  ssh_key: ""
  ssh_key_path: ""
  ssh_cert: ""
  ssh_cert_path: ""
monitoring:
  provider: ""
  options: {}
  node_selector: {}
  update_strategy: null
  replicas: null
restore:
  restore: false
  snapshot_name: ""
dns: null`
