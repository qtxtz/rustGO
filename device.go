package main

type Device struct {
	client *CGOClient
}

func NewDevice(client *CGOClient) *Device {
	return &Device{client: client}
}

func (d *Device) Info() (string, error) {
	return d.client.DeviceInfo()
}

func (d *Device) GetScreenSize() (string, error) {
	return d.client.GetScreenSize()
}

func (d *Device) IsScreenOn() bool {
	return d.client.IsScreenOn()
}

func (d *Device) Screenshot(path string) error {
	return d.client.Screenshot(path)
}

func (d *Device) Wait(ms int) error {
	return d.client.Wait(ms)
}

func (d *Device) GetClipboard() (string, error) {
	return d.client.GetClipboard()
}

func (d *Device) SetClipboard(text string) error {
	return d.client.SetClipboard(text)
}
