#!/usr/bin/env python3
"""Generate the run-scoped PostgreSQL package rebuild manifest.

The packaged extension catalog is the whitelist.  Explicit B10 support
components are the only allowed non-catalog additions.  Every other active
spec/recipe is classified as linux-owned, explicitly excluded, superseded, or
quarantined; it is never silently scheduled.
"""

from __future__ import annotations

import argparse
import collections
import csv
import hashlib
import json
import pathlib
import re
import subprocess
from typing import Any


PG_MIN = 14
PG_MAX = 18

RPM_PLATFORMS = [
    "el8.x86_64",
    "el8.aarch64",
    "el9.x86_64",
    "el9.aarch64",
    "el10.x86_64",
    "el10.aarch64",
]
DEB_PLATFORMS = [
    "d12.x86_64",
    "d12.aarch64",
    "d13.x86_64",
    "d13.aarch64",
    "u22.x86_64",
    "u22.aarch64",
    "u24.x86_64",
    "u24.aarch64",
    "u26.x86_64",
    "u26.aarch64",
]

ALIASES = {
    "rpm": {
        "pgedge-$v": "pgedge",
        "omnigres_$v": "omnigres",
        "onesparse_$v": "onesparse",
        "openhalodb-$v": "openhalodb",
        "postgresbson_$v": "postgresbson",
        "pgtde-$v": "pgtde",
    },
    "deb": {
        "pgedge-$v": "pgedge",
        "postgresql-$v-omnigres": "omnigres",
        "postgresql-$v-onesparse": "onesparse",
        "openhalodb-$v": "openhalodb",
        "postgresql-$v-pgbson": "postgresbson",
        "pgtde-$v": "pgtde",
    },
}

ALIAS_NOTES = {
    "pgedge-$v": "pgedge bundles spock, lolor, and snowflake",
    "omnigres_$v": "one omnigres distribution provides multiple omni extensions",
    "postgresql-$v-omnigres": "one omnigres distribution provides multiple omni extensions",
    "onesparse_$v": "catalog distribution one_sparse is packaged by the onesparse recipe",
    "postgresql-$v-onesparse": "catalog distribution one_sparse is packaged by the onesparse recipe",
    "openhalodb-$v": "catalog family openhalo is packaged by openhalodb",
    "postgresbson_$v": "extension pgbson is packaged from the postgresbson source/root",
    "postgresql-$v-pgbson": "extension pgbson is packaged from the postgresbson source/root",
    "pgtde-$v": "extension pg_tde is provided by the pgtde kernel bundle",
}

# Non-catalog components explicitly allowed by B10.  Platform filters keep a
# support package only where an allowed consumer needs it.
SUPPORT = {
    "rpm": [
        ("graphblas", "graphblas", [], ["graphblas-10.5.0.tar.gz"], None, "graphblas"),
        ("lagraph", "lagraph", [], ["lagraph-1.2.2.tar.gz"], None, "graphblas"),
        ("scws", "scws", [], ["scws-1.2.3.tar.bz2"], None, "zhparser"),
        ("libfq", "libfq", [], ["libfq-0.6.2.tar.gz"], None, "libfq"),
        ("libpgfeutils14", "libpgfeutils14", [14], ["postgresql-14.23.tar.gz"], None, "libpgfeutils"),
        ("libpgfeutils15", "libpgfeutils15", [15], ["postgresql-15.18.tar.gz"], None, "libpgfeutils"),
        ("libpgfeutils16", "libpgfeutils16", [16], ["postgresql-16.14.tar.gz"], None, "libpgfeutils"),
        ("libpgfeutils17", "libpgfeutils17", [17], ["postgresql-17.10.tar.gz"], None, "libpgfeutils"),
        ("libpgfeutils18", "libpgfeutils18", [18], ["postgresql-18.4.tar.gz"], None, "libpgfeutils"),
        (
            "inchi",
            "inchi",
            [],
            ["inchi_1.07.5+dfsg.orig.tar.xz"],
            ["el9.x86_64", "el9.aarch64", "el10.x86_64", "el10.aarch64"],
            "rdkit",
        ),
        (
            "antlr4-runtime413",
            "antlr4-runtime413",
            [],
            ["antlr4-cpp-runtime-4.13.2-source.zip"],
            None,
            "babelfish",
        ),
        ("zlog", "zlog", [], ["zlog-1.2.18.tar.gz"], None, "polarstore"),
        (
            "polarstore",
            "polarstore",
            [],
            ["polarstore-1.2.42-d0c5dc6.tar.gz"],
            None,
            "polarstore",
        ),
    ],
    "deb": [
        ("graphblas", "graphblas", [], ["graphblas-10.5.0.tar.gz"], None, "graphblas"),
        ("lagraph", "lagraph", [], ["lagraph-1.2.2.tar.gz"], None, "graphblas"),
        ("scws", "scws", [], ["scws-1.2.3.tar.bz2"], None, "zhparser"),
        ("libfq", "libfq", [], ["libfq-0.6.2.tar.gz"], None, "libfq"),
        (
            "pgsodium-libsodium",
            "pgsodium-libsodium",
            [],
            ["libsodium-1.0.22.tar.gz"],
            None,
            "pgsodium",
        ),
        (
            "inchi",
            "inchi",
            [],
            [
                "inchi_1.07.5+dfsg-1.dsc",
                "inchi_1.07.5+dfsg.orig.tar.xz",
                "inchi_1.07.5+dfsg-1.debian.tar.xz",
            ],
            ["u26.x86_64", "u26.aarch64"],
            "rdkit",
        ),
        (
            "antlr4-runtime413",
            "antlr4_runtime413",
            [],
            ["antlr4-cpp-runtime-4.13.2-source.zip"],
            None,
            "babelfish",
        ),
        ("zlog", "zlog", [], ["zlog-1.2.18.tar.gz"], None, "polarstore"),
        (
            "polarstore",
            "polarstore",
            [],
            ["polarstore-1.2.42-d0c5dc6.tar.gz"],
            None,
            "polarstore",
        ),
    ],
}

SUPPORT_RECIPES = {
    "graphblas",
    "lagraph",
    "onesparse",
    "scws",
    "zhparser",
    "libfq",
    "firebird_fdw",
    "libpgfeutils14",
    "libpgfeutils15",
    "libpgfeutils16",
    "libpgfeutils17",
    "libpgfeutils18",
    "pgactive",
    "pgsodium-libsodium",
    "pgsodium",
    "inchi",
    "rdkit",
    "antlr4-runtime413",
    "antlr4_runtime413",
    "zlog",
    "polarstore",
}

KERNEL_RECIPES = {
    "agensgraph",
    "babelfish",
    "cloudberry",
    "ivorysql",
    "openhalodb",
    "orioledb",
    "pgedge",
    "pgtde",
}

HEAVY_RECIPES = {
    "citus",
    "documentdb",
    "mobilitydb",
    "omnigres",
    "pg_duckdb",
    "pg_ducklake",
    "pg_lake",
    "pg_mooncake",
    "postgis",
    "timescaledb",
}

LINUX_OWNED = {
    "ansible",
    "haproxy",
    "nginx",
    "redis",
    "valkey",
    "docker-ce",
    "docker-ce-cli",
    "containerd.io",
    "docker-buildx-plugin",
    "docker-compose-plugin",
    "docker-ce-rootless-extras",
}

EXPLICIT_EXCLUDE = {
    "cri-dockerd",
    "keepalived",
    "percona-telemetry-agent",
    "python3-cdiff",
    "pg_analytics",
    "pg_lakehouse",
    "promscale",
    "polardb",
}

EXPLICIT_QUARANTINE = {
    "neon",
    "pg_ash",
    "pg_mon",
    "pg_timeit",
    "pg_typescript",
    "pgzint",
    "mysqlcompat",
    "spat",
    "postgresql_anonymizer",
    "pg_monetdb",
}

SUPERSEDED_RECIPES = {
    "spock": "provided by pgedge",
    "lolor": "provided by pgedge",
    "snowflake": "provided by pgedge",
    "pg_tde": "provided by pgtde",
}

EXTERNAL_DEPENDENCIES = [
    {
        "package": "apache-arrow",
        "version": "25.0.1",
        "release_policy": "use upstream repository closure; do not fabricate unsupported D12/U26 packages",
        "status": "external_dependency",
    },
    {
        "package": "groonga",
        "version": "16.1.0",
        "release_policy": "use upstream repository closure",
        "status": "external_dependency",
    },
]

# PGDG catalog rows intentionally omit Pigsty source provenance.  These
# allowlisted gap recipes nevertheless have explicit immutable build inputs in
# their current spec/Makefile, recorded here so gap jobs cannot run with an
# empty source field.
SOURCE_OVERRIDES = {
    ("rpm", "pg_auto_failover_$v"): ["pg_auto_failover-2.2.tar.gz"],
    ("rpm", "pgl_ddl_deploy_$v"): ["pgl_ddl_deploy-2.2.1.tar.gz"],
    ("deb", "postgresql-$v-pgl-ddl-deploy"): ["pgl_ddl_deploy-2.2.1.tar.gz"],
    ("rpm", "pg_strom_$v"): ["v3.5.tar.gz"],
    ("rpm", "pgmemcache_$v"): ["pgmemcache-2.3.0.tar.gz"],
    ("rpm", "h3-pg_$v"): ["h3-pg-4.2.3.tar.gz"],
}

DEB_SUPPORT_PRIMARY = {
    "graphblas": "libgraphblas10",
    "lagraph": "liblagraph1",
    "inchi": "libinchi1.07",
    "antlr4_runtime413": "libantlr4-runtime413",
}

# These recipes produce a distinct source tree/control file (and, for AGE,
# a distinct upstream version) for each PostgreSQL major.  Keep them as one-PG
# jobs so the build target and artifact guard can validate an exact package
# set rather than trusting the final control file from a multi-target loop.
DEB_SINGLE_PG_RECIPES = {"age", "babelfish", "orioledb", "pgedge", "pg_duckdb"}


def run_text(argv: list[str], cwd: pathlib.Path | None = None) -> str:
    return subprocess.check_output(argv, cwd=cwd, text=True).strip()


def run_bytes(argv: list[str], cwd: pathlib.Path | None = None) -> bytes:
    return subprocess.check_output(argv, cwd=cwd)


def git_head(repo: pathlib.Path) -> str:
    return run_text(["git", "rev-parse", "HEAD"], repo)


def git_status(repo: pathlib.Path) -> str:
    return run_text(["git", "status", "--porcelain=v1", "--untracked-files=all"], repo)


def load_catalog_db_export() -> tuple[bytes, str]:
    export = run_bytes(
        [
            "psql",
            "-X",
            "-q",
            "-h",
            "/tmp",
            "-p",
            "5432",
            "-d",
            "data",
            "-U",
            "postgres",
            "-c",
            "COPY (SELECT * FROM pgext.extension ORDER BY id) TO STDOUT WITH (FORMAT csv, HEADER true)",
        ]
    )
    snapshot_at = run_text(
        [
            "psql",
            "-X",
            "-h",
            "/tmp",
            "-p",
            "5432",
            "-d",
            "data",
            "-U",
            "postgres",
            "-A",
            "-t",
            "-c",
            "select clock_timestamp()",
        ]
    )
    return export, snapshot_at


def parse_pg_array(value: str) -> list[int]:
    if not value:
        return []
    result = []
    for token in value.strip("{}").split(","):
        if token and token.isdigit() and PG_MIN <= int(token) <= PG_MAX:
            result.append(int(token))
    return sorted(set(result))


def load_pgdg_state() -> tuple[set[tuple[str, str, int]], dict[str, str]]:
    # pgext.pkg is an aggregate availability matrix.  Its org/version summary
    # fields may be null after a staged parser refresh, so prove PGDG ownership
    # against the immutable lower-level package rows and repository dimension.
    query = (
        "select distinct p.ext,p.os,p.pg "
        "from pgext.pkg p "
        "join pgext.bin b on b.pg=p.pg and b.os=p.os and b.name=p.name "
        "join pgext.repository r on r.id=b.repo "
        "where r.org='pgdg' and p.state='AVAIL'"
    )
    output = run_text(
        [
            "psql",
            "-X",
            "-h",
            "/tmp",
            "-p",
            "5432",
            "-d",
            "data",
            "-U",
            "postgres",
            "-A",
            "-t",
            "-F",
            "|",
            "-c",
            query,
        ]
    )
    available = {
        (ext, os_name, int(pg))
        for ext, os_name, pg in (line.split("|") for line in output.splitlines() if line)
    }
    stamp_query = (
        "select type||'.'||org,min(update_at)::text||'/'||max(update_at)::text "
        "from pgext.repository join pgext.repo_data using(id) "
        "group by type,org order by type,org"
    )
    stamp_output = run_text(
        [
            "psql",
            "-X",
            "-h",
            "/tmp",
            "-p",
            "5432",
            "-d",
            "data",
            "-U",
            "postgres",
            "-A",
            "-t",
            "-F",
            "|",
            "-c",
            stamp_query,
        ]
    )
    stamps = dict(line.split("|", 1) for line in stamp_output.splitlines() if line)
    return available, stamps


def recipe_inventory(fmt: str, rpm_repo: pathlib.Path, deb_repo: pathlib.Path) -> set[str]:
    if fmt == "rpm":
        return {path.stem for path in (rpm_repo / "rpmbuild/SPECS").glob("*.spec")}
    return {
        path.name
        for path in (deb_repo / "debbuild").iterdir()
        if path.is_dir() and (path / "Makefile").is_file()
    }


def recipe_path(fmt: str, recipe: str) -> str:
    if fmt == "rpm":
        return f"rpmbuild/SPECS/{recipe}.spec"
    return f"debbuild/{recipe}"


def map_recipe(fmt: str, template: str, rows: list[dict[str, str]], inventory: set[str]) -> tuple[str | None, str]:
    if template in ALIASES[fmt]:
        recipe = ALIASES[fmt][template]
        return (recipe if recipe in inventory else None, "explicit_alias")

    candidates: list[str] = []
    preferred = sorted(
        rows,
        key=lambda row: (
            row.get("lead", "") != "t",
            row.get("lead_ext", "") != row.get("name", ""),
            int(row["id"]),
        ),
    )
    for row in preferred:
        candidates.extend([row.get("pkg", ""), row.get("lead_ext", ""), row.get("name", "")])
    base = template.replace("$v", "").rstrip("_-")
    candidates.extend([base, base.replace("-", "_"), base.replace("_", "-")])
    for candidate in dict.fromkeys(item for item in candidates if item):
        if candidate in inventory:
            return candidate, "catalog_identity"
    return None, "unmatched"


def is_noarch(fmt: str, recipe: str, rpm_repo: pathlib.Path, deb_repo: pathlib.Path) -> bool:
    if fmt == "rpm":
        text = (rpm_repo / recipe_path(fmt, recipe)).read_text(errors="replace")
        return re.search(r"(?mi)^\s*BuildArch:\s*noarch\s*$", text) is not None

    root = deb_repo / recipe_path(fmt, recipe) / "debian"
    controls = list(root.glob("control*")) if root.is_dir() else []
    arch_values: list[str] = []
    for path in controls:
        if not path.is_file():
            continue
        arch_values.extend(
            match.group(1).strip().lower()
            for match in re.finditer(r"(?mi)^Architecture:\s*(\S+)", path.read_text(errors="replace"))
        )
    return bool(arch_values) and set(arch_values) == {"all"}


def source_list(
    fmt: str,
    template: str,
    recipe: str,
    rows: list[dict[str, str]],
    pg_versions: list[int],
    platform: str,
) -> tuple[list[str], list[str]]:
    raw_values = sorted({row["source"].strip() for row in rows if row["source"].strip()})
    conflicts = raw_values if len(raw_values) > 1 else []

    if recipe == "age":
        sources = []
        if 17 in pg_versions and fmt == "deb":
            sources.append("age-1.7.0.tar.gz")
        if 18 in pg_versions:
            sources.append("age-PG18-v1.8.0-rc0.tar.gz")
        return sources, conflicts
    if recipe == "babelfish":
        source_by_pg = {
            17: "babelfish-17-17.10-5.7.0.tar.gz",
            18: "babelfish-18-18.4-6.2.0.tar.gz",
        }
        return [source_by_pg[pg] for pg in pg_versions if pg in source_by_pg], conflicts
    if recipe == "datasketches":
        return [
            "apache-datasketches-postgresql-1.7.0-src.tar.gz",
            "apache-datasketches-cpp-5.2.0-src.tar.gz",
        ], conflicts
    if recipe == "duckdb_fdw":
        return [
            "duckdb_fdw-2.0.1+git20260529.9354241.tar.gz",
            "duckdb-1.5.5-headers.tar.gz",
        ], conflicts
    if recipe == "etcd_fdw":
        return ["etcd_fdw-0.0.1.tar.gz", "wrappers-0.6.2.tar.gz"], conflicts
    if recipe == "orioledb":
        source_by_pg = {
            16: "postgres-patches16_47.tar.gz",
            17: "postgres-patches17_20.tar.gz",
            18: "postgres-patches18_1.tar.gz",
        }
        return ["orioledb-beta16.tar.gz"] + [
            source_by_pg[pg] for pg in pg_versions if pg in source_by_pg
        ], conflicts
    if recipe == "pg_ducklake":
        return ["pg_ducklake-1.0.2.tar.gz", "CRoaring-4.7.1-amalgamation.tar.gz"], conflicts
    if recipe == "pg_net":
        legacy = platform.startswith(("el8.", "el9.")) if fmt == "rpm" else platform.startswith("u22.")
        return ["pg_net-0.9.2.tar.gz" if legacy else "pg_net-0.20.5.tar.gz"], conflicts
    if recipe == "pg_search" and fmt == "rpm":
        return ["pg_search-0.25.6.tar.gz", "pg_search_collect_third_party_licenses.py"], conflicts
    if recipe == "pgtde":
        return [
            "percona-postgresql-18.6.tar.gz",
            "percona-pg_tde18-2.2.2.tar.gz",
            "pgtde-pg-config",
            "percona-postgis-3.5.7.tar.gz",
            "percona-pgvector_18-0.8.6.tar.gz",
            "percona-wal2json-2.6.tar.gz",
            "percona-pg_repack-1.5.3.tar.gz",
            "percona-pgaudit-18.0.tar.gz",
            "percona-pgaudit18_set_user-4.2.0.tar.gz",
            "percona-pg-stat-monitor18-2.3.2.tar.gz",
            "percona-pg_gather-33.tar.gz",
            "pgtde-sfcgal-config",
        ], conflicts
    if recipe == "rdkit":
        if fmt == "rpm":
            return ["rdkit_202603.6.orig.tar.xz", "better-enums-0.11.3-enum.h"], conflicts
        if platform.startswith("d12."):
            return [
                "rdkit_202303.3-3.pgdg120+1.dsc",
                "rdkit_202303.3.orig.tar.xz",
                "rdkit_202303.3-3.pgdg120+1.debian.tar.xz",
            ], conflicts
        if platform.startswith("u22."):
            return [
                "rdkit_202303.3-3.pgdg22.04+1.dsc",
                "rdkit_202303.3.orig.tar.xz",
                "rdkit_202303.3-3.pgdg22.04+1.debian.tar.xz",
            ], conflicts
        if platform.startswith("u26."):
            return [
                "rdkit_202603.6-1PGSTY~resolute.dsc",
                "rdkit_202603.6.orig.tar.xz",
                "rdkit_202603.6-1PGSTY~resolute.debian.tar.xz",
            ], conflicts
    if recipe == "pg_statviz":
        return ["pg_statviz-1.2.1.tar.gz"], conflicts
    if (fmt, template) in SOURCE_OVERRIDES:
        return list(SOURCE_OVERRIDES[(fmt, template)]), conflicts
    if template == "pgedge-$v":
        pg_version_map = {15: "15.19", 16: "16.15", 17: "17.11", 18: "18.6"}
        sources = [f"postgresql-{pg_version_map[pg]}.tar.gz" for pg in pg_versions if pg in pg_version_map]
        sources.extend(["spock-5.0.11.tar.gz", "lolor-1.2.2.tar.gz", "snowflake-2.6.0.tar.gz"])
        return sources, conflicts
    if template in {"omnigres_$v", "postgresql-$v-omnigres"}:
        return ["omnigres-20260212.tar.gz"], conflicts

    preferred = sorted(
        rows,
        key=lambda row: (
            row.get("lead", "") != "t",
            row.get("lead_ext", "") != row.get("name", ""),
            int(row["id"]),
        ),
    )
    chosen = next((row["source"].strip() for row in preferred if row["source"].strip()), "")
    return chosen.split() if chosen else [], conflicts


def dependency_group(recipe: str) -> str:
    groups = {
        "graphblas": "graphblas",
        "lagraph": "graphblas",
        "onesparse": "graphblas",
        "scws": "zhparser",
        "zhparser": "zhparser",
        "libfq": "libfq",
        "firebird_fdw": "libfq",
        "pgactive": "libpgfeutils",
        "pgsodium-libsodium": "pgsodium",
        "pgsodium": "pgsodium",
        "inchi": "rdkit",
        "rdkit": "rdkit",
        "antlr4-runtime413": "babelfish",
        "antlr4_runtime413": "babelfish",
        "babelfish": "babelfish",
        "zlog": "polarstore",
        "polarstore": "polarstore",
    }
    if recipe.startswith("libpgfeutils"):
        return "libpgfeutils"
    return groups.get(recipe, recipe)


def build_target(fmt: str, recipe: str, pg_versions: list[int]) -> str:
    if fmt != "deb":
        return "-"
    if recipe == "age":
        return ",".join([*(f"build{pg}" for pg in pg_versions), "move"])
    if recipe == "babelfish":
        return ",".join(f"build{pg}" for pg in pg_versions)
    if recipe == "pg_duckdb":
        return ",".join(f"build{pg}" for pg in pg_versions)
    if recipe == "orioledb":
        return ",".join(f"orioledb-{pg}" for pg in pg_versions)
    if recipe == "pgtde":
        return "pgtde-18"
    if recipe == "pgedge":
        return ",".join(f"pgedge-{pg}" for pg in pg_versions)
    return "-"


def classify(recipe: str, rows: list[dict[str, str]], noarch: bool, pgdg_gap: bool) -> str:
    # Required collision priority: kernel/tools > support > heavy > PGDG-gap >
    # Rust > noarch > standard.
    if recipe in KERNEL_RECIPES:
        return "kernel_tools"
    if recipe in SUPPORT_RECIPES:
        return "support"
    if recipe in HEAVY_RECIPES or any(row.get("lang") == "C++" for row in rows):
        return "heavy"
    # pg_statviz gap artifacts are pure SQL and byte-identical across
    # architectures; canonical noarch/all dedupe takes precedence for it.
    if recipe == "pg_statviz" and noarch:
        return "noarch"
    if pgdg_gap:
        return "pgdg_gap"
    if any(row.get("lang") == "Rust" or '"pgrx"' in row.get("extra", "") for row in rows):
        return "rust"
    if noarch:
        return "noarch"
    return "standard"


def batch_for(build_class: str, recipe: str) -> tuple[str, str]:
    if build_class == "support":
        return "B10", "explicit dependency order inside dependency_group"
    if build_class == "noarch":
        return "B20", "canonical aarch64 builder only"
    if build_class == "standard":
        shard = int(hashlib.sha256(recipe.encode()).hexdigest(), 16) % 2
        return ("B30" if shard == 0 else "B31"), "sha256(recipe UTF-8) mod 2"
    if build_class == "rust":
        return "B40", "rust/pgrx"
    if build_class == "pgdg_gap":
        return "B50", "PGDG metadata gap"
    if build_class == "heavy":
        return "B60", "heavy C/C++"
    return "B70", "kernel/tools"


def debug_policy(fmt: str, recipe: str, noarch: bool) -> str:
    if noarch or (fmt == "deb" and recipe in {"pgsodium-libsodium", "polarstore"}):
        return "N/A:no-native-ELF"
    if fmt == "rpm":
        return "required:non-empty-debuginfo+debugsource"
    return "required:non-empty-dbgsym"


def release_policy(fmt: str, recipe: str, platform: str, rows: list[dict[str, str]], pgdg_gap: bool) -> str:
    if recipe == "pg_net":
        if fmt == "deb" and platform.startswith("u22."):
            return "legacy 0.9.2-2PGSTY for Jammy"
        if fmt == "deb":
            return "current 0.20.5-2PGSTY"
        if platform.startswith(("el8.", "el9.")):
            return "legacy 0.9.2 with monotonic PGSTY release"
        return "current 0.20.5 with monotonic PGSTY release"
    if recipe == "rdkit":
        if fmt == "rpm":
            return "202603.6-1PGSTY on EL9/EL10; excluded on EL8"
        if platform.startswith(("d12.", "u22.")):
            return "legacy 202303.3 cartridge gap for PG17-18"
        return "202603.6-1PGSTY on Resolute PG14-17"
    if recipe == "pg_statviz":
        return "PGDG-gap only; package version=1.2.1; extension version=1.2; current release must end PGSTY"
    versions = sorted({row.get(f"{fmt}_ver", "") for row in rows if row.get(f"{fmt}_ver", "")})
    prefix = "PGDG-gap only; " if pgdg_gap else ""
    return prefix + "version=" + ",".join(versions or ["recipe-defined"]) + "; current release must end PGSTY"


def sha256_file(path: pathlib.Path) -> str:
    digest = hashlib.sha256()
    with path.open("rb") as stream:
        for chunk in iter(lambda: stream.read(1024 * 1024), b""):
            digest.update(chunk)
    return digest.hexdigest()


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--run-id", required=True)
    parser.add_argument("--rpm-repo", type=pathlib.Path, default=pathlib.Path.home() / "pgsty/rpm")
    parser.add_argument("--deb-repo", type=pathlib.Path, default=pathlib.Path.home() / "pgsty/deb")
    parser.add_argument("--pgext-repo", type=pathlib.Path, default=pathlib.Path.home() / "pgsty/pgext")
    parser.add_argument("--source-root", type=pathlib.Path, default=pathlib.Path.home() / "pgsty/repo/ext/src")
    parser.add_argument("--output", type=pathlib.Path, required=True)
    args = parser.parse_args()

    args.output.mkdir(parents=True, exist_ok=True)
    catalog_path = args.pgext_repo / "db/extension.csv"
    catalog_bytes = catalog_path.read_bytes()
    catalog_csv_sha256 = hashlib.sha256(catalog_bytes).hexdigest()
    catalog_db_export, catalog_db_snapshot_at = load_catalog_db_export()
    catalog_db_sha256 = hashlib.sha256(catalog_db_export).hexdigest()
    catalog_matches_db = catalog_bytes == catalog_db_export
    pgext_status = git_status(args.pgext_repo)

    with catalog_path.open(newline="") as stream:
        catalog = [row for row in csv.DictReader(stream) if row["state"] == "available"]

    pgdg_available, repo_stamps = load_pgdg_state()
    repo_sha = {
        "rpm": git_head(args.rpm_repo),
        "deb": git_head(args.deb_repo),
        "pgext": git_head(args.pgext_repo),
    }
    manifests: list[dict[str, Any]] = []
    unmatched: list[dict[str, Any]] = []
    alias_records: list[dict[str, Any]] = []
    source_conflicts: list[dict[str, Any]] = []
    referenced_recipes: dict[str, set[str]] = {"rpm": set(), "deb": set()}
    catalog_recipes: dict[str, set[str]] = {"rpm": set(), "deb": set()}
    pgdg_covered_recipes: dict[str, set[str]] = {"rpm": set(), "deb": set()}
    rule_effects: collections.Counter[str] = collections.Counter()

    if not catalog_matches_db:
        unmatched.append(
            {
                "format": "catalog",
                "package": "pgext.extension",
                "catalog_repo": "PGEXT",
                "catalog_extensions": [],
                "reason": "db/extension.csv does not byte-match the ordered live pgext.extension export",
            }
        )

    for fmt, platforms in (("rpm", RPM_PLATFORMS), ("deb", DEB_PLATFORMS)):
        inventory = recipe_inventory(fmt, args.rpm_repo, args.deb_repo)
        for side_repo in ("PIGSTY", "PGDG"):
            groups: dict[str, list[dict[str, str]]] = collections.defaultdict(list)
            for row in catalog:
                if row[f"{fmt}_repo"] == side_repo and row[f"{fmt}_pkg"]:
                    groups[row[f"{fmt}_pkg"]].append(row)

            for template, rows in sorted(groups.items()):
                pg_versions = sorted(
                    set().union(*(parse_pg_array(row[f"{fmt}_pg"]) for row in rows))
                )
                os_exclude: set[str] = set()
                for row in rows:
                    extra = json.loads(row.get("extra") or "{}")
                    os_exclude.update(extra.get("os_exclude", []))
                # Package ownership is format-specific.  A generic catalog
                # repo=PGDG value must never override rpm_repo/deb_repo=PIGSTY.
                pgdg_gap = side_repo == "PGDG"

                recipe, mapping_kind = map_recipe(fmt, template, rows, inventory)
                if recipe is None:
                    # A fully covered PGDG coordinate does not need a local
                    # recipe.  A real metadata gap without a recipe is a gate
                    # blocker and must not disappear from the unmatched list.
                    has_gap = side_repo == "PIGSTY"
                    if side_repo == "PGDG":
                        has_gap = any(
                            platform not in os_exclude
                            and any(
                                not any((row["name"], platform, pg) in pgdg_available for row in rows)
                                for pg in pg_versions
                            )
                            for platform in platforms
                        )
                    if has_gap:
                        unmatched.append(
                            {
                                "format": fmt,
                                "package": template,
                                "catalog_repo": side_repo,
                                "catalog_extensions": sorted(row["name"] for row in rows),
                                "reason": f"{side_repo} package has a build coordinate but no active spec/recipe or explicit alias",
                            }
                        )
                    continue

                catalog_recipes[fmt].add(recipe)
                if mapping_kind == "explicit_alias":
                    alias_records.append(
                        {
                            "format": fmt,
                            "package": template,
                            "recipe": recipe,
                            "extensions": sorted(row["name"] for row in rows),
                            "note": ALIAS_NOTES.get(template, "explicit product alias"),
                        }
                    )

                noarch = is_noarch(fmt, recipe, args.rpm_repo, args.deb_repo)
                jobs_before = len(manifests)
                for platform in platforms:
                    if platform in os_exclude:
                        rule_effects[f"{fmt}:os_exclude_jobs"] += 1
                        rule_effects[f"{fmt}:os_exclude_pg_abi"] += len(pg_versions)
                        continue
                    selected_pg = list(pg_versions)
                    names = {row["name"] for row in rows}

                    if recipe == "rdkit" and side_repo == "PIGSTY":
                        if fmt == "rpm":
                            if platform.startswith("el8."):
                                rule_effects["rpm:rdkit_platform_jobs_removed"] += 1
                                rule_effects["rpm:rdkit_pg_abi_removed"] += len(selected_pg)
                                continue
                            selected_pg = [14, 15, 16, 17, 18]
                        else:
                            if platform.startswith(("d13.", "u24.")):
                                rule_effects["deb:rdkit_pgdg_platform_jobs_removed"] += 1
                                rule_effects["deb:rdkit_pgdg_pg_abi_removed"] += len(selected_pg)
                                continue
                            selected_pg = [17, 18] if platform.startswith(("d12.", "u22.")) else [14, 15, 16, 17]

                    if pgdg_gap:
                        before_pgdg_filter = list(selected_pg)
                        selected_pg = [
                            pg
                            for pg in selected_pg
                            if not any((row["name"], platform, pg) in pgdg_available for row in rows)
                        ]
                        rule_effects[f"{fmt}:pgdg_covered_pg_abi_removed"] += (
                            len(before_pgdg_filter) - len(selected_pg)
                        )
                    if pg_versions and not selected_pg:
                        rule_effects[f"{fmt}:empty_after_policy_jobs_removed"] += 1
                        continue
                    if noarch and platform.endswith(".x86_64"):
                        rule_effects[f"{fmt}:canonical_noarch_jobs_removed"] += 1
                        rule_effects[f"{fmt}:canonical_noarch_pg_abi_removed"] += len(selected_pg)
                        continue

                    pg_slices = (
                        [[pg] for pg in selected_pg]
                        if fmt == "deb" and recipe in DEB_SINGLE_PG_RECIPES
                        else [selected_pg]
                    )
                    for job_pg in pg_slices:
                        sources, conflicts = source_list(
                            fmt, template, recipe, rows, job_pg, platform
                        )
                        if conflicts:
                            conflict_record = {
                                "format": fmt,
                                "package": template,
                                "recipe": recipe,
                                "pg_versions": job_pg,
                                "catalog_sources": conflicts,
                                "selected_sources": sources,
                            }
                            if conflict_record not in source_conflicts:
                                source_conflicts.append(conflict_record)
                        build_class = classify(recipe, rows, noarch, pgdg_gap)
                        batch, shard_policy = batch_for(build_class, recipe)
                        manifests.append(
                            {
                                "format": fmt,
                                "platform": platform,
                                "arch": platform.split(".", 1)[1],
                                "package": template,
                                "recipe": recipe_path(fmt, recipe),
                                "pg_versions": job_pg,
                                "build_class": build_class,
                                "batch": batch,
                                "source": sources,
                                "source_sha256": {},
                                "source_provenance": {},
                                "repo_sha": repo_sha[fmt],
                                "dependency_group": dependency_group(recipe),
                                "release_policy": release_policy(fmt, recipe, platform, rows, pgdg_gap),
                                "debug_policy": debug_policy(fmt, recipe, noarch),
                                "artifact_guard": "bin/verify-extension-job.sh run" if fmt == "deb" else "",
                                "artifact_package_template": template,
                                "job_kind": "extension",
                                "build_target": build_target(fmt, recipe, job_pg),
                                "status": "pending_gate",
                                "origin": "catalog_pgdg_gap" if side_repo == "PGDG" else "catalog_pigsty",
                                "catalog_extensions": sorted(names),
                                "mapping": mapping_kind,
                                "os_exclude": sorted(os_exclude),
                                "shard_policy": shard_policy,
                            }
                        )

                if len(manifests) > jobs_before:
                    referenced_recipes[fmt].add(recipe)
                elif side_repo == "PGDG":
                    pgdg_covered_recipes[fmt].add(recipe)

        # Explicitly allowlisted support components that are absent from the
        # extension catalog.  Existing catalog recipes are classified as
        # support above and are not duplicated here.
        for package, recipe, pgs, sources, allowed_platforms, dep_group in SUPPORT[fmt]:
            if recipe in catalog_recipes[fmt]:
                continue
            if recipe not in inventory:
                unmatched.append(
                    {
                        "format": fmt,
                        "package": package,
                        "catalog_repo": "explicit_B10",
                        "catalog_extensions": [],
                        "reason": "explicit B10 support recipe missing",
                    }
                )
                continue
            referenced_recipes[fmt].add(recipe)
            for platform in platforms:
                if allowed_platforms is not None and platform not in allowed_platforms:
                    continue
                rule_effects[f"{fmt}:explicit_support_jobs_added"] += 1
                rule_effects[f"{fmt}:explicit_support_pg_abi_added"] += len(pgs)
                manifests.append(
                    {
                        "format": fmt,
                        "platform": platform,
                        "arch": platform.split(".", 1)[1],
                        "package": package,
                        "recipe": recipe_path(fmt, recipe),
                        "pg_versions": pgs,
                        "build_class": "support",
                        "batch": "B10",
                        "source": sources,
                        "source_sha256": {},
                        "source_provenance": {},
                        "repo_sha": repo_sha[fmt],
                        "dependency_group": dep_group,
                        "release_policy": "explicit B10 support; current release must end PGSTY",
                        "debug_policy": debug_policy(fmt, recipe, noarch=False),
                        "artifact_guard": "bin/verify-extension-job.sh run" if fmt == "deb" else "",
                        "artifact_package_template": (
                            DEB_SUPPORT_PRIMARY.get(recipe, package) if fmt == "deb" else package
                        ),
                        "job_kind": "support",
                        "build_target": "-",
                        "status": "pending_gate",
                        "origin": "explicit_support",
                        "catalog_extensions": [],
                        "mapping": "explicit_B10_allowlist",
                        "os_exclude": [],
                        "shard_policy": "dependency ordered",
                    }
                )

    manifest_key_counts = collections.Counter(
        (row["format"], row["platform"], row["package"], tuple(row["pg_versions"]))
        for row in manifests
    )
    for key, count in sorted(manifest_key_counts.items()):
        if count > 1:
            unmatched.append(
                {
                    "format": key[0],
                    "package": key[2],
                    "catalog_repo": "mixed",
                    "catalog_extensions": [],
                    "reason": f"duplicate manifest key on {key[1]} PG{','.join(map(str, key[3])) or '-'} appears {count} times",
                }
            )

    # Hash each selected source once, then attach hashes to all manifest rows.
    selected_sources = sorted({source for row in manifests for source in row["source"]})
    source_hashes: dict[str, str | None] = {}
    source_provenance: dict[str, str] = {}
    for source in selected_sources:
        path = args.source_root / source
        provenance = f"source-root:{source}"
        if not path.is_file() and source == "pg_search_collect_third_party_licenses.py":
            path = args.rpm_repo / "rpmbuild/SPECS" / source
            provenance = f"rpm-repo:rpmbuild/SPECS/{source}"
        source_hashes[source] = sha256_file(path) if path.is_file() else None
        source_provenance[source] = provenance
    for row in manifests:
        row["source_sha256"] = {source: source_hashes[source] for source in row["source"]}
        row["source_provenance"] = {
            source: source_provenance[source] for source in row["source"]
        }
        if not row["source"]:
            row["status"] = "blocked_missing_source_field"
        elif any(source_hashes[source] is None for source in row["source"]):
            row["status"] = "blocked_missing_source_file"

    manifests.sort(key=lambda row: (row["format"], row["platform"], row["batch"], row["package"]))

    with (args.output / "build-manifest.jsonl").open("w") as stream:
        for row in manifests:
            stream.write(json.dumps(row, sort_keys=True, separators=(",", ":")) + "\n")

    csv_fields = [
        "format",
        "platform",
        "arch",
        "package",
        "recipe",
        "pg_versions",
        "build_class",
        "batch",
        "source",
        "source_sha256",
        "source_provenance",
        "repo_sha",
        "dependency_group",
        "release_policy",
        "debug_policy",
        "artifact_guard",
        "artifact_package_template",
        "job_kind",
        "build_target",
        "status",
        "origin",
        "catalog_extensions",
        "mapping",
        "shard_policy",
    ]
    with (args.output / "build-manifest.csv").open("w", newline="") as stream:
        writer = csv.DictWriter(stream, fieldnames=csv_fields)
        writer.writeheader()
        for row in manifests:
            flat = {key: row[key] for key in csv_fields}
            for key in ("pg_versions", "source", "catalog_extensions"):
                flat[key] = " ".join(map(str, flat[key]))
            flat["source_sha256"] = " ".join(
                f"{source}={digest or 'MISSING'}" for source, digest in sorted(row["source_sha256"].items())
            )
            flat["source_provenance"] = " ".join(
                f"{source}={location}"
                for source, location in sorted(row["source_provenance"].items())
            )
            writer.writerow(flat)

    with (args.output / "source-sha256.tsv").open("w") as stream:
        stream.write("source\tsha256\tstatus\tprovenance\n")
        for source, digest in sorted(source_hashes.items()):
            stream.write(
                f"{source}\t{digest or ''}\t{'present' if digest else 'missing'}"
                f"\t{source_provenance[source]}\n"
            )

    classifications: dict[str, list[dict[str, str]]] = {
        "pgdg_covered": [],
        "catalog_not_scheduled": [],
        "linux_owned": [],
        "explicit_exclude": [],
        "superseded": [],
        "explicit_quarantine": [],
        "unlisted_quarantine": [],
    }
    for fmt in ("rpm", "deb"):
        inventory = recipe_inventory(fmt, args.rpm_repo, args.deb_repo)
        for recipe in sorted(inventory - referenced_recipes[fmt]):
            item = {"format": fmt, "recipe": recipe_path(fmt, recipe)}
            if recipe in pgdg_covered_recipes[fmt]:
                item["reason"] = "refreshed PGDG metadata covers every declared coordinate"
                classifications["pgdg_covered"].append(item)
            elif recipe in catalog_recipes[fmt]:
                item["reason"] = "catalog mapping exists but all coordinates were removed by explicit platform/PG policy"
                classifications["catalog_not_scheduled"].append(item)
            elif recipe in LINUX_OWNED:
                classifications["linux_owned"].append(item)
            elif recipe in EXPLICIT_EXCLUDE:
                classifications["explicit_exclude"].append(item)
            elif recipe in SUPERSEDED_RECIPES:
                item["reason"] = SUPERSEDED_RECIPES[recipe]
                classifications["superseded"].append(item)
            elif recipe in EXPLICIT_QUARANTINE:
                classifications["explicit_quarantine"].append(item)
            else:
                item["reason"] = "active recipe/spec has no catalog package-template or explicit B10 product mapping"
                classifications["unlisted_quarantine"].append(item)

    (args.output / "unmatched.json").write_text(json.dumps(unmatched, indent=2, sort_keys=True) + "\n")
    (args.output / "aliases.json").write_text(json.dumps(alias_records, indent=2, sort_keys=True) + "\n")
    (args.output / "quarantine.json").write_text(json.dumps(classifications, indent=2, sort_keys=True) + "\n")
    (args.output / "source-conflicts.json").write_text(
        json.dumps(source_conflicts, indent=2, sort_keys=True) + "\n"
    )
    (args.output / "external-dependencies.json").write_text(
        json.dumps(EXTERNAL_DEPENDENCIES, indent=2, sort_keys=True) + "\n"
    )

    catalog_snapshot = {
        "pgext_head": repo_sha["pgext"],
        "catalog_path": "db/extension.csv",
        "catalog_csv_sha256": catalog_csv_sha256,
        "catalog_db_export_sha256": catalog_db_sha256,
        "catalog_matches_db": catalog_matches_db,
        "catalog_rows": max(0, len(catalog_db_export.splitlines()) - 1),
        "database_snapshot_at": catalog_db_snapshot_at,
        "worktree_dirty": bool(pgext_status),
        "worktree_status_sha256": hashlib.sha256(pgext_status.encode()).hexdigest(),
        "worktree_status": pgext_status.splitlines(),
    }
    (args.output / "catalog-snapshot.json").write_text(
        json.dumps(catalog_snapshot, indent=2, sort_keys=True) + "\n"
    )

    by_format = collections.Counter(row["format"] for row in manifests)
    abi_by_format = collections.Counter()
    for row in manifests:
        abi_by_format[row["format"]] += len(row["pg_versions"])
    summary = {
        "run_id": args.run_id,
        "production_gate": "closed",
        "repo_sha": repo_sha,
        "pgext_catalog_head": repo_sha["pgext"],
        "pgext_catalog": {
            key: value
            for key, value in catalog_snapshot.items()
            if key != "worktree_status"
        },
        "generator_sha256": sha256_file(pathlib.Path(__file__)),
        "pgdg_snapshot": repo_stamps,
        "rules": {
            "pg_versions": [14, 15, 16, 17, 18],
            "dedupe_key": "format+platform+package-template+PG-major-set; expanded artifact coordinates remain unique",
            "canonical_noarch_arch": "aarch64",
            "standard_shard": "sha256(recipe UTF-8) mod 2; B30=0, B31=1",
            "rdkit": "RPM EL9/10 202603.6; DEB D12/U22 legacy PG17-18, D13/U24 PGDG, U26 202603.6 PG14-17",
            "pg_net": "DEB Jammy 0.9.2; D12/D13/U24/U26 0.20.5",
        },
        "jobs": dict(by_format),
        "pg_abi_units": dict(abi_by_format),
        "jobs_by_origin": dict(collections.Counter((row["format"], row["origin"]) for row in manifests)),
        "jobs_by_class": dict(collections.Counter((row["format"], row["build_class"]) for row in manifests)),
        "jobs_by_batch": dict(collections.Counter((row["format"], row["batch"]) for row in manifests)),
        "status": dict(collections.Counter((row["format"], row["status"]) for row in manifests)),
        "rule_effects": dict(sorted(rule_effects.items())),
        "source_files": len(source_hashes),
        "source_files_missing": sorted(source for source, digest in source_hashes.items() if digest is None),
        "catalog_unmatched_count": len([item for item in unmatched if item["catalog_extensions"]]),
        "pigsty_unmatched_count": len(
            [item for item in unmatched if item.get("catalog_repo") == "PIGSTY"]
        ),
        "pgdg_gap_unmatched_count": len(
            [item for item in unmatched if item.get("catalog_repo") == "PGDG"]
        ),
        "support_unmatched_count": len([item for item in unmatched if not item["catalog_extensions"]]),
        "baseline": {
            "rpm": {"jobs": 1828, "pg_abi_units": 8347},
            "deb": {"jobs": 3114, "pg_abi_units": 14352},
        },
        "baseline_delta": {
            fmt: {
                "jobs": by_format[fmt] - ({"rpm": 1828, "deb": 3114}[fmt]),
                "pg_abi_units": abi_by_format[fmt] - ({"rpm": 8347, "deb": 14352}[fmt]),
                "rule_effects": {
                    key: value
                    for key, value in sorted(rule_effects.items())
                    if key.startswith(f"{fmt}:")
                },
                "explanation": "baseline delta is traced by the counters above plus catalog snapshot drift; counters are directional rule effects and are not asserted to sum across overlapping rules",
            }
            for fmt in ("rpm", "deb")
        },
    }

    # JSON object keys cannot be tuples; normalize the nested counters.
    for key in ("jobs_by_origin", "jobs_by_class", "jobs_by_batch", "status"):
        counter = summary[key]
        summary[key] = {f"{fmt}:{value}": count for (fmt, value), count in counter.items()}
    (args.output / "summary.json").write_text(json.dumps(summary, indent=2, sort_keys=True) + "\n")

    if unmatched or summary["source_files_missing"]:
        return 2
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
