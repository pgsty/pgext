## 用法

来源：

- [官方 README](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/README.md)
- [9.4 版扩展 SQL](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/pgorafunc--9.4.sql)
- [扩展控制文件](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/pgorafunc.control)

`pgorafunc` 为 PostgreSQL 提供一部分兼容 Oracle 的函数、操作符、数据类型与包式 API。它主要用于迁移依赖 Oracle 日期行为或 `dbms_output`、`utl_file` 等包的应用，并不是完整的 Oracle 兼容层。

### 核心流程

在目标数据库安装扩展；会话需要使用未限定的 Oracle 兼容名称时，将其 `oracle` 模式放到搜索路径中足够靠前的位置：

```sql
CREATE EXTENSION pgorafunc;

SET search_path TO oracle, "$user", public, pg_catalog;

SELECT oracle.add_months(oracle.date '2005-05-31 10:12:12', 1);
SELECT oracle.months_between(
  oracle.date '1995-02-02 10:00:00',
  oracle.date '1995-01-01 10:21:11'
);
SELECT oracle.to_date('02/16/09 04:12:12', 'MM/DD/YY HH24:MI:SS');
```

搜索路径的选择会影响实际行为：扩展既在 `oracle` 中安装兼容函数，也提供可与内置可见名称并存的对象；README 明确建议使用上述顺序解析 Oracle 风格名称。共享应用若不宜改变全局搜索路径，应显式限定模式。

### 重要对象

- `oracle.date` 保留日内时间，并提供 Oracle 风格的算术重载。
- 日期辅助函数包括 `add_months`、`last_day`、`next_day`、`months_between`、`trunc`、`round` 与 `to_date`。应针对待迁移工作负载验证其格式模型与舍入行为，不要仅凭 PostgreSQL 中的同名函数推断。
- `dbms_output` 提供会话本地的行队列，包含 `enable`、`put_line`、`get_line` 与 `get_lines` 等操作。
- `utl_file` 暴露服务器端文件操作，包括打开、读取、写入、复制、重命名与删除文件。
- 兼容对象还包括 `dual` 表及项目文档列出的其他 Oracle 风格包。

使用会话本地输出时，先启用队列、写入内容，再读取：

```sql
SELECT dbms_output.enable();
SELECT dbms_output.put_line('migration step complete');
SELECT * FROM dbms_output.get_lines(0);
```

### 运维说明

控制文件声明版本 9.4、扩展可重定位以及模块库，但没有声明预加载要求或扩展依赖。这是一套较旧且已停止维护的兼容代码；将其用于新迁移前，应验证所有必需函数以及目标 PostgreSQL 版本组合。

应把 `utl_file` 视为高权限服务器端能力。其操作针对 PostgreSQL 服务器进程可访问的文件执行，因此要限制执行权限，并在非生产环境验证路径与权限。`dbms_output` 队列仅属于当前会话，会随会话结束而消失。现有 Oracle 应用还可能依赖实现子集之外的隐式转换、区域设置规则或包语义。
