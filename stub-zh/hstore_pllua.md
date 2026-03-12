

## 用法

> [hstore_pllua: PL/Lua 的 hstore 转换支持](https://github.com/pllua/pllua)

`hstore_pllua` 扩展提供了一个转换器，允许 `hstore` 值在 PL/Lua 函数中作为原生 Lua 表直接传递和返回，而非不透明的 datum 对象。

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION pllua;
CREATE EXTENSION hstore_pllua;
```

### 在 PL/Lua 中使用 hstore

安装转换器后，`hstore` 参数会自动转换为 Lua 表，Lua 表也可以作为 `hstore` 值返回：

```sql
CREATE FUNCTION process_hstore(h hstore) RETURNS hstore LANGUAGE pllua AS $$
  -- h 是一个具有字符串键和字符串值的 Lua 表
  h["new_key"] = "new_value"
  h["modified"] = "true"
  return h
$$;

SELECT process_hstore('name => "Alice", age => "30"'::hstore);
```

### 遍历 hstore 键

```sql
CREATE FUNCTION hstore_keys(h hstore) RETURNS SETOF text LANGUAGE pllua AS $$
  for k, v in pairs(h) do
    coroutine.yield(k)
  end
$$;
```

如果没有此转换器，`hstore` 值需要使用 `hstore` 类型函数手动转换。
