package main

import "fmt"
import "html/template"
import "log"
import "net/http"
import "net/url"
import "strconv"

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

func display_all_accounts(w http.ResponseWriter, r *http.Request) {
	data := All_users{Get_all_accounts()}
	tmpl := template.Must(template.New(tbl).Parse(tbl))
	err := tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func create_account(w http.ResponseWriter, r *http.Request) {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	// needs to parse url param. Needs to write 500? on return
	name := m.Get("name")
	balance, err := strconv.Atoi(m.Get("balance"))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	NewAccount(name, balance)
	w.WriteHeader(200)
}

func get_account(w http.ResponseWriter, r *http.Request) {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	if id := m.Get("UUID"); id != "" {
		uuid, err := strconv.Atoi(id)
		if err != nil {
			goto Err
		}

		account, err := account_by_uuid(uuid)
		if err != nil {
			goto Err
		}
		data := All_users{[]Account{*account}}
		tmpl := template.Must(template.New(tbl).Parse(tbl))
		err = tmpl.Execute(w, data)
		if err != nil {
			// err should occur earlier, never here
			panic(err)
		}
	}
Err:
	w.WriteHeader(400)
	return

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
	http.HandleFunc("/show-all", display_all_accounts)
	http.HandleFunc("/show", get_account)
	http.HandleFunc("/create-account", create_account)

	log.Fatal(s.ListenAndServe())
}
