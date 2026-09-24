## 用法

来源：

- [README](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/README.md)
- [Control](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgvault/pgvault.control)
- [SQL](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgvault/pgvault--0.1.0.sql)
- [Bootstrap and commands](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/scripts/secrets.sh)

`pgvault` 是 PGEverything 的数据库级加密秘密存储。0.1.0 是尚未发布的源码快照，使用 `pgsodium` 认证加密与 `pgcrypto` 口令哈希。Vault 用户是应用层主体，与 PostgreSQL 登录角色相互独立。

### 启用与使用

上游说明的部署环境是 PostgreSQL 16 容器。在启动启用 pgsodium 的服务器前，应设置私有的 `PGE_VAULT_KEY`，不能使用不安全的开发默认值。必须在数据库外保存可恢复的密钥副本，仅恢复密文而没有匹配根密钥并不足够。

```sh
make secrets-init
make secret-user-create
make secrets-create
make secret-read
```

这些交互式辅助命令会询问数据库与凭据。初始化创建 `pgcrypto`、`age`、`pgsodium` 和 `pgvault`，注册派生密钥，并仅在首次运行时创建管理员。初始化 `age` 是为了满足容器预加载的 AGE 钩子；pgvault 控制文件声明的依赖是 `pgcrypto` 与 `pgsodium`，SQL 还使用 PL/pgSQL。扩展本身是纯 SQL，不新增预加载库。

### 接口与角色

| 接口 | 用途 |
| --- | --- |
| `vault.create_secret`、`vault.update_secret`、`vault.delete_secret` | 供具有写权限的用户修改秘密 |
| `vault.reveal_secret` | 供具有读权限的用户按 UUID 解密 |
| `vault.list_secrets` | 列出元数据，不返回明文 |
| `vault.create_user`、`vault.deactivate_user`、`vault.change_password` | 管理 vault 主体与凭据 |
| `vault.red_alert`、`vault.stand_down` | 禁用或恢复秘密访问 |

管理员可以管理用户，但在此 API 中没有秘密读写权限；其他用户获得读、写或两者兼有的权限。每次 API 调用都验证用户名和口令，这些权限并不是针对单个秘密的所有权规则。配置、用户和密文位于固定的 `vault` 模式，并包含在识别扩展配置表的逻辑备份中。

### 运行边界

凭据会作为 SQL 参数传入，应保护传输与语句日志。PostgreSQL 角色和对象所有权需要单独管理；应用层检查无法限制能够修改对象或权限的数据库超级用户。安全定义者函数的搜索路径包含 `public`，不能允许不可信用户写入。重启、恢复和复制时均须保留根密钥。上游没有明确许可证，也未公布更广泛的 PostgreSQL 兼容矩阵。
