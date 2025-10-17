# applet - manage small demonstrative examples

**contact**: [Guillaume Boulant](mailto:gboulant@gmail.com?subject=dingo-applet)

The applet package can be used to create an application that contains a
set of demonstrative examples. Each example is a simple go function
without argument and returning an error:

```go
func DEMO00_logscale() error {
    fmt.Println("Executing demo DEMO00_logscale")
    return nil
}
```

Then, you can register this function as a demonstrative example, giving
it a name identifier and a short description:  

```go
applet.NewExample("D00", "echelle logarithmique", DEMO00_logscale)
applet.NewExample("D01", "son de quintes", DEMO01_quintes)
...
```

And finaly, the main function should execute the `StartExampleApp`:

```go
applet.StartExampleApp("D01") // set D01 to be the default example to run
```

This function parses the command line arguments and executes the
selection made by the user. If the name of your executable program is
`mydemo`, then you can see the list of the demonstrative examples with:

```shell
./mydemo -l
```

That displays:

```text
D00            (echelle logarithmique)
D01            (son de quintes)
...
```

Then to play the example D00 for example:

```shell
./mydemo -n D01
```
