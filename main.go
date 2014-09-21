package app

import "net/http"
import "fmt"
import "regexp"
import "appengine"
import "strconv"

import . "flotilla"

var put_re = regexp.MustCompile(`^/(?P<value>[a-zA-Z._-][a-zA-Z0-9._-]*)$`)
var get_re = regexp.MustCompile(`^/(?P<value>[a-zA-Z._-][a-zA-Z0-9._-]*)|(?P<number>[0-9]+)$`)

func init() {
  Debug(true)
  Handle("/").Get(get).Put(put).Options(options)
}

func options(r *http.Request) {
  Status(http.StatusOK)
}

func get(r *http.Request) {
  c := appengine.NewContext(r)
  match := Components(get_re, r.URL.Path)
  if value := match["value"]; value != "" {
    number, e := getNumber(c, value)
    Check(e)
    Ensure(number != nil, http.StatusNotFound)
    Body(http.StatusOK, fmt.Sprintf("%v", *number), "text/plain")
  } else if number := match["number"]; number != "" {
    n, e := strconv.ParseInt(number, 10, 64)
    Check(e)
    value, e := getValue(c, n)
    Check(e)
    Ensure(value != nil, http.StatusNotFound)
    Body(http.StatusOK, fmt.Sprintf("%v", *value), "text/plain")
  } else {
    Status(http.StatusForbidden)
  }
}

func put(r *http.Request) {
  c := appengine.NewContext(r)
  match := Components(put_re, r.URL.Path)
  Ensure(match != nil, http.StatusForbidden)
  number, e := getNumber(c, match["value"])
  Check(e)
  if number != nil {
    Body(http.StatusOK, fmt.Sprintf("%v", *number), "text/plain")
  }
  number, e = allocate(c, match["value"])
  Check(e)
  Body(http.StatusCreated, fmt.Sprintf("%v", *number), "text/plain")
}
