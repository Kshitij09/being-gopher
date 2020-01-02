# GO Workspace
Go package hierarchy should be as shown below:
```
bin/
    hello                          # command executable
    outyet                         # command executable
src/
    github.com/golang/example/
        .git/                      # Git rep/ository metadata
        hello/
            hello.go                   # command source
        outyet/
            main.go                    # command source
            main_test.go               # test source
        stringutil/
            reverse.go                 # package source
            reverse_test.go            # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
        bmp/
            reader.go                  # package source
            writer.go                  # package source

    ... (many more repositories and packages omitted) ...
```

`src` contains Go source files, i.e., `$GOPATH/src`
`bin` contains executable commands, i.e., `$GOPATH/bin`


# Go environment variables:

`GOPATH` specifies the location of your workspace. To make it work with all of your shells, put the export part in `/etc/profile` file

`go env GOPATH` prints the effective current `GOPATH`

Put following lines in your `/etc/profile` file:

- `export GOPATH=$(go env GOPATH)`
- `export PATH=$PATH:$(go env GOPATH)/bin`


You can set go env variables using `-w` flag. For instance, settting `GOBIN` variable as follows

`go env -w GOBIN=$GOPATH/bin`

`go install github.com/<user>/<package>` should put the executable in `$GOPATH/bin` directory, if not, make sure you've set all of your environment variables correctly.
