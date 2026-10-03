## 用法

来源：

- [3.2.1 README](https://github.com/devrimgunduz/passwordcheck_cracklib/blob/3.2.1/README.md)
- [3.2.1 password hook](https://github.com/devrimgunduz/passwordcheck_cracklib/blob/3.2.1/passwordcheck_cracklib.c)
- [PostgreSQL passwordcheck manual](https://www.postgresql.org/docs/18/passwordcheck.html)

`passwordcheck_cracklib` 使用 CrackLib 检查通过 `CREATE ROLE` 和 `ALTER ROLE` 设置的密码。它是服务器钩子库，不创建 SQL 扩展对象。

### 启用钩子

将它加入现有预加载列表，然后重启 PostgreSQL：

```ini
shared_preload_libraries = '$libdir/passwordcheck_cracklib'
```

不要执行 `CREATE EXTENSION passwordcheck_cracklib`。PostgreSQL 操作系统账号必须能够使用 CrackLib 库和字典。

### 密码检查

```sql
CREATE ROLE app_user LOGIN PASSWORD 'password123';
```

弱明文密码会触发错误。钩子检查长度、密码与用户名的关系、字符组成以及 CrackLib 字典。通过这些检查并不保证密码能够抵抗所有攻击。

### 安全边界

字典检查要求在修改密码时获得明文密码。如果客户端提交已经计算好的密码散列，模块无法完成完整的强度检查；此时仅能检查密码是否等于用户名。应约束密码修改入口，并保护传输明文密码的连接。

模块不会追溯扫描已有密码。钩子还会调用先前安装的密码检查钩子；同时加载其他凭证策略库前，应检查它们的交互。
