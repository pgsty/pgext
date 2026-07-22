## 用法

来源：

- [Official README](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/README.md)
- [Official quick start](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/docs/quickstart.md)
- [API reference](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/docs/api-reference.md)
- [Configuration guide](https://github.com/zvectorlabs/pgauthz/blob/a24a3bae33d1bac320c233e99ceb9153a4d0667d/docs/configuration.md)

pgauthz 在 PostgreSQL 内实现 Zanzibar 风格的关系型授权。策略定义对象与主体类型、关系、计算权限、条件和通配符；随后由关系元组驱动检查。它面向 PostgreSQL 16 到 18，并在 `authz` 模式中保存策略、元组、修订、变更日志和断言数据。

### 核心流程

定义策略、写入关系，并在与应用数据相同的数据库事务边界内检查：

```sql
CREATE EXTENSION pgauthz;

SELECT pgauthz_define_policy('
  type user {}
  type document {
    relations
      define viewer: [user]
      define editor: [user]
    permissions
      define view = viewer + editor
  }
');

SELECT pgauthz_add_relation(
  'document', 'doc1', 'viewer', 'user', 'alice'
);

SELECT pgauthz_check(
  'document', 'doc1', 'view', 'user', 'alice'
);
```

当已声明条件需要运行时 JSON 上下文时使用 `pgauthz_check_with_context`，并在保护资源的同一条语句或 security-definer 边界中强制执行布尔结果。

### 重要对象

- `pgauthz_define_policy` 解析授权模型并为当前模型建立版本。
- `pgauthz_add_relation` 和 `pgauthz_write_relationships` 创建或修改关系元组。
- `pgauthz_check` 和 `pgauthz_check_with_context` 计算关系或计算权限。
- `pgauthz_list_objects` 和 `pgauthz_list_subjects` 枚举可达资源或主体。
- `pgauthz_expand` 返回用于调试的权限树。
- `pgauthz_read_relationships`、`pgauthz_read_latest_policy` 和 `pgauthz_read_changes` 暴露存储状态与变更。
- `authz.*` GUC 控制检查策略、缓存 TTL 与容量、修订量化、跟踪和 OpenTelemetry 导出。

### 安全与一致性说明

授权函数本身不会强制访问控制，除非每条数据路径都调用它，或使用正确的 RLS/策略包装。应撤销应用角色对受保护表和底层 `authz` 关系的直接访问。缓存默认禁用；启用 TTL 或修订量化会用新鲜度换取速度，因此必须明确测试撤权延迟。策略变化和关系写入都应作为安全敏感操作审计；还要校验条件上下文类型，并防止 OpenTelemetry 端点泄露授权元数据。
