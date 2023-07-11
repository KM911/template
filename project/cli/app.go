package cli

type Command struct {
	Name    string
	Aliases []string
	Usage   string
	Action  func()
	//Flags 	[]Flag
}

//type Flag struct {
//	Name     string
//	Aliases  []string
//}
type App struct {
	Name     string
	Commands []Command
	Action   func()
}