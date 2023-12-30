package main

type File struct {
	ID         string   `json:"_key,omitempty" validate:"-"`
	Name       string   `json:"name"`
	StoredName string   `json:"stored_name"`
	Tags       []string `json:"tags"`
	Type       string   `json:"type"`
	CreatedAt  int64    `json:"created_at"`
	Link       string   `json:"link"`
}

func (f *File) addLink() {
	//f.Link = fmt.Sprintf()
}
