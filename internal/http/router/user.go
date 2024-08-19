package router

import (
	"github.com/suryaapandi28/dev_kredit_plus/pkg/route"
)

const (
	Admin = "admin"
	User  = "user"
)

var (
	allRoles  = []string{Admin, User}
	onlyAdmin = []string{Admin}
	onlyUser  = []string{User}
)

func PublicRoutes() []*route.Route {
	return []*route.Route{}
}

func PrivateRoutes() []*route.Route {
	return []*route.Route{}
}
