## 用法

来源：

- [项目 README](https://github.com/tvondra/pg_check/blob/db34fe20ddf5c597fe97f634a765360b022d13ab/README.md)
- [扩展 control 文件](https://github.com/tvondra/pg_check/blob/db34fe20ddf5c597fe97f634a765360b022d13ab/pg_check.control)
- [0.1.0 版安装 SQL](https://github.com/tvondra/pg_check/blob/db34fe20ddf5c597fe97f634a765360b022d13ab/sql/pg_check--0.1.0.sql)

`pg_check` 0.1.0 对堆表和 B-tree 索引执行基础一致性检查。它检查页面头、重叠 item、元组属性边界和非法 varlena 长度，还可以交叉比对堆元组与索引项，以发现缺失或多余的索引项。

### 执行保守检查

```sql
CREATE EXTENSION pg_check;

SELECT pg_check_table('public.orders'::regclass, false, false);
SELECT pg_check_index('public.orders_pkey'::regclass);
```

`pg_check_table(regclass,bool,bool,bigint,bigint)` 返回问题数量，并可选择检查索引以及将索引与堆表交叉比对。`pg_check_index(regclass,bigint,bigint)` 可检查全部或指定块范围。详细信息会依据 `client_min_messages` 作为服务器消息输出；`pg_check.debug` 和 `pg_check.bitmap_format` 控制额外的交叉检查诊断。

### 锁与检查范围

不了解默认值时，不要省略两个布尔参数：索引检查和交叉检查默认均为 true。交叉检查会在表上取得 `SHARE ROW EXCLUSIVE`，在整个检查期间阻止数据变更；若其他会话按不同顺序获取表锁和索引锁，还可能发生死锁。非交叉检查使用 `ACCESS SHARE`。

该扩展只报告问题，不会修复问题；索引结构验证也只覆盖 B-tree 的页面和 item 基础检查。已审查 C 源码早于当前 PostgreSQL 版本，且没有现代兼容矩阵；应针对实际服务器构建完成验证，并在对重要数据使用此类底层诊断工具前做好备份。
