package main

import (
	"errors"
	"strconv"
	"windwalker/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetCollegeProfile(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	var profile models.CollegeProfile
	if err := db.Where("user_id = ?", id).First(&profile).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}
	NewSuccessResponse(ctx, profile)

}
func (s *Server) UpdateCollegeProfile(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id, _ := strconv.ParseUint(ctx.MustGet("ID").(string), 10, 64)

	var data models.CollegeProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("profile/invalid-input"))
		return
	}

	var profile models.CollegeProfile
	if err := db.
		Where(models.CollegeProfile{UserID: uint(id)}).
		Attrs(data).
		FirstOrCreate(&profile).
		Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	profile.Name = data.Name
	profile.Address = data.Address
	profile.Phone = data.Phone
	profile.Type = data.Type
	if err := db.Save(profile).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}

	NewSuccessResponse(ctx, "profile updated")
}

func (s *Server) GetStudents(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}
	var users []models.User
	if err := db.Model(&user).Association("AssociatedUsers").Find(&users); err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}
}

func (s *Server) GetCompanies(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	var companies []struct {
		ID                 int    `json:"id"`
		CompanyName        string `json:"company_name"`
		ContactPersonName  string `json:"contact_person_name"`
		ContactPersonEmail string `json:"contact_person_email"`
		ContactPersonPhone string `json:"contact_person_phone"`
		JobCount           int    `json:"job_count"`
		StudentCount       int    `json:"student_count"`
	}
	if err := db.Raw(`
		select id,
			   company_name,
			   contact_person_name,
			   contact_person_email,
			   contact_person_phone,
			   (select count(*) from users where company_id = companies.id) as job_count,
			   (select count(*) from posts where company_id = companies.id) as student_count
		from companies
		where deleted_at is null and user_id = ?
`, id).Scan(&companies).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}
	NewSuccessResponse(ctx, companies)
}

func (s *Server) CreateCompany(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id, _ := strconv.ParseUint(ctx.MustGet("ID").(string), 10, 64)
	var data models.Company
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("companies/malformed-input"))
		return
	}

	data.UserID = uint(id)
	if err := db.Create(&data).Error; err != nil {
		_ = ctx.Error(errors.New("misc/internal-server-error"))
		return
	}
	NewSuccessResponse(ctx, "record created")
}

func (s *Server) DeleteCompany(ctx *gin.Context) {
	db := GetDatabase(ctx)
	id := ctx.MustGet("ID").(string)
	cid := ctx.Param("id")

	if err := db.Where("user_id = ?", id).Delete(&models.Company{}, cid).Error; err != nil {
		_ = ctx.Error(errors.New("companies/not-found"))
		return
	}

	NewSuccessResponse(ctx, "record deleted")
}

func (s *Server) GetPosts(ctx *gin.Context) {
	// db := GetDatabase(ctx)
	// db.
}
