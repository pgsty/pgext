


## Usage

> [pllua: Lua as a procedural language](https://github.com/pllua/pllua)

`pllua` enables writing PostgreSQL functions in Lua (5.3, 5.4, or LuaJIT 2.1).

```sql
CREATE EXTENSION pllua;
```

### Create Functions

```sql
CREATE FUNCTION lua_max(a integer, b integer) RETURNS integer LANGUAGE pllua AS $$
  if a > b then return a else return b end
$$;

CREATE FUNCTION hello(name text) RETURNS text LANGUAGE pllua AS $$
  return "Hello, " .. name .. "!"
$$;
```

### Data Type Handling

Arguments are automatically converted: integers/floats to Lua numbers, text/varchar to strings, booleans to Lua booleans, NULL to nil. Other types remain as datum objects.

Construct typed values with `pgtype`:

```lua
pgtype.numeric(1234)
pgtype.date('2017-12-01')
pgtype.array.integer(1, 2, 3, 4)
pgtype.numrange(1, 2)
```

### Composite Types (Rows)

```lua
row.columnname     -- access by name
row[3]             -- access by attribute number
for colname, value, attnum in pairs(row) do ... end
```

### Set-Returning Functions

```sql
CREATE FUNCTION generate_n(n integer) RETURNS SETOF integer LANGUAGE pllua AS $$
  for i = 1, n do
    coroutine.yield(i)
  end
$$;
```

### SPI Database Access

```lua
-- Simple query
local rows = spi.execute("SELECT * FROM mytable WHERE id = $1", 42)

-- Row iterator
for row in spi.rows("SELECT * FROM mytable") do
  print(row.name)
end

-- Prepared statements
local stmt = spi.prepare("SELECT * FROM users WHERE id = $1", {'integer'})
local result = stmt:execute(42)
for row in stmt:rows(42) do ... end
```

### Cursors

```lua
local cursor = spi.newcursor()
cursor:open("SELECT * FROM items")
local rows = cursor:fetch(10)
cursor:move(5)
cursor:close()
```

### Trigger Functions

```sql
CREATE FUNCTION my_trigger() RETURNS trigger LANGUAGE pllua AS $$
  function(trigger, old, new)
    trigger.row = new
    return trigger.row
  end
$$;
```

Trigger fields: `trigger.event` (INSERT/UPDATE/DELETE), `trigger.when` (BEFORE/AFTER), `trigger.level` (ROW/STATEMENT), `trigger.new`, `trigger.old`, `trigger.row`.

### Error Handling

```lua
spi.error('division_by_zero', 'Cannot divide by zero')
spi.notice('informational message')
spi.warning('warning message')

-- Subtransactions with pcall
local ok, err = pcall(function()
  spi.execute("INSERT INTO mytable VALUES ($1)", val)
end)
```

### Logging

```lua
print("info message")
spi.debug("debug")
spi.notice("notice")
spi.warning("warning")
spi.error("error")
```
