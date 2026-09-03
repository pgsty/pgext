## 用法

来源：

- [官方文档](https://github.com/davinderatgithub/optimized-row-format/blob/d3d28fbb2b6ae5e826da92bd8bd0010275daba12/README.md)
- [扩展控制文件](https://github.com/davinderatgithub/optimized-row-format/blob/d3d28fbb2b6ae5e826da92bd8bd0010275daba12/optimized_row_format.control)
- [官方仓库](https://github.com/davinderatgithub/optimized-row-format)

`optimized_row_format` 面向紧凑且缓存友好行存储的实验性表访问方法。

### 启用

将 `optimized_row_format` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `optimized_row_format`：

```ini
shared_preload_libraries = 'optimized_row_format'
```

```sql
CREATE EXTENSION optimized_row_format;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Load the extension
CREATE EXTENSION optimized_row_format;

-- Create a test table
CREATE TABLE test_table (
    id integer,
    name text,
    value bigint,
    flag boolean
) USING optimized_row_format;

-- Insert test data
INSERT INTO test_table VALUES (1, 'test', 100, true);

-- Query the data
SELECT * FROM test_table;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `optimized_row_format_tableam_handler` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 17；不要推断未列出的主版本。
- 预加载 `optimized_row_format` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
