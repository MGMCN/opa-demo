package web

import (
	"context"
	"github.com/MGMCN/opa-demo/middlewares"
	"github.com/gofiber/fiber/v2"
)

type apiServer struct {
	app       *fiber.App
	opaServer *middlewares.OpaServer
}

func NewApiServer() *apiServer {
	return &apiServer{}
}

func (s *apiServer) InitServer(ctx context.Context) error {
	s.app = fiber.New()
	if err := s.initOpaServer(ctx); err != nil {
		// handle error
		return err
	}
	s.setupMiddlewares(s.opaChecker)
	s.setupRoutes("/test_post", s.post)
	return nil
}

func (s *apiServer) initOpaServer(ctx context.Context) error {
	s.opaServer = middlewares.NewOpaServer(ctx)
	return s.opaServer.InitServer()
}

func (s *apiServer) Serve(port string) error {
	return s.app.Listen(port)
}

func (s *apiServer) setupRoutes(path string, handler fiber.Handler) {
	s.app.Post(path, handler)
}

func (s *apiServer) setupMiddlewares(middleware fiber.Handler) {
	s.app.Use(middleware)
}

func (s *apiServer) opaChecker(c *fiber.Ctx) error {
	var err error

	inputData := map[string]interface{}{
		"path":   "/test_post",
		"method": "POST",
		"body": map[string]interface{}{
			"type": c.FormValue("type"),
		},
	}

	if s.opaServer.Check(inputData) {
		err = c.Next()
	} else {
		err = c.Send([]byte("only admin can access this method!"))
	}

	return err
}

func (s *apiServer) post(c *fiber.Ctx) error {
	var err error

	// Plain authentication approach
	//userType := c.FormValue("type")
	//if userType != "admin" {
	//	err = c.Send([]byte("only admin can access this method!"))
	//} else {
	//	err = c.Send([]byte("call get method!"))
	//}

	err = c.Send([]byte("call get method!"))

	if err != nil {
		return err
	}
	return nil
}
