## Todo

- [x] 过期处理（redis 中可以存储过期时间，或者 redis 中设置到期时间少于数据库过期时间） - 缓存数据库一致性
- [x] 添加布隆过滤器减少缓存击穿情况的发生
- [x] 容器化部署
- [x] 服务限流
- [x] 前端部署加入 docker compose
- [x] Nginx 部署
- [ ] 使用「布谷鸟过滤器」或「Counting Bloom Filter」替换 普通布隆过滤器（他们支持删除操作）
- [ ] 修改短链接到期时间的缓存数据库一致性问题（如延迟双删、给缓存加锁）
- [ ] 缓存请求超时处理
- [ ] 优化仪表盘分页查询方式，防止深度分页导致性能问题
- [ ] 优雅的退出
- [ ] 链接访问统计
- [ ] 优化连接关闭（返回 clean 接口，类似于 `contextWithCancel`，可参考 [Apache-Answer](https://github.com/apache/incubator-answer) ）
- [ ] 连接重试（如数据库），而非 panic（可参考 [Apache-Answer](https://github.com/apache/incubator-answer) ）
- [ ] 读写分离（参考 [beihai0xff/turl](https://github.com/beihai0xff/turl) ）
- [ ] 更好的集成测试与单元测试，可以对于有依赖的内容可以借助 [testcontainers](https://github.com/testcontainers/testcontainers-go)
- [ ] ~~批量生成短链接功能（如提交文件）~~
- [ ] ~~链接访问数据导出~~
- [ ] ~~相同 url 可以幂等~~
- [ ] ~~过期短链的处理（如轮询扫描全表，删除过期的）~~
- [ ] 可观测性监控

