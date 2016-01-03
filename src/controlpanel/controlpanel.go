package controlpanel

import (
	"db"
	"fmt"
	"formatter"
	"go/ast"
	"go/parser"
	"go/token"
	"net"
	"strconv"
	"strings"
	"time"
	"watcher"
)

//Control Panel that watches input from the chat.
func ControlPanel(conn net.Conn, bot *watcher.Watcher, c chan string, DBConn *db.Crud) {
	for {

		select {
		//Incase we got a message through the channel
		case message := <-c:
			//Call the readMessage function that handles code depending on commands
			readMessage(message, bot, conn, DBConn)
		//Incase we didn't we timeout after 1 Second
		case <-time.After(time.Second * 1):
			//fmt.Println("Def") //Timeout
		}
	}
}

//readMessage reads messages user sent to the chat
//If it recognize a command it calls a function depending on command.
func readMessage(msg string, bot *watcher.Watcher, conn net.Conn, DBConn *db.Crud) {
	//Splits the string into a slice so we can use multiple parts of the message
	splitted := strings.Split(msg, " ")

	//Switch through the incoming message
	switch strings.TrimSpace(splitted[3]) {
	//Incase user wrote !help, Display the help message
	case ":!help":
		//Calling function displaying help message
		helpCommand(bot, conn)
	case ":!status":
		//Calling the function displaying the status message
		statusCommand(bot, conn, DBConn)
	case ":!eval":
		evalCommand(bot, conn, msg)
	default:
		fmt.Println(msg)
	}
}

//helpCommand is a command function called by the readMessage function incase a
//user requested to display the help message.
func helpCommand(bot *watcher.Watcher, conn net.Conn) {
	conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + "BrokenBot Commands: !help - Display help message, !status - Bot Status \r\n"))
}

//statusCommand is a command function called by readMessage function incase a user
//Requested to view the status of the application.
func statusCommand(bot *watcher.Watcher, conn net.Conn, DBConn *db.Crud) {
	conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + "BrokenBot Status: Uptime: " + time.Since(bot.RanFor).String() + " Tracking: " + strconv.Itoa(db.GetAmmountOfUsers(DBConn)) + " users and tracking: " + strconv.Itoa(db.GetAmmountOfMessages(DBConn)) + " messages. \r\n"))
}

//evalCommand is a command function called by readMessage incase a user
//Requested a statement to be evaluated
func evalCommand(bot *watcher.Watcher, conn net.Conn, message string) {
	//Use the formatter library to extract the message from the user input
	newMess := formatter.ExtractMessage(message)

	//Using token and ast librariies to parse the string reading one value at the time.
	fs := token.NewFileSet()
	tr, _ := parser.ParseExpr(newMess[6:])

	//string to hold all the expressions
	var expression []string

	//The used interface to hold all the values typed in the eval statement
	var Evalvalues []interface{}
	//The used interface to hold all the operators typed in the eval statement
	var operands []interface{}
	ast.Print(fs, tr)
	//Gets the different values and what position they are at. Find a way to find the operators as well.
	ast.Inspect(tr, func(n ast.Node) bool {
		//String variable that every found identifier and litteral will be put in.
		var s string

		//Type switch to determit what type the current node is
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value //Adding the value into our string s
			expression = append(expression, s)
			//Evalvalues used to store all the given values, Strings Integers etc.
			Evalvalues = append(Evalvalues, s)
		case *ast.Ident:
			s = x.Name //Adding the value into our string s
			expression = append(expression, s)
			//Evalvalues used to store all the given values, Strings Integers etc.
			Evalvalues = append(Evalvalues, s)
		case *ast.BinaryExpr:
			s = x.Op.String() //Adding the operator by type assertion to our string
			expression = append(expression, s)
			//Storing all our operators in a separate interface slice.
			operands = append(operands, s)
		}
		if s != "" {
			//Printing out our created string
			fmt.Print("Pos: ", n.Pos(), " Value: ", s)
		}
		return true
	})

	//Check if float or string should be used
	useFloat, useString := checkType(Evalvalues)
	var floatSum float64 //variable to be used if float returns true
	var stringSum string //variable to be used if string returns true

	fmt.Println("useFloat: ", useFloat, " useString: ", useString)

	//Actions depending on if we use string or float
	if useString == true {
		//Looping through all values
		for i, _ := range Evalvalues {
			//If its not the first loop
			if i != 0 {
				operator := operands[i-1]
				//Check which operator is used
				switch operator {
				case "+":
					stringSum += Evalvalues[i].(string)
				}
			} else { //If its the first loop then just add the first value
				stringSum += Evalvalues[i].(string)
			}
		}
	} else if useFloat == true {
		//Looping through all the values
		for i, _ := range Evalvalues {
			//If its no the first loop
			if i != 0 {
				operator := operands[i-1]
				//Check which operatis is used
				switch operator {
				case "+":
					tempFloat, _ := strconv.ParseFloat(Evalvalues[i].(string), 64)
					floatSum += tempFloat
				case "-":
					tempFloat, _ := strconv.ParseFloat(Evalvalues[i].(string), 64)
					floatSum -= tempFloat
				case "*":
					tempFloat, _ := strconv.ParseFloat(Evalvalues[i].(string), 64)
					floatSum *= tempFloat
				}
			} else { //If its the first loop then add the first value
				tempFloat, _ := strconv.ParseFloat(Evalvalues[i].(string), 64)
				floatSum += tempFloat
			}
		}
	}

	//Write to the chat depending on if its a string or float used.
	if useString == true {
		conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + newMess[6:len(newMess)-2] + ": " + stringSum + " \r\n"))
	} else if useFloat == true {
		conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + newMess[6:len(newMess)-2] + ": " + strconv.FormatFloat(floatSum, 'f', -1, 64) + " \r\n"))
	}
}

//Test function to switch strings into integers
func testReturnNum(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

//Check which types are included, returns 2 booleans, one for string and one for float
func checkType(i []interface{}) (bool, bool) {
	//Creating our two return values, either float or string
	useFloat := false
	useString := false

	//Ranging through the interface slice
	for _, val := range i {
		//Attempting to parse the current value to a float
		newVal, err := strconv.ParseFloat(val.(string), 64)

		//If the parse to float fails we keep the value as a string and return useString as true
		if err != nil {
			useFloat = false
			useString = true
		} else { //Otherwise we will be using floats
			useFloat = true
		}

		fmt.Println(newVal, " :NewVal")
	}

	return useFloat, useString
}
