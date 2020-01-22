package zbox

import (
	"encoding/json"
	"fmt"

	"github.com/0chain/gosdk/zboxcore/blockchain"
	"github.com/0chain/gosdk/zboxcore/sdk"
)

type StatusCallback interface {
	sdk.StatusCallback
}

type Allocation struct {
	ID           string `json:"id"`
	DataShards   int    `json:"data_shards"`
	ParityShards int    `json:"parity_shards"`
	Size         int64  `json:"size"`
	Expiration   int64  `json:"expiration_date"`

	blobbers      []*blockchain.StorageNode
	sdkAllocation *sdk.Allocation
}

func (a *Allocation) ListDir(path string) (string, error) {
	listResult, err := a.sdkAllocation.ListDir(path)
	if err != nil {
		return "", err
	}
	retBytes, err := json.Marshal(listResult)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

func (a *Allocation) ListDirFromAuthTicket(authTicket string, lookupHash string) (string, error) {
	listResult, err := a.sdkAllocation.ListDirFromAuthTicket(authTicket, lookupHash)
	if err != nil {
		return "", err
	}
	retBytes, err := json.Marshal(listResult)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

func (a *Allocation) GetFileMeta(path string) (string, error) {
	fileMetaData, err := a.sdkAllocation.GetFileMeta(path)
	if err != nil {
		return "", err
	}
	retBytes, err := json.Marshal(fileMetaData)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

func (a *Allocation) GetFileMetaFromAuthTicket(authTicket string, lookupHash string) (string, error) {
	fileMetaData, err := a.sdkAllocation.GetFileMetaFromAuthTicket(authTicket, lookupHash)
	if err != nil {
		return "", err
	}
	retBytes, err := json.Marshal(fileMetaData)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

func (a *Allocation) DownloadFile(remotePath, localPath string, statusCb StatusCallback) error {
	return a.sdkAllocation.DownloadFile(localPath, remotePath, statusCb)
}

func (a *Allocation) DownloadThumbnail(remotePath, localPath string, statusCb StatusCallback) error {
	return a.sdkAllocation.DownloadThumbnail(localPath, remotePath, statusCb)
}

func (a *Allocation) UploadFile(localPath, remotePath string, statusCb StatusCallback) error {
	return a.sdkAllocation.UploadFile(localPath, remotePath, statusCb)
}

func (a *Allocation) UploadFileWithThumbnail(localPath, remotePath string, thumbnailpath string, statusCb StatusCallback) error {
	return a.sdkAllocation.UploadFileWithThumbnail(localPath, remotePath, thumbnailpath, statusCb)
}

func (a *Allocation) EncryptAndUploadFile(localPath, remotePath string, statusCb StatusCallback) error {
	return a.sdkAllocation.EncryptAndUploadFile(localPath, remotePath, statusCb)
}

func (a *Allocation) EncryptAndUploadFileWithThumbnail(localPath, remotePath string, thumbnailpath string, statusCb StatusCallback) error {
	return a.sdkAllocation.EncryptAndUploadFileWithThumbnail(localPath, remotePath, thumbnailpath, statusCb)
}

func (a *Allocation) UpdateFile(localPath, remotePath string, statusCb StatusCallback) error {
	return a.sdkAllocation.UpdateFile(localPath, remotePath, statusCb)
}

func (a *Allocation) UpdateFileWithThumbnail(localPath, remotePath string, thumbnailpath string, statusCb StatusCallback) error {
	return a.sdkAllocation.UpdateFileWithThumbnail(localPath, remotePath, thumbnailpath, statusCb)
}

func (a *Allocation) EncryptAndUpdateFile(localPath, remotePath string, statusCb StatusCallback) error {
	return a.sdkAllocation.EncryptAndUpdateFile(localPath, remotePath, statusCb)
}

func (a *Allocation) EncryptAndUpdateFileWithThumbnail(localPath, remotePath string, thumbnailpath string, statusCb StatusCallback) error {
	return a.sdkAllocation.EncryptAndUpdateFileWithThumbnail(localPath, remotePath, thumbnailpath, statusCb)
}

func (a *Allocation) DeleteFile(remotePath string) error {
	return a.sdkAllocation.DeleteFile(remotePath)
}

func (a *Allocation) RenameObject(remotePath string, destName string) error {
	return a.sdkAllocation.RenameObject(remotePath, destName)
}

func (a *Allocation) GetStats() (string, error) {
	stats := a.sdkAllocation.GetStats()
	retBytes, err := json.Marshal(stats)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

func (a *Allocation) GetShareAuthToken(path string, filename string, referenceType string, refereeClientID string) (string, error) {
	return a.sdkAllocation.GetAuthTicketForShare(path, filename, referenceType, refereeClientID)
}

func (a *Allocation) DownloadFromAuthTicket(localPath string, authTicket string, remoteLookupHash string, remoteFilename string, status StatusCallback) error {
	return a.sdkAllocation.DownloadFromAuthTicket(localPath, authTicket, remoteLookupHash, remoteFilename, status)
}

func (a *Allocation) DownloadThumbnailFromAuthTicket(localPath string, authTicket string, remoteLookupHash string, remoteFilename string, status StatusCallback) error {
	return a.sdkAllocation.DownloadThumbnailFromAuthTicket(localPath, authTicket, remoteLookupHash, remoteFilename, status)
}

func (a *Allocation) GetFileStats(path string) (string, error) {
	stats, err := a.sdkAllocation.GetFileStats(path)
	if err != nil {
		return "", err
	}
	result := make([]*sdk.FileStats, 0)
	for _, v := range stats {
		result = append(result, v)
	}
	retBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

func (a *Allocation) CancelDownload(remotepath string) error {
	return a.sdkAllocation.CancelDownload(remotepath)
}

func (a *Allocation) GetAllocationDiff(lastSyncCachePath string, localRootPath string, localFileFilters string, remoteExcludePaths string) (string, error) {
	var filterArray []string
	err := json.Unmarshal([]byte(localFileFilters), &filterArray)
	if err != nil {
		return "", fmt.Errorf("invalid local file filter JSON. %v", err)
	}
	var exclPathArray []string
	err = json.Unmarshal([]byte(remoteExcludePaths), &exclPathArray)
	if err != nil {
		return "", fmt.Errorf("invalid remote exclude path JSON. %v", err)
	}

	lFdiff, err := a.sdkAllocation.GetAllocationDiff(lastSyncCachePath, localRootPath, filterArray, exclPathArray)
	if err != nil {
		return "", fmt.Errorf("get allocation diff in sdk failed. %v", err)
	}
	retBytes, err := json.Marshal(lFdiff)
	if err != nil {
		return "", fmt.Errorf("failed to convert JSON. %v", err)
	}
	return string(retBytes), nil
}

func (a *Allocation) SaveRemoteSnapshot(pathToSave string, remoteExcludePaths string) error {
	var exclPathArray []string
	err := json.Unmarshal([]byte(remoteExcludePaths), &exclPathArray)
	if err != nil {
		return fmt.Errorf("invalid remote exclude path JSON. %v", err)
	}
	return a.sdkAllocation.SaveRemoteSnapshot(pathToSave, exclPathArray)
}

func (a *Allocation) CommitMetaTransaction(path, crudOperation string) (string, error) {
	metaTxnData, err := a.sdkAllocation.CommitMetaTransaction(path, crudOperation)
	if err != nil {
		return "", fmt.Errorf("Failed to commit metaTxn. %v", err)
	}
	retBytes, err := json.Marshal(metaTxnData)
	if err != nil {
		return "", fmt.Errorf("failed to convert JSON. %v", err)
	}
	return string(retBytes), nil
}
