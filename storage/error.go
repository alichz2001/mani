package main

type ServiceError string

func (e ServiceError) Error() string {
	return string(e)
}

const (
	FileSizeError     ServiceError = "file size error"
	StorageLimit      ServiceError = "storage limit"
	DiskError         ServiceError = "upload file error"
	DuplicateFileName ServiceError = "duplicate file name"
	DatabaseError     ServiceError = "database error"

	UnknownError ServiceError = "Unknown error"
)
