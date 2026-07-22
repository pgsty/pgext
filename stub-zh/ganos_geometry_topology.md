## 用法

来源：

- [ApsaraDB RDS 官方扩展矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [ApsaraDB RDS 官方扩展管理指南](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/manage-plug-ins)

`ganos_geometry_topology` 是阿里云 GanosBase 的拓扑组件，在受支持的 ApsaraDB RDS for PostgreSQL 实例上进行空间几何分析。它是服务商管理的软件，而不是可移植的社区软件包；版本、可用性、依赖与 SQL 接口均由所选 RDS 引擎决定。

### 启用与验证

使用目标数据库的 RDS 扩展管理页面和高权限 RDS 账户，并在创建前检查确切的可用版本：

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'ganos_geometry_topology';

CREATE EXTENSION ganos_geometry_topology;

SELECT extname, extversion, extnamespace::regnamespace
FROM pg_extension
WHERE extname = 'ganos_geometry_topology';
```

安装后，应使用与版本匹配的 GanosBase 拓扑文档并检查实际安装对象，不要假定其他拓扑软件包中的名称或语义完全相同：

```sql
SELECT n.nspname, p.proname, pg_get_function_identity_arguments(p.oid)
FROM pg_extension AS e
JOIN pg_depend AS d
  ON d.refclassid = 'pg_extension'::regclass
 AND d.refobjid = e.oid
 AND d.deptype = 'e'
JOIN pg_proc AS p
  ON d.classid = 'pg_proc'::regclass
 AND d.objid = p.oid
JOIN pg_namespace AS n ON n.oid = p.pronamespace
WHERE e.extname = 'ganos_geometry_topology'
ORDER BY 1, 2, 3;
```

使用 `pg_depend` 与 RDS 控制台清点其他对象类别；不要根据扩展名猜测 API。

### 服务边界

- RDS 支持矩阵是各 PostgreSQL 大版本和引擎小版本的权威依据。目录中的 `7.0` 之类版本不能证明每个实例都能安装同一版本。
- 阿里云要求使用高权限数据库账户安装 `ganos_geometry_topology`。应用角色只应获得实际需要的拓扑模式、表和函数权限。
- 服务商提醒 GanosBase 与 PostGIS 扩展不能安装在同一模式。启用任一系列前应规划模式，不要手工移动对象。
- 拓扑结构通常包含依赖模式、表和几何对象。卸载或升级应被视为数据迁移：清点依赖、备份、在克隆环境演练，并遵循 RDS 服务流程。
- 公开扩展列表只确认可用性，没有记录此版本的完整 SQL API。使用代码应绑定到已安装版本的 GanosBase 文档，并用代表性空间数据验证结果、坐标参考系、有效性、精度和索引计划。
