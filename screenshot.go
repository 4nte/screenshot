// Package screenshot captures screen-shot image as image.RGBA.
// Mac, Windows, Linux, FreeBSD, OpenBSD, NetBSD, and Solaris are supported.
package screenshot

import (
	"github.com/4nte/screenshot/internal/xwindow"
	"image"
)

// CaptureDisplay captures whole region of displayIndex'th display.
func (c *XgbConnection) CaptureDisplay(displayIndex int) (*image.RGBA, error) {
	rect := xwindow.GetDisplayBounds(c.conn, displayIndex)
	return c.CaptureRect(rect)
}

// CaptureRect captures specified region of desktop.
func (c *XgbConnection) CaptureRect(rect image.Rectangle) (*image.RGBA, error) {
	return c.Capture(rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy())
}
