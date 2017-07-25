package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/tsingson/go-sessions"
	"github.com/tsingson/go-sessions/sessiondb/goredis"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	db   *goredis.Database
	sess sessions.Sessions

	mySessionsConfig = sessions.Config{Cookie: "mysessioncookieid",
		Expires:                     time.Duration(2) * time.Hour,
		DisableSubdomainPersistence: false,
	}

	mySessions = sessions.New(mySessionsConfig)
)

func init() {
	db = goredis.New(goredis.Config{})
	pong, err := db.Redis.Ping().Result()
	fmt.Println("*****************************************")
	fmt.Println(pong, err)
	fmt.Println("*****************************************")
	mySessions.UseDatabase(db)
	spew.Dump(mySessions)

}

// set some values to the session
func setHandler(reqCtx *fasthttp.RequestCtx) {
	values := map[string]interface{}{
		"Name":   "go-sessions",
		"Days":   "1",
		"Secret": "dsads£2132215£%%Ssdsa",
	}

	sess := mySessions.StartFasthttp(reqCtx) // init the session

	// sessions.StartFasthttp returns the, same, Session interface we saw before too
	//sess.UseDatabase(db)

	for k, v := range values {
		sess.Set(k, v) // fill session, set each of the key-value pair
	}
	reqCtx.WriteString("Session saved, go to /get to view the results")
}

// get the values from the session
func getHandler(reqCtx *fasthttp.RequestCtx) {
	sess := mySessions.StartFasthttp(reqCtx) // init the session
	//sess.UseDatabase(db)

	sessValues := sess.GetAll() // get all values from this session

	reqCtx.WriteString(fmt.Sprintf("%#v", sessValues))
}

// clear all values from the session
func clearHandler(reqCtx *fasthttp.RequestCtx) {
	sess := mySessions.StartFasthttp(reqCtx) // init the session
	//sess.UseDatabase(db)

	sess.Clear()
}

// destroys the session, clears the values and removes the server-side entry and client-side sessionid cookie
func destroyHandler(reqCtx *fasthttp.RequestCtx) {
	mySessions.DestroyFasthttp(reqCtx)
}

func router() {
	fmt.Println("Open a browser tab and navigate to the localhost:8080/set")
	fasthttp.ListenAndServe(":8080", func(reqCtx *fasthttp.RequestCtx) {
		path := string(reqCtx.Path())

		if path == "/set" {
			setHandler(reqCtx)
		} else if path == "/get" {
			getHandler(reqCtx)
		} else if path == "/clear" {
			clearHandler(reqCtx)
		} else if path == "/destroy" {
			destroyHandler(reqCtx)
		} else {
			reqCtx.WriteString("Please navigate to /set or /get or /clear or /destroy")
		}
	})
}
