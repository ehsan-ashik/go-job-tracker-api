package filestorageservice

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

func UploadBlobFile(blobName string, blobData []byte) (string, error) {
	_, err := Client.UploadBuffer(context.TODO(), ContainerName, blobName, blobData, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", Client.URL(), ContainerName, blobName), nil
}

func DeleteBlobFile(blobName string) error {
	_, err := Client.DeleteBlob(context.TODO(), ContainerName, blobName, nil)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfBlobExists(blobName string) bool {
	pager := Client.NewListBlobsFlatPager(ContainerName, &container.ListBlobsFlatOptions{
		Prefix: &blobName,
	})

	for pager.More() {
		resp, _ := pager.NextPage(context.TODO())
		if len(resp.Segment.BlobItems) == 0 {
			return false
		}
	}
	return true
}
