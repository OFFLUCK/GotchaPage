package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Info struct {
	Nickname string   `json:"nickname"`
	Parsers  []string `json:"parsers"`
}

func requests(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: // Getting info about all pages (for checkboxes).
		{
			links, err := getPages()
			if err != nil {
				log.Fatalln("Get pages info")
				return
			}

<<<<<<< HEAD
			_, err = rw.Write(links)
			if err != nil {
				log.Fatalln("Write error")
			}
			log.Println("Successfully sent data!")
=======
	var formPage *template.Template = template.Must(template.ParseFiles(conf.Paths.MainTemplate))
	pageInfo := addCheckBoxesAndNickname(NewRequesterContainer(""), "")
	err := formPage.Execute(rw, pageInfo)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

// Adds check boxes and nickname value in text input field on page.
func addCheckBoxesAndNickname(container *RequesterContainer, nickname string) (pageInfo *HTMLInfo) {
	pageInfo = new(HTMLInfo)
	var isChecked string
	for _, page := range Pages {
		isChecked = ""
		if container.Requesters[page.ID].IsSelected() {
			isChecked = "checked"
>>>>>>> master
		}

	case http.MethodPost: // Sending info about requested users.
		{
			info, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatalln("Reading error")
				return
			}

			ans, err := getUsers(info)

			if err != nil {
				log.Fatalln("Getting answers error")
				return
			}

			_, err = rw.Write(ans)
			if err != nil {
				log.Fatalln("Write error")
			}

			log.Println("Successfully got data!")
		}
	default:
	}
}

<<<<<<< HEAD
// Returns info about all pages available.
func getPages() (pages []byte, err error) {
	pages, err = json.Marshal(Pages)
=======
// Adds answers to page.
func addAnswers(rw http.ResponseWriter, r *http.Request) {
	var answerPage *template.Template = template.Must(template.ParseFiles(conf.Paths.MainTemplate))

	err := r.ParseForm()
>>>>>>> master
	if err != nil {
		return nil, err
	}
	return pages, nil
}

// Checking which checkboxes are set.
func setUsedLinks(info *Info, container *RequesterContainer) {
	for _, parser := range info.Parsers {
		if _, ok := container.Requesters[parser]; ok {
			fmt.Println(parser)
			container.Requesters[parser].SetAvailability(true)
		}
	}
}

// Returns info about all users.
func getUsers(selected []byte) (links []byte, err error) {
	var info *Info = new(Info)
	log.Println(string(selected))
	err = json.Unmarshal(selected, info)
	if err != nil {
		log.Println("Unmarshal error")
		return nil, err
	}

	// Container initialization and execution.
	container := NewRequesterContainer(info.Nickname)
	setUsedLinks(info, container)

	users := container.GetLinks()
	links, err = json.Marshal(users)
	if err != nil {
		log.Fatalln("Marshal error")
		return nil, err
	}

	return links, nil
}
