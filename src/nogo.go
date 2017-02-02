package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var flagInit = flag.Bool("init", false, "Make current directory as GOPATH and create \"src\" direcotry if not exists")
var flagPrint = flag.Bool("print", false, "Print the GOPATH value, don't actually run \"go\"")

func main() {
	// Usage print command syntax and a brief usage sample
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: Runs the \"go\" command with a dynamic GOPATH\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	cwd, _ := os.Getwd()
	gopath := os.Getenv("GOPATH")
	_, err := os.Stat(filepath.Join(cwd, ".gopath"))
	if *flagInit {
		dotf, err := os.OpenFile(".gopath", os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot init: %v\n", err)
			os.Exit(1)
		}
		dotf.Close()
		fmt.Printf("%v\n", cwd)
		os.Mkdir("src", os.ModePerm) // ingore error here
		os.Exit(0)
	}

  for {
    if _, err = os.Stat(filepath.Join(cwd, ".gopath")); err == nil {
      // found .gopath, cwd + content of .gopath as GOPATH
      gopath = cwd
      if fcontent, err := ioutil.ReadFile(filepath.Join(cwd, ".gopath")); err == nil {
        lines := strings.Split(string(fcontent), "\n")
        for _, s := range lines {
          s = strings.TrimRight(s, "\r\n")
          if len(s) > 0 {
            if "#" == s[:1] {
              continue // line starts with # is comment, ignore
            }
            gopath = gopath + ":" + s
          }
        }
      }
      // we stop at first .gopath found
      break
    }
    if cwd == "/" {
      break
    }
    cwd = filepath.Dir(cwd)
  }
	if (*flagPrint) || (1 == len(os.Args)) {
		fmt.Printf("%v\n", gopath)
		os.Exit(0)
	}
	if "" != gopath {
		os.Setenv("GOPATH", gopath)
	}
	cmd := exec.Command("go", os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
