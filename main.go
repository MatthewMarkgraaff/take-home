package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type RepositoryInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"html_url"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./www")))
	http.HandleFunc("/form", gitFormHandler)
	log.Fatal(http.ListenAndServe(":2107", nil))
}

func gitFormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing input form: %v", err)
		return
	}
	userName := r.FormValue("userName")

	// Create url for repos using input userName from form
	base := "https://api.github.com/users/"
	urlString := base + userName + "/repos"
	url, err := url.Parse(urlString)
	if err != nil {
		fmt.Fprintf(w, "Error parsing url form: %v", err)
		return
	}

	// Get user repositories
	response, err := http.Get(url.String())
	if err != nil {
		fmt.Fprintf(w, "The Github HTTP request failed: %v", err)
		return
	}

	data, _ := ioutil.ReadAll(response.Body)
	var repositories []RepositoryInfo

	err = json.Unmarshal(data, &repositories)
	if err != nil {
		fmt.Fprintf(w, "Error unmarshalling response: %v", err)
		return
	}

	// Create reponse form
	html := fmt.Sprintf(`
			<html>
				<head>
  					<meta charset="UTF-8" />
					<link rel="stylesheet" href="styles.css">
				</head>
				<body>
				<h2>%s's public Github repositories: </h2>
				<table>
					<tbody>
						<tr>
							<td><b>Repository Name</b></td>
							<td><b>Repository Url</b></td>
						</tr>
			`, userName)

	// Table rows
	for _, repo := range repositories {
		html += fmt.Sprintf(`
			<tr>
				<td >%s</td>
				<td ><a href="%s">%s</a></td>
			</tr>`,
			repo.Name, repo.Url, repo.Url)
	}

	//  Close all tags and send
	html += `</tbody></table><body></html>`
	fmt.Fprint(w, html)
}
