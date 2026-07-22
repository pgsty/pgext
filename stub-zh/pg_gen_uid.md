## 用法

来源：

- [0.0.1 版本 SQL 实现](https://github.com/pindlebot/pg_gen_uid/blob/b5352b72551c56e9a83fcd2981c2c0c82168759b/pg_gen_uid--0.0.1.sql)
- [扩展 control 文件](https://github.com/pindlebot/pg_gen_uid/blob/b5352b72551c56e9a83fcd2981c2c0c82168759b/pg_gen_uid.control)

`pg_gen_uid` 安装一个 PL/pgSQL 辅助函数 `public.gen_uid(size integer)`，返回小写 ASCII 字母与十进制数字。它依赖 `pgcrypto` 取得随机字节。结果只是随机标签，不是 UUID、有序标识符，也不保证唯一。

### 核心流程

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_gen_uid;

CREATE TABLE invitation (
  code text PRIMARY KEY DEFAULT gen_uid(20),
  created_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO invitation DEFAULT VALUES
RETURNING code;
```

`gen_uid(20)` 返回 20 个字符，字符集为 `abcdefghijklmnopqrstuvwxyz0123456789`。标识符必须唯一时，应保留唯一约束并在碰撞后重试。扩展把函数硬编码在 `public`，因此设置权限与 `search_path` 时必须考虑该模式。

### 随机性与输入边界

每个随机字节都会对 36 取模。由于 256 不能被 36 整除，其中四个字符的出现概率略高；不要把结果用于需要均匀 token 或可量化密码学熵的场景。

向调用者暴露辅助函数前，应验证 `size` 非 NULL、为正数且处于合理上限内。函数一次请求全部随机字节，并为每个输出字符循环一次，而 `gen_random_bytes` 自身也有单次调用大小限制。已复核的仓库没有 README、测试、发行说明或声明许可证，因此应独立验证行为与部署策略。
