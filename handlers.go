package main

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"strings"
	"image"
	"os"
	"github.com/asmcos/requests"
	"image/png"
	"log"
	"net/http"
	"io"
	"path"
	"github.com/google/uuid"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	msg := "This bot will predict the objects within an image.\nCommands:\n1. Use /predict to predict the objects in the image of given url."
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle the /predict command here
func (a *application) predictHandler(m *tbot.Message) {

	url := strings.TrimPrefix(m.Text, "/predict ")

	filename := path.Base(url)
	fmt.Println(filename)

    	defer func() {
        	if r := recover(); r != nil {
			a.client.SendMessage(m.Chat.ID, "Unsupported image type")
            		fmt.Println("Recovered: ", r)
        	}
    	}()
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Image is received successfully: " + filename)

	out, erro := os.Create(filename)

	if erro != nil {
		log.Println(erro)
		return
	}

	defer out.Close()
	io.Copy(out, resp.Body)

	data := requests.Files{
           	"file": filename,
        }

	resps, errors := requests.Post("https://wpir-dnjf-8439.herokuapp.com/predict", data)

	if errors != nil{
		log.Println(errors)
        	return
        }

	fmt.Println("Image is predicted successfully: " + filename)

	e := os.Remove(filename)
    	if e != nil {
        	log.Fatal(e)
		return
    	}

	fmt.Println("Raw image is removed successfully: " + filename)

	img_data := resps.Text()

	img, _, errr := image.Decode(strings.NewReader(img_data))
    	if errr != nil {
		a.client.SendMessage(m.Chat.ID, "Unsupported image type")
        	fmt.Println(errr)
		return
    	}

	filename = uuid.New().String() + ".png"
    	outs, _ := os.Create(filename)
    	defer outs.Close()

    	errs := png.Encode(outs, img)
    	if errs != nil {
		fmt.Println(errs)
		return
    	}

	fmt.Println("Image is ready to be sent: " + filename)

	_, errrs := a.client.SendPhotoFile(m.Chat.ID, filename, tbot.OptCaption("Predicted Result!"))
	if errrs != nil {
		fmt.Println(errrs)
		return
	}

	fmt.Println("Image is sent successfully: " + filename)

	ers := os.Remove(filename)
    	if ers != nil {
        	log.Fatal(ers)
		return
    	}

	fmt.Println("Process complete")
}
