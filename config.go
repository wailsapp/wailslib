package wails

import (
	"net/url"
	"strings"

	"github.com/wailsapp/wails/runtime"
)

// AppConfig is the configuration structure used when creating a Wails App object
type AppConfig struct {
	// The width and height of your application in pixels
	Width, Height int

	// The title to put in the title bar
	Title string

	// The HTML your app should use. If you leave it blank, a default will be used:
	// <!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta http-equiv="IE=edge" content="IE=edge"></head><body><div id="app"></div><script type="text/javascript"></script></body></html>
	HTML string

	// The Javascript your app should use. Normally this should be generated by a bundler.
	JS string

	// The CSS your app should use. Normally this should be generated by a bundler.
	CSS string

	// The colour of your window. Can take "#fff", "rgb(255,255,255)", "rgba(255,255,255,1)" formats
	Colour string

	// Indicates whether your app should be resizable
	Resizable bool

	// Minimum width of a resizable window. If set, MinHeight should also be set.
	MinWidth int

	// Minimum height of a resizable window. If set, MinWidth should also be set.
	MinHeight int

	// Maximum width of a resizable window. If set, MaxHeight should also be set.
	MaxWidth int

	// Maximum height of a resizable window. If set, MaxWidth should also be set.
	MaxHeight int

	// Indicated if the devtools should be disabled
	DisableInspector bool
}

// GetWidth returns the desired width
func (a *AppConfig) GetWidth() int {
	return a.Width
}

// GetHeight returns the desired height
func (a *AppConfig) GetHeight() int {
	return a.Height
}

// GetTitle returns the desired window title
func (a *AppConfig) GetTitle() string {
	return a.Title
}

// GetHTML returns the default HTML
func (a *AppConfig) GetHTML() string {
	if len(a.HTML) > 0 {
		a.HTML = url.QueryEscape(a.HTML)
		a.HTML = "data:text/html," + strings.ReplaceAll(a.HTML, "+", "%20")
		a.HTML = strings.ReplaceAll(a.HTML, "%3D", "=")
	}
	return a.HTML
}

// GetResizable returns true if the window should be resizable
func (a *AppConfig) GetResizable() bool {
	return a.Resizable
}

// GetMinWidth returns the minimum width of the window
func (a *AppConfig) GetMinWidth() int {
	return a.MinWidth
}

// GetMinHeight returns the minimum height of the window
func (a *AppConfig) GetMinHeight() int {
	return a.MinHeight
}

// GetMaxWidth returns the maximum width of the window
func (a *AppConfig) GetMaxWidth() int {
	return a.MaxWidth
}

// GetMaxHeight returns the maximum height of the window
func (a *AppConfig) GetMaxHeight() int {
	return a.MaxHeight
}

// GetDisableInspector returns true if the inspector should be disabled
func (a *AppConfig) GetDisableInspector() bool {
	return a.DisableInspector
}

// GetColour returns the colour
func (a *AppConfig) GetColour() string {
	return a.Colour
}

// GetCSS returns the user CSS
func (a *AppConfig) GetCSS() string {
	return a.CSS
}

// GetJS returns the user Javascript
func (a *AppConfig) GetJS() string {
	return a.JS
}

func (a *AppConfig) merge(in *AppConfig) error {
	if in.CSS != "" {
		a.CSS = in.CSS
	}
	if in.Title != "" {
		a.Title = runtime.ProcessEncoding(in.Title)
	}

	if in.Colour != "" {
		a.Colour = in.Colour
	}

	if in.HTML != "" {
		a.HTML = in.HTML
	}

	if in.JS != "" {
		a.JS = in.JS
	}

	if in.HTML != "" {
		a.HTML = in.HTML
	}

	if in.Width != 0 {
		a.Width = in.Width
	}
	if in.Height != 0 {
		a.Height = in.Height
	}

	if in.MinWidth != 0 {
		a.MinWidth = in.MinWidth
	}

	if in.MinHeight != 0 {
		a.MinHeight = in.MinHeight
	}

	if in.MaxWidth != 0 {
		a.MaxWidth = in.MaxWidth
	}

	if in.MaxHeight != 0 {
		a.MaxHeight = in.MaxHeight
	}

	a.Resizable = in.Resizable
	a.DisableInspector = in.DisableInspector

	return nil
}

// Creates the default configuration
func newConfig(userConfig *AppConfig) (*AppConfig, error) {
	result := &AppConfig{
		Width:     800,
		Height:    600,
		Resizable: true,
		MinWidth:  -1,
		MinHeight: -1,
		MaxWidth:  -1,
		MaxHeight: -1,
		Title:     "My Wails App",
		Colour:    "#FFF", // White by default
		HTML:      defaultHTML,
	}

	if userConfig != nil {
		err := result.merge(userConfig)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

var defaultHTML = `<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
</head>

<body>
  <div id="app"></div>
</body>

</html>`
