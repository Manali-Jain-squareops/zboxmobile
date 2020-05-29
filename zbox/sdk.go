package zbox

import (
	"encoding/json"
	"math"
	"time"

	"github.com/0chain/gosdk/zboxcore/client"
	. "github.com/0chain/gosdk/zboxcore/logger"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zcncore"
)

type ChainConfig struct {
	ChainID           string   `json:"chain_id,omitempty"`
	PreferredBlobbers []string `json:"preferred_blobbers"`
	BlockWorker       string   `json:"block_worker"`
	SignatureScheme   string   `json:"signature_scheme"`
}

type StorageSDK struct {
	chainconfig *ChainConfig
	client      *client.Client
}

func SetLogFile(logFile string, verbose bool) {
	zcncore.SetLogFile(logFile, verbose)
	sdk.SetLogFile(logFile, verbose)
}

// SetLogLevel set the log level.
// lvl - 0 disabled; higher number (upto 4) more verbosity
func SetLogLevel(logLevel int) {
	zcncore.SetLogLevel(logLevel)
	sdk.SetLogLevel(logLevel)
}

func InitStorageSDK(clientjson string, configjson string) (*StorageSDK, error) {
	configObj := &ChainConfig{}
	err := json.Unmarshal([]byte(configjson), configObj)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	err = zcncore.InitZCNSDK(configObj.BlockWorker, configObj.SignatureScheme)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	err = sdk.InitStorageSDK(clientjson, configObj.BlockWorker, configObj.ChainID, configObj.SignatureScheme, configObj.PreferredBlobbers)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	Logger.Info("Init successful")
	return &StorageSDK{client: client.GetClient(), chainconfig: configObj}, nil
}

func (s *StorageSDK) CreateAllocation(datashards int, parityshards int, size, expiration, lock int64) (*Allocation, error) {
	readPrice := sdk.PriceRange{Min: 0, Max: math.MaxInt64}
	writePrice := sdk.PriceRange{Min: 0, Max: math.MaxInt64}
	sdkAllocationID, err := sdk.CreateAllocation(datashards, parityshards, size, expiration, readPrice, writePrice, 1*time.Hour, lock)
	if err != nil {
		return nil, err
	}
	sdkAllocation, err := sdk.GetAllocation(sdkAllocationID)
	if err != nil {
		return nil, err
	}
	return &Allocation{ID: sdkAllocation.ID, DataShards: sdkAllocation.DataShards, ParityShards: sdkAllocation.ParityShards, Size: sdkAllocation.Size, Expiration: sdkAllocation.Expiration, blobbers: sdkAllocation.Blobbers, sdkAllocation: sdkAllocation}, nil
}

func (s *StorageSDK) GetAllocation(allocationID string) (*Allocation, error) {
	sdkAllocation, err := sdk.GetAllocation(allocationID)
	if err != nil {
		return nil, err
	}
	return &Allocation{ID: sdkAllocation.ID, DataShards: sdkAllocation.DataShards, ParityShards: sdkAllocation.ParityShards, Size: sdkAllocation.Size, Expiration: sdkAllocation.Expiration, blobbers: sdkAllocation.Blobbers, sdkAllocation: sdkAllocation}, nil
}

func (s *StorageSDK) GetAllocations() (string, error) {
	sdkAllocations, err := sdk.GetAllocations()
	if err != nil {
		return "", err
	}
	result := make([]*Allocation, len(sdkAllocations))
	for i, sdkAllocation := range sdkAllocations {
		allocationObj := &Allocation{ID: sdkAllocation.ID, DataShards: sdkAllocation.DataShards, ParityShards: sdkAllocation.ParityShards, Size: sdkAllocation.Size, Expiration: sdkAllocation.Expiration, blobbers: sdkAllocation.Blobbers, sdkAllocation: sdkAllocation}
		result[i] = allocationObj
	}
	retBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

func (s *StorageSDK) GetAllocationFromAuthTicket(authTicket string) (*Allocation, error) {
	sdkAllocation, err := sdk.GetAllocationFromAuthTicket(authTicket)
	if err != nil {
		return nil, err
	}
	return &Allocation{ID: sdkAllocation.ID, DataShards: sdkAllocation.DataShards, ParityShards: sdkAllocation.ParityShards, Size: sdkAllocation.Size, Expiration: sdkAllocation.Expiration, blobbers: sdkAllocation.Blobbers, sdkAllocation: sdkAllocation}, nil
}

func (s *StorageSDK) GetAllocationStats(allocationID string) (string, error) {
	allocationObj, err := sdk.GetAllocation(allocationID)
	if err != nil {
		return "", err
	}
	stats := allocationObj.GetStats()
	retBytes, err := json.Marshal(stats)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

func (s *StorageSDK) FinalizeAllocation(allocationID string) (string, error) {
	return sdk.FinalizeAllocation(allocationID)
}

func (s *StorageSDK) CancelAlloctioan(allocationID string) (string, error) {
	return sdk.CancelAlloctioan(allocationID)
}

// READ POOL METHODS

//CreateReadPool is to create read pool for the wallet
func (s *StorageSDK) CreateReadPool() error {
	return sdk.CreateReadPool()
}

//GetReadPoolInfo is to get information about the read pool for the allocation
func (s *StorageSDK) GetReadPoolInfo(allocID string) (string, error) {
	readPool, err := sdk.GetReadPoolInfo("")
	if err != nil {
		return "", err
	}

	if len(allocID) > 0 {
		readPool.AllocFilter(allocID)
	}
	retBytes, err := json.Marshal(readPool)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

//ReadPoolLock is to lock tokens into the read pool
func (s *StorageSDK) ReadPoolLock(durInSeconds int64, tokens, fee float64, allocID, blobberID string) error {
	var duration time.Duration
	duration = time.Duration(durInSeconds) * time.Second
	return sdk.ReadPoolLock(duration, allocID, blobberID, zcncore.ConvertToValue(tokens), zcncore.ConvertToValue(fee))
}

//ReadPoolUnlock is to unlock tokens from read pool
func (s *StorageSDK) ReadPoolUnlock(poolID string, fee float64) error {
	return sdk.ReadPoolUnlock(poolID, zcncore.ConvertToValue(fee))
}

// WRITE POOL METHODS

//GetWritePoolInfo is to get information about the write pool for the allocation
func (s *StorageSDK) GetWritePoolInfo(allocID string) (string, error) {
	writePool, err := sdk.GetWritePoolInfo("")
	if err != nil {
		return "", err
	}
	if len(allocID) > 0 {
		writePool.AllocFilter(allocID)
	}
	retBytes, err := json.Marshal(writePool)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

//WritePoolLock is to lock tokens into the write pool
func (s *StorageSDK) WritePoolLock(durInSeconds int64, tokens, fee float64, allocID, blobberID string) error {
	var duration time.Duration
	duration = time.Duration(durInSeconds) * time.Second
	return sdk.WritePoolLock(duration, allocID, blobberID, zcncore.ConvertToValue(tokens), zcncore.ConvertToValue(fee))
}

//WritePoolUnlock is to unlock tokens from write pool
func (s *StorageSDK) WritePoolUnlock(poolID string, fee float64) error {
	return sdk.WritePoolUnlock(poolID, zcncore.ConvertToValue(fee))
}
