---
title: 按语言
description: 按实现语言组织的 PostgreSQL 扩展
excludeSearch: true
weight: 1
---

{{< language "C" >}} {{< language "C++" >}} {{< language "Rust" >}} {{< language "Java" >}} {{< language "Python" >}} {{< language "SQL" >}} {{< language "Data" >}}

## 概览

| 语言 | 数量 | 描述 |
|:-------:|:-----:|:------------|
| {{< language "C" >}} | 335 | 传统的 PostgreSQL 扩展开发语言 |
| {{< language "SQL" >}} | 37 | 纯 SQL 扩展和函数 |
| {{< language "Rust" >}} | 33 | 使用 pgrx 框架用 Rust 编写的扩展 |
| {{< language "Data" >}} | 10 | 仅包含数据的扩展 |
| {{< language "C++" >}} | 6 | 使用 C++ 特性和库的扩展 |
| {{< language "Python" >}} | 2 | 使用 Python 编写的扩展 |
| {{< language "Java" >}} | 1 | 在 JVM 上运行的扩展 |


## C

{{< language "C" >}} {{< badge content="335 个扩展" color="gray" >}}


传统的 PostgreSQL 扩展开发语言

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 3860 | [`acl`](/zh/e/acl) | [`pg_acl`](/zh/e/acl) | ACL数据类型 |
| 1505 | [`address_standardizer`](/zh/e/address_standardizer) | [`postgis`](/zh/e/postgis) | 地址标准化函数。 |
| 1506 | [`address_standardizer_data_us`](/zh/e/address_standardizer_data_us) | [`postgis`](/zh/e/postgis) | 地址标准化函数：美国数据集示例 |
| 5970 | [`adminpack`](/zh/e/adminpack) | [`adminpack`](/zh/e/adminpack) | PostgreSQL 管理函数集合 |
| 2760 | [`age`](/zh/e/age) | [`age`](/zh/e/age) | Apache AGE，图数据库扩展 （Deb可用） |
| 4750 | [`aggs_for_arrays`](/zh/e/aggs_for_arrays) | [`aggs_for_arrays`](/zh/e/aggs_for_arrays) | 计算数组聚合统计值的函数包 |
| 4740 | [`aggs_for_vecs`](/zh/e/aggs_for_vecs) | [`aggs_for_vecs`](/zh/e/aggs_for_vecs) | 针对数组类型的聚合函数集合扩展 |
| 5980 | [`amcheck`](/zh/e/amcheck) | [`amcheck`](/zh/e/amcheck) | 校验关系完整性 |
| 4760 | [`arraymath`](/zh/e/arraymath) | [`pg_arraymath`](/zh/e/arraymath) | 数组逐元素数学运算符包 |
| 3560 | [`asn1oid`](/zh/e/asn1oid) | [`asn1oid`](/zh/e/asn1oid) | ASN1OID数据类型支持 |
| 7970 | [`auth_delay`](/zh/e/auth_delay) | [`auth_delay`](/zh/e/auth_delay) | 在返回认证失败前暂停一会，避免爆破 |
| 6980 | [`auto_explain`](/zh/e/auto_explain) | [`auto_explain`](/zh/e/auto_explain) | 提供一种自动记录执行计划的手段 |
| 4881 | [`autoinc`](/zh/e/autoinc) | [`autoinc`](/zh/e/autoinc) | 用于自动递增字段的函数 |
| 9300 | [`babelfishpg_common`](/zh/e/babelfishpg_common) | [`babelfishpg_common`](/zh/e/babelfishpg_common) | SQL Server 数据类型兼容扩展 |
| 9330 | [`babelfishpg_money`](/zh/e/babelfishpg_money) | [`babelfishpg_money`](/zh/e/babelfishpg_money) | SQL Server 货币数据类型兼容扩展 |
| 9320 | [`babelfishpg_tds`](/zh/e/babelfishpg_tds) | [`babelfishpg_tds`](/zh/e/babelfishpg_tds) | SQL Server TDS线缆协议兼容扩展 |
| 9310 | [`babelfishpg_tsql`](/zh/e/babelfishpg_tsql) | [`babelfishpg_tsql`](/zh/e/babelfishpg_tsql) | SQL Server SQL语法兼容性扩展 |
| 4800 | [`base36`](/zh/e/base36) | [`pg_base36`](/zh/e/base36) | Base36编码解码扩展 |
| 4810 | [`base62`](/zh/e/base62) | [`pg_base62`](/zh/e/base62) | Base62编码解码扩展 |
| 5950 | [`basebackup_to_shell`](/zh/e/basebackup_to_shell) | [`basebackup_to_shell`](/zh/e/basebackup_to_shell) | 添加一种备份到Shell终端到基础备份方式 |
| 5940 | [`basic_archive`](/zh/e/basic_archive) | [`basic_archive`](/zh/e/basic_archive) | 归档模块样例 |
| 6340 | [`bgw_replstatus`](/zh/e/bgw_replstatus) | [`bgw_replstatus`](/zh/e/bgw_replstatus) | 用于汇报本机主从状态的后台工作进程 |
| 2990 | [`bloom`](/zh/e/bloom) | [`bloom`](/zh/e/bloom) | bloom 索引-基于指纹的索引 |
| 3261 | [`bool_plperl`](/zh/e/bool_plperl) | [`plperl`](/zh/e/plperl) | 在 bool 和 plperl 之间转换 |
| 3271 | [`bool_plperlu`](/zh/e/bool_plperlu) | [`plperlu`](/zh/e/plperlu) | 在 bool 和 plperlu 之间转换 |
| 4950 | [`btree_gin`](/zh/e/btree_gin) | [`btree_gin`](/zh/e/btree_gin) | 用GIN索引常见数据类型 |
| 4940 | [`btree_gist`](/zh/e/btree_gist) | [`btree_gist`](/zh/e/btree_gist) | 用GiST索引常见数据类型 |
| 4020 | [`bzip`](/zh/e/bzip) | [`pg_bzip`](/zh/e/bzip) | BZIP压缩解压缩函数包 |
| 3920 | [`chkpass`](/zh/e/chkpass) | [`chkpass`](/zh/e/chkpass) | 数据类型：自动加密的密码 |
| 3980 | [`citext`](/zh/e/citext) | [`citext`](/zh/e/citext) | 提供大小写不敏感的字符串类型 |
| 2400 | [`citus`](/zh/e/citus) | [`citus`](/zh/e/citus) | Citus 分布式数据库 |
| 2401 | [`citus_columnar`](/zh/e/citus_columnar) | [`citus`](/zh/e/citus) | Citus 列式存储引擎 |
| 3630 | [`collection`](/zh/e/collection) | [`pg_collection`](/zh/e/collection) | 在PlPGSQL中使用的内存优化高性能集合数据结构 |
| 2410 | [`columnar`](/zh/e/columnar) | [`hydra`](/zh/e/columnar) | 开源列式存储扩展 |
| 4630 | [`count_distinct`](/zh/e/count_distinct) | [`count_distinct`](/zh/e/count_distinct) | COUNT(DISTINCT …) 聚合的替代方案 |
| 3600 | [`country`](/zh/e/country) | [`pg_country`](/zh/e/country) | 国家代码数据类型，遵循ISO 3166-1标准 |
| 7110 | [`credcheck`](/zh/e/credcheck) | [`credcheck`](/zh/e/credcheck) | 明文凭证检查器 |
| 4450 | [`cryptint`](/zh/e/cryptint) | [`cryptint`](/zh/e/cryptint) | 加密INT与BIGINT类型 |
| 3950 | [`cube`](/zh/e/cube) | [`cube`](/zh/e/cube) | 用于存储多维立方体的数据类型 |
| 3620 | [`currency`](/zh/e/currency) | [`pg_currency`](/zh/e/currency) | 使用1字节表示的货币数据类型 |
| 8630 | [`db2_fdw`](/zh/e/db2_fdw) | [`db2_fdw`](/zh/e/db2_fdw) | 提供对DB2的外部数据源包装器 |
| 8970 | [`dblink`](/zh/e/dblink) | [`dblink`](/zh/e/dblink) | 从数据库内连接到其他 PostgreSQL 数据库 |
| 3220 | [`dbt2`](/zh/e/dbt2) | [`dbt2`](/zh/e/dbt2) | OSDL-DBT-2 测试组件 |
| 4650 | [`ddsketch`](/zh/e/ddsketch) | [`ddsketch`](/zh/e/ddsketch) | 实现DDSketch数据结构，实现在线的Quantile聚合 |
| 9660 | [`decoder_raw`](/zh/e/decoder_raw) | [`decoder_raw`](/zh/e/decoder_raw) | 逻辑复制解码输出插件：RAW SQL格式 |
| 9650 | [`decoderbufs`](/zh/e/decoderbufs) | [`decoderbufs`](/zh/e/decoderbufs) | 将WAL逻辑解码为ProtocolBuffer协议的消息 |
| 4980 | [`dict_int`](/zh/e/dict_int) | [`dict_int`](/zh/e/dict_int) | 用于整数的文本搜索字典模板 |
| 4900 | [`dict_xsyn`](/zh/e/dict_xsyn) | [`dict_xsyn`](/zh/e/dict_xsyn) | 用于扩展同义词处理的文本搜索字典模板 |
| 9000 | [`documentdb`](/zh/e/documentdb) | [`documentdb`](/zh/e/documentdb) | 微软DocumentDB的API层 |
| 9010 | [`documentdb_core`](/zh/e/documentdb_core) | [`documentdb`](/zh/e/documentdb) | 微软DocumentDB的核心API层实现 |
| 9020 | [`documentdb_distributed`](/zh/e/documentdb_distributed) | [`documentdb`](/zh/e/documentdb) | DocumentDB多节点模式的API层 |
| 2450 | [`duckdb_fdw`](/zh/e/duckdb_fdw) | [`duckdb_fdw`](/zh/e/duckdb_fdw) | DuckDB 外部数据源包装器 |
| 1690 | [`earthdistance`](/zh/e/earthdistance) | [`earthdistance`](/zh/e/earthdistance) | 计算地球表面上的大圆距离 |
| 3850 | [`emailaddr`](/zh/e/emailaddr) | [`pgemailaddr`](/zh/e/emailaddr) | Email地址数据类型 |
| 4270 | [`envvar`](/zh/e/envvar) | [`envvar`](/zh/e/envvar) | 获取环境变量的函数 |
| 4720 | [`extra_window_functions`](/zh/e/extra_window_functions) | [`extra_window_functions`](/zh/e/extra_window_functions) | 额外的窗口函数 |
| 8980 | [`file_fdw`](/zh/e/file_fdw) | [`file_fdw`](/zh/e/file_fdw) | 访问外部文件的外部数据包装器 |
| 4840 | [`financial`](/zh/e/financial) | [`pg_financial`](/zh/e/financial) | 金融领域聚合函数 |
| 5230 | [`fio`](/zh/e/fio) | [`pg_fio`](/zh/e/fio) | PostgreSQL文件IO函数包 |
| 8750 | [`firebird_fdw`](/zh/e/firebird_fdw) | [`firebird_fdw`](/zh/e/firebird_fdw) | Firebird外部数据源包装器 |
| 4710 | [`first_last_agg`](/zh/e/first_last_agg) | [`first_last_agg`](/zh/e/first_last_agg) | first() 与 last() 聚合函数 |
| 4280 | [`floatfile`](/zh/e/floatfile) | [`floatfile`](/zh/e/floatfile) | 将浮点数组存储到文件中而不是堆表中 |
| 4730 | [`floatvec`](/zh/e/floatvec) | [`floatvec`](/zh/e/floatvec) | 数组类型数学运算扩展 |
| 2180 | [`fuzzystrmatch`](/zh/e/fuzzystrmatch) | [`fuzzystrmatch`](/zh/e/fuzzystrmatch) | 确定字符串之间的相似性和距离 |
| 4010 | [`gzip`](/zh/e/gzip) | [`pg_gzip`](/zh/e/gzip) | 使用SQL执行Gzip压缩与解压缩 |
| 1530 | [`h3`](/zh/e/h3) | [`pg_h3`](/zh/e/h3) | H3六边形层级索引支持 |
| 1531 | [`h3_postgis`](/zh/e/h3_postgis) | [`pg_h3`](/zh/e/h3) | H3与PostGIS集成的扩展插件 |
| 4400 | [`hashlib`](/zh/e/hashlib) | [`pg_hashlib`](/zh/e/hashlib) | 稳定哈希函数包 |
| 3750 | [`hashtypes`](/zh/e/hashtypes) | [`hashtypes`](/zh/e/hashtypes) | 包括SHA1，MD5在内的多种哈希数据类型 |
| 8740 | [`hdfs_fdw`](/zh/e/hdfs_fdw) | [`hdfs_fdw`](/zh/e/hdfs_fdw) | hdfs 外部数据包装器 |
| 3970 | [`hstore`](/zh/e/hstore) | [`hstore`](/zh/e/hstore) | 用于存储（键，值）对集合的数据类型 |
| 3021 | [`hstore_pllua`](/zh/e/hstore_pllua) | [`pllua`](/zh/e/pllua) | Lua 程序语言的Hstore适配扩展 |
| 3031 | [`hstore_plluau`](/zh/e/hstore_plluau) | [`pllua`](/zh/e/pllua) | Lua 程序语言的Hstore适配扩展（不受信任的） |
| 3262 | [`hstore_plperl`](/zh/e/hstore_plperl) | [`plperl`](/zh/e/plperl) | 在 hstore 和 plperl 之间转换适配类型 |
| 3273 | [`hstore_plperlu`](/zh/e/hstore_plperlu) | [`plperlu`](/zh/e/plperlu) | 在 hstore 和 plperlu 之间转换适配类型 |
| 3293 | [`hstore_plpython3u`](/zh/e/hstore_plpython3u) | [`plpython3u`](/zh/e/plpython3u) | 在 hstore 和 plpython3u 之间转换 |
| 4070 | [`http`](/zh/e/http) | [`pg_http`](/zh/e/http) | HTTP客户端，允许在数据库内收发HTTP请求 (supabase) |
| 2830 | [`hypopg`](/zh/e/hypopg) | [`hypopg`](/zh/e/hypopg) | 假设索引，用于创建一个虚拟索引检验执行计划 |
| 4240 | [`icu_ext`](/zh/e/icu_ext) | [`icu_ext`](/zh/e/icu_ext) | 访问ICU库提供的函数 |
| 2860 | [`imgsmlr`](/zh/e/imgsmlr) | [`imgsmlr`](/zh/e/imgsmlr) | 使用Haar小波分析计算图片相似度 |
| 4882 | [`insert_username`](/zh/e/insert_username) | [`insert_username`](/zh/e/insert_username) | 用于跟踪谁更改了表的函数 |
| 4970 | [`intagg`](/zh/e/intagg) | [`intagg`](/zh/e/intagg) | 整数聚合器和枚举器（过时） |
| 4960 | [`intarray`](/zh/e/intarray) | [`intarray`](/zh/e/intarray) | 1维整数数组的额外函数、运算符和索引支持 |
| 3820 | [`ip4r`](/zh/e/ip4r) | [`ip4r`](/zh/e/ip4r) | PostgreSQL 的 IPv4/v6 和 IPv4/v6 范围索引类型 |
| 3930 | [`isn`](/zh/e/isn) | [`isn`](/zh/e/isn) | 用于国际产品编号标准的数据类型 |
| 8530 | [`jdbc_fdw`](/zh/e/jdbc_fdw) | [`jdbc_fdw`](/zh/e/jdbc_fdw) | 访问JDBC可访问的任何外部数据源 |
| 3263 | [`jsonb_plperl`](/zh/e/jsonb_plperl) | [`plperl`](/zh/e/plperl) | 在 jsonb 和 plperl 之间转换 |
| 3272 | [`jsonb_plperlu`](/zh/e/jsonb_plperlu) | [`plperlu`](/zh/e/plperlu) | 在 jsonb 和 plperlu 之间转换 |
| 3291 | [`jsonb_plpython3u`](/zh/e/jsonb_plpython3u) | [`plpython3u`](/zh/e/plpython3u) | 在 jsonb 和 plpython3u 之间转换 |
| 2810 | [`jsquery`](/zh/e/jsquery) | [`jsquery`](/zh/e/jsquery) | 用于内省 JSONB 数据类型的查询类型 |
| 8730 | [`kafka_fdw`](/zh/e/kafka_fdw) | [`kafka_fdw`](/zh/e/kafka_fdw) | Kafka外部数据源包装器 |
| 5930 | [`lo`](/zh/e/lo) | [`lo`](/zh/e/lo) | 大对象维护 |
| 8810 | [`log_fdw`](/zh/e/log_fdw) | [`log_fdw`](/zh/e/log_fdw) | 访问PostgreSQL日志文件的FDW |
| 7140 | [`logerrors`](/zh/e/logerrors) | [`logerrors`](/zh/e/logerrors) | 用于收集日志文件中消息统计信息的函数 |
| 7150 | [`login_hook`](/zh/e/login_hook) | [`login_hook`](/zh/e/login_hook) | 在用户登陆时执行login_hook.login()函数 |
| 4620 | [`lower_quantile`](/zh/e/lower_quantile) | [`lower_quantile`](/zh/e/lower_quantile) | Lower Quantile 聚合函数 |
| 3960 | [`ltree`](/zh/e/ltree) | [`ltree`](/zh/e/ltree) | 用于表示分层树状结构的数据类型 |
| 3292 | [`ltree_plpython3u`](/zh/e/ltree_plpython3u) | [`plpython3u`](/zh/e/plpython3u) | 在 ltree 和 plpython3u 之间转换 |
| 3550 | [`md5hash`](/zh/e/md5hash) | [`md5hash`](/zh/e/md5hash) | 提供128位MD5的原生数据类型 |
| 1650 | [`mobilitydb`](/zh/e/mobilitydb) | [`mobilitydb`](/zh/e/mobilitydb) | MobilityDB地理空间投影数据管理分析平台 |
| 4883 | [`moddatetime`](/zh/e/moddatetime) | [`moddatetime`](/zh/e/moddatetime) | 跟踪最后修改时间 |
| 8700 | [`mongo_fdw`](/zh/e/mongo_fdw) | [`mongo_fdw`](/zh/e/mongo_fdw) | MongoDB 外部数据包装器 |
| 8510 | [`multicorn`](/zh/e/multicorn) | [`multicorn`](/zh/e/multicorn) | 用Python编写自定义的外部数据源包装器 |
| 8600 | [`mysql_fdw`](/zh/e/mysql_fdw) | [`mysql_fdw`](/zh/e/mysql_fdw) | MySQL外部数据包装器 |
| 7210 | [`noset`](/zh/e/noset) | [`pg_noset`](/zh/e/noset) | 阻止非超级用户使用SET/RESET设置变量 |
| 3710 | [`numeral`](/zh/e/numeral) | [`numeral`](/zh/e/numeral) | 数值类型扩展 |
| 8520 | [`odbc_fdw`](/zh/e/odbc_fdw) | [`odbc_fdw`](/zh/e/odbc_fdw) | 访问ODBC可访问的任何外部数据源 |
| 1550 | [`ogr_fdw`](/zh/e/ogr_fdw) | [`ogr_fdw`](/zh/e/ogr_fdw) | GIS 数据外部数据源包装器 |
| 5960 | [`old_snapshot`](/zh/e/old_snapshot) | [`old_snapshot`](/zh/e/old_snapshot) | 支持 old_snapshot_threshold 的实用程序 |
| 2951 | [`omni`](/zh/e/omni) | [`omnigres`](/zh/e/omni) | PostgreSQL即平台，Omnigres主扩展与加载器 |
| 2952 | [`omni_auth`](/zh/e/omni_auth) | [`omnigres`](/zh/e/omni) | Omnigres 基础会话认证管理模块 |
| 2953 | [`omni_aws`](/zh/e/omni_aws) | [`omnigres`](/zh/e/omni) | Omnigres AWS S3 API封装 |
| 2954 | [`omni_cloudevents`](/zh/e/omni_cloudevents) | [`omnigres`](/zh/e/omni) | Omnigres CloudEvents 支持 |
| 2955 | [`omni_containers`](/zh/e/omni_containers) | [`omnigres`](/zh/e/omni) | Omnigres Docker容器管理模块 |
| 2956 | [`omni_credentials`](/zh/e/omni_credentials) | [`omnigres`](/zh/e/omni) | Omnigres 应用密钥管理模块 |
| 2958 | [`omni_email`](/zh/e/omni_email) | [`omnigres`](/zh/e/omni) | Omnigres Email 框架 |
| 2959 | [`omni_http`](/zh/e/omni_http) | [`omnigres`](/zh/e/omni) | Omnigres 基本HTTP类型 |
| 2960 | [`omni_httpc`](/zh/e/omni_httpc) | [`omnigres`](/zh/e/omni) | Omnigres HTTP客户端 |
| 2961 | [`omni_httpd`](/zh/e/omni_httpd) | [`omnigres`](/zh/e/omni) | Omnigres HTTP服务器 |
| 2962 | [`omni_id`](/zh/e/omni_id) | [`omnigres`](/zh/e/omni) | Omnigres ID身份数据类型 |
| 2963 | [`omni_json`](/zh/e/omni_json) | [`omnigres`](/zh/e/omni) | Omnigres JSON工具箱 |
| 2964 | [`omni_kube`](/zh/e/omni_kube) | [`omnigres`](/zh/e/omni) | Omnigres Kubernetes集成模块 |
| 2965 | [`omni_ledger`](/zh/e/omni_ledger) | [`omnigres`](/zh/e/omni) | Omnigres 金融账本模块 |
| 2966 | [`omni_manifest`](/zh/e/omni_manifest) | [`omnigres`](/zh/e/omni) | Omnigres 包管理清单模块 |
| 2967 | [`omni_mimetypes`](/zh/e/omni_mimetypes) | [`omnigres`](/zh/e/omni) | Omnigres MIME数据类型 |
| 2968 | [`omni_os`](/zh/e/omni_os) | [`omnigres`](/zh/e/omni) | Omnigres 操作系统集成模块 |
| 2969 | [`omni_polyfill`](/zh/e/omni_polyfill) | [`omnigres`](/zh/e/omni) | Omnigres Postgres多态API |
| 2970 | [`omni_python`](/zh/e/omni_python) | [`omnigres`](/zh/e/omni) | Omnigres 第一类Python支持模块 |
| 2971 | [`omni_regex`](/zh/e/omni_regex) | [`omnigres`](/zh/e/omni) | Omnigres PCRE兼容正则表达式模块 |
| 2972 | [`omni_rest`](/zh/e/omni_rest) | [`omnigres`](/zh/e/omni) | Omnigres REST API 工具包 |
| 2973 | [`omni_schema`](/zh/e/omni_schema) | [`omnigres`](/zh/e/omni) | Omnigres 高级模式管理组件 |
| 2974 | [`omni_seq`](/zh/e/omni_seq) | [`omnigres`](/zh/e/omni) | Omnigres 分布式整型序列号 |
| 2975 | [`omni_service`](/zh/e/omni_service) | [`omnigres`](/zh/e/omni) | Omnigres 服务管理器 |
| 2976 | [`omni_session`](/zh/e/omni_session) | [`omnigres`](/zh/e/omni) | Omnigres 会话管理器 |
| 2977 | [`omni_sql`](/zh/e/omni_sql) | [`omnigres`](/zh/e/omni) | Omnigres SQL编程组件 |
| 2979 | [`omni_sqlite`](/zh/e/omni_sqlite) | [`omnigres`](/zh/e/omni) | Omnigres 嵌入的SQLite支持 |
| 2980 | [`omni_test`](/zh/e/omni_test) | [`omnigres`](/zh/e/omni) | Omnigres 测试框架 |
| 2981 | [`omni_txn`](/zh/e/omni_txn) | [`omnigres`](/zh/e/omni) | Omnigres 事务管理器模块 |
| 2982 | [`omni_types`](/zh/e/omni_types) | [`omnigres`](/zh/e/omni) | Omnigres 高级数据类型模块 |
| 2983 | [`omni_var`](/zh/e/omni_var) | [`omnigres`](/zh/e/omni) | Omnigres 局部变量模块 |
| 2984 | [`omni_vfs`](/zh/e/omni_vfs) | [`omnigres`](/zh/e/omni) | Omnigres 虚拟文件系统 |
| 2985 | [`omni_vfs_types_v1`](/zh/e/omni_vfs_types_v1) | [`omnigres`](/zh/e/omni) | Omnigres 虚拟文件系统（v1） |
| 2986 | [`omni_web`](/zh/e/omni_web) | [`omnigres`](/zh/e/omni) | Omnigres Web工具箱 |
| 2987 | [`omni_worker`](/zh/e/omni_worker) | [`omnigres`](/zh/e/omni) | Omnigres 通用Worker池 |
| 2988 | [`omni_xml`](/zh/e/omni_xml) | [`omnigres`](/zh/e/omni) | Omnigres XML工具包 |
| 2989 | [`omni_yaml`](/zh/e/omni_yaml) | [`omnigres`](/zh/e/omni) | Omnigres YAML工具包 |
| 4640 | [`omnisketch`](/zh/e/omnisketch) | [`omnisketch`](/zh/e/omnisketch) | 实现OmniSketch数据结构，实现近似摘要聚合 |
| 8610 | [`oracle_fdw`](/zh/e/oracle_fdw) | [`oracle_fdw`](/zh/e/oracle_fdw) | 提供对Oracle的外部数据源包装器 |
| 9100 | [`orafce`](/zh/e/orafce) | [`orafce`](/zh/e/orafce) | 模拟 Oracle RDBMS 的一部分函数和包的函数和运算符 |
| 2920 | [`orioledb`](/zh/e/orioledb) | [`orioledb`](/zh/e/orioledb) | OrioleDB，下一代事务处理引擎 |
| 6900 | [`pageinspect`](/zh/e/pageinspect) | [`pageinspect`](/zh/e/pageinspect) | 检查数据库页面二进制内容 |
| 7990 | [`passwordcheck`](/zh/e/passwordcheck) | [`passwordcheck`](/zh/e/passwordcheck_cracklib) | 用于强制拒绝修改弱密码的扩展 |
| 7000 | [`passwordcheck_cracklib`](/zh/e/passwordcheck_cracklib) | [`passwordcheck`](/zh/e/passwordcheck_cracklib) | 使用cracklib加固PG用户密码 |
| 1030 | [`periods`](/zh/e/periods) | [`periods`](/zh/e/periods) | 为 PERIODs 和 SYSTEM VERSIONING 提供标准 SQL 功能 |
| 4550 | [`permuteseq`](/zh/e/permuteseq) | [`permuteseq`](/zh/e/permuteseq) | 伪随机数ID置换生成器 |
| 1880 | [`pg4ml`](/zh/e/pg4ml) | [`pg4ml`](/zh/e/pg4ml) | PG4ML是一个机器学习框架 |
| 7100 | [`pg_auth_mon`](/zh/e/pg_auth_mon) | [`pg_auth_mon`](/zh/e/pg_auth_mon) | 监控每个用户的连接尝试 |
| 1100 | [`pg_background`](/zh/e/pg_background) | [`pg_background`](/zh/e/pg_background) | 在后台运行 SQL 查询 |
| 2120 | [`pg_bigm`](/zh/e/pg_bigm) | [`pg_bigm`](/zh/e/pg_bigm) | 基于二字组的多语言全文检索扩展 |
| 6930 | [`pg_buffercache`](/zh/e/pg_buffercache) | [`pg_buffercache`](/zh/e/pg_buffercache) | 检查共享缓冲区缓存 |
| 9830 | [`pg_bulkload`](/zh/e/pg_bulkload) | [`pg_bulkload`](/zh/e/pg_bulkload) | 向 PostgreSQL 中高速加载数据 |
| 5160 | [`pg_catcheck`](/zh/e/pg_catcheck) | [`pg_catcheck`](/zh/e/pg_catcheck) | 用于诊断系统目录是否损坏的工具 |
| 5220 | [`pg_cheat_funcs`](/zh/e/pg_cheat_funcs) | [`pg_cheat_funcs`](/zh/e/pg_cheat_funcs) | 一些超级实用的作弊函数 |
| 5110 | [`pg_checksums`](/zh/e/pg_checksums) | [`pg_checksums`](/zh/e/pg_checksums) | 在离线模式下激活/启用/禁用数据库集群的校验和功能 |
| 5070 | [`pg_cooldown`](/zh/e/pg_cooldown) | [`pg_cooldown`](/zh/e/pg_cooldown) | 从缓冲区中移除特定关系的页面 |
| 5210 | [`pg_crash`](/zh/e/pg_crash) | [`pg_crash`](/zh/e/pg_crash) | 向数据库进程随机发送信号模拟故障 |
| 1070 | [`pg_cron`](/zh/e/pg_cron) | [`pg_cron`](/zh/e/pg_cron) | 定时任务调度器 |
| 4090 | [`pg_curl`](/zh/e/pg_curl) | [`pg_curl`](/zh/e/pg_curl) | 封装CURL，执行各种用URL传输数据的操作 |
| 5050 | [`pg_dirtyread`](/zh/e/pg_dirtyread) | [`pg_dirtyread`](/zh/e/pg_dirtyread) | 从表中读取尚未垃圾回收的行 |
| 3830 | [`pg_duration`](/zh/e/pg_duration) | [`pg_duration`](/zh/e/pg_duration) | 用于表示时间段的强化数据类型 |
| 9820 | [`pg_fact_loader`](/zh/e/pg_fact_loader) | [`pg_fact_loader`](/zh/e/pg_fact_loader) | 在 Postgres 中构建事实表 |
| 9530 | [`pg_failover_slots`](/zh/e/pg_failover_slots) | [`pg_failover_slots`](/zh/e/pg_failover_slots) | 在Failover过程中保留复制槽 |
| 6950 | [`pg_freespacemap`](/zh/e/pg_freespacemap) | [`pg_freespacemap`](/zh/e/pg_freespacemap) | 检查自由空间映射的内容（FSM） |
| 1590 | [`pg_geohash`](/zh/e/pg_geohash) | [`pg_geohash`](/zh/e/pg_geohash) | 使用GeoHash处理空间坐标的函数包 |
| 4560 | [`pg_hashids`](/zh/e/pg_hashids) | [`pg_hashids`](/zh/e/pg_hashids) | 加盐将整型ID转为短字符串ID |
| 2820 | [`pg_hint_plan`](/zh/e/pg_hint_plan) | [`pg_hint_plan`](/zh/e/pg_hint_plan) | 添加强制指定执行计划的能力 |
| 2880 | [`pg_incremental`](/zh/e/pg_incremental) | [`pg_incremental`](/zh/e/pg_incremental) | 增量处理流式事件 |
| 2870 | [`pg_ivm`](/zh/e/pg_ivm) | [`pg_ivm`](/zh/e/pg_ivm) | 增量维护的物化视图 |
| 6890 | [`pg_logicalinspect`](/zh/e/pg_logicalinspect) | [`pg_logicalinspect`](/zh/e/pg_logicalinspect) | 检视逻辑解码组件详情 |
| 4770 | [`pg_math`](/zh/e/pg_math) | [`pg_math`](/zh/e/pg_math) | 使用GSL库的数学统计函数 |
| 4080 | [`pg_net`](/zh/e/pg_net) | [`pg_net`](/zh/e/pg_net) | 用 SQL 进行异步非阻塞HTTP/HTTPS 请求的扩展 (supabase) |
| 5200 | [`pg_orphaned`](/zh/e/pg_orphaned) | [`pg_orphaned`](/zh/e/pg_orphaned) | 处理孤儿文件的扩展插件 |
| 6880 | [`pg_overexplain`](/zh/e/pg_overexplain) | [`pg_overexplain`](/zh/e/pg_overexplain) | 允许 EXPLAIN 转储更多详细 |
| 2510 | [`pg_partman`](/zh/e/pg_partman) | [`pg_partman`](/zh/e/pg_partman) | 用于按时间或 ID 管理分区表的扩展 |
| 5890 | [`pg_prewarm`](/zh/e/pg_prewarm) | [`pg_prewarm`](/zh/e/pg_prewarm) | 预热关系数据 |
| 6320 | [`pg_proctab`](/zh/e/pg_proctab) | [`pgnodemx`](/zh/e/pgnodemx) | 通过SQL接口访问操作系统进程表 |
| 6000 | [`pg_profile`](/zh/e/pg_profile) | [`pg_profile`](/zh/e/pg_profile) | PostgreSQL 数据库负载记录与AWR报表工具 |
| 4260 | [`pg_protobuf`](/zh/e/pg_protobuf) | [`pg_protobuf`](/zh/e/pg_protobuf) | 提供Protobuf函数支持 |
| 6240 | [`pg_qualstats`](/zh/e/pg_qualstats) | [`pg_qualstats`](/zh/e/pg_qualstats) | 收集有关 quals 的统计信息的扩展 |
| 3720 | [`pg_rational`](/zh/e/pg_rational) | [`pg_rational`](/zh/e/pg_rational) | 使用BIGINT表示的有理数数据类型 |
| 4300 | [`pg_readme`](/zh/e/pg_readme) | [`pg_readme`](/zh/e/pg_readme) | 为模式与扩展生成Markdown文档 |
| 4301 | [`pg_readme_test_extension`](/zh/e/pg_readme_test_extension) | [`pg_readme`](/zh/e/pg_readme) | 为模式与扩展生成Markdown文档 |
| 5120 | [`pg_readonly`](/zh/e/pg_readonly) | [`pg_readonly`](/zh/e/pg_readonly) | 将集群设置为只读 |
| 6380 | [`pg_relusage`](/zh/e/pg_relusage) | [`pg_relusage`](/zh/e/pg_relusage) | 打印查询引用的表与列 |
| 5010 | [`pg_repack`](/zh/e/pg_repack) | [`pg_repack`](/zh/e/pg_repack) | 在线垃圾清理与表膨胀治理 |
| 5020 | [`pg_rewrite`](/zh/e/pg_rewrite) | [`pg_rewrite`](/zh/e/pg_rewrite) | 在线重写整表，不阻塞读写 |
| 3880 | [`pg_rrule`](/zh/e/pg_rrule) | [`pg_rrule`](/zh/e/pg_rrule) | 日历重复规则RRULE数据类型 |
| 5810 | [`pg_savior`](/zh/e/pg_savior) | [`pg_savior`](/zh/e/pg_savior) | 阻止不带条件的全表更新以避免意外事故 |
| 6210 | [`pg_show_plans`](/zh/e/pg_show_plans) | [`pg_show_plans`](/zh/e/pg_show_plans) | 打印所有当前正在运行查询的执行计划 |
| 1840 | [`pg_similarity`](/zh/e/pg_similarity) | [`pg_similarity`](/zh/e/pg_similarity) | 提供17种距离度量函数 |
| 7170 | [`pg_snakeoil`](/zh/e/pg_snakeoil) | [`pg_snakeoil`](/zh/e/pg_snakeoil) | PostgreSQL动态链接库反病毒功能 |
| 3590 | [`pg_sphere`](/zh/e/pg_sphere) | [`pgsphere`](/zh/e/pg_sphere) | 球面对象函数、运算符与索引支持 |
| 5040 | [`pg_squeeze`](/zh/e/pg_squeeze) | [`pg_squeeze`](/zh/e/pg_squeeze) | 从关系中删除未使用空间 |
| 6220 | [`pg_stat_kcache`](/zh/e/pg_stat_kcache) | [`pg_stat_kcache`](/zh/e/pg_stat_kcache) | 内核统计信息收集 |
| 6230 | [`pg_stat_monitor`](/zh/e/pg_stat_monitor) | [`pg_stat_monitor`](/zh/e/pg_stat_monitor) | 提供查询聚合统计、客户端信息、执行计划详细信息和直方图 |
| 6990 | [`pg_stat_statements`](/zh/e/pg_stat_statements) | [`pg_stat_statements`](/zh/e/pg_stat_statements) | 跟踪所有执行的 SQL 语句的计划和执行统计信息 |
| 9130 | [`pg_statement_rollback`](/zh/e/pg_statement_rollback) | [`pg_statement_rollback`](/zh/e/pg_statement_rollback) | 在服务端提供类似Oracle/DB2的语句级回滚能力 |
| 6250 | [`pg_store_plans`](/zh/e/pg_store_plans) | [`pg_store_plans`](/zh/e/pg_store_plans) | 跟踪所有执行的 SQL 语句的计划统计信息 |
| 2530 | [`pg_strom`](/zh/e/pg_strom) | [`pg_strom`](/zh/e/pg_strom) | 使用GPU与NVMe加速大数据处理 |
| 5990 | [`pg_surgery`](/zh/e/pg_surgery) | [`pg_surgery`](/zh/e/pg_surgery) | 对损坏的关系进行手术 |
| 1080 | [`pg_task`](/zh/e/pg_task) | [`pg_task`](/zh/e/pg_task) | 在特定时间点在后台执行SQL命令 |
| 7060 | [`pg_tde`](/zh/e/pg_tde) | [`pg_tde`](/zh/e/pg_tde) | Percona加密存储引擎 |
| 3000 | [`pg_tle`](/zh/e/pg_tle) | [`pg_tle`](/zh/e/pg_tle) | AWS 可信语言扩展 |
| 6010 | [`pg_tracing`](/zh/e/pg_tracing) | [`pg_tracing`](/zh/e/pg_tracing) | PostgreSQL分布式Tracing |
| 2190 | [`pg_trgm`](/zh/e/pg_trgm) | [`pg_trgm`](/zh/e/pg_trgm) | 文本相似度测量函数与模糊检索 |
| 4540 | [`pg_uuidv7`](/zh/e/pg_uuidv7) | [`pg_uuidv7`](/zh/e/pg_uuidv7) | UUIDv7 支持 |
| 6960 | [`pg_visibility`](/zh/e/pg_visibility) | [`pg_visibility`](/zh/e/pg_visibility) | 检查可见性图（VM）和页面级可见性信息 |
| 6270 | [`pg_wait_sampling`](/zh/e/pg_wait_sampling) | [`pg_wait_sampling`](/zh/e/pg_wait_sampling) | 基于采样的等待事件统计 |
| 6940 | [`pg_walinspect`](/zh/e/pg_walinspect) | [`pg_walinspect`](/zh/e/pg_walinspect) | 用于检查 PostgreSQL WAL 日志内容的函数 |
| 9550 | [`pgactive`](/zh/e/pgactive) | [`pgactive`](/zh/e/pgactive) | PostgreSQL多主逻辑复制 |
| 5880 | [`pgagent`](/zh/e/pgagent) | [`pgagent`](/zh/e/pgagent) | PostgreSQL任务调度工具，与PGADMIN配合使用 |
| 7080 | [`pgaudit`](/zh/e/pgaudit) | [`pgaudit`](/zh/e/pgaudit) | 提供审计功能 |
| 7090 | [`pgauditlogtofile`](/zh/e/pgauditlogtofile) | [`pgauditlogtofile`](/zh/e/pgauditlogtofile) | pgAudit 子扩展，将审计日志写入单独的文件中 |
| 5150 | [`pgautofailover`](/zh/e/pgautofailover) | [`pgautofailover`](/zh/e/pgautofailover) | PG 自动故障迁移 |
| 7980 | [`pgcrypto`](/zh/e/pgcrypto) | [`pgcrypto`](/zh/e/pgcrypto) | 实用加解密函数 |
| 7120 | [`pgcryptokey`](/zh/e/pgcryptokey) | [`pgcryptokey`](/zh/e/pgcryptokey) | PG密钥管理 |
| 7180 | [`pgextwlist`](/zh/e/pgextwlist) | [`pgextwlist`](/zh/e/pgextwlist) | PostgreSQL扩展白名单功能 |
| 5060 | [`pgfincore`](/zh/e/pgfincore) | [`pgfincore`](/zh/e/pgfincore) | 检查和管理操作系统缓冲区缓存 |
| 4150 | [`pgjq`](/zh/e/pgjq) | [`pgjq`](/zh/e/pgjq) | 在Postgres中使用jq查询JSON |
| 9520 | [`pgl_ddl_deploy`](/zh/e/pgl_ddl_deploy) | [`pgl_ddl_deploy`](/zh/e/pgl_ddl_deploy) | 使用 pglogical 执行自动 DDL 部署 |
| 9500 | [`pglogical`](/zh/e/pglogical) | [`pglogical`](/zh/e/pglogical) | PostgreSQL逻辑复制：三方扩展实现 |
| 9501 | [`pglogical_origin`](/zh/e/pglogical_origin) | [`pglogical`](/zh/e/pglogical) | 用于从 Postgres 9.4 升级时的兼容性虚拟扩展 |
| 9510 | [`pglogical_ticker`](/zh/e/pglogical_ticker) | [`pglogical_ticker`](/zh/e/pglogical_ticker) | pglogical复制延迟以秒计的精确视图 |
| 9410 | [`pgmemcache`](/zh/e/pgmemcache) | [`pgmemcache`](/zh/e/pgmemcache) | 为PG提供memcached兼容接口 |
| 6350 | [`pgmeminfo`](/zh/e/pgmeminfo) | [`pgmeminfo`](/zh/e/pgmeminfo) | 显示内存使用情况 |
| 3700 | [`pgmp`](/zh/e/pgmp) | [`pgmp`](/zh/e/pgmp) | 多精度算术扩展 |
| 6310 | [`pgnodemx`](/zh/e/pgnodemx) | [`pgnodemx`](/zh/e/pgnodemx) | 使用SQL查询获取操作系统指标 |
| 9980 | [`pgoutput`](/zh/e/pgoutput) | [`pgoutput`](/zh/e/pgoutput) | PG内置的逻辑解码输出插件 |
| 4230 | [`pgpcre`](/zh/e/pgpcre) | [`pgpcre`](/zh/e/pgpcre) | PCRE/Perl风格的正则表达式支持 |
| 3530 | [`pgpdf`](/zh/e/pgpdf) | [`pgpdf`](/zh/e/pgpdf) | PDF数据类型，管理函数与全文检索 |
| 5900 | [`pgpool_adm`](/zh/e/pgpool_adm) | [`pgpool`](/zh/e/pgpool_adm) | PGPool 管理函数 |
| 5910 | [`pgpool_recovery`](/zh/e/pgpool_recovery) | [`pgpool`](/zh/e/pgpool_adm) | PGPool辅助扩展，从v4.3提供的恢复函数 |
| 5920 | [`pgpool_regclass`](/zh/e/pgpool_regclass) | [`pgpool`](/zh/e/pgpool_adm) | PGPool辅助扩展，RegClass替代 |
| 2910 | [`pgq`](/zh/e/pgq) | [`pgq`](/zh/e/pgq) | 通用队列的PG实现 |
| 4250 | [`pgqr`](/zh/e/pgqr) | [`pgqr`](/zh/e/pgqr) | 从数据库中直接生成QR二维码 |
| 2110 | [`pgroonga`](/zh/e/pgroonga) | [`pgroonga`](/zh/e/pgroonga) | 使用Groonga，面向所有语言的高速全文检索平台 |
| 2111 | [`pgroonga_database`](/zh/e/pgroonga_database) | [`pgroonga`](/zh/e/pgroonga) | PGGroonga 数据库管理模块 |
| 6910 | [`pgrowlocks`](/zh/e/pgrowlocks) | [`pgrowlocks`](/zh/e/pgrowlocks) | 显示行级锁信息 |
| 6280 | [`pgsentinel`](/zh/e/pgsentinel) | [`pgsentinel`](/zh/e/pgsentinel) | 活跃会话历史 |
| 7020 | [`pgsodium`](/zh/e/pgsodium) | [`pgsodium`](/zh/e/pgsodium) | 表数据加密存储 TDE |
| 8540 | [`pgspider_ext`](/zh/e/pgspider_ext) | [`pgspider_ext`](/zh/e/pgspider_ext) | 使用多种FDW访问远程数据库服务器 |
| 6970 | [`pgstattuple`](/zh/e/pgstattuple) | [`pgstattuple`](/zh/e/pgstattuple) | 显示元组级统计信息 |
| 3200 | [`pgtap`](/zh/e/pgtap) | [`pgtap`](/zh/e/pgtap) | PostgreSQL单元测试框架 |
| 9110 | [`pgtt`](/zh/e/pgtt) | [`pgtt`](/zh/e/pgtt) | 类似Oracle的全局临时表功能 |
| 4460 | [`pguecc`](/zh/e/pguecc) | [`pg_ecdsa`](/zh/e/pguecc) | PostgreSQL的uECC绑定，椭圆曲线加解密函数包 |
| 2850 | [`plan_filter`](/zh/e/plan_filter) | [`pg_plan_filter`](/zh/e/plan_filter) | 使用执行计划代价过滤阻止特定查询语句 |
| 3050 | [`pldbgapi`](/zh/e/pldbgapi) | [`pldebugger`](/zh/e/pldbgapi) | 用于调试 PL/pgSQL 函数的服务器端支持 |
| 3020 | [`pllua`](/zh/e/pllua) | [`pllua`](/zh/e/pllua) | Lua 程序语言 |
| 3030 | [`plluau`](/zh/e/plluau) | [`pllua`](/zh/e/pllua) | Lua 程序语言（不受信任的） |
| 3260 | [`plperl`](/zh/e/plperl) | [`plperl`](/zh/e/plperl) | PL/Perl 存储过程语言 |
| 3270 | [`plperlu`](/zh/e/plperlu) | [`plperlu`](/zh/e/plperlu) | PL/PerlU 存储过程语言（未受信/高权限） |
| 3280 | [`plpgsql`](/zh/e/plpgsql) | [`plpgsql`](/zh/e/plpgsql) | PL/pgSQL 程序设计语言 |
| 3060 | [`plpgsql_check`](/zh/e/plpgsql_check) | [`plpgsql_check`](/zh/e/plpgsql_check) | 对 plpgsql 函数进行扩展检查 |
| 3070 | [`plprofiler`](/zh/e/plprofiler) | [`plprofiler`](/zh/e/plprofiler) | 剖析 PL/pgSQL 函数 |
| 2520 | [`plproxy`](/zh/e/plproxy) | [`plproxy`](/zh/e/plproxy) | 作为过程语言实现的数据库分区 |
| 3290 | [`plpython3u`](/zh/e/plpython3u) | [`plpython3u`](/zh/e/plpython3u) | PL/Python3 存储过程语言（未受信/高权限） |
| 3100 | [`plr`](/zh/e/plr) | [`plr`](/zh/e/plr) | 从数据库中加载R语言解释器并执行R脚本 |
| 3080 | [`plsh`](/zh/e/plsh) | [`plsh`](/zh/e/plsh) | PL/sh 程序语言 |
| 3240 | [`pltcl`](/zh/e/pltcl) | [`pltcl`](/zh/e/pltcl) | PL/TCL 存储过程语言 |
| 3250 | [`pltclu`](/zh/e/pltclu) | [`pltcl`](/zh/e/pltcl) | PL/TCL 存储过程语言（未受信/高权限） |
| 1520 | [`pointcloud`](/zh/e/pointcloud) | [`pointcloud`](/zh/e/pointcloud) | 提供激光雷达点云数据类型支持 |
| 1521 | [`pointcloud_postgis`](/zh/e/pointcloud_postgis) | [`pointcloud`](/zh/e/pointcloud) | 将激光雷达点云与PostGIS几何类型相集成 |
| 1500 | [`postgis`](/zh/e/postgis) | [`postgis`](/zh/e/postgis) | PostGIS 几何和地理空间扩展 |
| 1502 | [`postgis_raster`](/zh/e/postgis_raster) | [`postgis`](/zh/e/postgis) | PostGIS 光栅类型和函数 |
| 1503 | [`postgis_sfcgal`](/zh/e/postgis_sfcgal) | [`postgis`](/zh/e/postgis) | PostGIS SFCGAL 函数 |
| 1504 | [`postgis_tiger_geocoder`](/zh/e/postgis_tiger_geocoder) | [`postgis`](/zh/e/postgis) | PostGIS tiger 地理编码器和反向地理编码器 |
| 1501 | [`postgis_topology`](/zh/e/postgis_topology) | [`postgis`](/zh/e/postgis) | PostGIS 拓扑空间类型和函数 |
| 8990 | [`postgres_fdw`](/zh/e/postgres_fdw) | [`postgres_fdw`](/zh/e/postgres_fdw) | 用于远程 PostgreSQL 服务器的外部数据包装器 |
| 5170 | [`pre_prepare`](/zh/e/pre_prepare) | [`preprepare`](/zh/e/pre_prepare) | 在服务端预先准备好PreparedStatement备用 |
| 3500 | [`prefix`](/zh/e/prefix) | [`pg_prefix`](/zh/e/prefix) | 前缀树数据类型 |
| 5090 | [`prioritize`](/zh/e/prioritize) | [`pg_prioritize`](/zh/e/prioritize) | 获取和设置 PostgreSQL 后端的优先级 |
| 1540 | [`q3c`](/zh/e/q3c) | [`q3c`](/zh/e/q3c) | Q3C天空索引插件 |
| 4610 | [`quantile`](/zh/e/quantile) | [`quantile`](/zh/e/quantile) | Quantile聚合函数 |
| 4780 | [`random`](/zh/e/random) | [`pg_random`](/zh/e/random) | 随机数生成器 |
| 8720 | [`redis`](/zh/e/redis) | [`pg_redis_pubsub`](/zh/e/redis) | 从PG向Redis发送Pub/Sub消息 |
| 8710 | [`redis_fdw`](/zh/e/redis_fdw) | [`redis_fdw`](/zh/e/redis_fdw) | 查询外部Redis数据源 |
| 4880 | [`refint`](/zh/e/refint) | [`refint`](/zh/e/refint) | 实现引用完整性的函数 |
| 9710 | [`repmgr`](/zh/e/repmgr) | [`repmgr`](/zh/e/repmgr) | PostgreSQL复制管理组件 |
| 3570 | [`roaringbitmap`](/zh/e/roaringbitmap) | [`roaringbitmap`](/zh/e/roaringbitmap) | 支持RoaringBitmap数据类型 |
| 2780 | [`rum`](/zh/e/rum) | [`rum`](/zh/e/rum) | RUM 索引访问方法 |
| 5820 | [`safeupdate`](/zh/e/safeupdate) | [`safeupdate`](/zh/e/safeupdate) | 强制在 UPDATE 和 DELETE 时提供 Where 条件 |
| 3940 | [`seg`](/zh/e/seg) | [`seg`](/zh/e/seg) | 表示线段或浮点间隔的数据类型 |
| 3510 | [`semver`](/zh/e/semver) | [`pg_semver`](/zh/e/semver) | 语义版本号数据类型 |
| 7960 | [`sepgsql`](/zh/e/sepgsql) | [`sepgsql`](/zh/e/sepgsql) | 基于SELinux标签的强制访问控制 |
| 4570 | [`sequential_uuids`](/zh/e/sequential_uuids) | [`sequential_uuids`](/zh/e/sequential_uuids) | 生成连续生成的UUID |
| 9120 | [`session_variable`](/zh/e/session_variable) | [`session_variable`](/zh/e/session_variable) | Oracle兼容的会话变量/常量操作函数 |
| 7160 | [`set_user`](/zh/e/set_user) | [`set_user`](/zh/e/set_user) | 增加了日志记录的 SET ROLE |
| 4440 | [`shacrypt`](/zh/e/shacrypt) | [`shacrypt`](/zh/e/shacrypt) | 实现SHA256-CRYPT与SHA512-CRYPT密钥加密算法 |
| 1850 | [`smlar`](/zh/e/smlar) | [`smlar`](/zh/e/smlar) | 高效的相似度搜索函数 |
| 9400 | [`spat`](/zh/e/spat) | [`spat`](/zh/e/spat) | 在PG中嵌入Redis风格的内存数据库 |
| 8640 | [`sqlite_fdw`](/zh/e/sqlite_fdw) | [`sqlite_fdw`](/zh/e/sqlite_fdw) | SQLite 外部数据包装器 |
| 6920 | [`sslinfo`](/zh/e/sslinfo) | [`sslinfo`](/zh/e/sslinfo) | 关于 SSL 证书的信息 |
| 7200 | [`sslutils`](/zh/e/sslutils) | [`sslutils`](/zh/e/sslutils) | 使用SQL管理SSL证书 |
| 7030 | [`supabase_vault`](/zh/e/supabase_vault) | [`pg_vault`](/zh/e/supabase_vault) | 在 Vault 中存储加密凭证的扩展 (supabase) |
| 7010 | [`supautils`](/zh/e/supautils) | [`supautils`](/zh/e/supautils) | 用于在云环境中确保数据库集群的安全 |
| 6290 | [`system_stats`](/zh/e/system_stats) | [`system_stats`](/zh/e/system_stats) | PostgreSQL 的系统统计函数 |
| 5840 | [`table_log`](/zh/e/table_log) | [`table_log`](/zh/e/table_log) | 记录某张表的修改日志并做表/行级时间点恢复 |
| 2590 | [`tablefunc`](/zh/e/tablefunc) | [`tablefunc`](/zh/e/tablefunc) | 交叉表函数 |
| 4920 | [`tcn`](/zh/e/tcn) | [`tcn`](/zh/e/tcn) | 用触发器通知变更 |
| 4700 | [`tdigest`](/zh/e/tdigest) | [`tdigest`](/zh/e/tdigest) | tdigest 聚合函数 |
| 8620 | [`tds_fdw`](/zh/e/tds_fdw) | [`tds_fdw`](/zh/e/tds_fdw) | TDS 数据库（Sybase/SQL Server）外部数据包装器 |
| 1040 | [`temporal_tables`](/zh/e/temporal_tables) | [`temporal_tables`](/zh/e/temporal_tables) | 时态表功能支持 |
| 9970 | [`test_decoding`](/zh/e/test_decoding) | [`test_decoding`](/zh/e/test_decoding) | 基于SQL的WAL逻辑解码样例 |
| 1000 | [`timescaledb`](/zh/e/timescaledb) | [`timescaledb`](/zh/e/timescaledb) | 时序数据库扩展插件 |
| 3890 | [`timestamp9`](/zh/e/timestamp9) | [`timestamp9`](/zh/e/timestamp9) | 纳秒分辨率时间戳 |
| 6360 | [`toastinfo`](/zh/e/toastinfo) | [`toastinfo`](/zh/e/toastinfo) | 显示TOAST字段的详细信息 |
| 4600 | [`topn`](/zh/e/topn) | [`topn`](/zh/e/topn) | top-n JSONB 的类型 |
| 4910 | [`tsm_system_rows`](/zh/e/tsm_system_rows) | [`tsm_system_rows`](/zh/e/tsm_system_rows) | 接受行数限制的 TABLESAMPLE 方法 |
| 4890 | [`tsm_system_time`](/zh/e/tsm_system_time) | [`tsm_system_time`](/zh/e/tsm_system_time) | 接受毫秒数限制的 TABLESAMPLE 方法 |
| 3730 | [`uint`](/zh/e/uint) | [`pguint`](/zh/e/uint) | 无符号整型数据类型 |
| 3740 | [`uint128`](/zh/e/uint128) | [`pg_uint128`](/zh/e/uint128) | 原生128位无符号整型数据类型 |
| 4990 | [`unaccent`](/zh/e/unaccent) | [`unaccent`](/zh/e/unaccent) | 删除重音的文本搜索字典 |
| 3520 | [`unit`](/zh/e/unit) | [`pgunit`](/zh/e/unit) | SI 国标单位扩展 |
| 3840 | [`uri`](/zh/e/uri) | [`pg_uri`](/zh/e/uri) | URI数据类型 |
| 4190 | [`url_encode`](/zh/e/url_encode) | [`url_encode`](/zh/e/url_encode) | 提供URL编码解码函数 |
| 4930 | [`uuid-ossp`](/zh/e/uuid-ossp) | [`uuid-ossp`](/zh/e/uuid-ossp) | 生成通用唯一标识符（UUIDs） |
| 4660 | [`vasco`](/zh/e/vasco) | [`vasco`](/zh/e/vasco) | 使用MIC发现数据中隐含的关联 |
| 1800 | [`vector`](/zh/e/vector) | [`pgvector`](/zh/e/vector) | 向量数据类型和 ivfflat / hnsw 访问方法 |
| 9630 | [`wal2json`](/zh/e/wal2json) | [`wal2json`](/zh/e/wal2json) | 用逻辑解码捕获 JSON 格式的 CDC 变更 |
| 9640 | [`wal2mongo`](/zh/e/wal2mongo) | [`wal2mongo`](/zh/e/wal2mongo) | 使用逻辑解码捕获MongoDB JSON格式的CDC变更 |
| 4670 | [`xicor`](/zh/e/xicor) | [`pgxicor`](/zh/e/xicor) | 在PG中计算XI相关系数 |
| 3990 | [`xml2`](/zh/e/xml2) | [`xml2`](/zh/e/xml2) | XPath 查询和 XSLT |
| 4430 | [`xxhash`](/zh/e/xxhash) | [`pg_xxhash`](/zh/e/xxhash) | xxhash哈希函数包 |
| 2130 | [`zhparser`](/zh/e/zhparser) | [`zhparser`](/zh/e/zhparser) | 中文分词，全文搜索解析器 |
| 4030 | [`zstd`](/zh/e/zstd) | [`pg_zstd`](/zh/e/zstd) | ZSTD压缩解压缩函数包 |

## SQL

{{< language "SQL" >}} {{< badge content="37 个扩展" color="gray" >}}


纯 SQL 扩展和函数

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 8800 | [`aws_s3`](/zh/e/aws_s3) | [`aws_s3`](/zh/e/aws_s3) | 从S3导入导出数据的外部数据源包装器 |
| 4320 | [`data_historization`](/zh/e/data_historization) | [`data_historization`](/zh/e/data_historization) | 用SQL将数据变更历史保存到分区表中 |
| 9540 | [`db_migrator`](/zh/e/db_migrator) | [`db_migrator`](/zh/e/db_migrator) | 使用FDW从其他DBMS迁移到PostgreSQL |
| 4310 | [`ddl_historization`](/zh/e/ddl_historization) | [`ddl_historization`](/zh/e/ddl_historization) | 用SQL将所有DDL变更写入到数据库表中 |
| 5080 | [`ddlx`](/zh/e/ddlx) | [`pg_ddlx`](/zh/e/ddlx) | 提取数据库对象的DDL |
| 3870 | [`debversion`](/zh/e/debversion) | [`debversion`](/zh/e/debversion) | Debian版本号数据类型 |
| 1050 | [`emaj`](/zh/e/emaj) | [`emaj`](/zh/e/emaj) | 让数据库的子集具有细粒度日志和时间旅行功能 |
| 1560 | [`geoip`](/zh/e/geoip) | [`geoip`](/zh/e/geoip) | IP 地理位置扩展（围绕 MaxMind GeoLite 数据集的包装器） |
| 2840 | [`index_advisor`](/zh/e/index_advisor) | [`index_advisor`](/zh/e/index_advisor) | 查询索引建议器 |
| 3611 | [`l10n_table_dependent_extension`](/zh/e/l10n_table_dependent_extension) | [`pg_xenophile`](/zh/e/pg_xenophile) | PostgreSQL l10n 工具包 |
| 6300 | [`meta`](/zh/e/meta) | [`pg_meta`](/zh/e/meta) | 标准化，更友好的PostgreSQL系统目录视图 |
| 9700 | [`mimeo`](/zh/e/mimeo) | [`mimeo`](/zh/e/mimeo) | 在PostgreSQL实例间进行表级复制 |
| 6800 | [`pagevis`](/zh/e/pagevis) | [`pagevis`](/zh/e/pagevis) | 使用ASCII字符可视化数据库物理页面布局 |
| 7190 | [`pg_auditor`](/zh/e/pg_auditor) | [`pg_auditor`](/zh/e/pg_auditor) | 审计数据变更并提供闪回能力 |
| 9260 | [`pg_dbms_job`](/zh/e/pg_dbms_job) | [`pg_dbms_job`](/zh/e/pg_dbms_job) | 添加 Oracle DBMS_JOB 兼容性支持的扩展 |
| 9250 | [`pg_dbms_lock`](/zh/e/pg_dbms_lock) | [`pg_dbms_lock`](/zh/e/pg_dbms_lock) | 为PG添加对 Oracle DBMS_LOCK 的完整兼容性支持 |
| 9240 | [`pg_dbms_metadata`](/zh/e/pg_dbms_metadata) | [`pg_dbms_metadata`](/zh/e/pg_dbms_metadata) | 添加 Oracle DBMS_METADATA 兼容性支持的扩展 |
| 5830 | [`pg_drop_events`](/zh/e/pg_drop_events) | [`pg_drop_events`](/zh/e/pg_drop_events) | 记录删表删列删视图的事务号，辅助PITR确定时间点 |
| 4220 | [`pg_extra_time`](/zh/e/pg_extra_time) | [`pg_extra_time`](/zh/e/pg_extra_time) | 一些关于日期与时间的扩展函数 |
| 2500 | [`pg_fkpart`](/zh/e/pg_fkpart) | [`pg_fkpart`](/zh/e/pg_fkpart) | 按外键实用程序进行表分区的扩展 |
| 4180 | [`pg_html5_email_address`](/zh/e/pg_html5_email_address) | [`pg_html5_email_address`](/zh/e/pg_html5_email_address) | 验证Email是否符合HTML5规范的扩展 |
| 7130 | [`pg_jobmon`](/zh/e/pg_jobmon) | [`pg_jobmon`](/zh/e/pg_jobmon) | 记录和监控函数 |
| 5140 | [`pg_permissions`](/zh/e/pg_permissions) | [`pg_permissions`](/zh/e/pg_permissions) | 查看对象权限并将其与期望状态进行比较 |
| 6330 | [`pg_sqlog`](/zh/e/pg_sqlog) | [`pg_sqlog`](/zh/e/pg_sqlog) | 提供访问PostgreSQL日志的SQL接口 |
| 6260 | [`pg_track_settings`](/zh/e/pg_track_settings) | [`pg_track_settings`](/zh/e/pg_track_settings) | 跟踪设置更改 |
| 5180 | [`pg_upless`](/zh/e/pg_upless) | [`pg_upless`](/zh/e/pg_upless) | 检测表上的无用UPDATE |
| 3610 | [`pg_xenophile`](/zh/e/pg_xenophile) | [`pg_xenophile`](/zh/e/pg_xenophile) | PostgreSQL i8n与l10n工具包 |
| 8650 | [`pgbouncer_fdw`](/zh/e/pgbouncer_fdw) | [`pgbouncer_fdw`](/zh/e/pgbouncer_fdw) | 用SQL查询pgbouncer统计信息，并执行pgbouncer命令 |
| 5190 | [`pgcozy`](/zh/e/pgcozy) | [`pgcozy`](/zh/e/pgcozy) | 根据先前的pg_buffercache快照预热内存缓冲区 |
| 3580 | [`pgfaceting`](/zh/e/pgfaceting) | [`pgfaceting`](/zh/e/pgfaceting) | 使用倒排索引的高速切面查询 |
| 4160 | [`pgjwt`](/zh/e/pgjwt) | [`pgjwt`](/zh/e/pgjwt) | JSON Web Token API 的PG实现 (supabase) |
| 2900 | [`pgmq`](/zh/e/pgmq) | [`pgmq`](/zh/e/pgmq) | 基于Postgres实现类似AWS SQS/RSMQ的消息队列 |
| 4200 | [`pgsql_tweaks`](/zh/e/pgsql_tweaks) | [`pgsql_tweaks`](/zh/e/pgsql_tweaks) | 一些日常会用到的便利函数与视图 |
| 4330 | [`schedoc`](/zh/e/schedoc) | [`pg_schedoc`](/zh/e/schedoc) | 在Django与DBT之间通过注释文档交换元数据 |
| 4470 | [`sparql`](/zh/e/sparql) | [`pgsparql`](/zh/e/sparql) | 使用SQL查询SPARQL数据源 |
| 1060 | [`table_version`](/zh/e/table_version) | [`table_version`](/zh/e/table_version) | PostgreSQL 版本控制表扩展 |
| 1020 | [`timeseries`](/zh/e/timeseries) | [`pg_timeseries`](/zh/e/timeseries) | 时序数据API封装 |

## Rust

{{< language "Rust" >}} {{< badge content="33 个扩展" color="gray" >}}


使用 pgrx 框架用 Rust 编写的扩展

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 7050 | [`anon`](/zh/e/anon) | [`pg_anon`](/zh/e/anon) | 数据匿名化处理工具 |
| 4850 | [`convert`](/zh/e/convert) | [`pg_convert`](/zh/e/convert) | 用于空间里程等的公英制转换函数 |
| 6370 | [`explain_ui`](/zh/e/explain_ui) | [`pg_explain_ui`](/zh/e/explain_ui) | 快速跳转至PEV查阅可视化执行计划 |
| 2420 | [`pg_analytics`](/zh/e/pg_analytics) | [`pg_analytics`](/zh/e/pg_analytics) | 由 DuckDB 驱动的数据分析引擎 |
| 4830 | [`pg_base58`](/zh/e/pg_base58) | [`pg_base58`](/zh/e/pg_base58) | Base58 编码/解码函数 |
| 2140 | [`pg_bestmatch`](/zh/e/pg_bestmatch) | [`pg_bestmatch`](/zh/e/pg_bestmatch) | 在数据库内生成BM25稀疏向量 |
| 2930 | [`pg_cardano`](/zh/e/pg_cardano) | [`pg_cardano`](/zh/e/pg_cardano) | Cardano相关工具包：加密函数，地址编解码，区块链处理 |
| 2790 | [`pg_graphql`](/zh/e/pg_graphql) | [`pg_graphql`](/zh/e/pg_graphql) | PG内的GraphQL支持 |
| 4500 | [`pg_idkit`](/zh/e/pg_idkit) | [`pg_idkit`](/zh/e/pg_idkit) | 生成各式各样的唯一标识符：UUIDv6, ULID, KSUID |
| 2800 | [`pg_jsonschema`](/zh/e/pg_jsonschema) | [`pg_jsonschema`](/zh/e/pg_jsonschema) | 提供JSON Schema校验能力 |
| 1090 | [`pg_later`](/zh/e/pg_later) | [`pg_later`](/zh/e/pg_later) | 执行查询，并在稍后异步获取查询结果 |
| 2460 | [`pg_parquet`](/zh/e/pg_parquet) | [`pg_parquet`](/zh/e/pg_parquet) | 在PostgreSQL与本地/S3中的Parquet文件复制数据 |
| 1570 | [`pg_polyline`](/zh/e/pg_polyline) | [`pg_polyline`](/zh/e/pg_polyline) | Google快速Polyline编码解码扩展 |
| 4290 | [`pg_render`](/zh/e/pg_render) | [`pg_render`](/zh/e/pg_render) | 使用SQL渲染HTML页面 |
| 2100 | [`pg_search`](/zh/e/pg_search) | [`pg_search`](/zh/e/pg_search) | ParadeDB BM25算法全文检索插件，ES全文检索 |
| 7040 | [`pg_session_jwt`](/zh/e/pg_session_jwt) | [`pg_session_jwt`](/zh/e/pg_session_jwt) | 使用JWT进行会话认证 |
| 4170 | [`pg_smtp_client`](/zh/e/pg_smtp_client) | [`pg_smtp_client`](/zh/e/pg_smtp_client) | 使用SMTP从PostgreSQL内发送邮件的客户端扩展 |
| 1860 | [`pg_summarize`](/zh/e/pg_summarize) | [`pg_summarize`](/zh/e/pg_summarize) | 使用LLM对文本字段进行总结 |
| 1870 | [`pg_tiktoken`](/zh/e/pg_tiktoken) | [`pg_tiktoken`](/zh/e/pg_tiktoken) | 在PostgreSQL中计算OpenAI使用的Token数 |
| 2160 | [`pg_tokenizer`](/zh/e/pg_tokenizer) | [`pg_tokenizer`](/zh/e/pg_tokenizer) | 用于全文检索的分词器 |
| 5130 | [`pgdd`](/zh/e/pgdd) | [`pgdd`](/zh/e/pgdd) | 提供通过标准SQL查询数据库目录集簇的能力 |
| 3540 | [`pglite_fusion`](/zh/e/pglite_fusion) | [`pglite_fusion`](/zh/e/pglite_fusion) | 在PG表中嵌入SQLite数据库作为数据类型 |
| 1890 | [`pgml`](/zh/e/pgml) | [`pgml`](/zh/e/pgml) | PostgresML：用SQL运行机器学习算法并训练模型 |
| 7070 | [`pgsmcrypto`](/zh/e/pgsmcrypto) | [`pgsmcrypto`](/zh/e/pgsmcrypto) | 为PostgreSQL提供商密算法支持：SM2,SM3,SM4 |
| 4510 | [`pgx_ulid`](/zh/e/pgx_ulid) | [`pgx_ulid`](/zh/e/pgx_ulid) | ULID数据类型与函数 |
| 3040 | [`plprql`](/zh/e/plprql) | [`plprql`](/zh/e/plprql) | 在PostgreSQL使用PRQL——管线式关系查询语言 |
| 1010 | [`timescaledb_toolkit`](/zh/e/timescaledb_toolkit) | [`timescaledb_toolkit`](/zh/e/timescaledb_toolkit) | 超表分析查询，时间序列流式处理，以及其他SQL工具 |
| 1680 | [`tzf`](/zh/e/tzf) | [`pg_tzf`](/zh/e/tzf) | 快速根据GPS经纬度坐标查找时区 |
| 1810 | [`vchord`](/zh/e/vchord) | [`vchord`](/zh/e/vchord) | 使用Rust重写的高性能向量扩展 |
| 2150 | [`vchord_bm25`](/zh/e/vchord_bm25) | [`vchord_bm25`](/zh/e/vchord_bm25) | BM25排序算法 |
| 1830 | [`vectorize`](/zh/e/vectorize) | [`pg_vectorize`](/zh/e/vectorize) | 在PostgreSQL中封装RAG向量检索服务 |
| 1820 | [`vectorscale`](/zh/e/vectorscale) | [`pgvectorscale`](/zh/e/vectorscale) | 使用DiskANN算法对向量进行高效索引 |
| 8500 | [`wrappers`](/zh/e/wrappers) | [`wrappers`](/zh/e/wrappers) | Supabase提供的外部数据源包装器捆绑包 |

## Data

{{< language "Data" >}} {{< badge content="10 个扩展" color="gray" >}}


仅包含数据的扩展

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 2170 | [`hunspell_cs_cz`](/zh/e/hunspell_cs_cz) | [`hunspell_cs_cz`](/zh/e/hunspell_cs_cz) | Hunspell捷克语全文检索词典 |
| 2171 | [`hunspell_de_de`](/zh/e/hunspell_de_de) | [`hunspell_de_de`](/zh/e/hunspell_de_de) | Hunspell德语全文检索词典 |
| 2172 | [`hunspell_en_us`](/zh/e/hunspell_en_us) | [`hunspell_en_us`](/zh/e/hunspell_en_us) | Hunspell英语全文检索词典 |
| 2173 | [`hunspell_fr`](/zh/e/hunspell_fr) | [`hunspell_fr`](/zh/e/hunspell_fr) | Hunspell法语全文检索词典 |
| 2174 | [`hunspell_ne_np`](/zh/e/hunspell_ne_np) | [`hunspell_ne_np`](/zh/e/hunspell_ne_np) | Hunspell尼泊尔语全文检索词典 |
| 2175 | [`hunspell_nl_nl`](/zh/e/hunspell_nl_nl) | [`hunspell_nl_nl`](/zh/e/hunspell_nl_nl) | Hunspell荷兰语全文检索词典 |
| 2176 | [`hunspell_nn_no`](/zh/e/hunspell_nn_no) | [`hunspell_nn_no`](/zh/e/hunspell_nn_no) | Hunspell挪威语全文检索词典 |
| 2177 | [`hunspell_pt_pt`](/zh/e/hunspell_pt_pt) | [`hunspell_pt_pt`](/zh/e/hunspell_pt_pt) | Hunspell葡萄牙语全文检索词典 |
| 2178 | [`hunspell_ru_ru`](/zh/e/hunspell_ru_ru) | [`hunspell_ru_ru`](/zh/e/hunspell_ru_ru) | Hunspell俄语全文检索词典 |
| 2179 | [`hunspell_ru_ru_aot`](/zh/e/hunspell_ru_ru_aot) | [`hunspell_ru_ru_aot`](/zh/e/hunspell_ru_ru_aot) | Hunspell俄语全文检索词典（来自AOT.ru小组） |

## C++

{{< language "C++" >}} {{< badge content="6 个扩展" color="gray" >}}


使用 C++ 特性和库的扩展

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 2770 | [`hll`](/zh/e/hll) | [`hll`](/zh/e/hll) | hyperloglog 数据类型 |
| 2430 | [`pg_duckdb`](/zh/e/pg_duckdb) | [`pg_duckdb`](/zh/e/pg_duckdb) | 在PostgreSQL中的嵌入式DuckDB扩展 |
| 2440 | [`pg_mooncake`](/zh/e/pg_mooncake) | [`pg_mooncake`](/zh/e/pg_mooncake) | PostgreSQL列式存储表 |
| 1510 | [`pgrouting`](/zh/e/pgrouting) | [`pgrouting`](/zh/e/pgrouting) | 提供寻路能力 |
| 3010 | [`plv8`](/zh/e/plv8) | [`plv8`](/zh/e/plv8) | PL/JavaScript (v8) 可信过程程序语言 |
| 2940 | [`rdkit`](/zh/e/rdkit) | [`rdkit`](/zh/e/rdkit) | 在PostgreSQL化学领域数据管理功能 |

## Python

{{< language "Python" >}} {{< badge content="2 个扩展" color="gray" >}}


使用 Python 编写的扩展

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 3210 | [`faker`](/zh/e/faker) | [`faker`](/zh/e/faker) | 插入生成的测试伪造数据，Python库的包装 |
| 6810 | [`powa`](/zh/e/powa) | [`powa`](/zh/e/powa) | PostgreSQL 工作负载分析器-核心 |

## Java

{{< language "Java" >}} {{< badge content="1 个扩展" color="gray" >}}


在 JVM 上运行的扩展

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 3090 | [`pljava`](/zh/e/pljava) | [`pljava`](/zh/e/pljava) | Java 程序语言 |
