package main

type Storages struct {
	client *CGOClient
}

func NewStorages(client *CGOClient) *Storages {
	return &Storages{client: client}
}

func (s *Storages) Exists(path string) bool {
	return s.client.FileExists(path)
}

func (s *Storages) Read(path string) (string, error) {
	return s.client.FileRead(path)
}

func (s *Storages) Write(path, content string) error {
	return s.client.FileWrite(path, content)
}
