## 用法

来源：

- [官方 README.md](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/README.md)
- [官方 pg_keyspace.control](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/extension/pg_keyspace.control)
- [官方 Cargo.toml](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/extension/Cargo.toml)
- [官方 lib.rs](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/extension/src/lib.rs)
- [官方 pg_keyspace--0.5.0--0.6.0.sql](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/extension/sql/pg_keyspace--0.5.0--0.6.0.sql)

`pg_keyspace` 0.6.0 在 PostgreSQL 后台进程中运行 RESP2/RESP3 键空间，并提供可选的 PostgREST 行缓存。上游在 PostgreSQL 17 上构建和测试该实现。control 与 SQL 版本为 0.6.0，Rust 包清单仍标为 0.1.0。

### 独立键空间

预加载并重启，再以超级用户安装。独立模式须显式关闭可选的列掩码集成：

```conf
shared_preload_libraries = 'pg_keyspace'
pg_keyspace.port = 6381
pg_keyspace.require_mask = off
```

```sql
CREATE EXTENSION pg_keyspace;
```

```sh
redis-cli -p 6381 SET demo:item example
redis-cli -p 6381 GET demo:item
```

```sql
SELECT convert_from(supacache.get('demo:item'), 'UTF8');
SELECT * FROM supacache.stats();
```

最小示例假定监听端口仅用于隔离测试。扩大访问范围前应配置 RESP 凭据、ACL、租户隔离与 TLS。`supacache.set_credential` 管理 RESP 凭据，不能将其等同于 PostgreSQL 连接认证。`supatype_mask` 和 `pg_guard` 均不是必需依赖。

### 持久性与行缓存

`pg_keyspace.durability` 可选择 ephemeral、relaxed、durable 或 replicated 模式；不同模式的确认与崩溃恢复语义不同，持久化数据存入 `supacache.kv`。保存不可重建的数据前，应选择并验证所需模式。

可选行缓存使用规划器钩子与配套的仅键解码插件。读穿缓存须显式启用，并依赖失效解码进程。其契约允许有界的缓存陈旧，并不无条件等同于事务快照；在缓存行之上重新检查 RLS，也不会使陈旧数据自动变新。

### 运行边界

尚未实现 Streams、Lua 脚本和阻塞式列表操作。多个后台进程具有各自的键空间，须明确配置客户端路由与恢复方式。流复制备库不提供 RESP 服务。跨实例发布订阅转发须显式启用，采用至多一次投递，不是键空间复制。替换既有 Redis 服务前，应演练升级并检查该版本的限制。
