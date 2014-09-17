package app

import "net/http"

import . "flotilla"

func init() {
  Handle("/").Get(get).Put(put).Options(options)
}

func options(r *http.Request) {
  Status(http.StatusOK)
}

func get(r *http.Request) {
  Status(http.StatusOK)
}

func put(r *http.Request) {
  Status(http.StatusOK)
}
