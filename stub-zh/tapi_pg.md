## 用法

来源：

- [docs/operate/postgres.md](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/docs/operate/postgres.md)
- [integrations/tapi-postgres/tapi_pg--0.1.sql](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/tapi_pg--0.1.sql)
- [integrations/tapi-postgres/tapi_pg.c](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/tapi_pg.c)
- [integrations/tapi-postgres/ddl.c](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/ddl.c)
- [integrations/tapi-postgres/txn.c](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/txn.c)
- [LICENSE](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/LICENSE)
- [integrations/tapi-postgres/tapi_pg.control](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/tapi_pg.control)

`tapi_pg` 0.1 将 C 外部数据封装器连接到 Rust Tapirs C API。使用 SQL 前需要运行中的集群、共享预加载和服务器重启。需按部署的发现端点与目录设置配置 `tapirs.cluster_config`。

### 核心用法

```ini
shared_preload_libraries = 'tapi_pg'
```

```sql
CREATE EXTENSION tapi_pg;
SELECT tapi_pg_status();
CREATE SERVER tapirs_store FOREIGN DATA WRAPPER tapi_pg;
CREATE FOREIGN TABLE tapirs_users (id integer, name text)
    SERVER tapirs_store
    OPTIONS (keyspace 'app', "table" 'users', primary_key 'id');
```

### 运行边界

`tapi_pg_status()` 检查集群连通性。外部表必须设置 `keyspace`、`table`、`primary_key`；可选的 `secondary_indexes` 在建表时声明，不通过 CREATE INDEX 创建。事务使用 Tapirs 乐观并发控制，序列化失败需要应用重试。已核验实现不支持保存点、两阶段提交和自动同步远程表重命名。此集成尚未接入集群 TLS，需要注意网络边界。安装要求超级用户权限及匹配 PostgreSQL 大版本的共享库；上游没有声明通用的大版本范围。集成 README 对表重命名的描述过强，操作指南与 DDL 源码明确说明该功能仍待实现。
