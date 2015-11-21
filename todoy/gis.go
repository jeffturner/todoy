package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// GoogleImageSearch models the JSON response of a google image search
type GoogleImageSearch struct {
	ResponseData struct {
		Results []struct {
			ResultClass         string `json:"GsearchResultClass"`
			Width               string `json:"width"`
			Height              string `json:"height"`
			ImageID             string `json:"imageId"`
			TBWidth             string `json:"tbWidth"`
			TBHeight            string `json:"tbHeight"`
			UnescapedURL        string `json:"unescapedUrl"`
			URL                 string `json:"url"`
			VisibleURL          string `json:"visibleUrl"`
			Title               string `json:"title"`
			TitleNoFormatting   string `json:"titleNoFormatting"`
			OrginalContextURL   string `json:"originalContextUrl"`
			Content             string `json:"content"`
			ContentNoFormatting string `json:"contentNoFormatting"`
			TBURL               string `json:"tbUrl"`
		} `json:"results"`
	} `json:"responseData"`
}

// Search calls out to GIS and returns somewhat random results based on params
func (gis *GoogleImageSearch) Search(userip string, term string) error {
	url := "https://ajax.googleapis.com/ajax/services/search/images?v=1.0" +
		"&userip=" + userip +
		"&rsz=8" +
		"&start=1" +
		"&as_filetype=jpg" +
		"&q=" + term

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Referer", "TODOY")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(gis)
	if err != nil {
		return err
	}
	return nil
}
