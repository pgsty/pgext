## 用法

来源：

- [项目 README](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/README.md)
- [扩展 control 文件](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/pksuid-extension/pksuid.control)
- [PostgreSQL 扩展实现](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/pksuid-extension/src/lib.rs)
- [Pksuid 类型实现](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/pksuid/src/lib.rs)
- [扩展 Cargo 清单](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/pksuid-extension/Cargo.toml)

`pksuid` 0.1.0 为 PostgreSQL 提供带前缀 KSUID 的类型和生成函数。其值把便于应用识别的前缀与按时间可排序、高熵的 KSUID 组合起来，例如生成客户域的 `client_...` 标识符，同时不依赖中心序列。

### 使用带前缀主键

```sql
CREATE EXTENSION pksuid;

CREATE TABLE client (
    id pksuid PRIMARY KEY DEFAULT pksuid_generate('client'),
    name text NOT NULL
);

INSERT INTO client (name) VALUES ('Ada');
SELECT * FROM client ORDER BY id;
```

源码定义的 SQL 函数是 `pksuid_generate(text)`，数据类型是 `pksuid`。README 中的 `CREATE EXTENSION prefixed_ksuid` 已经过时；实际 control 文件名和扩展目录名均为 `pksuid`。

### 比较语义

Rust 类型的相等与排序实现只比较底层 KSUID，不比较前缀。因此，前缀只是显示或业务域标签，不参与数据库相等判断和排序。但哈希实现会包含前缀，这与相等语义不一致。如果不同前缀的值可能包含相同 KSUID，应避免使用哈希索引或依赖哈希的分组。

源码把随机函数 `pksuid_generate(text)` 声明为 `IMMUTABLE` 和 `PARALLEL SAFE`。PostgreSQL 可能对使用常量前缀的不可变函数调用做常量折叠或在计划中复用，因此把它用作键默认值可能产生重复值。依赖它生成主键前，应修正函数易变性或使用正确声明的包装函数，并测试多行插入和预备语句。解析器按 `_` 分隔文本；前缀中应避免下划线，并由应用校验前缀策略。已审查的扩展 crate 版本为 0.1.0，提供 PostgreSQL 11 至 16 的 pgrx 构建 feature；每个目标大版本都应分别构建。
