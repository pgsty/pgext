## 用法

- 来源: [GitHub 仓库](https://github.com/PierreSenellart/provsql), [项目文档](https://provsql.org/docs/), [入门指南](https://provsql.org/docs/user/getting-provsql.html)
- ProvSQL 为 PostgreSQL 增加 m-semiring 溯源和不确定性管理能力，支持概率、Shapley 值以及半环求值。

```sql
shared_preload_libraries = 'provsql'
CREATE EXTENSION provsql CASCADE;
```

上游快速安装说明还列出了这些前置条件：PostgreSQL 10 及以上版本、C++17 编译器、PostgreSQL 头文件、`uuid-ossp` 以及 Boost 库。

## 核心流程

ProvSQL 通过 `shared_preload_libraries` 加载，然后使用 `CREATE EXTENSION provsql CASCADE;` 安装。

典型用途包括：

- 在不同半环上计算溯源
- 计算概率和期望值
- 计算 Shapley 值等博弈论贡献
- 使用内置的编译型半环处理常见场景

## 备注

项目主页和文档位于 [provsql.org](https://provsql.org/)。README 指向用户指南，涵盖完整的安装和测试流程。
