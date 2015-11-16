// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package databaseOverlay

import (
	"github.com/FactomProject/factomd/common/directoryBlock"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// ProcessDBlockBatche inserts the DBlock and update all it's dbentries in DB
func (db *Overlay) ProcessDBlockBatch(dblock interfaces.DatabaseBatchable) error {
	return db.ProcessBlockBatch([]byte{byte(DIRECTORYBLOCK)},
		[]byte{byte(DIRECTORYBLOCK_NUMBER)},
		[]byte{byte(DIRECTORYBLOCK_KEYMR)}, dblock)
}

// FetchHeightRange looks up a range of blocks by the start and ending
// heights.  Fetch is inclusive of the start height and exclusive of the
// ending height. To fetch all hashes from the start height until no
// more are present, use -1 as endHeight.
func (db *Overlay) FetchHeightRange(startHeight, endHeight int64) ([]interfaces.IHash, error) {
	return db.FetchBlockIndexesInHeightRange([]byte{byte(DIRECTORYBLOCK_NUMBER)}, startHeight, endHeight)
}

// FetchBlockHeightBySha returns the block height for the given hash.  This is
// part of the database.Db interface implementation.
func (db *Overlay) FetchBlockHeightByKeyMR(sha interfaces.IHash) (int64, error) {
	dblk, err := db.FetchDBlockByKeyMR(sha)
	if err != nil {
		return -1, err
	}

	var height int64 = -1
	if dblk != nil {
		height = int64(dblk.GetDatabaseHeight())
	}

	return height, nil
}

// FetchDBlock gets an entry by hash from the database.
func (db *Overlay) FetchDBlockByKeyMR(keyMR interfaces.IHash) (interfaces.IDirectoryBlock, error) {
	block, err := db.FetchBlock([]byte{byte(DIRECTORYBLOCK_KEYMR)}, keyMR, new(directoryBlock.DirectoryBlock))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IDirectoryBlock), nil
}

// FetchDBlockByHeight gets an directory block by height from the database.
func (db *Overlay) FetchDBlockByHeight(dBlockHeight uint32) (interfaces.IDirectoryBlock, error) {
	block, err := db.FetchBlockByHeight([]byte{byte(DIRECTORYBLOCK_NUMBER)}, []byte{byte(DIRECTORYBLOCK_KEYMR)}, dBlockHeight, new(directoryBlock.DirectoryBlock))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IDirectoryBlock), nil
}

// FetchDBHashByHeight gets a dBlockHash from the database.
func (db *Overlay) FetchDBHashByHeight(dBlockHeight uint32) (interfaces.IHash, error) {
	return db.FetchBlockIndexByHeight([]byte{byte(DIRECTORYBLOCK_NUMBER)}, dBlockHeight)
}

// FetchDBHashByMR gets a DBHash by MR from the database.
func (db *Overlay) FetchDBHashByMR(dBMR interfaces.IHash) (interfaces.IHash, error) {
	return db.FetchPrimaryIndexBySecondaryIndex([]byte{byte(DIRECTORYBLOCK_KEYMR)}, dBMR)
}

// FetchDBlockByMR gets a directory block by merkle root from the database.
func (db *Overlay) FetchDBlockByHash(dBMR interfaces.IHash) (interfaces.IDirectoryBlock, error) {
	block, err := db.FetchBlockBySecondaryIndex([]byte{byte(DIRECTORYBLOCK_KEYMR)}, []byte{byte(DIRECTORYBLOCK)}, dBMR, new(directoryBlock.DirectoryBlock))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IDirectoryBlock), nil
}

// FetchAllDBlocks gets all of the fbInfo
func (db *Overlay) FetchAllDBlocks() ([]interfaces.IDirectoryBlock, error) {
	list, err := db.FetchAllBlocksFromBucket([]byte{byte(DIRECTORYBLOCK)}, new(directoryBlock.DirectoryBlock))
	if err != nil {
		return nil, err
	}
	return toDBlocksList(list), nil
}

func toDBlocksList(source []interfaces.BinaryMarshallableAndCopyable) []interfaces.IDirectoryBlock {
	answer := make([]interfaces.IDirectoryBlock, len(source))
	for i, v := range source {
		answer[i] = v.(interfaces.IDirectoryBlock)
	}
	return answer
}

func (db *Overlay) SaveDirectoryBlockHead(dblock interfaces.DatabaseBatchable) error {
	return db.ProcessDBlockBatch(dblock)
}

func (db *Overlay) FetchDirectoryBlockHead() (interfaces.IDirectoryBlock, error) {
	blk := new(directoryBlock.DirectoryBlock)
	block, err := db.FetchChainHeadByChainID([]byte{byte(DIRECTORYBLOCK)}, primitives.NewHash(blk.GetChainID()), blk)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IDirectoryBlock), nil
}
