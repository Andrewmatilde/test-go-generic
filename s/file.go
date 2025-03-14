package s

type File struct {
	ID       int
	FileName string
	Data     string
}

type FileRequest struct {
	ID int `json:"id"`
}

type FileInterface interface {
	Get(request FileRequest) File
}

type FileClient struct{}

func (f *FileClient) Get(request FileRequest) File {
	return File{
		ID:       request.ID,
		FileName: "John",
		Data:     "Hello, World!",
	}
}
