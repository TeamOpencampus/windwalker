package controllers

import (
	"context"
	"errors"
	"windwalker/ent"
	"windwalker/ent/company"
	"windwalker/ent/jobpost"
	"windwalker/ent/user"
	"windwalker/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func (c *Controller) CreateCollegeProfile(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)

	var data models.CollegeProfile
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("profile/invalid-input"))
		return
	}

	if _, err := c.Client.CollegeProfile.
		Create().
		SetName(data.Name).
		SetPhone(data.Phone).
		SetAddress(data.Address).
		SetType(data.Type).
		SetUserID(id).
		Save(context.Background()); err != nil {
		_ = ctx.Error(errors.New("profile/save-failed"))
		return
	}

	RespondEmpty(ctx)
}

func (c *Controller) UpdateCollegeProfile(ctx *gin.Context) {
	// id := ctx.MustGet("ID").(string)
	// profile, err := c.Client.
	// 	CollegeProfile.
	// 	Query().
	// 	Where(collegeprofile.HasCollegeWith(user.ID(id))).
	// 	Only(context.Background())

	// if err != nil {
	// 	_ = ctx.Error(errors.New("profile/not-found"))
	// 	return
	// }

	// var data models.CollegeProfile
	// if err := ctx.ShouldBindJSON(&data); err != nil {
	// 	_ = ctx.Error(errors.New("profile/invalid-input"))
	// 	return
	// }

	// if _, err = profile.
	// 	Update().
	// 	SetName(data.Name).
	// 	SetPhone(data.Phone).
	// 	SetAddress(data.Address).
	// 	SetType(data.Type).
	// 	Save(context.Background()); err != nil {
	// 	_ = ctx.Error(errors.New("profile/update-failed"))
	// 	return
	// }

	// NewSuccessResponse(ctx, "record updated")
}

func (c *Controller) GetStudents(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	candidates, err := c.Client.
		User.
		Query().
		Where(user.ID(id), user.HasStudentProfile()).
		QueryCandidates().
		Where(user.Role("student")).
		WithStudentProfile(func(q *ent.StudentProfileQuery) {
			q.WithAcademicProfile()
			q.WithWorkProfile()
		}).
		All(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("profile/err-failed"))
		return
	}

	NewSuccessResponse(ctx, candidates)
}

func (c *Controller) GetCompanies(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)

	companies, err := c.Client.
		User.
		Query().
		Where(user.ID(id)).
		QueryCompanies().
		All(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("companies-query-failed"))
		return
	}

	NewSuccessResponse(ctx, companies)
}

func (c *Controller) CreateCompany(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)

	var data models.Company
	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("companies/invalid-input"))
		return
	}

	if _, err := c.Client.
		Company.
		Create().
		SetCompanyName(data.CompanyName).
		SetContactPersonName(data.ContactPersonName).
		SetContactPersonPhone(data.ContactPersonPhone).
		SetContactPersonEmail(data.ContactPersonEmail).
		SetUserID(id).
		Save(context.Background()); err != nil {
		_ = ctx.Error(errors.New("profile/save-failed"))
		return
	}

	RespondEmpty(ctx)
}

func (c *Controller) DeleteCompany(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	cid := ctx.Param("id")

	comID, err := xid.FromString(cid)
	if err != nil {
		_ = ctx.Error(errors.New("param-parse-fail"))
		return
	}

	company, err := c.Client.
		User.
		Query().
		Where(user.ID(id)).
		QueryCompanies().
		Where(company.ID(comID)).
		Only(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("profile/not-found"))
		return
	}

	if err = c.Client.
		Company.
		DeleteOne(company).
		Exec(context.Background()); err != nil {
		_ = ctx.Error(errors.New("profile/delete-failed"))
		return
	}
	RespondEmpty(ctx)
}

func (c *Controller) GetPosts(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	posts, err := c.Client.User.
		Query().
		Where(user.ID(id)).
		QueryJobPosts().
		WithCompany().
		All(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("posts-query-fail"))
		return
	}
	NewSuccessResponse(ctx, posts)
}

func (c *Controller) GetPostByID(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	jid := ctx.Param("id")
	jobID, err := xid.FromString(jid)
	if err != nil {
		_ = ctx.Error(errors.New("param-parse-failed"))
		return
	}

	post, err := c.Client.User.
		Query().
		Where(user.ID(id)).
		QueryJobPosts().
		Where(jobpost.ID(jobID)).
		WithCompany().
		WithCandidates(func(q *ent.UserQuery) {
			q.WithStudentProfile(func(q *ent.StudentProfileQuery) {
				q.WithAcademicProfile()
				q.WithWorkProfile()
			})
		}).
		Only(context.Background())

	if err != nil {
		_ = ctx.Error(errors.New("not-found"))
		return
	}

	NewSuccessResponse(ctx, post)
}

func (c *Controller) CreatePost(ctx *gin.Context) {
	var data struct {
		CompanyID   string   `json:"company_id" binding:"required"`
		Position    string   `json:"position" binding:"required"`
		Location    string   `json:"location" binding:"required"`
		Salary      string   `json:"salary" binding:"required"`
		Description string   `json:"description" binding:"required"`
		Tags        []string `json:"tags" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		_ = ctx.Error(errors.New("invalid-input"))
		return
	}

	id := ctx.MustGet("ID").(xid.ID)

	comID, err := xid.FromString(data.CompanyID)
	if err != nil {
		_ = ctx.Error(errors.New("param-parse-fail"))
		return
	}

	if _, err := c.Client.
		JobPost.
		Create().
		SetCompanyID(comID).
		SetUserID(id).
		SetPosition(data.Position).
		SetLocation(data.Location).
		SetSalary(data.Salary).
		SetDescription(data.Description).
		SetTags(data.Tags).
		Save(context.Background()); err != nil {
		_ = ctx.Error(errors.New("save-failed"))
		return
	}

	RespondEmpty(ctx)

}
