/*
 * Nnwdaf_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPCreateSessionManagementData - Creates and updates the session management data for a UE and for an individual PDU session
func HTTPCreateSessionManagementData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// HTTPDeleteSessionManagementData - Deletes the session management data for a UE and for an individual PDU session
func HTTPDeleteSessionManagementData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// HTTPQuerySessionManagementData - Retrieves the session management data for a UE and for an individual PDU session
func HTTPQuerySessionManagementData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
