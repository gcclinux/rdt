package main

type Config struct {
	InboundAddress  string `json:"inbound_address"`
	InboundPort     string `json:"inbound_port"`
	OutboundAddress string `json:"outbound_address"`
	OutboundPort    string `json:"outbound_port"`
	Verbose         bool   `json:"verbose"`
}
