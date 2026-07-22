## 用法

来源：

- [阿里云支持扩展矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [阿里云扩展管理文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/manage-plug-ins)
- [官方 Tiger Geocoder 文档](https://postgis.net/docs/Extras.html#Tiger_Geocoder)
- [官方 geocode 函数参考](https://postgis.net/docs/Geocode.html)

`ganos_tiger_geocoder` 是阿里云 GanosBase 中处理美国人口普查局 TIGER 数据、正向地理编码与反向地理编码的扩展。它是托管服务组件，并不是可以从上游源码仓库安装到社区 PostgreSQL 的扩展。

### 核心流程

在阿里云 RDS PostgreSQL 上，先确认实例引擎版本提供 7.0，让目标数据库归指定高权限账号所有，再通过阿里云扩展管理流程安装 `ganos_tiger_geocoder`。GanosBase 与 PostGIS 扩展不能安装在同一个模式中。

地理编码器必须先加载 TIGER 查询数据和州级数据，地址匹配才有实际意义。按厂商支持的加载流程填充数据并把 `tiger` 模式加入 `search_path` 后，可以通过 Tiger 接口执行正向地理编码：

```sql
SET search_path = public, tiger;

SELECT pprint_addy(addy) AS normalized_address,
       rating,
       ST_AsText(geomout) AS point
FROM geocode('1 Devonshire Place, Boston, MA 02109', 5);
```

主要用户接口包括返回地址候选与排序的 `geocode`、查询点附近地址的 `reverse_geocode`、处理地址的 `normalize_address` 与 `pprint_addy`，以及 `tiger` 模式中的加载和索引维护辅助函数。正向地理编码结果的 `rating` 越低，匹配通常越好。

### 托管服务边界

可用性和版本取决于阿里云产品、地域、PostgreSQL 大版本与小版本引擎，应在实际实例上查询 `pg_available_extensions`，不要只依据目录推断。服务文档要求该扩展的目标数据库归高权限账号所有。加载全国或州级 TIGER 数据会消耗大量存储和维护时间，地理编码质量也取决于已加载的人口普查数据年份与本地数据质量。

链接的 PostGIS 文档说明阿里云指向的 Tiger SQL 模型，但 Ganos 构建及生命周期由阿里云控制。安装、数据加载、升级、备份与删除都应遵循厂商控制台和支持文档；不要替换为社区 PostGIS 文件，也不要期望 `ganos_tiger_geocoder` 可在自管 PostgreSQL 上运行。
