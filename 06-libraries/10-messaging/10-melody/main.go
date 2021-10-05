package main

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type GopherInfo struct {
	ID, X, Y string
}

func main() {
	router := gin.Default()
	mldy := melody.New()
	gophers := make(map[*melody.Session]*GopherInfo)
	lock := new(sync.Mutex)
	counter := 0

	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	router.GET("/ws", func(c *gin.Context) {
		mldy.HandleRequest(c.Writer, c.Request)
	})

	mldy.HandleConnect(func(s *melody.Session) {
		lock.Lock()
		for _, info := range gophers {
			s.Write([]byte("set " + info.ID + " " + info.X + " " + info.Y))
		}
		gophers[s] = &GopherInfo{strconv.Itoa(counter), "0", "0"}
		s.Write([]byte("iam " + gophers[s].ID))
		counter += 1
		lock.Unlock()
	})

	mldy.HandleDisconnect(func(s *melody.Session) {
		lock.Lock()
		mldy.BroadcastOthers([]byte("dis "+gophers[s].ID), s)
		delete(gophers, s)
		lock.Unlock()
	})

	mldy.HandleMessage(func(s *melody.Session, msg []byte) {
		p := strings.Split(string(msg), " ")
		lock.Lock()
		info := gophers[s]
		if len(p) == 2 {
			info.X = p[0]
			info.Y = p[1]
			mldy.Broadcast([]byte("set "+info.ID+" "+info.X+" "+info.Y))
		}
		lock.Unlock()
	})

	router.Run(":5000")
}