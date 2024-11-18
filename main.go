package main

func main() {
	tds := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.load(&tds)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&tds)

	storage.save(tds)
}
