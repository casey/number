package app

import "appengine"
import "appengine/datastore"

type Counter struct {
  Count int64 `datastore:"count",noindex`
}

var next        int64 = 0
var end         int64 = 0
var reservation int64 = 1

func numberRoot(c appengine.Context, number int64) *datastore.Key {
  return datastore.NewKey(c, "Number",  "", number, nil)
}

func nameRoot(c appengine.Context, name string) *datastore.Key {
  return datastore.NewKey(c, "Name",  name, 0, nil)
}

func getName(c appengine.Context, number int64) (*string, error) {
  keys, e := datastore.NewQuery("Name").Ancestor(numberRoot(c, number)).KeysOnly().GetAll(c, nil)
  if e != nil {
    return nil, e
  } else if len(keys) == 0 {
    return nil, nil
  } else {
    result := new(string)
    *result = keys[0].StringID()
    return result, nil
  }
}

func getNumber(c appengine.Context, name string) (*int64, error) {
  keys, e := datastore.NewQuery("Number").Ancestor(nameRoot(c, name)).KeysOnly().GetAll(c, nil)
  if e != nil {
    return nil, e
  } else if len(keys) == 0 {
    return nil, nil
  } else {
    result := new(int64)
    *result = keys[0].IntID()
    return result, nil
  }
}

func allocate(c appengine.Context, name string) (*int64, error) {
  if next == end {
    var nextNext int64 = 0
    var nextEnd  int64 = 0
    
    e := datastore.RunInTransaction(c, func(c appengine.Context) error {
      key := datastore.NewKey(c, "Counter", "counter0", 0, nil)
      count := Counter{}
      e := datastore.Get(c, key, &count)
      if e == datastore.ErrNoSuchEntity {
        count.Count = 1
      } else if e != nil {
          return e
      }
      nextNext = count.Count
      nextEnd = nextNext + reservation
      count.Count = nextEnd
      
      _, e = datastore.Put(c, key, &count)
      return e
    }, nil)

    if e != nil {
      return nil, e
    }

    next = nextNext
    end = nextEnd
  }

  opts := datastore.TransactionOptions{XG: true}
  e := datastore.RunInTransaction(c, func(c appengine.Context) error {
    k1 := datastore.NewKey(c, "Name"  , name, 0   , numberRoot(c, next))
    k2 := datastore.NewKey(c, "Number", ""  , next, nameRoot  (c, name))
    if _, e := datastore.Put(c, k1, &struct{}{}); e != nil { return e }
    if _, e := datastore.Put(c, k2, &struct{}{}); e != nil { return e }
    return nil
  }, &opts)

  if e == nil {
    assigned := next
    next++
    return &assigned, nil
  } else {
    return nil, e
  }
}
