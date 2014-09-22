id
==

Assign numbers to keys.

[Test instance here.](http://rodarmor-id.appspot.com)

API
---

A KEY matches `/[a-zA-Z._-][a-zA-Z0-9._-]*/`.
A NUMBER matches `/[0-9]+/`.

* PUT /KEY -> Allocate number for KEY. Numbers are allocated in sequenceish starting at 0.
* GET /KEY -> Get number for KEY.
* GET /NUMBER -> Get key for NUMBER.
