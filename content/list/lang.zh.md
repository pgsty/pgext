---
title: "按语言"
description: "按实现语言组织的 PostgreSQL 扩展"
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

{{< language "C" >}} {{< badge content="335 个扩展" color="gray" icon="truck" >}}

传统的 PostgreSQL 扩展开发语言

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 1000 | {{< ext "timescaledb" >}} | {{< ext "timescaledb" >}} | 时序数据库扩展插件 |
| 1030 | {{< ext "periods" >}} | {{< ext "periods" >}} | 为 PERIODs 和 SYSTEM VERSIONING 提供标准 SQL 功能 |
| 1040 | {{< ext "temporal_tables" >}} | {{< ext "temporal_tables" >}} | 时态表功能支持 |
| 1070 | {{< ext "pg_cron" >}} | {{< ext "pg_cron" >}} | 定时任务调度器 |
| 1080 | {{< ext "pg_task" >}} | {{< ext "pg_task" >}} | 在特定时间点在后台执行SQL命令 |
| 1100 | {{< ext "pg_background" >}} | {{< ext "pg_background" >}} | 在后台运行 SQL 查询 |
| 1500 | {{< ext "postgis" >}} | {{< ext "postgis" >}} | PostGIS 几何和地理空间扩展 |
| 1501 | {{< ext "postgis_topology" >}} | {{< ext "postgis" >}} | PostGIS 拓扑空间类型和函数 |
| 1502 | {{< ext "postgis_raster" >}} | {{< ext "postgis" >}} | PostGIS 光栅类型和函数 |
| 1503 | {{< ext "postgis_sfcgal" >}} | {{< ext "postgis" >}} | PostGIS SFCGAL 函数 |
| 1504 | {{< ext "postgis_tiger_geocoder" >}} | {{< ext "postgis" >}} | PostGIS tiger 地理编码器和反向地理编码器 |
| 1505 | {{< ext "address_standardizer" >}} | {{< ext "postgis" >}} | 地址标准化函数。 |
| 1506 | {{< ext "address_standardizer_data_us" >}} | {{< ext "postgis" >}} | 地址标准化函数：美国数据集示例 |
| 1520 | {{< ext "pointcloud" >}} | {{< ext "pointcloud" >}} | 提供激光雷达点云数据类型支持 |
| 1521 | {{< ext "pointcloud_postgis" >}} | {{< ext "pointcloud" >}} | 将激光雷达点云与PostGIS几何类型相集成 |
| 1530 | {{< ext "h3" >}} | {{< ext "h3" "pg_h3" >}} | H3六边形层级索引支持 |
| 1531 | {{< ext "h3_postgis" >}} | {{< ext "h3" "pg_h3" >}} | H3与PostGIS集成的扩展插件 |
| 1540 | {{< ext "q3c" >}} | {{< ext "q3c" >}} | Q3C天空索引插件 |
| 1550 | {{< ext "ogr_fdw" >}} | {{< ext "ogr_fdw" >}} | GIS 数据外部数据源包装器 |
| 1590 | {{< ext "pg_geohash" >}} | {{< ext "pg_geohash" >}} | 使用GeoHash处理空间坐标的函数包 |
| 1650 | {{< ext "mobilitydb" >}} | {{< ext "mobilitydb" >}} | MobilityDB地理空间投影数据管理分析平台 |
| 1690 | {{< ext "earthdistance" >}} | {{< ext "earthdistance" >}} | 计算地球表面上的大圆距离 |
| 1800 | {{< ext "vector" >}} | {{< ext "vector" "pgvector" >}} | 向量数据类型和 ivfflat / hnsw 访问方法 |
| 1840 | {{< ext "pg_similarity" >}} | {{< ext "pg_similarity" >}} | 提供17种距离度量函数 |
| 1850 | {{< ext "smlar" >}} | {{< ext "smlar" >}} | 高效的相似度搜索函数 |
| 1880 | {{< ext "pg4ml" >}} | {{< ext "pg4ml" >}} | PG4ML是一个机器学习框架 |
| 2110 | {{< ext "pgroonga" >}} | {{< ext "pgroonga" >}} | 使用Groonga，面向所有语言的高速全文检索平台 |
| 2111 | {{< ext "pgroonga_database" >}} | {{< ext "pgroonga" >}} | PGGroonga 数据库管理模块 |
| 2120 | {{< ext "pg_bigm" >}} | {{< ext "pg_bigm" >}} | 基于二字组的多语言全文检索扩展 |
| 2130 | {{< ext "zhparser" >}} | {{< ext "zhparser" >}} | 中文分词，全文搜索解析器 |
| 2180 | {{< ext "fuzzystrmatch" >}} | {{< ext "fuzzystrmatch" >}} | 确定字符串之间的相似性和距离 |
| 2190 | {{< ext "pg_trgm" >}} | {{< ext "pg_trgm" >}} | 文本相似度测量函数与模糊检索 |
| 2400 | {{< ext "citus" >}} | {{< ext "citus" >}} | Citus 分布式数据库 |
| 2401 | {{< ext "citus_columnar" >}} | {{< ext "citus" >}} | Citus 列式存储引擎 |
| 2410 | {{< ext "columnar" >}} | {{< ext "columnar" "hydra" >}} | 开源列式存储扩展 |
| 2450 | {{< ext "duckdb_fdw" >}} | {{< ext "duckdb_fdw" >}} | DuckDB 外部数据源包装器 |
| 2510 | {{< ext "pg_partman" >}} | {{< ext "pg_partman" >}} | 用于按时间或 ID 管理分区表的扩展 |
| 2520 | {{< ext "plproxy" >}} | {{< ext "plproxy" >}} | 作为过程语言实现的数据库分区 |
| 2530 | {{< ext "pg_strom" >}} | {{< ext "pg_strom" >}} | 使用GPU与NVMe加速大数据处理 |
| 2590 | {{< ext "tablefunc" >}} | {{< ext "tablefunc" >}} | 交叉表函数 |
| 2760 | {{< ext "age" >}} | {{< ext "age" >}} | Apache AGE，图数据库扩展 （Deb可用） |
| 2780 | {{< ext "rum" >}} | {{< ext "rum" >}} | RUM 索引访问方法 |
| 2810 | {{< ext "jsquery" >}} | {{< ext "jsquery" >}} | 用于内省 JSONB 数据类型的查询类型 |
| 2820 | {{< ext "pg_hint_plan" >}} | {{< ext "pg_hint_plan" >}} | 添加强制指定执行计划的能力 |
| 2830 | {{< ext "hypopg" >}} | {{< ext "hypopg" >}} | 假设索引，用于创建一个虚拟索引检验执行计划 |
| 2850 | {{< ext "plan_filter" >}} | {{< ext "plan_filter" "pg_plan_filter" >}} | 使用执行计划代价过滤阻止特定查询语句 |
| 2860 | {{< ext "imgsmlr" >}} | {{< ext "imgsmlr" >}} | 使用Haar小波分析计算图片相似度 |
| 2870 | {{< ext "pg_ivm" >}} | {{< ext "pg_ivm" >}} | 增量维护的物化视图 |
| 2880 | {{< ext "pg_incremental" >}} | {{< ext "pg_incremental" >}} | 增量处理流式事件 |
| 2910 | {{< ext "pgq" >}} | {{< ext "pgq" >}} | 通用队列的PG实现 |
| 2920 | {{< ext "orioledb" >}} | {{< ext "orioledb" >}} | OrioleDB，下一代事务处理引擎 |
| 2951 | {{< ext "omni" >}} | {{< ext "omni" "omnigres" >}} | PostgreSQL即平台，Omnigres主扩展与加载器 |
| 2952 | {{< ext "omni_auth" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 基础会话认证管理模块 |
| 2953 | {{< ext "omni_aws" >}} | {{< ext "omni" "omnigres" >}} | Omnigres AWS S3 API封装 |
| 2954 | {{< ext "omni_cloudevents" >}} | {{< ext "omni" "omnigres" >}} | Omnigres CloudEvents 支持 |
| 2955 | {{< ext "omni_containers" >}} | {{< ext "omni" "omnigres" >}} | Omnigres Docker容器管理模块 |
| 2956 | {{< ext "omni_credentials" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 应用密钥管理模块 |
| 2958 | {{< ext "omni_email" >}} | {{< ext "omni" "omnigres" >}} | Omnigres Email 框架 |
| 2959 | {{< ext "omni_http" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 基本HTTP类型 |
| 2960 | {{< ext "omni_httpc" >}} | {{< ext "omni" "omnigres" >}} | Omnigres HTTP客户端 |
| 2961 | {{< ext "omni_httpd" >}} | {{< ext "omni" "omnigres" >}} | Omnigres HTTP服务器 |
| 2962 | {{< ext "omni_id" >}} | {{< ext "omni" "omnigres" >}} | Omnigres ID身份数据类型 |
| 2963 | {{< ext "omni_json" >}} | {{< ext "omni" "omnigres" >}} | Omnigres JSON工具箱 |
| 2964 | {{< ext "omni_kube" >}} | {{< ext "omni" "omnigres" >}} | Omnigres Kubernetes集成模块 |
| 2965 | {{< ext "omni_ledger" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 金融账本模块 |
| 2966 | {{< ext "omni_manifest" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 包管理清单模块 |
| 2967 | {{< ext "omni_mimetypes" >}} | {{< ext "omni" "omnigres" >}} | Omnigres MIME数据类型 |
| 2968 | {{< ext "omni_os" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 操作系统集成模块 |
| 2969 | {{< ext "omni_polyfill" >}} | {{< ext "omni" "omnigres" >}} | Omnigres Postgres多态API |
| 2970 | {{< ext "omni_python" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 第一类Python支持模块 |
| 2971 | {{< ext "omni_regex" >}} | {{< ext "omni" "omnigres" >}} | Omnigres PCRE兼容正则表达式模块 |
| 2972 | {{< ext "omni_rest" >}} | {{< ext "omni" "omnigres" >}} | Omnigres REST API 工具包 |
| 2973 | {{< ext "omni_schema" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 高级模式管理组件 |
| 2974 | {{< ext "omni_seq" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 分布式整型序列号 |
| 2975 | {{< ext "omni_service" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 服务管理器 |
| 2976 | {{< ext "omni_session" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 会话管理器 |
| 2977 | {{< ext "omni_sql" >}} | {{< ext "omni" "omnigres" >}} | Omnigres SQL编程组件 |
| 2979 | {{< ext "omni_sqlite" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 嵌入的SQLite支持 |
| 2980 | {{< ext "omni_test" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 测试框架 |
| 2981 | {{< ext "omni_txn" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 事务管理器模块 |
| 2982 | {{< ext "omni_types" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 高级数据类型模块 |
| 2983 | {{< ext "omni_var" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 局部变量模块 |
| 2984 | {{< ext "omni_vfs" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 虚拟文件系统 |
| 2985 | {{< ext "omni_vfs_types_v1" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 虚拟文件系统（v1） |
| 2986 | {{< ext "omni_web" >}} | {{< ext "omni" "omnigres" >}} | Omnigres Web工具箱 |
| 2987 | {{< ext "omni_worker" >}} | {{< ext "omni" "omnigres" >}} | Omnigres 通用Worker池 |
| 2988 | {{< ext "omni_xml" >}} | {{< ext "omni" "omnigres" >}} | Omnigres XML工具包 |
| 2989 | {{< ext "omni_yaml" >}} | {{< ext "omni" "omnigres" >}} | Omnigres YAML工具包 |
| 2990 | {{< ext "bloom" >}} | {{< ext "bloom" >}} | bloom 索引-基于指纹的索引 |
| 3000 | {{< ext "pg_tle" >}} | {{< ext "pg_tle" >}} | AWS 可信语言扩展 |
| 3020 | {{< ext "pllua" >}} | {{< ext "pllua" >}} | Lua 程序语言 |
| 3021 | {{< ext "hstore_pllua" >}} | {{< ext "pllua" >}} | Lua 程序语言的Hstore适配扩展 |
| 3030 | {{< ext "plluau" >}} | {{< ext "pllua" >}} | Lua 程序语言（不受信任的） |
| 3031 | {{< ext "hstore_plluau" >}} | {{< ext "pllua" >}} | Lua 程序语言的Hstore适配扩展（不受信任的） |
| 3050 | {{< ext "pldbgapi" >}} | {{< ext "pldbgapi" "pldebugger" >}} | 用于调试 PL/pgSQL 函数的服务器端支持 |
| 3060 | {{< ext "plpgsql_check" >}} | {{< ext "plpgsql_check" >}} | 对 plpgsql 函数进行扩展检查 |
| 3070 | {{< ext "plprofiler" >}} | {{< ext "plprofiler" >}} | 剖析 PL/pgSQL 函数 |
| 3080 | {{< ext "plsh" >}} | {{< ext "plsh" >}} | PL/sh 程序语言 |
| 3100 | {{< ext "plr" >}} | {{< ext "plr" >}} | 从数据库中加载R语言解释器并执行R脚本 |
| 3200 | {{< ext "pgtap" >}} | {{< ext "pgtap" >}} | PostgreSQL单元测试框架 |
| 3220 | {{< ext "dbt2" >}} | {{< ext "dbt2" >}} | OSDL-DBT-2 测试组件 |
| 3240 | {{< ext "pltcl" >}} | {{< ext "pltcl" >}} | PL/TCL 存储过程语言 |
| 3250 | {{< ext "pltclu" >}} | {{< ext "pltcl" >}} | PL/TCL 存储过程语言（未受信/高权限） |
| 3260 | {{< ext "plperl" >}} | {{< ext "plperl" >}} | PL/Perl 存储过程语言 |
| 3261 | {{< ext "bool_plperl" >}} | {{< ext "plperl" >}} | 在 bool 和 plperl 之间转换 |
| 3262 | {{< ext "hstore_plperl" >}} | {{< ext "plperl" >}} | 在 hstore 和 plperl 之间转换适配类型 |
| 3263 | {{< ext "jsonb_plperl" >}} | {{< ext "plperl" >}} | 在 jsonb 和 plperl 之间转换 |
| 3270 | {{< ext "plperlu" >}} | {{< ext "plperlu" >}} | PL/PerlU 存储过程语言（未受信/高权限） |
| 3271 | {{< ext "bool_plperlu" >}} | {{< ext "plperlu" >}} | 在 bool 和 plperlu 之间转换 |
| 3272 | {{< ext "jsonb_plperlu" >}} | {{< ext "plperlu" >}} | 在 jsonb 和 plperlu 之间转换 |
| 3273 | {{< ext "hstore_plperlu" >}} | {{< ext "plperlu" >}} | 在 hstore 和 plperlu 之间转换适配类型 |
| 3280 | {{< ext "plpgsql" >}} | {{< ext "plpgsql" >}} | PL/pgSQL 程序设计语言 |
| 3290 | {{< ext "plpython3u" >}} | {{< ext "plpython3u" >}} | PL/Python3 存储过程语言（未受信/高权限） |
| 3291 | {{< ext "jsonb_plpython3u" >}} | {{< ext "plpython3u" >}} | 在 jsonb 和 plpython3u 之间转换 |
| 3292 | {{< ext "ltree_plpython3u" >}} | {{< ext "plpython3u" >}} | 在 ltree 和 plpython3u 之间转换 |
| 3293 | {{< ext "hstore_plpython3u" >}} | {{< ext "plpython3u" >}} | 在 hstore 和 plpython3u 之间转换 |
| 3500 | {{< ext "prefix" >}} | {{< ext "prefix" "pg_prefix" >}} | 前缀树数据类型 |
| 3510 | {{< ext "semver" >}} | {{< ext "semver" "pg_semver" >}} | 语义版本号数据类型 |
| 3520 | {{< ext "unit" >}} | {{< ext "unit" "pgunit" >}} | SI 国标单位扩展 |
| 3530 | {{< ext "pgpdf" >}} | {{< ext "pgpdf" >}} | PDF数据类型，管理函数与全文检索 |
| 3550 | {{< ext "md5hash" >}} | {{< ext "md5hash" >}} | 提供128位MD5的原生数据类型 |
| 3560 | {{< ext "asn1oid" >}} | {{< ext "asn1oid" >}} | ASN1OID数据类型支持 |
| 3570 | {{< ext "roaringbitmap" >}} | {{< ext "roaringbitmap" >}} | 支持RoaringBitmap数据类型 |
| 3590 | {{< ext "pg_sphere" >}} | {{< ext "pg_sphere" "pgsphere" >}} | 球面对象函数、运算符与索引支持 |
| 3600 | {{< ext "country" >}} | {{< ext "country" "pg_country" >}} | 国家代码数据类型，遵循ISO 3166-1标准 |
| 3620 | {{< ext "currency" >}} | {{< ext "currency" "pg_currency" >}} | 使用1字节表示的货币数据类型 |
| 3630 | {{< ext "collection" >}} | {{< ext "collection" "pg_collection" >}} | 在PlPGSQL中使用的内存优化高性能集合数据结构 |
| 3700 | {{< ext "pgmp" >}} | {{< ext "pgmp" >}} | 多精度算术扩展 |
| 3710 | {{< ext "numeral" >}} | {{< ext "numeral" >}} | 数值类型扩展 |
| 3720 | {{< ext "pg_rational" >}} | {{< ext "pg_rational" >}} | 使用BIGINT表示的有理数数据类型 |
| 3730 | {{< ext "uint" >}} | {{< ext "uint" "pguint" >}} | 无符号整型数据类型 |
| 3740 | {{< ext "uint128" >}} | {{< ext "uint128" "pg_uint128" >}} | 原生128位无符号整型数据类型 |
| 3750 | {{< ext "hashtypes" >}} | {{< ext "hashtypes" >}} | 包括SHA1，MD5在内的多种哈希数据类型 |
| 3820 | {{< ext "ip4r" >}} | {{< ext "ip4r" >}} | PostgreSQL 的 IPv4/v6 和 IPv4/v6 范围索引类型 |
| 3830 | {{< ext "pg_duration" >}} | {{< ext "pg_duration" >}} | 用于表示时间段的强化数据类型 |
| 3840 | {{< ext "uri" >}} | {{< ext "uri" "pg_uri" >}} | URI数据类型 |
| 3850 | {{< ext "emailaddr" >}} | {{< ext "emailaddr" "pgemailaddr" >}} | Email地址数据类型 |
| 3860 | {{< ext "acl" >}} | {{< ext "acl" "pg_acl" >}} | ACL数据类型 |
| 3880 | {{< ext "pg_rrule" >}} | {{< ext "pg_rrule" >}} | 日历重复规则RRULE数据类型 |
| 3890 | {{< ext "timestamp9" >}} | {{< ext "timestamp9" >}} | 纳秒分辨率时间戳 |
| 3920 | {{< ext "chkpass" >}} | {{< ext "chkpass" >}} | 数据类型：自动加密的密码 |
| 3930 | {{< ext "isn" >}} | {{< ext "isn" >}} | 用于国际产品编号标准的数据类型 |
| 3940 | {{< ext "seg" >}} | {{< ext "seg" >}} | 表示线段或浮点间隔的数据类型 |
| 3950 | {{< ext "cube" >}} | {{< ext "cube" >}} | 用于存储多维立方体的数据类型 |
| 3960 | {{< ext "ltree" >}} | {{< ext "ltree" >}} | 用于表示分层树状结构的数据类型 |
| 3970 | {{< ext "hstore" >}} | {{< ext "hstore" >}} | 用于存储（键，值）对集合的数据类型 |
| 3980 | {{< ext "citext" >}} | {{< ext "citext" >}} | 提供大小写不敏感的字符串类型 |
| 3990 | {{< ext "xml2" >}} | {{< ext "xml2" >}} | XPath 查询和 XSLT |
| 4010 | {{< ext "gzip" >}} | {{< ext "gzip" "pg_gzip" >}} | 使用SQL执行Gzip压缩与解压缩 |
| 4020 | {{< ext "bzip" >}} | {{< ext "bzip" "pg_bzip" >}} | BZIP压缩解压缩函数包 |
| 4030 | {{< ext "zstd" >}} | {{< ext "zstd" "pg_zstd" >}} | ZSTD压缩解压缩函数包 |
| 4070 | {{< ext "http" >}} | {{< ext "http" "pg_http" >}} | HTTP客户端，允许在数据库内收发HTTP请求 (supabase) |
| 4080 | {{< ext "pg_net" >}} | {{< ext "pg_net" >}} | 用 SQL 进行异步非阻塞HTTP/HTTPS 请求的扩展 (supabase) |
| 4090 | {{< ext "pg_curl" >}} | {{< ext "pg_curl" >}} | 封装CURL，执行各种用URL传输数据的操作 |
| 4150 | {{< ext "pgjq" >}} | {{< ext "pgjq" >}} | 在Postgres中使用jq查询JSON |
| 4190 | {{< ext "url_encode" >}} | {{< ext "url_encode" >}} | 提供URL编码解码函数 |
| 4230 | {{< ext "pgpcre" >}} | {{< ext "pgpcre" >}} | PCRE/Perl风格的正则表达式支持 |
| 4240 | {{< ext "icu_ext" >}} | {{< ext "icu_ext" >}} | 访问ICU库提供的函数 |
| 4250 | {{< ext "pgqr" >}} | {{< ext "pgqr" >}} | 从数据库中直接生成QR二维码 |
| 4260 | {{< ext "pg_protobuf" >}} | {{< ext "pg_protobuf" >}} | 提供Protobuf函数支持 |
| 4270 | {{< ext "envvar" >}} | {{< ext "envvar" >}} | 获取环境变量的函数 |
| 4280 | {{< ext "floatfile" >}} | {{< ext "floatfile" >}} | 将浮点数组存储到文件中而不是堆表中 |
| 4300 | {{< ext "pg_readme" >}} | {{< ext "pg_readme" >}} | 为模式与扩展生成Markdown文档 |
| 4301 | {{< ext "pg_readme_test_extension" >}} | {{< ext "pg_readme" >}} | 为模式与扩展生成Markdown文档 |
| 4400 | {{< ext "hashlib" >}} | {{< ext "hashlib" "pg_hashlib" >}} | 稳定哈希函数包 |
| 4430 | {{< ext "xxhash" >}} | {{< ext "xxhash" "pg_xxhash" >}} | xxhash哈希函数包 |
| 4440 | {{< ext "shacrypt" >}} | {{< ext "shacrypt" >}} | 实现SHA256-CRYPT与SHA512-CRYPT密钥加密算法 |
| 4450 | {{< ext "cryptint" >}} | {{< ext "cryptint" >}} | 加密INT与BIGINT类型 |
| 4460 | {{< ext "pguecc" >}} | {{< ext "pguecc" "pg_ecdsa" >}} | PostgreSQL的uECC绑定，椭圆曲线加解密函数包 |
| 4540 | {{< ext "pg_uuidv7" >}} | {{< ext "pg_uuidv7" >}} | UUIDv7 支持 |
| 4550 | {{< ext "permuteseq" >}} | {{< ext "permuteseq" >}} | 伪随机数ID置换生成器 |
| 4560 | {{< ext "pg_hashids" >}} | {{< ext "pg_hashids" >}} | 加盐将整型ID转为短字符串ID |
| 4570 | {{< ext "sequential_uuids" >}} | {{< ext "sequential_uuids" >}} | 生成连续生成的UUID |
| 4600 | {{< ext "topn" >}} | {{< ext "topn" >}} | top-n JSONB 的类型 |
| 4610 | {{< ext "quantile" >}} | {{< ext "quantile" >}} | Quantile聚合函数 |
| 4620 | {{< ext "lower_quantile" >}} | {{< ext "lower_quantile" >}} | Lower Quantile 聚合函数 |
| 4630 | {{< ext "count_distinct" >}} | {{< ext "count_distinct" >}} | COUNT(DISTINCT …) 聚合的替代方案 |
| 4640 | {{< ext "omnisketch" >}} | {{< ext "omnisketch" >}} | 实现OmniSketch数据结构，实现近似摘要聚合 |
| 4650 | {{< ext "ddsketch" >}} | {{< ext "ddsketch" >}} | 实现DDSketch数据结构，实现在线的Quantile聚合 |
| 4660 | {{< ext "vasco" >}} | {{< ext "vasco" >}} | 使用MIC发现数据中隐含的关联 |
| 4670 | {{< ext "xicor" >}} | {{< ext "xicor" "pgxicor" >}} | 在PG中计算XI相关系数 |
| 4700 | {{< ext "tdigest" >}} | {{< ext "tdigest" >}} | tdigest 聚合函数 |
| 4710 | {{< ext "first_last_agg" >}} | {{< ext "first_last_agg" >}} | first() 与 last() 聚合函数 |
| 4720 | {{< ext "extra_window_functions" >}} | {{< ext "extra_window_functions" >}} | 额外的窗口函数 |
| 4730 | {{< ext "floatvec" >}} | {{< ext "floatvec" >}} | 数组类型数学运算扩展 |
| 4740 | {{< ext "aggs_for_vecs" >}} | {{< ext "aggs_for_vecs" >}} | 针对数组类型的聚合函数集合扩展 |
| 4750 | {{< ext "aggs_for_arrays" >}} | {{< ext "aggs_for_arrays" >}} | 计算数组聚合统计值的函数包 |
| 4760 | {{< ext "arraymath" >}} | {{< ext "arraymath" "pg_arraymath" >}} | 数组逐元素数学运算符包 |
| 4770 | {{< ext "pg_math" >}} | {{< ext "pg_math" >}} | 使用GSL库的数学统计函数 |
| 4780 | {{< ext "random" >}} | {{< ext "random" "pg_random" >}} | 随机数生成器 |
| 4800 | {{< ext "base36" >}} | {{< ext "base36" "pg_base36" >}} | Base36编码解码扩展 |
| 4810 | {{< ext "base62" >}} | {{< ext "base62" "pg_base62" >}} | Base62编码解码扩展 |
| 4840 | {{< ext "financial" >}} | {{< ext "financial" "pg_financial" >}} | 金融领域聚合函数 |
| 4880 | {{< ext "refint" >}} | {{< ext "refint" >}} | 实现引用完整性的函数 |
| 4881 | {{< ext "autoinc" >}} | {{< ext "autoinc" >}} | 用于自动递增字段的函数 |
| 4882 | {{< ext "insert_username" >}} | {{< ext "insert_username" >}} | 用于跟踪谁更改了表的函数 |
| 4883 | {{< ext "moddatetime" >}} | {{< ext "moddatetime" >}} | 跟踪最后修改时间 |
| 4890 | {{< ext "tsm_system_time" >}} | {{< ext "tsm_system_time" >}} | 接受毫秒数限制的 TABLESAMPLE 方法 |
| 4900 | {{< ext "dict_xsyn" >}} | {{< ext "dict_xsyn" >}} | 用于扩展同义词处理的文本搜索字典模板 |
| 4910 | {{< ext "tsm_system_rows" >}} | {{< ext "tsm_system_rows" >}} | 接受行数限制的 TABLESAMPLE 方法 |
| 4920 | {{< ext "tcn" >}} | {{< ext "tcn" >}} | 用触发器通知变更 |
| 4930 | {{< ext "uuid-ossp" >}} | {{< ext "uuid-ossp" >}} | 生成通用唯一标识符（UUIDs） |
| 4940 | {{< ext "btree_gist" >}} | {{< ext "btree_gist" >}} | 用GiST索引常见数据类型 |
| 4950 | {{< ext "btree_gin" >}} | {{< ext "btree_gin" >}} | 用GIN索引常见数据类型 |
| 4960 | {{< ext "intarray" >}} | {{< ext "intarray" >}} | 1维整数数组的额外函数、运算符和索引支持 |
| 4970 | {{< ext "intagg" >}} | {{< ext "intagg" >}} | 整数聚合器和枚举器（过时） |
| 4980 | {{< ext "dict_int" >}} | {{< ext "dict_int" >}} | 用于整数的文本搜索字典模板 |
| 4990 | {{< ext "unaccent" >}} | {{< ext "unaccent" >}} | 删除重音的文本搜索字典 |
| 5010 | {{< ext "pg_repack" >}} | {{< ext "pg_repack" >}} | 在线垃圾清理与表膨胀治理 |
| 5020 | {{< ext "pg_rewrite" >}} | {{< ext "pg_rewrite" >}} | 在线重写整表，不阻塞读写 |
| 5040 | {{< ext "pg_squeeze" >}} | {{< ext "pg_squeeze" >}} | 从关系中删除未使用空间 |
| 5050 | {{< ext "pg_dirtyread" >}} | {{< ext "pg_dirtyread" >}} | 从表中读取尚未垃圾回收的行 |
| 5060 | {{< ext "pgfincore" >}} | {{< ext "pgfincore" >}} | 检查和管理操作系统缓冲区缓存 |
| 5070 | {{< ext "pg_cooldown" >}} | {{< ext "pg_cooldown" >}} | 从缓冲区中移除特定关系的页面 |
| 5090 | {{< ext "prioritize" >}} | {{< ext "prioritize" "pg_prioritize" >}} | 获取和设置 PostgreSQL 后端的优先级 |
| 5110 | {{< ext "pg_checksums" >}} | {{< ext "pg_checksums" >}} | 在离线模式下激活/启用/禁用数据库集群的校验和功能 |
| 5120 | {{< ext "pg_readonly" >}} | {{< ext "pg_readonly" >}} | 将集群设置为只读 |
| 5150 | {{< ext "pgautofailover" >}} | {{< ext "pgautofailover" >}} | PG 自动故障迁移 |
| 5160 | {{< ext "pg_catcheck" >}} | {{< ext "pg_catcheck" >}} | 用于诊断系统目录是否损坏的工具 |
| 5170 | {{< ext "pre_prepare" >}} | {{< ext "pre_prepare" "preprepare" >}} | 在服务端预先准备好PreparedStatement备用 |
| 5200 | {{< ext "pg_orphaned" >}} | {{< ext "pg_orphaned" >}} | 处理孤儿文件的扩展插件 |
| 5210 | {{< ext "pg_crash" >}} | {{< ext "pg_crash" >}} | 向数据库进程随机发送信号模拟故障 |
| 5220 | {{< ext "pg_cheat_funcs" >}} | {{< ext "pg_cheat_funcs" >}} | 一些超级实用的作弊函数 |
| 5230 | {{< ext "fio" >}} | {{< ext "fio" "pg_fio" >}} | PostgreSQL文件IO函数包 |
| 5810 | {{< ext "pg_savior" >}} | {{< ext "pg_savior" >}} | 阻止不带条件的全表更新以避免意外事故 |
| 5820 | {{< ext "safeupdate" >}} | {{< ext "safeupdate" >}} | 强制在 UPDATE 和 DELETE 时提供 Where 条件 |
| 5840 | {{< ext "table_log" >}} | {{< ext "table_log" >}} | 记录某张表的修改日志并做表/行级时间点恢复 |
| 5880 | {{< ext "pgagent" >}} | {{< ext "pgagent" >}} | PostgreSQL任务调度工具，与PGADMIN配合使用 |
| 5890 | {{< ext "pg_prewarm" >}} | {{< ext "pg_prewarm" >}} | 预热关系数据 |
| 5900 | {{< ext "pgpool_adm" >}} | {{< ext "pgpool_adm" "pgpool" >}} | PGPool 管理函数 |
| 5910 | {{< ext "pgpool_recovery" >}} | {{< ext "pgpool_adm" "pgpool" >}} | PGPool辅助扩展，从v4.3提供的恢复函数 |
| 5920 | {{< ext "pgpool_regclass" >}} | {{< ext "pgpool_adm" "pgpool" >}} | PGPool辅助扩展，RegClass替代 |
| 5930 | {{< ext "lo" >}} | {{< ext "lo" >}} | 大对象维护 |
| 5940 | {{< ext "basic_archive" >}} | {{< ext "basic_archive" >}} | 归档模块样例 |
| 5950 | {{< ext "basebackup_to_shell" >}} | {{< ext "basebackup_to_shell" >}} | 添加一种备份到Shell终端到基础备份方式 |
| 5960 | {{< ext "old_snapshot" >}} | {{< ext "old_snapshot" >}} | 支持 old_snapshot_threshold 的实用程序 |
| 5970 | {{< ext "adminpack" >}} | {{< ext "adminpack" >}} | PostgreSQL 管理函数集合 |
| 5980 | {{< ext "amcheck" >}} | {{< ext "amcheck" >}} | 校验关系完整性 |
| 5990 | {{< ext "pg_surgery" >}} | {{< ext "pg_surgery" >}} | 对损坏的关系进行手术 |
| 6000 | {{< ext "pg_profile" >}} | {{< ext "pg_profile" >}} | PostgreSQL 数据库负载记录与AWR报表工具 |
| 6010 | {{< ext "pg_tracing" >}} | {{< ext "pg_tracing" >}} | PostgreSQL分布式Tracing |
| 6210 | {{< ext "pg_show_plans" >}} | {{< ext "pg_show_plans" >}} | 打印所有当前正在运行查询的执行计划 |
| 6220 | {{< ext "pg_stat_kcache" >}} | {{< ext "pg_stat_kcache" >}} | 内核统计信息收集 |
| 6230 | {{< ext "pg_stat_monitor" >}} | {{< ext "pg_stat_monitor" >}} | 提供查询聚合统计、客户端信息、执行计划详细信息和直方图 |
| 6240 | {{< ext "pg_qualstats" >}} | {{< ext "pg_qualstats" >}} | 收集有关 quals 的统计信息的扩展 |
| 6250 | {{< ext "pg_store_plans" >}} | {{< ext "pg_store_plans" >}} | 跟踪所有执行的 SQL 语句的计划统计信息 |
| 6270 | {{< ext "pg_wait_sampling" >}} | {{< ext "pg_wait_sampling" >}} | 基于采样的等待事件统计 |
| 6280 | {{< ext "pgsentinel" >}} | {{< ext "pgsentinel" >}} | 活跃会话历史 |
| 6290 | {{< ext "system_stats" >}} | {{< ext "system_stats" >}} | PostgreSQL 的系统统计函数 |
| 6310 | {{< ext "pgnodemx" >}} | {{< ext "pgnodemx" >}} | 使用SQL查询获取操作系统指标 |
| 6320 | {{< ext "pg_proctab" >}} | {{< ext "pgnodemx" >}} | 通过SQL接口访问操作系统进程表 |
| 6340 | {{< ext "bgw_replstatus" >}} | {{< ext "bgw_replstatus" >}} | 用于汇报本机主从状态的后台工作进程 |
| 6350 | {{< ext "pgmeminfo" >}} | {{< ext "pgmeminfo" >}} | 显示内存使用情况 |
| 6360 | {{< ext "toastinfo" >}} | {{< ext "toastinfo" >}} | 显示TOAST字段的详细信息 |
| 6380 | {{< ext "pg_relusage" >}} | {{< ext "pg_relusage" >}} | 打印查询引用的表与列 |
| 6880 | {{< ext "pg_overexplain" >}} | {{< ext "pg_overexplain" >}} | 允许 EXPLAIN 转储更多详细 |
| 6890 | {{< ext "pg_logicalinspect" >}} | {{< ext "pg_logicalinspect" >}} | 检视逻辑解码组件详情 |
| 6900 | {{< ext "pageinspect" >}} | {{< ext "pageinspect" >}} | 检查数据库页面二进制内容 |
| 6910 | {{< ext "pgrowlocks" >}} | {{< ext "pgrowlocks" >}} | 显示行级锁信息 |
| 6920 | {{< ext "sslinfo" >}} | {{< ext "sslinfo" >}} | 关于 SSL 证书的信息 |
| 6930 | {{< ext "pg_buffercache" >}} | {{< ext "pg_buffercache" >}} | 检查共享缓冲区缓存 |
| 6940 | {{< ext "pg_walinspect" >}} | {{< ext "pg_walinspect" >}} | 用于检查 PostgreSQL WAL 日志内容的函数 |
| 6950 | {{< ext "pg_freespacemap" >}} | {{< ext "pg_freespacemap" >}} | 检查自由空间映射的内容（FSM） |
| 6960 | {{< ext "pg_visibility" >}} | {{< ext "pg_visibility" >}} | 检查可见性图（VM）和页面级可见性信息 |
| 6970 | {{< ext "pgstattuple" >}} | {{< ext "pgstattuple" >}} | 显示元组级统计信息 |
| 6980 | {{< ext "auto_explain" >}} | {{< ext "auto_explain" >}} | 提供一种自动记录执行计划的手段 |
| 6990 | {{< ext "pg_stat_statements" >}} | {{< ext "pg_stat_statements" >}} | 跟踪所有执行的 SQL 语句的计划和执行统计信息 |
| 7000 | {{< ext "passwordcheck_cracklib" >}} | {{< ext "passwordcheck_cracklib" "passwordcheck" >}} | 使用cracklib加固PG用户密码 |
| 7010 | {{< ext "supautils" >}} | {{< ext "supautils" >}} | 用于在云环境中确保数据库集群的安全 |
| 7020 | {{< ext "pgsodium" >}} | {{< ext "pgsodium" >}} | 表数据加密存储 TDE |
| 7030 | {{< ext "supabase_vault" >}} | {{< ext "supabase_vault" "pg_vault" >}} | 在 Vault 中存储加密凭证的扩展 (supabase) |
| 7060 | {{< ext "pg_tde" >}} | {{< ext "pg_tde" >}} | Percona加密存储引擎 |
| 7080 | {{< ext "pgaudit" >}} | {{< ext "pgaudit" >}} | 提供审计功能 |
| 7090 | {{< ext "pgauditlogtofile" >}} | {{< ext "pgauditlogtofile" >}} | pgAudit 子扩展，将审计日志写入单独的文件中 |
| 7100 | {{< ext "pg_auth_mon" >}} | {{< ext "pg_auth_mon" >}} | 监控每个用户的连接尝试 |
| 7110 | {{< ext "credcheck" >}} | {{< ext "credcheck" >}} | 明文凭证检查器 |
| 7120 | {{< ext "pgcryptokey" >}} | {{< ext "pgcryptokey" >}} | PG密钥管理 |
| 7140 | {{< ext "logerrors" >}} | {{< ext "logerrors" >}} | 用于收集日志文件中消息统计信息的函数 |
| 7150 | {{< ext "login_hook" >}} | {{< ext "login_hook" >}} | 在用户登陆时执行login_hook.login()函数 |
| 7160 | {{< ext "set_user" >}} | {{< ext "set_user" >}} | 增加了日志记录的 SET ROLE |
| 7170 | {{< ext "pg_snakeoil" >}} | {{< ext "pg_snakeoil" >}} | PostgreSQL动态链接库反病毒功能 |
| 7180 | {{< ext "pgextwlist" >}} | {{< ext "pgextwlist" >}} | PostgreSQL扩展白名单功能 |
| 7200 | {{< ext "sslutils" >}} | {{< ext "sslutils" >}} | 使用SQL管理SSL证书 |
| 7210 | {{< ext "noset" >}} | {{< ext "noset" "pg_noset" >}} | 阻止非超级用户使用SET/RESET设置变量 |
| 7960 | {{< ext "sepgsql" >}} | {{< ext "sepgsql" >}} | 基于SELinux标签的强制访问控制 |
| 7970 | {{< ext "auth_delay" >}} | {{< ext "auth_delay" >}} | 在返回认证失败前暂停一会，避免爆破 |
| 7980 | {{< ext "pgcrypto" >}} | {{< ext "pgcrypto" >}} | 实用加解密函数 |
| 7990 | {{< ext "passwordcheck" >}} | {{< ext "passwordcheck_cracklib" "passwordcheck" >}} | 用于强制拒绝修改弱密码的扩展 |
| 8510 | {{< ext "multicorn" >}} | {{< ext "multicorn" >}} | 用Python编写自定义的外部数据源包装器 |
| 8520 | {{< ext "odbc_fdw" >}} | {{< ext "odbc_fdw" >}} | 访问ODBC可访问的任何外部数据源 |
| 8530 | {{< ext "jdbc_fdw" >}} | {{< ext "jdbc_fdw" >}} | 访问JDBC可访问的任何外部数据源 |
| 8540 | {{< ext "pgspider_ext" >}} | {{< ext "pgspider_ext" >}} | 使用多种FDW访问远程数据库服务器 |
| 8600 | {{< ext "mysql_fdw" >}} | {{< ext "mysql_fdw" >}} | MySQL外部数据包装器 |
| 8610 | {{< ext "oracle_fdw" >}} | {{< ext "oracle_fdw" >}} | 提供对Oracle的外部数据源包装器 |
| 8620 | {{< ext "tds_fdw" >}} | {{< ext "tds_fdw" >}} | TDS 数据库（Sybase/SQL Server）外部数据包装器 |
| 8630 | {{< ext "db2_fdw" >}} | {{< ext "db2_fdw" >}} | 提供对DB2的外部数据源包装器 |
| 8640 | {{< ext "sqlite_fdw" >}} | {{< ext "sqlite_fdw" >}} | SQLite 外部数据包装器 |
| 8700 | {{< ext "mongo_fdw" >}} | {{< ext "mongo_fdw" >}} | MongoDB 外部数据包装器 |
| 8710 | {{< ext "redis_fdw" >}} | {{< ext "redis_fdw" >}} | 查询外部Redis数据源 |
| 8720 | {{< ext "redis" >}} | {{< ext "redis" "pg_redis_pubsub" >}} | 从PG向Redis发送Pub/Sub消息 |
| 8730 | {{< ext "kafka_fdw" >}} | {{< ext "kafka_fdw" >}} | Kafka外部数据源包装器 |
| 8740 | {{< ext "hdfs_fdw" >}} | {{< ext "hdfs_fdw" >}} | hdfs 外部数据包装器 |
| 8750 | {{< ext "firebird_fdw" >}} | {{< ext "firebird_fdw" >}} | Firebird外部数据源包装器 |
| 8810 | {{< ext "log_fdw" >}} | {{< ext "log_fdw" >}} | 访问PostgreSQL日志文件的FDW |
| 8970 | {{< ext "dblink" >}} | {{< ext "dblink" >}} | 从数据库内连接到其他 PostgreSQL 数据库 |
| 8980 | {{< ext "file_fdw" >}} | {{< ext "file_fdw" >}} | 访问外部文件的外部数据包装器 |
| 8990 | {{< ext "postgres_fdw" >}} | {{< ext "postgres_fdw" >}} | 用于远程 PostgreSQL 服务器的外部数据包装器 |
| 9000 | {{< ext "documentdb" >}} | {{< ext "documentdb" >}} | 微软DocumentDB的API层 |
| 9010 | {{< ext "documentdb_core" >}} | {{< ext "documentdb" >}} | 微软DocumentDB的核心API层实现 |
| 9020 | {{< ext "documentdb_distributed" >}} | {{< ext "documentdb" >}} | DocumentDB多节点模式的API层 |
| 9100 | {{< ext "orafce" >}} | {{< ext "orafce" >}} | 模拟 Oracle RDBMS 的一部分函数和包的函数和运算符 |
| 9110 | {{< ext "pgtt" >}} | {{< ext "pgtt" >}} | 类似Oracle的全局临时表功能 |
| 9120 | {{< ext "session_variable" >}} | {{< ext "session_variable" >}} | Oracle兼容的会话变量/常量操作函数 |
| 9130 | {{< ext "pg_statement_rollback" >}} | {{< ext "pg_statement_rollback" >}} | 在服务端提供类似Oracle/DB2的语句级回滚能力 |
| 9300 | {{< ext "babelfishpg_common" >}} | {{< ext "babelfishpg_common" >}} | SQL Server 数据类型兼容扩展 |
| 9310 | {{< ext "babelfishpg_tsql" >}} | {{< ext "babelfishpg_tsql" >}} | SQL Server SQL语法兼容性扩展 |
| 9320 | {{< ext "babelfishpg_tds" >}} | {{< ext "babelfishpg_tds" >}} | SQL Server TDS线缆协议兼容扩展 |
| 9330 | {{< ext "babelfishpg_money" >}} | {{< ext "babelfishpg_money" >}} | SQL Server 货币数据类型兼容扩展 |
| 9400 | {{< ext "spat" >}} | {{< ext "spat" >}} | 在PG中嵌入Redis风格的内存数据库 |
| 9410 | {{< ext "pgmemcache" >}} | {{< ext "pgmemcache" >}} | 为PG提供memcached兼容接口 |
| 9500 | {{< ext "pglogical" >}} | {{< ext "pglogical" >}} | PostgreSQL逻辑复制：三方扩展实现 |
| 9501 | {{< ext "pglogical_origin" >}} | {{< ext "pglogical" >}} | 用于从 Postgres 9.4 升级时的兼容性虚拟扩展 |
| 9510 | {{< ext "pglogical_ticker" >}} | {{< ext "pglogical_ticker" >}} | pglogical复制延迟以秒计的精确视图 |
| 9520 | {{< ext "pgl_ddl_deploy" >}} | {{< ext "pgl_ddl_deploy" >}} | 使用 pglogical 执行自动 DDL 部署 |
| 9530 | {{< ext "pg_failover_slots" >}} | {{< ext "pg_failover_slots" >}} | 在Failover过程中保留复制槽 |
| 9550 | {{< ext "pgactive" >}} | {{< ext "pgactive" >}} | PostgreSQL多主逻辑复制 |
| 9630 | {{< ext "wal2json" >}} | {{< ext "wal2json" >}} | 用逻辑解码捕获 JSON 格式的 CDC 变更 |
| 9640 | {{< ext "wal2mongo" >}} | {{< ext "wal2mongo" >}} | 使用逻辑解码捕获MongoDB JSON格式的CDC变更 |
| 9650 | {{< ext "decoderbufs" >}} | {{< ext "decoderbufs" >}} | 将WAL逻辑解码为ProtocolBuffer协议的消息 |
| 9660 | {{< ext "decoder_raw" >}} | {{< ext "decoder_raw" >}} | 逻辑复制解码输出插件：RAW SQL格式 |
| 9710 | {{< ext "repmgr" >}} | {{< ext "repmgr" >}} | PostgreSQL复制管理组件 |
| 9820 | {{< ext "pg_fact_loader" >}} | {{< ext "pg_fact_loader" >}} | 在 Postgres 中构建事实表 |
| 9830 | {{< ext "pg_bulkload" >}} | {{< ext "pg_bulkload" >}} | 向 PostgreSQL 中高速加载数据 |
| 9970 | {{< ext "test_decoding" >}} | {{< ext "test_decoding" >}} | 基于SQL的WAL逻辑解码样例 |
| 9980 | {{< ext "pgoutput" >}} | {{< ext "pgoutput" >}} | PG内置的逻辑解码输出插件 |


## SQL

{{< language "SQL" >}} {{< badge content="37 个扩展" color="gray" icon="truck" >}}

纯 SQL 扩展和函数

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 1020 | {{< ext "timeseries" >}} | {{< ext "timeseries" "pg_timeseries" >}} | 时序数据API封装 |
| 1050 | {{< ext "emaj" >}} | {{< ext "emaj" >}} | 让数据库的子集具有细粒度日志和时间旅行功能 |
| 1060 | {{< ext "table_version" >}} | {{< ext "table_version" >}} | PostgreSQL 版本控制表扩展 |
| 1560 | {{< ext "geoip" >}} | {{< ext "geoip" >}} | IP 地理位置扩展（围绕 MaxMind GeoLite 数据集的包装器） |
| 2500 | {{< ext "pg_fkpart" >}} | {{< ext "pg_fkpart" >}} | 按外键实用程序进行表分区的扩展 |
| 2840 | {{< ext "index_advisor" >}} | {{< ext "index_advisor" >}} | 查询索引建议器 |
| 2900 | {{< ext "pgmq" >}} | {{< ext "pgmq" >}} | 基于Postgres实现类似AWS SQS/RSMQ的消息队列 |
| 3580 | {{< ext "pgfaceting" >}} | {{< ext "pgfaceting" >}} | 使用倒排索引的高速切面查询 |
| 3610 | {{< ext "pg_xenophile" >}} | {{< ext "pg_xenophile" >}} | PostgreSQL i8n与l10n工具包 |
| 3611 | {{< ext "l10n_table_dependent_extension" >}} | {{< ext "pg_xenophile" >}} | PostgreSQL l10n 工具包 |
| 3870 | {{< ext "debversion" >}} | {{< ext "debversion" >}} | Debian版本号数据类型 |
| 4160 | {{< ext "pgjwt" >}} | {{< ext "pgjwt" >}} | JSON Web Token API 的PG实现 (supabase) |
| 4180 | {{< ext "pg_html5_email_address" >}} | {{< ext "pg_html5_email_address" >}} | 验证Email是否符合HTML5规范的扩展 |
| 4200 | {{< ext "pgsql_tweaks" >}} | {{< ext "pgsql_tweaks" >}} | 一些日常会用到的便利函数与视图 |
| 4220 | {{< ext "pg_extra_time" >}} | {{< ext "pg_extra_time" >}} | 一些关于日期与时间的扩展函数 |
| 4310 | {{< ext "ddl_historization" >}} | {{< ext "ddl_historization" >}} | 用SQL将所有DDL变更写入到数据库表中 |
| 4320 | {{< ext "data_historization" >}} | {{< ext "data_historization" >}} | 用SQL将数据变更历史保存到分区表中 |
| 4330 | {{< ext "schedoc" >}} | {{< ext "schedoc" "pg_schedoc" >}} | 在Django与DBT之间通过注释文档交换元数据 |
| 4470 | {{< ext "sparql" >}} | {{< ext "sparql" "pgsparql" >}} | 使用SQL查询SPARQL数据源 |
| 5080 | {{< ext "ddlx" >}} | {{< ext "ddlx" "pg_ddlx" >}} | 提取数据库对象的DDL |
| 5140 | {{< ext "pg_permissions" >}} | {{< ext "pg_permissions" >}} | 查看对象权限并将其与期望状态进行比较 |
| 5180 | {{< ext "pg_upless" >}} | {{< ext "pg_upless" >}} | 检测表上的无用UPDATE |
| 5190 | {{< ext "pgcozy" >}} | {{< ext "pgcozy" >}} | 根据先前的pg_buffercache快照预热内存缓冲区 |
| 5830 | {{< ext "pg_drop_events" >}} | {{< ext "pg_drop_events" >}} | 记录删表删列删视图的事务号，辅助PITR确定时间点 |
| 6260 | {{< ext "pg_track_settings" >}} | {{< ext "pg_track_settings" >}} | 跟踪设置更改 |
| 6300 | {{< ext "meta" >}} | {{< ext "meta" "pg_meta" >}} | 标准化，更友好的PostgreSQL系统目录视图 |
| 6330 | {{< ext "pg_sqlog" >}} | {{< ext "pg_sqlog" >}} | 提供访问PostgreSQL日志的SQL接口 |
| 6800 | {{< ext "pagevis" >}} | {{< ext "pagevis" >}} | 使用ASCII字符可视化数据库物理页面布局 |
| 7130 | {{< ext "pg_jobmon" >}} | {{< ext "pg_jobmon" >}} | 记录和监控函数 |
| 7190 | {{< ext "pg_auditor" >}} | {{< ext "pg_auditor" >}} | 审计数据变更并提供闪回能力 |
| 8650 | {{< ext "pgbouncer_fdw" >}} | {{< ext "pgbouncer_fdw" >}} | 用SQL查询pgbouncer统计信息，并执行pgbouncer命令 |
| 8800 | {{< ext "aws_s3" >}} | {{< ext "aws_s3" >}} | 从S3导入导出数据的外部数据源包装器 |
| 9240 | {{< ext "pg_dbms_metadata" >}} | {{< ext "pg_dbms_metadata" >}} | 添加 Oracle DBMS_METADATA 兼容性支持的扩展 |
| 9250 | {{< ext "pg_dbms_lock" >}} | {{< ext "pg_dbms_lock" >}} | 为PG添加对 Oracle DBMS_LOCK 的完整兼容性支持 |
| 9260 | {{< ext "pg_dbms_job" >}} | {{< ext "pg_dbms_job" >}} | 添加 Oracle DBMS_JOB 兼容性支持的扩展 |
| 9540 | {{< ext "db_migrator" >}} | {{< ext "db_migrator" >}} | 使用FDW从其他DBMS迁移到PostgreSQL |
| 9700 | {{< ext "mimeo" >}} | {{< ext "mimeo" >}} | 在PostgreSQL实例间进行表级复制 |


## Rust

{{< language "Rust" >}} {{< badge content="33 个扩展" color="gray" icon="truck" >}}

使用 pgrx 框架用 Rust 编写的扩展

| ID | 扩展 | 包 | PGRX Ver | 描述 |
|:---:|:---|:---|:---|:---|
| 1010 | {{< ext "timescaledb_toolkit" >}} | {{< ext "timescaledb_toolkit" >}} | 0.12.9 | 超表分析查询，时间序列流式处理，以及其他SQL工具 |
| 1090 | {{< ext "pg_later" >}} | {{< ext "pg_later" >}} | 0.12.5 | 执行查询，并在稍后异步获取查询结果 |
| 1570 | {{< ext "pg_polyline" >}} | {{< ext "pg_polyline" >}} | 0.12.7 | Google快速Polyline编码解码扩展 |
| 1680 | {{< ext "tzf" >}} | {{< ext "tzf" "pg_tzf" >}} | 0.14.1 | 快速根据GPS经纬度坐标查找时区 |
| 1810 | {{< ext "vchord" >}} | {{< ext "vchord" >}} | 0.16.0 | 使用Rust重写的高性能向量扩展 |
| 1820 | {{< ext "vectorscale" >}} | {{< ext "vectorscale" "pgvectorscale" >}} | 0.12.9 | 使用DiskANN算法对向量进行高效索引 |
| 1830 | {{< ext "vectorize" >}} | {{< ext "vectorize" "pg_vectorize" >}} | 0.13.1 | 在PostgreSQL中封装RAG向量检索服务 |
| 1860 | {{< ext "pg_summarize" >}} | {{< ext "pg_summarize" >}} | 0.12.4 | 使用LLM对文本字段进行总结 |
| 1870 | {{< ext "pg_tiktoken" >}} | {{< ext "pg_tiktoken" >}} | 0.12.6 | 在PostgreSQL中计算OpenAI使用的Token数 |
| 1890 | {{< ext "pgml" >}} | {{< ext "pgml" >}} | 0.12.9 | PostgresML：用SQL运行机器学习算法并训练模型 |
| 2100 | {{< ext "pg_search" >}} | {{< ext "pg_search" >}} | 0.15.0 | ParadeDB BM25算法全文检索插件，ES全文检索 |
| 2140 | {{< ext "pg_bestmatch" >}} | {{< ext "pg_bestmatch" >}} | 0.12.7 | 在数据库内生成BM25稀疏向量 |
| 2150 | {{< ext "vchord_bm25" >}} | {{< ext "vchord_bm25" >}} | 0.13.1 | BM25排序算法 |
| 2160 | {{< ext "pg_tokenizer" >}} | {{< ext "pg_tokenizer" >}} | 0.13.1 | 用于全文检索的分词器 |
| 2420 | {{< ext "pg_analytics" >}} | {{< ext "pg_analytics" >}} | 0.13.0 | 由 DuckDB 驱动的数据分析引擎 |
| 2460 | {{< ext "pg_parquet" >}} | {{< ext "pg_parquet" >}} | 0.16.0 | 在PostgreSQL与本地/S3中的Parquet文件复制数据 |
| 2790 | {{< ext "pg_graphql" >}} | {{< ext "pg_graphql" >}} | 0.12.9 | PG内的GraphQL支持 |
| 2800 | {{< ext "pg_jsonschema" >}} | {{< ext "pg_jsonschema" >}} | 0.12.9 | 提供JSON Schema校验能力 |
| 2930 | {{< ext "pg_cardano" >}} | {{< ext "pg_cardano" >}} | 0.14.1 | Cardano相关工具包：加密函数，地址编解码，区块链处理 |
| 3040 | {{< ext "plprql" >}} | {{< ext "plprql" >}} | 0.11.4 | 在PostgreSQL使用PRQL——管线式关系查询语言 |
| 3540 | {{< ext "pglite_fusion" >}} | {{< ext "pglite_fusion" >}} | 0.14.1 | 在PG表中嵌入SQLite数据库作为数据类型 |
| 4170 | {{< ext "pg_smtp_client" >}} | {{< ext "pg_smtp_client" >}} | 0.12.7 | 使用SMTP从PostgreSQL内发送邮件的客户端扩展 |
| 4290 | {{< ext "pg_render" >}} | {{< ext "pg_render" >}} | 0.12.8 | 使用SQL渲染HTML页面 |
| 4500 | {{< ext "pg_idkit" >}} | {{< ext "pg_idkit" >}} | 0.15.0 | 生成各式各样的唯一标识符：UUIDv6, ULID, KSUID |
| 4510 | {{< ext "pgx_ulid" >}} | {{< ext "pgx_ulid" >}} | 0.12.7 | ULID数据类型与函数 |
| 4830 | {{< ext "pg_base58" >}} | {{< ext "pg_base58" >}} | 0.12.1 | Base58 编码/解码函数 |
| 4850 | {{< ext "convert" >}} | {{< ext "convert" "pg_convert" >}} | 0.14.1 | 用于空间里程等的公英制转换函数 |
| 5130 | {{< ext "pgdd" >}} | {{< ext "pgdd" >}} | 0.14.1 | 提供通过标准SQL查询数据库目录集簇的能力 |
| 6370 | {{< ext "explain_ui" >}} | {{< ext "explain_ui" "pg_explain_ui" >}} | 0.12.4 | 快速跳转至PEV查阅可视化执行计划 |
| 7040 | {{< ext "pg_session_jwt" >}} | {{< ext "pg_session_jwt" >}} | 0.12.9 | 使用JWT进行会话认证 |
| 7050 | {{< ext "anon" >}} | {{< ext "anon" "pg_anon" >}} | 0.14.3 | 数据匿名化处理工具 |
| 7070 | {{< ext "pgsmcrypto" >}} | {{< ext "pgsmcrypto" >}} | 0.12.6 | 为PostgreSQL提供商密算法支持：SM2,SM3,SM4 |
| 8500 | {{< ext "wrappers" >}} | {{< ext "wrappers" >}} | 0.14.3 | Supabase提供的外部数据源包装器捆绑包 |


## Data

{{< language "Data" >}} {{< badge content="10 个扩展" color="gray" icon="truck" >}}

仅包含数据的扩展

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 2170 | {{< ext "hunspell_cs_cz" >}} | {{< ext "hunspell_cs_cz" >}} | Hunspell捷克语全文检索词典 |
| 2171 | {{< ext "hunspell_de_de" >}} | {{< ext "hunspell_de_de" >}} | Hunspell德语全文检索词典 |
| 2172 | {{< ext "hunspell_en_us" >}} | {{< ext "hunspell_en_us" >}} | Hunspell英语全文检索词典 |
| 2173 | {{< ext "hunspell_fr" >}} | {{< ext "hunspell_fr" >}} | Hunspell法语全文检索词典 |
| 2174 | {{< ext "hunspell_ne_np" >}} | {{< ext "hunspell_ne_np" >}} | Hunspell尼泊尔语全文检索词典 |
| 2175 | {{< ext "hunspell_nl_nl" >}} | {{< ext "hunspell_nl_nl" >}} | Hunspell荷兰语全文检索词典 |
| 2176 | {{< ext "hunspell_nn_no" >}} | {{< ext "hunspell_nn_no" >}} | Hunspell挪威语全文检索词典 |
| 2177 | {{< ext "hunspell_pt_pt" >}} | {{< ext "hunspell_pt_pt" >}} | Hunspell葡萄牙语全文检索词典 |
| 2178 | {{< ext "hunspell_ru_ru" >}} | {{< ext "hunspell_ru_ru" >}} | Hunspell俄语全文检索词典 |
| 2179 | {{< ext "hunspell_ru_ru_aot" >}} | {{< ext "hunspell_ru_ru_aot" >}} | Hunspell俄语全文检索词典（来自AOT.ru小组） |


## C++

{{< language "C++" >}} {{< badge content="6 个扩展" color="gray" icon="truck" >}}

使用 C++ 特性和库的扩展

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 1510 | {{< ext "pgrouting" >}} | {{< ext "pgrouting" >}} | 提供寻路能力 |
| 2430 | {{< ext "pg_duckdb" >}} | {{< ext "pg_duckdb" >}} | 在PostgreSQL中的嵌入式DuckDB扩展 |
| 2440 | {{< ext "pg_mooncake" >}} | {{< ext "pg_mooncake" >}} | PostgreSQL列式存储表 |
| 2770 | {{< ext "hll" >}} | {{< ext "hll" >}} | hyperloglog 数据类型 |
| 2940 | {{< ext "rdkit" >}} | {{< ext "rdkit" >}} | 在PostgreSQL化学领域数据管理功能 |
| 3010 | {{< ext "plv8" >}} | {{< ext "plv8" >}} | PL/JavaScript (v8) 可信过程程序语言 |


## Python

{{< language "Python" >}} {{< badge content="2 个扩展" color="gray" icon="truck" >}}

使用 Python 编写的扩展

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 3210 | {{< ext "faker" >}} | {{< ext "faker" >}} | 插入生成的测试伪造数据，Python库的包装 |
| 6810 | {{< ext "powa" >}} | {{< ext "powa" >}} | PostgreSQL 工作负载分析器-核心 |


## Java

{{< language "Java" >}} {{< badge content="1 个扩展" color="gray" icon="truck" >}}

在 JVM 上运行的扩展

| ID | 扩展 | 包 | 描述 |
|:---:|:---|:---|:---|
| 3090 | {{< ext "pljava" >}} | {{< ext "pljava" >}} | Java 程序语言 |

