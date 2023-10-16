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

// @Summary 	Get doctors
// @Description This api can get doctors
// @Tags 		Doctor
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} models.AllDoctors
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/get-doctors [get]
func (h *handlerV1) HandleGetDoctors(ctx *gin.Context) {
	doctors, err := h.Storage.UnicalPro().GetAllDoctors(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get all doctors",
		})
		return
	}

	var allDoctors models.AllDoctors
	allDoctors.Doctors = make([]*models.Doctor, 0)
	for _, v := range doctors.Doctors {
		allDoctors.Doctors = append(allDoctors.Doctors, &models.Doctor{
			ID:       v.ID,
			Fullname: v.Fullname,
			Type:     v.Type,
			About:    v.About,
			ImageUrl: v.ImageUrl,
		})
	}
	if len(allDoctors.Doctors) == 0 {
		allDoctors.Doctors = nil
		ctx.JSON(http.StatusOK, allDoctors)
		return
	}

	ctx.JSON(http.StatusOK, allDoctors)
}

// @Summary create doctor
// @Description This api can doctor register
// @Tags Doctor
// @Security    BearerAuth
// @Accept json
// @Produce json
// @Param body body models.CreateDoctor true "Doctor"
// @Success 201 {object} models.Doctor
// @Failure 400 string Error response
// @Router /v1/create-doctor [post]
func (h *handlerV1) HandleCreateDoctor(ctx *gin.Context) {
	var req models.CreateDoctor
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	doctor, err := h.Storage.UnicalPro().CreateDoctor(ctx, &repo.Doctor{
		Fullname: req.Fullname,
		Type:     req.Type,
		About:    req.About,
		ImageUrl: req.ImageUrl,
	})

	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while creting doctor",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Doctor{
		ID:       doctor.ID,
		Fullname: doctor.Fullname,
		Type:     doctor.Type,
		About:    doctor.About,
		ImageUrl: doctor.ImageUrl,
	})
}

// @Summary 	Delete doctor
// @Description This api can delete doctor
// @Tags 		Doctor
// @Security    BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/delete-doctor/{id}  [delete]
func (h *handlerV1) HandleDeleteDoctor(ctx *gin.Context) {
	id := ctx.Param("id")
	doctorID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	err = h.Storage.UnicalPro().DeleteDoctor(ctx, int64(doctorID))
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

// @Summary 	Delete doctor
// @Description This api can update doctor
// @Tags 		Doctor
// @Security    BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param body 	body models.CreateDoctor true "Body"
// @Success 	200 {object} models.Doctor
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/update-doctor/{id}  [post]
func (h *handlerV1) HandleUpdateDoctor(ctx *gin.Context) {
	id := ctx.Param("id")
	doctorID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}
	var req models.CreateDoctor
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	err = h.Storage.UnicalPro().UpdateDoctor(ctx, &repo.Doctor{
		ID:       int64(doctorID),
		Fullname: req.Fullname,
		Type:     req.Type,
		About:    req.About,
		ImageUrl: req.ImageUrl,
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

// @Summary 		Get doctor by id
// @Description 	This api can get doctor by id
// @Tags 			Doctor
// @Accept 			json
// @Produce         json
// @Param 			id path string true "ID"
// @Success         200			{object}  models.Doctor
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/doctorById-get/{id} [get]
func (h *handlerV1) DoctorGet(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID parameter is empty",
		})
		return
	}

	doctorID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while parsing ID",
		})
		return
	}

	response, err := h.Storage.UnicalPro().GetDoctorById(ctx, int64(doctorID))
	log.Println(err)
	if err != nil {
		h.log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": "doctor not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while getting doctor",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Doctor{
		ID:       response.ID,
		Fullname: response.Fullname,
		Type:     response.Type,
		About:    response.About,
		ImageUrl: response.ImageUrl,
	})
}
