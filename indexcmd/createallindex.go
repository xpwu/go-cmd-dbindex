package indexcmd

import (
  "fmt"
  "github.com/xpwu/go-cmd/arg"
  "github.com/xpwu/go-cmd/cmd"
)

func init() {
  cmd.RegisterCmd("createallindex", "create all index", func(args *arg.Arg) {
    arg.ReadConfig(args)
    args.Parse()

    mu.Lock()
    defer mu.Unlock()

    fmt.Print("--- create all index ... ---")
    for i, v := range allNewCreators {
      c := v.newCreator()
      fmt.Printf("\tdoing %d: %s, creator: %s\n", i, v.help, getGoInfo(c))
      c.CreateIndex()
    }
    fmt.Print("--- create all index -- end ---")
  })
}
