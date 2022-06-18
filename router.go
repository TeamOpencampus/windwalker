package main

import "windwalker/middleware"

func (s *Server) Routes() {
	router := s.engine

	router.NoRoute(s.NoRoute)
	router.NoMethod(s.NoMethod)

	v1 := router.Group("/v1")
	{
		v1.GET("/health", s.HealthCheck)
		v1.POST("/login", s.Login)
		v1.POST("/register", s.Register)
		v1.POST("/reset")
		secure := v1.Group("/secure", middleware.WithAuthentication("secret_key"))
		{
			secure.POST("/logout", s.Logout)
			secure.GET("/onboarding", s.GetOnboarding)
			secure.POST("/verify", s.VerifyEmail)

			profile := secure.Group("/profile")
			{
				profile.GET("/", s.GetProfile)                           // combined profile
				profile.POST("/basic", s.CreateBasicProfile)             // create and update basic profile
				profile.POST("/academic", s.CreateAcademicProfile)       // create new academic profile entry
				profile.PUT("/academic/:id", s.UpdateAcademicProfile)    // update existing academic profile entry
				profile.DELETE("/academic/:id", s.DeleteAcademicProfile) // delete existing academic profile entry
				profile.POST("/work", s.CreateWorkProfile)               // create new work profile entry
				profile.PUT("/work/:id", s.UpdateWorkProfile)            // update existing work profile entry
				profile.DELETE("/work/:id", s.DeleteWorkProfile)         // delete existing work profile entry
			}

			feed := secure.Group("/feed")
			{
				feed.GET("/")    // get latest feed
				feed.GET("/:id") // get feed item detatils
			}

			college := secure.Group("/college", middleware.WithAuthorization("college"))
			{
				college.GET("/students", s.GetStudents)           // get list of student
				college.GET("/companies", s.GetCompanies)         // get list of companies
				college.POST("/companies", s.CreateCompany)       // create new company
				college.DELETE("/companies/:id", s.DeleteCompany) // delete a company
				college.GET("/posts", s.GetPosts)                 // get list of posts

				// profile := college.Group("/profile")
				{

				}
			}
		}
	}
}
