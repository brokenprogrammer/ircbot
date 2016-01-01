# ircbot
An feature rich irc bot capable of watching over channels and logging messages.

##Features
* Remember users and their messages using Sqlite3
* Listens to commands from both terminal and irc chat
* Print out logs of user messages

##Installing
To set up your own developing envornment to run the ircbot is easy!
Clone the repo by typing this into git:
```
	git clone https://github.com/brokenprogrammer/ircbot.git
```
Or simply just download the zip archive and extract it to where you want your workspace.

If you want the project to be in a separate directory go requires you to export your gopath variable.
This is how i do it to set my Gopath to my ircbot directory which is in my home directory on Ubuntu:
```
	export GOPATH=$HOME/ircbot
```
When you have the bot in your directory you have to set up the config file to work your way, The configurations
file is found under (yourpath/src/config/config.go)

Editing the config is pretty straight forward. Right now its only tested with sqlite3 as database driver. The two things you probably would like to change is "Channel" - The irc channel the bot connects to and "Nick" - The username the bot will connect with.

###Third Party Packages
Sqlite3 for go
```
	go get github.com/mattn/go-sqlite3
	go install github.com/mattn/go-sqlite3
```
##Running ircbot
Running the bot is realy easy, just run it using go by entering the same directory as the main.go file (yourpath/src)
then type:
```
	go run main.go
```
###Console Commands
Commands you can type into the terminal while the bot is running to make the bot do tasks:
```	
	Hello - Makes the bot print a hello message into the irc chat.
	Block Username - Replace Username with the user you wish to block, blocked users cannot use the chat commands.
	UnBlock Username - Replace Username with the user you wish to unblock.
```
###Chat Commands
Commands users can type in the irc chat to make the bot do specific tasks:
```	
	!help - Displays the chat commands the bot reacts to.
	!status - Displays the time the bot has been running as well as the ammount of tracked users and messages.
```