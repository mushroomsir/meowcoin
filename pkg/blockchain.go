package pkg

import (
	"encoding/hex"

	"github.com/boltdb/bolt"
	"github.com/mushroomsir/meowcoin/tools"
)

const dbFile = "data/blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "2018年1月26日起始"

func NewBlockchain(address string) *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	tools.Check(err)
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
			genesis := NewGenesisBlock(cbtx)
			b, err := tx.CreateBucket([]byte(blocksBucket))
			tools.Check(err)
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func (bc *Blockchain) FindUnspentTransactions(address string) {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()
	for {
		block := bci.Next()
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		}
	}
}
func (bc *Blockchain) CloseDB() {
	tools.Print(bc.db.Close())
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}
	return bci
}
