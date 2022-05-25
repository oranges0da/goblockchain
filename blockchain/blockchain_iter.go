package blockchain

import (
	badger "github.com/dgraph-io/badger/v3"
	"github.com/oranges0da/goblockchain/block"
	"github.com/oranges0da/goblockchain/utils"
)

type BlockchainIter struct {
	currentHash []byte
	db          *badger.DB
}

func (chain *Blockchain) NewIter() *BlockchainIter {
	iter := &BlockchainIter{
		currentHash: chain.LastHash,
		db:          chain.Database,
	}

	return iter
}

func (iter *BlockchainIter) Next() *block.Block {
	var block *block.Block

	err := iter.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.currentHash)
		utils.Handle(err)

		encoded, err := item.Value()

		return err
	})

	iter.currentHash = block.PrevHash
}
