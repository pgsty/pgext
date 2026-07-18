## 用法

来源：

- [已复核 commit 的 pg_xid README](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/README.md)
- [已复核 commit 的 pg_xid 1.0 安装 SQL](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/pg_xid--1.0.sql)
- [已复核 commit 的 pg_xid C 入口](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/pg_xid.c)

`pg_xid` 1.0 生成与 Mongo ObjectID 兼容的 12 字节标识符。`xid` 以 `bytea` 返回原始值；`xid_encoded` 返回其 20 字符编码 `text` 形式。布局组合秒级时间戳、机器标识、进程 ID 和进程内计数器。

```sql
CREATE EXTENSION pg_xid;

SELECT encode(xid(), 'hex');
SELECT xid_encoded();
```

时间戳前缀让这些标识符无需中央生成器即可大致按创建时间排序。

### 注意事项

- 这些标识符不是秘密，也不是密码学随机令牌；其布局会暴露创建时间及由进程、机器派生的组成部分。
- 按上游说明，24 位计数器在每个主机/进程每秒回绕前可生成 16,777,216 个值；仍需按实际负载测试生成器边界。
- 上游没有发布系列或当前 PostgreSQL 兼容矩阵。应在每个目标大版本和架构上构建并测试 C 扩展的并发行为。
- 原始形式和编码形式应保持一致存储。将其用作分布式排序键之前，要验证文本排序、二进制排序及应用序列化。
