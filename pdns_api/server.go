package pdns_api

import (
	"context"
	"errors"
	"fmt"
	stdLog "log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/facebookgo/pidfile"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/pir5/pdns-api/docs" // docs is generated by Swag CLI, you have to import it.
	pmiddleware "github.com/pir5/pdns-api/middleware"
)

const (
	AllowDomainsKey = "AllowDomains"
)

type HTTPError struct {
	echo.HTTPError
}

var CmdServer = &Command{
	Run:       runServer,
	UsageLine: "server",
	Short:     "Start API Server",
	Long: `
Start API Server
	`,
}

func init() {
	// Set your flag here like below.
	// cmdServer.Flag.BoolVar(&flagA, "a", false, "")

}

var globalConfig Config

// runServer executes sub command and return exit code.
func runServer(cmdFlags *GlobalFlags, args []string) error {
	c, err := NewConfig(*cmdFlags.ConfPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	globalConfig = c

	logger := log.New("pdns-api")
	pidfile.SetPidfilePath(*cmdFlags.PidPath)
	if *cmdFlags.LogPath != "" {
		f, err := os.OpenFile(*cmdFlags.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return errors.New("error opening file :" + err.Error())
		}
		logger.SetOutput(f)
	} else {
		logger.SetLevel(log.DEBUG)
	}

	e := echo.New()
	e.Logger = logger
	e.StdLogger = stdLog.New(e.Logger.Output(), e.Logger.Prefix()+": ", 0)

	e.GET("/status", status)

	if err := pidfile.Write(); err != nil {
		return err
	}

	defer func() {
		if err := os.Remove(pidfile.GetPidfilePath()); err != nil {
			e.Logger.Fatalf("Error removing %s: %s", pidfile.GetPidfilePath(), err)
		}
	}()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}}` + "\n",
		Output: logger.Output(),
	}))

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: func(key string, c echo.Context) (bool, error) {
			if globalConfig.IsTokenAuth() {
				for _, v := range globalConfig.Auth.Tokens {
					if key == v {
						return true, nil
					}
				}
			}
			return true, nil

		},
		Skipper: func(c echo.Context) bool {
			return !globalConfig.IsTokenAuth() && !globalConfig.IsHTTPAuth()
		},
	}))
	e.Use(pmiddleware.HeaderAuthWithConfig(pmiddleware.HeaderAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if globalConfig.IsHTTPAuth() {
				domains, err := globalConfig.Auth.HttpAuth.Authenticate(username, password)
				if err != nil {
					return false, err
				}
				c.Set(AllowDomainsKey, domains)
				return true, nil
			}
			return false, nil
		},
		Skipper: func(c echo.Context) bool {
			return !globalConfig.IsTokenAuth() && !globalConfig.IsHTTPAuth()
		},
	}))

	go func() {
		if err := e.Start(globalConfig.Listen); err != nil {
			e.Logger.Fatalf("shutting down the server: %s", err)
		}
	}()

	v1 := e.Group("/v1")

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		globalConfig.DB.UserName,
		globalConfig.DB.Password,
		globalConfig.DB.Host,
		globalConfig.DB.Port,
		globalConfig.DB.DBName,
	))
	if err != nil {
		return err
	}
	DomainEndpoints(v1, db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false))
	RecordEndpoints(v1, db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false))

	v1.GET("/viron", vironGlobalMenu)
	v1.GET("/viron_authtype", vironAuthType)

	v1.GET("/swagger/*", echoSwagger.WrapHandler)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
func status(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// vironAuthType
// @Summary get auth type
// @Description get auth type
// @ID viron_authtype#get
// @Accept  json
// @Produce  json
// @Router /viron_authtype [get]
// @Tags viron
func vironAuthType(c echo.Context) error {
	encodedJSON := []byte(`
[
    {
      "type": "email",
      "provider": "viron-demo",
      "url": "/signin",
      "method": "POST",
    },
    {
      "type": "signout",
      "provider": "",
      "url": "",
      "method": "POST",
    },
]
`)
	return c.JSONBlob(http.StatusOK, encodedJSON)

}

//vironGlobalMenu
// @Summary get global menu
// @Description get global menu
// @ID viron#get
// @Accept json
// @Produce json
// @Router /viron [get]
// @Tags viron
func vironGlobalMenu(c echo.Context) error {
	encodedJSON := []byte(`{
  "theme": "standard",
  "color": "white",
  "name": "Viron example - local",
  "tags": [
    "domains",
    "records"
  ],
  "thumbnail": "https://avatars3.githubusercontent.com/u/23251378?v=3&s=200",
  "pages": [
    {
      "section": "manage",
      "id": "domains",
      "name": "ドメイン操作",
      "components": [
        {
          "api": {
            "method": "get",
            "path": "/domains"
          },
	  "query": [
	    { key: "id", type: "integer" },
	    { key: "name", type: "string" },
          ],
	  "primary": "id",
          "name": "Domain",
	  "style": "table",
          "pagination": true,
	  "table_labels": [
	    "id",
            "name",
            "type",
            "notified_serial",
            "last_check"
	  ]
        }
      ]
    },
    {
      "section": "manage",
      "id": "records",
      "name": "レコード操作",
      "components": [
        {
          "api": {
            "method": "get",
            "path": "/records"
          },
	  "query": [
	    { key: "id", type: "integer" },
	    { key: "domain_id", type: "integer" },
	    { key: "name", type: "string" },
          ],
          "name": "Record",
	  "style": "table",
	  "primary": "id",
	  "table_labels": [
	    "id",
	    "domain_id",
            "name",
            "type",
            "content",
            "ttl",
            "prio",
            "disabled",
            "ordername",
            "auth"
	  ]
        }
      ]
    }
  ]
}`)
	return c.JSONBlob(http.StatusOK, encodedJSON)

}
