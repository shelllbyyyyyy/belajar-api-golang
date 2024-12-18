package common

import (
	"api/first-go/configs"
	"api/first-go/util"
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		
		traceID := uuid.New().String()
		c.Set("X-Trace-ID", traceID)

		
		startTime := time.Now()
		data := logrus.Fields{
			"trace_id": traceID,
			"method":   c.Method(),
			"path":     c.OriginalURL(),
		}

		ctx := context.WithValue(c.UserContext(), "data: ", data)
		c.SetUserContext(ctx)

		log.WithFields(data).Info("Incoming request")

		err := c.Next()

		duration := time.Since(startTime).Milliseconds()
		data["response_time"] = duration
		data["status"] = c.Response().StatusCode()

		if c.Response().StatusCode() >= 200 && c.Response().StatusCode() <= 299 {
			log.WithFields(data).Info("Request processed successfully")
		} else {
			data["response_body"] = string(c.Response().Body())
			log.WithFields(data).Error("Request failed")
		}

		return err
	}
}


func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return NewResponse(
				WithError(ErrorUnauthorized),
			).Send(c)
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("token invalid")
			return NewResponse(
				WithError(ErrorUnauthorized),
			).Send(c)
		}

		token := bearer[1]

		id, email, err := util.ValidateToken(token, configs.Cfg.App.Encryption.JWTSecret)
		if err != nil {
			log.Println(err.Error())
			return NewResponse(
				WithError(ErrorUnauthorized),
			).Send(c)
		}

		c.Locals("email", email)
		c.Locals("id", id)

		return c.Next()
	}
}

