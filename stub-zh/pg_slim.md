## 用法

来源：

- [README.md](https://github.com/matteuccimarco/pg-slim/blob/1ec4df3a35c395d9d0d7ba25c963775fbfa80088/README.md)
- [sql/pg_slim--0.1.0.sql](https://github.com/matteuccimarco/pg-slim/blob/1ec4df3a35c395d9d0d7ba25c963775fbfa80088/sql/pg_slim--0.1.0.sql)
- [pg_slim.control](https://github.com/matteuccimarco/pg-slim/blob/1ec4df3a35c395d9d0d7ba25c963775fbfa80088/pg_slim.control)

`pg_slim` 0.1.0 新增 `slim` 类型、JSONB 转换、字段访问、包含判断及比较算子。这是早期存储格式实现，不能保证替代 JSONB 后一定节省空间。

### 核心用法

```sql
CREATE EXTENSION pg_slim;
SELECT slim_decode(slim_encode('{"name":"Alice","age":30}'::jsonb));
SELECT slim_encode('{"name":"Alice"}'::jsonb)->>'name';
```

### 运行边界

`slim_encode` 与 `slim_decode` 在 JSONB 和 slim 之间转换；`slim_typeof`、`slim_array_length`、`slim_object_keys` 和 `slim_pretty` 用于检查值。扩展提供 B-tree 与 hash 算子类，但尚未实现 GIN 算子类。部分操作会在内部解码为 JSONB，更新须替换整个值。

由超级用户安装。源码未要求预加载，也未声明完整主版本支持矩阵；安装示例使用 PostgreSQL 15。迁移应用数据前应衡量实际存储与查询成本，并验证双向转换。
