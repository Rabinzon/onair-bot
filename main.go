package main

import (
	"./config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

const (
	WAIT        = "WAIT"
	DELAY       = "DELAY"
	TIME_FORMAT = "15"
	ONLINE_TIME = "09"
)

var (
	state      = WAIT
	locMsc, _  = time.LoadLocation(config.LocMsc)
	commandReg = regexp.MustCompile(config.Command)
)

func reportErr(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusExpectationFailed)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
}

func jsonResponse(w http.ResponseWriter, code int, data interface{}) {
	out, err := json.Marshal(data)
	if err != nil {
		reportErr(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", out)
}

func isTime(day time.Weekday, hour string) bool {
	return day == time.Monday && hour == ONLINE_TIME
}

func wait() {
	timer := time.NewTimer(time.Hour)
	<-timer.C
	state = WAIT
	return
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	ev := config.Event{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reportErr(nil, w)
		return
	}

	if err = json.Unmarshal(body, &ev); err != nil {
		reportErr(nil, w)
		return
	}

	if commandReg.MatchString(ev.Text) {
		jsonResponse(w, http.StatusCreated, config.ComRes)
		return
	}

	currentTime := time.Now().In(locMsc)
	weekday := currentTime.Weekday()
	hour := currentTime.Format(TIME_FORMAT)

	switch {
	case state == WAIT:
		if isTime(weekday, hour) {
			state = DELAY
			jsonResponse(w, http.StatusCreated, config.BotRes)
		} else {
			reportErr(nil, w)
		}
	case state == DELAY:
		go wait()
		reportErr(nil, w)
	}
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusCreated, config.BotInfo)
}

func main() {
	log.Println(config.BotName)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/event", eventHandler)
	log.Fatal(http.ListenAndServe(config.Port, nil))
}
