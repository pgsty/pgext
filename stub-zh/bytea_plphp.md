## 用法

来源：

- [PGXN 2.6.0 README](https://pgxn.org/dist/plphp/2.6.0/README.html)
- [bytea_plphp 控制文件](https://api.pgxn.org/src/plphp/plphp-2.6.0/bytea_plphp/bytea_plphp.control)
- [bytea_plphp 1.0 SQL 定义](https://api.pgxn.org/src/plphp/plphp-2.6.0/bytea_plphp/bytea_plphp--1.0.sql)
- [PL/php 2.6.0 变更日志](https://api.pgxn.org/src/plphp/plphp-2.6.0/CHANGELOG.md)

`bytea_plphp` 为使用 `plphp` 编写的函数在 PostgreSQL `bytea` 与二进制安全的 PHP 字符串之间转换。代码需要保留包括内嵌 NUL 在内的原始字节、且不希望经过 PostgreSQL 文本 `\x...` 表示时可以使用它。

### 核心流程

```sql
CREATE EXTENSION bytea_plphp CASCADE;

CREATE FUNCTION bytea_roundtrip(bytea)
RETURNS bytea
TRANSFORM FOR TYPE bytea
LANGUAGE plphp
AS $$
    return $args[0];
$$;

SELECT encode(bytea_roundtrip(decode('000102ff', 'hex')), 'hex');
```

`CASCADE` 会在可用时创建 `plphp`。只有声明 `TRANSFORM FOR TYPE bytea` 的函数才会收到原始二进制 PHP 字符串。该转换也可用于受支持的嵌套复合、数组、集合返回与触发器上下文。

### 边界

二进制字符串可能包含 NUL 与无效文本编码，不应交给假设 UTF-8 或 C 风格终止符的 PHP 文本 API。返回 PHP 字符串会逐字节写入，因此需要十六进制或 base64 解码时必须显式完成。大值应测试后端内存压力，并避免在日志中记录二进制载荷或秘密。

`bytea_plphp` 继承 `plphp` 的非可信、仅超级用户安全模型。转换控制版本为 `1.0`，在 PL/php 2.6.0 发行包中引入，要求 PostgreSQL 11 至 18 上已有 PHP embed SAPI 与 `plphp`，无需预加载。
