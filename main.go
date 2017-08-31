package main

import "fmt"

func main() {

	server := NewServer()
	server.Start()
	<-server.exit
}

// display version info
func displayVersion() {
	fmt.Println("Git Commit:", GitCommit)
	fmt.Println("Version:", Version)
	if VersionPrerelease != "" {
		fmt.Println("Version PreRelease:", VersionPrerelease)
	}
}

// display header when program starts
func displayHeader() {

	fmt.Println()
	fmt.Println(".................. ............. /' /)")
	fmt.Println("................./´ /)........./¯ //")
	fmt.Println("..............,/¯// ......... /...//")
	fmt.Println("............./...//. ......./¯ //")
	fmt.Println(".........../´¯/'´ ¯/´¯ /.../ /")
	fmt.Println("........./'.../... ./... /.../ //")
	fmt.Println("........('(...´(... ....... ,../'. .')")
	fmt.Println(".........\\.......... ..... ..\\/..../")
	fmt.Println("..........''...\\.... ..... . _.•´")
	fmt.Println("............\\....... ..... ..(")
	fmt.Println("..............\\..... ..... ..")
	fmt.Println()
	fmt.Println("http://www.github.com/anguenot/ipfix-forwarder")
	fmt.Println()

	displayVersion()
	fmt.Println()
}
