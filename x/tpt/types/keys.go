package types

const (
	// ModuleName defines the module name
	ModuleName = "tpt"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tpt"
)

var (
	ParamsKey = []byte("p_tpt")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
