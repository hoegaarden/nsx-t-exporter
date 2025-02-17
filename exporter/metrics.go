package exporter

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// GetMetricsDescription - creates Prometheus metrics description
func GetMetricsDescription() map[string]*prometheus.Desc {

	APIMetrics := make(map[string]*prometheus.Desc)

	APIMetrics["ManagementClusterStatus"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster_management", "status"),
		"NSX-T management cluster status - STABLE=1, INITIALIZING=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["ManagementClusterLastSuccessfulConnection"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster_management", "last_successful_data_fetch"),
		"NSX-T last successful data fetch in UNIX timestamp converted to float64",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["ControlClusterStatus"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster_control", "status"),
		"NSX-T control cluster status - STABLE=1, NO_CONTROLLERS=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["ManagementClusterNodesOnline"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster_management", "online_nodes"),
		"NSX-T management cluster online nodes",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["ManagementClusterNodesOffline"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster_management", "offline_nodes"),
		"NSX-T management cluster offline nodes",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["ManagementNodeConnectivity"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "connectivity"),
		"NSX-T management node connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeVersion"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "version"),
		"NSX-T management node version",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip", "nsxv3_node_version"}, nil,
	)

	APIMetrics["ManagementNodeCpuCores"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "cpu_cores"),
		"NSX-T management node cpu cores",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeLoadAverage"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "load_average"),
		"NSX-T management node average load",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip", "minutes"}, nil,
	)

	APIMetrics["ManagementNodeMemoryUse"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "memory_use"),
		"NSX-T management node memory use",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeMemoryTotal"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "memory_total"),
		"NSX-T management node memory total",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeMemoryCached"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "memory_cached"),
		"NSX-T management node cached memory",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeSwapUse"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "swap_use"),
		"NSX-T management node swap use",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeSwapTotal"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "swap_total"),
		"NSX-T management node swap total",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeStorageUse"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "storage_use"),
		"NSX-T management node storage use",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip", "filesystem"}, nil,
	)

	APIMetrics["ManagementNodeStorageTotal"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "storage_total"),
		"NSX-T management node storage total",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip", "filesystem"}, nil,
	)

	APIMetrics["ManagerNodeFirewallTotalSectionCount"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node_firewall", "total_section_count"),
		"NSX-T management node firewall all sections type count",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagerNodeFirewallSectionCount"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node_firewall", "section_count"),
		"NSX-T management node firewall L3 section count",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagerNodeFirewallRuleCount"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node_firewall", "rule_count"),
		"NSX-T management node firewall L3 rule count",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ControlNodeConnectivity"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "control_node", "connectivity"),
		"NSX-T control node connectivity - CONNECTED=1, DISCONNECTED=0, UNKNOWN=-1",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ControlNodeManagementConnectivity"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "control_node", "management_connectivity"),
		"NSX-T control node management connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["TransportNodesUp"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "transport_nodes", "up"),
		"NSX-T transport nodes with state up",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["TransportNodesDegraded"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "transport_nodes", "degraded"),
		"NSX-T transport nodes with state degraded",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["TransportNodesDown"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "transport_nodes", "down"),
		"NSX-T transport nodes with state down",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["TransportNodesUnknown"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "transport_nodes", "unknown"),
		"NSX-T transport nodes with state unknown",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["TransportNodeState"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "transport_node", "state"),
		"NSX-T transport node state - SUCCESS=1, IN_PROGRESS=0, PENDING=-1, FAILED=-2, PARTIAL_SUCCESS=-3, ORPHANED=-4, UNKNOWN=-5",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_id"}, nil,
	)

	APIMetrics["TransportNodeDeploymentState"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "transport_node", "deployment_state"),
		"NSX-T transport node deployment state - SUCCESS=1, IN_PROGRESS=0, PENDING=-1, FAILED=-2, PARTIAL_SUCCESS=-3, ORPHANED=-4, UNKNOWN=-5",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_id"}, nil,
	)

	APIMetrics["LogicalSwitchAdminState"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "logical_switch", "admin_state"),
		"NSX-T logical switch admin state - UP=1, DOWN=0",
		[]string{"nsxv3_manager_hostname", "name", "id"}, nil,
	)

	APIMetrics["IPPoolFreePercentage"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "ip_pool", "percentage_free"),
		"NSX-T IP Pools Percentage Free, 0 = no ips available, 100 = no ips allocated",
		[]string{"nsxv3_manager_hostname", "name", "id", "free", "total"}, nil,
	)

	APIMetrics["LoadBalancers"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "load_balancer", "virtual_server_count"),
		"NSX-T Load Balancers and their virtual servers",
		[]string{"nsxv3_manager_hostname", "enabled", "name", "id", "size"}, nil,
	)

	APIMetrics["LoadBalancerSummarySmall"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "load_balancer", "count_small"),
		"NSX-T number of SMALL load balancers",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["LoadBalancerSummaryMedium"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "load_balancer", "count_medium"),
		"NSX-T number of MEDIUM load balancers",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["LoadBalancerSummaryLarge"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "load_balancer", "count_large"),
		"NSX-T number of LARGE load balancers",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["LogicalPortOperationalState"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "logical_port", "operational_state"),
		"NSX-T logical port operational state - UP=1, DOWN=0, UNKNOWN=-1",
		[]string{"nsxv3_manager_hostname", "id", "transport_node_id"}, nil,
	)

	APIMetrics["LogicalSwitchState"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "logical_switch", "state"),
		"NSX-T logical switch overall state -  SUCCESS=1, IN_PROGRESS=0, FAILED=-1, PARTIAL_SUCCESS=-2, ORPHANED=-3, UNKNOWN=-4",
		[]string{"nsxv3_manager_hostname", "id"}, nil,
	)

	APIMetrics["SchedulerTotalComplete"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "scheduler", "total_complete"),
		"NSX-T Scheduler total completed jobs",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["SchedulerTotalExecuting"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "scheduler", "total_executing"),
		"NSX-T Scheduler total executing jobs",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["SchedulerTotalQueued"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "scheduler", "total_queued"),
		"NSX-T Scheduler total queued jobs",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["SchedulerTotalScheduled"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "scheduler", "total_scheduled"),
		"NSX-T Scheduler total scheduled jobs",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["SchedulerTotalSuspended"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "scheduler", "total_suspended"),
		"NSX-T Scheduler total suspended jobs",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	return APIMetrics
}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(data *Nsxv3Data, ch chan<- prometheus.Metric) error {
	if !data.ExtractedActualValues {
		log.Warn("Metrics processing completed with error, will not report any metrics.")
		return nil
	}

	small := Nsxv3LoadBalancerSummary{size: "SMALL", count: 0}
	medium := Nsxv3LoadBalancerSummary{size: "MEDIUM", count: 0}
	large := Nsxv3LoadBalancerSummary{size: "LARGE", count: 0}

	// Prometheus scrape metric callback (concurrent)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementClusterStatus"],
		prometheus.GaugeValue,
		data.ClusterManagementStatus,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ControlClusterStatus"],
		prometheus.GaugeValue,
		data.ClusterControlStatus,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementClusterNodesOnline"],
		prometheus.GaugeValue,
		data.ClusterOnlineNodes,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementClusterNodesOffline"],
		prometheus.GaugeValue,
		data.ClusterOfflineNodes,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementClusterLastSuccessfulConnection"],
		prometheus.GaugeValue,
		data.LastSuccessfulDataFetch,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["TransportNodesUp"],
		prometheus.GaugeValue,
		data.TransportNodesState.UpCount,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["TransportNodesDegraded"],
		prometheus.GaugeValue,
		data.TransportNodesState.DegradedCount,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["TransportNodesDown"],
		prometheus.GaugeValue,
		data.TransportNodesState.DownCount,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["TransportNodesUnknown"],
		prometheus.GaugeValue,
		data.TransportNodesState.UnknownCount,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["SchedulerTotalComplete"],
		prometheus.GaugeValue,
		data.Scheduler.TotalComplete,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["SchedulerTotalExecuting"],
		prometheus.GaugeValue,
		data.Scheduler.TotalExecuting,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["SchedulerTotalQueued"],
		prometheus.GaugeValue,
		data.Scheduler.TotalQueued,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["SchedulerTotalScheduled"],
		prometheus.GaugeValue,
		data.Scheduler.TotalScheduled,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["SchedulerTotalSuspended"],
		prometheus.GaugeValue,
		data.Scheduler.TotalSuspended,
		data.ClusterHost)

	for _, element := range data.ManagementNodes {
		e.processManagementNodeMetrics(data.ClusterHost, &element, ch)
	}

	for _, element := range data.ControlNodes {
		e.processControlNodeMetrics(data.ClusterHost, &element, ch)
	}

	for _, element := range data.TransportNodes {
		e.processTransportNodeMetrics(data.ClusterHost, &element, ch)
	}

	for _, element := range data.LogicalSwitchesAdminStates {
		e.processLogicalSwitchAdminStateMetrics(data.ClusterHost, &element, ch)
	}

	for _, element := range data.LogicalSwitchesStates {
		e.processLogicalSwitchStateMetrics(data.ClusterHost, &element, ch)
	}

	for _, element := range data.IPPools {
		e.processIPPoolList(data.ClusterHost, &element, ch)
	}

	for _, element := range data.LoadBalancers {

		e.processLoadBalancers(data.ClusterHost, &element, ch)

		/* TODO: there's got to be a better way to do this, but for now.. */
		if element.size == "SMALL" {
			small.count += 1
		} else if element.size == "MEDIUM" {
			medium.count += 1
		} else if element.size == "LARGE" {
			large.count += 1
		}
	}

	e.processLoadBalancerSmallSummaries(data.ClusterHost, &small, ch)
	e.processLoadBalancerMediumSummaries(data.ClusterHost, &medium, ch)
	e.processLoadBalancerLargeSummaries(data.ClusterHost, &large, ch)

	for _, element := range data.LogicalPortOperationalStates {
		e.processLogicalPortOperationalStateMetrics(data.ClusterHost, &element, ch)
	}

	return nil

}

func (e *Exporter) processManagementNodeMetrics(host string, data *Nsxv3ManagementNodeData, ch chan<- prometheus.Metric) error {
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeConnectivity"],
		prometheus.GaugeValue,
		data.Connectivity,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeVersion"],
		prometheus.CounterValue,
		1,
		host, data.IP, data.Version)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeCpuCores"],
		prometheus.GaugeValue,
		data.CPUCores,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeLoadAverage"],
		prometheus.GaugeValue,
		data.LoadAverage[0],
		host, data.IP, "1") // 1 minute
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeLoadAverage"],
		prometheus.GaugeValue,
		data.LoadAverage[1],
		host, data.IP, "5") // 5 minutes
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeLoadAverage"],
		prometheus.GaugeValue,
		data.LoadAverage[2],
		host, data.IP, "15") // 15 minutes

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeMemoryUse"],
		prometheus.GaugeValue,
		data.MemoryUse,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeMemoryTotal"],
		prometheus.GaugeValue,
		data.MemoryTotal,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeMemoryCached"],
		prometheus.GaugeValue,
		data.MemoryCached,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeSwapUse"],
		prometheus.GaugeValue,
		data.SwapUse,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeSwapTotal"],
		prometheus.GaugeValue,
		data.SwapTotal,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagerNodeFirewallTotalSectionCount"],
		prometheus.GaugeValue,
		data.TotalSectionCount,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagerNodeFirewallSectionCount"],
		prometheus.GaugeValue,
		data.L3DFWSectionCount,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagerNodeFirewallRuleCount"],
		prometheus.GaugeValue,
		data.L3DFWRuleCount,
		host, data.IP)

	for _, element := range data.Storage {
		ch <- prometheus.MustNewConstMetric(
			e.APIMetrics["ManagementNodeStorageTotal"],
			prometheus.GaugeValue,
			element.totalMetric,
			host, data.IP, element.filesystem)
		ch <- prometheus.MustNewConstMetric(
			e.APIMetrics["ManagementNodeStorageUse"],
			prometheus.GaugeValue,
			element.usedMetric,
			host, data.IP, element.filesystem)
	}

	return nil
}

func (e *Exporter) processControlNodeMetrics(host string, data *Nsxv3ControlNodeData, ch chan<- prometheus.Metric) error {

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ControlNodeConnectivity"],
		prometheus.GaugeValue,
		data.Connectivity,
		host, data.IP)

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ControlNodeManagementConnectivity"],
		prometheus.GaugeValue,
		data.ManagementConnectivity,
		host, data.IP)

	return nil
}

func (e *Exporter) processTransportNodeMetrics(host string, data *Nsxv3TransportNodeData, ch chan<- prometheus.Metric) error {

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["TransportNodeState"],
		prometheus.GaugeValue,
		data.State,
		host, data.ID)

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["TransportNodeDeploymentState"],
		prometheus.GaugeValue,
		data.DeploymentState,
		host, data.ID)

	return nil
}

func (e *Exporter) processLogicalSwitchAdminStateMetrics(host string, data *Nsxv3LogicalSwitchAdminStateData, ch chan<- prometheus.Metric) error {

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["LogicalSwitchAdminState"],
		prometheus.GaugeValue,
		data.adminStateMetric,
		host, data.name, data.id)

	return nil
}

func (e *Exporter) processLogicalPortOperationalStateMetrics(host string, data *Nsxv3LogicalPortOperationalStateData, ch chan<- prometheus.Metric) error {
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["LogicalPortOperationalState"],
		prometheus.GaugeValue,
		data.operationalStateMetric,
		host, data.id, data.hostID)

	return nil
}

func (e *Exporter) processLogicalSwitchStateMetrics(host string, data *Nsxv3LogicalSwitchStateData, ch chan<- prometheus.Metric) error {

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["LogicalSwitchState"],
		prometheus.GaugeValue,
		data.stateMetric,
		host, data.id)

	return nil
}

func (e *Exporter) processIPPoolList(host string, data *Nsxv3IPPoolItem, ch chan<- prometheus.Metric) error {

	var percentFree float64 = 0

	if data.totalIds == 0 {
		percentFree = 0
	} else if data.freeIds == 0 {
		percentFree = 0
	} else {
		percentFree = 1 * 100 * (data.freeIds / data.totalIds)
	}

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["IPPoolFreePercentage"],
		prometheus.GaugeValue,
		percentFree,
		host,
		data.name,
		data.id,
		strconv.FormatFloat(data.freeIds, 'f', 0, 64),
		strconv.FormatFloat(data.totalIds, 'f', 0, 64))

	return nil
}

func (e *Exporter) processLoadBalancers(host string, data *Nsxv3LoadBalancer, ch chan<- prometheus.Metric) error {

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["LoadBalancers"],
		prometheus.GaugeValue,
		float64(data.count),
		host,
		strconv.FormatBool(data.enabled),
		data.name,
		data.id,
		data.size)

	return nil
}

func (e *Exporter) processLoadBalancerSmallSummaries(host string, data *Nsxv3LoadBalancerSummary, ch chan<- prometheus.Metric) error {

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["LoadBalancerSummarySmall"],
		prometheus.GaugeValue,
		float64(data.count),
		host)

	return nil
}

func (e *Exporter) processLoadBalancerMediumSummaries(host string, data *Nsxv3LoadBalancerSummary, ch chan<- prometheus.Metric) error {
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["LoadBalancerSummaryMedium"],
		prometheus.GaugeValue,
		float64(data.count),
		host)

	return nil
}

func (e *Exporter) processLoadBalancerLargeSummaries(host string, data *Nsxv3LoadBalancerSummary, ch chan<- prometheus.Metric) error {
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["LoadBalancerSummaryLarge"],
		prometheus.GaugeValue,
		float64(data.count),
		host)

	return nil
}
