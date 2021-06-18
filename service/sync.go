package service

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"yousyncclient/client"
	"yousyncclient/pb"
	"yousyncclient/utils"
)

const MaxChunkSize int64 = 1024 * 1024

func SyncFolder(syncPath string, folderId int64) error {
	err := filepath.Walk(syncPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(syncPath, path)
		if err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		nBytes, nChunks := int64(0), int64(0)
		r := bufio.NewReader(f)
		buf := make([]byte, 0, MaxChunkSize)
		for {
			n, err := r.Read(buf[:cap(buf)])
			buf = buf[:n]
			if n == 0 {
				if err == nil {
					continue
				}
				if err == io.EOF {
					break
				}
				return err
			}
			// check chunk
			checksum := utils.SHA256Checksum(buf)
			result, err := client.DefaultSyncClient.Client.CheckChunk(context.Background(), &pb.ChunkInfo{
				Index:    uint64(nChunks),
				Size:     uint64(len(buf)),
				Offset:   uint64(nChunks * MaxChunkSize),
				CheckSum: checksum,
				Path:     rel,
				FolderId: uint64(folderId),
			})
			if err != nil {
				return err
			}
			fmt.Println(fmt.Sprintf("%s chunk=%d same=%t", rel, nChunks, result.Success))
			if !result.Success {
				_, err = client.DefaultSyncClient.Client.SyncFileChunk(context.Background(), &pb.Chunk{
					Index:    uint64(nChunks),
					Size:     uint64(len(buf)),
					Offset:   uint64(nChunks * MaxChunkSize),
					Path:     rel,
					Data:     buf,
					FolderId: uint64(folderId),
					LastChunk: uint64(MaxChunkSize) >= uint64(len(buf)),
				})
				if err != nil {
					return err
				}
			}
			nChunks++
			nBytes += int64(len(buf))
			// process buf
			if err != nil && err != io.EOF {
				return err
			}
			// write chunk to remote

		}
		log.Println("Bytes:", nBytes, "Chunks:", nChunks)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
