

## 用法

> [pg_csv: 灵活的 CSV 处理聚合函数](https://github.com/PostgREST/pg_csv)

提供 CSV 聚合函数，可与 SQL 表达式组合使用，不同于需要特殊协议的 COPY 命令。符合 RFC 4180 标准，正确处理引号。

```sql
CREATE EXTENSION pg_csv;
```

### 函数

| 函数 | 说明 |
|---|---|
| `csv_agg(record)` | 将行聚合为带表头的 CSV 文本 |
| `csv_agg(record, csv_options(...))` | 使用自定义选项进行聚合 |
| `csv_options(delimiter, bom, header, nullstr)` | 配置 CSV 输出选项 |

### 示例

```sql
CREATE TABLE projects AS SELECT * FROM (VALUES
  (1, 'Death Star OS', 1),
  (2, 'Windows 95 Rebooted', 1),
  (3, 'Project "Comma,Please"', 2)
) AS _(id, name, client_id);

-- 基本 CSV 输出
SELECT csv_agg(x) FROM projects x;
-- id,name,client_id
-- 1,Death Star OS,1
-- 2,Windows 95 Rebooted,1
-- 3,"Project ""Comma,Please""",2

-- 管道符分隔
SELECT csv_agg(x, csv_options(delimiter := '|')) FROM projects x;

-- 制表符分隔
SELECT csv_agg(x, csv_options(delimiter := E'\t')) FROM projects x;

-- 添加 BOM 以兼容 Excel
SELECT csv_agg(x, csv_options(bom := true)) FROM projects x;
```
