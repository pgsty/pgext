## 用法

来源：

- [README.md](https://github.com/postsql/pl-jsonschema/blob/e4f2472671ed5d1d336d8afcdf6b6cde05bb286f/README.md)
- [Makefile](https://github.com/postsql/pl-jsonschema/blob/e4f2472671ed5d1d336d8afcdf6b6cde05bb286f/Makefile)
- [pl_jsonschema.sql.in](https://github.com/postsql/pl-jsonschema/blob/e4f2472671ed5d1d336d8afcdf6b6cde05bb286f/pl_jsonschema.sql.in)
- [pl_jsonschema.control](https://github.com/postsql/pl-jsonschema/blob/e4f2472671ed5d1d336d8afcdf6b6cde05bb286f/pl_jsonschema.control)

`pl_jsonschema` 安装受信任语言 `pl/jsonschema`，函数体是由 Ajv 编译的 JSON Schema。带连字符的 `pl-jsonschema` 是同一实现的另一安装名，两者不能同时安装。

### 核心用法

```sql
CREATE EXTENSION pljs;
SET pl_jsonschema.engine = 'pljs';
CREATE EXTENSION pl_jsonschema;
CREATE FUNCTION valid_profile(value json) RETURNS boolean
LANGUAGE "pl/jsonschema" IMMUTABLE STRICT AS $$
{"type":"object","properties":{"name":{"type":"string"}},"required":["name"]}
$$;
SELECT valid_profile('{"name":"Alice"}'::json);
```

### 运行边界

先安装具备 JavaScript 语言处理器能力的 `plv8` 或 `pljs`。两者同时存在时默认选择 `plv8`，也可在安装前设置 `pl_jsonschema.engine`。扩展控制文件本身未标记为 trusted，因此创建扩展要求超级用户；创建出来的过程语言则是受信任语言。

验证器必须是标量函数，只有一个 json 类型的 IN 参数，并返回 boolean；不支持过程、集合返回函数或其他参数模式。`$id` 与 `$ref` 在会话内关联模式，替换验证器会使缓存失效。若模式语义适合，可用 immutable、strict 验证器构造检查约束。

上游在 PostgreSQL 19 测试，并要求相应的 JavaScript 处理器能力；不能假设旧版 PLV8/PLJS 软件包都支持。这个 SQL/JavaScript 扩展没有自己的共享库，也不要求预加载。
