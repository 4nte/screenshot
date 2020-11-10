package screenshot

import (
	"github.com/4nte/screenshot/internal/xwindow"
	"github.com/BurntSushi/xgb"
	"image"
	"log"
	"sync"
)

type XgbConnection struct {
	conn *xgb.Conn
	mu   sync.Mutex
}

func NewXgbConnection() (*XgbConnection, error) {
	conn, err := xgb.NewConn()
	if err != nil {
		return nil, err
	}

	window := &XgbConnection{
		conn: conn,
	}

	// Handle connection close
	go func() {
		for true {
			event, err := conn.WaitForEvent()
			if event == nil && err == nil {
				window.mu.Lock()

				log.Println("x Connection closed")
				// Connection is closed, open a new one.
				newConn, err := xgb.NewConn()
				if err != nil {
					panic(err)
				}
				log.Println("Opened a new X connection")

				window.conn = newConn
				window.mu.Unlock()

			}
		}
	}()

	return window, nil
}

// Capture returns screen capture of specified desktop region.
// x and y represent distance from the upper-left corner of main display.
// Y-axis is downward direction. This means coordinates system is similar to Windows OS.
func (c *XgbConnection) Capture(x, y, width, height int) (*image.RGBA, error) {
	return xwindow.Capture(c.conn, x, y, width, height)
}

// NumActiveDisplays returns the number of active displays.
func (c *XgbConnection) NumActiveDisplays() int {
	return xwindow.NumActiveDisplays(c.conn)
}

// GetDisplayBounds returns the bounds of displayIndex'th display.
// The main display is displayIndex = 0.
func (c *XgbConnection) GetDisplayBounds(displayIndex int) image.Rectangle {
	return xwindow.GetDisplayBounds(c.conn, displayIndex)
}
