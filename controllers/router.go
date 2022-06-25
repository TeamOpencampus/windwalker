package controllers

import (
	"windwalker/middleware"

	"github.com/gin-gonic/gin"
)

func (c *Controller) SetupRoutes(router *gin.Engine) {
	router.NoRoute(NoRoute)
	router.NoMethod(NoMethod)
	router.GET("/health", HealthCheck)

	v1 := router.Group("/v1")
	{
		v1.POST("/login", c.Login)
		v1.POST("/register", c.Register)
		// v1.POST("/reset")
		secure := v1.Group("/secure", middleware.WithAuthentication("secret_key"))
		{
			secure.GET("/me", c.GetMe)
			secure.GET("/colleges", c.GetColleges)
			secure.POST("/colleges/join/:id", c.JoinCollege)
			secure.GET("/jobs-feed", c.GetJobsFeed)
			secure.GET("/jobs-feed/:id", c.GetJobByID)
			secure.POST("/jobs-feed/:id/apply", c.ApplyForJob)
			secure.POST("/jobs-feed/:id/retract", c.RetractJobApplication)
			// secure.POST("/verify", c.VerifyEmail)

			profile := secure.Group("/profile")
			{
				profile.POST("/basic", c.CreateBasicProfile)             // create and update basic profile
				profile.POST("/academic", c.CreateAcademicProfile)       // create new academic profile entry
				profile.PUT("/academic/:id", c.UpdateAcademicProfile)    // update existing academic profile entry
				profile.DELETE("/academic/:id", c.DeleteAcademicProfile) // delete existing academic profile entry
				profile.POST("/work", c.CreateWorkProfile)               // create new work profile entry
				profile.PUT("/work/:id", c.UpdateWorkProfile)            // update existing work profile entry
				profile.DELETE("/work/:id", c.DeleteWorkProfile)         // delete existing work profile entry
			}

			college := secure.Group("/college", middleware.WithAuthorization("college"))
			{
				college.POST("/profile", c.CreateCollegeProfile)
				college.GET("/students", c.GetStudents)           // get list of student
				college.GET("/companies", c.GetCompanies)         // get list of companies
				college.POST("/companies", c.CreateCompany)       // create new company
				college.DELETE("/companies/:id", c.DeleteCompany) // delete a company
				college.GET("/posts", c.GetPosts)                 // get list of posts
				college.GET("/posts/:id", c.GetPostByID)          // get list of posts
				college.POST("/posts", c.CreatePost)              // get list of posts
			}
		}
	}
}
