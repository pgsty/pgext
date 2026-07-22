## 用法

来源：

- [官方 README](https://github.com/sqids/sqids-postgresql/blob/709167f4c3c97f7af096a94c854e8e5071ddeaf5/README.md)
- [Rust SQL API 实现](https://github.com/sqids/sqids-postgresql/blob/709167f4c3c97f7af096a94c854e8e5071ddeaf5/src/lib.rs)
- [构建与 PostgreSQL 特性矩阵](https://github.com/sqids/sqids-postgresql/blob/709167f4c3c97f7af096a94c854e8e5071ddeaf5/Cargo.toml)

`pg_sqids` 在 PostgreSQL 中把一个或多个非负整数编码为简短、URL 安全的 Sqid，也可将其解码回来。它适合公开标识符和短链接，但只是可逆混淆，不是加密或访问控制。

### 核心流程

```sql
CREATE EXTENSION pg_sqids;

SELECT sqids_encode(1, 2, 3);
-- 86Rf07

SELECT sqids_decode('86Rf07');
-- {1,2,3}

SELECT sqids_encode(10::smallint, 1, 2, 3);
-- 86Rf07xd4z
```

第一个 `smallint` 重载设置最小输出长度。其他重载可接收自定义字母表、`text[]` 屏蔽列表，或字母表、最小长度和屏蔽列表的组合；解码时必须使用相同选项。

```sql
SELECT sqids_encode(
    'k3G7QAe51FCsPW92uEOyq4Bg6Sp8YzVTmnU0liwDdHXLajZrfxNhobJIRcMvKt',
    ARRAY['XRKUdQ'],
    1, 2, 3
);
```

### 输入与标识规则

- 负数会被拒绝。实现会把可变数字参数中的空值映射为零，因此当标识必须互异时，应用输入应拒绝空值。
- 最小长度必须介于 0 和 255 之间。
- 多个不同 Sqid 字符串可能解码为同一数字序列。如果文本必须规范化，应使用相同选项解码、重新编码，再与传入字符串比较。
- 屏蔽列表会改变生成结果；它不是任意解码输入的内容过滤器。

知道算法和选项的人可以看出 Sqid 中的数字及数量。不要编码秘密、授权范围或敏感用户数量，而且仍须在解码后的主键上执行数据库授权。

### 兼容性

0.1.0 版使用 `pgrx` 0.18.1，上游 Cargo 特性覆盖 PostgreSQL 13 到 18，默认构建特性为 PostgreSQL 13。必须安装针对准确服务器大版本构建的二进制。控制文件同时标记扩展可信并要求超级用户安装；实际安装策略取决于打包后的控制文件和 PostgreSQL 版本。
