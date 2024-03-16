package cid_local

import (
    "bytes"
    chunker "github.com/ipfs/boxo/chunker"
    "github.com/ipfs/go-cid"
    mh "github.com/multiformats/go-multihash"
    "io"
)

func Cid(data []byte) string {
    chunks := chunker.NewSizeSplitter(bytes.NewReader(data), 262144)
    var buf bytes.Buffer
    for {
        chunk, err := chunks.NextBytes()
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
        buf.Write(chunk)
    }
    hash, err := mh.Sum(buf.Bytes(), mh.SHA2_256, -1)
    if err != nil {panic(err)}
    c := cid.NewCidV1(cid.Raw, hash)
    return c.String()
}
