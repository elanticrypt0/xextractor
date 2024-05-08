package main

import "fmt"

func AppBanner() {

	appBanner := `
	██╗  ██╗     ███████╗██╗  ██╗████████╗██████╗  █████╗  ██████╗████████╗ ██████╗ ██████╗ 
	╚██╗██╔╝     ██╔════╝╚██╗██╔╝╚══██╔══╝██╔══██╗██╔══██╗██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗
	 ╚███╔╝█████╗█████╗   ╚███╔╝    ██║   ██████╔╝███████║██║        ██║   ██║   ██║██████╔╝
	 ██╔██╗╚════╝██╔══╝   ██╔██╗    ██║   ██╔══██╗██╔══██║██║        ██║   ██║   ██║██╔══██╗
	██╔╝ ██╗     ███████╗██╔╝ ██╗   ██║   ██║  ██║██║  ██║╚██████╗   ██║   ╚██████╔╝██║  ██║
	╚═╝  ╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝
		
		Extract emails and domains from a file
		AUTHOR: k23dev
	`

	fmt.Println(appBanner)
	fmt.Printf("\n")

	fmt.Println("Please, write the file's path using the -file flag or -clean to remove duplicates in a file.")
	fmt.Println("Select what you want to extract with: -ex [email|domain|x(both)]")
	fmt.Println("Use --help to see all options")
	fmt.Printf("\n")
}
