package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"log"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Now().In(loc)
	formatter:= fmt.Sprintf("%02d%02d%d",t.Day(), t.Month(), t.Year() )
	anandabajarformatter:= fmt.Sprintf("anandabajar/%s",formatter)
	telegraphformatter:= fmt.Sprintf("telegraph/%s",formatter)
	fmt.Println("Downloading todays paper:", formatter)
	// Create the folder
	_, err := os.Stat(anandabajarformatter)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(anandabajarformatter, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	_, err1 := os.Stat(telegraphformatter)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(telegraphformatter, 0755)
		if errDir != nil {
			log.Fatal(err1)
		}
	}
		

	for i:=1; i<= 14; i++ {
		anandabajarURL := fmt.Sprintf("https://epaper.anandabazar.com/epaperimages////%s////%s-md-hr-%dll.png", formatter, formatter, i)
		err:= DownloadFile(fmt.Sprintf("%s/anandabajar-%s-%02d.png",anandabajarformatter, formatter,i), anandabajarURL)  
		if err != nil {
		panic(err)
		}
		fmt.Println("Downloaded Anandabajar: " + anandabajarURL)
	}

	for i:=1; i<= 10; i++ {
		telegraphURL := fmt.Sprintf("https://epaper.telegraphindia.com/epaperimages////%s////%s-md-hr-%dll.png", formatter, formatter, i)
		err:= DownloadFile(fmt.Sprintf("%s/telegraph-%s-%02d.png",telegraphformatter, formatter,i), telegraphURL)  
		if err != nil {
		panic(err)
		}
		fmt.Println("Downloaded Telegraph: " + telegraphURL)
	}
	
}


// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
// https://golangcode.com/download-a-file-from-a-url/
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
} 

