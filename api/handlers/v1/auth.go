package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/api/models"

	"gitlab.com/QuvonchbekOtajonov/clinic-back/pkg/utils"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/storage/repo"
)

// @Summary 	Login
// @Description This api can login
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param body 	body models.LoginReq true "Body"
// @Success 201 {object} models.LoginRes
// @Failure 400 string Error response
// @Router /v1/login [post]
func (h *handlerV1) Login(ctx *gin.Context) {
	var req models.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "fill all required fields",
		})
		return
	}

	user, err := h.Storage.UnicalPro().GetUserByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "user not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password is wrong",
		})
		return
	}

	token, _, err := utils.CreateToken(h.cfg, &utils.TokenParams{
		UserName: user.Username,
		Duration: h.cfg.TokenTime,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create token",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.LoginRes{
		Token: token,
		Username: user.Username,
	})
}

// @Summary 	Create user
// @Description This api can create new user
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param body 	body models.LoginReq true "Body"
// @Success 201 {object} models.LoginReq
// @Failure 400 string Error response
// @Router /v1/create-user [post]
func (h *handlerV1) CreateNewUser(ctx *gin.Context) {
	var req models.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "fill all required fields",
		})
		return
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password is wrong",
		})
		return
	}

	token, _, err := utils.CreateToken(h.cfg, &utils.TokenParams{
		UserName: req.Username,
		Duration: h.cfg.TokenTime,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create token",
		})
		return
	}

	request := repo.UserReq{
		Username: req.Username,
		Password: password,
	}

	user, err := h.Storage.UnicalPro().CreateNewUser(ctx, &request)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "user not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.LoginRes{
		Token: token,
		Username: user.Username,
	})
}
