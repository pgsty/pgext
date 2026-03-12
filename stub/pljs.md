


## Usage

> [pljs: PL/JavaScript trusted procedural language](https://github.com/plv8/pljs)

`pljs` enables writing PostgreSQL functions in JavaScript using the QuickJS engine.

```sql
CREATE EXTENSION pljs;
DO $$ pljs.elog(NOTICE, "Hello, World!") $$ LANGUAGE pljs;
```

### Create Functions

```sql
CREATE FUNCTION pljs_add(a int, b int) RETURNS int AS $$
  return a + b;
$$ LANGUAGE pljs;

SELECT pljs_add(1, 2);  -- 3
```

### Database Access

```sql
CREATE FUNCTION get_users() RETURNS SETOF json AS $$
  var rows = pljs.execute('SELECT * FROM users');
  for (var i = 0; i < rows.length; i++) {
    pljs.return_next(JSON.stringify(rows[i]));
  }
$$ LANGUAGE pljs;
```

Execute with arguments:

```js
var rows = pljs.execute('SELECT * FROM tbl WHERE id = $1', [42]);
var affected = pljs.execute('DELETE FROM tbl WHERE price > $1', [1000]);
```

### Prepared Statements

```js
var plan = pljs.prepare('SELECT * FROM tbl WHERE col = $1', ['int']);
var rows = plan.execute([1]);
plan.free();
```

### Cursors

```js
var plan = pljs.prepare('SELECT * FROM tbl WHERE col = $1', ['int']);
var cursor = plan.cursor([1]);
var row;
while (row = cursor.fetch()) {
    // process row
}
cursor.close();
plan.free();
```

### Subtransactions

```js
try {
  pljs.subtransaction(function() {
    pljs.execute("INSERT INTO tbl VALUES(1)");
    pljs.execute("INSERT INTO tbl VALUES(1/0)"); // error - rolls back
  });
} catch(e) {
  // handle error
}
```

### Logging

```js
pljs.elog(DEBUG1, 'debug message');
pljs.elog(NOTICE, 'notice message');
pljs.elog(WARNING, 'warning message');
pljs.elog(ERROR, 'error message');
```

### Find Other PLJS Functions

```sql
CREATE FUNCTION callee(a int) RETURNS int AS $$ return a * a $$ LANGUAGE pljs;
CREATE FUNCTION caller(a int, t int) RETURNS int AS $$
  var func = pljs.find_function("callee");
  return func(a);
$$ LANGUAGE pljs;
```

### Window Functions

```sql
CREATE FUNCTION my_window_func(val int) RETURNS int AS $$
  var winobj = pljs.get_window_object();
  var pos = winobj.get_current_position();
  var total = winobj.get_partition_row_count();
  return winobj.get_func_arg_in_current(0);
$$ LANGUAGE pljs WINDOW;
```

Window object methods: `get_current_position()`, `get_partition_row_count()`, `set_mark_position(pos)`, `rows_are_peers(pos1, pos2)`, `get_func_arg_in_partition(argno, relpos, seektype, mark_pos)`, `get_func_arg_in_frame(argno, relpos, seektype, mark_pos)`, `get_func_arg_in_current(argno)`, `get_partition_local()`, `set_partition_local(obj)`.

### Utility Functions

```sql
SELECT pljs_info();     -- memory and stack usage as JSON
SELECT pljs_version();  -- extension version
```
