## 用法

来源：

- [官方 PL/V8 方言文档](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/doc/plv8.md#dialects)
- [官方 LiveScript 示例](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/doc/plv8.md#livescript-example)
- [官方构建规则](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/Makefile)

`plls` 是 PL/V8 随附的可选 LiveScript 过程语言方言。它编译 LiveScript 函数体，并交给 V8 后端的 PL/V8 处理器。版本 `3.1.10` 属于遗留方言，主要用于兼容现有数据库函数。

### 核心流程

服务器软件包必须包含可选的 `plls.control` 与版本化 SQL 文件。超级用户创建语言后，获准用户即可定义 LiveScript 函数：

```sql
CREATE EXTENSION plls;

CREATE FUNCTION plls_test(keys text[], vals text[])
RETURNS text AS $$
  return JSON.stringify { [key, vals[idx]] for key, idx in keys }
$$ LANGUAGE plls IMMUTABLE STRICT;

SELECT plls_test(ARRAY['name', 'age'], ARRAY['Tom', '29']);
```

结果是由 LiveScript 函数体组装的 JSON 字符串。

### 对象与运行时

`CREATE EXTENSION plls` 会在 `pg_catalog` 中安装可信语言 `plls`，以及调用、内联、校验处理器函数。生成的 control/SQL 文件加载 `$libdir/plv8`；LiveScript 没有单独的 PostgreSQL 共享库，也不需要预加载设置。

### 注意事项

使用 `DISABLE_DIALECT` 构建 PL/V8 时会省略 `plls` 文件。应确认已安装软件包包含该方言，并与 PL/V8 二进制和版本匹配。虽然最终语言标记为 trusted，创建扩展仍需要超级用户。

LiveScript 及其 PL/V8 集成属于遗留接口。执行的代码会在数据库后端中消耗内存与 CPU。应限制函数定义权限，审查易变性与 strict 声明，并在升级前对完全一致的 PostgreSQL/PL/V8/V8/LiveScript 组合执行回归测试。
