

## 用法

> [hstore_plluau: 不可信 PL/Lua 的 hstore 转换支持](https://github.com/pllua/pllua)

`hstore_plluau` 扩展提供了一个转换器，允许 `hstore` 值在不可信 PL/Lua（`plluau`）函数中作为原生 Lua 表直接传递和返回。

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION plluau;
CREATE EXTENSION hstore_plluau;
```

### 在 plluau 中使用 hstore

安装转换器后，`hstore` 参数会自动转换为 Lua 表：

```sql
CREATE FUNCTION process_hstore(h hstore) RETURNS hstore LANGUAGE plluau AS $$
  -- h 是一个具有字符串键和字符串值的 Lua 表
  h["processed"] = "true"
  -- 还可以执行不受限操作（文件 I/O 等）
  return h
$$;

SELECT process_hstore('key1 => "val1", key2 => "val2"'::hstore);
```

这是 `hstore_pllua` 的不可信对应版本。行为完全相同——`hstore` 值变为 Lua 表，反之亦然——但函数在不受限的 `plluau` 环境中运行。
