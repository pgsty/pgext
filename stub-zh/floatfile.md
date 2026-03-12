

## 用法

> [floatfile: PostgreSQL 的浮点数组文件 I/O](https://github.com/pjungwir/floatfile)

将浮点数组存储在独立文件中，无需 MVCC 开销即可实现极快的查询。非常适合需要低延迟读取和低成本追加写入的时间序列数据。

### 函数

#### 将浮点数组保存到文件

```sql
SELECT save_floatfile('my_data', ARRAY[1.0, 2.0, 3.0, 4.0]::float[]);
```

#### 从文件加载浮点数组

```sql
SELECT load_floatfile('my_data');
-- {1,2,3,4}
```

#### 向现有文件追加数据

```sql
SELECT extend_floatfile('my_data', ARRAY[5.0, 6.0]::float[]);
```

#### 删除浮点文件

```sql
SELECT drop_floatfile('my_data');
```

### 表空间变体

所有函数都接受可选的表空间名称作为第一个参数：

```sql
SELECT save_floatfile('my_ts', 'my_data', ARRAY[1.0, 2.0]::float[]);
SELECT load_floatfile('my_ts', 'my_data');
SELECT extend_floatfile('my_ts', 'my_data', ARRAY[3.0]::float[]);
SELECT drop_floatfile('my_ts', 'my_data');
```

### 直方图函数

直接从浮点文件计算直方图（当数组超过 PostgreSQL 的 1GB 限制时非常有用）：

```sql
SELECT floatfile_to_hist('my_data', 0.0, 1.0, 10);
-- 返回 int[] 直方图计数

SELECT floatfile_to_hist2d('xs_file', 'ys_file', 0.0, 0.0, 1.0, 1.0, 10, 10);
-- 返回二维 int[] 直方图
```

### 并发控制

函数使用 PostgreSQL 咨询锁：`load_floatfile` 获取共享锁；`save`、`extend` 和 `drop` 获取排他锁。
