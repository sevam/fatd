// Package testdata creates a runtime.Context for use in the runtime test. This
// intermediary package allows the test data to depend directly on a header
// file shared with the API test C code.
package testdata

// #include "./src/runtime_test.h"
import "C"
import (
	"time"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/Factom-Asset-Tokens/factom/fat0"
	"github.com/Factom-Asset-Tokens/fatd/internal/db"
	"github.com/Factom-Asset-Tokens/fatd/internal/db/addresses"
	"github.com/Factom-Asset-Tokens/fatd/internal/runtime"
)

// Context returns a runtime.Context populated with the test data expected by
// the api_test.wasm.
func Context(chain db.Chain) runtime.Context {
	var tx fat0.Transaction

	sender := factom.FAAddress(genBytes32(C.GET_SENDER_ERR))
	address := factom.FAAddress(genBytes32(C.GET_ADDRESS_ERR))
	tx.Inputs = fat0.AddressAmountMap{sender: C.GET_AMOUNT_EXP}
	tx.Outputs = fat0.AddressAmountMap{address: C.GET_AMOUNT_EXP}

	hash := genBytes32(C.GET_ENTRY_HASH_ERR)
	tx.Entry.Hash = &hash
	tx.Entry.Timestamp = time.Unix(C.GET_TIMESTAMP_EXP, 0)

	chain.Issuance.Precision = C.GET_PRECISION_EXP
	if _, err := addresses.Add(chain.Conn, &address, C.GET_BALANCE_EXP); err != nil {
		panic(err)
	}
	address = factom.FAAddress(genBytes32(C.GET_BALANCE_OF_ERR))
	addresses.Add(chain.Conn, &address, C.GET_BALANCE_OF_EXP)

	return runtime.Context{
		DBlock:      factom.DBlock{Height: uint32(C.GET_HEIGHT_EXP)},
		Chain:       chain,
		Transaction: tx,
	}
}

// ErrMap is a map between run_all() return codes and their error string.
var ErrMap = map[int32]string{
	int32(C.SUCCESS):             "success",
	int32(C.GET_HEIGHT_ERR):      "error: ext_get_height",
	int32(C.GET_SENDER_ERR):      "error: ext_get_sender",
	int32(C.GET_AMOUNT_ERR):      "error: ext_get_amount",
	int32(C.GET_ENTRY_HASH_ERR):  "error: ext_get_entry_hash",
	int32(C.GET_TIMESTAMP_ERR):   "error: ext_get_timestamp",
	int32(C.GET_PRECISION_ERR):   "error: ext_get_precision",
	int32(C.GET_BALANCE_ERR):     "error: ext_get_balance",
	int32(C.GET_BALANCE_OF_ERR):  "error: ext_get_balance_of",
	int32(C.SEND_ERR_BALANCE):    "error: ext_send: balance",
	int32(C.SEND_ERR_BALANCE_OF): "error: ext_send: balance_of",
	int32(C.BURN_ERR_BALANCE):    "error: ext_burn: balance",
	int32(C.BURN_ERR_BALANCE_OF): "error: ext_burn: balance_of",
}

func genBytes32(val byte) factom.Bytes32 {
	var hash factom.Bytes32
	for i := range hash[:] {
		hash[i] = byte(i) + val
	}
	return hash
}