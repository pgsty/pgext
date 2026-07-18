## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/README.md)
- [SQL 函数定义](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/sql/rpg_base36.sql)
- [Rust 实现](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/src/lib.rs)
- [扩展 control 文件](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/rpg_base36.control)

`rpg_base36` 是将 PostgreSQL 整数转换为大写 Base36 的 Rust 演示扩展。负数输入会保留前导负号。

```sql
CREATE EXTENSION rpg_base36;

SELECT base36_encode(123456);
SELECT base36_encode(-35);
```

### 注意事项

扩展 control 版本是 0.0.1，而 Cargo 包报告 0.1.0；没有发行版证明存在更新的扩展版本。该项目早于 `pgrx`，使用旧式 `rpgffi` binding 与 C glue，因此应视为历史演示，而不是仍在维护的生产包。其 SQL 文件还残留 `CREATE EXTENSION rust_base36` 守卫提示，但实际扩展名是 `rpg_base36`。
