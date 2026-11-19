## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tsfunc/pg_tsfunc.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tsfunc/pg_tsfunc--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tsfunc/meson.build)

`pg_tsfunc` `1.0` 提供时序重采样、降采样与数值编解码。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_tsfunc;
SELECT ts_decode_floats(ts_encode_floats(ARRAY[1.0,2.0,3.0]::float8[]));
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

ts_lttb 与 ts_minmax 用于降采样，ts_resample 对齐观测，ts_asof 与 ts_asof_join 提供时间查找。时间戳、浮点编解码器将序列存为 bytea，并不提供表访问方法。科学计算需核对顺序、重复值处理和最大间隙设置。数组和聚合操作仍会占用后端内存，压缩率取决于数据。

此模块对应上游旧名 `alohadb_tsfunc`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
