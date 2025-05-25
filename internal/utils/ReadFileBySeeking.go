package utils

import (
	"io"
	"os"
)

type ReadFileBySeekingParams struct {
	Offset   int64
	Bytes    int64
	FilePath string
}

type ReadFileBySeekingResponse struct {
	Data   string
	Offset int64
}

func ReadFileBySeeking(params ReadFileBySeekingParams) (*ReadFileBySeekingResponse, error) {
	if params.Bytes == 0 {
		params.Bytes = 4096
	}

	file, err := os.Open(params.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Seek(params.Offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, params.Bytes)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}

	newOffset := params.Offset + int64(n)

	return &ReadFileBySeekingResponse{
		Data:   string(buf[:n]),
		Offset: newOffset,
	}, nil
}
