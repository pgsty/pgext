## 用法

来源：

- [官方上游 README](https://github.com/heterodb/toybox/blob/e4e9b0526432809bc2f6d50df35b23237ad6c74e/README.md)
- [官方扩展控制文件 (g_cube.control)](https://github.com/heterodb/toybox/blob/e4e9b0526432809bc2f6d50df35b23237ad6c74e/g_cube/g_cube.control)
- [官方实现源代码](https://github.com/heterodb/toybox/blob/e4e9b0526432809bc2f6d50df35b23237ad6c74e/g_cube/g_cube.c)

`g_cube` 是一个 PG-Strom 插件，它教会 PG-Strom 如何在 GPU 上表示和执行 PostgreSQL `cube` 和 `earth` 值的选定操作。它扩展了现有类型而不是安装一个新的用户可面对的类型。

### 核心工作流

构建库及其 CUDA 胖二进制文件，针对相同的 PostgreSQL 和 PG-Strom 安装，然后通过 PG-Strom 插件机制加载它。在测试 GPU 执行之前，安装 `cube`，并在需要 `earthdistance` 值时安装 `earth`。

审查过的源代码从 `_PG_init` 注册其描述符；它不提供版本化的 SQL 或独立的 `CREATE EXTENSION g_cube` 工作流。

### 加速能力

- `cube_contains(cube, cube)`
- `cube_contained(cube, cube)`
- `cube_ll_coord(cube, integer)`
- `cube` 和 `earth` 之间的转换
- `cube` 值的 Arrow 解码

### 要求与注意事项

- 审查过的控制文件、注册表或目录证据标识版本 `1.0`。
- 控制文件将扩展标记为可重定位。
- Makefile 需要 `nvcc`、PG-Strom 服务器头文件以及在构建时选择的 GPU 架构。
- C API 直接调用 PG-Strom 的用户额外接口，因此 PostgreSQL、PG-Strom、CUDA 和此插件必须版本兼容。
- 确认目标构建中的 CPU 回退、数值等价性、设备可用性和 Arrow 错误处理。
