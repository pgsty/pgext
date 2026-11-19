## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ndarray/pg_ndarray.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ndarray/pg_ndarray--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ndarray/meson.build)

`pg_ndarray` `1.0` 提供分块压缩的多维数组。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_ndarray;
SELECT ndarray_create('demo', 'float32', '{4,4}', '{2,2}', fill_value => '0');
SELECT ndarray_read('demo', '{0,0}', '{2,2}', '{1,1}');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

数组元数据和 zstd 压缩块保存在普通表中。ndarray_write、ndarray_reduce、ndarray_map 和 ndarray_storage 提供更新、计算和检查，内核 tensor 辅助函数属于 AlohaDB。切片采用从零开始的 start/stop/step 边界，未写入块返回配置的填充值。大规模使用前需检查维数、分块大小、写锁与角色访问权限。

此模块对应上游旧名 `alohadb_ndarray`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
