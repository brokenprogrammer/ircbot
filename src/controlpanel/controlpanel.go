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
	var expressionString string

	ast.Print(fs, tr)
	//Gets the different values and what position they are at. Find a way to find the operators as well.
	ast.Inspect(tr, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
			expression = append(expression, s)
		case *ast.Ident:
			s = x.Name
			expression = append(expression, s)
		case *ast.BinaryExpr:
			s = x.Op.String()
			expression = append(expression, s)
		}
		if s != "" {
			fmt.Print("Pos: ", n.Pos(), " Value: ", s)
		}
		return true
	})

	for _, v := range expression {
		expressionString += v + " "
	}

	//TODO: CHECK IF THE VALUES ARE BOOLEANS / STRINGS FLOATS OR INTS
	//TODO: MAKE IT SO MULTIPLE EXPRESSIONS ARE POSSIBLE

	var astring string

	if len(expression) >= 3 {
		var1 := expression[1]
		var2 := expression[2]
		solved := expression[1] + expression[2]
		switch expression[0] {
		case "+":
			solved = var1 + var2
			astring = solved
		}
	} else {
		solved := expression[0]
		astring = solved
	}
	fmt.Println(astring)
	theReal := testReturnNum(expression[1]) + testReturnNum(expression[2])

	conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + newMess[6:len(newMess)-2] + ": " + strconv.Itoa(theReal) + " \r\n"))
}

func testReturnNum(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}
