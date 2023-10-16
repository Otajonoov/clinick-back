package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/pkg/utils"
)

func (h *handlerV1) AuthMiddleWare(ctx *gin.Context) {
	accessToken := ctx.GetHeader(h.cfg.HeaderKey)

	if len(accessToken) == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "authorization header is not provided",
		})
		return
	}

	payload, err := utils.VerifyToken(h.cfg, accessToken)
	if err != nil {
		h.log.Error(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is invalid",
		})
		return
	}

	ctx.Set(h.cfg.PayloadKey, payload)
	ctx.Next()
}

func (h *handlerV1) GetAuthPayload(ctx *gin.Context) (*utils.Payload, error) {
	i, exist := ctx.Get(h.cfg.PayloadKey)
	fmt.Println(h.cfg.PayloadKey)
	if !exist {
		return nil, errors.New("not found payload")
	}

	payload, ok := i.(*utils.Payload)
	if !ok {
		return nil, errors.New("unknown user")
	}
	return payload, nil
}
