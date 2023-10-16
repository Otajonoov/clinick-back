package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/api/models"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/storage/repo"
)

// @Summary 	Get services
// @Description This api can get services
// @Tags 		Service
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} models.AllServices
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/get-services [get]
func (h *handlerV1) GetAllServices(ctx *gin.Context) {
	services, err := h.Storage.UnicalPro().GetAllServices(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get all doctors",
		})
		return
	}

	var allServices models.AllServices
	allServices.Services = make([]*models.Service, 0)
	for _, v := range services.Services {
		allServices.Services = append(allServices.Services, &models.Service{
			ID:          v.ID,
			ServiceName: v.ServiceName,
			About:       v.About,
			ImageUrl:    v.ImageUrl,
		})
	}
	if len(allServices.Services) == 0 {
		allServices.Services = nil
		ctx.JSON(http.StatusOK, allServices)
		return
	}

	ctx.JSON(http.StatusOK, allServices)
}

// @Summary 	Create service
// @Description This api can service register
// @Tags 		Service
// @Security    BearerAuth
// @Accept 		json
// @Produce 	json
// @Param body body models.CreateService true "Service"
// @Success 201 {object} models.Service
// @Failure 400 string Error response
// @Router /v1/create-service [post]
func (h *handlerV1) CreateService(ctx *gin.Context) {
	var req models.CreateService
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	service, err := h.Storage.UnicalPro().CreateService(ctx, &repo.Service{
		ServiceName: req.ServiceName,
		About:       req.About,
		ImageUrl:    req.ImageUrl,
	})

	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while creting service",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Service{
		ID:          service.ID,
		ServiceName: service.ServiceName,
		About:       service.About,
		ImageUrl:    service.ImageUrl,
	})
}

// @Summary 	Delete service
// @Description This api can delete service
// @Tags 		Service
// @Security    BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/delete-service/{id}  [delete]
func (h *handlerV1) DeleteService(ctx *gin.Context) {
	id := ctx.Param("id")
	doctorID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	err = h.Storage.UnicalPro().DeleteService(ctx, int64(doctorID))
	log.Println(err)
	if err != nil {
		h.log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": "doctor not found to delete",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while deleting doctor",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"is_ok": "success",
	})
}

// @Summary 	Delete service
// @Description This api can update service
// @Tags 		Service
// @Security    BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param body 	body models.CreateService true "Body"
// @Success 	200 {object} models.Service
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/update-service/{id}  [post]
func (h *handlerV1) UpdateService(ctx *gin.Context) {
	id := ctx.Param("id")
	doctorID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}
	var req models.CreateService
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	err = h.Storage.UnicalPro().UpdateService(ctx, &repo.Service{
		ID:          int64(doctorID),
		ServiceName: req.ServiceName,
		About:       req.About,
		ImageUrl:    req.ImageUrl,
	})

	if err != nil {
		h.log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": "doctor not found to delete",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while deleting doctor",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"is_ok": "success",
	})
}

// @Summary 		Get service by id
// @Description 	This api can get service by id
// @Tags 			Service
// @Accept 			json
// @Produce         json
// @Param 			id path string true "ID"
// @Success         200			{object}  models.Service
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/serviceById-get/{id} [get]
func (h *handlerV1) ServiceGet(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID parameter is empty",
		})
		return
	}

	serviceID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while parsing ID",
		})
		return
	}

	response, err := h.Storage.UnicalPro().GetServiceById(ctx, int64(serviceID))
	log.Println(err)
	if err != nil {
		h.log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": "service not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while getting service",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Service{
		ID:          response.ID,
		ServiceName: response.ServiceName,
		About:       response.About,
		ImageUrl:    response.ImageUrl,
	})
}
