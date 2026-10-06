## 用法

来源：

- [packages/encrypted-secrets/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets/README.md)
- [packages/encrypted-secrets/sql/pgpm-encrypted-secrets--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets/sql/pgpm-encrypted-secrets--0.47.0.sql)
- [packages/encrypted-secrets/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets/Makefile)
- [packages/encrypted-secrets/pgpm-encrypted-secrets.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets/pgpm-encrypted-secrets.control)

`pgpm-encrypted-secrets` 基于独立的秘密值表扩展，增加读取、写入或更新、验证和删除操作，支持 PGP 编码与 crypt 散列模式。

### 核心用法

```sql
CREATE EXTENSION "pgpm-encrypted-secrets" CASCADE;
SELECT encrypted_secrets.secrets_getter('00000000-0000-0000-0000-000000000001'::uuid, 'demo', NULL);
```

### 运行边界

`encrypted_secrets.secrets_upsert`、`encrypted_secrets.secrets_verify` 和 `encrypted_secrets.secrets_delete` 接收所有者 UUID 与秘密名称。PGP 实现以同一行保存的所有者 UUID 作为口令，无法对读取该标识符的人保护密文。必须另行实施应用授权和 SQL 权限，且不要记录真实秘密参数。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`pgcrypto`、`plpgsql`、`pgpm-encrypted-secrets-table`、`pgpm-verify`。 安装 SQL 要求平台角色 `authenticated` 已存在，须先使用上游角色初始化流程，并审查授权。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
