## 用法

来源：

- [阿里云 RDS 扩展矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [阿里云 RDS 网格模型指南](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/grid-model)
- [阿里云 RDS ST_AsGrid 参考](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/st-asgrid)

ganos_geomgrid 是阿里云 ApsaraDB RDS for PostgreSQL 中由云厂商管理的 GanosBase 组件。它提供 GeoSOT 地理网格与 H3 六边形网格，用于多分辨率编码、索引、聚合与空间关系查询。云厂商矩阵为较新的受支持引擎列出 7.0 版，其他 PostgreSQL 大版本则使用较旧的 6.x 版。

### 基本用法

在 RDS 实例上创建该扩展，并让 CASCADE 安装由云厂商管理的依赖：

```sql
CREATE EXTENSION ganos_geomgrid WITH SCHEMA public CASCADE;

SELECT ST_AsText(g)
FROM unnest(
    ST_AsGrid(
        ST_GeomFromText(
            'POINT(116.31522216796875 39.910277777777778)',
            4490
        ),
        15
    )
) AS g;
```

ST_AsGrid 会返回与几何对象相交的 GeoSOT 网格，精度范围为 1 至 32。阿里云文档将 SRID 4490 规定为原生坐标系；其他受支持的输入会自动转换到该坐标系。该扩展还公开 geomgrid 与 h3grid 类型，以及转换、层级、空间关系和索引操作。

### 注意事项

- 这是阿里云服务扩展，不是独立发布的开源软件包。可用性、确切版本、权限与依赖行为取决于 RDS PostgreSQL 大版本和当前实例策略。
- 部署前应在云厂商矩阵中确认确切引擎版本是否受支持。如果 CREATE EXTENSION 被拒绝，应升级实例小版本或联系阿里云支持。
- 网格精度同时影响准确性与结果规模。对大型或复杂几何对象使用高精度可能产生大量网格，应对存储、索引大小和查询成本做基准测试。
- GeoSOT 处理以 CGC2000/SRID 4490 为中心。在混用其他空间参考系的数据前，应验证坐标转换与大地基准假设。
