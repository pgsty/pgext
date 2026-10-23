## 用法

来源：

- [官方资料](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/extensions/pg_feature/pg_feature.control)
- [官方资料](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/README.md)
- [官方资料](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/extensions/pg_feature/Cargo.toml)
- [官方资料](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/extensions/pg_feature/src/lib.rs)
- [官方资料](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/CHANGELOG.md)
- [官方资料](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/LICENSE)
- [官方资料](https://github.com/matroidbe/pg_extensions-releases/blob/f6caae20c3500dcb2b8419ed68f9b380333b552b/extensions/pg_feature/README.md)

`pg_feature` 0.4.0 在 `pgft` 中提供关系型特征合成：登记表关系与时间索引，再通过聚合和时间截点生成特征。

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

### 0.4.0 的一致性修复

仓库发布 0.7.3 包含扩展版本 0.4.0，修复训练矩阵与实时特征 SQL 的计数一致性、二层计算特征、时间截点与窗口处理。计数列应保留空值语义，不能用计数全部行替代。升级会保留已有截点间隔数据，不执行级联删除。仍应以有代表性的数据核对生成 SQL 与训练、在线结果。PostgreSQL 特性范围为 14–18；已有 0.3.0 及以后实例使用匹配升级文件，早期版本仍按上述迁移边界处理。
