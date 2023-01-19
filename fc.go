package main

import (
	"fmt"
	"crypto/ed25519"
	"context"

	//"lukechampine.com/us/ed25519"
	"lukechampine.com/us/hostdb"
  "lukechampine.com/us/renter"
  "lukechampine.com/us/renter/proto"
  "lukechampine.com/us/renter/renterutil"
  //"lukechampine.com/us/renterhost"

  "lukechampine.com/frand"
  //"gitlab.com/NebulousLabs/Sia/crypto"
)

func main() {

const siadPassword = "siaP4ss" // found in ~/.sia/apipassword
siad := renterutil.NewSiadClient("localhost:9980", siadPassword)

const hostPubKey = "ed25519:25407e439611ef4310034a07b6f618292117857bd7d9b2025db6e6da1102efde"

hostIP, _ := siad.ResolveHostKey(hostPubKey)
host, _ := hostdb.Scan(context.Background(), hostIP, hostPubKey)

fmt.Printf("%+v\n", host)

const uploadedBytes = 1e9
const downloadedBytes = 2e9
const duration = 1000
uploadFunds := host.UploadBandwidthPrice.Mul64(uploadedBytes)
downloadFunds := host.DownloadBandwidthPrice.Mul64(downloadedBytes)
storageFunds := host.StoragePrice.Mul64(uploadedBytes).Mul64(duration)
totalFunds := uploadFunds.Add(downloadFunds).Add(storageFunds)

currentHeight, _ := siad.ChainHeight()
start, end := currentHeight, currentHeight+duration

key := ed25519.NewKeyFromSeed(frand.Bytes(32))

fmt.Println("Generated renter key.")

rev, contractTxn, err := proto.FormContract(siad, siad, key, host, totalFunds, start, end)
if err != nil {
		fmt.Println(err)
}

contract := renter.Contract{
    HostKey:   rev.HostKey(),
    ID:        rev.ID(),
    RenterKey: key,
}

fmt.Println("Formed contract with ID ", contract.ID )
fmt.Printf("%+v\n", contractTxn)

//contractPath := "mycontract.contract"
//_ = renter.SaveContract(contract, contractPath)

}
