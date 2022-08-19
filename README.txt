##Tabla de contenidos

1.[Informacion General](#Informacion-general)

2.[Expliacion detallada](#Explicacion-detallada)

3.[Creditos](#Creditos)

***
##Informacion General
***
El archivo main.go es una priority Queue en la cual como ejemoplo creamos una lista de elementos los cuales basan
Su prioridad en los dias restante ante de que expire, entre menos dias le restan mayor su priorida al salir de la Queue.

***
##Expliacion detallada
***

Importamos 2 paquetes: 
-FMT: Para utilizado en la prueba al sacar lso objetos en base a la prioridad.
-Container/Heap:  El cual es requerido para poder crear la Priority Queue.

El proyecto se compone de los siguiente elementos:
-Estructura tipo Item
-Arreglo tipo Priority Queue
-Funcion Len
-Funcion Less
-Funcion Pop
-Funcion Push
-Funcion Swap
-Funcion main
 
 Estructura tipo Item:

 Contiene los productos los cuales seran introducidos dentro de la priority queue. contineen el nombre, los dias restantes 
 antes de expirar y el componente index requerido por el heap.

 Arreglo tipo Priority

 Es nuestra priority queue en si. Se compone de un array de elementos tipo Item al cual le agregaremos las funciones de 
 nuestra priority queue.

 Funcion Len

 Nos devuelve el tamaño de la cola.

 Funcion Less

 Es la funcion en la que define que tendra prioridad, en este al menor ser los dias restantes antes de expirar mayor prioridad tendra.

 Funcion Pop

 Elimina elementos de la priority queue.

 Funcion Push

 Agrega elementos a la priority queue.

 Funcion Swap

 Intercambia elementos dentro de la priority queue.

 Funcion main:

 Nuestra funcion principal donde se ejecuta la priority queue.

 ***
 ##Creditos
 ***

 Codigo creado por Jeremmy Figueroa Aguilera para la clase de Algoritmos y Estructuras de datos 
 en Universidad Nacional Autonoma de Honduras.

 Agradecimientos espciales para siguienes enlaces de apoyo que fueron utilizados para crear este codigo.
 -https://www.ionos.es/digitalguide/paginas-web/desarrollo-web/archivo-readme/
 -https://www.youtube.com/watch?v=wptevk0bshY
 -https://www.youtube.com/watch?v=8Wuzvdvi7JI
 -https://pkg.go.dev/container/heap

