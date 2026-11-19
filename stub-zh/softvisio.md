## 用法

来源：

- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/README.md)
- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/softvisio.control)
- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/softvisio--2.0.4.sql)
- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/Makefile)
- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/package.json)
- [Official source](https://www.postgresql.org/docs/18/functions-string.html)

softvisio 2.0.4 以 SQL/PL/pgSQL 提供数值域、Unicode 辅助函数和大对象大小查询函数，不需要原生库或预加载。项目将包标记为专有软件。安装 SQL 使用 PostgreSQL 18 提供的 casefold 函数，上游没有公布完整的大版本支持矩阵。

### 核心流程

```sql
CREATE EXTENSION softvisio;
SELECT 9007199254740991::int53;
SELECT 12.34::number2;
SELECT normalize_nfc('é'), normalize_nfc_casefold('Straße');
```

int53 将 bigint 限制在 JavaScript 能精确表示的整数范围内。number1 至 number16 是 numeric(16, scale) 数值域，并额外限制绝对值上界且不包含边界。先按声明的小数位舍入，再执行域约束；这些类型不保证任意精度或无损的二进制浮点转换。

### 使用边界

lo_size(oid) 以只读方式打开大对象，定位到末尾后关闭描述符；调用者仍需普通的大对象权限，并在事务中使用。规范化辅助函数标记为 strict、immutable、parallel safe。Unicode 规范化要求 UTF8，大小写折叠遵循输入的排序规则。

控制文件标记 relocatable 和 superuser=false，安装仍需正常的数据库与模式创建权限。安装脚本发现可见的同名域时直接复用，不核对其定义；应使用受控模式与 search_path，并事先检查现有 int53/number 域。所附 Makefile 包含经过核对、通向 2.0.4 的升级链。升级前备份依赖对象并检查可用的扩展更新路径。
