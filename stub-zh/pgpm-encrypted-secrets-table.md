## 用法

来源：

- [packages/encrypted-secrets-table/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets-table/README.md)
- [packages/encrypted-secrets-table/sql/pgpm-encrypted-secrets-table--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets-table/sql/pgpm-encrypted-secrets-table--0.47.0.sql)
- [packages/encrypted-secrets-table/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets-table/Makefile)
- [packages/encrypted-secrets-table/pgpm-encrypted-secrets-table.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets-table/pgpm-encrypted-secrets-table.control)

`pgpm-encrypted-secrets-table` 提供供上层秘密值模块使用的 `secrets_schema.secrets_table` 及触发器，保存所有者标识、名称和编码后的值。

### 核心用法

```sql
CREATE EXTENSION "pgpm-encrypted-secrets-table" CASCADE;
SELECT count(*) FROM secrets_schema.secrets_table;
```

### 运行边界

PGP 路径以所有者 UUID 作为口令，并非独立保护的密钥；能同时读取标识符和密文的用户因此可以解密。Crypt 值是单向密码散列。这种设计不是密钥库，也不能防御数据库读取者，应限制表和函数权限，不能将标识符当作秘密。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`pgcrypto`、`plpgsql`、`pgpm-verify`。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
