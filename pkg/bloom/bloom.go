package bloom

import (
	"hash"
	"hash/fnv"
	"math"
	"sync"
)

// Filter 布隆过滤器实现
type Filter struct {
	lock   sync.RWMutex
	m      uint64   // 位数组大小
	k      uint     // 哈希函数个数
	b      []uint64 // 位数组
	hasher hash.Hash64
}

// New 创建布隆过滤器
func New(n uint64, fp float64) *Filter {
	// 计算最优位数组大小
	m := uint64(-float64(n) * math.Log(fp) / math.Pow(math.Log(2), 2))
	// 计算最优哈希函数数量
	k := uint(math.Ceil(math.Log(2) * float64(m) / float64(n)))

	return &Filter{
		m:      m,
		k:      k,
		b:      make([]uint64, (m+63)/64),
		hasher: fnv.New64(),
	}
}

// Add 添加元素
func (f *Filter) Add(data []byte) {
	f.lock.Lock()
	defer f.lock.Unlock()

	h := f.hash(data)
	for i := uint(0); i < f.k; i++ {
		pos := (h + uint64(i)*h) % f.m
		f.b[pos/64] |= 1 << (pos % 64)
	}
}

// MayExist 判断元素是否存在
func (f *Filter) MayExist(data []byte) bool {
	f.lock.RLock()
	defer f.lock.RUnlock()

	h := f.hash(data)
	for i := uint(0); i < f.k; i++ {
		pos := (h + uint64(i)*h) % f.m
		if f.b[pos/64]&(1<<(pos%64)) == 0 {
			return false
		}
	}
	return true
}

func (f *Filter) hash(data []byte) uint64 {
	f.hasher.Reset()
	f.hasher.Write(data)
	return f.hasher.Sum64()
}
