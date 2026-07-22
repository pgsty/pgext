## 用法

来源：

- [扩展 control 文件](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/conda_pgsql_rust_ext.control)
- [Rust 实现](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/src/lib.rs)
- [上游回归示例](https://github.com/JeanChristopheMorinPerso/data_experiments/tree/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/tests/pg_regress/sql)

`conda_pgsql_rust_ext` 提供由 `rattler_conda_types` 支撑的 `condaversion` 类型。当 Conda 软件包版本需要按照 Conda 版本规则解析、比较、索引或聚合，而不能使用普通文本排序时，可使用该扩展。

### 核心流程

创建扩展，在强类型列中保存版本，然后使用常规比较与聚合语法：

```sql
CREATE EXTENSION conda_pgsql_rust_ext;

CREATE TABLE package_release (
    package text NOT NULL,
    version condaversion NOT NULL
);

INSERT INTO package_release VALUES
    ('demo', '1.1dev1'),
    ('demo', '1.1.0'),
    ('demo', '1!0.4.1');

SELECT version::text
FROM package_release
WHERE package = 'demo'
ORDER BY version;

SELECT min(version), max(version)
FROM package_release;
```

该类型接受 Conda 版本字符串，在文本输出中保留原始拼写，并提供相等、排序、哈希以及 `min(condaversion)` 和 `max(condaversion)` 聚合。无效的 Conda 语法会触发输入错误。

### 运维说明

已核验的上游修订版将 crate 版本标为 `0.0.0`。control 文件将扩展标记为不可重定位、不受信任且仅限超级用户安装；其中没有声明预加载要求。应把这一早期版本号视为 API 稳定性警告，并在将其用于约束或唯一索引前，使用软件包元数据中的实际版本格式测试比较行为。
