## 用法

来源：

- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_xarray/pg_xarray.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_xarray/pg_xarray.control)
- [extensions/pg_xarray/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_xarray/src/lib.rs)
- [extensions/pg_xarray/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_xarray/pgbrew.toml)
- [extensions/pg_xarray/demo/01_register_local.sql](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_xarray/demo/01_register_local.sql)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)

`pg_xarray` 0.3.0 在 `pgx` 中登记科学数据集、变量、数据块和网格，提供科学数组查询与可选的 WMS 服务。

### 核心用法

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_xarray;
SELECT pgx.register_dataset('weather', 'zarr');
SELECT * FROM pgx.list_datasets();
```

### 运行边界

依赖 `postgis`，需要超级用户安装。数据集／文件登记指向服务器可读取的本地或对象存储资源，应限制路径和凭据。SQL 接口包括 `register_dataset`、`register_file`、`register_variable`、`register_chunk` 和 `list_datasets`。WMS 需要预加载、重启并显式启用 `pg_xarray.wms_enabled`，默认关闭。启用 SQL 目录并不会开启公开瓦片服务。备份时需同时保留外部数组数据与目录。这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
