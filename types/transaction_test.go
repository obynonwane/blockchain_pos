package types

import (
	"testing"

	"github.com/obynonwane/blockchain_project/crypto"
	"github.com/obynonwane/blockchain_project/proto"
	"github.com/obynonwane/blockchain_project/util"
)

func TestNewTransaction(t *testing.T) {
	// private key of the sender
	fromPrivKey := crypto.GeneratePrivateKey()
	// address of the sender
	fromAddress := fromPrivKey.Public().Bytes()

	// private key of the receiver
	// in other to get the address of teh receiver - testing purposes
	toPrivKey := crypto.GeneratePrivateKey()
	//address generated from private key
	toAddress := toPrivKey.Public().Bytes()

	// input to the txn
	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}
	// output to the receiver of txns
	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}

	//out put to the sender who receives 95 back - ie using UTXO
	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

}
