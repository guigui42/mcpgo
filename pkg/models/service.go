package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Technologies represents the technologies used by a service
type Technologies struct {
	Languages []string `json:"languages"`
	Framework []string `json:"framework"`
	DB        []string `json:"db"`
}

// Service represents a service with its details
type Service struct {
	ID            int          `json:"id"`
	TitleEN       string       `json:"title_en"`
	TitleFR       string       `json:"title_fr"`
	DescriptionEN string       `json:"description_en"`
	DescriptionFR string       `json:"description_fr"`
	Owner         string       `json:"owner"`
	EmailContact  string       `json:"email_contact"`
	Git           string       `json:"git"`
	Technologies  Technologies `json:"technologies"`
}

// Cache for services to avoid reading the file multiple times
var (
	servicesData []Service
	servicesMux  sync.RWMutex
	dataLoaded   bool
)

// ReadServices is the direct equivalent of the TypeScript readServices function
// This function reads the services data from the JSON file, validates it and returns it
func ReadServices() ([]Service, error) {
	return loadServices()
}

// loadServices loads service data from the JSON file
func loadServices() ([]Service, error) {
	servicesMux.RLock()
	if dataLoaded {
		defer servicesMux.RUnlock()
		return servicesData, nil
	}
	servicesMux.RUnlock()

	// Get the executable path
	execPath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("error getting executable path: %v", err)
	}

	// Try multiple possible locations for the data file
	possiblePaths := []string{
		filepath.Join("data", "services.json"),                         // Relative to current working directory
		filepath.Join(filepath.Dir(execPath), "data", "services.json"), // Relative to executable
		filepath.Join("..", "data", "services.json"),                   // One level up
		filepath.Join("..", "..", "data", "services.json"),             // Two levels up
	}

	var jsonFile *os.File

	for _, path := range possiblePaths {
		jsonFile, err = os.Open(path)
		if err == nil {
			break
		}
	}

	if jsonFile == nil {
		return nil, fmt.Errorf("could not find services.json in any of the expected locations")
	}
	defer jsonFile.Close()

	// Read and parse JSON
	var services []Service
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&services); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	// Cache the data
	servicesMux.Lock()
	servicesData = services
	dataLoaded = true
	servicesMux.Unlock()

	return services, nil
}

// GetAllServices returns all services
func GetAllServices() ([]Service, error) {
	return loadServices()
}

// GetServiceByID returns a service by ID
func GetServiceByID(id int) (*Service, error) {
	services, err := loadServices()
	if err != nil {
		return nil, err
	}

	for _, service := range services {
		if service.ID == id {
			return &service, nil
		}
	}
	return nil, nil // Return nil if not found
}

// GetServicesByOwner returns services filtered by owner
func GetServicesByOwner(owner string) ([]Service, error) {
	services, err := loadServices()
	if err != nil {
		return nil, err
	}

	var result []Service
	for _, service := range services {
		if service.Owner == owner {
			result = append(result, service)
		}
	}
	return result, nil
}
