package indexcmd

import (
  "fmt"
  "github.com/xpwu/go-cmd/arg"
  "github.com/xpwu/go-cmd/cmd"
)

func init() {
  cmd.RegisterCmd("listindex", "list all index", func(args *arg.Arg) {
    mu.Lock()
    defer mu.Unlock()

    for i, v := range allNewCreators {
      c := v.newCreator()
      fmt.Printf("\t%d -> %s, creator: %s\n", i, v.help, getGoInfo(c))
    }
  })
}
