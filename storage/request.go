package main

import (
	"errors"
	"mime/multipart"
)

type UploadFileReq struct {
	Name     string
	FileType string
	Tags     []string
	File     *multipart.FileHeader
}

func (ufr *UploadFileReq) addVal(title string, values []string) error {
	switch title {
	case "name":
		if len(values) > 1 {
			return errors.New("file should have one name")
		}
		ufr.Name = values[0]
		break
	case "type":
		if len(values) > 1 {
			return errors.New("file should have one type")
		}
		ufr.FileType = values[0]
		break
	case "tags":
		ufr.Tags = values
		break
	default:
		return nil
	}
	return nil
}
