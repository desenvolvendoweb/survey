package main

import (
    "strings"
    "strconv"

    _ "github.com/lib/pq"
)


type User struct {
    ID        int
    UserName  string
    Email     string
}

type Users []User

func getListUsers() (*Users) {

    db := Connect()

    users := make(Users, 0)

  	rows, err := db.Query("SELECT id, username, email FROM auth_user")

  	checkErr(err)

  	for rows.Next() {

  		var id       int
  		var username string
  		var email    string

  		err = rows.Scan(&id, &username, &email)
      checkErr(err)

  		users = append(users, User{ID: id, UserName: username, Email: email})

  	}

    err = rows.Err()
    checkErr(err)

    return &users

}


type Survey struct {
    ID           int
    Title        string
    Alternatives *Alternatives
}

type Surveys []Survey

func getListSurveys() (*Surveys) {

    db := Connect()

    surveys := make(Surveys, 0)

    rows, err := db.Query("SELECT id, title FROM vote_survey WHERE status=true "+
      "ORDER BY date DESC")

    checkErr(err)

    for rows.Next() {

      var id        int
      var title     string

      err = rows.Scan(&id, &title)
      checkErr(err)

      alternatives := getListAlternatives(id)

      surveys = append(surveys, Survey{ID: id, Title: title, Alternatives: alternatives})

    }

    err = rows.Err()
    checkErr(err)

    return &surveys

}

func getListSurvey(id int) (*Survey) {

    db := Connect()

    survey := Survey{}

    rows, err := db.Query("SELECT id, title FROM vote_survey WHERE status=true "+
      "AND id=$1 ORDER BY date DESC", id)

    checkErr(err)

    for rows.Next() {

      var id        int
      var title     string

      err = rows.Scan(&id, &title)
      checkErr(err)

      alternatives := getListAlternatives(id)

      survey = Survey{ID: id, Title: title, Alternatives: alternatives}

    }

    err = rows.Err()
    checkErr(err)

    return &survey

}


type Alternative struct {
    ID        int
    Title     string
    Votes     int
}

type Alternatives []Alternative

func getListAlternatives(survey_id int) (*Alternatives) {

    db := Connect()

    alternatives := make(Alternatives, 0)

    rows, err := db.Query("SELECT id, title, votes FROM vote_alternative "+
      "WHERE status=true AND survey_id=$1 ORDER BY date DESC", survey_id)

    checkErr(err)

    for rows.Next() {

      var id        int
      var title     string
      var votes     int

      err = rows.Scan(&id, &title, &votes)
      checkErr(err)

      alternatives = append(alternatives, Alternative{ID: id, Title: title, Votes: votes})

    }

    err = rows.Err()
    checkErr(err)

    return &alternatives

}


type Result struct {
    ID      int
    Type    string
    Msg     string
}

func voteSurvey(id, alternative_id int) (*Result) {

    db := Connect()

    stmt, err := db.Prepare("UPDATE vote_alternative SET votes=votes+$1 WHERE id=$2")
    checkErr(err)

    res, err := stmt.Exec(1, alternative_id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    result := Result{}

    if affect > 0 {

        msg := []string{strconv.FormatInt(int64(affect), 10), "row changed"}

        result = Result{ID: alternative_id, Type: "success", Msg: strings.Join(msg, " ")}

    } else {

        result = Result{ID: alternative_id, Type: "error", Msg: "alternative id not exists !"}

    }

    return &result

}

func checkSurveyAlternative(id, alternative_id int) (*Result) {

    db := Connect()

    rows := db.QueryRow("SELECT COUNT(*) AS count FROM vote_survey AS VS "+
      "INNER JOIN vote_alternative AS VA ON VA.survey_id = VS.id "+
      "WHERE VS.status=true AND VA.status=true AND VS.id=$1 AND VA.id=$2",
      id, alternative_id)

    var count int
    rows.Scan(&count)

    if count == 0 {
        result := Result{ID: alternative_id, Type: "error", Msg: "survey id or alternative id not exists !"}
        return &result
    }

    return nil

}
