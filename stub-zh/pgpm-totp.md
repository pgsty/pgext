## 用法

来源：

- [packages/totp/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/totp/README.md)
- [packages/totp/sql/pgpm-totp--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/totp/sql/pgpm-totp--0.47.0.sql)
- [packages/totp/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/totp/Makefile)
- [packages/totp/pgpm-totp.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/totp/pgpm-totp.control)

`pgpm-totp` 通过 pgcrypto HMAC 与 PGPM Base32 辅助函数实现 `totp.generate` 和 `totp.verify`。默认周期为 30 秒、六位数字和 SHA-1；SQL 还接受显式时间、摘要算法、编码及时间步偏移。

### 核心用法

```sql
CREATE EXTENSION "pgpm-totp";
WITH demo AS (SELECT totp.random_base32(20) AS secret)
SELECT totp.verify(secret, totp.generate(secret, 30, 6), 30, 6)
FROM demo;
```

### 运行边界

`totp.random_base32` 使用随机字节生成密钥文本，`totp.url` 构造包含密钥的登记 URI，二者都应保密。校验只检查一个选定时间步，不记录重放、不限制尝试速率，也不会搜索可接受的时间窗口，这些策略由应用负责。URI 编码器拒绝需要三个或更多 UTF-8 字节的字符。 依赖为 `pgcrypto`, `plpgsql`, `pgpm-base32`, `pgpm-verify`。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
