## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/theory/pg-hash64/blob/2f969d14998f82c98322bb6f53f788b5be5d8b18/README.md)
- [扩展 control 文件](https://github.com/theory/pg-hash64/blob/2f969d14998f82c98322bb6f53f788b5be5d8b18/hash64.control)
- [C 实现](https://github.com/theory/pg-hash64/blob/2f969d14998f82c98322bb6f53f788b5be5d8b18/hash64.c)

`hash64` 提供返回 `bigint` 的稳定 64 位哈希函数 `hash64(text)`。与 PostgreSQL 内部哈希函数不同，上游承诺结果不会随 PostgreSQL 大版本变化，因此适合在混合版本部署中进行确定性分片或分区分配。

```sql
CREATE EXTENSION hash64;
SELECT hash64('customer-42');
SELECT mod(hash64('customer-42') & 9223372036854775807, 32) AS shard_id;
```

哈希既不唯一，也不是密码学认证器。必须保留原始键用于等值检查、容许碰撞，并且不能把结果用于密码存储、签名或对抗性完整性检查。

0.1.0 已较旧，且没有当前 PostgreSQL 兼容矩阵。将其作为跨集群路由契约前，应在所有目标架构上构建，并用固定语料验证预期输出；否则整数规范化或平台行为变化会导致数据移动到其他分片。
