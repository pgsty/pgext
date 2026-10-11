## 用法

来源：

- [examples/pgaudit_zig/README.md](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/README.md)
- [examples/pgaudit_zig/build.zig](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/build.zig)
- [examples/pgaudit_zig/src/main.zig](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/src/main.zig)
- [examples/pgaudit_zig/extension/pgaudit_zig--0.1.sql](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/extension/pgaudit_zig--0.1.sql)
- [examples/pgaudit_zig/extension/pgaudit_zig.control](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/extension/pgaudit_zig.control)

`pgaudit_zig` 0.1 演示如何用 `pgzx` 编写执行器审计钩子。上游明确将它定位为不完整的教学示例，并建议生产审计使用 `pgaudit`。

### 核心用法

```sql
LOAD 'pg_audit_zig';
CREATE EXTENSION pgaudit_zig;
CREATE TABLE audit_demo (id integer);
SELECT * FROM audit_demo;
```

### 运行边界

SQL 扩展名是 `pgaudit_zig`，构建出的库名则是 `pg_audit_zig`。必须在会话中显式加载该库才能注册钩子；单独执行 `CREATE EXTENSION` 只会运行没有功能对象的 SQL 脚本，不会启用日志。这些操作需要相应的超级用户权限。钩子记录命令类型、引用的关系及语句文本。超级用户可设置布尔参数 `pgaudit_zig.log_statement`，默认值为 true，用于控制是否包含语句文本；日志中的文本可能含敏感值。审阅代码将记录送入 PostgreSQL 日志。该示例未承诺完整审计覆盖、防篡改能力或明确的兼容版本范围。
