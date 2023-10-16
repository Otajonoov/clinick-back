package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/api/models"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/pkg/logger"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/storage/repo"
)

func customerParams(c *gin.Context) (*models.CustomersFindReq, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	return &models.CustomersFindReq{
		Limit: int64(limit),
		Page:  int64(page),
	}, nil
}

// @Summary 	Get customers
// @Description This api can get customers
// @Tags 		Customer
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.CustomersFindReq false "Filter"
// @Success 	200 {object} models.CustomersResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/get-customers [get]
func (h *handlerV1) GetAllCustomers(ctx *gin.Context) {
	var (
		CustomersResp models.CustomersResp
	)

	req, err := customerParams(ctx)
	if err != nil {
		h.log.Error("Error finding doctor", logger.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	response, err := h.Storage.UnicalPro().GetAllCustomers(repo.CustomersFindReq{
		Limit: req.Limit,
		Page:  req.Page,
	})
	if err != nil {
		h.log.Error("Error finding patients", logger.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, customer := range response.Customers {
		CustomersResp.Customers = append(CustomersResp.Customers, &models.Customer{
			ID:       customer.ID,
			Fullname: customer.Fullname,
			Stars:    customer.Stars,
			About:    customer.About,
			ImageUrl: customer.ImageUrl,
		})
	}
	CustomersResp.Count = response.Count

	ctx.JSON(http.StatusCreated, CustomersResp)
}

// @Summary 	create customer
// @Description This api can customer register
// @Tags 		Customer
// @Security    BearerAuth
// @Accept 		json
// @Produce 	json
// @Param body 	body models.CreateCustomer true "Body"
// @Success 201 {object} models.Customer
// @Failure 400 string Error response
// @Router /v1/create-customer [post]
func (h *handlerV1) CreateCustomer(ctx *gin.Context) {
	var req models.CreateCustomer
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	customer, err := h.Storage.UnicalPro().CreateCustomer(ctx, &repo.Customer{
		Fullname: req.Fullname,
		Stars:    req.Stars,
		About:    req.About,
		ImageUrl: req.ImageUrl,
	})

	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while creting customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Customer{
		ID:        customer.ID,
		Fullname:  customer.Fullname,
		Stars:     customer.Stars,
		About:     customer.About,
		ImageUrl:  customer.ImageUrl,
		CreatedAt: customer.CreatedAt,
	})
}

// @Summary 	Delete customer
// @Description This api can customer doctor
// @Tags 		Customer
// @Security    BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/delete-customer/{id}  [delete]
func (h *handlerV1) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	doctorID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	err = h.Storage.UnicalPro().DeleteCustomer(ctx, int64(doctorID))
	log.Println(err)
	if err != nil {
		h.log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": "customer not found to delete",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while deleting customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"is_ok": "success",
	})
}

// @Summary 	Delete customer
// @Description This api can update customer
// @Tags 		Customer
// @Security    BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param body 	body models.CreateCustomer true "Body"
// @Success 	200 {object} models.Customer
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/update-customer/{id}  [post]
func (h *handlerV1) UpdateCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	doctorID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}
	var req models.CreateCustomer
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	err = h.Storage.UnicalPro().UpdateCustomer(ctx, &repo.Customer{
		ID:       int64(doctorID),
		Fullname: req.Fullname,
		Stars:    req.Stars,
		About:    req.About,
		ImageUrl: req.ImageUrl,
	})

	if err != nil {
		h.log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": "customer not found to update",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while update customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"is_ok": "success",
	})
}

// @Summary 		Get Customer by id
// @Description 	This api can get Customer by id
// @Tags 			Customer
// @Accept 			json
// @Produce         json
// @Param 			id path string true "ID"
// @Success         200			{object}  models.Customer
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/customerById-get/{id} [get]
func (h *handlerV1) CustomerGet(ctx *gin.Context) {
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

	response, err := h.Storage.UnicalPro().GetCustomerById(ctx, int64(serviceID))
	log.Println(err)
	if err != nil {
		h.log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": "customer not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while getting customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Customer{
		ID:       response.ID,
		Fullname: response.Fullname,
		Stars:    response.Stars,
		About:    response.About,
		ImageUrl: response.ImageUrl,
	})
}
