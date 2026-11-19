## 用法

来源：

- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_columnar/README.md)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_columnar/kingres_columnar.control)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_columnar/kingres_columnar--0.1.sql)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_columnar` 0.1 是实验性的仅插入列式表访问方法。所审查源码明确警告不要保存重要数据：尚未解决的缺陷涉及可串行化事务与写缓冲区。此处将其列为预览，不代表经过生产存储验证。

### 创建测试表

源码面向 PostgreSQL 18/19。由于库会注册 WAL 资源管理器，所有写节点、备库与恢复节点都要预加载并重启。安装需要超级用户。

```ini
shared_preload_libraries = 'kingres_columnar'
```

```sql
CREATE EXTENSION kingres_columnar;
CREATE TABLE public.events (ts timestamptz, value float8) USING columnar;
SELECT * FROM kingres.columnar_info('public.events');
SELECT * FROM kingres.columnar_check(false);
```

### 存储与运行

`columnar` 访问方法使用 `none`、`pglz`、`lz4` 或 `zstd` 压缩条带，支持串行投影扫描和最小/最大值跳过。它不支持索引，也不实现普通行更新或删除。`kingres.columnar_set_options` 仅影响新条带，限所有者调用且非事务性。`kingres.columnar_compact` 合并小条带，`kingres.columnar_stripes` 与 `kingres.columnar_options` 用于检查。读取工具要求 SELECT 权限。

`kingres.columnar_check(false)` 报告损坏，修复模式与 `kingres.columnar_reconcile` 是超级用户操作。实验性的 WAL 资源管理器 ID 为 128，正式发布前需要保留专用编号。一旦存在列存 WAL，移除库可能导致崩溃恢复、归档恢复或还原失败。没有预加载时可读取部分已有数据，但拒绝写入。该开发分支内部也会改变格式，必须严格匹配修订。
