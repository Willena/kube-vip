package kubevip

// Environment variables
const (

	// vipArp - defines if the arp broadcast should be enabled
	vipArp = "vip_arp"

	// vip_arpRate - defines the rate of gARP broadcasts
	vipArpRate = "vip_arpRate"

	// vipLeaderElection - defines if the kubernetes algorithm should be used
	vipLeaderElection = "vip_leaderelection"

	// vipLeaseName - defines the name of the lease lock
	vipLeaseName = "vip_leasename"

	// vipLeaderElection - defines if the kubernetes algorithm should be used
	vipLeaseDuration = "vip_leaseduration"

	// vipLeaderElection - defines if the kubernetes algorithm should be used
	vipRenewDeadline = "vip_renewdeadline"

	// vipLeaderElection - defines if the kubernetes algorithm should be used
	vipRetryPeriod = "vip_retryperiod"

	// vipLeaderElection - defines the annotations given to the lease lock
	vipLeaseAnnotations = "vip_leaseannotations"

	// vipLogLevel - defines the level of logging to produce (5 being the most verbose)
	vipLogLevel = "vip_loglevel"

	// vipInterface - defines the interface that the vip should bind too
	vipInterface = "vip_interface"

	// vipInterfaceLoGlobal - defines if the lo interface (if used) should have a global scope
	vipInterfaceLoGlobal = "vip_interfaceloglobal"

	// vipServicesInterface - defines the interface that the service vips should bind too
	vipServicesInterface = "vip_servicesinterface"

	// vipSubnet - defines the subnet that the vip will use
	vipSubnet = "vip_subnet"

	// egressPodCidr - defines the cidr that egress will ignore
	egressPodCidr = "egress_podcidr"

	// egressServiceCidr - defines the cidr that egress will ignore
	egressServiceCidr = "egress_servicecidr"

	// egressWithNftables - enables using nftables over iptables
	egressWithNftables = "egress_withnftables"

	/////////////////////////////////////
	// TO DO:
	// Determine how to tidy this mess up
	/////////////////////////////////////

	// vipAddress - defines the address that the vip will expose
	// DEPRECATED: will be removed in a next release
	vipAddress = "vip_address"

	// address - defines the address that would be used as a vip
	// it may be an IP or a DNS name, in case of a DNS name
	// kube-vip will try to resolve it and use the IP as a VIP
	address = "address"

	// port - defines the port for the VIP
	port = "port"

	// annotations
	annotations = "annotation"

	// vipDdns - defines if use dynamic dns to allocate IP for "address"
	vipDdns = "vip_ddns"

	// vipLeaseNodeName defines the node name that is used to acquire leases
	nodeName = "vip_nodename"

	// vipSingleNode - defines the vip start as a single node cluster
	vipSingleNode = "vip_singlenode"

	// vipStartLeader - will start this instance as the leader of the cluster
	vipStartLeader = "vip_startleader"

	// bgpEnable defines if BGP should be enabled
	bgpEnable = "bgp_enable"
	// bgpRouterID defines the routerID for the BGP server
	bgpRouterID = "bgp_routerid"
	// bgpRouterInterface defines the interface that we can find the address for
	bgpRouterInterface = "bgp_routerinterface"
	// bgpRouterAS defines the AS for the BGP server
	bgpRouterAS = "bgp_as"
	// bgpPeerAddress defines the address for a BGP peer
	bgpPeerAddress = "bgp_peeraddress"
	// bgpPeers defines the address for a BGP peer
	bgpPeers = "bgp_peers"
	// bgpPeerAS defines the AS for a BGP peer
	bgpPeerAS = "bgp_peeras"
	// bgpPeerAS defines the AS for a BGP peer
	bgpPeerPassword = "bgp_peerpass" // nolint
	// bgpMultiHop enables mulithop routing
	bgpMultiHop = "bgp_multihop"
	// bgpSourceIF defines the source interface for BGP peering
	bgpSourceIF = "bgp_sourceif"
	// bgpSourceIP defines the source address for BGP peering
	bgpSourceIP = "bgp_sourceip"
	// bgpHoldTime defines bgp timers hold time
	bgpHoldTime = "bgp_hold_time"
	// bgpKeepaliveInterval defines bgp timers keepalive interval
	bgpKeepaliveInterval = "bgp_keepalive_interval"

	// zebraEnable defines if Zebra integraton should be enabled
	zebraEnable = "zebra_enable"
	// zebraUrl specifies path to the unix domain socket for connecting to Zebra daemon
	zebraURL = "zebra_url"
	// zebraVersion specifies Zebra API Version
	zebraVersion = "zebra_version"
	// zebraSoftwareName specifies Software Name for Zebra
	zebraSoftwareName = "zebra_software_name"

	// mpbgpNexthop defines MPBGP mode
	mpbgpNexthop = "mpbgp_nexthop"
	// mpbgpIPv4 defines fixed IPv4 to be used with MPBGP
	mpbgpIPv4 = "mpbgp_ipv4"
	// mpbgpIPv6 defines fixed IPv6 to be used with MPBGP
	mpbgpIPv6 = "mpbgp_ipv6"

	// vipWireguard - defines if wireguard will be used for vips
	vipWireguard = "vip_wireguard" //nolint

	// vipRoutingTable - defines if table mode will be used for vips
	vipRoutingTable = "vip_routingtable" //nolint

	// vipRoutingTableID - defines which table mode will be used for vips
	vipRoutingTableID = "vip_routingtableid" //nolint

	// vipRoutingTableType - defines which table type will be used for vip routes
	// 						 valid values for this variable can be found in:
	//						 https://pkg.go.dev/golang.org/x/sys/unix#RTN_UNSPEC
	//						 Note that route type have the prefix `RTN_`, and you
	//						 specify the integer value, not the name. For example:
	//						 you should say `vip_routingtabletype=2` for RTN_LOCAL
	vipRoutingTableType = "vip_routingtabletype" //nolint

	// vipRoutingProtocol - defines what value will be used as protocol when creating routes
	vipRoutingProtocol = "vip_routingprotocol" //nolint

	// vipCleanRoutingTable - defines if routing table will be cleaned of redundant routes on kube-vip's start
	vipCleanRoutingTable = "vip_cleanroutingtable" //nolint

	// cpNamespace defines the namespace the control plane pods will run in
	cpNamespace = "cp_namespace"

	// cpEnable enables the control plane feature
	cpEnable = "cp_enable"

	// cpDetect will attempt to automatically find a working address for the control plane from loopback
	cpDetect = "cp_detect"

	// kubernetesAddr，is the address of the Kubernetes API server on this machine
	kubernetesAddr = "kubernetes_addr"

	// svcEnable enables the Kubernetes service feature
	svcEnable = "svc_enable"

	// svcNamespace defines the namespace the service pods will run in
	svcNamespace = "svc_namespace"

	// svcElection enables election per Kubernetes service
	svcElection = "svc_election"

	// svcLeaseName Name of the lease that is used for leader election for services (in arp mode)
	svcLeaseName = "svc_leasename"

	// lbClassOnly enables load-balancer for class "kube-vip.io/kube-vip-class" only
	lbClassOnly = "lb_class_only"

	// lbClassName enables load-balancer for a specific class only
	lbClassName = "lb_class_name"

	// lbClassLegacyHandling enables legacy handing of load-balancer class
	lbClassLegacyHandling = "lb_class_legacy_handling"

	// lbEnable defines if the load-balancer should be enabled
	lbEnable = "lb_enable"

	// lbPort defines the port of load-balancer
	lbPort = "lb_port"

	// lbForwardingMethod defines the forwarding method of load-balancer
	lbForwardingMethod = "lb_fwdmethod"

	// EnableServiceSecurity defines if the load-balancer should only allow traffic to service ports
	EnableServiceSecurity = "enable_service_security"

	// EnableNodeLabeling, will enable node labeling as the node becomes leader
	EnableNodeLabeling = "enable_node_labeling"

	// prometheusServer defines the address prometheus listens on
	prometheusServer = "prometheus_server"

	// vipConfigMap defines the configmap that kube-vip will watch for service definitions
	// vipConfigMap = "vip_configmap"

	// k8sConfigFile defines the path to the configfile used to speak with the API server
	k8sConfigFile = "k8s_config_file"

	// dnsMode defines mode that DNS lookup will be performed with (first, ipv4, ipv6, dual)
	dnsMode = "dns_mode"

	// disableServiceUpdates disables service updating
	disableServiceUpdates = "disable_service_updates"

	// enableEndpointSlices enables use of EndpointSlices instead of Endpoints
	enableEndpointSlices = "enable_endpointslices"

	// mirrorDestInterface is the network interface where all traffics that go through service interface
	// will be mirrored to. The source interface is ServicesInterface by default, fall back to Interface if not set.
	// + optional
	mirrorDestInterface = "mirror_dest_interface"

	// iptablesBackend iptables backend, can be specified as `nft` or `legacy`. If not set, it defaults to automatic detection.
	iptablesBackend = "iptables_backend"

	// backendHealthCheckInterval Interval in seconds for checking backend health.
	backendHealthCheckInterval = "backend_health_check_interval"

	// healthCheckPort, if set to non-zero will be the port the health check will listen on
	healthCheckPort = "health_check_port"

	// enableUPNP enables UPNP functions
	enableUPNP = "enable_upnp"

	// egressClean enables egress cleaning on kube-vip's start
	egressClean = "egress_clean"
)
