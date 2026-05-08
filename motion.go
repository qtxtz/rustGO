package main

type Motion struct {
	client *CGOClient
}

func NewMotion(client *CGOClient) *Motion {
	return &Motion{client: client}
}

func (m *Motion) TouchDown(x, y, fingerID int) error {
	return m.client.TouchDown(x, y, fingerID)
}

func (m *Motion) TouchMove(x, y, fingerID int) error {
	return m.client.TouchMove(x, y, fingerID)
}

func (m *Motion) TouchUp(x, y, fingerID int) error {
	return m.client.TouchUp(x, y, fingerID)
}

func (m *Motion) Click(x, y int) error {
	return m.client.Tap(x, y)
}

func (m *Motion) LongClick(x, y, duration int) error {
	return m.client.LongPress(x, y, duration)
}

func (m *Motion) Swipe(x1, y1, x2, y2, duration int) error {
	return m.client.Swipe(x1, y1, x2, y2, duration)
}

func (m *Motion) InputText(text string) error {
	return m.client.InputText(text)
}

func (m *Motion) KeyEvent(keyCode int) error {
	return m.client.KeyEvent(keyCode)
}

func (m *Motion) Home() error {
	return m.client.KeyEvent(3)
}

func (m *Motion) Back() error {
	return m.client.KeyEvent(4)
}

func (m *Motion) Recents() error {
	return m.client.KeyEvent(187)
}
