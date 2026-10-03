## 用法

来源：

- [README.md](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/README.md)
- [copy_arrow.control](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/copy_arrow.control)
- [copy_arrow--0.0.1.sql](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/copy_arrow--0.0.1.sql)
- [meson.build](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/meson.build)
- [sql/copy_to_arrow/int32.sql](https://github.com/kou/pg-copy-arrow/blob/b513d764112b48c2803280e64c56c046b2b6eb6b/sql/copy_to_arrow/int32.sql)

`copy_arrow` 0.0.1 是实验性的 Apache Arrow 流式导出与 COPY 格式扩展。它面向上游 copy-format-extendable-2024-11-v26 PostgreSQL 分支，不能据此认定兼容普通 PostgreSQL 发行版。

### 核心工作流

```sql
CREATE EXTENSION copy_arrow;
CREATE TABLE arrow_example (value integer);
INSERT INTO arrow_example SELECT generate_series(1, 10);
SELECT format_arrow(copy_to_arrow('arrow_example'::regclass));
```

### 对象与要求

`copy_to_arrow(regclass)` 与 `scan_to_arrow(regclass)` 返回二进制 `bytea` 数据；`format_arrow(bytea)` 将其格式化为文本。所核验的回归用例导出一个整数列。构建需要 Apache Arrow C++ 和打过补丁的服务端头文件。安装会创建原生 C 函数，因此需要超级用户；控制文件允许重定位。上游未要求预加载或独立后台进程。该实现属于概念验证，使用前应核对具体补丁分支与数据类型。
