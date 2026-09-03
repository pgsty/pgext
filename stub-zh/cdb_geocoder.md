## 用法

来源：

- [官方文档](https://github.com/CartoDB/data-services/blob/11ecbff14d81b71af4cc5f87c503bd71084fcefc/geocoder/extension/README.md)
- [扩展控制文件](https://github.com/CartoDB/data-services/blob/11ecbff14d81b71af4cc5f87c503bd71084fcefc/geocoder/extension/cdb_geocoder.control)
- [官方仓库](https://github.com/CartoDB/data-services)

`cdb_geocoder` 已归档的 CARTO 地理编码扩展，支持行政区、邮编、IP 与地名查询。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `cdb_geocoder`：

```sql
CREATE EXTENSION cdb_geocoder CASCADE;
```

经审查的控制文件或官方流程要求 `cartodb`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
set statement_timeout = '20min';
INSERT INTO ip_address_locations (the_geom, network_start_ip) SELECT the_geom, ('::ffff:' || split_part(network, '/', 1))::inet FROM latest_ip_address_locations;
INSERT INTO ip_address_locations (the_geom, network_start_ip) SELECT the_geom, split_part(network, '/', 1)::inet FROM latest_ip6_address_locations;
```

### 主要对象

官方来源通过动态方式或供应商工具定义扩展接口；授权前应检查实际安装版本。

### 运维与边界

- 目录生命周期为 archived；生产使用前应测试升级、备份恢复与服务器兼容性。
