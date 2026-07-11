-----------------------------------
-- Schema: pgext
-----------------------------------
-- DROP SCHEMA IF EXISTS pgext CASCADE;
CREATE SCHEMA IF NOT EXISTS pgext;
COMMENT ON SCHEMA pgext IS 'PostgreSQL Extension Manager Schema';

-----------------------------------
-- Status Record
-----------------------------------
-- Singleton table tracking the last update time of various catalog components
-- There is only one record in this table with id = 0
CREATE TABLE IF NOT EXISTS pgext.status
(
    id         INTEGER PRIMARY KEY CHECK (id = 0),      -- Singleton constraint: only record with id=0 allowed
    init_time  TIMESTAMPTZ NOT NULL DEFAULT now(),      -- Initial database schema creation timestamp
    ext_time   TIMESTAMPTZ          DEFAULT now(),      -- Last extension catalog update time
    repo_time  TIMESTAMPTZ          DEFAULT now(),      -- Last repository metadata refresh time
    fetch_time TIMESTAMPTZ          DEFAULT NULL,       -- Last package metadata fetch operation time
    parse_time TIMESTAMPTZ          DEFAULT NULL,       -- Last package metadata parse completion time
    recap_time TIMESTAMPTZ          DEFAULT NULL,       -- Last catalog recapitulation time
    pkg_time   TIMESTAMPTZ          DEFAULT NULL,       -- Last availability pkg rebuild time
    extra      JSONB                DEFAULT '{}'::JSONB -- Additional status metadata
);
COMMENT ON TABLE pgext.status IS 'Extension Catalog Status Singleton Table';
COMMENT ON COLUMN pgext.status.id IS 'Singleton constraint: must always be 0, ensures only one status record exists';
COMMENT ON COLUMN pgext.status.init_time IS 'Database schema initialization timestamp, set once during creation';
COMMENT ON COLUMN pgext.status.ext_time IS 'Last time the extension catalog was updated with new extension data';
COMMENT ON COLUMN pgext.status.repo_time IS 'Last time repository metadata was refreshed or synchronized';
COMMENT ON COLUMN pgext.status.fetch_time IS 'Last time package repository metadata was fetched from upstream sources';
COMMENT ON COLUMN pgext.status.parse_time IS 'Last time fetched package metadata was parsed and processed';
COMMENT ON COLUMN pgext.status.recap_time IS 'Last time the catalog was recapitulated or summarized';
COMMENT ON COLUMN pgext.status.pkg_time IS 'Last time the extension availability pkg was rebuilt';
COMMENT ON COLUMN pgext.status.extra IS 'Additional status information and metadata in JSONB format';
INSERT INTO pgext.status(id, init_time)
VALUES (0, now());

-----------------------------------
-- PostgreSQL Versions
-----------------------------------
-- Store PostgreSQL major version information and release details
CREATE TABLE IF NOT EXISTS pgext.pg
(
    pg      INTEGER PRIMARY KEY,           -- PostgreSQL major version number (e.g., 14, 15, 16, 17, 18)
    active  BOOLEAN NOT NULL DEFAULT true, -- Whether this major version is actively supported
    version TEXT,                          -- Latest full version string (e.g., '17.6', '16.10')
    minor   INTEGER,                       -- Latest minor version number (e.g., 6, 10)
    src_url TEXT                           -- Source code tarball download URL
);
COMMENT ON TABLE pgext.pg IS 'PostgreSQL Major Version Info';
COMMENT ON COLUMN pgext.pg.pg IS 'PostgreSQL major version number (e.g., 14, 15, 16, 17, 18)';
COMMENT ON COLUMN pgext.pg.active IS 'Whether this major version is actively supported and maintained';
COMMENT ON COLUMN pgext.pg.version IS 'Latest full version string in format x.y (e.g., 17.6, 16.10)';
COMMENT ON COLUMN pgext.pg.minor IS 'Latest minor version number as integer (e.g., 6 for 17.6, 10 for 16.10)';
COMMENT ON COLUMN pgext.pg.src_url IS 'Official PostgreSQL source code tarball download URL';

-----------------------------------
-- Linux Versions
-----------------------------------
-- Store Linux distribution version information for extension packaging
CREATE TABLE IF NOT EXISTS pgext.os
(
    os        TEXT PRIMARY KEY,                     -- Unique OS identifier (e.g., 'u24.x86_64', 'el9.aarch64')
    active    BOOLEAN NOT NULL DEFAULT true,        -- Whether this OS distribution is actively supported
    os_type   TEXT    NOT NULL,                     -- Package manager type: 'rpm' or 'deb'
    os_mark   TEXT    NULL,                         -- Short OS marker (e.g., 'el9i', 'u24a', 'd12i')
    os_code   TEXT    NOT NULL,                     -- OS version code (e.g., 'el8', 'el9', 'u20', 'u22', 'u24', 'd12')
    os_major  INTEGER NULL,                         -- Major version number as integer (e.g., 9 for EL9, 24 for Ubuntu 24)
    os_arch   TEXT    NOT NULL,                     -- CPU architecture: 'x86_64' or 'aarch64'
    os_vendor TEXT    NOT NULL,                     -- OS vendor name: 'EL', 'Debian', 'Ubuntu'
    os_full   TEXT    NOT NULL,                     -- Full descriptive name (e.g., 'RockyLinux 9 x86', 'Ubuntu 24.04 x86')
    os_minor  TEXT    NULL,                         -- Latest minor version (e.g., '9.6', '24.04.3')
    tf_img    TEXT    NULL,                         -- Terraform cloud image identifier
    vg_box    TEXT    NULL,                         -- Vagrant box name for VM provisioning
    extra     JSONB   NOT NULL DEFAULT '{}'::JSONB, -- Additional metadata in JSON format
    comment   TEXT    NOT NULL DEFAULT ''           -- Extra notes or status (e.g., 'EOL', 'OK', 'TBD')
);
COMMENT ON TABLE pgext.os IS 'Linux Distribution Major Version Info';
COMMENT ON COLUMN pgext.os.os IS 'Unique OS identifier combining code and architecture (e.g., u24.x86_64, el9.aarch64)';
COMMENT ON COLUMN pgext.os.active IS 'Whether this OS distribution is actively supported and receives updates';
COMMENT ON COLUMN pgext.os.os_type IS 'Package manager type: rpm for RedHat/CentOS/Rocky, deb for Debian/Ubuntu';
COMMENT ON COLUMN pgext.os.os_mark IS 'Short OS marker for identification (e.g., el9i for EL9 x86, u24a for Ubuntu 24 ARM)';
COMMENT ON COLUMN pgext.os.os_code IS 'OS version code using 3-4 char format (el7/el8/el9/el10, u20/u22/u24/u26, d11/d12/d13)';
COMMENT ON COLUMN pgext.os.os_major IS 'Major version number as integer (e.g., 9 for EL9, 24 for Ubuntu 24.04)';
COMMENT ON COLUMN pgext.os.os_arch IS 'CPU architecture: x86_64 for Intel/AMD 64-bit, aarch64 for ARM 64-bit';
COMMENT ON COLUMN pgext.os.os_vendor IS 'Operating system vendor/family: EL (Enterprise Linux), Debian, Ubuntu';
COMMENT ON COLUMN pgext.os.os_full IS 'Full human-readable distribution name (e.g., RockyLinux 9 x86, Ubuntu 24.04 ARM)';
COMMENT ON COLUMN pgext.os.os_minor IS 'Latest minor version string (e.g., 9.6 for Rocky 9.6, 24.04.3 for Ubuntu)';
COMMENT ON COLUMN pgext.os.tf_img IS 'Terraform cloud image identifier, aliyun as example';
COMMENT ON COLUMN pgext.os.vg_box IS 'Vagrant box name for local VM development and testing';
COMMENT ON COLUMN pgext.os.extra IS 'Additional metadata stored as JSONB for flexibility';
COMMENT ON COLUMN pgext.os.comment IS 'Status notes (e.g., EOL=End of Life, OK=Stable, TBD=To Be Determined)';


-----------------------------------
-- Active PGSQL & Linux Versions
-----------------------------------
CREATE OR REPLACE VIEW pgext.active_pg AS
SELECT pg
FROM pgext.pg
WHERE active
ORDER BY pg DESC;
COMMENT ON VIEW pgext.active_pg IS 'Active PostgreSQL Major Versions';

CREATE OR REPLACE VIEW pgext.active_os AS
SELECT os
FROM pgext.os
WHERE active
ORDER BY os_major;
COMMENT ON VIEW pgext.active_os IS 'Active Linux Distribution Major Versions';


-----------------------------------
-- Extension Repository
-----------------------------------
-- Define PostgreSQL extension repositories with access information for both default and mirror sources
-- DROP TABLE IF EXISTS pgext.repository CASCADE;
CREATE TABLE IF NOT EXISTS pgext.repository
(
    id           TEXT PRIMARY KEY,              -- Repository identifier in format: os_code.os_arch.org (e.g., 'el9.x86_64.pgdg')
    os           TEXT REFERENCES pgext.os (os), -- OS platform identifier combining code and arch (e.g., 'u22.arm', 'el7.amd')
    org          TEXT,                          -- Repository organization/source (e.g., 'pgdg', 'pigsty', 'epel')
    type         TEXT,                          -- Package manager type: 'rpm' for YUM/DNF or 'deb' for APT
    os_code      TEXT,                          -- OS distribution code in 3-char format (e.g., 'u22', 'el7', 'd12')
    os_arch      TEXT,                          -- CPU architecture: 'x86_64' or 'aarch64' (simplified as 'amd' or 'arm')
    default_url  TEXT,                          -- Base URL for default repository location (upstream/official)
    default_meta TEXT,                          -- Metadata file URL for default repository (repomd.xml or Packages)
    mirror_url   TEXT,                          -- Base URL for China/regional mirror repository
    mirror_meta  TEXT,                          -- Metadata file URL for mirror repository
    comment      TEXT,                          -- Additional notes or status information about repository
    extra        JSONB DEFAULT '{}'::JSONB      -- Extended metadata in JSON format for flexibility
);

COMMENT ON TABLE pgext.repository IS 'PostgreSQL Extension Repository Definitions with Default and Mirror Sources';
COMMENT ON COLUMN pgext.repository.id IS 'Unique repository identifier in format os_code.os_arch.org (e.g., el9.x86_64.pgdg, u24.aarch64.pigsty)';
COMMENT ON COLUMN pgext.repository.os IS 'Operating system platform identifier, reference pgext.os';
COMMENT ON COLUMN pgext.repository.org IS 'Repository organization/maintainer: pgdg (PostgreSQL Global Development Group), pigsty (Pigsty Project), epel (Extra Packages for Enterprise Linux)';
COMMENT ON COLUMN pgext.repository.type IS 'Package manager type: rpm for YUM/DNF repositories (RedHat-based), deb for APT repositories (Debian-based)';
COMMENT ON COLUMN pgext.repository.os_code IS 'Operating system version code in 3-4 character format: el7/el8/el9/el10 (Enterprise Linux), u20/u22/u24/u26 (Ubuntu), d11/d12/d13 (Debian)';
COMMENT ON COLUMN pgext.repository.os_arch IS 'CPU architecture: x86_64 (Intel/AMD 64-bit) or aarch64 (ARM 64-bit), often simplified as amd or arm';
COMMENT ON COLUMN pgext.repository.default_url IS 'Base URL for default/upstream repository (e.g., https://download.postgresql.org/pub/repos/)';
COMMENT ON COLUMN pgext.repository.default_meta IS 'Full URL to repository metadata file: repodata/repomd.xml for YUM, Packages(.gz) for APT';
COMMENT ON COLUMN pgext.repository.mirror_url IS 'Base URL for China/regional mirror repository for faster access in specific regions';
COMMENT ON COLUMN pgext.repository.mirror_meta IS 'Full URL to mirror repository metadata file, using same format as default_meta';
COMMENT ON COLUMN pgext.repository.comment IS 'Additional notes, status information, or special instructions for this repository';
COMMENT ON COLUMN pgext.repository.extra IS 'Extended metadata in JSONB format for storing additional repository-specific information';


-----------------------------------
-- Repository Raw Metadata
-----------------------------------
-- Store raw repository metadata files (sqlite.bz2 for YUM, Packages for APT)
-- DROP TABLE IF EXISTS pgext.repo_data;
CREATE TABLE IF NOT EXISTS pgext.repo_data
(
    id            TEXT PRIMARY KEY REFERENCES pgext.repository (id), -- Repository ID (references pgext.repository.id)
    etag          TEXT,                                              -- HTTP ETag header for cache validation
    size          BIGINT,                                            -- Size of repository metadata in bytes
    extra         JSONB,                                             -- Additional metadata for caching
    data          BYTEA,                                             -- Raw repository metadata file content
    last_modified TIMESTAMPTZ,                                       -- HTTP Last-Modified timestamp from source
    update_at     TIMESTAMP DEFAULT now()::DATE                      -- Local record update timestamp
);

COMMENT ON TABLE pgext.repo_data IS 'Raw repository metadata cache table';
COMMENT ON COLUMN pgext.repo_data.id IS 'Repository identifier in format os_code.os_arch.org (e.g., d12.arm.pigsty, el9.x86_64.pgdg)';
COMMENT ON COLUMN pgext.repo_data.etag IS 'HTTP ETag header value for efficient cache validation and conditional requests';
COMMENT ON COLUMN pgext.repo_data.size IS 'Size of repository metadata file in bytes, used for cache management';
COMMENT ON COLUMN pgext.repo_data.extra IS 'Additional metadata for caching (compression, format, etc.) in JSONB format';
COMMENT ON COLUMN pgext.repo_data.data IS 'Raw binary repository metadata (repomd.xml for YUM, Packages file for APT)';
COMMENT ON COLUMN pgext.repo_data.last_modified IS 'HTTP Last-Modified timestamp from upstream repository source';
COMMENT ON COLUMN pgext.repo_data.update_at IS 'Local timestamp when this record was last updated or refreshed';

-----------------------------------
-- Extension Category (16 categories)
-----------------------------------
-- Classification of PostgreSQL extensions into functional categories
CREATE TABLE IF NOT EXISTS pgext.category
(
    id      INTEGER PRIMARY KEY,       -- Category ID, derived from minimum extension ID in category
    name    TEXT UNIQUE,               -- Category name (e.g., 'TIME', 'GIS', 'RAG', 'FTS', 'OLAP')
    icon1   TEXT,                      -- FontAwesome icon class for UI display
    icon2   TEXT,                      -- Lucide icon name for alternative UI frameworks
    extra   JSONB DEFAULT '{}'::JSONB, -- Additional category metadata
    zh_desc TEXT,                      -- Chinese description of category purpose
    en_desc TEXT                       -- English description of category purpose
);
COMMENT ON TABLE pgext.category IS 'PG Extension Category Constant Table';
COMMENT ON COLUMN pgext.category.id IS 'Category identifier, set to minimum extension.id within this category';
COMMENT ON COLUMN pgext.category.name IS 'Short category name (TIME, GIS, RAG, FTS, OLAP, FEAT, LANG, TYPE, etc.)';
COMMENT ON COLUMN pgext.category.icon1 IS 'FontAwesome icon class (e.g., fas fa-clock) for web UI display';
COMMENT ON COLUMN pgext.category.icon2 IS 'Lucide icon name (e.g., Clock) for React/modern UI frameworks';
COMMENT ON COLUMN pgext.category.extra IS 'Additional category metadata in JSONB format';
COMMENT ON COLUMN pgext.category.zh_desc IS 'Chinese language description of category purpose and scope';
COMMENT ON COLUMN pgext.category.en_desc IS 'English language description of category purpose and scope';

-----------------------------------
-- Extension
-----------------------------------
-- Central catalog of PostgreSQL extensions with metadata and package information
-- DROP TABLE IF EXISTS pgext.extension CASCADE;
CREATE TABLE IF NOT EXISTS pgext.extension
(
    id          INTEGER PRIMARY KEY,      -- Unique extension identifier
    name        TEXT NOT NULL UNIQUE,     -- Extension name as it appears in pg_extension
    pkg         TEXT NOT NULL,            -- Normalized package name (may differ from extension name)
    lead_ext    TEXT,                     -- Leading/primary extension in multi-extension packages
    category    TEXT,                     -- Category classification (TIME, GIS, RAG, FTS, etc.)
    state       TEXT,                     -- Extension state: available, deprecated, removed, not-ready
    url         TEXT,                     -- Extension homepage or source repository URL
    license     TEXT,                     -- Software license (PostgreSQL, MIT, Apache-2.0, etc.)
    tags        TEXT[],                   -- Additional classification tags
    version     TEXT,                     -- Latest available version of this extension
    repo        TEXT,                     -- Source repository type (github, gitlab, etc.)
    lang        TEXT,                     -- Primary programming language (C, SQL, PLpgSQL, Rust, etc.)
    contrib     BOOLEAN,                  -- Whether this is a PostgreSQL contrib extension
    lead        BOOLEAN,                  -- Whether this is the lead extension in its package
    has_bin     BOOLEAN,                  -- Whether extension includes binary executables
    has_lib     BOOLEAN,                  -- Whether extension includes shared libraries (.so/.dll)
    need_ddl    BOOLEAN,                  -- Whether requires CREATE EXTENSION DDL to install
    need_load   BOOLEAN,                  -- Whether requires LOAD or shared_preload_libraries
    trusted     BOOLEAN,                  -- Whether non-superuser can install (trusted extension)
    relocatable BOOLEAN,                  -- Whether extension can be relocated to different schema
    schemas     TEXT[],                   -- Fixed schema names if not relocatable
    pg_ver      TEXT[],                   -- Supported PostgreSQL major versions
    requires    TEXT[],                   -- Extension dependencies (other extensions required)
    require_by  TEXT[],                   -- Extensions that depend on this extension
    see_also    TEXT[],                   -- Related or similar extensions
    rpm_ver     TEXT,                     -- Latest RPM package version
    rpm_repo    TEXT,                     -- RPM repository source (PGDG, PIGSTY, etc.)
    rpm_pkg     TEXT,                     -- RPM package name template ($v for PG version)
    rpm_pg      TEXT[],                   -- PostgreSQL versions available in RPM
    rpm_deps    TEXT[],                   -- RPM package dependencies
    deb_ver     TEXT,                     -- Latest DEB package version
    deb_repo    TEXT,                     -- DEB repository source (PGDG, PIGSTY, etc.)
    deb_pkg     TEXT,                     -- DEB package name template ($v for PG version)
    deb_deps    TEXT[],                   -- DEB package dependencies
    deb_pg      TEXT[],                   -- PostgreSQL versions available in DEB
    source      TEXT,                     -- Source tarball name if built by Pigsty
    extra       JSONB,                    -- Additional metadata in JSONB format
    en_desc     TEXT,                     -- English description of extension functionality
    zh_desc     TEXT,                     -- Chinese description of extension functionality
    comment     TEXT,                     -- Additional notes or special instructions
    mtime       DATE DEFAULT CURRENT_DATE -- Last modification date of this record
);

CREATE INDEX IF NOT EXISTS ext_name_pkg_idx ON pgext.extension (name, pkg);

COMMENT ON TABLE pgext.extension IS 'Packaged PostgreSQL extension standard catalog; pgext.universe is its ecosystem superset';
COMMENT ON COLUMN pgext.extension.id IS 'Unique integer identifier for each extension';
COMMENT ON COLUMN pgext.extension.name IS 'Extension name as it appears in PostgreSQL system catalog (pg_extension)';
COMMENT ON COLUMN pgext.extension.pkg IS 'Normalized package name used for package management (may differ from extension name)';
COMMENT ON COLUMN pgext.extension.lead_ext IS 'Primary/leading extension name in multi-extension packages';
COMMENT ON COLUMN pgext.extension.category IS 'Functional category: TIME, GIS, RAG, FTS, OLAP, FEAT, LANG, TYPE, UTIL, FUNC, ADMIN, STAT, SEC, FDW, SIM, ETL';
COMMENT ON COLUMN pgext.extension.state IS 'Extension lifecycle state: available, deprecated, removed, not-ready';
COMMENT ON COLUMN pgext.extension.url IS 'Extension homepage or source code repository URL (GitHub, GitLab, etc.)';
COMMENT ON COLUMN pgext.extension.license IS 'Software license (e.g., PostgreSQL, MIT, Apache-2.0, GPL, BSD)';
COMMENT ON COLUMN pgext.extension.tags IS 'Additional classification tags as string array for flexible categorization';
COMMENT ON COLUMN pgext.extension.version IS 'Latest available version of this extension';
COMMENT ON COLUMN pgext.extension.repo IS 'Source repository hosting platform (github, gitlab, bitbucket, etc.)';
COMMENT ON COLUMN pgext.extension.lang IS 'Primary implementation language (C, SQL, PLpgSQL, Rust, Go, Python, etc.)';
COMMENT ON COLUMN pgext.extension.contrib IS 'Whether this is an official PostgreSQL contrib extension';
COMMENT ON COLUMN pgext.extension.lead IS 'Whether this is the primary/lead extension in a multi-extension package';
COMMENT ON COLUMN pgext.extension.has_bin IS 'Whether the extension package includes binary executables/utilities';
COMMENT ON COLUMN pgext.extension.has_lib IS 'Whether the extension includes shared library files (.so on Linux, .dll on Windows)';
COMMENT ON COLUMN pgext.extension.need_ddl IS 'Whether extension requires CREATE EXTENSION DDL command to install';
COMMENT ON COLUMN pgext.extension.need_load IS 'Whether requires explicit LOAD or shared_preload_libraries configuration';
COMMENT ON COLUMN pgext.extension.trusted IS 'Whether non-superuser can install this extension (trusted extension feature)';
COMMENT ON COLUMN pgext.extension.relocatable IS 'Whether extension can be relocated to a different schema after installation';
COMMENT ON COLUMN pgext.extension.schemas IS 'Fixed schema names if extension is not relocatable (must be installed in specific schema)';
COMMENT ON COLUMN pgext.extension.pg_ver IS 'Array of supported PostgreSQL major versions (e.g., {14,15,16,17,18})';
COMMENT ON COLUMN pgext.extension.requires IS 'Array of extension dependencies (other extensions that must be installed first)';
COMMENT ON COLUMN pgext.extension.require_by IS 'Array of extensions that depend on this extension (reverse dependency list)';
COMMENT ON COLUMN pgext.extension.see_also IS 'Array of related or similar extensions (for discovery and comparison)';
COMMENT ON COLUMN pgext.extension.rpm_ver IS 'Latest available RPM package version';
COMMENT ON COLUMN pgext.extension.rpm_repo IS 'RPM repository source (PGDG, PIGSTY, EPEL, etc.)';
COMMENT ON COLUMN pgext.extension.rpm_pkg IS 'RPM package name template where $v is replaced with PostgreSQL major version';
COMMENT ON COLUMN pgext.extension.rpm_pg IS 'Array of PostgreSQL versions available as RPM packages';
COMMENT ON COLUMN pgext.extension.rpm_deps IS 'Array of RPM package dependencies (system libraries and other packages)';
COMMENT ON COLUMN pgext.extension.deb_ver IS 'Latest available DEB package version';
COMMENT ON COLUMN pgext.extension.deb_repo IS 'DEB repository source (PGDG, PIGSTY, etc.)';
COMMENT ON COLUMN pgext.extension.deb_pkg IS 'DEB package name template where $v is replaced with PostgreSQL major version';
COMMENT ON COLUMN pgext.extension.deb_deps IS 'Array of DEB package dependencies (system libraries and other packages)';
COMMENT ON COLUMN pgext.extension.deb_pg IS 'Array of PostgreSQL versions available as DEB packages';
COMMENT ON COLUMN pgext.extension.source IS 'Source code tarball filename if built and distributed by Pigsty';
COMMENT ON COLUMN pgext.extension.extra IS 'Additional extension metadata stored as JSONB for extensibility';
COMMENT ON COLUMN pgext.extension.en_desc IS 'English description of extension functionality and purpose';
COMMENT ON COLUMN pgext.extension.zh_desc IS 'Chinese description of extension functionality and purpose';
COMMENT ON COLUMN pgext.extension.comment IS 'Additional notes, special instructions, or warnings';
COMMENT ON COLUMN pgext.extension.mtime IS 'Last modification timestamp of this record';

-- BEGIN PGEXT UNIVERSE DDL
-----------------------------------
-- Extension Universe
-----------------------------------
-- The complete PostgreSQL extension ecosystem. pgext.extension is the
-- packaged standard catalog and every extension row is also present here.
CREATE TABLE IF NOT EXISTS pgext.universe
(
    id                INTEGER PRIMARY KEY,
    name              TEXT NOT NULL UNIQUE,
    pkg               TEXT NOT NULL,
    lead_ext          TEXT,
    category          TEXT,
    state             TEXT,
    url               TEXT,
    license           TEXT,
    tags              TEXT[],
    version           TEXT,
    repo              TEXT,
    lang              TEXT,
    contrib           BOOLEAN,
    lead              BOOLEAN,
    has_bin           BOOLEAN,
    has_lib           BOOLEAN,
    need_ddl          BOOLEAN,
    need_load         BOOLEAN,
    trusted           BOOLEAN,
    relocatable       BOOLEAN,
    schemas           TEXT[],
    pg_ver            TEXT[],
    requires          TEXT[],
    require_by        TEXT[],
    see_also          TEXT[],
    rpm_ver           TEXT,
    rpm_repo          TEXT,
    rpm_pkg           TEXT,
    rpm_pg            TEXT[],
    rpm_deps          TEXT[],
    deb_ver           TEXT,
    deb_repo          TEXT,
    deb_pkg           TEXT,
    deb_deps          TEXT[],
    deb_pg            TEXT[],
    source            TEXT,
    ext_type          TEXT,
    ext_kernel        TEXT,
    ext_vendor        TEXT,
    star_cnt          INTEGER,
    watch_cnt         INTEGER,
    fork_cnt          INTEGER,
    last_commit_date  DATE,
    last_release_date DATE,
    last_update_date  DATE,
    extra             JSONB,
    en_desc           TEXT,
    zh_desc           TEXT,
    comment           TEXT,
    mtime             DATE DEFAULT CURRENT_DATE
);

CREATE INDEX IF NOT EXISTS universe_name_pkg_idx ON pgext.universe (name, pkg);

COMMENT ON TABLE pgext.universe IS 'PostgreSQL 扩展生态全集。该表是 pgext.extension 打包标准目录的超集，还包含源码可用但尚未打包的扩展、云厂商或托管服务专有扩展，以及经过发现和校准后确认应纳入目录的扩展。它不是任务队列或中间暂存表；面向网站报告、扩展分析、后续打包评估和人工维护。';
COMMENT ON COLUMN pgext.universe.id IS '扩展记录的全局数值 ID；作为 universe 主键，用于稳定排序、引用和与历史扩展目录兼容。';
COMMENT ON COLUMN pgext.universe.name IS '实际 PostgreSQL extension 名称，通常来自 *.control 文件名、CREATE EXTENSION 名称或官方文档；在 universe 中唯一。';
COMMENT ON COLUMN pgext.universe.pkg IS '归一化后的项目包名或源码包名；一个包可以提供多个 extension，因此它可能不同于 name。';
COMMENT ON COLUMN pgext.universe.lead_ext IS '同一软件包包含多个 extension 时的主扩展名称；单扩展包通常等于 name 或为空。';
COMMENT ON COLUMN pgext.universe.category IS '功能分类代码，限定在 TIME、GIS、RAG、FTS、OLAP、FEAT、LANG、TYPE、UTIL、FUNC、ADMIN、STAT、SEC、FDW、SIM、ETL 等目录分类中。';
COMMENT ON COLUMN pgext.universe.state IS '扩展生命周期或收录状态；已打包基线通常为 available，新增非打包扩展通常为 n/a，历史值可能包含 deprecated、removed、not-ready。';
COMMENT ON COLUMN pgext.universe.url IS '扩展主页、源码仓库、发行页面或云厂商文档的规范 URL；用于人工核查、仓库元数据匹配和报告跳转。';
COMMENT ON COLUMN pgext.universe.license IS '归一化后的许可证名称，例如 PostgreSQL、MIT、Apache-2.0、GPL-3.0、Commercial、Unknown。';
COMMENT ON COLUMN pgext.universe.tags IS '扩展标签数组，用于补充分类检索，例如 pgrx、vector、fdw、auth、timeseries、cloud 等。';
COMMENT ON COLUMN pgext.universe.version IS '已知的最新扩展版本或项目版本；优先来自 control 文件、发布信息、包元数据或官方文档。';
COMMENT ON COLUMN pgext.universe.repo IS '扩展在当前目录中的来源或打包归属标记；基线扩展常见值包括 PGDG、PIGSTY、CONTRIB、MIXED，新增未打包扩展通常为 n/a。';
COMMENT ON COLUMN pgext.universe.lang IS '主要实现语言或扩展形态，例如 C、C++、SQL、Rust、Python、Java、Data；用于生态分析和打包风险判断。';
COMMENT ON COLUMN pgext.universe.contrib IS '是否属于 PostgreSQL 官方 contrib 扩展；仅核心 PostgreSQL 随附扩展应为 true。';
COMMENT ON COLUMN pgext.universe.lead IS '是否为所在 pkg 的主扩展；网站包页面和安装矩阵通常以主扩展代表一个软件包。';
COMMENT ON COLUMN pgext.universe.has_bin IS '扩展软件包是否包含需要安装或暴露给用户的命令行二进制程序。';
COMMENT ON COLUMN pgext.universe.has_lib IS '扩展是否包含 PostgreSQL 可加载共享库；C/C++/Rust 等编译型扩展通常为 true，纯 SQL 扩展通常为 false。';
COMMENT ON COLUMN pgext.universe.need_ddl IS '是否需要执行 CREATE EXTENSION 或等价 DDL 才能启用；仅提供外部工具或数据文件的包可能为 false。';
COMMENT ON COLUMN pgext.universe.need_load IS '是否需要显式 LOAD、shared_preload_libraries、session_preload_libraries 或类似预加载配置；普通 module_pathname 不应单独视为需要预加载。';
COMMENT ON COLUMN pgext.universe.trusted IS '扩展是否可由非超级用户在满足权限条件时安装；对应 PostgreSQL trusted extension 语义。';
COMMENT ON COLUMN pgext.universe.relocatable IS '扩展对象是否允许安装到非默认 schema 或迁移 schema；来自 control 文件 relocatable 字段或文档判断。';
COMMENT ON COLUMN pgext.universe.schemas IS '扩展固定或推荐使用的 schema 列表；不可重定位或有固定对象路径要求的扩展会填写。';
COMMENT ON COLUMN pgext.universe.pg_ver IS '扩展声明支持的 PostgreSQL 主版本数组；不等同于二进制包实际可用性。';
COMMENT ON COLUMN pgext.universe.requires IS '安装或使用该扩展前需要先安装的其他 PostgreSQL extension 名称数组。';
COMMENT ON COLUMN pgext.universe.require_by IS '依赖该扩展的其他 PostgreSQL extension 名称数组；通常由全局依赖关系反向生成。';
COMMENT ON COLUMN pgext.universe.see_also IS '相关、替代或可比较的扩展名称数组，用于目录导航和人工评估。';
COMMENT ON COLUMN pgext.universe.rpm_ver IS 'RPM/YUM/DNF 渠道中已知的最新包版本；为空表示未维护 RPM 包信息。';
COMMENT ON COLUMN pgext.universe.rpm_repo IS 'RPM 包来源仓库或组织，例如 PGDG、PIGSTY、EPEL；用于解释 rpm_pkg 的来源。';
COMMENT ON COLUMN pgext.universe.rpm_pkg IS 'RPM 包名模板，通常用 $v 表示 PostgreSQL 主版本占位符，例如 postgresql$v-pgvector。';
COMMENT ON COLUMN pgext.universe.rpm_pg IS '该扩展存在 RPM 包的 PostgreSQL 主版本数组；来自包仓库解析或人工维护。';
COMMENT ON COLUMN pgext.universe.rpm_deps IS 'RPM 包层面的依赖包名数组，记录系统库、其他 RPM 包或扩展包依赖。';
COMMENT ON COLUMN pgext.universe.deb_ver IS 'DEB/APT 渠道中已知的最新包版本；为空表示未维护 DEB 包信息。';
COMMENT ON COLUMN pgext.universe.deb_repo IS 'DEB 包来源仓库或组织，例如 PGDG、PIGSTY；用于解释 deb_pkg 的来源。';
COMMENT ON COLUMN pgext.universe.deb_pkg IS 'DEB 包名模板，通常用 $v 表示 PostgreSQL 主版本占位符，例如 postgresql-$v-pgvector。';
COMMENT ON COLUMN pgext.universe.deb_deps IS 'DEB 包层面的依赖包名数组，记录系统库、其他 DEB 包或扩展包依赖。';
COMMENT ON COLUMN pgext.universe.deb_pg IS '该扩展存在 DEB 包的 PostgreSQL 主版本数组；来自包仓库解析或人工维护。';
COMMENT ON COLUMN pgext.universe.source IS 'Pigsty 自行构建时使用的源码包文件名或源码包标识；不是发现来源 URL，也不是仓库主页。';
COMMENT ON COLUMN pgext.universe.ext_type IS '扩展形态标签，从原 extra.type 上提而来；取值为 standard、preload、puresql、headless，用于区分普通 SQL+共享库扩展、预加载扩展、纯 SQL 扩展和无 CREATE EXTENSION 前端的共享库模块。';
COMMENT ON COLUMN pgext.universe.ext_kernel IS '扩展绑定或要求的 PostgreSQL 内核、发行版或兼容内核标识，从原 extra.kernel 上提而来，例如 ivorysql、babelfish、pgedge、orioledb、percona-pg。';
COMMENT ON COLUMN pgext.universe.ext_vendor IS '扩展所属厂商、云服务商或专有发行方，从原 extra.vendor 上提而来，例如 AWS、HighGo、pgEdge、Percona、HaloTech。';
COMMENT ON COLUMN pgext.universe.star_cnt IS '上游仓库 star 数；来自 GitHub/Gitea 等公开仓库元数据，非仓库或低置信度来源保持 NULL。';
COMMENT ON COLUMN pgext.universe.watch_cnt IS '上游仓库 watch/subscriber 数；优先表示 GitHub subscribers_count 或 Gitea watchers_count。';
COMMENT ON COLUMN pgext.universe.fork_cnt IS '上游仓库 fork 数；来自 GitHub/Gitea/GitLab 等公开仓库元数据。';
COMMENT ON COLUMN pgext.universe.last_commit_date IS '最新已知上游代码提交日期，或 contrib 策略日期；由 timestamptz 按 UTC 转换为 date。';
COMMENT ON COLUMN pgext.universe.last_release_date IS '最新已知上游 release/tag/PGXN 发布日期，或 contrib 发布日期；由 timestamptz 按 UTC 转换为 date。';
COMMENT ON COLUMN pgext.universe.last_update_date IS '最新已知活动日期；可来自 commit、release、tag、PGXN、官方文档或低置信度 metadata freshness。';
COMMENT ON COLUMN pgext.universe.extra IS '稳定扩展元数据 JSONB；type、kernel、vendor 已上提为 ext_type、ext_kernel、ext_vendor，star 已移除并上提为 star_cnt，pgrx 使用规范语义化版本号。';
COMMENT ON COLUMN pgext.universe.en_desc IS '英文功能描述，面向网站、报告和搜索展示；优先来自 control comment、仓库描述或官方文档摘要。';
COMMENT ON COLUMN pgext.universe.zh_desc IS '中文功能描述，面向中文网站、报告和人工维护；应是对 en_desc 或官方描述的准确中文表达。';
COMMENT ON COLUMN pgext.universe.comment IS '最终目录层面的人工备注、限制说明或维护提示；不用于保存采集过程日志。';
COMMENT ON COLUMN pgext.universe.mtime IS '该扩展记录最后一次人工或脚本维护日期；用于判断目录数据新鲜度。';
-- END PGEXT UNIVERSE DDL

-----------------------------------
-- Extension Documentation
-----------------------------------
-- Local Markdown documentation for each curated extension.
-- English and Chinese docs are synchronized from stub/ and stub-zh/.
CREATE TABLE IF NOT EXISTS pgext.doc
(
    id          INTEGER PRIMARY KEY REFERENCES pgext.extension (id) ON DELETE CASCADE,
    ext         TEXT NOT NULL UNIQUE REFERENCES pgext.extension (name) ON DELETE CASCADE,
    pkg         TEXT NOT NULL,
    repo_url    TEXT,
    license_url TEXT,
    control_url TEXT,
    author_url  TEXT,
    home_url    TEXT,
    cargo_url   TEXT,
    en_doc      TEXT,
    zh_doc      TEXT
);

CREATE INDEX IF NOT EXISTS doc_ext_pkg_idx ON pgext.doc (ext, pkg);

COMMENT ON TABLE pgext.doc IS 'PostgreSQL extension local documentation synchronized from stub directories';
COMMENT ON COLUMN pgext.doc.id IS 'Extension identifier copied from pgext.extension.id';
COMMENT ON COLUMN pgext.doc.ext IS 'Extension name copied from pgext.extension.name';
COMMENT ON COLUMN pgext.doc.pkg IS 'Package name copied from pgext.extension.pkg';
COMMENT ON COLUMN pgext.doc.repo_url IS 'Source repository URL seeded from pgext.extension.url';
COMMENT ON COLUMN pgext.doc.license_url IS 'Optional upstream license URL';
COMMENT ON COLUMN pgext.doc.control_url IS 'Optional upstream control file URL';
COMMENT ON COLUMN pgext.doc.author_url IS 'Optional upstream author or maintainer URL';
COMMENT ON COLUMN pgext.doc.home_url IS 'Optional upstream homepage URL';
COMMENT ON COLUMN pgext.doc.cargo_url IS 'pgrx Cargo.toml file URL for rust extension';
COMMENT ON COLUMN pgext.doc.en_doc IS 'English Markdown documentation from stub/<ext>.md';
COMMENT ON COLUMN pgext.doc.zh_doc IS 'Chinese Markdown documentation from stub-zh/<ext>.md';


-----------------------------------
-- YUM Packages
-----------------------------------
-- Parsed YUM/RPM repository package metadata from repomd.xml primary database
-- DROP TABLE IF EXISTS pgext.dnf;
CREATE TABLE IF NOT EXISTS pgext.dnf
(
    repo             TEXT    NOT NULL REFERENCES pgext.repository (id), -- Repository identifier
    pkg_key          INTEGER NOT NULL,                                  -- Package key from repodata
    pkg_id           TEXT    NOT NULL,                                  -- Package checksum identifier
    name             TEXT    NOT NULL,                                  -- RPM package name
    arch             TEXT,                                              -- Target architecture (x86_64, aarch64, noarch)
    version          TEXT,                                              -- Package version
    epoch            TEXT,                                              -- RPM epoch value
    release          TEXT,                                              -- RPM release string
    summary          TEXT,                                              -- Short package description
    description      TEXT,                                              -- Detailed package description
    url              TEXT,                                              -- Project homepage URL
    time_file        INTEGER,                                           -- File timestamp (Unix epoch)
    time_build       INTEGER,                                           -- Build timestamp (Unix epoch)
    rpm_license      TEXT,                                              -- Package license
    rpm_vendor       TEXT,                                              -- Package vendor/maintainer
    rpm_group        TEXT,                                              -- RPM package group classification
    rpm_buildhost    TEXT,                                              -- Hostname where package was built
    rpm_sourcerpm    TEXT,                                              -- Source RPM filename
    rpm_header_start INTEGER,                                           -- RPM header start offset
    rpm_header_end   INTEGER,                                           -- RPM header end offset
    rpm_packager     TEXT,                                              -- Person/organization who packaged
    size_package     INTEGER,                                           -- Package file size in bytes
    size_installed   INTEGER,                                           -- Installed size in bytes
    size_archive     INTEGER,                                           -- Archive size in bytes
    location_href    TEXT,                                              -- Relative package file path
    location_base    TEXT,                                              -- Base URL for package download
    checksum_type    TEXT,                                              -- Checksum algorithm (sha256, etc.)
    PRIMARY KEY (repo, pkg_key)
);
COMMENT ON TABLE pgext.dnf IS 'YUM/RPM Package Metadata Parsed from Repository Primary Database';
COMMENT ON COLUMN pgext.dnf.repo IS 'Repository identifier (foreign key to pgext.repository)';
COMMENT ON COLUMN pgext.dnf.pkg_key IS 'Unique package key within repository from repodata primary.xml';
COMMENT ON COLUMN pgext.dnf.pkg_id IS 'Package checksum identifier used for integrity verification';
COMMENT ON COLUMN pgext.dnf.name IS 'RPM package name (e.g., postgresql17-pgvector)';
COMMENT ON COLUMN pgext.dnf.arch IS 'Target CPU architecture: x86_64, aarch64, or noarch for arch-independent packages';
COMMENT ON COLUMN pgext.dnf.version IS 'Package version string';
COMMENT ON COLUMN pgext.dnf.epoch IS 'RPM epoch for version comparison (usually 0 or empty)';
COMMENT ON COLUMN pgext.dnf.release IS 'RPM release identifier (e.g., 1.el9)';
COMMENT ON COLUMN pgext.dnf.summary IS 'One-line package summary description';
COMMENT ON COLUMN pgext.dnf.description IS 'Detailed multi-line package description';
COMMENT ON COLUMN pgext.dnf.url IS 'Project homepage or documentation URL';
COMMENT ON COLUMN pgext.dnf.time_file IS 'File modification timestamp in Unix epoch seconds';
COMMENT ON COLUMN pgext.dnf.time_build IS 'Package build timestamp in Unix epoch seconds';
COMMENT ON COLUMN pgext.dnf.rpm_license IS 'Software license(s) of the package';
COMMENT ON COLUMN pgext.dnf.rpm_vendor IS 'Package vendor/maintainer organization';
COMMENT ON COLUMN pgext.dnf.rpm_group IS 'RPM package group for categorization';
COMMENT ON COLUMN pgext.dnf.rpm_buildhost IS 'Hostname of the build system';
COMMENT ON COLUMN pgext.dnf.rpm_sourcerpm IS 'Source RPM filename used to build this package';
COMMENT ON COLUMN pgext.dnf.rpm_header_start IS 'Start byte offset of RPM header in package file';
COMMENT ON COLUMN pgext.dnf.rpm_header_end IS 'End byte offset of RPM header in package file';
COMMENT ON COLUMN pgext.dnf.rpm_packager IS 'Name and email of the person/org who created the package';
COMMENT ON COLUMN pgext.dnf.size_package IS 'Size of the RPM package file in bytes';
COMMENT ON COLUMN pgext.dnf.size_installed IS 'Disk space required after installation in bytes';
COMMENT ON COLUMN pgext.dnf.size_archive IS 'Size of the compressed payload archive in bytes';
COMMENT ON COLUMN pgext.dnf.location_href IS 'Relative path to package file in repository';
COMMENT ON COLUMN pgext.dnf.location_base IS 'Base URL prefix for package download location';
COMMENT ON COLUMN pgext.dnf.checksum_type IS 'Checksum algorithm type (sha256, sha1, md5)';


-----------------------------------
-- APT Packages
-----------------------------------
-- Parsed APT/DEB repository package metadata from Packages files
-- DROP TABLE IF EXISTS pgext.apt;
CREATE TABLE IF NOT EXISTS pgext.apt
(
    repo         TEXT NOT NULL REFERENCES pgext.repository (id), -- Repository identifier
    package      TEXT NOT NULL,                                  -- Package name
    version      TEXT NOT NULL,                                  -- Package version string
    architecture TEXT NOT NULL,                                  -- Target architecture (amd64, arm64, all)
    size         INTEGER,                                        -- Package file size in bytes
    size_install INTEGER,                                        -- Installed size in kilobytes
    priority     TEXT,                                           -- Package priority (required, important, standard, optional, extra)
    section      TEXT,                                           -- Package section/category
    filename     TEXT,                                           -- Relative path to .deb file
    sha256       TEXT,                                           -- SHA256 checksum for integrity verification
    sha1         TEXT,                                           -- SHA1 checksum (legacy)
    md5sum       TEXT,                                           -- MD5 checksum (legacy)
    maintainer   TEXT,                                           -- Package maintainer name and email
    homepage     TEXT,                                           -- Project homepage URL
    depends      TEXT,                                           -- Required package dependencies
    source       TEXT,                                           -- Source package name
    provides     TEXT,                                           -- Virtual package names provided
    recommends   TEXT,                                           -- Recommended additional packages
    suggests     TEXT,                                           -- Suggested related packages
    conflicts    TEXT,                                           -- Packages that conflict with this one
    breaks       TEXT,                                           -- Packages broken by this version
    replaces     TEXT,                                           -- Packages replaced by this one
    enhances     TEXT,                                           -- Packages enhanced by this one
    pre_depends  TEXT,                                           -- Pre-installation dependencies
    build_ids    TEXT,                                           -- Debug build identifiers
    package_type TEXT,                                           -- Package type (deb, udeb for installer)
    auto_built   TEXT,                                           -- Automated build flag
    multi_arch   TEXT,                                           -- Multi-architecture support (same, foreign, allowed)
    description  TEXT,                                           -- Full package description
    extra        JSONB,                                          -- Additional fields in JSON format
    PRIMARY KEY (repo, package, version, architecture)
);
COMMENT ON TABLE pgext.apt IS 'APT/DEB Package Metadata Parsed from Repository Packages Files';
COMMENT ON COLUMN pgext.apt.repo IS 'Repository identifier (foreign key to pgext.repository)';
COMMENT ON COLUMN pgext.apt.package IS 'Debian package name (e.g., postgresql-17-pgvector)';
COMMENT ON COLUMN pgext.apt.version IS 'Package version string including upstream and Debian revision';
COMMENT ON COLUMN pgext.apt.architecture IS 'Target architecture: amd64 (x86_64), arm64 (aarch64), or all (arch-independent)';
COMMENT ON COLUMN pgext.apt.size IS 'Size of the .deb package file in bytes';
COMMENT ON COLUMN pgext.apt.size_install IS 'Approximate disk space required after installation in kilobytes';
COMMENT ON COLUMN pgext.apt.priority IS 'Package priority level: required, important, standard, optional, or extra';
COMMENT ON COLUMN pgext.apt.section IS 'Package category/section for organization (database, admin, libs, etc.)';
COMMENT ON COLUMN pgext.apt.filename IS 'Relative path to .deb file within repository';
COMMENT ON COLUMN pgext.apt.sha256 IS 'SHA256 cryptographic checksum for package integrity verification';
COMMENT ON COLUMN pgext.apt.sha1 IS 'SHA1 checksum (legacy, less secure than SHA256)';
COMMENT ON COLUMN pgext.apt.md5sum IS 'MD5 checksum (legacy, primarily for backward compatibility)';
COMMENT ON COLUMN pgext.apt.maintainer IS 'Package maintainer name and email address';
COMMENT ON COLUMN pgext.apt.homepage IS 'Project homepage or documentation URL';
COMMENT ON COLUMN pgext.apt.depends IS 'Comma-separated list of required package dependencies';
COMMENT ON COLUMN pgext.apt.source IS 'Source package name if different from binary package';
COMMENT ON COLUMN pgext.apt.provides IS 'Virtual package names that this package provides';
COMMENT ON COLUMN pgext.apt.recommends IS 'Recommended packages that enhance functionality';
COMMENT ON COLUMN pgext.apt.suggests IS 'Suggested optional packages for additional features';
COMMENT ON COLUMN pgext.apt.conflicts IS 'Packages that cannot be installed alongside this one';
COMMENT ON COLUMN pgext.apt.breaks IS 'Package versions broken by installing this package';
COMMENT ON COLUMN pgext.apt.replaces IS 'Packages that this package replaces or supersedes';
COMMENT ON COLUMN pgext.apt.enhances IS 'Packages whose functionality is enhanced by this package';
COMMENT ON COLUMN pgext.apt.pre_depends IS 'Dependencies that must be configured before installing this package';
COMMENT ON COLUMN pgext.apt.build_ids IS 'Debug symbol build identifiers for debugging';
COMMENT ON COLUMN pgext.apt.package_type IS 'Package type: deb (standard) or udeb (installer micro-package)';
COMMENT ON COLUMN pgext.apt.auto_built IS 'Whether package was automatically built by build system';
COMMENT ON COLUMN pgext.apt.multi_arch IS 'Multi-architecture support: same, foreign, allowed, or no';
COMMENT ON COLUMN pgext.apt.description IS 'Full package description with synopsis and detailed information';
COMMENT ON COLUMN pgext.apt.extra IS 'Additional non-standard fields stored as JSONB for flexibility';

-----------------------------------
-- BIN: RPM/DEB binary packages
-----------------------------------
-- DROP TABLE IF EXISTS pgext.bin;
CREATE TABLE IF NOT EXISTS pgext.bin
(
    pg         INTEGER REFERENCES pgext.pg (pg),                -- PostgreSQL major version (14-18)
    os         TEXT REFERENCES pgext.os (os),                   -- OS identifier (e.g., 'el9.x86_64', 'u24.aarch64')
    name       TEXT,                                            -- Full package name with version
    repo       TEXT REFERENCES pgext.repository (id),           -- Source repository identifier
    ver        TEXT,                                            -- Complete version string (version-release)
    version    TEXT,                                            -- Cleansed semantic version number
    file       TEXT,                                            -- Package filename
    sha256     TEXT,                                            -- SHA256 checksum
    href       TEXT,                                            -- href URL for download (+default_url/mirror_url)
    size       integer,                                         -- Package file size in bytes
    size_full  integer                                          -- Installed size in bytes
);

CREATE INDEX IF NOT EXISTS package_pg_os_idx ON pgext.bin USING BTREE (pg,os);
CREATE INDEX IF NOT EXISTS package_name_ver_idx ON pgext.bin USING BTREE (name,pg,os);

COMMENT ON TABLE pgext.bin IS 'Binary Packages (RPM/DEB) for PostgreSQL Extensions';
COMMENT ON COLUMN pgext.bin.pg IS 'PostgreSQL major version this package is built for (14, 15, 16, 17, 18)';
COMMENT ON COLUMN pgext.bin.os IS 'Operating system and architecture identifier (e.g., el9.x86_64, u24.aarch64)';
COMMENT ON COLUMN pgext.bin.repo IS 'Source repository identifier this package was parsed from';
COMMENT ON COLUMN pgext.bin.name IS 'Full package name including PostgreSQL version (e.g., postgresql-17-pgvector)';
COMMENT ON COLUMN pgext.bin.ver IS 'Complete raw version string combining version and release';
COMMENT ON COLUMN pgext.bin.version IS 'Cleansed semantic version number';
COMMENT ON COLUMN pgext.bin.file IS 'Package filename for download (e.g., postgresql17-pgvector-0.7.0-1.el9.x86_64.rpm)';
COMMENT ON COLUMN pgext.bin.sha256 IS 'SHA256 cryptographic checksum for package integrity verification';
COMMENT ON COLUMN pgext.bin.href IS 'Relative href URL for downloading the package (combined with base URL)';
COMMENT ON COLUMN pgext.bin.size IS 'Package file size in bytes (download size)';
COMMENT ON COLUMN pgext.bin.size_full IS 'Installed size';


-----------------------------------
-- Extension Package Availability
-----------------------------------
-- DROP TYPE IF EXISTS pgext.pkg_state CASCADE;
CREATE TYPE pgext.pkg_state AS ENUM ('AVAIL', 'MISS', 'HIDE', 'BREAK','THROW', 'FORK');

-- Cross-reference table showing extension package availability across PG versions and OS platforms
-- DROP TABLE IF EXISTS pgext.pkg CASCADE;
CREATE TABLE IF NOT EXISTS pgext.pkg
(
    pg      INTEGER REFERENCES pgext.pg,            -- PostgreSQL major version (14-18)
    os      TEXT REFERENCES pgext.os,               -- OS identifier (e.g., 'el9.x86_64', 'u24.aarch64')
    name    TEXT,                                   -- Versioned package name (e.g., 'postgresql-17-pgvector')
    pkg     TEXT,                                   -- Normalized extension package name (like pgvector, postgis)
    ext     TEXT REFERENCES pgext.extension (name) ON DELETE CASCADE, -- Leading extension name in this package
    state   pgext.pkg_state DEFAULT 'MISS',         -- State: AVAIL, MISS, HIDE, BREAK, THROW
    hide    BOOLEAN         DEFAULT false,          -- Hide this entry in the extension category listing
    org     TEXT            DEFAULT NULL,           -- Repository source of latest package (pgdg or pigsty)
    version TEXT            DEFAULT NULL,           -- Latest available version for this combination
    count   BIGINT          DEFAULT 0,              -- Number of available package variants
    PRIMARY KEY (pg, os, pkg)
);
CREATE INDEX IF NOT EXISTS pkg_pg_os_name_idx ON pgext.pkg (pg, os, name);
CREATE INDEX IF NOT EXISTS pkg_pkg_os_pg_idx ON pgext.pkg (pkg, os, pg);

COMMENT ON TABLE pgext.pkg IS 'PostgreSQL Extension Package Availability Table';
COMMENT ON COLUMN pgext.pkg.pg IS 'PostgreSQL major version number (18,17,16,...)';
COMMENT ON COLUMN pgext.pkg.os IS 'OS and architecture identifier in format os_code.arch (e.g., el9.x86_64, u24.aarch64)';
COMMENT ON COLUMN pgext.pkg.pkg IS 'Normalized extension package name like pgvector, postgis';
COMMENT ON COLUMN pgext.pkg.ext IS 'Leading extension name within this package';
COMMENT ON COLUMN pgext.pkg.name IS 'Full versioned package name (e.g., postgresql-17-pgvector, postgresql-16-postgis)';
COMMENT ON COLUMN pgext.pkg.state IS 'State code: AVAIL, MISS, HIDE, BREAK, THROW';
COMMENT ON COLUMN pgext.pkg.org IS 'Latest package source organization (pgdg or pigsty)';
COMMENT ON COLUMN pgext.pkg.version IS 'Latest available package version for this specific PG+OS combination';
COMMENT ON COLUMN pgext.pkg.count IS 'Count of available package variants (including different repos and versions)';

-----------------------------------
-- Domain: Semantic Version
-----------------------------------
-- DROP DOMAIN pgext.version CASCADE;
SET search_path TO pgext, public;
CREATE DOMAIN pgext.version AS TEXT;

-- Optimized version comparison following RPM/DEB rules
-- Handles: epoch, tilde (~) for pre-releases, caret (^), and standard segments
CREATE OR REPLACE FUNCTION pgext.version_compare(v1 TEXT, v2 TEXT) RETURNS INTEGER AS $$
DECLARE
    epoch1 INT := 0; epoch2 INT := 0; pos INT;
    ver1 TEXT; ver2 TEXT;
    p1 INT := 1; p2 INT := 1; l1 INT; l2 INT;
    c1 CHAR; c2 CHAR; seg1 TEXT; seg2 TEXT;
    num1 BIGINT; num2 BIGINT; cmp INT;
BEGIN
    IF v1 = v2 THEN RETURN 0; END IF;
    IF v1 IS NULL THEN RETURN -1; END IF;
    IF v2 IS NULL THEN RETURN 1; END IF;

    -- Extract epoch if present (highest priority)
    pos := position(':' IN v1);
    IF pos > 0 THEN epoch1 := left(v1, pos-1)::INT; ver1 := substr(v1, pos+1); ELSE ver1 := v1; END IF;
    pos := position(':' IN v2);
    IF pos > 0 THEN epoch2 := left(v2, pos-1)::INT; ver2 := substr(v2, pos+1); ELSE ver2 := v2; END IF;
    IF epoch1 < epoch2 THEN RETURN -1; ELSIF epoch1 > epoch2 THEN RETURN 1; END IF;

    l1 := length(ver1); l2 := length(ver2);

    -- Main comparison loop
    WHILE p1 <= l1 OR p2 <= l2 LOOP
            c1 := CASE WHEN p1 <= l1 THEN substr(ver1, p1, 1) ELSE NULL END;
            c2 := CASE WHEN p2 <= l2 THEN substr(ver2, p2, 1) ELSE NULL END;

            -- Handle tilde (sorts before everything, including NULL)
            IF c1 = '~' OR c2 = '~' THEN
                IF c1 = '~' AND c2 = '~' THEN p1 := p1 + 1; p2 := p2 + 1; CONTINUE;
                ELSIF c1 = '~' THEN RETURN -1;
                ELSE RETURN 1; END IF;
            END IF;

            -- Handle caret (sorts after normal chars but before NULL)
            IF c1 = '^' OR c2 = '^' THEN
                IF c1 = '^' AND c2 = '^' THEN p1 := p1 + 1; p2 := p2 + 1; CONTINUE;
                ELSIF c1 = '^' THEN RETURN CASE WHEN c2 IS NULL THEN -1 ELSE 1 END;
                ELSE RETURN CASE WHEN c1 IS NULL THEN 1 ELSE -1 END; END IF;
            END IF;

            -- Skip common separators
            WHILE p1 <= l1 AND substr(ver1, p1, 1) IN ('.', '-', '_', '+') LOOP p1 := p1 + 1; END LOOP;
            WHILE p2 <= l2 AND substr(ver2, p2, 1) IN ('.', '-', '_', '+') LOOP p2 := p2 + 1; END LOOP;

            -- Check for end of strings
            IF p1 > l1 AND p2 > l2 THEN RETURN 0;
            ELSIF p1 > l1 THEN RETURN -1;
            ELSIF p2 > l2 THEN RETURN 1; END IF;

            c1 := substr(ver1, p1, 1); c2 := substr(ver2, p2, 1);

            -- Extract and compare segments
            IF c1 ~ '[0-9]' AND c2 ~ '[0-9]' THEN
                -- Both numeric: extract and compare as numbers
                num1 := 0; WHILE p1 <= l1 AND substr(ver1, p1, 1) ~ '[0-9]' LOOP
                        num1 := num1 * 10 + substr(ver1, p1, 1)::INT; p1 := p1 + 1; END LOOP;
                num2 := 0; WHILE p2 <= l2 AND substr(ver2, p2, 1) ~ '[0-9]' LOOP
                        num2 := num2 * 10 + substr(ver2, p2, 1)::INT; p2 := p2 + 1; END LOOP;
                IF num1 < num2 THEN RETURN -1; ELSIF num1 > num2 THEN RETURN 1; END IF;
            ELSIF c1 ~ '[0-9]' THEN
                RETURN 1;  -- Numbers sort after letters in RPM
            ELSIF c2 ~ '[0-9]' THEN
                RETURN -1; -- Numbers sort after letters in RPM
            ELSIF c1 ~ '[a-zA-Z]' AND c2 ~ '[a-zA-Z]' THEN
                -- Both alpha: extract and compare lexically
                seg1 := ''; WHILE p1 <= l1 AND substr(ver1, p1, 1) ~ '[a-zA-Z]' LOOP
                        seg1 := seg1 || substr(ver1, p1, 1); p1 := p1 + 1; END LOOP;
                seg2 := ''; WHILE p2 <= l2 AND substr(ver2, p2, 1) ~ '[a-zA-Z]' LOOP
                        seg2 := seg2 || substr(ver2, p2, 1); p2 := p2 + 1; END LOOP;
                cmp := CASE WHEN seg1 < seg2 THEN -1 WHEN seg1 > seg2 THEN 1 ELSE 0 END;
                IF cmp != 0 THEN RETURN cmp; END IF;
            ELSE
                -- Handle other characters
                IF c1 < c2 THEN RETURN -1; ELSIF c1 > c2 THEN RETURN 1; END IF;
                p1 := p1 + 1; p2 := p2 + 1;
            END IF;
        END LOOP;
    RETURN 0;
END;
$$ LANGUAGE plpgsql IMMUTABLE STRICT PARALLEL SAFE;

-- Comparison operators (SQL for performance)
CREATE OR REPLACE FUNCTION pgext.version_lt(pgext.version, pgext.version) RETURNS BOOLEAN AS 'SELECT pgext.version_compare($1::TEXT, $2::TEXT) < 0' LANGUAGE SQL IMMUTABLE STRICT PARALLEL SAFE;
CREATE OR REPLACE FUNCTION pgext.version_le(pgext.version, pgext.version) RETURNS BOOLEAN AS 'SELECT pgext.version_compare($1::TEXT, $2::TEXT) <= 0' LANGUAGE SQL IMMUTABLE STRICT PARALLEL SAFE;
CREATE OR REPLACE FUNCTION pgext.version_eq(pgext.version, pgext.version) RETURNS BOOLEAN AS 'SELECT pgext.version_compare($1::TEXT, $2::TEXT) = 0' LANGUAGE SQL IMMUTABLE STRICT PARALLEL SAFE;
CREATE OR REPLACE FUNCTION pgext.version_ne(pgext.version, pgext.version) RETURNS BOOLEAN AS 'SELECT pgext.version_compare($1::TEXT, $2::TEXT) <> 0' LANGUAGE SQL IMMUTABLE STRICT PARALLEL SAFE;
CREATE OR REPLACE FUNCTION pgext.version_ge(pgext.version, pgext.version) RETURNS BOOLEAN AS 'SELECT pgext.version_compare($1::TEXT, $2::TEXT) >= 0' LANGUAGE SQL IMMUTABLE STRICT PARALLEL SAFE;
CREATE OR REPLACE FUNCTION pgext.version_gt(pgext.version, pgext.version) RETURNS BOOLEAN AS 'SELECT pgext.version_compare($1::TEXT, $2::TEXT) > 0' LANGUAGE SQL IMMUTABLE STRICT PARALLEL SAFE;
CREATE OR REPLACE FUNCTION pgext.version_cmp(pgext.version, pgext.version) RETURNS INTEGER AS 'SELECT pgext.version_compare($1::TEXT, $2::TEXT)' LANGUAGE SQL IMMUTABLE STRICT PARALLEL SAFE;

-- Operator definitions (compact inline form)
CREATE OPERATOR pgext.< (LEFTARG = pgext.version, RIGHTARG = pgext.version, PROCEDURE = pgext.version_lt, COMMUTATOR = >, NEGATOR = >=, RESTRICT = scalarltsel, JOIN = scalarltjoinsel);
CREATE OPERATOR pgext.<= (LEFTARG = pgext.version, RIGHTARG = pgext.version, PROCEDURE = pgext.version_le, COMMUTATOR = >=, NEGATOR = >, RESTRICT = scalarlesel, JOIN = scalarlejoinsel);
CREATE OPERATOR pgext.= (LEFTARG = pgext.version, RIGHTARG = pgext.version, PROCEDURE = pgext.version_eq, COMMUTATOR = =, NEGATOR = <>, RESTRICT = eqsel, JOIN = eqjoinsel);
CREATE OPERATOR pgext.<> (LEFTARG = pgext.version, RIGHTARG = pgext.version, PROCEDURE = pgext.version_ne, COMMUTATOR = <>, NEGATOR = =, RESTRICT = neqsel, JOIN = neqjoinsel);
CREATE OPERATOR pgext.>= (LEFTARG = pgext.version, RIGHTARG = pgext.version, PROCEDURE = pgext.version_ge, COMMUTATOR = <=, NEGATOR = <, RESTRICT = scalargesel, JOIN = scalargejoinsel);
CREATE OPERATOR pgext.> (LEFTARG = pgext.version, RIGHTARG = pgext.version, PROCEDURE = pgext.version_gt, COMMUTATOR = <, NEGATOR = <=, RESTRICT = scalargtsel, JOIN = scalargtjoinsel);

-- B-tree operator class for indexing and sorting
CREATE OPERATOR CLASS pgext.version_ops DEFAULT FOR TYPE pgext.version USING btree AS
    OPERATOR 1 pgext.<, OPERATOR 2 pgext.<=, OPERATOR 3 pgext.=, OPERATOR 4 pgext.>=, OPERATOR 5 pgext.>,
    FUNCTION 1 pgext.version_cmp(pgext.version, pgext.version);


-----------------------------------
-- View: Extension Summary
-----------------------------------
CREATE OR REPLACE VIEW pgext.summary AS
SELECT 1::INTEGER AS id, 'ALL Extensions' AS title, count(*) AS total,
       count(*) FILTER ( WHERE rpm_repo = 'PGDG' or deb_repo = 'PGDG' ) AS pgdg, count(*) FILTER ( WHERE rpm_repo = 'PIGSTY' or deb_repo = 'PIGSTY' ) AS pigsty,count(*) FILTER ( WHERE contrib ) AS contrib, 0 AS miss,
       count(*) FILTER ( WHERE pg_ver @> '{18}') AS pg18,count(*) FILTER ( WHERE pg_ver @> '{17}') AS pg17,count(*) FILTER ( WHERE pg_ver @> '{16}') AS pg16,count(*) FILTER ( WHERE pg_ver @> '{15}') AS pg15,count(*) FILTER ( WHERE pg_ver @> '{14}') AS pg14
FROM pgext.extension
UNION ALL
SELECT 2::INTEGER AS id, 'EL Extensions' AS title, count(*) AS total,
       count(*) FILTER ( WHERE rpm_repo = 'PGDG' or deb_repo = 'PGDG' ) AS pgdg, count(*) FILTER ( WHERE rpm_repo = 'PIGSTY' or deb_repo = 'PIGSTY' ) AS pigsty,count(*) FILTER ( WHERE contrib ) AS contrib, (SELECT count(*) FROM pgext.extension) - count(*) AS miss,
       count(*) FILTER ( WHERE rpm_pg @> '{18}') AS pg18,count(*) FILTER ( WHERE rpm_pg @> '{17}') AS pg17,count(*) FILTER ( WHERE rpm_pg @> '{16}') AS pg16,count(*) FILTER ( WHERE rpm_pg @> '{15}') AS pg15,count(*) FILTER ( WHERE rpm_pg @> '{14}') AS pg14
FROM pgext.extension WHERE rpm_repo IS NOT NULL
UNION ALL
SELECT 3::INTEGER AS id, 'Debian Extensions' AS title, count(*) AS total,
       count(*) FILTER ( WHERE deb_repo = 'PGDG' or deb_repo = 'PGDG' ) AS pgdg, count(*) FILTER ( WHERE deb_repo = 'PIGSTY' or deb_repo = 'PIGSTY' ) AS pigsty,count(*) FILTER ( WHERE contrib ) AS contrib, (SELECT count(*) FROM pgext.extension) - count(*) AS miss,
       count(*) FILTER ( WHERE deb_pg @> '{18}') AS pg18,count(*) FILTER ( WHERE deb_pg @> '{17}') AS pg17,count(*) FILTER ( WHERE deb_pg @> '{16}') AS pg16,count(*) FILTER ( WHERE deb_pg @> '{15}') AS pg15,count(*) FILTER ( WHERE deb_pg @> '{14}') AS pg14
FROM pgext.extension WHERE deb_repo IS NOT NULL
UNION ALL
SELECT 4::INTEGER AS id, 'ALL Packages' AS title, count(*) AS total,
       count(*) FILTER ( WHERE rpm_repo = 'PGDG' or deb_repo = 'PGDG' ) AS pgdg, count(*) FILTER ( WHERE rpm_repo = 'PIGSTY' or deb_repo = 'PIGSTY' ) AS pigsty,count(*) FILTER ( WHERE contrib ) AS contrib, 0 AS miss,
       count(*) FILTER ( WHERE pg_ver @> '{18}') AS pg18,count(*) FILTER ( WHERE pg_ver @> '{17}') AS pg17,count(*) FILTER ( WHERE pg_ver @> '{16}') AS pg16,count(*) FILTER ( WHERE pg_ver @> '{15}') AS pg15,count(*) FILTER ( WHERE pg_ver @> '{14}') AS pg14
FROM pgext.extension WHERE lead AND NOT contrib
UNION ALL
SELECT 5::INTEGER AS id, 'EL Packages' AS title, count(*) AS total,
       count(*) FILTER ( WHERE rpm_repo = 'PGDG' or deb_repo = 'PGDG' ) AS pgdg, count(*) FILTER ( WHERE rpm_repo = 'PIGSTY' or deb_repo = 'PIGSTY' ) AS pigsty,count(*) FILTER ( WHERE contrib ) AS contrib, (SELECT count(*) FROM pgext.extension WHERE lead) - count(*) AS miss,
       count(*) FILTER ( WHERE rpm_pg @> '{18}') AS pg18,count(*) FILTER ( WHERE rpm_pg @> '{17}') AS pg17,count(*) FILTER ( WHERE rpm_pg @> '{16}') AS pg16,count(*) FILTER ( WHERE rpm_pg @> '{15}') AS pg15,count(*) FILTER ( WHERE rpm_pg @> '{14}') AS pg14
FROM pgext.extension WHERE rpm_repo IS NOT NULL AND lead AND NOT contrib
UNION ALL
SELECT 5::INTEGER AS id, 'Debian Packages' AS title, count(*) AS total,
       count(*) FILTER ( WHERE deb_repo = 'PGDG' or deb_repo = 'PGDG' ) AS pgdg, count(*) FILTER ( WHERE deb_repo = 'PIGSTY' or deb_repo = 'PIGSTY' ) AS pigsty,count(*) FILTER ( WHERE contrib ) AS contrib, (SELECT count(*) FROM pgext.extension WHERE lead) - count(*) AS miss,
       count(*) FILTER ( WHERE deb_pg @> '{18}') AS pg18,count(*) FILTER ( WHERE deb_pg @> '{17}') AS pg17,count(*) FILTER ( WHERE deb_pg @> '{16}') AS pg16,count(*) FILTER ( WHERE deb_pg @> '{15}') AS pg15,count(*) FILTER ( WHERE deb_pg @> '{14}') AS pg14
FROM pgext.extension WHERE deb_repo IS NOT NULL AND lead AND NOT contrib;


-----------------------------------
-- View used by pig the PM
-----------------------------------
CREATE OR REPLACE VIEW pgext.pig AS
    SELECT
        e.id, e.name, e.pkg, e.lead_ext, e.category, e.state, e.url, e.license,
        e.tags, e.version, e.repo, e.lang, e.contrib, e.lead, e.has_bin, e.has_lib,
        e.need_ddl, e.need_load, e.trusted, e.relocatable, e.schemas, e.pg_ver,
        e.requires, e.require_by, e.see_also,
        e.rpm_ver, e.rpm_repo, e.rpm_pkg, e.rpm_pg, e.rpm_deps,
        e.deb_ver, e.deb_repo, e.deb_pkg, e.deb_deps, e.deb_pg,
        e.source,
        CASE WHEN e.lead THEN
                 COALESCE(e.extra, '{}'::jsonb) || jsonb_build_object('matrix', m.arr)
             ELSE e.extra END AS extra,
        e.en_desc, e.zh_desc, e.comment, e.mtime
    FROM pgext.extension e
             LEFT JOIN LATERAL (
        SELECT array_agg(
                       split_part(p.os,'.',1) || CASE WHEN p.os LIKE '%x86%' THEN 'i' ELSE 'a' END
                           ||':'|| p.pg ||':'|| left(p.state::TEXT,1)
                           ||':'|| CASE p.hide WHEN true THEN 't' ELSE 'f' END
                           ||':'|| COALESCE(p.count,0)
                           ||':'|| COALESCE(CASE p.org WHEN 'pgdg' then 'G' when 'pigsty' then 'P' else '' END)
                           ||':'|| COALESCE(p.version,'')
                       ORDER BY p.os, p.pg DESC
               ) AS arr
        FROM pgext.pkg p WHERE p.pkg = e.pkg
        ) m ON e.lead = true
    ORDER BY id;

COMMENT ON VIEW pgext.pig IS 'Extensions with Package Availability Matrix';
