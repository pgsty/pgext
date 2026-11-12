## 用法

来源：

- [I-Rex 扩展指南](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/README.md)
- [控制文件](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/blmfl.control)
- [0.0.1 版 SQL](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/blmfl--0.0.1.sql)
- [用法测试](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/sql/blmfl_test.sql)
- [C 实现](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/blmfl.c)
- [PostgreSQL 16 环境](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/docker/postgres/Dockerfile)

`blmfl` 0.0.1 是 I-Rex SQL 调试研究制品附带的实验性布隆过滤器与范围摘要扩展。它提供聚合与成员测试，不会自动优化普通 PostgreSQL 查询。

### 启用并创建摘要

安装原生扩展需要管理权限。所提供的环境使用 PostgreSQL 16，该制品没有确立其他大版本的兼容性。无需共享预加载。

```sql
CREATE EXTENSION blmfl;

SET blmfl.bloomfilter_bitsize = 256;
SET blmfl.estimated_count = 10;
SET blmfl.num_hashes = 15;

CREATE TEMP TABLE example_values (value integer);
INSERT INTO example_values VALUES (5), (24), (90);

CREATE TEMP TABLE example_summary AS
SELECT blmfl_any(value) AS summary FROM example_values;

SELECT blmfl_test_any(summary, 24),
       blmfl_fpr(summary)
FROM example_summary;
```

应显式设置正数大小与计数：三个 GUC 的默认值均为零，而源码会使用它们进行除法和取模。这些会话设置只决定随后构造的摘要，不会修改已有值。

### SQL 接口

- `BLMFL_RESULT` 保存过滤器字节、大小、估计数量、哈希次数和累计项目数。
- `blmfl_any` 与 `blmfl_test_any` 接受带类型的参数；`blmfl` 与 `blmfl_test` 提供 bytea 变体。构造与探测时应保持类型一致。
- `blmfl_optimal_k`、`blmfl_fpr` 和 `blmfl_approx_unique_count` 提供容量选择与估算辅助函数。
- `blmfl_merge` 与 `blmfl_agg` 合并摘要。实现检查位长度是否相同，但没有验证所有兼容性条件。
- `sarg_agg` 构造或合并上下界摘要；`is_iid_*`、`min_iid` 和 `max_iid` 辅助处理 I-Rex 的内部标识符。

### 研究用途边界

源码审查发现，带类型参数的哈希循环只保留最后一个参数的哈希，且忽略每次哈希的种子，不能将其视为经过验证的多列布隆过滤器。合并缓冲区分配后未显式清零，空输入的最终处理会报错。因此，估算与合并摘要均需独立检查正确性。

应在隔离的研究环境中使用该制品。未经精确验证，不要依靠这些摘要排除业务数据行。所核对修订未发现许可证授权、受支持的升级链或生产兼容性保证。
