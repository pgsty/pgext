


## 用法

> [pg_ivm: PostgreSQL 增量视图维护](https://github.com/sraoss/pg_ivm)

`pg_ivm` 扩展提供增量视图维护（IVM）功能，通过仅应用增量变更而非完全重算来更新物化视图。视图在基表修改时通过 AFTER 触发器立即更新。

```sql
CREATE EXTENSION pg_ivm;
```

### 配置

将 `pg_ivm` 添加到预加载库以确保正确维护：

```ini
shared_preload_libraries = 'pg_ivm'
```

### 函数

#### create_immv

```sql
pgivm.create_immv(immv_name text, view_definition text) RETURNS bigint
```

创建增量可维护物化视图（IMMV）。触发器会自动创建以保持视图更新。如果可能，会自动创建唯一索引。

```sql
SELECT pgivm.create_immv('myview', 'SELECT * FROM mytab');
```

#### refresh_immv

```sql
pgivm.refresh_immv(immv_name text, with_data bool) RETURNS bigint
```

完全替换 IMMV 内容。`with_data = false` 时，IMMV 变为未填充状态且触发器被移除。`with_data = true` 时，触发器和索引会重新创建。

```sql
SELECT pgivm.refresh_immv('myview', true);
```

#### get_immv_def

```sql
pgivm.get_immv_def(immv regclass) RETURNS text
```

返回 IMMV 的重构 SELECT 命令。

### IMMV 元数据目录

`pgivm.pg_ivm_immv` 目录存储 IMMV 信息：

| 列 | 类型 | 描述 |
|----|------|------|
| `immvrelid` | regclass | IMMV 的 OID |
| `viewdef` | text | 视图定义的查询树 |
| `ispopulated` | bool | IMMV 当前是否已填充 |

### 示例

创建带聚合的 IMMV：

```sql
SELECT pgivm.create_immv('immv_agg',
    'SELECT bid, count(*), sum(abalance), avg(abalance)
     FROM pgbench_accounts JOIN pgbench_branches USING(bid) GROUP BY bid');
```

基表更新会自动反映：

```sql
UPDATE pgbench_accounts SET abalance = abalance + 1000 WHERE aid = 4112345;
SELECT * FROM immv_agg WHERE bid = 42;  -- 聚合值已自动更新
```

列出所有 IMMV：

```sql
SELECT immvrelid AS immv, pgivm.get_immv_def(immvrelid) AS def
FROM pgivm.pg_ivm_immv;
```

使用 `DROP TABLE` 删除 IMMV：

```sql
DROP TABLE myview;
```

### 禁用/启用维护

在批量修改前禁用即时维护，然后刷新：

```sql
SELECT pgivm.refresh_immv('myview', false);   -- 禁用
-- ... 批量修改 ...
SELECT pgivm.refresh_immv('myview', true);    -- 刷新并重新启用
```

### 支持的查询功能

- 内连接和外连接（包括自连接）
- `DISTINCT` 子句
- 聚合函数：`count`、`sum`、`avg`、`min`、`max`
- `FROM` 子句中的简单子查询
- `WHERE` 子句中的 `EXISTS` 子查询
- 简单 CTE（`WITH` 查询）
- `GROUP BY` 子句
