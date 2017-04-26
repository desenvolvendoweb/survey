package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"
    "regexp"
    "index/suffixarray"

    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func ListUser(w http.ResponseWriter, r *http.Request) {

    users := getListUsers()

    err := json.NewEncoder(w).Encode(users)
    checkErr(err)

}

func ListSurveys(w http.ResponseWriter, r *http.Request) {

    surveys := getListSurveys()

    err := json.NewEncoder(w).Encode(surveys)
    checkErr(err)

}

func ListSurvey(w http.ResponseWriter, r *http.Request) {

    vars    := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    survey := getListSurvey(id)

    err = json.NewEncoder(w).Encode(survey)
    checkErr(err)

}

func VoteSurvey(w http.ResponseWriter, r *http.Request) {

    type ResponseHttpRequestJson struct {
        Alternative_id int
    }

    vars    := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    decoder := json.NewDecoder(r.Body)

    var data ResponseHttpRequestJson
  	err = decoder.Decode(&data); if err != nil {
  	     panic(err)
  	}

    check := checkSurveyAlternative(id, data.Alternative_id); if check != nil {
        err = json.NewEncoder(w).Encode(check)
        checkErr(err)
        return
    }

    var store = sessions.NewCookieStore([]byte("something-very-secret"))

    session, err := store.Get(r, "session-name")

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    items := fmt.Sprint(session.Values["survey_itens"])

    suff_items  := suffixarray.New([]byte(items))
    occurrences := suff_items.FindAllIndex(regexp.MustCompile(strconv.FormatInt(int64(id), 10)), -1)

    if len(occurrences) > 0 {

        result := Result{ID: data.Alternative_id, Type: "error", Msg: "Vote already done!"}
        err = json.NewEncoder(w).Encode(result)
        return

    } else {

        // Set some session values.
        session.Values["survey_itens"] = strings.Join([]string{items, strconv.FormatInt(int64(id), 10)}, ",")
        // Save it before we write to the response/return from the handler.
        session.Save(r, w)

    }

    vote := voteSurvey(id, data.Alternative_id)

    err = json.NewEncoder(w).Encode(vote)
    checkErr(err)

}
