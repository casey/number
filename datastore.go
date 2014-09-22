package app

import "appengine"
import "appengine/datastore"

type Entry struct {
  Number int64  `datastore:"number"`
  Value  string `datastore:"value"`
}

type Counter struct {
  Count int64 `datastore:"count",noindex`
}

var next        int64 = 0
var end         int64 = 0
var reservation int64 = 1

func query(c appengine.Context, lhs string, rhs interface{}) (*Entry, error) {
  results := make([]Entry, 1)
  keys, e := datastore.NewQuery("Entry").Filter(lhs, rhs).Limit(len(results)).GetAll(c, &results)
  if e != nil {
    return nil, e
  } else if len(keys) > 0 {
    return &results[0], nil
  } else {
    return nil, nil
  }
}

func getValue(c appengine.Context, number int64) (*string, error) {
  result, e := query(c, "number =", number)
  if result != nil {
    return &result.Value, nil
  } else {
    return nil, e
  }
}

func getNumber(c appengine.Context, value string) (*int64, error) {
  result, e := query(c, "value =", value)
  if result != nil {
    return &result.Number, nil
  } else {
    return nil, e
  }
}

func allocate(c appengine.Context, value string) (*int64, error) {
  if next == end {
    var nextNext int64 = 0
    var nextEnd  int64 = 0
    
    e := datastore.RunInTransaction(c, func(c appengine.Context) error {
      key := datastore.NewKey(c, "Counter", "counter0", 0, nil)
      count := Counter{}
      e := datastore.Get(c, key, &count)
      if e != nil && e != datastore.ErrNoSuchEntity {
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

  key := datastore.NewKey(c, "Entry", "", 0, nil)
  entry := Entry{next, value}
  _, e := datastore.Put(c, key, &entry)
  
  if e == nil {
    next++
    return &entry.Number, nil
  } else {
    return nil, e
  }
}
