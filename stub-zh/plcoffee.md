## 用法

来源：

- [官方 PL/V8 方言文档](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/doc/plv8.md#dialects)
- [官方 CoffeeScript 示例](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/doc/plv8.md#coffeescript-example)
- [官方构建规则](https://github.com/plv8/plv8/blob/3be0e9d2b112cbf0a4783bbdecd442acbe6b0fb3/Makefile)

`plcoffee` 是 PL/V8 随附的可选 CoffeeScript 过程语言方言。它编译 CoffeeScript 函数体，并交给与 PL/V8 相同的 V8 后端处理器执行。版本 `3.1.10` 属于遗留方言：适合维护现有函数，不应作为新数据库代码的默认语言。

### 核心流程

服务器软件包必须包含可选的 `plcoffee.control` 与版本化 SQL 文件。超级用户创建语言后，获准用户即可用它定义函数：

```sql
CREATE EXTENSION plcoffee;

CREATE FUNCTION plcoffee_test(keys text[], vals text[])
RETURNS text AS $$
  return JSON.stringify(keys.reduce(((o, key, idx) ->
    o[key] = vals[idx]; return o), {}), {})
$$ LANGUAGE plcoffee IMMUTABLE STRICT;

SELECT plcoffee_test(ARRAY['name', 'age'], ARRAY['Tom', '29']);
```

结果是由 CoffeeScript 函数体组装的 JSON 字符串。

### 对象与运行时

`CREATE EXTENSION plcoffee` 会在 `pg_catalog` 中安装可信语言 `plcoffee`，以及调用、内联、校验处理器函数。control 与 SQL 文件由 PL/V8 构建过程生成，加载共享模块 `$libdir/plv8`；没有单独的 CoffeeScript 服务器库，也不需要预加载设置。

### 注意事项

PL/V8 可以使用 `DISABLE_DIALECT` 构建，此时不会产生 `plcoffee` 扩展文件。应确认软件包确实包含该方言，并与已安装的 PL/V8 二进制和版本匹配。虽然最终过程语言标记为 trusted，创建扩展仍仅限超级用户。

CoffeeScript 及其 PL/V8 方言都属于遗留接口。JavaScript 会在 PostgreSQL 后端内消耗内存与 CPU，因此应限制函数定义权限，审查易变性与 strict 声明，并使用完全一致的 PostgreSQL、PL/V8、V8、CoffeeScript 构建组合测试升级。
