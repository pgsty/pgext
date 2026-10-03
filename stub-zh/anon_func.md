## 用法

来源：

- [cc/postgres/anon_func.control](https://github.com/google/differential-privacy/blob/49f9a34a5c6da7459ca4f99c391f3b8e595c3bfe/cc/postgres/anon_func.control)
- [cc/postgres/README.md](https://github.com/google/differential-privacy/blob/49f9a34a5c6da7459ca4f99c391f3b8e595c3bfe/cc/postgres/README.md)
- [cc/postgres/anon_func--1.0.0.sql](https://github.com/google/differential-privacy/blob/49f9a34a5c6da7459ca4f99c391f3b8e595c3bfe/cc/postgres/anon_func--1.0.0.sql)
- [cc/postgres/install_extension.sh](https://github.com/google/differential-privacy/blob/49f9a34a5c6da7459ca4f99c391f3b8e595c3bfe/cc/postgres/install_extension.sh)

`anon_func` 将 Google differential-privacy 的匿名计数、求和、均值、方差等算法暴露为聚合函数。扩展 SQL 版本为 1.0.0，与整个仓库的库版本不同。

### 核心用法

```sql
CREATE EXTENSION anon_func;
SELECT anon_count(value, 1.0)
FROM (VALUES (1), (2), (3)) AS sample(value);
```

### 运行边界

官方安装说明面向 PostgreSQL 12，需要编译后的 C++ 库；这里不扩展兼容性声明。安装需要特权，没有预加载要求。Epsilon 和可选上下界作为聚合参数字面值传入。实现假定每条输入来自不同用户，因此调用方需预先限制个人贡献量。空组返回 NULL，上游明确说明这种情况不具备差分隐私。重复查询会消耗隐私预算，仅靠这些封装不能建立完整的隐私策略。
