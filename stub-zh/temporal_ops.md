## 用法

来源：

- [README.md](https://github.com/pjungwir/temporal_ops/blob/516fe77bf732e1858329dc583a7cfbea43a4e81b/README.md)
- [temporal_ops--1.0.0.sql](https://github.com/pjungwir/temporal_ops/blob/516fe77bf732e1858329dc583a7cfbea43a4e81b/temporal_ops--1.0.0.sql)
- [temporal_ops.c](https://github.com/pjungwir/temporal_ops/blob/516fe77bf732e1858329dc583a7cfbea43a4e81b/temporal_ops.c)
- [temporal_ops.control](https://github.com/pjungwir/temporal_ops/blob/516fe77bf732e1858329dc583a7cfbea43a4e81b/temporal_ops.control)

`temporal_ops` 提供考虑时间区间的连接：半连接保留左侧记录被覆盖的时间段，反连接保留未覆盖部分，外连接还表示未匹配的时间。它将 C 查询生成器与可从 SQL 调用的包装函数结合起来。

### 核心用法

```sql
CREATE EXTENSION temporal_ops;
CREATE TABLE temporal_left (id integer, valid_at daterange);
CREATE TABLE temporal_right (id integer, valid_at daterange);
INSERT INTO temporal_left VALUES (1, '[2026-01-01,2026-02-01)');
INSERT INTO temporal_right VALUES (1, '[2026-01-10,2026-01-20)');
SELECT (j.item).id, j.valid_at
FROM temporal_semijoin('temporal_left', 'id', 'valid_at',
                       'temporal_right', 'id', 'valid_at')
AS j(item temporal_left, valid_at daterange);
```

### 运行边界

当前源码使用 PostgreSQL 19 的 `SupportRequestInlineInFrom` 规划器接口，不声明兼容更早主版本。C 函数和防泄漏函数需要管理员安装。`temporal_semijoin`、`temporal_antijoin` 和 `temporal_outer_join` 接收 `regclass` 类型的关系名、对应的键列与区间列名；数组重载支持复合键。返回记录需要调用者显式声明行结构，`_sql` 变体可查看生成的 SQL。关系必须存在，调用者需要读取权限。并集、差集和聚合仍属于计划中的能力。未找到项目许可证声明。
