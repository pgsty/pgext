## 用法

来源：

- [已复核 commit 的 README](https://github.com/drforr/plsix/blob/813fd801fa1d11ead0d834f7a4d5144e57d86952/README.md)
- [扩展 control 文件](https://github.com/drforr/plsix/blob/813fd801fa1d11ead0d834f7a4d5144e57d86952/plsix.control)
- [语言安装 SQL](https://github.com/drforr/plsix/blob/813fd801fa1d11ead0d834f7a4d5144e57d86952/plsix--2.sql)
- [语言 handler 实现](https://github.com/drforr/plsix/blob/813fd801fa1d11ead0d834f7a4d5144e57d86952/plsix.c)

`plsix` 是一个面向 Perl 6（现称 Raku）的非受信 PostgreSQL 过程语言 handler。它派生自 PL/sh：函数体会被复制到一个临时可执行脚本，参数被传给子进程，标准输出则被转换为声明的 PostgreSQL 返回类型。

### 存储函数

首行用于选择已安装的解释器。这个历史项目要求存在 `perl6` 可执行文件。

```sql
CREATE EXTENSION plsix;

CREATE FUNCTION plsix_echo(integer)
RETURNS text
LANGUAGE plsix
AS $$
#!/usr/bin/env perl6
say "argument is " ~ @ARGS[0];
$$;

SELECT plsix_echo(7);
```

标准输出为空时返回 SQL `NULL`；标准错误和非零子进程状态会被报告为错误。handler 还支持 `DO`、trigger 与 event trigger 入口。Trigger 上下文通过继承的 `PLSH_TG_*` 环境变量导出。

### 注意事项

- 这是一种非受信语言，会以 PostgreSQL 服务器账号的权限执行任意操作系统代码。只有超级用户可以用它创建函数；必须审计每个函数体和解释器路径。
- handler 没有 SPI 集成。需要访问数据库的脚本必须另建客户端连接，递归 `psql` 调用具有独立的事务行为。
- Trigger 脚本能够接收上下文，却无法通过此接口返回修改后的行。需要改写行的 trigger 语义时不要使用 `plsix`。
- README 保留了从 PL/sh 复制而来的名称和安装示例，而实际扩展及语言名称是 `plsix`。应以版本 `2` 的安装 SQL 作为 API 权威来源。
- 源码最后更新于 2019 年，使用旧的 `perl6` 可执行文件名，也没有当前 PostgreSQL 或 Raku 兼容矩阵。生产代码应优先选择仍受维护的过程语言。
