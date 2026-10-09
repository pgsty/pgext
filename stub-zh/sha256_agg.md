## 用法

来源：

- [README.md](https://github.com/danolivo/sha256_agg/blob/bc0ead44962c4afa53c2b0446e5d73492f422405/README.md)
- [sha256_agg--1.0.sql](https://github.com/danolivo/sha256_agg/blob/bc0ead44962c4afa53c2b0446e5d73492f422405/sha256_agg--1.0.sql)
- [Makefile](https://github.com/danolivo/sha256_agg/blob/bc0ead44962c4afa53c2b0446e5d73492f422405/Makefile)
- [sha256_agg.control](https://github.com/danolivo/sha256_agg/blob/bc0ead44962c4afa53c2b0446e5d73492f422405/sha256_agg.control)

`sha256_agg` 对按指定顺序排列的整行或选定字段计算单一摘要，用于比较不同副本中的逻辑表内容。

### 核心用法

```sql
CREATE EXTENSION sha256_agg;
CREATE TABLE digest_demo (id integer PRIMARY KEY, value text);
INSERT INTO digest_demo VALUES (1, 'one'), (2, 'two');
SELECT encode(sha256_agg(t ORDER BY id), 'hex') FROM digest_demo AS t;
SELECT sha256_agg_kernel();
```

### 运行边界

必须在聚合内部按唯一键排序。空输入返回 NULL，但所有字段均为 NULL 的一行仍会改变摘要。跨服务器比较要求字节序相同。`sha256_agg_encode` 返回参与摘要计算的行字节，`sha256_agg_kernel` 显示当前实现；`sha256_agg.hardware_acceleration` 可关闭硬件加速而不改变结果。不支持部分聚合或并行聚合。上游面向 PostgreSQL 14+，报告在 18 和 master 上测试；控制文件声明为受信任且可迁移模式。
