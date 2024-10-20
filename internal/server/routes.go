package server

import (
	"easy-life-back-go/internal/server/common"
	"easy-life-back-go/internal/server/routes/auth"
)

func registerRoutes(params *common.RoutesParams) {
	auth.RegisterRoutes(params)
}
