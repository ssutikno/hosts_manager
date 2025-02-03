package controllers

import (
	"net/http"

	"host-manager/services"

	"github.com/gin-gonic/gin"
)

// HostController handles host-related requests
type HostController struct {
	HostService services.HostService
}

// NewHostController creates a new HostController
func NewHostController(hostService services.HostService) *HostController {
	return &HostController{HostService: hostService}
}

// ListHosts displays a list of registered hosts
func (hc *HostController) ListHosts(c *gin.Context) {
	hosts, err := hc.HostService.GetAllHosts()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "hosts/list.html", gin.H{"hosts": hosts})
}

// AddHost handles adding a new host
func (hc *HostController) AddHost(c *gin.Context) {
	// ... (Get host data from request, validate, and save)
}

// Other host management functions (GetHost, UpdateHost, DeleteHost, etc.)
// ...
