## 用法

来源：

- [上游 README](https://github.com/citusdata/cstore_fdw/blob/90e22b62fbee6852529104fdd463f532cf7a3311/README.md)
- [扩展 control 文件](https://github.com/citusdata/cstore_fdw/blob/90e22b62fbee6852529104fdd463f532cf7a3311/cstore_fdw.control)

`cstore_fdw` 以列式格式存储 PostgreSQL 外部表，适合批量装载的分析数据。它支持列裁剪、基于最小值/最大值的跳过索引以及可选的 `pglz` 压缩。

创建扩展前必须预加载其动态库：

```conf
shared_preload_libraries = 'cstore_fdw'
```

修改后重启 PostgreSQL，再创建 FDW 对象：

```sql
CREATE EXTENSION cstore_fdw;
CREATE SERVER cstore_server FOREIGN DATA WRAPPER cstore_fdw;

CREATE FOREIGN TABLE measurements (
    measured_at timestamptz,
    sensor_id bigint,
    reading double precision
)
SERVER cstore_server
OPTIONS (compression 'pglz', stripe_row_count '150000');
```

应使用 `COPY` 或 `INSERT INTO ... SELECT ...` 批量装载数据，随后执行 `ANALYZE`。上游 `1.7` 版本支持 PostgreSQL 9.3 至 12，不支持 `UPDATE`、`DELETE` 或单行插入。

### 已弃用状态

上游已将列式存储迁移到 Citus，并改用 PostgreSQL 表访问方法 API。上游建议把现有 `cstore_fdw` 数据迁移到 Citus 列式表，以获得流复制、回滚和更简便的 `pg_upgrade` 等能力。`cstore_fdw` 主要适合维护或迁移遗留部署。
