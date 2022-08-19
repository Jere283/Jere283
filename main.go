package main

import (
	"container/heap"
	"fmt"
)

//Definimos las estrcutura en la cual para el ejemplo se crearan productos que deberan tener una
//Fecha de vencicimiento que determina la prioridad
type Item struct {
	Nombre string
	Exp    int //esto indicara los dias restantes antes de que expire el producto.

	Index int // esto sera lo requerido por los heaps.
}

//Array de elementos tipo item al cual le agregaremos las funciones de nuestra priority Q
type PriorityQueue []*Item

//El tamano de nuestra priority Q
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Queremos que Pop nos dé el más bajo basado en el número de vencimiento como prioridad
	// Cuanto menor sea el no., mayor será la prioridad
	return pq[i].Exp < pq[j].Exp
}

//funcion pop pare remover el elemento de mayor prioridad.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

//insertamos los eleemntos dentro la Priority Q, basado en su valor.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

//funcion swap, encargada de intercambiar los contenido dentor de la mima.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func main() {
	//Primero creamos la lsita de objetos, EXP nos dira cuanto dias restantes tiene, por lo tanto entre menos dias tengan mayor sera la prioridad.
	listItems := []*Item{
		{Nombre: "Yogurt", Exp: 30},
		{Nombre: "Yogurt_Griego", Exp: 45},
		{Nombre: "Arroz", Exp: 100},
		{Nombre: "Manzana", Exp: 1},
		{Nombre: "Barra energetica", Exp: 2},
		{Nombre: "Lata de atun", Exp: 365},
	}
	//creamos la priority Q que contine obejtos tipo item.
	priorityQueue := make(PriorityQueue, len(listItems))

	//introducimos la lista de objetos dentro de la Priority Q.
	for i, item := range listItems {
		priorityQueue[i] = item
		priorityQueue[i].Index = i
	}

	//Inicializamos heap.
	heap.Init(&priorityQueue)

	// Imprimimos bajo el orden de prioridad basado en la fecha de expiracion.
	for priorityQueue.Len() > 0 {
		item := heap.Pop(&priorityQueue).(*Item)
		fmt.Printf("Nombre: %s, Dias restantes: %d\n", item.Nombre, item.Exp)
	}

}
