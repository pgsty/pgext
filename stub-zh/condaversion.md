## 用法

来源：

- [扩展 control 文件](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/condaversion.control)
- [版本 1.0 安装 SQL](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/condaversion--1.0.sql)
- [上游回归测试](https://github.com/JeanChristopheMorinPerso/data_experiments/tree/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/tests/sql)

`condaversion` 版本 `1.0` 增加了一个由 C++ 与 libmamba 版本解析实现的 Conda 感知标量类型。它适合用于排序和索引优先级不同于词法文本顺序的 Conda 软件包版本。

### 核心流程

```sql
CREATE EXTENSION condaversion;

CREATE TABLE conda_release (
    name text NOT NULL,
    version condaversion NOT NULL
);

INSERT INTO conda_release VALUES
    ('demo', '1.0'),
    ('demo', '1.0.0'),
    ('demo', '1.1a1');

SELECT version::text
FROM conda_release
ORDER BY version;

CREATE INDEX conda_release_version_idx
    ON conda_release (version);
```

该扩展定义了输入/输出及二进制发送/接收函数、比较操作符 `=`、`!=`、`<`、`>`、`<=` 和 `>=`、涉及文本类型的隐式或赋值转换，以及默认 B-tree 与哈希操作符类。上游测试明确展示了 `1.0` 与 `1.0.0` 会被比较为相等，而不是两个不同字符串。

### 运维说明

control 文件支持重定位，且没有声明预加载依赖。构建模块需要上游 C++/libmamba 环境；如果缺少匹配的共享库，仅有 SQL 扩展文件并不足够。生产使用前，应在目标 PostgreSQL 与 libmamba 版本上验证备份恢复、索引行为和 ABI 兼容性。
