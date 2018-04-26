package cmd

import "github.com/fatih/color"

var (
	// Version is the default version of BDT
	Version = "v0.0.1-beta"
	logo    = `
	██████╗ ██████╗         ████████╗██████╗  █████╗ ███╗   ██╗███████╗██╗      █████╗ ████████╗ ██████╗ ██████╗
	██╔══██╗██╔══██╗        ╚══██╔══╝██╔══██╗██╔══██╗████╗  ██║██╔════╝██║     ██╔══██╗╚══██╔══╝██╔═══██╗██╔══██╗
	██████╔╝██║  ██║           ██║   ██████╔╝███████║██╔██╗ ██║███████╗██║     ███████║   ██║   ██║   ██║██████╔╝
	██╔══██╗██║  ██║           ██║   ██╔══██╗██╔══██║██║╚██╗██║╚════██║██║     ██╔══██║   ██║   ██║   ██║██╔══██╗
	██████╔╝██████╔╝           ██║   ██║  ██║██║  ██║██║ ╚████║███████║███████╗██║  ██║   ██║   ╚██████╔╝██║  ██║
	╚═════╝ ╚═════╝            ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝
	BDT %s
	https://github.com/mojo-zd/bd`
)

func DisplayLogo() {
	color.Cyan(logo, Version)
}
