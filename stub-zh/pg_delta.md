## 用法

来源：

- [extensions/pg_delta/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_delta/README.md)
- [extensions/pg_delta/pg_delta.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_delta/pg_delta.control)
- [extensions/pg_delta/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_delta/src/lib.rs)
- [extensions/pg_delta/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_delta/pgbrew.toml)
- [docs/pg_delta.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/docs/pg_delta.md)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)

`pg_delta` 0.3.0 通过读取、导出和受管理的流将 PostgreSQL 与 Delta Lake 集成。新增索引模式登记 Delta 事务日志并裁剪文件，以 FDW 原地查询。

### 核心用法

```conf
shared_preload_libraries = 'pg_delta'
```

```sql
CREATE EXTENSION pg_delta;
SELECT * FROM delta.list_tables();
SELECT delta.status();
```

### 运行边界

安装需要超级用户。流管理器要求预加载 `pg_delta` 并重启，应明确配置工作数据库和存储凭据。`delta` 模式提供表／流创建、刷新、状态、历史与导出接口，索引模式使用 `pg_delta_server`。云路径和 SQL 定义属于特权操作。手册将基于逻辑复制的 CDC 导出标为未实现；轮询和快照模式各有更新／删除及恢复语义。这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
