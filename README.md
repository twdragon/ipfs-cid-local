# IPFS CID Calculator

This repository contains an example program that calculates [IPFS CID](https://docs.ipfs.tech) for the single file specified in the command line. The program is inspired by [this StackOverflow answer](https://stackoverflow.com/a/76799856/9560245) by [**lajosdeme**](https://github.com/lajosdeme) 

## Default CID parameters
The program uses the following values of hashing/DAG construction parameters. They are considered as defaults for the current [IPFS Kubo](https://github.com/ipfs/kubo/) implementation. 

|Parameter|Value|
|---------|-----|
|Chunker|`size-262144`|
|Hash|`sha2-256`|
|Codec|`raw`|
|Version|1|
|Encoding|`base32`|

## Compiling
```bash
go build
```
## Usage
```bash
cid-local <filename>
```

