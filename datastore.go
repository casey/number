package app

import "appengine"
import "appengine/datastore"

type entry struct {
  Number int64  `datastore:"number"`
  Value  string `datastore:"value"`
}

type counter struct {
  Count int64 `datastore:"count"`
}

var next        int64 = 0
var end         int64 = 0
var reservation int64 = 1

func getValue(c appengine.Context, number int64) (*string, error) {
  results := make([]entry, 1)
  _, e := datastore.NewQuery("Entry").Filter("number =", number).Limit(len(results)).GetAll(c, &results)
  if e == datastore.ErrNoSuchEntity {
    return nil, nil
  } else if e == nil {
    return &results[0].Value, nil
  } else {
    return nil, e
  }
}

func getNumber(c appengine.Context, value string) (*int64, error) {
  results := make([]entry, 1)
  _, e := datastore.NewQuery("Entry").Filter("value =", value).Limit(len(results)).GetAll(c, &results)
  if e == datastore.ErrNoSuchEntity {
    return nil, nil
  } else if e == nil {
    return &results[0].Number, nil
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
      count := counter{}
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
  entry := entry{next, value}
  _, e := datastore.Put(c, key, &entry)
  
  if e == nil {
    next++
    return &entry.Number, nil
  } else {
    return nil, e
  }
}
