## 用法

来源：

- [ApsaraDB RDS for PostgreSQL 支持的扩展](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [ST_SrEqual](https://www.alibabacloud.com/help/doc-detail/176141.html)
- [ST_SrReg](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/st-srreg)
- [ST_SrFromEsriWkt](https://www.alibabacloud.com/help/doc-detail/176143.html)

`ganos_spatialref` 是阿里云 GanosBase 空间参考扩展。它可以按语义比较坐标系定义、注册自定义空间参考系，并把 Esri WKT 转为 OGC WKT。它是托管 GanosBase 组件，并非独立社区软件包。

### 核心流程

在提供该扩展的 ApsaraDB RDS for PostgreSQL 实例中：

```sql
CREATE EXTENSION ganos_spatialref;

SELECT ST_SrEqual(
  '+proj=longlat +datum=WGS84 +no_defs',
  '+proj=longlat +ellps=WGS84 +datum=WGS84 +no_defs'
);

SELECT srid, auth_name, auth_srid
FROM spatial_ref_sys
WHERE ST_SrEqual(
  srtext::cstring,
  '+proj=longlat +ellps=GRS80 +no_defs'
)
LIMIT 1;
```

`ST_SrEqual(sr1, sr2, strict DEFAULT true)` 接受 OGC WKT 或 PROJ.4 字符串。它比较投影、椭球体和轴参数，而不是直接比较文本；适用时，严格模式还会比较参考椭球名称。

### 注册与转换

`ST_SrReg(sr)` 注册 WKT/PROJ.4，并返回已有或新分配的 SRID。`ST_SrReg(auth_name, auth_id, sr)` 还会添加权威机构元数据。注册会修改 `spatial_ref_sys`，因此应先验证定义，并只向可信空间管理员开放执行权限。`ST_SrFromEsriWkt(sr)` 把 Esri WKT 转为 OGC WKT，其中包括移除 Esri datum 名称的 `D_` 前缀等规范化。

### 服务与版本边界

可用性取决于 RDS 版本、PostgreSQL 大版本和引擎小版本。引用的服务商矩阵为 PostgreSQL 17 列出 7.0，为多个较旧大版本列出 6.9，对更旧组合列出 6.3；并非每个受支持 PostgreSQL 版本都提供该扩展。使用前应通过 `pg_available_extension_versions` 确认当前实例。坐标系注册、内置定义和函数行为都属于托管 GanosBase 版本，因此引擎升级后应重新测试转换，不能假定所有实例都是目录中的 7.0。
