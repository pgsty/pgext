## 用法

来源：

- [官方 control 文件](https://github.com/ttfkam/pg_audit/blob/c7f3069058215873ae8e497e1476c1e195309479/audit.control)
- [官方 1.0.0 版 SQL](https://github.com/ttfkam/pg_audit/blob/c7f3069058215873ae8e497e1476c1e195309479/audit--1.0.0.sql)
- [官方代码仓库](https://github.com/ttfkam/pg_audit)

`audit` 是一个小型纯 SQL 行变更审计原型。它提供 `audit` 表和触发器函数，设计目标是在 `INSERT`、`UPDATE` 与 `DELETE` 后把新行或旧行保存为 `jsonb`。版本 `1.0.0` 是年代较久且未完成的产物；依赖它之前必须检查并修补其 SQL。

### 核心流程

预期用法是安装扩展，然后把 `audit()` 触发器函数挂到每个需要行级历史的表上：

```sql
CREATE EXTENSION audit;

CREATE TRIGGER accounts_audit
AFTER INSERT OR UPDATE OR DELETE ON accounts
FOR EACH ROW EXECUTE FUNCTION audit();

SELECT id, relation::regclass, op, row_data, archived
FROM audit
ORDER BY id DESC;
```

插入和更新时，触发器预期保存 `NEW`；删除时保存 `OLD`。`archived` 标志需要由应用管理，并不代表自动保留策略。

### 主要对象

- `audit`：审计行表，包含 `jsonb` 载荷、关系 OID、操作枚举与归档标志。
- `operation`：取值为 `INSERT`、`UPDATE`、`DELETE` 的枚举。
- `audit()`：行级触发器函数。
- `audit_id(jsonb)`：把 `id` 属性提取为 `bigint`。
- `audit_oid(name, name)`：设计用于把模式名和关系名解析为 OID。
- `oid_metadata`：关系与模式元数据视图。

### 注意事项

发布的 `1.0.0` SQL 内部存在不一致：表字段名是 `op`，触发器却向 `operation` 字段插入；`audit_oid` 中的参数名还遮蔽了查询列名。脚本末尾的 `\echo` 也写成了另一个扩展名。因此，原样发布的触发器示例需要先审查并修正上游 SQL，才能可靠工作。

触发器会保存完整行镜像，但没有清理、抗篡改、操作者/会话元数据或 DDL 审计能力。应限制审计表访问权限，明确制定保留策略，并先在一次性数据库中测试存储量与写放大。
