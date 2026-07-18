## 用法

来源：

- [1.0 版本 SQL 对象](https://github.com/amutu/fnvhash/blob/f536e411fa8b23cca1fd664e017c80ee6df42448/fnvhash--1.0.sql)
- [包装实现](https://github.com/amutu/fnvhash/blob/f536e411fa8b23cca1fd664e017c80ee6df42448/fnvhash.c)
- [扩展 control 文件](https://github.com/amutu/fnvhash/blob/f536e411fa8b23cca1fd664e017c80ee6df42448/fnvhash.control)

`fnvhash` 为 `varchar`、`text` 与 `bytea` 暴露 Fowler–Noll–Vo 哈希变体。32 位函数 `fnv032`、`fnv132` 和 `fnv1a32` 返回 `bigint`；64 位对应函数 `fnv064`、`fnv164` 和 `fnv1a64` 返回文本表示。

```sql
CREATE EXTENSION fnvhash;
SELECT fnv1a32('example'::text),
       fnv1a64('example'::text),
       fnv1a32(E'\\x0001'::bytea);
```

FNV 是快速的非密码学哈希，碰撞是正常现象，32 位尤其如此。绝不能把结果用于密码、认证、签名、防篡改，也不能作为唯一性保证。身份重要时应保留原值并使用数据库约束。

除源码外，上游 README 没有给出兼容性或字节序契约。若用于分区、跨系统关联或持久协议字段，应固定实现，并在不同架构、编码、PostgreSQL 版本与客户端表示之间验证精确输出。修改算法或返回格式时必须协调迁移数据。
