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
	go additin()
	http.HandleFunc("/", trash.Home)
	http.HandleFunc("/person", trash.Person)
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.ListenAndServe(":7777", nil)

}

// ////////////////////////////////////////////////////////////////////////////////////////
func randomint() int64 {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 10 and 100
	randomNumber := rand.Intn(91) + 10
	return int64(randomNumber)
}

// // ///////////////////////////////////////////////////////////////////////////////////////
func cur_time() string {
	// Get the current time
	currentTime := time.Now()

	// Convert the time to a string using a layout
	timeString := currentTime.Format("2006-01-02 15:04:05")
	return timeString
}

///////////////////////////////////////////////////////////////////////////////////////////

func additin() {
	// Create a ticker that triggers every 15 seconds
	ticker := time.NewTicker(15 * time.Second)

	// Start a goroutine to run the function repeatedly
	go func() {
		for {
			// Call your function here
			adc()

			// Wait for the ticker to trigger
			<-ticker.C
		}
	}()

	// Sleep for a while to allow the function to run
	time.Sleep(1 * time.Hour)

	// Stop the ticker when done
	ticker.Stop()
}

func adc() {
	post := trash.Post{
		GasValue:    randomint(),
		Humidity:    25, //
		Pressure:    randomint(),
		Temperature: 16, //
		UserId:      randomint() % 6,
		WaterValue:  0,
		Time:        cur_time(),
	}
	repoo.Add(&post)
}
