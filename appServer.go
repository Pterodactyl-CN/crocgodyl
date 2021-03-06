package crocgodyl

import (
	"encoding/json"
	"strconv"
	"time"
)

// Application Server API

// Servers is the struct for the servers on the panel.
type Servers struct {
	Object string   `json:"object,omitempty"`
	Server []Server `json:"data,omitempty"`
	Meta   Meta     `json:"meta,omitempty"`
}

// Server is the struct for a server on the panel.
type Server struct {
	Object     string           `json:"object,omitempty"`
	Attributes ServerAttributes `json:"attributes,omitempty"`
}

// ServerAttributes are the attributes for a server.
type ServerAttributes struct {
	ID            int                 `json:"id,omitempty"`
	ExternalID    interface{}         `json:"external_id,omitempty"`
	UUID          string              `json:"uuid,omitempty"`
	Identifier    string              `json:"identifier,omitempty"`
	Name          string              `json:"name,omitempty"`
	Description   string              `json:"description,omitempty"`
	Suspended     bool                `json:"suspended,omitempty"`
	Limits        ServerLimits        `json:"limits,omitempty"`
	FeatureLimits ServerFeatureLimits `json:"feature_limits,omitempty"`
	User          int                 `json:"user,omitempty"`
	Node          int                 `json:"node,omitempty"`
	Allocation    int                 `json:"allocation,omitempty"`
	Nest          int                 `json:"nest,omitempty"`
	Egg           int                 `json:"egg,omitempty"`
	Pack          interface{}         `json:"pack,omitempty"`
	Container     ServerContainer     `json:"container,omitempty"`
	Relationships ServerRelationship  `json:"relationships,omitempty"`
	UpdatedAt     time.Time           `json:"updated_at,omitempty"`
	CreatedAt     time.Time           `json:"created_at,omitempty"`
}

// ServerChange is the struct for the required data for creating/modifying a server.
type ServerChange struct {
	Name          string              `json:"name,omitempty"`
	User          int                 `json:"user,omitempty"`
	Egg           int                 `json:"egg,omitempty"`
	DockerImage   string              `json:"docker_image,omitempty"`
	Startup       string              `json:"startup,omitempty"`
	Environment   map[string]string   `json:"environment,omitempty"`
	Limits        ServerLimits        `json:"limits,omitempty"`
	FeatureLimits ServerFeatureLimits `json:"feature_limits,omitempty"`
	Allocation    ServerAllocation    `json:"allocation,omitempty"`
}

// ServerLimits are the system resource limits for a server
type ServerLimits struct {
	Memory int `json:"memory,omitempty"`
	Swap   int `json:"swap,omitempty"`
	Disk   int `json:"disk,omitempty"`
	Io     int `json:"io,omitempty"`
	CPU    int `json:"cpu,omitempty"`
}

// ServerFeatureLimits this is the limit on Databases and extra Allocations on a server
type ServerFeatureLimits struct {
	Databases   int `json:"databases,omitempty"`
	Allocations int `json:"allocations,omitempty"`
}

// ServerContainer is the config on the docker container the server runs in.
type ServerContainer struct {
	StartupCommand string            `json:"startup_command,omitempty"`
	Image          string            `json:"image,omitempty"`
	Installed      bool              `json:"installed,omitempty"`
	Environment    map[string]string `json:"environment,omitempty"`
}

// ServerRelationship are the relationships for a server.
type ServerRelationship struct {
	Allocations struct {
		Object string          `json:"object,omitempty"`
		Data   []ServerRelData `json:"data,omitempty"`
	} `json:"allocations,omitempty"`
}

// ServerRelData is the data for the server relationship
type ServerRelData struct {
	Object     string                  `json:"object,omitempty"`
	Attributes ServerRelDataAttributes `json:"attributes,omitempty"`
}

// ServerRelDataAttributes are the attributes for the server relationship data
type ServerRelDataAttributes struct {
	ID       int         `json:"id,omitempty"`
	IP       string      `json:"ip,omitempty"`
	Alias    interface{} `json:"alias,omitempty"`
	Port     int         `json:"port,omitempty"`
	Assigned bool        `json:"assigned,omitempty"`
}

// ServerAllocation is only used when creating a server
type ServerAllocation struct {
	Default int `json:"default,omitempty"`
}

// GetServers returns all available servers.
func GetServers() (Servers, error) {
	var servers Servers

	// get json bytes from the panel.
	sbytes, err := queryPanelAPI("servers", "get", nil)
	if err != nil {
		return servers, err
	}

	// Get server info from the panel
	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(sbytes, &servers)
	if err != nil {
		return servers, err
	}

	return servers, nil
}

// GetServer returns Information on a single server.
func GetServer(serverid int) (Server, error) {
	var server Server

	// get json bytes from the panel.
	sbytes, err := queryPanelAPI("servers/"+strconv.Itoa(serverid), "get", nil)
	if err != nil {
		return server, err
	}

	// Get server info from the panel
	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(sbytes, &server)
	if err != nil {
		return server, err
	}

	return server, nil
}

// GetServerAllocations will return a list of allocations for the server in a []int array
func GetServerAllocations(serverid int) ([]int, error) {
	var allServerAlloc []int

	// get json bytes from the panel.
	sabytes, err := queryPanelAPI("servers/"+strconv.Itoa(serverid)+"?include=allocations", "get", nil)
	if err != nil {
		return allServerAlloc, err
	}

	// Get server info from the panel
	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(sabytes, &allServerAlloc)
	if err != nil {
		return allServerAlloc, err
	}

	return allServerAlloc, nil
}

// CreateServer creates a new server via the API.
// A complete ServerChange is required.
func CreateServer(newServer ServerChange) (Server, error) {
	var serverDetails Server

	nsbytes, err := json.Marshal(newServer)
	if err != nil {
		return serverDetails, err
	}

	// get json bytes from the panel.
	sbytes, err := queryPanelAPI("servers", "post", nsbytes)
	if err != nil {
		return serverDetails, err
	}

	// Get server info from the panel
	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(sbytes, &serverDetails)
	if err != nil {
		return serverDetails, err
	}

	return serverDetails, nil
}

// EditServerDetails creates a new server via the API.
// The server name and user are required when updating a server.
func EditServerDetails(newServer ServerChange, serverid int) (Server, error) {
	var serverDetails Server

	esbytes, err := json.Marshal(newServer)
	if err != nil {
		return serverDetails, err
	}

	// get json bytes from the panel.
	sbytes, err := queryPanelAPI("servers/"+strconv.Itoa(serverid)+"/details", "patch", esbytes)
	if err != nil {
		return serverDetails, err
	}

	// Get server info from the panel
	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(sbytes, &serverDetails)
	if err != nil {
		return serverDetails, err
	}

	return serverDetails, nil
}

//TODO: bug dane about this too

// EditServerBuild creates a new server via the API.
// The server name and user are required when updating a server.
func EditServerBuild(newServer ServerChange, serverid int) (Server, error) {
	var serverDetails Server

	esbytes, err := json.Marshal(newServer)
	if err != nil {
		return serverDetails, err
	}

	// get json bytes from the panel.
	sbytes, err := queryPanelAPI("servers/"+strconv.Itoa(serverid)+"/build", "patch", esbytes)
	if err != nil {
		return serverDetails, err
	}

	// Get server info from the panel
	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(sbytes, &serverDetails)
	if err != nil {
		return serverDetails, err
	}

	return serverDetails, nil
}

// EditServerStartup creates a new server via the API.
// The server name and user are required when updating a server.
func EditServerStartup(newServer ServerChange, serverid int) (Server, error) {
	var serverDetails Server

	esbytes, err := json.Marshal(newServer)
	if err != nil {
		return serverDetails, err
	}

	// get json bytes from the panel.
	sbytes, err := queryPanelAPI("servers/"+strconv.Itoa(serverid)+"/startup", "patch", esbytes)
	if err != nil {
		return serverDetails, err
	}

	// Get server info from the panel
	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(sbytes, &serverDetails)
	if err != nil {
		return serverDetails, err
	}

	return serverDetails, nil
}

// DeleteServer deletes a server.
// It only requires a server id as a string
func DeleteServer(serverid int) error {
	// get json bytes from the panel.
	_, err := queryPanelAPI("servers/"+strconv.Itoa(serverid), "delete", nil)
	if err != nil {
		return err
	}

	return nil
}
