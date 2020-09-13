package mq

import (
	"cloudDisk/common"
)

type TransferData struct{
	FileHash string
	CurLocation string
	Destination string
	DestStoreType common.StoreType
}