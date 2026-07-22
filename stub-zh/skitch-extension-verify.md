## 用法

来源：

- [扩展控制文件](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/verify/skitch-extension-verify.control)
- [0.0.7 版扩展 SQL](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/verify/sql/skitch-extension-verify--0.0.7.sql)
- [官方 verify 软件包目录](https://github.com/airpage-app/pg-utils/tree/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/verify)

`skitch-extension-verify` 0.0.7 是一组 PL/pgSQL 部署检查函数。目录查询找到目标对象时，每个函数返回 `true`；检查失败时则抛出异常，因此它适合受控迁移验证，而不是交互式发现。

### 核心流程

扩展依赖 `plpgsql`、`uuid-ossp` 和 `skitch-extension-utils`：

```sql
CREATE EXTENSION "skitch-extension-verify" CASCADE;

SELECT verify_schema('app');
SELECT verify_table('app.orders');
SELECT verify_constraint('app.orders', 'orders_pkey');
SELECT verify_table_grant('app.orders', 'SELECT', 'reporter');
```

对于会把名称拆为模式和实体的检查，应使用模式限定对象字符串。成功调用返回 `true`；对象不存在时通常会以函数的 `Nonexistent ...` 异常中止当前语句。

### 验证函数

- 对象：`verify_schema`、`verify_table`、`verify_view`、`verify_type`、`verify_domain`、`verify_extension` 和 `verify_function`。
- 表细节：`verify_constraint`、`verify_index`、`verify_trigger`、`verify_policy` 和 `verify_security`。
- 角色与权限：`verify_role`、`verify_membership` 和 `verify_table_grant`。

这些函数依赖 `skitch-extension-utils` 提供的 `get_schema_from_str`、`get_entity_from_str`、`list_indexes` 和 `list_memberships` 等辅助例程。

### 正确性边界

这些函数并不是完整的安全或模式差异引擎。在所审计 SQL 中，`verify_extension` 查询 `pg_available_extensions`，所以它证明扩展可供安装，而不是已经安装在当前数据库。`verify_function` 检查匹配函数行是否存在，却不要求返回的 `has_function_privilege` 值为 true。

若干例程使用宽泛或脆弱的目录连接，而且所有函数都被错误声明为 `IMMUTABLE`，尽管它们读取会变化的目录和权限。`verify_security` 的失败分支还引用了未定义的 `_name` 变量。应把它们视为迁移断言，并针对确切检查和 PostgreSQL 版本审计 SQL，而不是视为合规证据。应避免复用预备计划或常量表达式，并用直接目录查询独立验证安全敏感结果。
