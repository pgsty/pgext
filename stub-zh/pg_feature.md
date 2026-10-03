## 用法

来源：

- [extensions/pg_feature/pg_feature.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_feature/pg_feature.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_feature/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_feature/Cargo.toml)
- [extensions/pg_feature/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_feature/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_feature/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_feature/README.md)

`pg_feature` 0.3.0 在 `pgft` 中提供关系型特征合成：登记表关系与时间索引，再通过聚合和时间截点生成特征。

### 核心用法

```sql
CREATE EXTENSION pg_feature;
CREATE TABLE feature_events (id bigint PRIMARY KEY, created_at timestamptz, value numeric);
SELECT pgft.set_time_index('feature_events', 'created_at');
SELECT * FROM pgft.synthesize_features(
  'event_features', 'feature_events', 'id',
  cutoff_time => TIMESTAMPTZ '2026-10-01 00:00:00+00', max_depth => 1
);
SELECT * FROM pgft.list_features('event_features');
```

### 运行边界

控制文件不限定超级用户安装，但仍需具备创建相应对象的权限。无需预加载或 Python 运行时。`set_time_index` 与 `add_relationship` 建立元数据，`synthesize_features` 创建输出，`list_features` 查看定义。登记的时间戳和截点应避免未来数据泄漏。物化输出需要显式刷新；用于大表或敏感数据前，应检查生成的 SQL 与表关系。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
