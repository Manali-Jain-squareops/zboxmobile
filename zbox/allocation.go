package zbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"sort"
	"sync"
	"time"

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
func (a *Allocation) GetAuthToken(path string, filename string, referenceType string, refereeClientID string, refereeEncryptionPublicKey string, expiration int64) (string, error) {
	return a.sdkAllocation.GetAuthTicket(path, filename, referenceType, refereeClientID, refereeEncryptionPublicKey, expiration)
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

func (a *Allocation) GetFirstSegment(localPath, remotePath string) error {
	if len(remotePath) == 0 {
		return errors.New("Error: remotepath / authticket flag is missing")
	}

	if len(localPath) == 0 {
		return errors.New("Error: localpath is missing")
	}

	dir := path.Dir(localPath)

	file, err := os.Create(localPath)

	if err != nil {
		return err
	}

	downloader := &M3u8Downloader{
		localDir:      dir,
		localPath:     localPath,
		remotePath:    remotePath,
		authTicket:    "",
		allocationID:  a.ID,
		rxPay:         false,
		downloadQueue: make(chan MediaItem, 100),
		playlist:      sdk.NewMediaPlaylist(5, file),
		done:          make(chan error, 1),
	}

	downloader.allocationObj = a.sdkAllocation

	listResult, err := a.sdkAllocation.ListDir(downloader.remotePath)
	list := listResult.Children
	if err != nil {
		return err
	}

	downloader.Lock()
	n := len(downloader.items)
	max := len(list)
	fmt.Println("max")
	fmt.Println(max)
	if n < max {
		sort.Sort(SortedListResult(list))

		for i := n; i < max; i++ {
			item := MediaItem{
				Name: list[i].Name,
				Path: list[i].Path,
			}
			downloader.items = append(downloader.items, item)
			downloader.playlist.Append(item.Name)
			downloader.playlist.Wait = append(downloader.playlist.Wait, item.Name)

			downloader.addToDownload(item)
		}
	}
	downloader.Unlock()

	downloader.playlist.Writer.Truncate(0)
	downloader.playlist.Writer.Seek(0, 0)
	downloader.playlist.Writer.Write(downloader.playlist.Encode())
	downloader.playlist.Writer.Sync()

	return nil
}

// GetMinStorageCost - getting back min cost for allocation
func (a *Allocation) PlayStreaming(localPath, remotePath, authTicket, lookupHash string, delay int, statusCb StatusCallback) error {
	downloader, err := createM3u8Downloader(localPath, remotePath, authTicket, a.ID, lookupHash, false, delay)
	if err != nil {
		return err
	}

	downloader.status = statusCb
	return downloader.Start()
}

// M3u8Downloader download files from blobber's dir, and build them into a local m3u8 playlist
type M3u8Downloader struct {
	sync.RWMutex
	delay int

	localDir     string
	localPath    string
	remotePath   string
	authTicket   string
	allocationID string
	rxPay        bool

	allocationObj *sdk.Allocation

	lookupHash    string
	items         []MediaItem
	downloadQueue chan MediaItem
	playlist      *sdk.MediaPlaylist
	done          chan error
	status        StatusCallback
}

func createM3u8Downloader(localPath, remotePath, authTicket, allocationID, lookupHash string, rxPay bool, delay int) (*M3u8Downloader, error) {
	if len(remotePath) == 0 && (len(authTicket) == 0) {
		return nil, errors.New("Error: remotepath / authticket flag is missing")
	}

	if len(localPath) == 0 {
		return nil, errors.New("Error: localpath is missing")
	}
	dir := path.Dir(localPath)

	file, err := os.Create(localPath)

	if err != nil {
		return nil, err
	}

	downloader := &M3u8Downloader{
		localDir:      dir,
		localPath:     localPath,
		remotePath:    remotePath,
		authTicket:    authTicket,
		allocationID:  allocationID,
		rxPay:         rxPay,
		downloadQueue: make(chan MediaItem, 100),
		playlist:      sdk.NewMediaPlaylist(delay, file),
		done:          make(chan error, 1),
	}

	if len(remotePath) > 0 {
		if len(allocationID) == 0 { // check if the flag "path" is set
			return nil, errors.New("Error: allocation flag is missing") // If not, we'll let the user know
		}

		allocationObj, err := sdk.GetAllocation(allocationID)

		if err != nil {
			return nil, fmt.Errorf("Error fetching the allocation: %s", err)
		}

		downloader.allocationObj = allocationObj

	} else if len(authTicket) > 0 {
		allocationObj, err := sdk.GetAllocationFromAuthTicket(authTicket)
		if err != nil {
			return nil, fmt.Errorf("Error fetching the allocation: %s", err)
		}

		downloader.allocationObj = allocationObj

		at := sdk.InitAuthTicket(authTicket)
		isDir, err := at.IsDir()
		if isDir && len(lookupHash) == 0 {
			lookupHash, err = at.GetLookupHash()
			if err != nil {
				return nil, fmt.Errorf("Error getting the lookuphash from authticket: %s", err)
			}

			downloader.lookupHash = lookupHash
		}
		if !isDir {
			return nil, fmt.Errorf("Invalid operation. Auth ticket is not for a directory: %s", err)
		}

	}

	return downloader, nil
}

// Start start to download ,and build playlist
func (d *M3u8Downloader) Start() error {
	if d.status != nil {
		d.status.Started(d.allocationID, d.localPath, 0, 0)
	}

	go d.autoDownload()
	go d.autoRefreshList()
	go d.playlist.Play()

	err := <-d.done

	return err
}

func (d *M3u8Downloader) addToDownload(item MediaItem) {
	d.downloadQueue <- item
}

func (d *M3u8Downloader) autoDownload() {
	//for {
	item := <-d.downloadQueue
	d.playlist.Append(item.Path)

	/*
		for i := 0; i < 3; i++ {
			if path, err := d.download(item); err == nil {
				d.playlist.Append(path)
				if d.status != nil {
					d.status.InProgress(d.allocationID, path, 1, len(d.items), nil)
				}
				break
			}
		}*/
	//}
}

func (d *M3u8Downloader) autoRefreshList() {
	for {
		list, err := d.getList()
		if err != nil {
			continue
		}

		d.Lock()
		n := len(d.items)
		max := len(list)
		if n < max {
			sort.Sort(SortedListResult(list))

			for i := n; i < max; i++ {
				item := MediaItem{
					Name: list[i].Name,
					Path: list[i].Path,
				}
				d.items = append(d.items, item)
				d.addToDownload(item)
			}
		}
		d.Unlock()

		time.Sleep(5 * time.Second)
	}
}

/*
func (d *M3u8Downloader) download(item MediaItem) (string, error) {
	wg := &sync.WaitGroup{}
	statusBar := &StatusBarMocked{wg: wg}
	wg.Add(1)

	localPath := d.localDir + string(os.PathSeparator) + item.Name
	remotePath := item.Path

	if len(d.remotePath) > 0 {
		err := d.allocationObj.DownloadFile(localPath, remotePath, statusBar)
		if err != nil {
			return "", err
		}
		wg.Wait()
	}

	if !statusBar.success {
		return "", statusBar.err
	}

	return localPath, nil
}*/

func (d *M3u8Downloader) getList() ([]*sdk.ListResult, error) {
	//get list from remote allocations's path
	if len(d.remotePath) > 0 {
		ref, err := d.allocationObj.ListDir(d.remotePath)
		if err != nil {
			return nil, err
		}
		return ref.Children, nil
	}

	//get list from authticket
	if len(d.authTicket) > 0 {
		ref, err := d.allocationObj.ListDirFromAuthTicket(d.authTicket, d.lookupHash)
		if err != nil {
			return nil, err
		}
		return ref.Children, nil
	}

	return nil, fmt.Errorf("file not found")
}

// SortedListResult sort files order by name
type SortedListResult []*sdk.ListResult

func (a SortedListResult) Len() int {
	return len(a)
}
func (a SortedListResult) Less(i, j int) bool {

	l := a[i]
	r := a[j]

	if len(l.Name) < len(r.Name) {
		return true
	}

	if len(l.Name) > len(r.Name) {
		return false
	}

	return l.Name < r.Name
}

func (a SortedListResult) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// MediaItem ts file
type MediaItem struct {
	Name string
	Path string
}

type StatusBarMocked struct {
	wg      *sync.WaitGroup
	success bool
	err     error
}

func (s *StatusBarMocked) Started(allocationId, filePath string, op int, totalBytes int) {}

func (s *StatusBarMocked) InProgress(allocationId, filePath string, op int, completedBytes int, data []byte) {
}

func (s *StatusBarMocked) Error(allocationID string, filePath string, op int, err error) {
	s.success = false
	s.err = err
	s.wg.Done()
}

func (s *StatusBarMocked) Completed(allocationId, filePath string, filename string, mimetype string, size int, op int) {
	s.success = true
	s.wg.Done()
}

func (s *StatusBarMocked) CommitMetaCompleted(request, response string, err error) {}

func (s *StatusBarMocked) RepairCompleted(filesRepaired int) {}

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
