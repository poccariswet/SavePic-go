package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func SavePic(filename string, text string, i int) error {
	res, err := http.Get(text)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	picture, err := os.Create(fmt.Sprintf(filename+"/%d.jpg", i))
	if err != nil {
		return err
	}

	defer picture.Close()
	io.Copy(picture, res.Body)

	return nil
}
