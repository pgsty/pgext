## 用法

来源：

- [README.md](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/README.md)
- [pg_local_cache.control](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/pg_local_cache.control)
- [sql/pg_local_cache--3.0.0.sql](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/sql/pg_local_cache--3.0.0.sql)
- [sql/pg_local_cache--2.0.4--3.0.0.sql](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/sql/pg_local_cache--2.0.4--3.0.0.sql)
- [docs/UPGRADING.md](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/docs/UPGRADING.md)
- [CHANGELOG.md](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/CHANGELOG.md)

`pg_local_cache` 3.0.0 在有容量边界的共享内存中按完整主键缓存整行。缓存读取现通过认证的 RESP MGET 完成；SQL `local_cache.mget` 已被移除，普通 SELECT 也不会被改写。PostgreSQL 仍是持久化数据的权威来源。

### 核心用法

```ini
shared_preload_libraries = 'pg_local_cache'
pg_local_cache.database = 'app'
pg_local_cache.role = 'local_cache_worker'
pg_local_cache.bind_address = '127.0.0.1'
pg_local_cache.port = 6380
pg_local_cache.auth_token_file = '/secure/path/token'
```

```sql
CREATE EXTENSION pg_local_cache;
CREATE TABLE public.items (id bigint PRIMARY KEY, value text);
INSERT INTO public.items VALUES (42, 'example');
SELECT local_cache.attach_table('public.items'::regclass);
SELECT local_cache.health();
```

RESP 连接认证后：

```text
MGET CRUD:app.public.items:{"id":42} CRUD:app.public.items:{"id":7}
```

### 运行边界

将库追加到已有预加载列表，建立专用工作角色并按文档授予元数据与表权限，保护令牌文件，然后重启。超级用户在固定模式 `local_cache` 中创建扩展，仅注册受支持的永久主键表。RESP MGET 保留键顺序与重复项，缺失行返回 null。工作进程使用配置好的数据库角色，不继承客户端 SQL 权限、事务或快照。连接、投影、行锁及需要会话语义的读取仍应使用普通 SQL。

`local_cache.attach_table` 安装失效触发器，`local_cache.detach_table` 删除映射，`local_cache.reconcile_table` 在 DDL 或权限变化后重新验证映射。`local_cache.health`、`local_cache.stats` 和 `local_cache.metrics` 报告就绪状态与资源使用。普通 PostgreSQL 写入会使相关缓存失效。RLS、分区表及继承表不在文档支持范围内，也不提供 TTL 或分布式缓存协调。

支持 PostgreSQL 14–18 上的单个可写主库。非回环监听要求启用 `pg_local_cache.tls`，或明确选择 `pg_local_cache.allow_plaintext_network`。TLS 使用独立证书与私钥配置，设置 CA 可启用双向 TLS；可重新加载的 `pg_local_cache.enabled` 开关可禁用缓存。

从 2.x 升级前，先将 SQL MGET 调用迁移到 RESP 授权模型。安装匹配的 3.0.0 文件，配置监听安全并重启，确认库版本与监听就绪后，在各受影响数据库执行 `ALTER EXTENSION pg_local_cache UPDATE`。依赖被移除函数或旧 metrics 返回类型的对象可能阻止迁移；脚本有意不使用 CASCADE。上游未提供降级脚本。
