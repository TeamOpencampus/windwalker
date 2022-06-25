package controllers

import (
	"context"
	"errors"
	"windwalker/ent"
	"windwalker/ent/studentacademicprofile"
	"windwalker/ent/studentprofile"
	"windwalker/ent/user"
	"windwalker/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func (c *Controller) GetMe(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)

	profile, err := c.Client.User.
		Query().
		Where(user.ID(id)).
		WithCollegeProfile().
		WithStudentProfile(func(q *ent.StudentProfileQuery) {
			q.Limit(1)
			q.WithAcademicProfile()
			q.WithWorkProfile()
		}).
		WithEnrolledIn(func(q *ent.UserQuery) {
			q.WithCollegeProfile()
		}).
		Only(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("profile/not-found"))
		return
	}

	NewSuccessResponse(ctx, profile)
}

func (c *Controller) CreateBasicProfile(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)

	var data models.BasicProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("profile/invalid-input"))
		return
	}

	_, err := c.Client.StudentProfile.
		Create().
		SetName(data.Name).
		SetPhone(data.Phone).
		SetGender(data.Gender).
		SetCaste(data.Caste).
		SetNationality(data.Nationality).
		SetUserID(id).
		Save(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("profile/save-failed"))
		return
	}

	RespondEmpty(ctx)
}

func (c *Controller) CreateAcademicProfile(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)

	var data models.AcademicProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("profile/invalid-input"))
		return
	}

	acm, err := c.Client.StudentAcademicProfile.
		Create().
		SetCourse(data.Course).
		SetInstitute(data.Institute).
		SetBoard(data.Board).
		SetRegNo(data.RegNo).
		SetDepartment(data.Department).
		SetStartDate(data.StartDate).
		SetEndDate(data.EndDate).
		SetMarks(data.Marks).
		Save(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("profile/save-failed"))
		return
	}

	profile, err := c.Client.User.
		Query().
		Where(user.ID(id)).
		QueryStudentProfile().
		Only(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("profile/not-found"))
		return
	}

	if _, err = profile.
		Update().
		AddAcademicProfile(acm).
		Save(context.Background()); err != nil {
		_ = ctx.Error(errors.New("profile/add-failed"))
		return
	}

	RespondEmpty(ctx)
}

func (c *Controller) UpdateAcademicProfile(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	pid := ctx.Param("id")

	profile, err := c.Client.StudentProfile.
		Query().
		Where(studentprofile.HasUserWith(user.ID(id))).
		QueryAcademicProfile().
		Where(studentacademicprofile.ID(pid)).
		Only(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("profile/not-found"))
		return
	}

	var data models.AcademicProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("profile/invalid-input"))
		return
	}

	if _, err = profile.
		Update().
		SetCourse(data.Course).
		SetInstitute(data.Institute).
		SetBoard(data.Board).
		SetRegNo(data.RegNo).
		SetDepartment(data.Department).
		SetStartDate(data.StartDate).
		SetEndDate(data.EndDate).
		SetMarks(data.Marks).
		Save(context.Background()); err != nil {
		_ = ctx.Error(errors.New("profile/update-failed"))
		return
	}

	RespondEmpty(ctx)
}

func (c *Controller) DeleteAcademicProfile(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	pid := ctx.Param("id")

	profile, err := c.Client.
		StudentProfile.
		Query().
		Where(studentprofile.HasUserWith(user.ID(id))).
		QueryAcademicProfile().
		Where(studentacademicprofile.ID(pid)).
		Only(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("profile/not-found"))
		return
	}

	if err := c.Client.
		StudentAcademicProfile.
		DeleteOne(profile).
		Exec(context.Background()); err != nil {
		_ = ctx.Error(errors.New("profile/delete-failed"))
		return
	}

	RespondEmpty(ctx)
}

func (c *Controller) CreateWorkProfile(ctx *gin.Context) {
	// db := GetDatabase(ctx)
	// id, _ := strconv.ParseUint(ctx.MustGet("ID").(xid.ID), 10, 64)

	// var data models.WorkProfile
	// if err := ctx.ShouldBindJSON(&data); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"status": "error",
	// 		"data":   "Malformed data provided",
	// 	})
	// 	return
	// }
	// var profile models.Profile
	// if err := db.Where("user_id = ?", id).First(&profile).Error; err != nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"status": "error",
	// 		"data":   "Failed to retrieve profile: " + err.Error(),
	// 	})
	// 	return
	// }

	// if err := db.Model(&profile).Association("WorkProfile").Append(&data); err != nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"status": "error",
	// 		"data":   "Failed to add work details: " + err.Error(),
	// 	})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"status": "ok",
	// 	"data":   "Successfully added work details.",
	// })

}

func (c *Controller) UpdateWorkProfile(ctx *gin.Context) {
	// var result struct{ Count uint }
	// db := GetDatabase(ctx)
	// id := ctx.MustGet("ID").(xid.ID)
	// pid := ctx.Param("id")

	// if err := db.Raw(
	// 	`select count(*) as count from work_profiles as wp
	// 	join profiles p on p.id = wp.profile_id
	// 	join users u on u.id = p.user_id where u.id =  ? and wp.id = ?`,
	// 	id, pid,
	// ).
	// 	Scan(&result).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"status": "error",
	// 		"data":   "Failed to handle transaction.",
	// 	})
	// 	return
	// }

	// if result.Count == 0 {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"status": "error",
	// 		"data":   "The resource does not exist or do to belong to you.",
	// 	})
	// 	return
	// }

	// var data models.AcademicProfile
	// if err := ctx.ShouldBindJSON(&data); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"status": "error",
	// 		"data":   "Malformed data found",
	// 	})
	// 	return
	// }

	// var profile models.WorkProfile
	// db.Where("id = ?", pid).First(&profile)
	// if err := db.Model(&profile).Select(
	// 	"Company",
	// 	"StartDate",
	// 	"EndDate",
	// 	"Salary",
	// 	"Position",
	// ).Updates(data).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"status": "error",
	// 		"data":   "Failed to update record",
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"status": "ok",
	// 	"data":   "Record updated successfully",
	// })
}

func (c *Controller) DeleteWorkProfile(ctx *gin.Context) {
	// var result struct{ Count uint }
	// db := GetDatabase(ctx)
	// id := ctx.MustGet("ID").(xid.ID)
	// pid := ctx.Param("id")

	// if err := db.Raw(
	// 	`select count(*) as count from work_profiles as wp
	// 	join profiles p on p.id = wp.profile_id
	// 	join users u on u.id = p.user_id where u.id =  ? and wp.id = ?`,
	// 	id, pid,
	// ).
	// 	Scan(&result).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"status": "error",
	// 		"data":   "Failed to handle transaction.",
	// 	})
	// 	return
	// }

	// if result.Count == 0 {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"status": "error",
	// 		"data":   "The resource does not exist or do to belong to you.",
	// 	})
	// 	return
	// }

	// if err := db.Delete(&models.WorkProfile{}, pid).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"status": "error",
	// 		"data":   "Failed to delete the record.",
	// 	})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"status": "ok",
	// 	"data":   "Record deleted successfully",
	// })
}
