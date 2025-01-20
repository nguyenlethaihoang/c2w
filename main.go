package main

import (
	"image/color"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	os.Setenv("FYNE_SCALE", "1.0")

	myApp := app.New()
	myWindow := myApp.NewWindow("c2w - cURL to Windows")

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Paste your cURL command here...")

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("The result will be displayed here...")
	output.Wrapping = fyne.TextWrapWord

	convertBtn := widget.NewButtonWithIcon("Convert", theme.ViewRefreshIcon(), func() {
		inputText := input.Text
		result := convertCURL(inputText)
		output.SetText(result)

		customDialog := dialog.NewCustom(
			"Success",
			"OK",
			container.NewVBox(
				widget.NewIcon(theme.ConfirmIcon()),
				widget.NewLabel("Conversion successful!"),
			),
			myWindow,
		)
		customDialog.Show()
	})

	copyBtn := widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {
		myWindow.Clipboard().SetContent(output.Text)

		customDialog := dialog.NewCustom(
			"Copy",
			"Close",
			container.NewVBox(
				widget.NewIcon(theme.ConfirmIcon()),
				widget.NewLabel("Result copied to clipboard!"),
			),
			myWindow,
		)
		customDialog.Show()
	})

	clearBtn := widget.NewButtonWithIcon("Clear", theme.DeleteIcon(), func() {
		input.SetText("")
		output.SetText("")
	})

	toggleThemeBtn := widget.NewButtonWithIcon("Toggle Theme", theme.ColorPaletteIcon(), func() {
		currentTheme := myApp.Settings().Theme()
		if _, isCustom := currentTheme.(*customDarkTheme); isCustom {
			myApp.Settings().SetTheme(theme.LightTheme())
		} else {
			myApp.Settings().SetTheme(&customDarkTheme{})
		}
		myWindow.Content().Refresh()
	})

	buttons := container.NewHBox(convertBtn, copyBtn, clearBtn, toggleThemeBtn)
	content := container.NewHSplit(
		container.NewVScroll(input),
		container.NewVScroll(output),
	)
	myWindow.SetContent(container.NewBorder(nil, buttons, nil, nil, content))

	myWindow.Resize(fyne.NewSize(800, 400))
	myWindow.ShowAndRun()
}

func convertCURL(input string) string {
	input = strings.ReplaceAll(input, "--request", "-X")
	input = strings.ReplaceAll(input, "--header", "-H")
	input = strings.ReplaceAll(input, "--data", "-d")

	input = strings.ReplaceAll(input, "'", "\"")
	input = strings.ReplaceAll(input, "\\", "")
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")

	index := 0
	newString := ""
	check := false
	for {
		if index >= len(input) {
			break
		}
		char := rune(input[index])
		if char == '{' {
			check = true
		}
		if char == '}' {
			check = false
		}
		if check && char == '"' {
			newString += `\"`
		} else {
			newString += string(char)
		}
		index++
	}

	return newString
}

type customDarkTheme struct{}

func (t *customDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 30, G: 30, B: 30, A: 255}
	case theme.ColorNameForeground:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0, G: 150, B: 255, A: 255}
	case theme.ColorNameHover:
		return color.RGBA{R: 0, G: 150, B: 255, A: 100}
	case theme.ColorNameFocus:
		return color.RGBA{R: 0, G: 150, B: 255, A: 255}
	case theme.ColorNameButton:
		return color.RGBA{R: 0, G: 150, B: 255, A: 255}
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 50, G: 50, B: 50, A: 255}
	case theme.ColorNameInputBorder:
		return color.RGBA{R: 100, G: 100, B: 100, A: 255}
	case theme.ColorNamePlaceHolder:
		return color.RGBA{R: 150, G: 150, B: 150, A: 255}
	case theme.ColorNameScrollBar:
		return color.RGBA{R: 100, G: 100, B: 100, A: 255}
	case theme.ColorNameShadow:
		return color.RGBA{R: 0, G: 0, B: 0, A: 100}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (t *customDarkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *customDarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *customDarkTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
