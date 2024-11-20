package main

import (
	"fmt"
)

const unitario_luz = 100
const unitario_agua = 10000
const unitario_gas = 5000

type sujetos []string

var usuario = sujetos{"juan", "camilo", "lucho", "diana", "rafaela"}

type factura struct {
	estado  bool
	luz     float32
	agua    float32
	gas     float32
	persona string
	// saldo   float32
}

func (valor *factura) Consumir(luz float32, agua float32, gas float32) {
	valor.luz += luz
	valor.agua += agua
	valor.gas += gas
}

// func (valor *factura) pagar(pago float32) {
// 	valor.saldo -= pago
// }

func main() {

	var facturas []factura

	for _, i := range usuario {
		factura := factura{persona: i}

		facturas = append(facturas, factura)
	}

	//limpiar terminal
	fmt.Print("\033[H\033[2J")
	// inicializacion de variables
	hoja_factura := facturas[0]
	hoja_factura.Consumir(250, 3, 5)

	// ejecuacion de codigo
	fmt.Printf("Bienvenido Factura EPM\n\n")
	fmt.Printf("Factura de un usuario\n")

	fmt.Println("Usuario: ", hoja_factura.persona, "activa: ", hoja_factura.estado)
	fmt.Printf("variable     consumo        precio U.         subtotal\n")

	subtotal_luz := (hoja_factura.luz * unitario_luz)
	subtotal_agua := (hoja_factura.agua * unitario_agua)
	subtotal_gas := (hoja_factura.gas * unitario_gas)
	total := subtotal_luz + subtotal_agua + subtotal_gas

	fmt.Printf("luz :        %03.0f (Kwh)      %05d (Kwh/$)       %05.0f ($)\n", hoja_factura.luz, unitario_luz, subtotal_luz)
	fmt.Printf("agua:        %03.0f (Kwh)      %05d (m^3/$)       %05.0f ($)\n", hoja_factura.agua, unitario_agua, subtotal_agua)
	fmt.Printf("gas :        %03.0f (Kwh)      %05d (m^3/$)       %05.0f ($)\n", hoja_factura.gas, unitario_gas, subtotal_gas)
	fmt.Printf("\ntotal: %v $", total)
}
