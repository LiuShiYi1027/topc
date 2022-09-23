package cgroups

type PageStats struct {
	Total uint64           `json:"total,omitempty"`
	Nodes map[uint8]uint64 `json:"nodes,omitempty"`
}

type PageUsageByNUMAInner struct {
	Total       PageStats `json:"total,omitempty"`
	File        PageStats `json:"file,omitempty"`
	Anon        PageStats `json:"anon,omitempty"`
	Unevictable PageStats `json:"unevictable,omitempty"`
}

type PageUsageByNUMA struct {
	// Embedding is used as types can't be recursive.
	PageUsageByNUMAInner
	Hierarchical PageUsageByNUMAInner `json:"hierarchical,omitempty"`
}

type MemoryData struct {
	Usage    uint64 `json:"usage,omitempty"`
	MaxUsage uint64 `json:"max_usage,omitempty"`
	Failcnt  uint64 `json:"failcnt"`
	Limit    uint64 `json:"limit"`
}

type MemoryStats struct {
	// memory used for cache
	Cache uint64 `json:"cache,omitempty"`
	// usage of memory
	Usage MemoryData `json:"usage,omitempty"`
	// usage of memory + swap
	SwapUsage MemoryData `json:"swap_usage,omitempty"`
	// usage of kernel memory
	KernelUsage MemoryData `json:"kernel_usage,omitempty"`
	// usage of kernel TCP memory
	KernelTCPUsage MemoryData `json:"kernel_tcp_usage,omitempty"`
	// usage of memory pages by NUMA node
	// see chapter 5.6 of memory controller documentation
	PageUsageByNUMA PageUsageByNUMA `json:"page_usage_by_numa,omitempty"`
	// if true, memory usage is accounted for throughout a hierarchy of cgroups.
	UseHierarchy bool `json:"use_hierarchy"`

	Stats map[string]uint64 `json:"stats,omitempty"`
}
