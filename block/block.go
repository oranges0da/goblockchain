package block

import (
	"bytes"
	"crypto/sha256"
	"log"
)

type Block struct {
	BlockID   int
	Nonce     int
	IsGenesis bool
	PrevHash  []byte
	Data      []byte
	Hash      []byte
}

func (b *Block) GetHash() []byte {
	concat_data := [][]byte{b.PrevHash, b.Data}

	data := bytes.Join(concat_data, []byte{})

	hash := sha256.Sum256(data)
	log.Printf("Block %d: %x\n", b.BlockID, hash)

	return hash[:]
}

func NewBlock(BlockId int, PrevHash []byte, data string) *Block {
	block := &Block{
		BlockID:   BlockId,
		IsGenesis: false,
		PrevHash:  PrevHash,
		Data:      []byte(data),
	}

	hash := block.GetHash()

	block.Hash = hash

	return block
}

func Genesis() *Block {
	block := &Block{
		BlockID:   0,
		IsGenesis: true,
		PrevHash:  []byte{},
		Data:      []byte("Genesis Block"),
	}

	hash := block.GetHash()

	block.Hash = hash

	return block
}
