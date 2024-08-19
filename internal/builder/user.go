package builder

import (
	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/internal/http/router"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/encrypt"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/route"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/token"

	// "github.com/labstack/echo/"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB, redisDB *redis.Client, tokenUseCase token.TokenUseCase, encryptTool encrypt.EncryptTool,
	entityCfg *entity.Config) []*route.Route {

	return router.PublicRoutes()
}

func BuildPrivateRoutes(db *gorm.DB, redisDB *redis.Client, encryptTool encrypt.EncryptTool, entityCfg *entity.Config, tokenUseCase token.TokenUseCase) []*route.Route {

	return router.PrivateRoutes()
}
