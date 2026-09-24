## 用法

来源：

- [AlohaDB README](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [Control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vault/alohadb_vault.control)
- [SQL](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vault/alohadb_vault--1.0.sql)
- [Source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vault/alohadb_vault.c)

`alohadb_vault` 在基于 PostgreSQL 18 的 AlohaDB 发行版中保存按名称索引的加密秘密。扩展版本 1.0 使用 pgcrypto 的 PGP 对称加密、基于角色的表策略和审计表。文档对应的平台是 AlohaDB，不能据此认定已有独立的社区 PostgreSQL 软件包。

### 核心用法

由超级用户创建 `pgcrypto` 和本扩展。使用前，通过受保护的服务器或会话配置设置仅超级用户可修改的 `alohadb.vault_passphrase`；空口令会报错。模块加载时注册该设置，没有要求共享预加载。

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION alohadb_vault;
SET alohadb.vault_passphrase = 'example-only-change-me';
SELECT alohadb_vault_store('example_key', 'example_value', 'demonstration');
SELECT alohadb_vault_fetch('example_key');
SELECT * FROM alohadb_vault_list();
SELECT alohadb_vault_delete('example_key');
```

本例使用可丢弃的演示值。实际口令应避开 SQL 历史和日志，并与加密数据分别备份。安装时创建管理角色的代码块需要 PL/pgSQL。

### 对象与权限

| 对象 | 用途 |
| --- | --- |
| `alohadb_vault_store(text,text,text)` | 插入或替换命名秘密，描述可省略 |
| `alohadb_vault_fetch(text)` | 按键返回解密后的明文 |
| `alohadb_vault_delete(text)` | 删除秘密 |
| `alohadb_vault_list()` | 返回键、描述与时间戳 |
| `alohadb_vault_rotate_key(text,text)` | 用旧、新口令重新加密数据，返回行数 |
| `alohadb_vault_secrets`、`alohadb_vault_audit` | 密文与操作记录 |
| `alohadb_vault_admin` | 被授予底层表权限、纳入 RLS 策略的管理角色 |

### 维护边界

C 函数以调用者权限执行 SQL，只应向指定操作人员授予管理角色成员资格。RLS 与审计记录不能隔离超级用户或对象所有者。实现使用未限定模式的表名和 pgcrypto 函数名，因此必须控制调用者的搜索路径与模式写权限。

密钥轮换重写密文，但不会修改 `alohadb.vault_passphrase`；应将配置密钥切换与成功提交的轮换事务协调，并保留可恢复备份。扩展不可迁移模式，安装时还会创建集群级角色。不能假定删除数据库中的扩展会一并删除该角色，也不能将其当作密钥管理流程。
