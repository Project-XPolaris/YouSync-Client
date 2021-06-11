package api

import (
	"github.com/allentom/haruka"
	"net/http"
	"yousyncclient/service"
)

type SyncFolderRequestBody struct {
	Path     string `json:"path"`
	FolderId int64  `json:"folderId"`
}

var syncFileHandler haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody SyncFolderRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	err = service.SyncFolder(requestBody.Path, requestBody.FolderId)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	context.JSON(haruka.JSON{
		"success": true,
	})
}
