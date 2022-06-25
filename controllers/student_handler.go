package controllers

import (
	"context"
	"errors"
	"windwalker/ent"
	"windwalker/ent/jobpost"
	"windwalker/ent/user"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func (c *Controller) GetColleges(ctx *gin.Context) {
	colleges, err := c.Client.
		User.
		Query().
		Where(
			user.Role("college"),
			user.HasCollegeProfile(),
		).
		WithCollegeProfile().
		All(context.Background())
	if err != nil {
		_ = ctx.Error(errors.New("college-query-failed"))
		return
	}

	NewSuccessResponse(ctx, colleges)
}

func (c *Controller) JoinCollege(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	cid := ctx.Param("id")
	colID, err := xid.FromString(cid)
	if err != nil {
		_ = ctx.Error(errors.New("param-parse-failed"))
		return
	}
	if _, err := c.Client.
		User.
		UpdateOneID(id).
		AddEnrolledInIDs(colID).
		Save(context.Background()); err != nil {
		_ = ctx.Error(errors.New("save-failed"))
		return
	}
}

func (c *Controller) GetJobsFeed(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	posts, err := c.Client.
		User.
		Query().
		Where(user.ID(id)).
		QueryEnrolledIn().
		QueryJobPosts().
		WithCompany().
		WithCandidates(func(q *ent.UserQuery) {
			q.Where(user.ID(id))
		}).
		All(context.Background())
	if err != nil {
		_ = ctx.Error(errors.New("posts-query-failed"))
		return
	}

	NewSuccessResponse(ctx, posts)
}

func (c *Controller) GetJobByID(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	jid := ctx.Param("id")
	jobID, err := xid.FromString(jid)
	if err != nil {
		_ = ctx.Error(errors.New("param-parse-failed"))
		return
	}

	post, err := c.Client.
		User.
		Query().
		Where(user.ID(id)).
		QueryEnrolledIn().
		QueryJobPosts().
		Where(jobpost.ID(jobID)).
		WithCompany().
		WithCandidates(func(q *ent.UserQuery) {
			q.Where(user.ID(id))
		}).
		Only(context.Background())
	if err != nil {
		_ = ctx.Error(errors.New("posts-query-failed"))
		return
	}

	NewSuccessResponse(ctx, post)
}

func (c *Controller) ApplyForJob(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	jid := ctx.Param("id")
	jobID, err := xid.FromString(jid)
	if err != nil {
		_ = ctx.Error(errors.New("param-parse-failed"))
		return
	}

	if _, err := c.Client.
		JobPost.
		UpdateOneID(jobID).
		AddCandidateIDs(id).
		Save(context.Background()); err != nil {
		_ = ctx.Error(errors.New("save-failed"))
		return
	}

	RespondEmpty(ctx)
}

func (c *Controller) RetractJobApplication(ctx *gin.Context) {
	id := ctx.MustGet("ID").(xid.ID)
	jid := ctx.Param("id")
	jobID, err := xid.FromString(jid)
	if err != nil {
		_ = ctx.Error(errors.New("param-parse-failed"))
		return
	}

	if _, err := c.Client.
		JobPost.
		UpdateOneID(jobID).
		RemoveCandidateIDs(id).
		Save(context.Background()); err != nil {
		_ = ctx.Error(errors.New("save-failed"))
		return
	}

	RespondEmpty(ctx)
}
