## 用法

来源：

- [已归档项目 README](https://github.com/supabase/supa_audit/blob/4f527f062e85ec0549c45009c3860d50ddf7f282/README.md)
- [扩展 control 文件](https://github.com/supabase/supa_audit/blob/4f527f062e85ec0549c45009c3860d50ddf7f282/supa_audit.control)
- [0.3.1 版 SQL 实现](https://github.com/supabase/supa_audit/blob/4f527f062e85ec0549c45009c3860d50ddf7f282/supa_audit--0.3.1.sql)
- [回归测试](https://github.com/supabase/supa_audit/tree/4f527f062e85ec0549c45009c3860d50ddf7f282/test/sql)

`supa_audit` 0.3.1 是一个已归档、适用于 PostgreSQL 12 及以上版本的纯 SQL 触发器审计扩展。它把 insert、update、delete 和 truncate 事件存入 `audit.record_version`。对行事件，它会记录当前与先前行的 JSONB，并由表 OID 和主键值派生 UUID 记录标识符。

### 为表启用跟踪

```sql
CREATE EXTENSION supa_audit CASCADE;

CREATE TABLE public.account (
  id bigint PRIMARY KEY,
  display_name text NOT NULL
);

SELECT audit.enable_tracking('public.account'::regclass);

SELECT id, record_id, old_record_id, op, ts, record, old_record
FROM audit.record_version
WHERE table_oid = 'public.account'::regclass::oid
ORDER BY id;
```

必要时，`CASCADE` 会安装所需的 `uuid-ossp` 扩展。跟踪要求表有主键。查询应使用带索引的 `table_oid`，并单独定义保留和归档任务，因为审计表会随写入增长。

### 审计数据具有事务性且敏感

触发器行与源事务一起提交和回滚。它不是独立、防篡改的取证日志：表 owner 或特权角色可以禁用触发器、修改历史或绕过该路径。truncate 事件只记录一个标记，不保存每个被删除的行。重要事件应导出到单独控制的存储。

每个被跟踪列，包括凭据、token、个人数据和大型值，都可能被复制到当前及旧 JSONB 中。应限制 `audit` 模式访问、加密并最小化保留数据，并把它纳入隐私删除与备份策略。上游警告审计会降低写吞吐，不建议对峰值每秒超过 3,000 次操作的表启用。表 OID 参与记录标识符计算且不是可移植身份，因此必须验证转储恢复后的连续性。项目已归档；继续部署前需要审查其 security-definer 函数以及未固定 UUID 函数搜索路径的问题。
