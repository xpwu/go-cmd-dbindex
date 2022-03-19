package indexcmd

import (
  "reflect"
  "sync"
)

type Creator interface {
  CreateIndex()
}

type value struct {
  newCreator func() Creator
  help       string
}

var (
  allNewCreators = make([]value, 0, 10)
  mu             = sync.Mutex{}
)

func Add(f func() Creator, help string) {
  mu.Lock()
  defer mu.Unlock()

  allNewCreators = append(allNewCreators, value{
    newCreator: f,
    help:       help,
  })
}

func getGoInfo(creator Creator) string {
  typ := reflect.TypeOf(creator)
  if typ.Kind() == reflect.Ptr {
    typ = typ.Elem()
  }

  return typ.String()
}

