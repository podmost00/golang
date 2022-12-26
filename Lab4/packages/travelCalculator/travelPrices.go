package travelCalculator

import "github.com/andlabs/ui"

func travelPrices() {
	containerWindow := ui.NewWindow("Список цін", 10, 10, false)
	containerWindow.SetMargined(false)

	pricesBox := ui.NewVerticalBox()
	pricesBox.SetPadded(true)

	firstLabel := ui.NewLabel("Болгарія, літо – 100 $")
	secondLabel := ui.NewLabel("Болгарія, зима – 150 $")
	thirdLabel := ui.NewLabel("Німеччина, літо – 160 $")
	fourthLabel := ui.NewLabel("Німеччина, зима – $200")
	fifthLabel := ui.NewLabel("Польща, літо – $120;")
	sixthLabel := ui.NewLabel("Польща, зима – $180;")
	seventhLabel := ui.NewLabel("Вартість індивідуального гіда – $50 в день на всю кількість путівок")
	eighthLabel := ui.NewLabel("Вартість націнки в розмірі 20% за проживання в номері люкс")

	okButton := ui.NewButton("OK")

	pricesBox.Append(firstLabel, false)
	pricesBox.Append(secondLabel, false)
	pricesBox.Append(thirdLabel, false)
	pricesBox.Append(fourthLabel, false)
	pricesBox.Append(fifthLabel, false)
	pricesBox.Append(sixthLabel, false)
	pricesBox.Append(seventhLabel, false)
	pricesBox.Append(eighthLabel, false)

	pricesBox.Append(okButton, false)

	pricesGroup := ui.NewGroup("Ціни за 1 день подорожі")
	pricesGroup.SetChild(pricesBox)

	containerWindow.SetChild(pricesGroup)
	containerWindow.Show()

	okButton.OnClicked(func(*ui.Button) {
		containerWindow.Destroy()
	})

	containerWindow.OnClosing(func(*ui.Window) bool {
		containerWindow.Hide()
		return true
	})
}
