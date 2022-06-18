package main

import (
	"errors"
	"net/http"
	"strconv"
	"windwalker/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetProfile(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	var profile models.Profile
	if err := db.
		Preload("AcademicProfile").
		Preload("WorkProfile").
		Where("user_id = ?", id).
		First(&profile).
		Error; err != nil {
		_ = ctx.Error(errors.New("profile/not-found"))
		return
	}

	NewSuccessResponse(ctx, profile)
}

func (s *Server) CreateBasicProfile(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id, _ := strconv.ParseUint(ctx.MustGet("ID").(string), 10, 64)

	var data models.BasicProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("profile/invalid-input"))
		return
	}

	var profile models.Profile
	if err := db.
		Where(models.Profile{UserID: uint(id)}).
		Attrs(models.Profile{BasicProfile: data}).
		FirstOrCreate(&profile).
		Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}
	profile.BasicProfile = data
	if err := db.Save(profile).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	NewSuccessResponse(ctx, "basic profile updated")
}

func (s *Server) CreateAcademicProfile(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id, _ := strconv.ParseUint(ctx.MustGet("ID").(string), 10, 64)

	var data models.AcademicProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("profile/invalid-input"))
		return
	}
	var profile models.Profile
	if err := db.Where("user_id = ?", id).First(&profile).Error; err != nil {
		_ = ctx.Error(errors.New("profile/not-found"))
		return
	}

	if err := db.
		Model(&profile).
		Association("AcademicProfile").
		Append(&data); err != nil {
		_ = ctx.Error(errors.New("profile/internal-server-error"))
		return
	}

	NewSuccessResponse(ctx, "record created")
}

func (s *Server) UpdateAcademicProfile(ctx *gin.Context) {
	var result struct{ Count uint }
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	pid := ctx.Param("id")

	if err := db.Raw(
		`select count(*) as count from academic_profiles as ap
		join profiles p on p.id = ap.profile_id
		join users u on u.id = p.user_id where u.id =  ? and ap.id = ?`,
		id, pid,
	).
		Scan(&result).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	if result.Count == 0 {
		_ = ctx.Error(errors.New("profile/resource-nonexistent-or-not-owner"))
		return
	}

	var data models.AcademicProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("profile/invalid-input"))
		return
	}

	var profile models.AcademicProfile
	db.Where("id = ?", pid).First(&profile)
	if err := db.Model(&profile).Select(
		"Course",
		"Institute",
		"Board",
		"RegNo",
		"Department",
		"StartDate",
		"EndDate",
		"Marks",
	).Updates(data).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	NewSuccessResponse(ctx, "record updated")
}

func (s *Server) DeleteAcademicProfile(ctx *gin.Context) {
	var result struct{ Count uint }
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	pid := ctx.Param("id")

	if err := db.Raw(
		`select count(*) as count from academic_profiles as ap
		join profiles p on p.id = ap.profile_id
		join users u on u.id = p.user_id where u.id =  ? and ap.id = ?`,
		id, pid,
	).
		Scan(&result).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	if result.Count == 0 {
		_ = ctx.Error(errors.New("profile/resource-nonexistent-or-not-owner"))
		return
	}

	if err := db.Delete(&models.AcademicProfile{}, pid).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}
	NewSuccessResponse(ctx, "record deleted")
}

func (s *Server) CreateWorkProfile(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id, _ := strconv.ParseUint(ctx.MustGet("ID").(string), 10, 64)

	var data models.WorkProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   "Malformed data provided",
		})
		return
	}
	var profile models.Profile
	if err := db.Where("user_id = ?", id).First(&profile).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"data":   "Failed to retrieve profile: " + err.Error(),
		})
		return
	}

	if err := db.Model(&profile).Association("WorkProfile").Append(&data); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"data":   "Failed to add work details: " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "Successfully added work details.",
	})

}

func (s *Server) UpdateWorkProfile(ctx *gin.Context) {
	var result struct{ Count uint }
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	pid := ctx.Param("id")

	if err := db.Raw(
		`select count(*) as count from work_profiles as wp
		join profiles p on p.id = wp.profile_id
		join users u on u.id = p.user_id where u.id =  ? and wp.id = ?`,
		id, pid,
	).
		Scan(&result).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   "Failed to handle transaction.",
		})
		return
	}

	if result.Count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"data":   "The resource does not exist or do to belong to you.",
		})
		return
	}

	var data models.AcademicProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   "Malformed data found",
		})
		return
	}

	var profile models.WorkProfile
	db.Where("id = ?", pid).First(&profile)
	if err := db.Model(&profile).Select(
		"Company",
		"StartDate",
		"EndDate",
		"Salary",
		"Position",
	).Updates(data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   "Failed to update record",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "Record updated successfully",
	})
}

func (s *Server) DeleteWorkProfile(ctx *gin.Context) {
	var result struct{ Count uint }
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	pid := ctx.Param("id")

	if err := db.Raw(
		`select count(*) as count from work_profiles as wp
		join profiles p on p.id = wp.profile_id
		join users u on u.id = p.user_id where u.id =  ? and wp.id = ?`,
		id, pid,
	).
		Scan(&result).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   "Failed to handle transaction.",
		})
		return
	}

	if result.Count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"data":   "The resource does not exist or do to belong to you.",
		})
		return
	}

	if err := db.Delete(&models.WorkProfile{}, pid).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   "Failed to delete the record.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "Record deleted successfully",
	})
}
