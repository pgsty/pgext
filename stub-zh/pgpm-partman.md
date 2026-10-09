## 用法

来源：

- [packages/partman/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/partman/README.md)
- [packages/partman/sql/pgpm-partman--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/partman/sql/pgpm-partman--0.47.0.sql)
- [packages/partman/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/partman/Makefile)
- [packages/partman/pgpm-partman.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/partman/pgpm-partman.control)

`pgpm-partman` 为已有的 pg_partman 安装增加分区辅助函数与元模式触发器。`partman.create_parent_with_retention` 配置父表及保留策略，按 ID 操作的辅助函数通过应用目录解析表名。

### 核心用法

```sql
CREATE EXTENSION "pgpm-partman";
SELECT parent_table, retention, retention_keep_table
FROM partman.part_config;
```

### 运行边界

须先在 `partman` 模式中安装 pg_partman；固定版本的封装 SQL 不会创建它。`partman.verify_parent_by_id` 检查登记状态，`partman.run_maintenance_by_id` 执行维护。安装会在 `metaschema_public.partition` 上附加 INSERT 与 DELETE 触发器；删除分区记录会移除配置，并可能删除配置中记录的模板表。多个维护函数采用 `SECURITY DEFINER`，应限制执行权限及目录写入权限。兼容性取决于实际安装的 pg_partman API 和相匹配的元模式布局。 依赖为 `plpgsql`, `pg_partman`, `metaschema-schema`。 安装脚本会向预先存在的角色 `authenticated` 授权。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
