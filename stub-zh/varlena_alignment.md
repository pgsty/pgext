## 用法

来源：

- [官方缺陷分析](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/README.md)
- [复现测试](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/src/lib.rs)
- [固定的 pgrx 依赖与 PostgreSQL 特性](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/Cargo.toml)

`varlena_alignment` 版本 `0.0.0` 是用于复现历史 pgrx 对齐缺陷的开发者项目。它演示了在 `set_varsize_4b` 中把按一字节对齐的 PostgreSQL `varlena` 指针转换为按四字节对齐的 `varattrib_4b` 可能触发未定义行为。

### 预期流程

仓库只定义了 `#[pg_test]` 测试。它不导出应用 SQL 函数、类型、运算符或配置，因此通过 `CREATE EXTENSION` 安装后没有面向用户的工作流。其用途是针对固定依赖运行 pgrx 测试框架：

```sh
cargo pgrx test
```

一个测试在故意偏移的分配内存上调用 `set_varsize_1b`，预期成功。另一个测试在同一个按一字节对齐的 `varlena` 指针上调用 `set_varsize_4b`，复现 Rust 调试指针检查可见的更严格对齐失败。

### 开发者边界

该项目固定使用 `pgrx` 版本 `0.11.1`，并声明从 `pg11` 到 `pg16` 的 PostgreSQL 特性集。它是关于该实现上下文的证据，不是通用运行时诊断，也不是修复。不要在生产环境安装，也不要把 release 构建通过视为不安全转换有效的证明；应直接评估相关 pgrx 版本与内存对齐契约。
