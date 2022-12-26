package main

import "fmt"

func main() {
	var x, y, z uint8

	x = 9
	y = 28
	z = x

	fmt.Println("Битовые операции")

	fmt.Printf("^x      = ^(%d)      = ^(%.8b)            = %.8b = %d\n", x, x, ^x, ^x)
	fmt.Printf("x << 2  = (%d << 2)  = (%.8b << 2)        = %.8b = %d\n", x, x, x<<2, x<<2)
	fmt.Printf("x >> 2  = (%d >> 2)  = (%.8b >> 2)        = %.8b = %d\n", x, x, x>>2, x>>2)
	fmt.Printf("x & y   = (%d & %d)  = (%.8b & %.8b)  = %.8b = %d\n", x, y, x, y, x&y, x&y)
	fmt.Printf("x | y   = (%d | %d)  = (%.8b | %.8b)  = %.8b = %d\n", x, y, x, y, x|y, x|y)
	fmt.Printf("x ^ y   = (%d ^ %d)  = (%.8b ^ %.8b)  = %.8b = %d\n", x, y, x, y, x^y, x^y)
	fmt.Printf("x &^ y  = (%d &^ %d) = (%.8b &^ %.8b) = %.8b = %d\n", x, y, x, y, x&^y, x&^y)
	fmt.Printf("x %% y   = (%d %% %d)  = (%.8b %% %.8b)  = %.8b = %d\n", x, y, x, y, x%y, x%y)

	fmt.Println("\nБитовые операции с присваиванием")

	x = z
	x &= y
	fmt.Printf("x &= y   = (%d &= %d)  = (%.8b &= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x |= y
	fmt.Printf("x |= y   = (%d |= %d)  = (%.8b |= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x ^= y
	fmt.Printf("x ^= y   = (%d ^= %d)  = (%.8b ^= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x &^= y
	fmt.Printf("x &^= y  = (%d &^= %d) = (%.8b &^= %.8b) = %.8b = %d\n", z, y, z, y, x, x)
	x = z
	x %= y
	fmt.Printf("x %%= y   = (%d %%= %d)  = (%.8b %%= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)

	//Задание.
	//1. Пояснить результаты операций
	// Додав пояснення у звіт
}
