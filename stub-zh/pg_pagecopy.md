## 用法

来源：

- [上游来源 1](https://github.com/pgtoolz/pg_pagecopy/blob/203a9361d598a7433c086c621f965a2d909e5fb4/README.md)
- [上游来源 2](https://github.com/pgtoolz/pg_pagecopy/blob/203a9361d598a7433c086c621f965a2d909e5fb4/pg_pagecopy.control)

`pg_pagecopy` 0.0.1 在兼容的 PostgreSQL 实例间传输表页面，也可在本地克隆表。

### 核心流程

`pagecopy.pull` 连接源数据库并传输指定关系，`pagecopy.clone` 在同一个数据库内执行克隆，`pagecopy.check` 报告适用性。目标不存在时可从源模式创建，已有目标则按传输约定被覆盖。

```sql
CREATE EXTENSION pg_pagecopy;
CREATE TABLE pagecopy_source (id integer, payload text);
INSERT INTO pagecopy_source VALUES (1, 'example');
SELECT * FROM pagecopy.check('public.pagecopy_source');
SELECT pagecopy.clone('public.pagecopy_source', 'public.pagecopy_copy');
VACUUM (ANALYZE) public.pagecopy_copy;
```

### 运行边界

两端均需安装扩展并具备超级用户权限。上游支持 PostgreSQL 13–18 的同主版本及向上迁移，并执行兼容性检查，拒绝降级。源端冻结及 ACCESS EXCLUSIVE 锁可能阻塞应用。仅支持特定关系和索引形态。需保护连接秘密，保留目标的可恢复副本，并在传输后执行清理与统计分析。
