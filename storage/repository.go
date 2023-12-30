package main

import (
	"context"
	"errors"
	"fmt"
)

func (db *ManiDB) SaveFile(ctx context.Context, file *File) error {

	docMeta, err := db.StorageCollection.CreateDocument(ctx, file)
	if err != nil {
		return errors.New("save file error")
	}

	file.ID = docMeta.Key

	return nil
}

func (db *ManiDB) SearchFileByNameOrTags(name string, tags []string) ([]*File, error) {
	aql := `
LET matched_docs = (                                          
    FOR i IN files                                                  
    FILTER i.status != -1                                            
    FILTER i.name == "%s" OR LENGTH(INTERSECTION(%s, i.tags)) > 0
    RETURN i
)
LET result = (
    (LENGTH(matched_docs) != 0)
    ? (matched_docs)
    : (
        FOR k IN files
        FILTER k.status != -1
        SORT k.created_at
        LIMIT 1
        RETURN k
      )
)
FOR j IN result
UPDATE j WITH { status: -1 } IN files
RETURN j`

	q := fmt.Sprintf(aql, name, stringifyListForAQL(tags))

	ctx := context.Background()
	cursor, err := db.Database.Query(ctx, q, nil)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close() }()

	files := make([]*File, cursor.Count())

	var tmp *File

	for cursor.HasMore() {
		tmp = &File{}
		_, _ = cursor.ReadDocument(ctx, tmp)
		files = append(files, tmp)
	}

	return files, nil
}
