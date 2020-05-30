package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarAlista(tl *task) {
	t.tasks = append(t.tasks, tl) // append  va a tomar como parametro un slice y un elemeto que vamos a agregar al slice
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descripcion", tarea.descripcion)
	}

}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Descripcion", tarea.descripcion)
		}

	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() { // usamos el operador asterisco para completar la funcion del apuntador para actualizar el valor como tal
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizacionNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{ // usamos ampersan como apuntador del struct
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go de platzi en esta semana",
	}

	t2 := &task{ // usamos ampersan como apuntador del struct
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python de platzi en esta semana",
	}

	t3 := &task{ // usamos ampersan como apuntador del struct
		nombre:      "Completar mi curso de Flutter",
		descripcion: "Completar mi curso de Flutter de platzi en esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	lista.agregarAlista(t3)
	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	fmt.Println("Tareas Completadas")
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)
	mapaTareas["Briam"] = lista

	t4 := &task{ // usamos ampersan como apuntador del struct
		nombre:      "Completar mi curso de Deno",
		descripcion: "Completar mi curso de Deno de platzi en esta semana",
	}

	t5 := &task{ // usamos ampersan como apuntador del struct
		nombre:      "Completar mi curso de Typescript",
		descripcion: "Completar mi curso de Typescript de platzi en esta semana",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Cristhofer"] = lista2

	fmt.Println("Tareas de Briam")
	mapaTareas["Briam"].imprimirLista()
	fmt.Println("Tareas de Cristhofer")
	mapaTareas["Cristhofer"].imprimirLista()

	// con esta linea agregamos una tarea a la lista
	//for i := 0; i < len(lista.tasks); i++ {
	//	fmt.Println("Index", i, "nombre", lista.tasks[i].nombre)
	//}

	//for index, tarea := range lista.tasks { // misma accion que con el for pero usando range
	//	fmt.Println("index", index, "nombre", tarea.nombre)

	//}

	//for i := 0; i < 10; i++ { // usando for con break para que finalice el for donde lo indique el break
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println(i)

	//}

	//for i := 0; i < 10; i++ { // usando for con break para que finalice el for donde lo indique el break
	//	if i == 5 {
	//		continue
	//	}
	//	fmt.Println(i)

	//}
}
