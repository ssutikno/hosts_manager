package services

import (
	"errors"
	"sync"
)

// Host represents a host entity.
type Host struct {
	ID   string
	Name string
	IP   string
}

// HostService provides methods for managing hosts.
type HostService struct {
	mu    sync.Mutex
	hosts map[string]Host
}

// NewHostService creates a new instance of HostService.
func NewHostService() *HostService {
	return &HostService{
		hosts: make(map[string]Host),
	}
}

// AddHost adds a new host to the service.
func (s *HostService) AddHost(host Host) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.hosts[host.ID] = host
}

// GetHost retrieves a host by ID.
func (s *HostService) GetHost(id string) (Host, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	host, exists := s.hosts[id]
	if !exists {
		return Host{}, errors.New("host not found")
	}
	return host, nil
}

// ListHosts returns all hosts.
func (s *HostService) ListHosts() []Host {
	s.mu.Lock()
	defer s.mu.Unlock()
	var hostList []Host
	for _, host := range s.hosts {
		hostList = append(hostList, host)
	}
	return hostList
}

// RemoveHost removes a host by ID.
func (s *HostService) RemoveHost(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.hosts[id]; !exists {
		return errors.New("host not found")
	}
	delete(s.hosts, id)
	return nil
}