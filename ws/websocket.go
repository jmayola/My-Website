package ws

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

// Add more data to this type if needed
type client struct {
	isClosing bool
	mu        sync.Mutex
}

var (
	clients    = make(map[*websocket.Conn]*client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
	register   = make(chan *websocket.Conn)
	broadcast  = make(chan string)
	unregister = make(chan *websocket.Conn)
)

func GetWs(c *websocket.Conn) {
	fmt.Println(c.Locals("Host")) // "Localhost:3000"
	defer func() {
		c.Close()
	}()
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}

			return // Calls the deferred function, i.e. closes the connection on error
		}
		log.Printf("recv: %s", msg)
		err = c.WriteMessage(mt, []byte(strings.ToUpper(string(msg))))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
