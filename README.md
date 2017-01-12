Nogo
======

Run golang's `go` command with a dynamic and contextuel GOPATH, à la git.

Why ?
-------
In golang, except really simple programs, project needs GOPATH to hold things like `bin`,`pkg`,`vendor` ... 
and developers may not always will to keep a single system wide GOPATH, but one GOPATH per projet

Nogo do this simply, à la git

How to use
-------------
Simply use `nogo` à la place de `go`.

What it does
-------------
  - If GOPATH is defined, use it
  - Else : from the current directory to /, find a file named `.gopath`
    - The first found, the owner directory will be the GOPATH
    - In addition, content of `.gopath` will be appended to GOPATH
      - #comment is ignored in `.gopath`
  - Passes your command through to `go` with GOPATH
  - If nothing found, launch `go` anyway without

Usage
-------
 - nogo : print the actual usage
 - nogo -init : create a `.gopath` in current directory, then create a `src`subdirectory
 - nogo -print : print the found GOPATH

Exemple
-------

 > cd /some_path/myproject
 > nogo -init
 > echo /corporate/common/golibs >> .gopath
 > echo ~/golibs >> .gopath
 > git clone http://github.com/usr/projet src/
 > cd src
 > nogo build .

Credits
-------
Ideas from [golang]:https://github.com/golang/go/issues/17271#issuecomment-265932522 and [gopath]:https://github.com/nickcarenza/gopath
