

## 用法

> [pllua: Lua 过程语言](https://github.com/pllua/pllua)

`pllua` 允许在 PostgreSQL 中使用 Lua（5.3、5.4 或 LuaJIT 2.1）编写函数。

```sql
CREATE EXTENSION pllua;
```

### 创建函数

```sql
CREATE FUNCTION lua_max(a integer, b integer) RETURNS integer LANGUAGE pllua AS $$
  if a > b then return a else return b end
$$;

CREATE FUNCTION hello(name text) RETURNS text LANGUAGE pllua AS $$
  return "Hello, " .. name .. "!"
$$;
```

### 数据类型处理

参数会自动转换：整数/浮点数转为 Lua 数字，text/varchar 转为字符串，布尔值转为 Lua 布尔值，NULL 转为 nil。其他类型保持为 datum 对象。

使用 `pgtype` 构造类型化的值：

```lua
pgtype.numeric(1234)
pgtype.date('2017-12-01')
pgtype.array.integer(1, 2, 3, 4)
pgtype.numrange(1, 2)
```

### 复合类型（行）

```lua
row.columnname     -- 按名称访问
row[3]             -- 按属性编号访问
for colname, value, attnum in pairs(row) do ... end
```

### 集合返回函数

```sql
CREATE FUNCTION generate_n(n integer) RETURNS SETOF integer LANGUAGE pllua AS $$
  for i = 1, n do
    coroutine.yield(i)
  end
$$;
```

### SPI 数据库访问

```lua
-- 简单查询
local rows = spi.execute("SELECT * FROM mytable WHERE id = $1", 42)

-- 行迭代器
for row in spi.rows("SELECT * FROM mytable") do
  print(row.name)
end

-- 预备语句
local stmt = spi.prepare("SELECT * FROM users WHERE id = $1", {'integer'})
local result = stmt:execute(42)
for row in stmt:rows(42) do ... end
```

### 游标

```lua
local cursor = spi.newcursor()
cursor:open("SELECT * FROM items")
local rows = cursor:fetch(10)
cursor:move(5)
cursor:close()
```

### 触发器函数

```sql
CREATE FUNCTION my_trigger() RETURNS trigger LANGUAGE pllua AS $$
  function(trigger, old, new)
    trigger.row = new
    return trigger.row
  end
$$;
```

触发器字段：`trigger.event`（INSERT/UPDATE/DELETE）、`trigger.when`（BEFORE/AFTER）、`trigger.level`（ROW/STATEMENT）、`trigger.new`、`trigger.old`、`trigger.row`。

### 错误处理

```lua
spi.error('division_by_zero', 'Cannot divide by zero')
spi.notice('informational message')
spi.warning('warning message')

-- 使用 pcall 进行子事务
local ok, err = pcall(function()
  spi.execute("INSERT INTO mytable VALUES ($1)", val)
end)
```

### 日志记录

```lua
print("info message")
spi.debug("debug")
spi.notice("notice")
spi.warning("warning")
spi.error("error")
```
