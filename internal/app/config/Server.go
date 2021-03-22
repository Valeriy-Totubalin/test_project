package config

import "time"

type Server struct {
	port         string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func NewServer() *Server {
	return &Server{
		port:         getEnv("SRV_PORT", "8080"),
		readTimeout:  getEnvAsSeconds("SRV_READ_TIMEOUT", 10*time.Second),
		writeTimeout: getEnvAsSeconds("SRV_WRITE_TIMEOUT", 10*time.Second),
	}
}

func (s *Server) GetPort() string {
	return s.port
}

func (s *Server) GetReadTimeout() time.Duration {
	return s.readTimeout
}

func (s *Server) GetWriteTimeout() time.Duration {
	return s.writeTimeout
}
