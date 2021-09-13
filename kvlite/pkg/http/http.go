package http

import (
	"kvlite/pkg/kv_agent"

	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	server   *fiber.App
	kv_agent kv_agent.KVAgent
}

func (s *HTTPServer) Listen(port string) {
	s.server.Listen(port)
}

func (s *HTTPServer) Shutdown() {
	s.server.Shutdown()
}

func (s *HTTPServer) Init(agent kv_agent.KVAgent) {
	api := fiber.New()

	api.Get("/:key", s.Get)
	api.Put("/:key", s.Set)
	api.Delete("/:key", s.Delete)

	s.server = api
	s.kv_agent = agent

}

// API endpoint

func (s *HTTPServer) Get(c *fiber.Ctx) error {
	if c.Params("key") == "" {
		return c.SendString("KVLITE Error: Need a key to get")
	}
	// KV agent
	var value []byte
	value, closer, err := s.kv_agent.Get(c.Params("key"))
	if err != nil {
		return c.SendString("KVLITE Error: " + err.Error())
	}

	// Res
	closer.Close()
	return c.JSON(
		fiber.Map{
			"message": "KVLITE Success: Get key",
			"key":     c.Params("key"),
			"value":   string(value),
		},
	)
}

type SetReq struct {
	Value string `json:value`
}

func (s *HTTPServer) Set(c *fiber.Ctx) error {

	// Check key
	if c.Params("key") == "" {
		return c.SendString("KVLITE Error: Need a key to set")
	}

	// Parse body value
	req := &SetReq{}
	err := c.BodyParser(req)
	if err != nil || req.Value == "" {
		return c.SendString("KVLITE Error: " + err.Error())
	}

	// KV agent
	err = s.kv_agent.Set(
		[]byte(c.Params("key")),
		[]byte(req.Value),
	)

	if err != nil {
		return c.SendString("KVLITE Error: " + err.Error())
	}
	// Res
	return c.JSON(
		fiber.Map{
			"message": "KVLITE Success: Set key",
			"key":     c.Params("key"),
			"value":   req.Value,
		},
	)

}

func (s *HTTPServer) Delete(c *fiber.Ctx) error {
	if c.Params("key") == "" {
		return c.SendString("KVLITE Error: Need a key to delete")
	}

	// KV agent
	err := s.kv_agent.Delete([]byte(c.Params("key")))
	if err != nil {
		return c.SendString("KVLITE Error: " + err.Error())
	}

	// Res
	return c.JSON(
		fiber.Map{
			"message": "KVLITE Success: Delete key",
			"key":     c.Params("key"),
		},
	)
}
