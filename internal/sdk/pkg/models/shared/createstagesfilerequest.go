// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateStagesFileRequestFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

type CreateStagesFileRequest struct {
	// File to upload
	File *CreateStagesFileRequestFile `multipartForm:"file"`
}
