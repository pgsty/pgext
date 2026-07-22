## 用法

来源：

- [官方 README](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/README.md)
- [官方扩展 SQL](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/dbms_sql--1.0.sql)
- [官方扩展控制文件](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/dbms_sql.control)

`dbms_sql` 实现了面向 PostgreSQL 的 Oracle `DBMS_SQL` 游标 API 子集。它帮助迁移需要解析动态语句、绑定标量或数组值、执行语句、描述结果列并抓取行的 PL/SQL；它并不承诺完整兼容 Oracle。

### 核心流程

打开游标，解析带命名绑定的语句，绑定值，执行并关闭游标：

```sql
CREATE EXTENSION dbms_sql;

DO $$
DECLARE
  c integer;
  inserted bigint;
BEGIN
  c := dbms_sql.open_cursor();
  CALL dbms_sql.parse(c, 'insert into target_table(id, label) values(:id, :label)');
  CALL dbms_sql.bind_variable(c, 'id', 42);
  CALL dbms_sql.bind_variable(c, 'label', 'example'::text);
  inserted := dbms_sql.execute(c);
  CALL dbms_sql.close_cursor(c);
  RAISE NOTICE 'affected rows: %', inserted;
END
$$;
```

处理查询结果时，应调用 `define_column` 或 `define_array`，执行语句，循环调用 `fetch_rows`，通过 `column_value` 取值，并在正常和异常路径中始终关闭游标。

### 主要接口

- `open_cursor`、`is_open` 和 `close_cursor` 管理扩展游标句柄。
- `parse` 准备动态 SQL，`bind_variable` 与 `bind_array` 提供命名参数。
- `execute`、`execute_and_fetch`、`fetch_rows` 和 `last_row_count` 驱动执行与抓取。
- `define_column`、`define_array` 和 `column_value` 描述目标缓冲区并复制已抓取值。
- `describe_columns` 与 `describe_columns_f` 返回 `dbms_sql.desc_rec` 元数据，其中 `col_type` 数字是 PostgreSQL 类型 OID，不是 Oracle 类型码。

### 迁移边界

- 应根据已实现签名以及 PostgreSQL 事务和错误语义审计每个迁移调用。不支持的 Oracle 重载、类型转换、权限行为和游标生命周期必须明确改写。
- 安装 SQL 会创建无模式限定的 `varchar2` 域，作为临时兼容对象。上游要求与 Orafce 一起使用时从安装脚本中删除该语句，因为 Orafce 会提供自己的 `varchar2`。
- 游标句柄是后端本地执行状态，不能跨会话、连接池签出、故障转移或事务重试边界。长连接中应确定性关闭游标，避免状态泄漏。
- 动态 SQL 仍需遵守通常的注入防护规则。数据值应使用绑定参数，无法绑定的标识符或 SQL 片段必须验证，并以最小权限调用者运行。
- 该仓库已经归档。迁移或升级前，应固定当前审阅提交，并针对准确的服务器大版本测试其直接使用 PostgreSQL 内部接口的行为。
