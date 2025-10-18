# applet - manage small little applications

**contact**: [Guillaume Boulant](mailto:gboulant@gmail.com?subject=dingo-applet)

The applet package can be used to create an executable program that can
play a set of different little applications (called applets). The
typical use case is to create a tutorial program that demonstrates
different features of a go package with little examples. Each applet is
a standard go function without argument and returning an error:

```go
func DEMO00_logscale() error {
    fmt.Println("Executing demo DEMO00_logscale")
    // do something
    // ...
    return nil
}
```

Then, you can register the function as an applet, giving it a name
identifier and a short description:  

```go
applet.NewExample("D00", "echelle logarithmique", DEMO00_logscale)
applet.NewExample("D01", "son de quintes", DEMO01_quintes)
...
```

And finaly, the main program should execute the `StartExampleApp`:

```go
applet.StartExampleApp("D01") // set D01 to be the default example to run
```

This function parses the command line arguments and executes the
selection of the user. For example, if the name of your executable
program is `mydemo`, then you can see the list of the applets with:

```shell
./mydemo -l
```

That displays:

```text
D00            (echelle logarithmique)
D01            (son de quintes)
...
```

Then to play the applet D01 for example:

```shell
./mydemo -n D01
```
