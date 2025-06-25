package main

import (
	"fmt"
)

// interfaces
type consumo interface {
	Consumir(float32, float32, float32)
}

// constantes
const unitario_luz = 100
const unitario_agua = 10000
const unitario_gas = 5000

// creacion de datos
type usuario_general []string
type consumo_general [][]float32
type activo_general []bool

var usuarios = usuario_general{"juan", "camilo", "lucho", "diana", "rafaela"}
var consumos = consumo_general{{250, 3, 5}, {30, 2, 25}, {240, 8, 1}, {70, 8, 9}, {100, 2, 3}}
var activos = activo_general{true, false, false, true, true}

// creacion de la estructrura factura
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

func imprimir_factura(hoja_factura factura) {
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
	fmt.Printf("\ntotal: %v $\n", total)
}

func main() {

	var facturas []factura
	var consumo_interfaz consumo
	var interfaz_factura_aparte consumo = &factura{persona: "anonimo"}
	interfaz_factura_aparte.Consumir(10, 20, 30)

	// inicializando instancia de clases
	for i := 0; i < len(usuarios); i++ {
		factura := factura{persona: usuarios[i], luz: 0, agua: 0, gas: 0, estado: activos[i]}
		consumo_interfaz = &factura
		consumo_interfaz.Consumir(consumos[i][0], consumos[i][1], consumos[i][2])
		facturas = append(facturas, factura)
	}

	//limpiar terminal
	fmt.Print("\033[H\033[2J")

	//fmt.Println("3333", interfaz_factura_aparte, consumo_interfaz)
	// ejecuacion de codigo
	fmt.Printf("Bienvenido Factura EPM\n\n")

	imprimir_factura(facturas[1])

}
