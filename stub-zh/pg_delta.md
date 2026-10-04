## 用法

来源：

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [extensions/pg_delta/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_delta/pgbrew.toml)
- [extensions/pg_delta/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_delta/README.md)
- [extensions/pg_delta/pg_delta.control](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_delta/pg_delta.control)
- [extensions/pg_delta/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_delta/src/lib.rs)
- [extensions/pg_delta/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_delta/pgbrew.toml)
- [docs/pg_delta.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/docs/pg_delta.md)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/LICENSE)

`pg_delta` 0.3.2 通过读取、导出和受管理的流将 PostgreSQL 与 Delta Lake 集成。新增索引模式登记 Delta 事务日志并裁剪文件，以 FDW 原地查询。

### 核心用法

```ini
shared_preload_libraries = 'pg_delta'
```

```sql
CREATE EXTENSION pg_delta;
SELECT * FROM delta.list_tables();
SELECT delta.status();
```

### 运行边界

安装需要超级用户。流管理器要求预加载 `pg_delta` 并重启，应明确配置工作数据库和存储凭据。`delta` 模式提供表／流创建、刷新、状态、历史与导出接口，索引模式使用 `pg_delta_server`。云路径和 SQL 定义属于特权操作。手册将基于逻辑复制的 CDC 导出标为未实现；轮询和快照模式各有更新／删除及恢复语义。这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

### 当前版本与升级

扩展版本 0.3.2 通过 `delta.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。 已撤回的仓库 0.4.0 二进制应替换为 0.4.1；上游二进制不代表 Pigsty 软件包可用性。
