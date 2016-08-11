package main

import (
//  "time"
//  . "response"
)

func Routes() {
	v1 := router.Group("/v1")
	{
	  v1book := v1.Group("/book")
	  {
	    v1book.GET("/hotel", undefineFunc)
	  }
		v1feed := v1.Group("/feed")
		{
		  v1feed.GET("/", GetFeedv1)
//			v1feed.GET("/:name", GetFeedWithIdv1)
//  		v1feed.GET("/category", undefineFunc)
		}
		v1profile := v1.Group("/profile")
		{
		  v1profile.GET("/brief/:uid", undefineFunc)
		  v1profile.GET("/", undefineFunc)
		  v1profile.PUT("/", undefineFunc)
		  v1profile.PUT("/privacy", undefineFunc)
		}
		v1sso := v1.Group("/sso")
		{
		  v1sso.GET("/login", undefineFunc)
		  v1sso.GET("/logout", undefineFunc)
		  v1sso.GET("/logout-all", undefineFunc)
		  v1sso.GET("/refreshToken", undefineFunc)
		}
		v1lbs := v1.Group("/lbs")
		{
		  v1lbs.GET("/addrtime", undefineFunc)
		  v1lbs.GET("/exchange", undefineFunc)
		  v1lbs.GET("/latlng", undefineFunc)
		  v1lbs.GET("/weather/d", undefineFunc)
		  v1lbs.GET("/weather/3h", undefineFunc)
		  v1lbs.GET("/nearby", undefineFunc)
		  v1lbs.GET("/nearby/category", undefineFunc)
		}
		v1review := v1.Group("/review")
		{
		  v1review.GET("/", undefineFunc)
		  v1review.POST("/", undefineFunc)
		  v1review.PUT("/", undefineFunc)
		  v1review.DELETE("/", undefineFunc)
		}
		v1locsug := v1.Group("/locsug")
		{
		  v1locsug.GET("/bybridsearch", undefineFunc)
		  v1locsug.GET("/recommend", undefineFunc)
		  v1locsug.GET("/g2lyrQuery", undefineFunc)
		}
		v1other := v1.Group("/ot")
		{
		  v1other.GET("/version", undefineFunc)
		  v1other.POST("/tracking", undefineFunc)
		}
	}

	v2 := router.Group("/v2")
	{
//	  v2.GET("/", Response{Status:"Working", Timestamp:time.Now().Unix()})
		v2.GET("/", GetHelloV2)
		v2.GET("/get", Getv2)
		v2.GET("/loop", GetLoopv2)
	}

}
