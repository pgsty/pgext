## 用法

来源：

- [0.0.1 版本 SQL 实现](https://github.com/pindlebot/pg_gen_uid/blob/b5352b72551c56e9a83fcd2981c2c0c82168759b/pg_gen_uid--0.0.1.sql)
- [扩展 control 文件](https://github.com/pindlebot/pg_gen_uid/blob/b5352b72551c56e9a83fcd2981c2c0c82168759b/pg_gen_uid.control)

`pg_gen_uid` 安装 PL/pgSQL 辅助函数 `public.gen_uid(size)`，按指定长度返回小写 ASCII 字母与数字组成的字符串。它通过 `pgcrypto` 获取随机字节。

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_gen_uid;
SELECT gen_uid(16);
```

已复核仓库中没有 README、测试、发行说明或明确许可证。函数固定安装在 `public`，每个字节会对 36 字符字母表取模，因此存在轻微分布偏差。生成值不能替代数据库唯一性约束；标识符必须唯一时，应建立唯一索引并在冲突时重试。

向不受信调用者开放前应校验请求长度，因为函数会分配随机字节，并为每个输出字符循环一次。
