package main

type App struct {
	client *CGOClient
}

func NewApp(client *CGOClient) *App {
	return &App{client: client}
}

func (a *App) Launch(packageName string) error {
	return a.client.LaunchApp(packageName)
}

func (a *App) Close(packageName string) error {
	return a.client.CloseApp(packageName)
}

func (a *App) CurrentPackage() (string, error) {
	return a.client.CurrentPackage()
}

func (a *App) CurrentActivity() (string, error) {
	return a.client.CurrentActivity()
}

func (a *App) Shell(cmd string) (string, error) {
	return a.client.Shell(cmd)
}
