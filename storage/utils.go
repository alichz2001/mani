package main

import (
	"errors"
	"mime/multipart"
	"strings"
)

func MapAndValidate(form *multipart.Form) (*UploadFileReq, error) {
	ufr := &UploadFileReq{}

	for title, values := range form.Value {
		err := ufr.addVal(title, values)
		if err != nil {
			return nil, err
		}
	}

	if len(form.File) != 1 {
		return nil, errors.New("only accept one file")
	}

	ufr.File = form.File["file"][0]

	return ufr, nil
}

func genFileName(fileName string) string {
	//TODO some encodeing and filtering character
	return fileName
}

func stringifyListForAQL(list []string) string {
	return "[\"" + strings.Join(list, "\",\"") + "\"]"

}
