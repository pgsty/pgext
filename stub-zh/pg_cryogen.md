## 用法

来源：

- [上游存储与配置文档](https://github.com/adjust/pg_cryogen/blob/4c8b4f382215861e613b1e9be66233a10e058af5/README.md)
- [扩展控制文件](https://github.com/adjust/pg_cryogen/blob/4c8b4f382215861e613b1e9be66233a10e058af5/pg_cryogen.control)

`pg_cryogen` 是适用于 PostgreSQL 12 及以上版本的实验性只追加表访问方法。它用 `lz4` 或 `zstd` 压缩堆元组，并实现索引扫描与位图扫描。

```sql
CREATE EXTENSION pg_cryogen;
SET pg_cryogen.compression_method = 'zstd';
CREATE TABLE events (id bigint, payload jsonb) USING pg_cryogen;
INSERT INTO events VALUES (1, '{"ok":true}');
```

只有 `INSERT` 与 `COPY` 能修改数据，且已复核实现每个事务只允许向一个此类表插入。每条语句至少创建一个独立数据块，因此应优先批量加载。仓库已归档；必须固定 PostgreSQL 大版本、验证备份恢复，并且不能假定它支持更新、删除、vacuum 行为、复制或未来的访问方法 API。
