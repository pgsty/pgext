## 用法

来源：

- [官方上游 README](https://github.com/janssenproject/jans/blob/2d578f00e5408a213d70402f10205b09467d16ed/jans-cedarling/README.md)
- [官方扩展控制文件 (cedarling_pg.control)](https://github.com/janssenproject/jans/blob/2d578f00e5408a213d70402f10205b09467d16ed/jans-cedarling/cedarling_pg/cedarling_pg.control)
- [官方扩展 SQL (cedarling_pg--0.1.0.sql)](https://github.com/janssenproject/jans/blob/2d578f00e5408a213d70402f10205b09467d16ed/jans-cedarling/cedarling_pg/sql/cedarling_pg--0.1.0.sql)

`cedarling_pg` — PostgreSQL 集成用于 Cedarling 授权决策，包括 JWT 意识检查和 RLS 方向的帮助程序。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION cedarling_pg;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `cedarling.entity_map` 是由扩展安装或管理的表。
- `cedarling.mask_rules` 是由扩展安装或管理的表。
- `cedarling.policy_history` 是由扩展安装或管理的表。
- `cedarling.policy_versions` 是由扩展创建的模式。
- `cedarling` 是由扩展创建的模式。

### 要求与注意事项

- 该目录记录版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
