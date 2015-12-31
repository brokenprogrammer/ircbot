package config

type Config struct {
	Server   string //Server adress
	Port     string //Server Port
	Nick     string //Watchers Nickname
	Channel  string //Channel to join
	DBDriver string //The database driver used.
	DBPath   string //The path for our sqlite3 database.
}

//The Config structure used to return values we are intrested in from the Config
var MainCFG Config

//Function to initialize the configurations. Must be a better way to do this?
//TODO: Return a pointer to the MainCFG ?
func InitCfg() {
	MainCFG = Config{
		Server:   "irc.freenode.net", //Server our IRC uses
		Port:     "6667",             //Port for IRC
		Nick:     "BrokenBot",        //Name of our bot
		Channel:  "#gobotter",        //Channel bot is connecting to
		DBDriver: "sqlite3",          //The driver used by our database
		DBPath:   "./db/ircbot.db",   //The path for our sqlite3 database
	}
}
