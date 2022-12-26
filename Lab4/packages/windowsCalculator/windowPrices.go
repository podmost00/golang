package windowsCalculator

import (
	"github.com/andlabs/ui"
)

func windowPrices() {
	mainWindow := ui.NewWindow("Список цін", 10, 10, false)
	mainWindow.SetMargined(false)

	verticalBox := ui.NewVerticalBox()
	verticalBox.SetPadded(true)

	firstLabel := ui.NewLabel("Однокамерний, дерев'яний – 2.5 грн")
	secondLabel := ui.NewLabel("Двокамерний, дерев'яний – 3 грн")
	thirdLabel := ui.NewLabel("Однокамерний, металевий – 0.5 грн")
	fourthLabel := ui.NewLabel("Двокамерний, металевий – 1 грн")
	fifthLabel := ui.NewLabel("Однокамерний, металопластиковий – 1.5 грн")
	sixthLabel := ui.NewLabel("Двокамерний, металопластиковий - 2 грн")
	seventhLabel := ui.NewLabel("Вартість підвіконня – 350 грн")

	okButton := ui.NewButton("OK")

	verticalBox.Append(firstLabel, false)
	verticalBox.Append(secondLabel, false)
	verticalBox.Append(thirdLabel, false)
	verticalBox.Append(fourthLabel, false)
	verticalBox.Append(fifthLabel, false)
	verticalBox.Append(sixthLabel, false)
	verticalBox.Append(seventhLabel, false)

	verticalBox.Append(okButton, false)

	group := ui.NewGroup("Ціни за 1 сантиметр квадратний склопакета")
	group.SetChild(verticalBox)

	mainWindow.SetChild(group)
	mainWindow.Show()

	okButton.OnClicked(func(*ui.Button) {
		mainWindow.Destroy()
	})

	mainWindow.OnClosing(func(*ui.Window) bool {
		mainWindow.Hide()
		return true
	})
}
