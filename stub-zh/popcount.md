## 用法

来源：

- [官方 README](https://github.com/eschmar/postgresql-popcount/blob/34fadc73163f22ac7b3f567ebded20d80703a597/readme.md)
- [1.0.0 版扩展 SQL](https://github.com/eschmar/postgresql-popcount/blob/34fadc73163f22ac7b3f567ebded20d80703a597/popcount--1.0.0.sql)
- [PGXN 发行页面](https://pgxn.org/dist/popcount/)

`popcount` 为 PostgreSQL `bit(n)` 值提供多种 C 实现，用于计算置位比特数。它用于比较查找表、Hamming weight、编译器 intrinsic 与汇编策略。PostgreSQL 14 及更高版本已经提供内置 `bit_count`，上游也建议现代服务器直接使用内置函数。

### 核心流程

在 PostgreSQL 14 之前的版本中，安装扩展，并且只有在目标 CPU 与位宽上验证后才选择实现：

```sql
CREATE EXTENSION popcount;

SELECT popcount(B'101101');
SELECT popcount64(B'101101');
SELECT popcountAsm64(B'101101');
```

所有函数都不可变且 strict，接受变长 `bit` 值并返回 `integer`。因此空值输入会直接返回空值，不调用 C 函数。

### 可用实现

- `popcount` 使用 8 位查找表策略。
- `popcount32` 与 `popcount64` 使用 32 位和 64 位 Hamming weight 算法。
- `popcountAsm` 与 `popcountAsm64` 使用面向 32 位和 64 位 intrinsic/指令的路径。
- `popcount256` 使用展开的汇编导向实现。

上游基准在其测试硬件上更偏向 `popcountAsm64`，但这不是可移植性或性能保证。CPU 指令可用性、编译标志、对齐、位宽、PostgreSQL 版本与负载都会影响安全性和速度。

在 PostgreSQL 14 或更高版本上，应优先使用内置操作：

```sql
SELECT bit_count(B'101101');
```

### 运维说明

1.0.0 版可重定位，且没有声明预加载或扩展依赖。由于扩展包含特定架构优化代码，必须在每个目标构建与部署架构上运行安装测试，不能在不兼容 CPU 之间复制二进制。

扩展添加六个全局函数名，因此模式放置与搜索路径行为很重要。选择实现前，应使用代表性的 `bit(n)` 宽度做基准，并比较所有实现的正确性。迁移到内置 `bit_count` 可以消除一个原生扩展依赖，但应测试应用查询的返回类型与计划差异。
