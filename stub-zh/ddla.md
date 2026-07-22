## 用法

来源：

- [已审查提交上的 README](https://github.com/asotolongo/ddla/blob/385eeaafad77ae401a64a1f60181869ca2d1b973/README)
- [扩展 control 文件](https://github.com/asotolongo/ddla/blob/385eeaafad77ae401a64a1f60181869ca2d1b973/ddla.control)
- [0.1 版安装 SQL](https://github.com/asotolongo/ddla/blob/385eeaafad77ae401a64a1f60181869ca2d1b973/ddla--0.1.sql)

`ddla` 0.1 是一个纯 SQL 的 DDL 审计扩展。安装会创建 `ddla` schema、`ddla.ddl_logs` 表以及两个作用于整个数据库的事件触发器。触发器记录已完成的 DDL 命令和被删除对象，并保存命令标签、对象标识与类型、时间戳、当前用户、客户端地址和完整的当前查询文本。

### 查看审计日志

```sql
CREATE EXTENSION ddla;

CREATE TABLE ddl_audit_example (id bigint PRIMARY KEY);

SELECT * FROM ddla.get_ddl_cmd_by_cmd('CREATE TABLE');
SELECT * FROM ddla.get_ddl_cmd_stats();
```

查询函数还可以按对象类型、时间范围或用户筛选。`ddla.reset_logs()` 会清空审计表，`ddla.reset_id_seq_logs()` 还会重置其序列。

### 运维注意事项

上游要求 PostgreSQL 9.6 或更高版本。control 文件要求由超级用户安装，事件触发器会处理数据库中的全部 DDL。应为 `ddla.ddl_logs` 规划保留策略；由于 `query` 可能含敏感文本，还应限制访问，并在 DDL 密集型系统上评估开销。

它并不是防篡改审计链：安装 SQL 向 `PUBLIC` 授予 `ddla` schema 的使用权、`ddla.ddl_logs` 的插入权以及其序列的全部权限，因此普通角色可以写入伪造记录。尽管 control 文件声明扩展可重定位，安装 SQL 仍将 `ddla` schema 写死。把日志作为证据前，应撤销不必要的授权，并保护或外送日志。上游 README 明确指出 0.1 版存在已知缺陷；将其作为审计控制前，应在目标 PostgreSQL 版本上验证实际行为。
