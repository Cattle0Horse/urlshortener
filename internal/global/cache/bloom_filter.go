package cache

import (
	"hash/fnv"
	"math"
	"sync"
)

type BloomFilter struct {
	bitset  []bool
	size    uint
	hashNum uint
	mutex   sync.RWMutex
}

func NewBloomFilter(size uint, hashNum uint) *BloomFilter {
	return &BloomFilter{
		bitset:  make([]bool, size),
		size:    size,
		hashNum: hashNum,
	}
}

func (bf *BloomFilter) Add(item string) {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()

	for i := uint(0); i < bf.hashNum; i++ {
		hash := bf.hash(item, i)
		index := hash % bf.size
		bf.bitset[index] = true
	}
}

func (bf *BloomFilter) Contains(item string) bool {
	bf.mutex.RLock()
	defer bf.mutex.RUnlock()

	for i := uint(0); i < bf.hashNum; i++ {
		hash := bf.hash(item, i)
		index := hash % bf.size
		if !bf.bitset[index] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) hash(item string, seed uint) uint {
	h := fnv.New32a()
	h.Write([]byte(item))
	h.Write([]byte{byte(seed)})
	return uint(h.Sum32())
}

// OptimalHashNum 计算最优哈希函数数量
func OptimalHashNum(size uint, itemCount uint) uint {
	return uint(math.Ceil(float64(size) / float64(itemCount) * math.Ln2))
}

// OptimalSize 计算最优位数组大小
func OptimalSize(itemCount uint, falsePositiveRate float64) uint {
	return uint(math.Ceil(-float64(itemCount) * math.Log(falsePositiveRate) / math.Pow(math.Ln2, 2)))
}
