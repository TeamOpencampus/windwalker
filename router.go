package main

func (s *Server) Routes() {
	router := s.engine

	v1 := router.Group("/v1")
	{
		v1.GET("/health", s.HealthCheck)
		v1.POST("/login", s.Login)
		v1.POST("/register", s.Register)
		v1.POST("/reset")
		secure := v1.Group("/secure", WithAuthentication())
		{
			secure.POST("/logout", s.Logout)
			secure.GET("/onboarding")
			secure.POST("/verify")

			profile := secure.Group("/profile")
			{
				profile.GET("/")                // combined profile
				profile.POST("/basic")          // create basic profile
				profile.PUT("/basic")           // update basic profile
				profile.POST("/academic")       // create new academic profile entry
				profile.PUT("/academic/:id")    // update existing academic profile entry
				profile.DELETE("/academic/:id") // delete existing academic profile entry
				profile.POST("/work")           // create new work profile entry
				profile.PUT("/work/:id")        // update existing work profile entry
				profile.DELETE("/work/:id")     // delete existing work profile entry
			}
			feed := secure.Group("/feed")
			{
				feed.GET("/")    // get latest feed
				feed.GET("/:id") // get feed item detatils
			}

			college := secure.Group("/college", WithAuthorization("college"))
			{
				college.GET("/students")  // get list of student
				college.GET("/companies") // get list of companies
				college.GET("/posts")     // get list of posts
			}
		}
	}
}
