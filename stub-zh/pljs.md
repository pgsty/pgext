

## 用法

> [pljs: PL/JavaScript 可信过程语言](https://github.com/plv8/pljs)

`pljs` 允许使用 QuickJS 引擎在 PostgreSQL 中编写 JavaScript 函数。

```sql
CREATE EXTENSION pljs;
DO $$ pljs.elog(NOTICE, "Hello, World!") $$ LANGUAGE pljs;
```

### 创建函数

```sql
CREATE FUNCTION pljs_add(a int, b int) RETURNS int AS $$
  return a + b;
$$ LANGUAGE pljs;

SELECT pljs_add(1, 2);  -- 3
```

### 数据库访问

```sql
CREATE FUNCTION get_users() RETURNS SETOF json AS $$
  var rows = pljs.execute('SELECT * FROM users');
  for (var i = 0; i < rows.length; i++) {
    pljs.return_next(JSON.stringify(rows[i]));
  }
$$ LANGUAGE pljs;
```

使用参数执行：

```js
var rows = pljs.execute('SELECT * FROM tbl WHERE id = $1', [42]);
var affected = pljs.execute('DELETE FROM tbl WHERE price > $1', [1000]);
```

### 预备语句

```js
var plan = pljs.prepare('SELECT * FROM tbl WHERE col = $1', ['int']);
var rows = plan.execute([1]);
plan.free();
```

### 游标

```js
var plan = pljs.prepare('SELECT * FROM tbl WHERE col = $1', ['int']);
var cursor = plan.cursor([1]);
var row;
while (row = cursor.fetch()) {
    // 处理行
}
cursor.close();
plan.free();
```

### 子事务

```js
try {
  pljs.subtransaction(function() {
    pljs.execute("INSERT INTO tbl VALUES(1)");
    pljs.execute("INSERT INTO tbl VALUES(1/0)"); // 出错 - 回滚
  });
} catch(e) {
  // 处理错误
}
```

### 日志记录

```js
pljs.elog(DEBUG1, 'debug message');
pljs.elog(NOTICE, 'notice message');
pljs.elog(WARNING, 'warning message');
pljs.elog(ERROR, 'error message');
```

### 查找其他 PLJS 函数

```sql
CREATE FUNCTION callee(a int) RETURNS int AS $$ return a * a $$ LANGUAGE pljs;
CREATE FUNCTION caller(a int, t int) RETURNS int AS $$
  var func = pljs.find_function("callee");
  return func(a);
$$ LANGUAGE pljs;
```

### 窗口函数

```sql
CREATE FUNCTION my_window_func(val int) RETURNS int AS $$
  var winobj = pljs.get_window_object();
  var pos = winobj.get_current_position();
  var total = winobj.get_partition_row_count();
  return winobj.get_func_arg_in_current(0);
$$ LANGUAGE pljs WINDOW;
```

窗口对象方法：`get_current_position()`、`get_partition_row_count()`、`set_mark_position(pos)`、`rows_are_peers(pos1, pos2)`、`get_func_arg_in_partition(argno, relpos, seektype, mark_pos)`、`get_func_arg_in_frame(argno, relpos, seektype, mark_pos)`、`get_func_arg_in_current(argno)`、`get_partition_local()`、`set_partition_local(obj)`。

### 实用函数

```sql
SELECT pljs_info();     -- 以 JSON 格式返回内存和栈使用情况
SELECT pljs_version();  -- 扩展版本
```
