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
	Miners            []string `json:"miners"`
	Sharders          []string `json:"sharders"`
	PreferredBlobbers []string `json:"preferred_blobbers"`
	SignatureScheme   string   `json:"signaturescheme"`
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
	err = zcncore.InitZCNSDK(configObj.Miners, configObj.Sharders, configObj.SignatureScheme)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	err = sdk.InitStorageSDK(clientjson, configObj.Miners, configObj.Sharders, configObj.ChainID, configObj.SignatureScheme, configObj.PreferredBlobbers)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	Logger.Info("Init successful")
	return &StorageSDK{client: client.GetClient(), chainconfig: configObj}, nil
}

func (s *StorageSDK) CreateAllocation(datashards int, parityshards int, size int64, expiration int64) (*Allocation, error) {
	readPrice := sdk.PriceRange{Min: 0, Max: math.MaxInt64}
	writePrice := sdk.PriceRange{Min: 0, Max: math.MaxInt64}
	sdkAllocationID, err := sdk.CreateAllocation(datashards, parityshards, size, expiration, readPrice, writePrice, 0)
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

// READ POOL METHODS

//CreateReadPool is to create read pool for the wallet
func (s *StorageSDK) CreateReadPool() error {
	return sdk.CreateReadPool()
}

//GetReadPoolInfo is to get information about the read pool for the client
func (s *StorageSDK) GetReadPoolInfo(clientID string) (string, error) {
	readPool, err := sdk.GetReadPoolInfo(clientID)
	if err != nil {
		return "", err
	}

	retBytes, err := json.Marshal(readPool)
	if err != nil {
		return "", err
	}
	return string(retBytes), nil
}

//ReadPoolLock is to lock tokens into the read pool
func (s *StorageSDK) ReadPoolLock(durInSeconds, tokens, fee int64, allocID string) error {
	var duration time.Duration
	duration = time.Duration(durInSeconds) * time.Second
	return sdk.ReadPoolLock(duration, allocID, "", tokens, fee)
}

//ReadPoolUnlock is to unlock tokens from read pool
func (s *StorageSDK) ReadPoolUnlock(poolID string, fee int64) error {
	return sdk.ReadPoolUnlock(poolID, fee)
}
