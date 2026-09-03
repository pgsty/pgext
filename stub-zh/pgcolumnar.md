## 用法

来源：

- [官方文档](https://github.com/commandprompt/pgcolumnar/blob/0e4884c18a678bf0a990e6a7dfcdd47248d9111d/README.md)
- [扩展控制文件](https://github.com/commandprompt/pgcolumnar/blob/0e4884c18a678bf0a990e6a7dfcdd47248d9111d/pgcolumnar.control)
- [官方仓库](https://github.com/commandprompt/pgcolumnar)

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
