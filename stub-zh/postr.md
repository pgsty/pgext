## 用法

来源：

- [官方 README](https://github.com/mrcsparker/postr/blob/e1dbaa89db80e76a7d0cdff28cf3314047d5db2b/README.md)
- [扩展控制文件](https://github.com/mrcsparker/postr/blob/e1dbaa89db80e76a7d0cdff28cf3314047d5db2b/postr.control)
- [官方示例索引](https://github.com/mrcsparker/postr/blob/e1dbaa89db80e76a7d0cdff28cf3314047d5db2b/examples/README.md)

`postr` 将 Ruby 嵌入 PostgreSQL，并注册名为 `ruby` 的非受信过程语言。SQL 函数、触发器、集合返回函数与内联代码块可在后端进程中执行 Ruby；但它明确是仅限超级用户的原型，并非面向多租户的沙箱语言。

### 核心流程

由可信管理员安装扩展后，以 `def call(...)` 入口定义 Ruby 函数，再通过普通 SQL 调用：

```sql
CREATE EXTENSION postr;

CREATE FUNCTION public.ruby_greet(name text, excited boolean)
RETURNS text
LANGUAGE ruby
AS $ruby$
  def call(name:, excited:)
    message = "hello, #{name}"
    excited ? "#{message}!" : message
  end
$ruby$;

SELECT public.ruby_greet('postgres', true);
```

参数既可以按名称访问，也可通过 `args` 映射和按调用顺序排列的 `argv` 使用。运行时会把常见 PostgreSQL 标量、数组、复合、网络和内置范围值映射为 Ruby 对象。不支持的伪类型与输出形状会在函数验证阶段被拒绝。

### 数据库访问与结果

Ruby 代码可通过 `Postr` 前置模块回调当前 PostgreSQL 事务：

- `Postr.exec` 执行参数化 SQL 并返回处理行数。
- `Postr.select` 返回行哈希或逐行 yield，`Postr.first` 与 `Postr.value` 则选择单行或单值。
- `Postr.param` 为存在歧义的 Ruby 值提供明确的 PostgreSQL 类型。
- `Postr.notice`、`Postr.warn` 与 `Postr.quote_ident` 暴露服务器消息和标识符引用能力。
- `Postr.transaction` 在内部子事务中执行 Ruby 块；它类似保存点，不是独立提交。

该语言还支持 `DO LANGUAGE ruby`、行级与语句级触发器、迭代或 yield 驱动的 `SETOF` 结果、record、`OUT`/`INOUT`/`TABLE` 签名以及可变参数。每种形状的具体返回契约应查阅官方示例。

### 受控加载

Ruby `require` 受 PostgreSQL 设置限制。管理员可用 `postr.extra_load_paths` 与 `postr.allowed_requires` 配置批准的特性，或用 `postr.gem_home` 与 `postr.allowed_gems` 配置随附 gem：

```sql
SET LOCAL postr.extra_load_paths = '/opt/postr/ruby';
SET LOCAL postr.allowed_requires = 'postr_helpers';
```

这些允许列表可以减少意外模块暴露，但不能形成 Ruby 沙箱。策略状态和已编译调用对象按后端缓存，因此部署测试中应使用新会话验证配置与代码变更。

### 安全与兼容性

控制文件将 0.1.0 版标为不可重定位且仅限超级用户。Ruby 在数据库后端内部执行，会使用其内存、CPU 与进程权限；只有高度可信的角色才能定义或修改 `LANGUAGE ruby` 代码。外部库与 gem 处于同一信任边界。

项目的打包冒烟测试目标是 Linux 上的 PostgreSQL 17，而源码特性覆盖 PostgreSQL 13 至 18。Ruby 支持范围以嵌入库 `magnus` 的文档为准，README 也明确区分本地 Ruby 4 测试与兼容性保证。应固定 PostgreSQL、Ruby、扩展与 gem 版本，施加语句和资源控制，并将后端崩溃或阻塞 Ruby 调用视为数据库可用性风险。
