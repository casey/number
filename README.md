id
==

Assign numbers to keys.

[Test instance here.](http://rodarmor-id.appspot.com)

API
---

A VALUE matches `/[a-zA-Z._-][a-zA-Z0-9._-]*/`.
A NUMBER matches `/[0-9]+/`.

* GET /VALUE -> get number for VALUE.
* GET /NUMBER -> get value for NUMBER.
* PUT /VALUE -> Set value for KEY. Numbers are assigned in sequenceish starting at 0.
