package builder

import (
	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/internal/http/handler"
	"github.com/suryaapandi28/dev_kredit_plus/internal/http/router"
	"github.com/suryaapandi28/dev_kredit_plus/internal/repository"
	"github.com/suryaapandi28/dev_kredit_plus/internal/service"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/cache"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/email"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/encrypt"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/route"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/token"

	// "github.com/labstack/echo/"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB, redisDB *redis.Client, tokenUseCase token.TokenUseCase, encryptTool encrypt.EncryptTool,
	entityCfg *entity.Config) []*route.Route {
	cacheable := cache.NewCacheable(redisDB)
	emailService := email.NewEmailSender(entityCfg)

	konsumenRepository := repository.NewKonsumenRepository(db, cacheable)
	konsumenService := service.NewKonsumenService(konsumenRepository)
	konsumenHandler := handler.NewKonsumenHandler(konsumenService)

	accountRepository := repository.NewAccountRepository(db, cacheable)
	accountService := service.NewAccountService(accountRepository, tokenUseCase, encryptTool, emailService)
	accountHandler := handler.NewAccountHandler(accountService)
	return router.PublicRoutes(konsumenHandler, accountHandler)
}

func BuildPrivateRoutes(db *gorm.DB, redisDB *redis.Client, encryptTool encrypt.EncryptTool, entityCfg *entity.Config, tokenUseCase token.TokenUseCase) []*route.Route {
	cacheable := cache.NewCacheable(redisDB)
	//Konsumen
	konsumenRepository := repository.NewKonsumenRepository(db, cacheable)
	konsumenService := service.NewKonsumenService(konsumenRepository)
	konsumenHandler := handler.NewKonsumenHandler(konsumenService)

	transactonRepository := repository.NewTransactionRepository(db, cacheable)
	transactonService := service.NewTransactionService(transactonRepository)
	transactonHandler := handler.NewTransactionHandler(transactonService, konsumenRepository)

	return router.PrivateRoutes(konsumenHandler, transactonHandler)
}
