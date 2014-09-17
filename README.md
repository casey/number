id
==

Assign numbers to keys.

[Test instance here.](http://rodarmor-id.appspot.com)

API
---

A KEY matches `/[a-zA-Z._-][a-zA-Z0-9._-]*/`.
A NUMBER matches `/[0-9]+/`.

* GET /KEY -> get number for KEY.
* GET /NUMBER -> get key for NUMBER.
* PUT /KEY -> Set number for KEY. Numbers are assigned in sequence starting at 0.
