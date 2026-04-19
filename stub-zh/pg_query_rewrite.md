## 用法

来源：[README](https://github.com/pierreforstmann/pg_query_rewrite/blob/master/README.md)

`pg_query_rewrite` 会把一个完全匹配的源 SQL 语句重写为预定义的目标语句，规则存储在共享内存中。

### 所需设置

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_query_rewrite'
pg_query_rewrite.max_rules = 10

CREATE EXTENSION pg_query_rewrite;
```

### 规则管理

```sql
SELECT pgqr_add_rule('select 10;', 'select 11;');
SELECT pgqr_rules();
SELECT pgqr_remove_rule('select 10;');
SELECT pgqr_truncate();
```

- `pgqr_add_rule(source, target)`：添加一条重写规则。
- `pgqr_remove_rule(source)`：删除一条规则。
- `pgqr_truncate()`：删除全部规则。
- `pgqr_rules()`：查看当前共享内存中的规则和重写计数。

### 注意事项

- 匹配是精确的：大小写、空白和分号都必须逐字符一致。
- 不支持参数化 SQL。
- 最大语句长度硬编码为 32 KB。
- 规则不会持久化；重启后会消失，除非你重新应用。
