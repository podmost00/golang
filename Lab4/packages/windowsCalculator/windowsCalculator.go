package windowsCalculator

import (
	"github.com/andlabs/ui"
	"strconv"
)

func WindowsCalculator() {
	window := ui.NewWindow("Калькулятор склопакета", 100, 100, false)
	containerBox := ui.NewVerticalBox()
	containerBox.SetPadded(true)

	container := ui.NewHorizontalBox()
	container.SetPadded(true)

	inputGroup := ui.NewGroup("Розмір вікна")
	inputGroup.SetMargined(true)

	inputBox := ui.NewHorizontalBox()
	inputBox.SetPadded(true)

	inputForm := ui.NewForm()
	inputForm.SetPadded(true)

	height := ui.NewEntry()
	height.SetText("50")
	width := ui.NewEntry()
	width.SetText("50")

	materialCombobox := ui.NewCombobox()
	materialCombobox.Append("Дерево")
	materialCombobox.Append("Метал")
	materialCombobox.Append("Металопластик")
	materialCombobox.SetSelected(0)

	langCombobox := ui.NewCombobox()
	langCombobox.Append("Обрахувати мовою C")
	langCombobox.Append("Обрахувати мовою Go")
	langCombobox.SetSelected(1)

	inputForm.Append("Ширина, см", height, false)
	inputForm.Append("Висота, см", width, false)
	inputForm.Append("Матеріали", materialCombobox, false)

	inputBox.Append(inputForm, false)
	inputGroup.SetChild(inputBox)

	resultBox := ui.NewVerticalBox()
	resultBox.SetPadded(true)

	resultGroup := ui.NewGroup("Ваша ціна")
	resultGroup.SetMargined(true)

	glassGroup := ui.NewGroup("Cклопакет")
	glassGroup.SetMargined(true)

	glassBox := ui.NewVerticalBox()
	glassBox.SetPadded(true)

	glassCombobox := ui.NewCombobox()
	glassCombobox.Append("Однокамерний")
	glassCombobox.Append("Двокамерний")
	glassCombobox.SetSelected(0)

	glassBox.Append(glassCombobox, false)
	sillCheck := ui.NewCheckbox("Підвіконня")

	glassBox.Append(sillCheck, false)
	glassGroup.SetChild(glassBox)

	container.Append(inputGroup, false)
	container.Append(glassGroup, false)

	containerBox.Append(container, false)

	priceBox := ui.NewVerticalBox()
	priceBox.SetPadded(true)

	resultLabel := ui.NewLabel("")
	priceBox.Append(resultLabel, false)

	resultGroup.SetChild(priceBox)
	resultBox.Append(resultGroup, false)

	buttonBox := ui.NewVerticalBox()
	buttonBox.SetPadded(true)

	button := ui.NewButton("Розрахувати")
	buttonBox.Append(button, false)

	errorWindow := ui.NewWindow("Помилка", 10, 10, false)

	buttonPrices := ui.NewButton("Список цін")
	buttonBox.Append(buttonPrices, false)

	buttonPrices.OnClicked(func(*ui.Button) {
		windowPrices()
	})

	button.OnClicked(func(*ui.Button) {
		selectedMaterial := materialCombobox.Selected()
		selectedGlass := glassCombobox.Selected()
		selectedLang := langCombobox.Selected()

		heightText := height.Text()
		widthText := width.Text()

		heightNumber, errHeight := strconv.ParseFloat(heightText, 64)
		widthNumber, errWidth := strconv.ParseFloat(widthText, 64)

		if errHeight == nil && errWidth == nil {
			var resultPrice string

			if selectedLang == 0 {
				resultPrice = calculatePriceC(selectedMaterial, selectedGlass, sillCheck.Checked(), heightNumber, widthNumber)
			} else {
				result, err := calculatePrice(selectedMaterial, selectedGlass, sillCheck.Checked(), heightNumber, widthNumber)

				if err != nil {
					ui.MsgBoxError(errorWindow, "Помилка!", "Перевірте введенні дані")
					return
				}

				resultPrice = result
			}

			resultLabel.SetText(resultPrice)
		} else {
			ui.MsgBoxError(errorWindow, "Помилка!", "Перевірте введенні дані")
		}

	})

	containerBox.Append(langCombobox, false)
	containerBox.Append(buttonBox, false)
	containerBox.Append(resultBox, false)

	window.SetChild(containerBox)
	window.Show()
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
}
