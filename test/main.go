package main

import (
	"hackathon/trash"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	repoo trash.PostRepository = trash.NewPostRepository()
)

func main() {
	path := `D:\Adil\go\test\hackathon-1018f-firebase-adminsdk-tfuqo-8825075380.json`
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", path)

	http.HandleFunc("/", trash.Home)
	http.HandleFunc("/person", trash.Person)
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.ListenAndServe(":7777", nil)

	// post := trash.Post{
	// 	GasValue:    randomint(),
	// 	Humidity:    float64(randomint()),
	// 	Pressure:    randomint(),
	// 	Temperature: float64(randomint()),
	// 	UserId:      randomint() % 6,
	// 	WaterValue:  randomint(),
	// 	Time:        cur_time(),
	// }
	// repoo.Add(&post)
}

// ////////////////////////////////////////////////////////////////////////////////////////
func randomint() int64 {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 10 and 100
	randomNumber := rand.Intn(91) + 10
	return int64(randomNumber)
}

// ///////////////////////////////////////////////////////////////////////////////////////
func cur_time() string {
	// Get the current time
	currentTime := time.Now()

	// Convert the time to a string using a layout
	timeString := currentTime.Format("2006-01-02 15:04:05")
	return timeString
}

///////////////////////////////////////////////////////////////////////////////////////////
