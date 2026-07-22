## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/README.md)
- [SQL 函数定义](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/sql/rpg_base36.sql)
- [Rust 实现](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/src/lib.rs)
- [Cargo 包元数据](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/Cargo.toml)
- [扩展 control 文件](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/rpg_base36.control)

`rpg_base36` 是一个历史性的 Rust 玩具扩展，只有一个函数 `base36_encode(integer)`，把 PostgreSQL 32 位整数转换为大写的 36 进制。零变为 `0`，普通负数保留前导负号。它没有解码器或可选字母表。

### 核心流程

```sql
CREATE EXTENSION rpg_base36;

SELECT base36_encode(0);       -- 0
SELECT base36_encode(35);      -- Z
SELECT base36_encode(123456);  -- 2N9C
SELECT base36_encode(-35);     -- -Z
```

规范扩展名是 `rpg_base36`，但 SQL 加载保护消息仍含有旧名称 `rust_base36`。control 文件和 SQL 扩展版本为 0.0.1，Rust crate 元数据则写为 0.1.0；crate 版本不能证明存在更高版本的数据库扩展发行版。

### 兼容性与输入边界

实现会在转换前调用 Rust 整数绝对值。32 位最小整数 `-2147483648` 没有对应的正 32 位整数，在已复核的旧工具链中可能溢出或 panic；应拒绝该值。

该项目早于 `pgrx`，使用旧式 `rpgffi` 绑定、C 胶水和不安全的 PostgreSQL 函数调用访问。必须针对确切 PostgreSQL 大版本、Rust 编译器和目标架构构建与测试。应把它视为学习示例，而不是受维护的生产依赖。
