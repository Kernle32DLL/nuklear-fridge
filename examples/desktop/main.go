package main

import (
	"log"
	"runtime"
	"time"

	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/Kernle32DLL/nuklear-fridge/nkf/color"
	"github.com/Kernle32DLL/nuklear-fridge/nkf/label"
	"github.com/Kernle32DLL/nuklear-fridge/nkf/property"
	"github.com/Kernle32DLL/nuklear-fridge/nkf/row"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"github.com/xlab/closer"
	"strconv"
)

const (
	winWidth  = 400
	winHeight = 500

	maxVertexBuffer  = 512 * 1024
	maxElementBuffer = 128 * 1024
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		closer.Fatalln(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(winWidth, winHeight, "Nuklear Demo", nil, nil)
	if err != nil {
		closer.Fatalln(err)
	}
	win.MakeContextCurrent()

	width, height := win.GetSize()
	log.Printf("glfw: created window %dx%d", width, height)

	if err := gl.Init(); err != nil {
		closer.Fatalln("opengl: init failed:", err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))

	ctx := nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)

	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	sansFont := nk.NkFontAtlasAddFromBytes(atlas, MustAsset("assets/FreeSans.ttf"), 16, nil)
	// sansFont := nk.NkFontAtlasAddDefault(atlas, 16, nil)
	nk.NkFontStashEnd()
	if sansFont != nil {
		nk.NkStyleSetFont(ctx, sansFont.Handle())
	}

	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	closer.Bind(func() {
		close(exitC)
		<-doneC
	})

	state := &State{
		bgColor: nk.NkRgba(28, 48, 62, 255),
	}
	fridge := buildFridge(state)

	fpsTicker := time.NewTicker(time.Second / 30)
	for {
		select {
		case <-exitC:
			nk.NkPlatformShutdown()
			glfw.Terminate()
			fpsTicker.Stop()
			close(doneC)
			return
		case <-fpsTicker.C:
			if win.ShouldClose() {
				close(exitC)
				continue
			}
			glfw.PollEvents()

			bg := make([]float32, 4)
			nk.NkColorFv(bg, state.bgColor)
			width, height := win.GetSize()
			gl.Viewport(0, 0, int32(width), int32(height))
			gl.Clear(gl.COLOR_BUFFER_BIT)
			gl.ClearColor(bg[0], bg[1], bg[2], bg[3])

			fridge.Render(ctx)

			win.SwapBuffers()
		}
	}
}

func buildFridge(stateReference *State) *nkf.Fridge {
	fridge := nkf.NewFridge(true, maxVertexBuffer, maxElementBuffer)

	window := nkf.NewWindow(50, 50, 230, 250, "Fridge Demo",
		nk.WindowBorder|nk.WindowMovable|nk.WindowScalable|nk.WindowClosable|nk.WindowTitle)

	buttonRow := row.NewStaticAsChild(window, 30, 80, 1)
	label.NewButtonAsChild(buttonRow, "button", func(buttonLabel *label.Button) {
		log.Println("[INFO] button pressed!")
	})

	optionRow := row.NewDynamicAsChild(window, 30, 2)
	{
		var option1, option2 *label.Option
		updateState := func() {
			option1.SetActive(stateReference.opt == Easy)
			option2.SetActive(stateReference.opt == Hard)
		}

		option1 = label.NewOptionAsChild(optionRow, "easy", stateReference.opt == Easy, func(optionLabel *label.Option) {
			if stateReference.opt != Easy {
				log.Println("[INFO] easy selected!")
				stateReference.opt = Easy
			}

			updateState()
		})
		option2 = label.NewOptionAsChild(optionRow, "hard", stateReference.opt == Hard, func(optionLabel *label.Option) {
			if stateReference.opt != Hard {
				log.Println("[INFO] hard selected!")
				stateReference.opt = Hard
			}

			updateState()
		})
	}

	propertyRow := row.NewDynamicAsChild(window, 25, 1)
	property.NewIntegerAsChild(propertyRow, "Compression:", 0, stateReference.prop, 100, 10, 1, func(integer *property.Integer) {
		log.Println("[INFO] compression changed to " + strconv.FormatInt(int64(integer.Value()), 10))
	})

	labelRow := row.NewDynamicAsChild(window, 20, 1)
	nkf.NewLabelAsChild(labelRow, "background: ", nk.TextLeft)

	colorRow := row.NewDynamicAsChild(window, 25, 1)
	colorComboBox := color.NewComboBoxAsChild(colorRow, -1, 400, stateReference.bgColor, nil)
	{
		colorPickerRow := row.NewDynamicAsChild(colorComboBox, 120, 1)
		{
			var colorPicker *color.Picker
			var rProp, gProp, bProp, aProp *property.Integer

			colorPicker = color.NewPickerAsChild(colorPickerRow, stateReference.bgColor, nk.ColorFormatRGBA, func(_ *color.Picker) {
				stateReference.bgColor = colorPicker.Color()

				// Re-apply color to properties
				r, g, b, a := stateReference.bgColor.RGBAi()
				rProp.SetValue(r)
				gProp.SetValue(g)
				bProp.SetValue(b)
				aProp.SetValue(a)

				// Re-apply color to combo box
				colorComboBox.SetColor(stateReference.bgColor)
			})

			colorPropertiesRow := row.NewDynamicAsChild(colorComboBox, 25, 1)

			propertiesChanged := func(_ *property.Integer) {
				stateReference.bgColor.SetRGBAi(rProp.Value(), gProp.Value(), bProp.Value(), aProp.Value())

				// Re-apply color to picker
				colorPicker.SetColor(stateReference.bgColor)

				// Re-apply color to combo box
				colorComboBox.SetColor(stateReference.bgColor)
			}

			r, g, b, a := stateReference.bgColor.RGBAi()
			rProp = property.NewIntegerAsChild(colorPropertiesRow, "#R:", 0, r, 255, 1, 1, propertiesChanged)
			gProp = property.NewIntegerAsChild(colorPropertiesRow, "#G:", 0, g, 255, 1, 1, propertiesChanged)
			bProp = property.NewIntegerAsChild(colorPropertiesRow, "#B:", 0, b, 255, 1, 1, propertiesChanged)
			aProp = property.NewIntegerAsChild(colorPropertiesRow, "#A:", 0, a, 255, 1, 1, propertiesChanged)
		}
	}

	fridge.AddWindow(window)

	return fridge
}

type Option uint8

const (
	Easy Option = 0
	Hard Option = 1
)

type State struct {
	bgColor nk.Color
	prop    int32
	opt     Option
}

func onError(code int32, msg string) {
	log.Printf("[glfw ERR]: error %d: %s", code, msg)
}
