## 用法

来源：

- [官方 README](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/README.rst)
- [扩展控制文件](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/fast_guid.control)
- [0.1 版扩展 SQL](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/fast_guid--0.1.sql)
- [C 实现](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/fast_guid.c)

`fast_guid` 0.1 返回由时间、分片和序列派生的 `bigint`。它不返回 PostgreSQL `uuid`，也不会自行分配序列号；调用方必须负责提供不会碰撞的 `smallint` 输入。

### 核心流程

使用序列提供低位值：

```sql
CREATE EXTENSION fast_guid;
CREATE SEQUENCE global_id_sequence;

SELECT fast_guid(
  7::smallint,
  (nextval('global_id_sequence') % 1024)::smallint
);
```

函数读取从 `2011-01-01 00:00:00 UTC` 开始的墙钟毫秒数，将 `shard_id` 对 8192 取模、`seq_id` 对 1024 取模，再组合为有符号 `bigint`。

### 接口

`fast_guid(shard_id smallint, seq_id smallint) RETURNS bigint` 被声明为 `STRICT`，易变性则采用 PostgreSQL 默认的 `VOLATILE`。因此 NULL 输入会直接得到 NULL，而不会调用 C 函数。

上游注释描述的预期布局是 41 位时间、13 位分片和 10 位序列。实际源码将分片值左移 13 位而不是 10 位，因此使用方不能按注释中的连续 41/13/10 布局解码 ID，除非先验证所安装实现。

### 碰撞与排序注意事项

同一毫秒内使用相同取模后分片值和序列值的两次调用可能返回相同结果。每分片每毫秒超过 1024 个 ID、序列重置、时钟回拨、重复分片分配或负 `smallint` 输入，都会破坏唯一性或顺序。该函数没有跨节点协调，也不持久保存时钟状态。

即使保留此函数，也应使用主键唯一约束；不要把输出视为加密随机、全局协调或不透明的标识符。新系统应优先采用仍在维护、明确规定并测试位布局、溢出行为和分布式节点分配的生成器。
