package server

import (
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"server/log"
	"server/settings"
	"server/torr/utils"
	"server/web"
)

// returns the port of the server or -1 if failed
// specify port 0 if you want a random port
func Start(pathdb string, port int, roSets, searchWA bool) int {
	settings.Path = pathdb
	if !settings.InitSets(roSets, searchWA) {
		return -1
	}
	if roSets {
		log.TLogln("Enabled Read-only DB mode!")
	}

	// http checks
	log.TLogln("Check web port", port)

	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if l != nil {
		l.Close()
	}
	if err != nil {
		log.TLogln("Port", port, "already in use! Please set different sslport for HTTP. Abort")
		return -1
		//os.Exit(1)
	}
	realPort := l.Addr().(*net.TCPAddr).Port
	// remove old disk caches
	go cleanCache()
	// set settings http Start web server.
	settings.Port = strconv.Itoa(realPort)
	if web.Start() {
		return realPort
	} else {
		return -1
	}
}

func cleanCache() {
	if !settings.BTsets.UseDisk || settings.BTsets.TorrentsSavePath == "/" || settings.BTsets.TorrentsSavePath == "" {
		return
	}

	dirs, err := os.ReadDir(settings.BTsets.TorrentsSavePath)
	if err != nil {
		return
	}

	torrs := settings.ListTorrent()

	log.TLogln("Remove unused cache in dir:", settings.BTsets.TorrentsSavePath)
	for _, d := range dirs {
		if len(d.Name()) != 40 {
			// Not a hash
			continue
		}

		if !settings.BTsets.RemoveCacheOnDrop {
			for _, t := range torrs {
				if d.IsDir() && d.Name() != t.InfoHash.HexString() {
					log.TLogln("Remove unused cache:", d.Name())
					removeAllFiles(filepath.Join(settings.BTsets.TorrentsSavePath, d.Name()))
					break
				}
			}
		} else {
			if d.IsDir() {
				log.TLogln("Remove unused cache:", d.Name())
				removeAllFiles(filepath.Join(settings.BTsets.TorrentsSavePath, d.Name()))
			}
		}
	}
}

func removeAllFiles(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, f := range files {
		name := filepath.Join(path, f.Name())
		os.Remove(name)
	}
	os.Remove(path)
}

func WaitServer() string {
	err := web.Wait()
	if err != nil {
		return err.Error()
	}
	return ""
}

func Stop() {
	web.Stop()
	settings.CloseDB()
}

func AddTrackers(trackers string) {
	tracks := strings.Split(trackers, ",\n")
	utils.SetDefTrackers(tracks)
}
