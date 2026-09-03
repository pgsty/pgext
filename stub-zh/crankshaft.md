## 用法

来源：

- [官方文档](https://github.com/CartoDB/crankshaft/blob/4d7bc1acb5cd3167cb0c9e8341beb1f32964ae3e/README.md)
- [扩展控制文件](https://github.com/CartoDB/crankshaft/blob/4d7bc1acb5cd3167cb0c9e8341beb1f32964ae3e/release/crankshaft.control)
- [官方仓库](https://github.com/CartoDB/crankshaft)

`crankshaft` 已归档的 CARTO 空间分析扩展，提供聚类、分群与空间统计函数。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `crankshaft`：

```sql
CREATE EXTENSION crankshaft CASCADE;
```

经审查的控制文件或官方流程要求 `plpython3u`, `postgis`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT cdb_crankshaft.cdb_crankshaft_version();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `_cdb_random_seeds` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `cdb_crankshaft.cdb_crankshaft_version` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `cdb_crankshaft.cdb_moran_local` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `cdb_crankshaft.cdb_moran_local_rate` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 扩展会在 `cdb_crankshaft` 下固定或创建模式对象；权限与备份审查应包含这些对象。
- 目录生命周期为 archived；生产使用前应测试升级、备份恢复与服务器兼容性。
