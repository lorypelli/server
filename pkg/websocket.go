package pkg

import (
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/lorypelli/server/internal"
)

var reload = []byte("reload")

func serveWebsocket(dir string, network bool) {
	watcher := Watch(dir)
	app := fiber.New()
	app.Get("/", websocket.New(func(conn *websocket.Conn) {
		defer conn.Close()
		changes, unsubscribe := watcher.Subscribe()
		defer unsubscribe()
		closed := make(chan struct{})
		go func() {
			defer close(closed)
			for {
				if _, _, err := conn.ReadMessage(); err != nil {
					return
				}
			}
		}()
		for {
			select {
			case <-closed:
				return
			case <-changes:
				if err := conn.WriteMessage(websocket.TextMessage, reload); err != nil {
					return
				}
			}
		}
	}))
	listen(app, network, internal.WSPort)
}
