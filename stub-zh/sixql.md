## 用法

来源：

- [官方 README](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/readme.md)
- [扩展实现](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/src/lib.rs)
- [Cargo 与 PostgreSQL 兼容性](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/Cargo.toml)

`sixql` 0.0.0 是已经放弃的 pgrx 实验，尝试从 PostgreSQL 渲染 SIXEL 终端图形。固定版本实现只是原型，并非绘图 API：函数返回 NULL，生成的转义数据通过服务器日志或进程标准输出写出，而不是经 PostgreSQL 线协议返回。

### 可用实验

超级用户可以在带 `libsixel` 的测试服务器上安装原生模块，然后调用两个生成的 SQL 函数：

```sql
CREATE EXTENSION sixql;

SELECT hello_sixql();

SELECT line_plot(ARRAY[0.0, 1.0, 0.5]::real[]);
```

`hello_sixql()` 构造固定 SIXEL 样本并返回 NULL。`line_plot(real[])` 渲染固定的 100×100 正弦曲线，也返回 NULL；此版本会忽略输入点。两个函数都不会返回普通 SQL 客户端可消费的图像、字节数组、文件路径或表值。

### 终端与服务器边界

SIXEL 需要兼容终端和原生 `libsixel`，但 PostgreSQL 线协议与普通日志路径并不构成可靠的二进制终端通道。服务器日志可能转义、分割、添加前缀、抑制或重定向控制序列，从服务器进程发出也不会指向用户的 `psql` 终端。不要在共享日志中启用原始控制序列渲染。

Cargo 功能通过 pgrx 0.8.3 面向 PostgreSQL 11–15。代码使用 unsafe FFI 回调与实验性 SIXEL crate，而且没有有效的 PostgreSQL 回归测试。它只能作为未来显式返回图像载荷设计的源码素材；不适合生产或不可信调用者。
