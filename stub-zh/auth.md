## 用法

来源：

- [docs/auth.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/auth.rst)
- [auth/auth--0.1.sql](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/auth/auth--0.1.sql)
- [docs/install.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/install.rst)
- [auth/auth.control](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/auth/auth.control)

`auth` 是 Oliphant 的权限管理扩展，可检查、复制或迁移授权，并在替换表前保存表权限。

### 核心用法

```sql
CREATE EXTENSION auth;
SELECT * FROM role_auths(current_user::name);
```

### 运行边界

这些是最初面向 PostgreSQL 9.x 编写文档的纯 SQL/PLpgSQL 工具，尚未验证当前主版本兼容性。安装沿用 control 默认权限，需要管理员；不含共享库，也不需要预加载。 `copy_role_auths` 为目标角色补齐缺少的授权，`move_role_auths` 还会撤销源角色的相应授权。调用者仍需要授权和撤权能力，通常为超级用户。对象所有权不会随之转移。`store_table_auths` 写入 `stored_table_auths`，`restore_table_auths` 重放保存的授权；恢复前应检查这些记录，尤其是在替换表或手动修改之后。
