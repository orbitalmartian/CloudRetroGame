package config

import "flag"

type Version int

type Monitoring struct {
	Port             int
	URLPrefix        string
	MetricEnabled    bool `json:"metric_enabled"`
	ProfilingEnabled bool `json:"profiling_enabled"`
}

func (c *Monitoring) IsEnabled() bool { return c.MetricEnabled || c.ProfilingEnabled }

type Server struct {
	Address string
	Https   bool
	Tls     struct {
		Address   string
		Domain    string
		HttpsKey  string
		HttpsCert string
	}
}

type Recording struct {
	Enabled       bool
	CompressLevel int
	Name          string
	Folder        string
	Zip           bool
}

func (s *Server) WithFlags() {
	flag.StringVar(&s.Address, "address", s.Address, "HTTP server address (host:port)")
	flag.StringVar(&s.Tls.Address, "httpsAddress", s.Tls.Address, "HTTPS server address (host:port)")
	flag.StringVar(&s.Tls.HttpsKey, "httpsKey", s.Tls.HttpsKey, "HTTPS key")
	flag.StringVar(&s.Tls.HttpsCert, "httpsCert", s.Tls.HttpsCert, "HTTPS chain")
}

func (s *Server) GetAddr() string {
	if s.Https {
		return s.Tls.Address
	}
	return s.Address
}
