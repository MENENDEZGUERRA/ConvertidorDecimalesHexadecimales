package main

import (
	"fmt"
	"math"
)

func main() {
	// Solicitar al usuario que ingrese una opción
	fmt.Println("Ingrese una opción:")
	fmt.Println("1. Convertir de hexadecimal a binario")
	fmt.Println("2. Convertir de binario a hexadecimal")
	var option int
	fmt.Scan(&option)

	if option == 1 {
		// Si el usuario eligió la opción 1, solicitar que ingrese un número hexadecimal de 3 dígitos
		fmt.Println("Ingrese un número hexadecimal de 3 dígitos:")
		var hex string
		fmt.Scan(&hex)

		// Convertir cada dígito hexadecimal a su equivalente binario de 4 bits
		var bin string
		for _, c := range hex {
			switch c {
			case '0':
				bin += "0000"
			case '1':
				bin += "0001"
			case '2':
				bin += "0010"
			case '3':
				bin += "0011"
			case '4':
				bin += "0100"
			case '5':
				bin += "0101"
			case '6':
				bin += "0110"
			case '7':
				bin += "0111"
			case '8':
				bin += "1000"
			case '9':
				bin += "1001"
			case 'a', 'A':
				bin += "1010"
			case 'b', 'B':
				bin += "1011"
			case 'c', 'C':
				bin += "1100"
			case 'd', 'D':
				bin += "1101"
			case 'e', 'E':
				bin += "1110"
			case 'f', 'F':
				bin += "1111"
			}
		}

		// Imprimir el resultado
		fmt.Printf("El número %s en binario es: %s", hex, bin)
	} else if option == 2 {
		// Si el usuario eligió la opción 2, solicitar que ingrese un número binario de 12 dígitos
		fmt.Println("Ingrese un número binario de 12 dígitos:")
		var binary string
		fmt.Scan(&binary)

		// Convertir el número binario a su equivalente decimal
		var dec int
		for i, c := range binary {
			if c == '1' {
				dec += int(math.Pow(2, float64(len(binary)-1-i)))
			}
		}

		// Convertir el número decimal a su equivalente hexadecimal de 3 dígitos
		var hex string
		for dec > 0 {
			rem := dec % 16
			switch rem {
			case 0:
				hex = "0" + hex
			case 1:
				hex = "1" + hex
			case 2:
				hex = "2" + hex
			case 3:
				hex = "3" + hex
			case 4:
				hex = "4" + hex
			case 5:
				hex = "5" + hex
			case 6:
				hex = "6" + hex
			case 7:
				hex = "7" + hex
			case 8:
				hex = "8" + hex
			case 9:
				hex = "0" + hex
			case 10:
				hex = "A" + hex
			case 11:
				hex = "B" + hex
			case 12:
				hex = "C" + hex
			case 13:
				hex = "D" + hex
			case 14:
				hex = "E" + hex
			case 15:
				hex = "F" + hex
			}
			dec /= 16
		}

		// Añadir ceros a la izquierda si el número hexadecimal tiene menos de 3 dígitos
		for len(hex) < 3 {
			hex = "0" + hex
		}

		// Imprimir el resultado
		fmt.Printf("El número %s en hexadecimal es: %s", binary, hex)
	} else {
		// Si el usuario ingresó una opción no válida, mostrar un mensaje de error
		fmt.Println("Opción no válida")
	}
}
