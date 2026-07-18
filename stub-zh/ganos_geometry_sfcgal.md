## 用法

来源：

- [阿里云 RDS PostgreSQL 扩展矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_geometry_sfcgal` 在目录中是阿里云 RDS PostgreSQL 提供的 Ganos 7.0 组件。它属于云厂商托管软件；本次复核没有取得可移植扩展所需的公开源码树、control 文件或 SQL API。

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'ganos_geometry_sfcgal';
```

该查询只用于发现当前 RDS 引擎实际暴露的内容。可用性可能随 RDS 引擎版本、版本类型、地域与维护策略变化。只能按精确实例对应的厂商支持流程启用，不要照搬其他 SFCGAL 项目的安装命令或函数名。

公开矩阵能够确认可用性，但已复核材料没有记录独立函数接口。依赖它之前，应检查 `pg_available_extension_versions`、控制台提供的厂商文档，并在可丢弃的 RDS 实例上验证。迁移设计还必须考虑该组件在阿里云之外可能不存在。
