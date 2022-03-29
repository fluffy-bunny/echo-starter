package main

import (
	"echo-starter/internal/templates"
	"fmt"
	"net/http"

	contracts_config "echo-starter/internal/contracts/config"
	contracts_handler "echo-starter/internal/contracts/handler"
	middleware_container "echo-starter/internal/middleware/container"
	middleware_session "echo-starter/internal/middleware/session"
	services_container "echo-starter/internal/services/container"
	"encoding/base64"

	"github.com/gorilla/securecookie"
	"github.com/quasoft/memstore"

	"echo-starter/internal/shared"
	echostarter_utils "echo-starter/internal/utils"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/core"
	"github.com/google/uuid"

	"echo-starter/internal/wellknown"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var version = "Development"

func main() {
	appInstanceID := uuid.New().String()

	fmt.Println("Version:" + version)
	DumpPath("./")
	var err error
	startup := NewStartup()
	configOptions := startup.GetConfigOptions()
	err = core.LoadConfig(configOptions)
	if err != nil {
		panic(err)
	}
	appConfig := configOptions.Destination.(*contracts_config.Config)
	if core_utils.IsEmptyOrNil(appConfig.Oidc.CallbackURL) {
		appConfig.Oidc.CallbackURL = fmt.Sprintf("http://localhost:%v%s",
			appConfig.Port,
			wellknown.OIDCCallbackPath)
	}
	fmt.Println(echostarter_utils.PrettyJSON(appConfig))
	builder, _ := di.NewBuilder(di.App, di.Request, "transient")
	services_container.AddContainerAccessorFunc(builder)
	err = startup.ConfigureServices(builder)
	if err != nil {
		panic(err)
	}
	shared.RootContainer = builder.Build()

	e := echo.New()
	//Set Renderer
	e.Renderer = templates.GetTemplateRender("./templates")

	// SECURE COOKIE
	if core_utils.IsEmptyOrNil(appConfig.SecureCookieHashKey) {
		fmt.Println("WARNING: SECURE_COOKIE_HASH_KEY must be set for production......")
		key := securecookie.GenerateRandomKey(32)
		encodedString := base64.StdEncoding.EncodeToString(key)
		appConfig.SecureCookieHashKey = encodedString
		fmt.Printf("SECURE_COOKIE_HASH_KEY: %v\n", appConfig.SecureCookieHashKey)
	}
	if core_utils.IsEmptyOrNil(appConfig.SecureCookieEncryptionKey) {
		fmt.Println("WARNING: SECURE_COOKIE_ENCRYPTION_KEY must be set for production......")
		key := securecookie.GenerateRandomKey(32)
		encodedString := base64.StdEncoding.EncodeToString(key)
		appConfig.SecureCookieEncryptionKey = encodedString
		fmt.Printf("SECURE_COOKIE_ENCRYPTION_KEY: %v\n", appConfig.SecureCookieEncryptionKey)
	}
	e.Use(middleware.Logger())
	e.Use(middleware_container.EnsureScopedContainer(shared.RootContainer))

	sessionMemStore := memstore.NewMemStore(
		[]byte(appConfig.SessionKey), []byte(appConfig.SessionEncryptionKey),
	)
	sessionMemStore.Options.Secure = true
	sessionMemStore.Options.HttpOnly = true
	sessionMemStore.Options.SameSite = http.SameSiteStrictMode
	sessionMemStore.Options.MaxAge = appConfig.SessionMaxAgeSeconds

	e.Use(session.Middleware(sessionMemStore))
	e.Use(middleware_session.EnsureSlidingSession(shared.RootContainer))

	if appConfig.ApplicationEnvironment == contracts_config.Environment_Development {
		e.Use(middleware_session.EnsureDevelopmentSession(appInstanceID))
	}

	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "header:X-XSRF-TOKEN",
		CookiePath:     "/api",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	app := e.Group("")
	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "form:csrf",
		CookiePath:     "/",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	startup.Configure(e, shared.RootContainer)
	e.Use(middleware.Recover())

	e.Static("/css", "./css")
	e.Static("/assets", "./assets")
	e.Static("/js", "./js")
	// TODO: need to register API handler separately
	handlerFactory := contracts_handler.GetIHandlerFactoryFromContainer(shared.RootContainer)
	handlerFactory.RegisterHandlers(app)

	port := startup.GetPort()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
