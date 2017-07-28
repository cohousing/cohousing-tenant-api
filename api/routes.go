package api

import (
	"github.com/cohousing/cohousing-tenant-api/api/admin"
	"github.com/cohousing/cohousing-tenant-api/api/tenant"
	"github.com/cohousing/cohousing-tenant-api/api/utils"
	"github.com/cohousing/location"
	"github.com/gin-gonic/gin"
)

func CreateRouter() {
	router := gin.Default()
	router.Use(utils.ContextResolver())
	router.Use(location.Default())

	apiV1 := router.Group("api/v1")

	CreateHomeRoutes(apiV1)
	tenant.CreateTenantRoutes(apiV1)
	admin.CreateAdminRoutes(apiV1)

	router.Run(":8080")
}
