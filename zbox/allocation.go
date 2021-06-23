package zbox

import (
	"encoding/json"
	"fmt"

	"github.com/0chain/gosdk/zboxcore/blockchain"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/sdk"
)

// StatusCallback - callback for file operations
type StatusCallback interface {
	sdk.StatusCallback
}

// Allocation - structure for allocation object
type Allocation struct {
	ID           string `json:"id"`
	DataShards   int    `json:"data_shards"`
	ParityShards int    `json:"parity_shards"`
	Size         int64  `json:"size"`
	Expiration   int64  `json:"expiration_date"`

	blobbers      []*blockchain.StorageNode
	sdkAllocation *sdk.Allocation
}

// MinMaxCost - keeps cost for allocation update/creation
type MinMaxCost struct {
	minW float64
	minR float64
	maxW float64
	maxR float64
}

// ListDir - listing files from path
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

// ListDirFromAuthTicket - listing files from path with auth ticket
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

// GetFileMeta - getting file meta details from file path
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

// GetFileMetaFromAuthTicket - getting file meta details from file path and auth ticket
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

// DownloadFile - start download file from remote path to localpath
func (a *Allocation) DownloadFile(remotePath, localPath string, statusCb StatusCallback) error {
	return a.sdkAllocation.DownloadFile(localPath, remotePath, statusCb)
}

// DownloadFileByBlock - start download file from remote path to localpath by blocks number
func (a *Allocation) DownloadFileByBlock(remotePath, localPath string, startBlock, endBlock int64, numBlocks int, statusCb StatusCallback) error {
	return a.sdkAllocation.DownloadFileByBlock(localPath, remotePath, startBlock, endBlock, numBlocks, statusCb)
}

// DownloadThumbnail - start download file thumbnail from remote path to localpath
func (a *Allocation) DownloadThumbnail(remotePath, localPath string, statusCb StatusCallback) error {
	return a.sdkAllocation.DownloadThumbnail(localPath, remotePath, statusCb)
}

// UploadFile - start upload file thumbnail from localpath to remote path
func (a *Allocation) UploadFile(localPath, remotePath, fileAttrs string, statusCb StatusCallback) error {
	var attrs fileref.Attributes
	if len(fileAttrs) > 0 {
		err := json.Unmarshal([]byte(fileAttrs), &attrs)
		if err != nil {
			return fmt.Errorf("failed to convert fileAttrs. %v", err)
		}
	}
	return a.sdkAllocation.UploadFile(localPath, remotePath, attrs, statusCb)
}

// RepairFile - repairing file if it's exist in remote path
func (a *Allocation) RepairFile(localPath, remotePath string, statusCb StatusCallback) error {
	return a.sdkAllocation.RepairFile(localPath, remotePath, statusCb)
}

// UploadFileWithThumbnail - start upload file with thumbnail
func (a *Allocation) UploadFileWithThumbnail(localPath, remotePath, fileAttrs string, thumbnailpath string, statusCb StatusCallback) error {
	var attrs fileref.Attributes
	if len(fileAttrs) > 0 {
		err := json.Unmarshal([]byte(fileAttrs), &attrs)
		if err != nil {
			return fmt.Errorf("failed to convert fileAttrs. %v", err)
		}
	}
	return a.sdkAllocation.UploadFileWithThumbnail(localPath, remotePath, thumbnailpath, attrs, statusCb)
}

// EncryptAndUploadFile - start upload encrypted file
func (a *Allocation) EncryptAndUploadFile(localPath, remotePath, fileAttrs string, statusCb StatusCallback) error {
	var attrs fileref.Attributes
	if len(fileAttrs) > 0 {
		err := json.Unmarshal([]byte(fileAttrs), &attrs)
		if err != nil {
			return fmt.Errorf("failed to convert fileAttrs. %v", err)
		}
	}
	return a.sdkAllocation.EncryptAndUploadFile(localPath, remotePath, attrs, statusCb)
}

// EncryptAndUploadFileWithThumbnail - start upload encrypted file with thumbnail
func (a *Allocation) EncryptAndUploadFileWithThumbnail(localPath, remotePath, fileAttrs string, thumbnailpath string, statusCb StatusCallback) error {
	var attrs fileref.Attributes
	if len(fileAttrs) > 0 {
		err := json.Unmarshal([]byte(fileAttrs), &attrs)
		if err != nil {
			return fmt.Errorf("failed to convert fileAttrs. %v", err)
		}
	}
	return a.sdkAllocation.EncryptAndUploadFileWithThumbnail(localPath, remotePath, thumbnailpath, attrs, statusCb)
}

// UpdateFile - update file from local path to remote path
func (a *Allocation) UpdateFile(localPath, remotePath, fileAttrs string, statusCb StatusCallback) error {
	var attrs fileref.Attributes
	if len(fileAttrs) > 0 {
		err := json.Unmarshal([]byte(fileAttrs), &attrs)
		if err != nil {
			return fmt.Errorf("failed to convert fileAttrs. %v", err)
		}
	}
	return a.sdkAllocation.UpdateFile(localPath, remotePath, attrs, statusCb)
}

// UpdateFileWithThumbnail - update file from local path to remote path with Thumbnail
func (a *Allocation) UpdateFileWithThumbnail(localPath, remotePath, fileAttrs string, thumbnailpath string, statusCb StatusCallback) error {
	var attrs fileref.Attributes
	if len(fileAttrs) > 0 {
		err := json.Unmarshal([]byte(fileAttrs), &attrs)
		if err != nil {
			return fmt.Errorf("failed to convert fileAttrs. %v", err)
		}
	}
	return a.sdkAllocation.UpdateFileWithThumbnail(localPath, remotePath, thumbnailpath, attrs, statusCb)
}

// EncryptAndUpdateFile - update file from local path to remote path from encrypted folder
func (a *Allocation) EncryptAndUpdateFile(localPath, remotePath, fileAttrs string, statusCb StatusCallback) error {
	var attrs fileref.Attributes
	if len(fileAttrs) > 0 {
		err := json.Unmarshal([]byte(fileAttrs), &attrs)
		if err != nil {
			return fmt.Errorf("failed to convert fileAttrs. %v", err)
		}
	}
	return a.sdkAllocation.EncryptAndUpdateFile(localPath, remotePath, attrs, statusCb)
}

// EncryptAndUpdateFileWithThumbnail - update file from local path to remote path from encrypted folder with Thumbnail
func (a *Allocation) EncryptAndUpdateFileWithThumbnail(localPath, remotePath, fileAttrs string, thumbnailpath string, statusCb StatusCallback) error {
	var attrs fileref.Attributes
	if len(fileAttrs) > 0 {
		err := json.Unmarshal([]byte(fileAttrs), &attrs)
		if err != nil {
			return fmt.Errorf("failed to convert fileAttrs. %v", err)
		}
	}
	return a.sdkAllocation.EncryptAndUpdateFileWithThumbnail(localPath, remotePath, thumbnailpath, attrs, statusCb)
}

// DeleteFile - delete file from remote path
func (a *Allocation) DeleteFile(remotePath string) error {
	return a.sdkAllocation.DeleteFile(remotePath)
}

// RenameObject - rename or move file
func (a *Allocation) RenameObject(remotePath string, destName string) error {
	return a.sdkAllocation.RenameObject(remotePath, destName)
}

// GetStats - get allocation stats
func (a *Allocation) GetStats() (string, error) {
	stats := a.sdkAllocation.GetStats()
	retBytes, err := json.Marshal(stats)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

// GetBlobberStats - get blobbers stats
func (a *Allocation) GetBlobberStats() (string, error) {
	stats := a.sdkAllocation.GetBlobberStats()
	retBytes, err := json.Marshal(stats)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

// GetShareAuthToken - get auth ticket from refereeClientID
func (a *Allocation) GetShareAuthToken(path string, filename string, referenceType string, refereeClientID string) (string, error) {
	return a.sdkAllocation.GetAuthTicketForShare(path, filename, referenceType, refereeClientID)
}

// GetAuthToken - get auth token from refereeClientID
func (a *Allocation) GetAuthToken(path string, filename string, referenceType string, refereeClientID string, refereeEncryptionPublicKey string) (string, error) {
	return a.sdkAllocation.GetAuthTicket(path, filename, referenceType, refereeClientID, refereeEncryptionPublicKey)
}

// DownloadFromAuthTicket - download file from Auth ticket
func (a *Allocation) DownloadFromAuthTicket(localPath string, authTicket string, remoteLookupHash string, remoteFilename string, rxPay bool, status StatusCallback) error {
	return a.sdkAllocation.DownloadFromAuthTicket(localPath, authTicket, remoteLookupHash, remoteFilename, rxPay, status)
}

// DownloadFromAuthTicketByBlocks - download file from Auth ticket by blocks number
func (a *Allocation) DownloadFromAuthTicketByBlocks(localPath string, authTicket string, startBlock, endBlock int64, numBlocks int, remoteLookupHash string, remoteFilename string, rxPay bool, status StatusCallback) error {
	return a.sdkAllocation.DownloadFromAuthTicketByBlocks(localPath, authTicket, startBlock, endBlock, numBlocks, remoteLookupHash, remoteFilename, rxPay, status)
}

// DownloadThumbnailFromAuthTicket - downloadThumbnail from Auth ticket
func (a *Allocation) DownloadThumbnailFromAuthTicket(localPath string, authTicket string, remoteLookupHash string, remoteFilename string, rxPay bool, status StatusCallback) error {
	return a.sdkAllocation.DownloadThumbnailFromAuthTicket(localPath, authTicket, remoteLookupHash, remoteFilename, rxPay, status)
}

// GetFileStats - get file stats from path
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

// CancelDownload - cancel file download
func (a *Allocation) CancelDownload(remotepath string) error {
	return a.sdkAllocation.CancelDownload(remotepath)
}

// CancelUpload - cancel file upload
func (a *Allocation) CancelUpload(localpath string) error {
	return a.sdkAllocation.CancelUpload(localpath)
}

// GetDiff - cancel file diff
func (a *Allocation) GetDiff(lastSyncCachePath string, localRootPath string, localFileFilters string, remoteExcludePaths string) (string, error) {
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

// SaveRemoteSnapshot - saving remote snapshot
func (a *Allocation) SaveRemoteSnapshot(pathToSave string, remoteExcludePaths string) error {
	var exclPathArray []string
	err := json.Unmarshal([]byte(remoteExcludePaths), &exclPathArray)
	if err != nil {
		return fmt.Errorf("invalid remote exclude path JSON. %v", err)
	}
	return a.sdkAllocation.SaveRemoteSnapshot(pathToSave, exclPathArray)
}

// CommitMetaTransaction - authTicket - Optional, Only when you do download using authTicket and lookUpHash.
// lookupHash - Same as above.
// fileMeta - Optional, Only when you do delete and have already fetched fileMeta before delete operation.
func (a *Allocation) CommitMetaTransaction(path, crudOperation, authTicket, lookupHash, fileMeta string, statusCb StatusCallback) error {
	var fileMetaData *sdk.ConsolidatedFileMeta
	if len(fileMeta) > 0 {
		err := json.Unmarshal([]byte(fileMeta), fileMetaData)
		if err != nil {
			return fmt.Errorf("failed to convert fileMeta. %v", err)
		}
	}
	return a.sdkAllocation.CommitMetaTransaction(path, crudOperation, authTicket, lookupHash, fileMetaData, statusCb)
}

// StartRepair - start repair files from path
func (a *Allocation) StartRepair(localRootPath, pathToRepair string, statusCb StatusCallback) error {
	return a.sdkAllocation.StartRepair(localRootPath, pathToRepair, statusCb)
}

// CancelRepair - cancel repair files from path
func (a *Allocation) CancelRepair() error {
	return a.sdkAllocation.CancelRepair()
}

// CopyObject - copy object from path to dest
func (a *Allocation) CopyObject(path string, destPath string) error {
	return a.sdkAllocation.CopyObject(path, destPath)
}

// MoveObject - move object from path to dest
func (a *Allocation) MoveObject(path string, destPath string) error {
	return a.sdkAllocation.MoveObject(path, destPath)
}

// GetMinWriteRead - getting back cost for allocation
func (a *Allocation) GetMinWriteRead() (string, error) {
	minW, minR, err := a.sdkAllocation.GetMinWriteRead()
	maxW, maxR, err := a.sdkAllocation.GetMaxWriteRead()

	minMaxCost := &MinMaxCost{}
	minMaxCost.maxR = maxR
	minMaxCost.maxW = maxW
	minMaxCost.minR = minR
	minMaxCost.minW = minW

	retBytes, err := json.Marshal(minMaxCost)
	if err != nil {
		return "", fmt.Errorf("failed to convert JSON. %v", err)
	}

	return string(retBytes), nil
}

// GetMaxStorageCost - getting back max cost for allocation
func (a *Allocation) GetMaxStorageCost(size int64) (string, error) {
	cost, err := a.sdkAllocation.GetMaxStorageCost(size)
	return fmt.Sprintf("%f", cost), err
}

// GetMinStorageCost - getting back min cost for allocation
func (a *Allocation) GetMinStorageCost(size int64) (string, error) {
	cost, err := a.sdkAllocation.GetMinStorageCost(size)
	return fmt.Sprintf("%f", cost), err
}

// GetMaxStorageCostWithBlobbers - getting cost for listed blobbers
func (a *Allocation) GetMaxStorageCostWithBlobbers(size int64, blobbersJson string) (string, error) {
	var selBlobbers *[]*sdk.BlobberAllocation
	err := json.Unmarshal([]byte(blobbersJson), selBlobbers)
	if err != nil {
		return "", err
	}

	cost, err := a.sdkAllocation.GetMaxStorageCostFromBlobbers(size, *selBlobbers)
	return fmt.Sprintf("%f", cost), err
}
