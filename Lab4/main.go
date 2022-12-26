package main

import (
	"Lab4_Graphical_interface/packages/travelCalculator"
	"Lab4_Graphical_interface/packages/windowsCalculator"
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"os"
)

func main() {
	fmt.Println("Lab4. Графічний інтерфейс")
	fmt.Println("\n1) Обчислення вартості склопакета;\n2) Обчислення вартості туру.")

	var menu int
	fmt.Print("Введіть обраний пункт меню: ")
	fmt.Fscan(os.Stdin, &menu)

	var err error
	switch menu {
	case 1:
		err = ui.Main(windowsCalculator.WindowsCalculator)
	case 2:
		err = ui.Main(travelCalculator.TravelCalculator)
	default:
		fmt.Println("Обрано неправльний пункт меню!")
	}

	if err != nil {
		panic(err)
	}
}
