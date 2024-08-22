package router

import (
	"net/http"

	"github.com/suryaapandi28/dev_kredit_plus/internal/http/handler"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/route"
)

const (
	Admin = "admin"
	User  = "user"
)

var (
	onlyUser = []string{User}
)

func PublicRoutes(konsumenHandler handler.KonsumenHandler, accountHandler handler.AccountHandler) []*route.Route {

	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "Account/Registered",
			Handler: accountHandler.Registered,
		},
		{
			Method:  http.MethodPost,
			Path:    "Account/Verification",
			Handler: accountHandler.Verification_account,
		},
	}

}

func PrivateRoutes(konsumenHandler handler.KonsumenHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "Konsumen/Create",
			Handler: konsumenHandler.CreateUser,
			Roles:   onlyUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "Konsumen/CheckLimit",
			Handler: konsumenHandler.CheckLimitTenor,
			Roles:   onlyUser,
		},
	}
}
