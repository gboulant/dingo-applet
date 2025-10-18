package applet

import (
	"flag"
	"fmt"
	"os"
)

// Applet defines a data structure to describe a little executable procedure
type Applet struct {
	Name    string
	Execute func() error
	Comment string
}

// String return a on-line string representation of the Applet. It can
// be used for listing the applets
func (p Applet) String() string {
	return fmt.Sprintf("%-14s (%s)", p.Name, p.Comment)
}

var applets []Applet

// AddApplet creates a new Applet program and registers the created
// program into the catalog of applets. After this registration, the
// applet program can be obtain from ist name using the function
// GetApplet
func AddApplet(name string, comment string, function func() error) *Applet {
	p := Applet{name, function, comment}
	applets = append(applets, p)
	return &p
}

func GetAppletNames() []string {
	names := make([]string, len(applets))
	for i, p := range applets {
		names[i] = p.Name
	}
	return names
}

// GetApplet returns (if exists) the Applet selected by its identifier
// name, return nil and an error if no applet is register with this
// name.
func GetApplet(name string) (*Applet, error) {
	for _, example := range applets {
		if example.Name == name {
			return &example, nil
		}
	}
	return nil, fmt.Errorf("no example program with name %s", name)
}

// ListApplets prints on the standard output the list of applets
// registered with AddApplet
func ListApplets() {
	for _, p := range applets {
		fmt.Println(p)
	}
}

// StartApplication parses the command line arguments and execute the
// selected applet if specified, or print the list of applets if the
// option -l is specified.
func StartApplication(defaultAppletName string) {
	var listApplets bool
	var appletName string
	flag.BoolVar(&listApplets, "l", false, "list the applets")
	flag.StringVar(&appletName, "n", defaultAppletName, "name of the applet to execute")
	flag.Parse()

	if listApplets {
		ListApplets()
		os.Exit(0)
	}

	p, err := GetApplet(appletName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Executing applet %s ...\n", p.Name)
	if err := p.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
