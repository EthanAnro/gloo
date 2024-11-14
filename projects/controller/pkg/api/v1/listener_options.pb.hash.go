// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/listener_options.proto

package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *ListenerOptions) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controllerrollerroller/pkg/api/v1.ListenerOptions")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAccessLoggingService()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AccessLoggingService")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAccessLoggingService(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AccessLoggingService")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetExtensions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Extensions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExtensions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Extensions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PerConnectionBufferLimitBytes")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPerConnectionBufferLimitBytes(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PerConnectionBufferLimitBytes")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetSocketOptions() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetProxyProtocol()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ProxyProtocol")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetProxyProtocol(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ProxyProtocol")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetConnectionBalanceConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ConnectionBalanceConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetConnectionBalanceConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ConnectionBalanceConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetListenerAccessLoggingService()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ListenerAccessLoggingService")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetListenerAccessLoggingService(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ListenerAccessLoggingService")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTcpStats()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TcpStats")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTcpStats(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TcpStats")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *ConnectionBalanceConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controllerrollerroller/pkg/api/v1.ConnectionBalanceConfig")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetExactBalance()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ExactBalance")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExactBalance(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ExactBalance")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *ConnectionBalanceConfig_ExactBalance) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controllerrollerroller/pkg/api/v1.ConnectionBalanceConfig_ExactBalance")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}