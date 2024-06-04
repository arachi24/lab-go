package azureblob

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/spf13/viper"
)

func Upload(content []byte, contentType string) (*string, error) {
	reader := bytes.NewReader(content)
	blobName := contentType
	client, err := ConnectAzure()
	if err != nil {
		return nil, err
	}

	containerName := viper.GetString("BLOB_CONTAINER")
	uploadResp, err := client.UploadStream(context.TODO(),
		containerName,
		blobName,
		reader,
		&azblob.UploadStreamOptions{})
	if err != nil {
		return nil, err
	}

	url := client.URL() + containerName + "/" + blobName
	_ = uploadResp

	return &url, nil
}

func ConnectAzure() (*azblob.Client, error) {
	accountName := viper.GetString("BLOB_ACCOUNT_NAME")
	accountKey := viper.GetString("BLOB_ACCOUNT_KEY")
	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, err
	}
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", accountName), cred, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
