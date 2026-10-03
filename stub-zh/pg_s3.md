## 用法

来源：

- [extensions/pg_s3/pg_s3.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_s3/pg_s3.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_s3/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_s3/Cargo.toml)
- [extensions/pg_s3/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_s3/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_s3/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_s3/pgbrew.toml)

`pg_s3` 0.3.0 提供实验性的 S3 兼容 HTTP API，在 `pgs3` 中保存桶／对象元数据，在本地磁盘保存二进制内容。

### 核心用法

```conf
shared_preload_libraries = 'pg_s3'
pg_s3.database = 'postgres'
pg_s3.host = '127.0.0.1'
pg_s3.port = 9100
```

```sql
CREATE EXTENSION pg_s3;
SELECT pgs3.create_bucket('sample-bucket');
SELECT * FROM pgs3.list_buckets();
```

### 运行边界

控制文件要求超级用户安装。需要预加载并重启。通过 `pg_s3.database`、`pg_s3.host`、`pg_s3.port` 和 `pg_s3.data_directory` 配置服务，默认包括所有接口、9100 端口与 postgres 数据库。SQL API 管理桶、对象元数据及列表。备份必须一致地覆盖数据库元数据和对象文件。应限制文件路径与网络访问；源码可用不代表完整的 S3 API、认证或多节点持久性兼容。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
