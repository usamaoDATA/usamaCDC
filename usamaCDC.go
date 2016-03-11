package greetings 

//GreetingsString is used AS IS
var GreetingsString = "HI"

//PrintGreetings IS A THING
func PrintGreetings (name string) string {	
	return GreetingsString + "-" + name; 
}
