package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type config struct {
  targetPaths []string
  startingPath string
  template *template.Template
}

func newConfig(targetPaths []string, tpl string) (config, error) {
  startingPath, e := os.Getwd()

  if e != nil {
    return config{}, e
  }

  t, e := template.New("t").Parse(tpl)

  if e != nil {
    return config{}, e
  }

  return config{ targetPaths: targetPaths, startingPath: startingPath, template: t }, nil
}

func main() {
  tplPtr := flag.String("t", "{{.Target}}", "the template to use when a match is found. Valid tokens {{.Path}}, {{.Target}}")
  flag.Parse()
  args := flag.Args()

  if len(args) < 1 {
    fmt.Fprintln(os.Stderr, "not enough arguments")
    os.Exit(1)
  }

  cfg, e := newConfig(args, *tplPtr)

  if e != nil {
    fmt.Fprintf(os.Stderr, "Error: %s\n", e)
    os.Exit(2)
  }

  nearest := findNearest(cfg)

  fmt.Println(nearest)
}

func findNearest(cfg config) string {
  curPath := cfg.startingPath
  rootPath := string(filepath.Separator)

  for {
    for _, targetPath := range cfg.targetPaths {
      testPath := filepath.Join(curPath, targetPath)

      if _, e := os.Stat(testPath); e == nil {
        var buffer bytes.Buffer
        cfg.template.Execute(
          &buffer,
          map[string]string{
            "Path": curPath,
            "Target": testPath,
          },
        )
        return buffer.String()
      }
    }

    curPath = filepath.Dir(curPath)

    if curPath == rootPath {
      return ""
    }
  }
}
