package indexcmd

import (
  "fmt"
  "github.com/xpwu/go-cmd/arg"
  "github.com/xpwu/go-cmd/cmd"
  "os"
)

func init() {
  cmd.RegisterCmd("createindex", "create index by num, which is listed by <listindex> cmd",
    func(args *arg.Arg) {
      num := -1
      args.Int(&num, "i", "which one?")
      arg.ReadConfig(args)

      args.Parse()

      mu.Lock()
      defer mu.Unlock()

      if num < 0 || num >= len(allNewCreators) {
        fmt.Printf("ERROR: i must be >0 && < len(all),%d\n", len(allNewCreators))
        os.Exit(1)
      }

      c := allNewCreators[num].newCreator()
      fmt.Printf("\tcreating %d: %s, creator: %s\n", num, allNewCreators[num].help, getGoInfo(c))
      c.CreateIndex()
      fmt.Print("succeed")

    })
}
