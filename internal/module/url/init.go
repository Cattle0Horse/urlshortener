// Provide init function and global variables for url module,
// it will be called by `cmd/server/server.go`
package url

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/constant"
	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/logger"
	"github.com/Cattle0Horse/url-shortener/internal/global/redis"
	"github.com/Cattle0Horse/url-shortener/pkg/bloomfilter"
	pkgcache "github.com/Cattle0Horse/url-shortener/pkg/cache"
	"github.com/Cattle0Horse/url-shortener/pkg/tddl"
	"github.com/Cattle0Horse/url-shortener/pkg/tools"
	"github.com/Cattle0Horse/url-shortener/test"
)

var (
	log     *slog.Logger
	baseUrl string
	tddlGen tddl.TDDL
	bloom   bloomfilter.Interface
	cache   pkgcache.Interface

	cacheTTL time.Duration
)

type ModuleUrl struct{}

func (p *ModuleUrl) GetName() string {
	return "Url"
}

func (p *ModuleUrl) Init() {
	switch test.IsTest() {
	case false:
		log = logger.NewModule("Url")
	case true:
		log = logger.Get()
	}
	sc := config.Get().Server
	// http协议
	if sc.Port == "8080" {
		baseUrl = "http://" + sc.Host + sc.Prefix
	} else {
		baseUrl = "http://" + sc.Host + ":" + sc.Port + sc.Prefix
	}

	var err error
	tddlGen, err = tddl.New(database.DB)
	tools.PanicOnErr(err)

	// 创建布隆过滤器接口
	uc := config.Get().Url
	bloom = bloomfilter.NewRedisBloomFilter(
		redis.Client,
		constant.ShortCodeBloomFilterCacheKey,
		uc.BloomFilterSize,
		uc.BloomFilterFalsePositiveRate,
	)

	err = bloom.Create(context.Background())
	if err != nil {
		if errors.Is(err, bloomfilter.ErrBloomFilterAlreadyExists) {
			log.Info("bloom filter already exists")
		} else {
			panic(err)
		}
	}

	cache, err = pkgcache.NewProxy(redis.Client)
	tools.PanicOnErr(err)

	cacheTTL = config.Get().Cache.Redis.TTL

}

func selfInit() {
	u := &ModuleUrl{}
	u.Init()
}
