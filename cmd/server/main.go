package main

import "fmt"
//App - the struct which contains things like pointers
//to database connections 
type App struct{

}

//Run sets up our Applications
func (app *App)Run()error{
  fmt.Println("Setting Up Our App")
  return nil
}

func main(){
	fmt.Println("GO Rest Api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}