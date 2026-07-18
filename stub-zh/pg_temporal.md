## 用法

来源：

- [上游 README](https://github.com/Wundero/pg_temporal/blob/48e7a74f38d59d59f1006180bf3a1528c0b5f632/README.md)
- [扩展 control 文件](https://github.com/Wundero/pg_temporal/blob/48e7a74f38d59d59f1006180bf3a1528c0b5f632/pg_temporal.control)
- [Cargo 软件包元数据](https://github.com/Wundero/pg_temporal/blob/48e7a74f38d59d59f1006180bf3a1528c0b5f632/Cargo.toml)
- [Rust 实现](https://github.com/Wundero/pg_temporal/blob/48e7a74f38d59d59f1006180bf3a1528c0b5f632/src/lib.rs)

`pg_temporal` `1.0.0` 版添加由 Rust Jiff 库支持的 `ZonedDateTime` 类型，用于表示带偏移量和命名时区的本地日期时间。该类型支持输入输出、相等、排序和哈希；`zdt_now()` 返回当前分区时间。

### 示例

```sql
CREATE EXTENSION pg_temporal;
SELECT zdt_now(), pg_typeof(zdt_now());
```

当前 API 很小：只有类型和当前时间函数，没有文档化的转换或算术。安装仅限超级用户。解析使用随包 Jiff/tz 数据，非法文本会进入 unwrap 路径；行为可能不同于 PostgreSQL `timestamptz` 及其时区数据库。把该类型用于持久模式前，应验证可接受序列化、比较语义、升级和时区数据变更。
