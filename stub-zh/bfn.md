## 用法

来源：

- [已复核 commit 的 bfn README](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/readme.md)
- [已复核 commit 的 bfn Rust API](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/src/lib.rs)
- [已复核 commit 的 bfn.control](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/bfn.control)

`bfn` 是一组以 Rust/`pgrx` 编写的实用函数。2.0.1 版提供 UUIDv7 生成与时间戳提取、日期范围与月边界、字符串规范化与校验、JSONB 地址构造，以及哈希或随机值操作等辅助能力。

上游安装方式将扩展放入名为 `bfn` 的模式。代表性函数包括 `bfn_version`、`new_uuid`、`uuid_to_ts`、`date_range`、`first_day_of_month` 和 `san_trim`。

### 安装与检查

```sql
CREATE SCHEMA bfn;
CREATE EXTENSION bfn SCHEMA bfn;

SELECT bfn.bfn_version();
SELECT bfn.new_uuid();
SELECT bfn.uuid_to_ts(bfn.new_uuid());
SELECT bfn.date_range(DATE '2026-01-01', DATE '2026-01-03');
```

### 注意事项

- 已复核 control 文件设置了 `superuser = true`，因此安装需要具有相应权限的角色。
- 已复核 Cargo 清单面向 PostgreSQL 16、17 和 18，并标识版本 2.0.1，与 catalog 版本一致。
- 这是较宽泛的实用函数集合，而非单一用途扩展。将某个辅助函数写入应用 SQL 前，应查阅固定版本的 Rust API，确认其确切的空值处理和参数类型。
