## 用法

来源：

- [官方 README](https://github.com/glynastill/pg_jsonb_opx/blob/adb985a8c3192e2577bf7ef863560baf4e566c5e/README.md)
- [扩展控制文件](https://github.com/glynastill/pg_jsonb_opx/blob/adb985a8c3192e2577bf7ef863560baf4e566c5e/jsonb_opx.control)

`jsonb_opx` 提供 PostgreSQL 9.4 初版 `jsonb` 所缺少的 hstore 风格修改运算符。版本 `1.0` 支持浅层删除、拼接与替换；它主要适合维护依赖这套历史运算符语义的应用。

### 核心流程

```sql
CREATE EXTENSION jsonb_opx;

SELECT '{"a":1,"b":2}'::jsonb - 'a';
SELECT '{"a":1}'::jsonb || '{"b":2}'::jsonb;
SELECT '{"a":1,"b":2}'::jsonb #= '{"a":9}'::jsonb;
```

`-` 的重载会删除匹配的标量或标量数组，不会遍历嵌套对象。`||` 拼接两个 `jsonb` 值，`#=` 使用另一个 `jsonb` 文档替换顶层值。

### 用户接口

- `jsonb_delete(jsonb, text|numeric|boolean)`：按标量值执行浅层删除。
- `jsonb_delete(jsonb, text[]|numeric[]|boolean[])`：按值数组执行浅层删除。
- `jsonb_delete(jsonb, jsonb)`：由另一个 JSON 值驱动浅层删除。
- `jsonb_concat(jsonb, jsonb)`：`||` 背后的实现函数。
- `jsonb_replace(jsonb, jsonb)`：`#=` 背后的实现函数。
- `jsonb_delete_path(jsonb, text[])` 与 `jsonb_replace_path(jsonb, text[], jsonb)`：上游版本 `1.1` 描述的嵌套路径函数。

### 版本与兼容性说明

控制文件的默认版本是 `1.0`。README 还描述了带有 `#-` 和嵌套替换的实验性 `1.1` API，但版本 `1.0` 的软件包并不意味着一定包含它；使用前应查询 `pg_available_extension_versions()`。现代 PostgreSQL 已内置多个拼写相近的 `jsonb` 运算符，因此需要在确切的服务器版本上验证解析与语义。上游明确将项目标为实验性质，并要求用户在生产使用前自行验证。它没有声明预加载或固定模式要求。
