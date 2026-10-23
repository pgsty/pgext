## 用法

来源：

- [官方资料](https://github.com/ringerc/postgres_otel_api/blob/4192d644f04a9f52285bae55e565649e19b7cef2/otel_plpgsql/otel_plpgsql.control)
- [官方资料](https://github.com/ringerc/postgres_otel_api/blob/4192d644f04a9f52285bae55e565649e19b7cef2/otel_plpgsql/otel_plpgsql--0.1.sql)
- [官方资料](https://github.com/ringerc/postgres_otel_api/blob/4192d644f04a9f52285bae55e565649e19b7cef2/otel_plpgsql/otel_plpgsql.c)
- [官方资料](https://github.com/ringerc/postgres_otel_api/blob/4192d644f04a9f52285bae55e565649e19b7cef2/README.md)

`otel_plpgsql` 0.1 通过插件钩子观测 PL/pgSQL 函数、过程与语句的执行，属于实验性的 postgres_otel_api 扩展家族。版本化 SQL 脚本仅登记安装信息，真正启用观测的是加载动态库。

### 启用方式

在目标后端会话中加载提供方及适用的导出器，再加载本观测模块。管理员测试会话可执行：

```sql
LOAD 'otel_api';
LOAD 'otel_plpgsql';
SHOW otel_plpgsql.enabled;
SHOW otel_plpgsql.trace_statements;
```

持久保存 span 还需要导出器，上述命令本身不会配置导出器。也可以使用会话预加载；加入服务器全局预加载配置时需要重启。应使用加载本地库所需的权限，并限制导出目标。

### Span 行为

`otel_plpgsql.enabled` 控制观测开关，`otel_plpgsql.trace_statements` 控制语句 span。本次核对的源码默认开启两者。函数 span 以当前活动 span 或传播的根上下文为父级；否则最外层调用可以开始新的链路。错误与事务边界会影响 span 生命周期，应验证异常处理和可控制事务的过程所生成的链路。

项目文档提供未打补丁 PostgreSQL 14+ 的构建路径，更高级的协议传播则需要配套内核补丁。已有的 PL/pgSQL 观测工具可能共用或竞争插件钩子，应测试共存行为。观测可能暴露代码名称与执行信息，也会给每次调用增加开销。

单独构建的桩变体会编译掉遥测调用，用于比较开销；它既不是导出器，也不能用于获得链路。本扩展没有 SQL 查询接口，无需编造用于激活的 SQL 语句。
