## 用法

来源：

- [README](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/README.md)
- [Control file / 控制文件](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/pgcolumnar.control)
- [SQL](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/pgcolumnar--1.0-alpha4.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/CHANGELOG.md)
- [docs/limitations.md](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/docs/limitations.md)

`pgcolumnar` 原生列式表访问方法，支持压缩、向量化扫描与 Parquet 工作流。

### 启用

将 `pgcolumnar` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pgcolumnar`：

```ini
shared_preload_libraries = 'pgcolumnar'
```

```sql
CREATE EXTENSION pgcolumnar;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE EXTENSION pgcolumnar;

CREATE TABLE events (id bigint, ts timestamptz, kind int, payload text)
  USING pgcolumnar;

INSERT INTO events
  SELECT g, now(), g % 8, 'p' || g
  FROM generate_series(1, 1000000) g;

SELECT count(*), avg(kind) FROM events WHERE kind = 3;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pgcolumnar.projection` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pgcolumnar.bloom` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pgcolumnar.export_arrow` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgcolumnar.export_parquet` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgcolumnar.options` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pgcolumnar.storage` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pgcolumnar.vacuum_sorted` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 15, 16, 17, 18, 19；不要推断未列出的主版本。
- 预加载 `pgcolumnar` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
- 扩展会在 `pgcolumnar` 下固定或创建模式对象；权限与备份审查应包含这些对象。
- 目录生命周期为 preview；生产使用前应测试升级、备份恢复与服务器兼容性。

### Alpha 4 升级

控制版本为 `1.0-alpha4`，PGXN 发行版本为 1.0.0-alpha.4。安装匹配文件后，应在每个数据库中运行 `ALTER EXTENSION pgcolumnar UPDATE`，仅替换共享库并不足够。Alpha 4 增加 Hilbert 聚簇和更多正确性修复。PostgreSQL 19 的验证基于 beta2。应保留可重新加载的原始数据，上游仍未普遍承诺未来磁盘格式变更的兼容性。部分旧限制页面的发布标签尚未更新，版本以控制文件和当前变更日志为准。
