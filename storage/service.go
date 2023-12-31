package main

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type StorageServiceImpl struct {
	repo            *ManiDB
	maxFileSize     int64
	maxStorageSize  int64
	usedStorageSize int64
	basePath        string
}

type FetchFileReq struct {
	Name string
	Tags []string
}

type StorageService interface {
	Upload(*UploadFileReq) error
	Fetch(*FetchFileReq) ([]*File, error)
	doUploadFile(*multipart.FileHeader, string) error
}

func NewStorageService(repo *ManiDB) *StorageServiceImpl {

	err := os.MkdirAll(StorageBasePath, 0755)
	if err != nil {
		panic(err)
	}

	return &StorageServiceImpl{
		repo:           repo,
		maxStorageSize: MaxStorageSize,
		maxFileSize:    MaxFileSize,
		basePath:       StorageBasePath,
		//TODO cal used size in startup
		usedStorageSize: 0,
	}
}

func (s *StorageServiceImpl) Upload(ctx context.Context, uploadReq *UploadFileReq) (*File, error) {

	//TODO check  duplicate file name
	fileName := genFileName(uploadReq.Name)

	UploadErr := s.doUploadFile(uploadReq.File, fileName)
	if UploadErr != nil {

		if errors.Is(UploadErr, os.ErrExist) {
			return nil, DuplicateFileName
		}

		return nil, UnknownError
	}

	file := &File{
		Name:       uploadReq.Name,
		Tags:       uploadReq.Tags,
		StoredName: fileName,
		Type:       uploadReq.FileType,
		CreatedAt:  time.Now().Unix(),
	}

	saveErr := s.repo.SaveFile(ctx, file)
	if saveErr != nil {
		_ = s.doUploadFileRollback(fileName)
		return nil, saveErr
	}

	return file, nil
}

func (s *StorageServiceImpl) Fetch(req *GetFileReq) ([]*File, error) {
	files, err := s.repo.SearchFileByNameOrTags(req.Name, req.Tags)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(files); i++ {
		files[i].addLink()
	}

	return files, nil
}

func (s *StorageServiceImpl) doUploadFile(file *multipart.FileHeader, fileName string) error {

	// check file size
	if file.Size > s.maxFileSize {
		return FileSizeError
	}

	//check storage size
	if file.Size > s.maxStorageSize-s.usedStorageSize {
		return StorageLimit
	}

	tmpFile, err := file.Open()
	if err != nil {
		return err
	}
	defer func() { _ = tmpFile.Close() }()

	targetFile, err := os.OpenFile(s.basePath+fileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	defer func() { _ = targetFile.Close() }()

	wroteBytes, writeErr := io.Copy(targetFile, tmpFile)
	if writeErr != nil || wroteBytes != file.Size {
		return DiskError
	}

	return nil
}

func (s *StorageServiceImpl) doUploadFileRollback(fileName string) error {
	//TODO
	return nil
}
