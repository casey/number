number
======

Assign numbers to names with low storage overhead.

[Test instance here.](http://rodarmor-number.appspot.com)


API
---

A NAME matches `/[a-zA-Z._-][a-zA-Z0-9._-]*/`.
A NUMBER matches `/[0-9]+/`.

* PUT /NAME -> Allocate number for NAME. Numbers are allocated in sequenceish starting at 0.
* GET /NAME -> Get number for NAME.
* GET /NUMBER -> Get name for NUMBER.

To Do
-----

* Some kind of self-healing for filling in holes.
