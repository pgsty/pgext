## 用法

来源：

- [官方来源](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/README.md)
- [官方来源](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/pg_array_utils.control)
- [官方来源](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/pg_array_utils--1.0.sql)
- [官方来源](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/design_doc.md)
- [官方来源](https://github.com/UuAcC/pgpro_practice_repos/blob/0ec613aeb5e6cfeb6bba202a5592940bf5f86cab/pg_array_utils/sql/init__.sql)

`pg_array_utils` 1.0 是 PostgreSQL 教学仓库中的 C 数组转换扩展。文档环境为 PostgreSQL 18.4，尚未确认更广泛的主版本支持与分发许可证。

### 数组转换

扩展可迁移 schema，无需预加载。其 C 函数使用普通扩展安装权限。

```sql
CREATE EXTENSION pg_array_utils;
SELECT array_flatten(ARRAY[[1,2],[3,4]]);
SELECT array_chunk(ARRAY[1,2,3,4], 2);
SELECT array_compact(ARRAY[1,NULL,2]);
SELECT array_slice(ARRAY[1,2,3,4], 1, 4, 2);
```

### 函数与限制

`array_flatten` 展平维度，`array_chunk` 对一维数组分块，`array_zip` 将可变数量的数组组合为 JSONB 元组，`array_compact` 删除指定空值，`array_pad` 填充到目标大小，`array_slice` 按步长切片，`array_repeat` 重复标量或数组。填充与重复操作分别提供标量和数组重载。函数声明为 `IMMUTABLE STRICT`，SQL NULL 参数会直接返回 NULL。

`array_chunk` 与 `array_compact` 拒绝多维输入。应遵循文档中的下标、维度及类型规则，不要假设采用 Python 切片语义。大结果仍受 PostgreSQL 分配上限限制，这不是流式数组处理器。
