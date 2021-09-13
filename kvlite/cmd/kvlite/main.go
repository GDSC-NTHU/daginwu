package main

func main() {
	app := &App{}
	ctl := &AppCtl{
		app: app,
	}
	go ctl.Init()
	app.New()
	app.Init()
	app.Run()
}
