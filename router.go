package main

import (
	"windwalker/handlers"
	"windwalker/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.NoRoute(handlers.NoRoute)
	router.NoMethod(handlers.NoMethod)

	v1 := router.Group("/v1")
	{
		v1.GET("/health", handlers.HealthCheck)
		v1.POST("/login", handlers.Login)
		v1.POST("/register", handlers.Register)
		v1.POST("/reset")
		secure := v1.Group("/secure", middleware.WithAuthentication("secret_key"))
		{
			secure.POST("/logout", handlers.Logout)
			secure.GET("/onboarding", handlers.GetOnboarding)
			secure.POST("/verify", handlers.VerifyEmail)

			profile := secure.Group("/profile")
			{
				profile.GET("/", handlers.GetProfile)                           // combined profile
				profile.POST("/basic", handlers.CreateBasicProfile)             // create and update basic profile
				profile.POST("/academic", handlers.CreateAcademicProfile)       // create new academic profile entry
				profile.PUT("/academic/:id", handlers.UpdateAcademicProfile)    // update existing academic profile entry
				profile.DELETE("/academic/:id", handlers.DeleteAcademicProfile) // delete existing academic profile entry
				profile.POST("/work", handlers.CreateWorkProfile)               // create new work profile entry
				profile.PUT("/work/:id", handlers.UpdateWorkProfile)            // update existing work profile entry
				profile.DELETE("/work/:id", handlers.DeleteWorkProfile)         // delete existing work profile entry
			}

			feed := secure.Group("/feed")
			{
				feed.GET("/")    // get latest feed
				feed.GET("/:id") // get feed item detatils
			}

			college := secure.Group("/college", middleware.WithAuthorization("college"))
			{
				college.GET("/profile", handlers.GetCollegeProfile)
				college.POST("/profile", handlers.UpdateCollegeProfile)
				college.GET("/students", handlers.GetStudents)           // get list of student
				college.GET("/companies", handlers.GetCompanies)         // get list of companies
				college.POST("/companies", handlers.CreateCompany)       // create new company
				college.DELETE("/companies/:id", handlers.DeleteCompany) // delete a company
				college.GET("/posts", handlers.GetPosts)                 // get list of posts
			}
		}
	}
}
