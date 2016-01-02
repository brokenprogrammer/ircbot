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

func evalCommand(bot *watcher.Watcher, conn net.Conn, message string) {
	newMess := formatter.ExtractMessage(message)
	fs := token.NewFileSet()
	tr, _ := parser.ParseExpr(newMess[6:])

	var expression []string
	//var expressionString string
	var Evalvalues []interface{}
	var operands []interface{}
	ast.Print(fs, tr)
	//Gets the different values and what position they are at. Find a way to find the operators as well.
	ast.Inspect(tr, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
			expression = append(expression, s)
			Evalvalues = append(Evalvalues, s)
		case *ast.Ident:
			s = x.Name
			expression = append(expression, s)
			Evalvalues = append(Evalvalues, s)
		case *ast.BinaryExpr:
			s = x.Op.String()
			expression = append(expression, s)
			operands = append(operands, s)
		}
		if s != "" {
			fmt.Print("Pos: ", n.Pos(), " Value: ", s)
		}
		return true
	})

	useFloat, useString := checkType(Evalvalues)
	var floatSum float64
	var stringSum string

	fmt.Println("useFloat: ", useFloat, " useString: ", useString)

	if useString == true {
		for i, _ := range Evalvalues {
			if i != 0 {
				operator := operands[i-1]
				switch operator {
				case "+":
					stringSum += Evalvalues[i].(string)
				}
			} else {
				stringSum += Evalvalues[i].(string)
			}
		}
	} else if useFloat == true {
		for i, _ := range Evalvalues {
			if i != 0 {
				operator := operands[i-1]
				switch operator {
				case "+":
					tempFloat, _ := strconv.ParseFloat(Evalvalues[i].(string), 64)
					floatSum += tempFloat
				}
			} else {
				tempFloat, _ := strconv.ParseFloat(Evalvalues[i].(string), 64)
				floatSum += tempFloat
			}
		}
	}

	if useString == true {
		conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + newMess[6:len(newMess)-2] + ": " + stringSum + " \r\n"))
	} else if useFloat == true {
		conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + newMess[6:len(newMess)-2] + ": " + strconv.FormatFloat(floatSum, 'f', -1, 64) + " \r\n"))
	}
}

func testReturnNum(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func checkType(i []interface{}) (bool, bool) {
	useFloat := false
	useString := false

	for _, val := range i {
		newVal, err := strconv.ParseFloat(val.(string), 64)

		if err != nil {
			useFloat = false
			useString = true
		} else {
			useFloat = true
		}

		fmt.Println(newVal, " :NewVal")
	}

	return useFloat, useString
}
