## 用法

来源：

- [PGXN 2.6.0 README](https://pgxn.org/dist/plphp/2.6.0/README.html)
- [jsonb_plphp 控制文件](https://api.pgxn.org/src/plphp/plphp-2.6.0/jsonb_plphp/jsonb_plphp.control)
- [jsonb_plphp 1.0 SQL 定义](https://api.pgxn.org/src/plphp/plphp-2.6.0/jsonb_plphp/jsonb_plphp--1.0.sql)
- [PL/php 语言参考](https://pgxn.org/dist/plphp/2.6.0/doc/plphp.html)

`jsonb_plphp` 为使用 `plphp` 编写的函数提供 `jsonb` 与 PHP 原生数组、标量、布尔值和 null 之间的 PostgreSQL 转换。它是非可信 PL/php 运行时的配套扩展，不是独立过程语言。

### 核心流程

```sql
CREATE EXTENSION jsonb_plphp CASCADE;

CREATE FUNCTION add_seen(jsonb)
RETURNS jsonb
TRANSFORM FOR TYPE jsonb
LANGUAGE plphp
AS $$
    $value = $args[0];
    $value['seen'] = true;
    return $value;
$$;

SELECT add_seen('{"id": 1}'::jsonb);
```

`CASCADE` 会在可用时创建 `plphp`。只有声明 `TRANSFORM FOR TYPE jsonb` 的函数才会收到 PHP 原生值。转换也适用于受支持的嵌套复合类型、数组、集合返回与触发器上下文；SPI 结果行仍走常规文本转换路径。

### 边界

`jsonb_plphp` 与 `plphp` 具有相同的超级用户安装与信任模型：PHP 代码可访问 PostgreSQL 服务器账户的文件、网络和进程环境。JSON 对象会变为 PHP 关联数组，因此应审查键转换以及数组、对象、数值与 null 的区别。大型或深层嵌套文档需要测试内存用量与转换行为。

该转换的控制版本是 `1.0`，随 PL/php 2.6.0 发行包提供。它无需预加载，但 PHP embed SAPI 与 `plphp` 共享库必须已能在 PostgreSQL 11 至 18 上使用。
