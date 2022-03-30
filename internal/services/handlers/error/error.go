package error

import (
	contracts_handler "echo-starter/internal/contracts/handler"
	"echo-starter/internal/templates"
	"echo-starter/internal/wellknown"
	"net/http"
	"reflect"

	contracts_core_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
)

type (
	service struct {
		Logger          contracts_logger.ILogger                        `inject:"logger"`
		ClaimsPrincipal contracts_core_claimsprincipal.IClaimsPrincipal `inject:"claimsPrincipal"`
	}
)

func assertImplementation() {
	var _ contracts_handler.IHandler = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIHandler registers the *service as a singleton.
func AddScopedIHandler(builder *di.Builder) {
	contracts_handler.AddScopedIHandlerEx(builder,
		reflectType,
		[]contracts_handler.HTTPVERB{
			contracts_handler.GET,
		},
		wellknown.ErrorPath)
}

func (s *service) Ctor() {}
func (s *service) GetMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

type params struct {
	Message string `param:"msg" query:"msg" header:"msg" form:"msg" json:"msg" xml:"msg"`
}

func (s *service) Do(c echo.Context) error {

	u := new(params)
	if err := c.Bind(u); err != nil {
		return err
	}

	return templates.Render(c, s.ClaimsPrincipal, http.StatusOK, "views/error/index", map[string]interface{}{
		"params": u,
	})

}
