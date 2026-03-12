

## 用法

> [plluau: Lua 不可信过程语言](https://github.com/pllua/pllua)

`plluau` 是 `pllua` 的不可信变体，允许 Lua 函数访问文件系统、加载任意模块，以及执行在可信版本中受限的操作。

```sql
CREATE EXTENSION plluau;
```

### 创建函数

```sql
CREATE FUNCTION read_file(path text) RETURNS text LANGUAGE plluau AS $$
  local f = io.open(path, "r")
  if f then
    local content = f:read("*a")
    f:close()
    return content
  end
  return nil
$$;
```

### 与 pllua（可信版本）的区别

| 功能 | pllua（可信） | plluau（不可信） |
|------|--------------|-----------------|
| 文件 I/O | 受限 | 完全访问 |
| 模块加载 | 仅白名单 | 无限制 |
| 操作系统访问 | 否 | 是 |
| 适用于 | 用户自定义函数 | 管理员/超级用户函数 |

### 与 pllua 相同的 API

`plluau` 共享与 `pllua` 相同的 SPI 接口、触发器支持、集合返回函数和数据类型处理。所有 SPI 函数（`spi.execute`、`spi.prepare`、`spi.rows`）、基于协程的集合返回以及触发器函数的工作方式完全相同。

```sql
CREATE FUNCTION run_command(cmd text) RETURNS text LANGUAGE plluau AS $$
  local handle = io.popen(cmd)
  local result = handle:read("*a")
  handle:close()
  return result
$$;
```

### 初始化

通过 GUC 参数配置（仅限超级用户）：

```sql
SET pllua.on_untrusted_init = 'myvar = {}';
```

由于 `plluau` 提供的访问权限不受限制，只有超级用户才能创建 `plluau` 函数。
