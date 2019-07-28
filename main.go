package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"time"
)

func main() {
	response, _ := http.Get("https://alistairfink.com/alistairfink/api/abouts")

	if response.StatusCode != 200 {
		layout := "January 2, 2006 at 3:04pm (ET)"
		torontoTime, _ := time.LoadLocation("America/Toronto")
		from := "alistairfinkraspberrypi@gmail.com"
		pass := ""
		to := "alistairfink@gmail.com"
		body := "Alistairfink.com responded with error code " + strconv.Itoa(response.StatusCode) + ".\nTimestamp: " + time.Now().In(torontoTime).Format(layout) + "."

		msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Alistairfink.com Webserver Responded With Error!\n\n" +
			body

		err := smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))

		if err != nil {
			log.Printf("smtp error: %s", err)
			return
		}
	}

}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	response, _ := http.Get("https://alistairfink.com/alistairfink/api/abouts")

	if response.StatusCode != 200 {

	}

	fmt.Fprint(w, response.StatusCode == 404)
}
