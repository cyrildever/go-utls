module github.com/cyrildever/go-utls

go 1.26.10

require (
	github.com/beevik/ntp v1.6.0
	github.com/btcsuite/btcd v0.25.0
	github.com/btcsuite/btcd/btcec/v2 v2.5.1
	github.com/btcsuite/btcd/btcutil v1.2.0
	github.com/elliotchance/orderedmap v1.8.0
	github.com/ethereum/go-ethereum v1.17.10
	github.com/fatih/color v1.19.0
	github.com/go-stack/stack v1.8.1
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/sammyne/base58 v0.1.0
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/valyala/fasthttp v1.79.0
	github.com/vmihailenco/msgpack/v5 v5.4.1
	gitlab.com/metakeule/fmtdate v1.2.2
	go.mongodb.org/mongo-driver v1.17.10
	golang.org/x/crypto v0.59.0
	golang.org/x/net v0.63.0
	golang.org/x/text v0.44.0
	gotest.tools v2.2.0+incompatible
)

require (
	github.com/ProjectZKM/Ziren/crates/go-runtime/zkvm_runtime v0.0.0-20261119055430-c9988a098d92 // indirect
	github.com/bits-and-blooms/bitset v1.30.0 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.2.0 // indirect
	github.com/consensys/gnark-crypto v0.22.0 // indirect
	github.com/crate-crypto/go-eth-kzg v1.5.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.1 // indirect
	github.com/ethereum/c-kzg-4844/v2 v2.1.8 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/kcalvinalvin/anet v0.0.0-20251112173137-d8ddc1f6dbee // indirect
	github.com/klauspost/compress v1.20.5 // indirect
	github.com/mattn/go-colorable v0.1.16 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/molecule-man/go-brrr v1.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/supranational/blst v0.3.17 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/sync v0.25.0 // indirect
	golang.org/x/sys v0.50.0 // indirect
)

// btcd >= v0.26.0 drops the chaincfg and wire packages still imported by btcd/btcutil v1.2.0
exclude (
	github.com/btcsuite/btcd v0.26.0-beta.rc1
	github.com/btcsuite/btcd v0.26.0
	github.com/btcsuite/btcd v0.26.1-beta.rc1
	github.com/btcsuite/btcd v0.26.2
	github.com/btcsuite/btcd v0.27.0
)
