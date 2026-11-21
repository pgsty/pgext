## 用法

来源：

- [README.md](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/README.md)
- [bridge.md](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/docs/bridge.md)
- [limitations.md](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/docs/limitations.md)
- [tessera.control](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/bridge/tessera.control)
- [tessera--0.1.sql](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/bridge/tessera--0.1.sql)
- [LICENSE](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/LICENSE)

`tessera` 是面向普通 PostgreSQL 堆表的批量执行器桥接扩展。项目通过 C 模块和 Rust 内核处理支持的分析计划节点。SQL 扩展版本 0.1 加载桥接库，暴露 C rendezvous API，而非 SQL 函数。上游面向 PostgreSQL 18、19 及开发中的 master。

### 启用模块

安装匹配的桥接、节点和内核库。常规使用按以下顺序预加载，并重启 PostgreSQL：

```ini
shared_preload_libraries = 'tessera, tessera_nodes, tessera_kernels'
```

保留其他已有库。以超级用户创建扩展，并在代表性数据上比较计划：

```sql
CREATE EXTENSION tessera;
CREATE TABLE tessera_demo (x integer);
INSERT INTO tessera_demo SELECT generate_series(1, 100000);
ANALYZE tessera_demo;
EXPLAIN ANALYZE SELECT count(*), sum(x) FROM tessera_demo WHERE x > 1000;
SET tessera.enable = off;
EXPLAIN ANALYZE SELECT count(*), sum(x) FROM tessera_demo WHERE x > 1000;
```

桥接库必须先于节点及内核模块加载。开发时也支持按相同顺序会话预加载。CREATE EXTENSION 只在创建扩展的后端加载桥接库；目录中存在扩展，不会让后续连接自动加载它。新连接必须预加载或显式加载这些库。

### 计划与限制

支持的节点包括堆表与索引扫描、批量过滤及聚合、哈希连接、排序与集合运算。不支持的节点在混合计划中仍由 PostgreSQL 执行。窗口函数、外部表、其他表访问方法以及写入或锁行语句的扫描仍使用普通执行器。应通过 EXPLAIN 检查实际选择，性能收益并无保证。

可见性遵循 PostgreSQL 快照。哈希与排序使用 work_mem 和 hash_mem_multiplier，并可在 temp_file_limit 约束下溢写。批量循环检查取消信号。切换 PostgreSQL 安装时须重建模块，升级时须替换主进程或会话已加载的库。项目仍在积极开发，使用前应验证结果等价性与实际负载行为。
