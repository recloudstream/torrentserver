package torrServer

import (
	server "server"
)

func StartTorrentServer(pathdb string, port int) int {
	return server.Start(pathdb, port, false, false)
}

func WaitTorrentServer() {
	server.WaitServer()
}

func StopTorrentServer() {
	server.Stop()
}

func AddTrackers(trackers string) {
	server.AddTrackers(trackers)
}
