package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func test01() gin.HandlerFunc {
	log.Debug().Msg("init test01 middleware")
	return func(c *gin.Context) {
		c.Set("test", "debug")
		c.Next()
	}
}

func test02() gin.HandlerFunc {
	log.Debug().Msg("init test02 middleware")
	return func(c *gin.Context) {
		test, ok := c.Get("test")
		log.Debug().Msgf("%#v", test)
		log.Debug().Msgf("%#v", ok)
		if test, ok := c.Get("test"); ok {
			log.Debug().Str("test", test.(string)).Msg("test value")
		}

		c.Next()
	}
}
