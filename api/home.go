package api

import (
	"fmt"
	"github.com/cohousing/location"
	"github.com/gin-gonic/gin"
	"net/http"
	domain2 "github.com/cohousing/cohousing-api-utils/domain"
	"github.com/cohousing/cohousing-tenant-api/domain"
)

type TenantHome struct {
	Context string `json:"context"`
	ApiUrl  string `json:"apiurl"`
	Name    string `json:"name"`
	domain2.DefaultHalResource
}

func CreateHomeRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		basePath := router.BasePath()
		t := GetTenantFromContext(c)
		url := location.Get(c)

		tenantHome := &TenantHome{
			Context: t.Context,
			ApiUrl:  fmt.Sprintf("%s://%s%s", url.Scheme, url.Host, basePath),
			Name:    t.Name,
		}

		tenantHome.AddLink(domain2.REL_SELF, basePath)
		tenantHome.AddLink(domain.REL_APARTMENTS, ApartmentBasePath)
		tenantHome.AddLink(domain.REL_RESIDENTS, ResidentBasePath)
		tenantHome.AddLink(domain.REL_USERS, UserBasePath)
		tenantHome.AddLink(domain.REL_GROUPS, GroupBasePath)

		c.JSON(http.StatusOK, tenantHome)
	})
}
