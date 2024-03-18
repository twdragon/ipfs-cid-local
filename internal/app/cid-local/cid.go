package cid_local

import (
    "bytes"
    "os"
    chunker "github.com/ipfs/boxo/chunker"
    "github.com/ipfs/go-cid"
//    mh "github.com/multiformats/go-multihash"
//    "io"
    dsync "github.com/ipfs/go-datastore/sync"
    multicodec "github.com/multiformats/go-multicodec"
//    unixfile "github.com/ipfs/boxo/ipld/unixfs/file"
	"github.com/ipfs/boxo/ipld/unixfs/importer/balanced"
	uih "github.com/ipfs/boxo/ipld/unixfs/importer/helpers"
	"github.com/ipfs/go-datastore"
	blockstore "github.com/ipfs/boxo/blockstore"
	"github.com/ipfs/boxo/blockservice"
	offline "github.com/ipfs/boxo/exchange/offline"
	"github.com/ipfs/boxo/ipld/merkledag"
)

func Cid(filename string) string {
    fileData, err := os.ReadFile(filename)
    if err != nil {panic(err)}
    fileReader := bytes.NewReader(fileData)
    ds := dsync.MutexWrap(datastore.NewNullDatastore())
    bs := blockstore.NewBlockstore(ds)
	bs = blockstore.NewIdStore(bs)
    bsrv := blockservice.New(bs, offline.Exchange(bs))
	dsrv := merkledag.NewDAGService(bsrv)
	// Create a UnixFS graph from our file, parameters described here but can be visualized at https://dag.ipfs.tech/
	ufsImportParams := uih.DagBuilderParams{
		Maxlinks:  uih.DefaultLinksPerBlock, // Default max of 174 links per block
		RawLeaves: true,                     // Leave the actual file bytes untouched instead of wrapping them in a dag-pb protobuf wrapper
		CidBuilder: cid.V1Builder{ // Use CIDv1 for all links
			Codec:    uint64(multicodec.Raw),
			MhType:   uint64(multicodec.Sha2_256), // Use SHA2-256 as the hash function
			MhLength: -1,                          // Use the default hash length for the given hash function (in this case 256 bits)
		},
		Dagserv: dsrv,
		NoCopy:  false,
	}
	ufsBuilder, err := ufsImportParams.New(chunker.NewSizeSplitter(fileReader, chunker.DefaultBlockSize)) // Split the file up into fixed sized 256KiB chunks
	if err != nil {
		return cid.Undef.String()
	}
	nd, err := balanced.Layout(ufsBuilder) // Arrange the graph with a balanced layout
	if err != nil {
		return cid.Undef.String()
	}
	return nd.Cid().String()
}
