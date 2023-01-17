package main

import (
		"crypto/ed25519"
		"context"

    //"lukechampine.com/us/ed25519hash"
    "lukechampine.com/us/renter"
    "lukechampine.com/us/renter/proto"
    "lukechampine.com/us/renter/renterutil"
    //"lukechampine.com/us/renterhost"

    "lukechampine.com/frand"
    //"gitlab.com/NebulousLabs/Sia/crypto"
)

func main() {

const siadPassword = "c4c5f8884a5d6c4f731c24c75378b2b3" // found in ~/.sia/apipassword
siad := renterutil.NewSiadClient(":9980", siadPassword)

const hostPubKey = "ed25519:25407e439611ef4310034a07b6f618292117857bd7d9b2025db6e6da1102efde"

hostIP, _ := siad.ResolveHostKey(hostPubKey)
host, _ := hostdb.Scan(context.Background(), hostIP, hostPubKey)

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
rev, _ := proto.FormContract(siad, siad, key, host, totalFunds, start, end)
contract := renter.Contract{
    HostKey:   rev.HostKey(),
    ID:        rev.ID(),
    RenterKey: key,
}

//contractPath := "mycontract.contract"
//_ = renter.SaveContract(contract, contractPath)

}
