package filestorageservice

import (
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

var Client *azblob.Client
var ContainerName = os.Getenv("AZURE_CONTAINER_NAME")

func ServiceClientSharedKey(accountName string, accountKey string) {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", accountName)

	client, err := azblob.NewClientWithSharedKeyCredential(accountURL, credential, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	Client = client
}
