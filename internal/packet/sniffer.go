package packet

import (
	"net"

	"github.com/umitaltintas/SpoofDPI/internal/cache"
)

type Sniffer interface {
	StartCapturing()
	RegisterUntracked(addrs []net.IPAddr, port int)
	GetOptimalTTL(key string) uint8
	Cache() cache.Cache
}
