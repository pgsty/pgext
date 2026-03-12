

## 用法

> [pg_bulkload: pg_bulkload 是 PostgreSQL 的高速数据加载工具](https://github.com/ossc-db/pg_bulkload)

一个用于 PostgreSQL 的高速数据加载工具，通过绕过共享缓冲区实现海量数据的快速加载，并内置 ETL 功能用于输入验证和数据转换。

### 基本用法

使用控制文件加载数据：

```bash
pg_bulkload sample_csv.ctl
```

输出：

```
NOTICE: BULK LOAD START
NOTICE: BULK LOAD END
    0 Rows skipped.
    8 Rows successfully loaded.
    0 Rows not loaded due to parse errors.
    0 Rows not loaded due to duplicate errors.
    0 Rows replaced with new rows.
```

### 控制文件示例

```ini
# sample_csv.ctl
OUTPUT = my_table
INPUT = /path/to/data.csv
TYPE = CSV
DELIMITER = ,
QUOTE = "\""
ESCAPE = "\""
NULL = ""
SKIP = 1              # 跳过表头行
PARSE_ERRORS = 100    # 允许最多 100 个解析错误
DUPLICATE_ERRORS = 0  # 遇到重复键错误时拒绝
ON_DUPLICATE_KEEP = NEW  # 或 OLD
TRUNCATE = NO
```

### 加载模式

- **DIRECT**：绕过共享缓冲区，直接写入数据文件（最快）
- **PARALLEL**：使用多个进程进行加载
- **CSV/BINARY/FIXED**：支持多种输入格式

### SQL 接口

```sql
-- 在 SQL 中加载数据
SELECT pg_bulkload(
    'OUTPUT = my_table, INPUT = /path/to/data.csv, TYPE = CSV'
);
```

### 主要功能

- 绕过 PostgreSQL 共享缓冲区以实现最大吞吐量
- 输入数据验证，支持可配置的错误阈值
- 重复键处理（保留新值、保留旧值或拒绝）
- 支持 CSV、定长和二进制输入格式
- 跳过行、过滤函数用于数据转换
- 支持并行加载

### 文档

完整文档：http://ossc-db.github.io/pg_bulkload/index.html
