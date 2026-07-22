## 用法

来源：

- [官方 README](https://github.com/petere/pgvihash/blob/70e637321ab07746ecaebf6c47385b06c84c050c/README)
- [官方 SQL 定义](https://github.com/petere/pgvihash/blob/70e637321ab07746ecaebf6c47385b06c84c050c/vihash.sql)
- [官方 C 实现](https://github.com/petere/pgvihash/blob/70e637321ab07746ecaebf6c47385b06c84c050c/vihash.c)
- [官方 PGXN 发布页](https://pgxn.org/dist/pgvihash/)

`vihash` 提供一个稳定的 32 位文本哈希 `vihashtext(text)`，实现复制自 PostgreSQL 8.3 的 `hashtext`。它的用途是在迁移期间多个 PostgreSQL 大版本并存时，保持分片/分区路由决策不变。

### 核心流程

在所有参与数据库安装扩展，并在修改路由代码前验证结果完全一致：

```sql
CREATE EXTENSION vihash;

SELECT vihashtext('tenant-42');

SELECT key, vihashtext(key) AS routing_hash
FROM (VALUES ('tenant-1'), ('tenant-2')) AS v(key);
```

对于 PL/Proxy 时代的代码，只有在规划好的重分布/兼容迁移中，才应把 `RUN ON hashtext(arg)` 替换为 `RUN ON vihashtext(arg)`。

### 兼容契约

`vihashtext(text) RETURNS integer` 对数据库编码后的原始字节做哈希，设计目标是跨版本保持不变。输出与 PostgreSQL 8.0–8.3 内置 `hashtext` 相同，但不匹配 PostgreSQL 8.4 开始改变后的内置实现。

### 注意事项

如果数据已经使用 PostgreSQL 8.4 或更高版本的 `hashtext` 分布，改用 `vihashtext` 会改变路由并需要重分布。切换前应在新旧节点用固定语料集测试，保持数据库编码一致，并在映射分片号时统一处理有符号 32 位结果。

这是迁移辅助工具，不是密码学或抗碰撞哈希。上游建议新设计采用提供更多稳定算法的库，而且仓库已归档。在没有碰撞处理时，绝不能把结果用于认证、完整性、唯一性或对抗性分区键。
