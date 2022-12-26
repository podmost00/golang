package travelCalculator

import (
	"github.com/andlabs/ui"
	"strconv"
)

func TravelCalculator() {
	containerWindow := ui.NewWindow("Калькулятор вартості туру", 60, 70, false)
	containerWindow.SetMargined(false)

	containerBox := ui.NewHorizontalBox()
	containerBox.SetPadded(true)

	tourBox := ui.NewVerticalBox()
	tourBox.SetPadded(true)

	tourGroup := ui.NewGroup("Основні параметри")

	tourForm := ui.NewForm()
	tourForm.SetPadded(true)

	dayEntry := ui.NewEntry()
	dayEntry.SetText("7")

	countryCombobox := ui.NewCombobox()
	countryCombobox.Append("Болгарія")
	countryCombobox.Append("Німеччина")
	countryCombobox.Append("Польща")
	countryCombobox.SetSelected(1)

	seasonCombobox := ui.NewCombobox()
	seasonCombobox.Append("Літо")
	seasonCombobox.Append("Зима")
	seasonCombobox.SetSelected(1)

	langCombobox := ui.NewCombobox()
	langCombobox.Append("C")
	langCombobox.Append("Go")
	langCombobox.SetSelected(1)

	tourForm.Append("Кількість днів", dayEntry, false)
	tourForm.Append("Країна", countryCombobox, false)
	tourForm.Append("Пора року", seasonCombobox, false)
	tourForm.Append("Мова програмування", langCombobox, false)

	tourBox.Append(tourForm, false)

	tourGroup.SetChild(tourBox)

	containerBox.Append(tourGroup, false)

	additionalGroup := ui.NewGroup("Додаткові параметри")
	additionalGroup.SetMargined(true)

	additionalBox := ui.NewVerticalBox()
	additionalBox.SetPadded(true)

	luxuryCheckbox := ui.NewCheckbox("Номер люкс")
	guideCheckbox := ui.NewCheckbox("Індивідуальний гід")

	additionalBox.Append(luxuryCheckbox, false)
	additionalBox.Append(guideCheckbox, false)

	additionalGroup.SetChild(additionalBox)

	tourBox.Append(additionalGroup, false)

	resultBox := ui.NewVerticalBox()
	resultBox.SetPadded(true)

	resultGroup := ui.NewGroup("Ваша ціна")
	resultGroup.SetMargined(true)

	priceBox := ui.NewVerticalBox()
	priceBox.SetPadded(true)

	resultLabel := ui.NewLabel("")
	priceBox.Append(resultLabel, false)

	resultGroup.SetChild(priceBox)
	resultBox.Append(resultGroup, false)

	buttonsBox := ui.NewHorizontalBox()

	priceButton := ui.NewButton("Розрахувати ціну")
	pricesListButton := ui.NewButton("Список цін")

	buttonsBox.Append(priceButton, false)
	buttonsBox.Append(pricesListButton, false)

	resultBox.Append(buttonsBox, false)

	containerBox.Append(resultBox, false)

	pricesListButton.OnClicked(func(*ui.Button) {
		travelPrices()
	})

	priceButton.OnClicked(func(*ui.Button) {
		dayText := dayEntry.Text()
		dayValue, parseErr := strconv.ParseInt(dayText, 0, 64)

		errorWindow := ui.NewWindow("Помилка", 10, 10, false)

		selectedCountry := countryCombobox.Selected()
		selectedSeason := seasonCombobox.Selected()
		selectedLang := langCombobox.Selected()

		isLuxury := luxuryCheckbox.Checked()
		isWithGuide := guideCheckbox.Checked()

		if parseErr == nil {
			var resultPrice string

			if selectedLang == 0 {
				resultPrice = calculatePriceC(dayValue, selectedCountry, selectedSeason, isLuxury, isWithGuide)
			} else {
				res, resErr := calculatePrice(dayValue, selectedCountry, selectedSeason, isLuxury, isWithGuide)

				if resErr != nil {
					ui.MsgBoxError(errorWindow, "Помилка!", "Перевірте правильність введених параметрів")
				}

				resultPrice = res
			}

			resultLabel.SetText(resultPrice)
		} else {
			ui.MsgBoxError(errorWindow, "Помилка!", "Перевірте правильність введених параметрів")
		}
	})

	containerWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	containerWindow.SetChild(containerBox)
	containerWindow.Show()
}
