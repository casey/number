package app

import "net/http"
import "fmt"
import "regexp"
import "appengine"
import "strconv"

import . "flotilla"

var put_re = regexp.MustCompile(`^/(?P<name>[a-zA-Z._-][a-zA-Z0-9._-]{0,63})$`)
var get_re = regexp.MustCompile(`^/(?P<name>[a-zA-Z._-][a-zA-Z0-9._-]{0,63})|(?P<number>[0-9]+)$`)

func displayNumber(number int64) string {
  return fmt.Sprintf("%v", number - 1)
}

func parseNumber(number string) (int64, error) {
  n, e := strconv.ParseInt(number, 10, 64)
  return n + 1, e
}

func init() {
  Handle("/").Get(get).Put(put).Options(options)
}

func options(r *http.Request) {
  Status(http.StatusOK)
}

func get(r *http.Request) {
  c := appengine.NewContext(r)
  match := Components(get_re, r.URL.Path)
  if name := match["name"]; name != "" {
    number, e := getNumber(c, name)
    Check(e)
    Ensure(number != nil, http.StatusNotFound)
    Body(http.StatusOK, displayNumber(*number), "text/plain")
  } else if number := match["number"]; number != "" {
    n, e := parseNumber(number)
    Check(e)
    name, e := getName(c, n)
    Check(e)
    Ensure(name != nil, http.StatusNotFound)
    Body(http.StatusOK, fmt.Sprintf("%v", *name), "text/plain")
  } else {
    Status(http.StatusForbidden)
  }
}

func put(r *http.Request) {
  c := appengine.NewContext(r)
  match := Components(put_re, r.URL.Path)
  Ensure(match != nil, http.StatusForbidden)
  number, e := getNumber(c, match["name"])
  Check(e)
  if number != nil {
    Body(http.StatusOK, displayNumber(*number), "text/plain")
  }
  number, e = allocate(c, match["name"])
  Check(e)
  Body(http.StatusCreated, displayNumber(*number), "text/plain")
}
