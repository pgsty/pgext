## 用法

来源：

- [PGXN 2.6.0 README](https://pgxn.org/dist/plphp/2.6.0/README.html)
- [hstore_plphp 控制文件](https://api.pgxn.org/src/plphp/plphp-2.6.0/hstore_plphp/hstore_plphp.control)
- [hstore_plphp 1.0 SQL 定义](https://api.pgxn.org/src/plphp/plphp-2.6.0/hstore_plphp/hstore_plphp--1.0.sql)
- [PL/php 语言参考](https://pgxn.org/dist/plphp/2.6.0/doc/plphp.html)

`hstore_plphp` 为使用 `plphp` 编写的函数在 PostgreSQL `hstore` 与 PHP 原生关联数组之间转换。PHP 字符串键与字符串或 null 值会映射为相应的 `hstore` 条目。

### 核心流程

```sql
CREATE EXTENSION hstore_plphp CASCADE;

CREATE FUNCTION add_label(hstore)
RETURNS hstore
TRANSFORM FOR TYPE hstore
LANGUAGE plphp
AS $$
    $value = $args[0];
    $value['label'] = 'reviewed';
    return $value;
$$;

SELECT add_label('id=>1'::hstore);
```

`CASCADE` 会在可用时创建所需的 `hstore` 与 `plphp` 扩展。只有声明 `TRANSFORM FOR TYPE hstore` 的函数使用 PHP 原生数组。PHP null 值映射为 `hstore` 内的 SQL NULL 值，而不是文本字符串 `NULL`。

### 边界

`hstore_plphp` 继承 `plphp` 的非可信、仅超级用户安全模型，不会为 PHP 提供沙箱。转换覆盖受支持的参数、结果、嵌套复合或数组值、集合返回行与触发器行，而 SPI 结果行仍使用常规文本转换。替换显式文本解析前应测试键值编码与 null 行为。

该转换的控制版本是 `1.0`，随 PL/php 2.6.0 提供。它无需预加载，但要求 PostgreSQL 11 至 18 上已有可用的 PHP embed SAPI 与 `plphp`。
