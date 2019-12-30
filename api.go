package main

import "fmt"
import "html/template"
import "log"
import "net/http"

//import "net/url"
import "time"

type All_users struct {
	// does this need to be exported?
	Accounts []Account
}

const tbl = `
<table>
  <tr>
	<th>UUID</th>
	<th>Name</th>
	<th>Balance</th>
  </tr>
  {{range .Accounts}}
   <tr>
    <td>{{.UUID}}</td>
    <td>{{.Name}}</td>
    <td>{{.Balance}}</td>
  </tr>
  {{end}}
</table>
`

func display_all_users(w http.ResponseWriter, r *http.Request) {
	data := All_users{Get_all_accounts()}
	tmpl := template.Must(template.New(tbl).Parse(tbl))
	err := tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func create_user(w http.ResponseWriter, r *http.Request) {
	//m, _ := url.ParseQuery(r.URL.RawQuery)
	// needs to parse url param. Needs to write 500? on return

}

func InitServer() {
	s := &http.Server{
		Addr: ":8080",
		//Handler:        handle,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/server-config", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "SERVER CONFIG:\n%+v\n", *s)
	})
	http.HandleFunc("/show-all", display_all_users)

	log.Fatal(s.ListenAndServe())
}
