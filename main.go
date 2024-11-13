package main

var (
	AppName    string
	AppVersion string
)

func main() {
	tui := TermUI{}

	if err := tui.Run(); err != nil {
		panic(err)
	}
}
