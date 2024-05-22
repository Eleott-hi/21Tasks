package ex02

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

func main() {
	cocoa.TerminateAfterWindowsClose = true
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		// Create the window
		rect := core.Rect(0, 0, 300, 200)
		window := cocoa.NSWindow_Init(rect, cocoa.NSClosableWindowMask|cocoa.NSTitledWindowMask, cocoa.NSBackingStoreBuffered, false)
		window.SetTitle("School 21")
		window.MakeKeyAndOrderFront(nil)
	})
	app.Run()
}
