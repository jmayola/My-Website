package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

func SessionStore() *session.Store {
	storage := redis.New()
	sessionconf := session.Config{
		Storage: storage,
	}
	return session.New(sessionconf)
}
