#!/usr/bin/env python3
"""
Generate PostgreSQL Extension Index Pages (_index.md and _index.zh.md)

This script generates two tables:
1. Extensions table - listing all extensions with their PG versions, attributes, and descriptions
2. Packages table - listing all packages with version, tags, and repository URLs

Usage:
    python ext-index.py [PGURL] [OUTPUT_DIR]

    PGURL: PostgreSQL connection URL (default: postgres:///)
    OUTPUT_DIR: Output directory for index files (default: content/e/)
"""

import argparse
import psycopg2
from pathlib import Path
from typing import List, Dict, Optional, Any
from dataclasses import dataclass


# ============================================================
# Data Models
# ============================================================
@dataclass
class Extension:
    """Extension with all metadata for the index table"""
    id: int
    name: str
    pkg: str
    category: str
    state: str
    url: str
    license: str
    lang: str
    version: str
    pg_ver: List[int]
    contrib: bool
    lead: bool
    has_bin: bool
    has_lib: bool
    need_ddl: bool
    need_load: bool
    trusted: bool
    relocatable: bool
    en_desc: str
    zh_desc: str


@dataclass
class Package:
    """Package with metadata for the package table"""
    id: int
    pkg: str
    lead_ext: str
    version: str
    category: str
    license: str
    lang: str
    url: str
    rpm_pkg: str
    deb_pkg: str


# ============================================================
# Database Operations
# ============================================================
def fetch_active_pg_versions(conn) -> List[int]:
    """Fetch active PostgreSQL major versions from pgext.active_pg view"""
    with conn.cursor() as cur:
        cur.execute("""SELECT pg FROM pgext.active_pg ORDER BY pg DESC""")
        return [row[0] for row in cur.fetchall()]

def parse_array(value: Any) -> List[str]:
    """Parse PostgreSQL array to Python list"""
    if isinstance(value, list):
        return value
    if not value or not isinstance(value, str):
        return []
    if not value.startswith('{') or not value.endswith('}'):
        return []
    items = []
    for raw in value[1:-1].split(','):
        item = raw.strip()
        if len(item) >= 2 and item[0] == item[-1] == '"':
            item = item[1:-1]
        if item:
            items.append(item)
    return items


def parse_pg_versions(values: List[str]) -> List[int]:
    """Convert a list of version strings to sorted integer PG major versions"""
    pg_versions = []
    for val in values:
        stripped = val.strip()
        if stripped.endswith('+'):
            stripped = stripped[:-1]
        try:
            pg_versions.append(int(stripped))
        except ValueError:
            continue
    return sorted(set(pg_versions), reverse=True)


def fetch_extensions(conn) -> List[Extension]:
    """Fetch all extensions for the Extensions table"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT id, name, pkg, category, state, url, license, lang, version,
                   pg_ver, contrib, lead, has_bin, has_lib, need_ddl, need_load,
                   trusted, relocatable, en_desc, zh_desc
            FROM pgext.extension
            WHERE state != 'not-ready'
            ORDER BY id
        """)

        extensions = []
        for row in cur.fetchall():
            (id, name, pkg, category, state, url, license, lang, version,
             pg_ver_raw, contrib, lead, has_bin, has_lib, need_ddl, need_load,
             trusted, relocatable, en_desc, zh_desc) = row

            pg_ver = parse_pg_versions(parse_array(pg_ver_raw))

            extensions.append(Extension(
                id=id, name=name, pkg=pkg, category=category or "UNKNOWN",
                state=state, url=url or "", license=license or "Unknown",
                lang=lang or "Unknown", version=version or "Unknown",
                pg_ver=pg_ver, contrib=contrib, lead=lead,
                has_bin=has_bin, has_lib=has_lib, need_ddl=need_ddl,
                need_load=need_load, trusted=trusted, relocatable=relocatable,
                en_desc=en_desc or "", zh_desc=zh_desc or ""
            ))

        return extensions


def fetch_packages(conn) -> List[Package]:
    """Fetch unique packages (lead extensions only) for the Packages table, sorted by id"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT id, pkg, name, version, category, license, lang, url, rpm_pkg, deb_pkg
            FROM pgext.extension
            WHERE lead = true AND state != 'not-ready'
            ORDER BY id
        """)

        packages = []
        for row in cur.fetchall():
            (id, pkg, lead_ext, version, category, license, lang, url, rpm_pkg, deb_pkg) = row

            packages.append(Package(
                id=id, pkg=pkg, lead_ext=lead_ext, version=version or "Unknown",
                category=category or "UNKNOWN", license=license or "Unknown",
                lang=lang or "Unknown", url=url or "",
                rpm_pkg=rpm_pkg or "", deb_pkg=deb_pkg or ""
            ))

        return packages


# ============================================================
# Shortcode Helpers
# ============================================================
def ext_shortcode(name: str, label: Optional[str] = None) -> str:
    """Generate ext shortcode: {{< ext "name" "label" >}}"""
    if label and label != name:
        return f'{{{{< ext "{name}" "{label}" >}}}}'
    return f'{{{{< ext "{name}" >}}}}'


def generate_attribute_badge(ext: Extension) -> str:
    """Generate attribute badge string like '---s-d-r'"""
    attrs = ''.join([
        "c" if ext.contrib else "-",
        "b" if ext.has_bin else "-",
        "s" if ext.has_lib else "-",
        "L" if ext.need_load else "-",
        "d" if ext.need_ddl else "-",
        "t" if ext.trusted else "-",
        "r" if ext.relocatable else "-"
    ])
    return f'{{{{< badge content="{attrs}" color="blue" >}}}}'


def generate_pg_badges(supported_versions: List[int], active_versions: List[int]) -> str:
    """Generate PostgreSQL version badges using pgver shortcode"""
    if not active_versions:
        return '{{< badge content="N/A" color="gray" >}}'

    # Build comma-separated version list
    versions = ','.join(str(pg) for pg in active_versions)

    # Build comma-separated color list (g=green for supported, r=red for unsupported)
    colors = ','.join('g' if pg in supported_versions else 'r' for pg in active_versions)

    return f'{{{{< pgver "{versions}" "{colors}" >}}}}'

# ============================================================
# Content Generation
# ============================================================
def sanitize_text(text: str) -> str:
    """Sanitize text for markdown table"""
    if not text:
        return ""
    return text.replace('|', '\\|').replace('\n', ' ').strip()


def generate_extensions_table(extensions: List[Extension], active_pg_versions: List[int], locale: str) -> str:
    """Generate Extensions table content with ID column"""
    lines = []

    # Header
    if locale == "en":
        lines.append("| Extension | PG Versions | Attribute | Description |")
        lines.append("|:----------|:------------|:---------:|:------------|")
    else:  # zh
        lines.append("| 扩展 | PG 版本列表          | 属性 | 描述 |")
        lines.append("|:-----|:-------------------|:----:|:-----|")

    # Rows
    for ext in extensions:
        ext_link = ext_shortcode(ext.name, ext.pkg)
        pg_badges = generate_pg_badges(ext.pg_ver, active_pg_versions)
        attr_badge = generate_attribute_badge(ext)

        if locale == "en":
            desc = sanitize_text(ext.en_desc)[:100]
        else:
            desc = sanitize_text(ext.zh_desc or ext.en_desc)[:100]

        lines.append(f"| {ext_link} | {pg_badges} | {attr_badge} | {desc} |")

    return '\n'.join(lines)


def generate_packages_table(packages: List[Package], locale: str) -> str:
    """Generate Packages table content with RPM and DEB columns"""
    lines = []

    # Header
    if locale == "en":
        lines.append("| Package | Version | Repo | Category | RPM | DEB |")
        lines.append("|:--------|:--------|:-----|:---------|:-----|:-----|")
    else:  # zh
        lines.append("| 包 | 版本 | 仓库 | 分类 | RPM | DEB |")
        lines.append("|:---|:-----|:-----|:-----|:-----|:-----|")

    # Rows
    for pkg in packages:
        pkg_link = ext_shortcode(pkg.lead_ext, pkg.pkg)
        version = f"`{pkg.version}`"

        # URL badge
        if pkg.url:
            url_badge = f'{{{{< badge content="{"Link" if locale == "en" else "链接"}" link="{pkg.url}" >}}}}'
        else:
            url_badge = '{{< badge content="N/A" color="gray" >}}'

        # Category badge
        cate_badge = f'{{{{< category "{pkg.category.upper()}" >}}}}' if pkg.category else '-'

        # RPM and DEB packages
        rpm_display = f"`{pkg.rpm_pkg}`" if pkg.rpm_pkg else "-"
        deb_display = f"`{pkg.deb_pkg}`" if pkg.deb_pkg else "-"

        lines.append(f"| {pkg_link} | {version} | {url_badge} | {cate_badge} | {rpm_display} | {deb_display} |")

    return '\n'.join(lines)


def generate_content(extensions: List[Extension], packages: List[Package], active_pg_versions: List[int], locale: str) -> str:
    """Generate complete content for the index page"""
    ext_count = len(extensions)
    pkg_count = len(packages)
    is_zh = locale == "zh"

    # Locale-specific content
    content = {
        'en': {
            'title': 'Extensions',
            'desc': 'PostgreSQL extensions enhance the database with additional functionality.',
            'pkg_header': f'\n## Packages\n\nThere are {pkg_count} available PostgreSQL packages:\n',
            'ext_header': f'\n## Extensions\n\nThere are {ext_count} available PostgreSQL extensions:\n',
            'attr_legend': '\n**Attribute Legend**: `c`:contrib `b`:bin `s`:lib `l`:load `d`:ddl `t`:trusted `r`:relocatable'
        },
        'zh': {
            'title': '扩展',
            'desc': 'PostgreSQL 扩展通过额外的功能增强数据库。',
            'pkg_header': f'\n## 包\n\n共有 {pkg_count} 个可用的 PostgreSQL 包：\n',
            'ext_header': f'\n## 扩展\n\n共有 {ext_count} 个可用的 PostgreSQL 扩展：\n',
            'attr_legend': '\n**属性说明**: `c`:contrib `b`:二进制 `s`:动态库 `l`:需加载 `d`:需DDL `t`:无需特权 `r`:可重定位'
        }
    }

    loc = content[locale]

    # Build content
    lines = [
        f"""---
title: "{loc['title']}"
breadcrumbs: false
excludeSearch: true
comments: false
weight: 900
---

{loc['desc']}
""",
        loc['pkg_header'],
        generate_packages_table(packages, locale),
        loc['ext_header'],
        generate_extensions_table(extensions, active_pg_versions, locale),
        loc['attr_legend']
    ]

    return '\n'.join(lines)


# ============================================================
# Main Logic
# ============================================================
def generate_index_pages(pgurl: str, output_dir: Path):
    """Generate both English and Chinese index pages"""
    # Connect to database
    conn = psycopg2.connect(pgurl)

    try:
        # Fetch data
        active_pg_versions = fetch_active_pg_versions(conn)
        extensions = fetch_extensions(conn)
        packages = fetch_packages(conn)

        # Generate English index
        en_content = generate_content(extensions, packages, active_pg_versions, "en")

        # Generate Chinese index
        zh_content = generate_content(extensions, packages, active_pg_versions, "zh")

        # Write files
        output_dir.mkdir(parents=True, exist_ok=True)
        (output_dir / "_index.md").write_text(en_content, encoding='utf-8')
        (output_dir / "_index.zh.md").write_text(zh_content, encoding='utf-8')

        print(f"✓ Generated: {output_dir}/_index.md")
        print(f"✓ Generated: {output_dir}/_index.zh.md")
        print(f"  Extensions: {len(extensions)}, Packages: {len(packages)}")

    finally:
        conn.close()


def main():
    """Main entry point"""
    parser = argparse.ArgumentParser(
        description='Generate extension index pages for Hugo'
    )
    parser.add_argument(
        'pgurl', nargs='?', default='postgres:///',
        help='PostgreSQL connection URL (default: postgres:///)'
    )
    parser.add_argument(
        'output_dir', nargs='?', default=None,
        help='Output directory (default: content/e/)'
    )

    args = parser.parse_args()

    # Determine output directory
    if args.output_dir:
        output_dir = Path(args.output_dir)
    else:
        output_dir = Path(__file__).parent.parent / 'content' / 'e'

    generate_index_pages(args.pgurl, output_dir)


if __name__ == "__main__":
    main()
