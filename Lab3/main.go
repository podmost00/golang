package main

import (
	"Lab3_Patterns/packages/patterns/factory_method"
	"Lab3_Patterns/packages/patterns/strategy"
	"Lab3_Patterns/packages/patterns/wrappper"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Lab3. Паттерни проектування")
	fmt.Println("\n1) Фабричний метод;\n2) Стратегія;\n3) Декоратор;")

	var menu int
	fmt.Print("\nОберіть паттерн: ")
	fmt.Fscan(os.Stdin, &menu)

	switch menu {
	case 1:
		fmt.Println("\nОбрано паттерн фабричний метод.")

		types := []string{factory_method.BallType, factory_method.TShortsType, factory_method.GoalkeeperEquipType}
		fmt.Print("\nДоступні категорії: ")
		for _, typeItem := range types {
			fmt.Print(typeItem, " ")
		}

		var category string
		fmt.Print("\nВведіть доступну категорії: ")
		fmt.Fscan(os.Stdin, &category)

		result, err := factory_method.Create(category)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\nСтворений фабричним методом екземпляр товару:")
			fmt.Println(result.GetInfo())
		}

	case 2:
		fmt.Println("\nОбрано паттерн стратегія.")
		deliveries := []strategy.Delivery{&strategy.RoadStrategy{}, &strategy.CourierStrategy{}, &strategy.BikeStrategy{}}
		manager := strategy.DeliveryManager{}

		var distance float64
		fmt.Print("\nВведіть відстань для доставки: ")
		fmt.Fscan(os.Stdin, &distance)

		fmt.Println("\nПобудовані маршрути залежно від типу доставки:")
		for _, delivery := range deliveries {
			manager.SetDelivery(delivery)
			manager.Route(distance)
		}
	case 3:
		fmt.Println("\nОбрано паттерн декоратор.")

		fmt.Println("Типи футболок: 1) проста, 2) аматорська, 3) професійна.")
		var tShortType int
		fmt.Fscan(os.Stdin, &tShortType)

		simple := wrappper.SimpleTShort{}

		switch tShortType {
		case 1:
			fmt.Printf("\nЦіна вашої простої футболки %.2f", simple.GetPrice())
		case 2:
			var number int
			fmt.Print("Введіть номер: ")
			fmt.Fscan(os.Stdin, &number)

			amateur := wrappper.AmateurTShort{
				Number:  number,
				Wrapper: simple,
			}

			fmt.Printf("\nЦіна вашої аматорський футболки %.2f під номером %d", amateur.GetPrice(), amateur.Number)
		case 3:
			var number int
			fmt.Print("Введіть номер: ")
			fmt.Fscan(os.Stdin, &number)

			var name string
			fmt.Print("Введіть прізвище: ")
			fmt.Fscan(os.Stdin, &name)

			professional := wrappper.ProfessionalTShort{
				Number:  number,
				Name:    name,
				Wrapper: simple,
			}

			fmt.Printf("\nЦіна вашої професійної футболки %.2f під номером %d та прізвищем %s", professional.GetPrice(), professional.Number, professional.Name)
		default:
			fmt.Println("Обрано неправильний тип футболок.")
		}

	default:
		fmt.Println("Паттерну під таким номером немає.")
	}
}
