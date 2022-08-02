package conf

import "time"

var (
	LenStackBuf = 4096

	// log
	LogLevel string
	LogPath  string
	LogFlag  int

	// console
	ConsolePort   int
	ConsolePrompt string = "Leaf# "
	ProfilePath   string

	// cluster
	ServerName        string
	ListenAddr        string
	ConnAddrs         map[string]string
	PendingWriteNum   int
	HeartBeatInterval int

	// client
	HeartBeatTime time.Duration
)
