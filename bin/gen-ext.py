#!/usr/bin/env python3
"""
PostgreSQL Extension Hugo Content Generator

Generates Hugo/Hextra-compatible Markdown files for PostgreSQL extensions
by reading from the pgext database schema.

Usage:
    python generate.py [PGURL] [OUTPUT_DIR]

    PGURL: PostgreSQL connection URL (default: postgres:///)
    OUTPUT_DIR: Output directory for generated files (default: content/e/)
"""

import argparse
import psycopg2
from typing import List, Optional, Any, Dict, Tuple
from dataclasses import dataclass
from pathlib import Path
import logging

# Configure logging
logging.basicConfig(level=logging.INFO, format='%(levelname)s: %(message)s')
logger = logging.getLogger(__name__)


@dataclass
class ExtensionData:
    """Extension metadata"""
    id: int
    name: str
    pkg: str
    lead_ext: Optional[str]
    category: str
    state: str
    url: Optional[str]
    license: str
    tags: List[str]
    version: str
    repo: Optional[str]
    lang: str
    contrib: bool
    lead: bool
    has_bin: bool
    has_lib: bool
    need_ddl: bool
    need_load: bool
    trusted: bool
    relocatable: bool
    schemas: List[str]
    pg_ver: List[int]
    requires: List[str]
    require_by: List[str]
    see_also: List[str]
    rpm_ver: Optional[str]
    rpm_repo: Optional[str]
    rpm_pkg: Optional[str]
    rpm_pg: List[int]
    rpm_deps: List[str]
    deb_ver: Optional[str]
    deb_repo: Optional[str]
    deb_pkg: Optional[str]
    deb_deps: List[str]
    deb_pg: List[int]
    source: Optional[str]
    en_desc: str
    zh_desc: Optional[str]
    comment: Optional[str]


@dataclass
class PackageInfo:
    """Package availability information"""
    pg: int
    os: str
    name: str
    pkg: str
    ext: str
    state: str
    org: Optional[str]
    version: Optional[str]


@dataclass
class BinaryPackage:
    """Binary package details"""
    pg: int
    os: str
    name: str
    repo: str
    org: str
    type: str
    version: str
    file: str
    sha256: str
    href: str
    default_url: str
    mirror_url: Optional[str]
    size: int
    size_full: int


@dataclass
class PackageDetail:
    """Detailed package information with URL"""
    name: str
    version: str
    os: str
    org: str
    size: int
    file: str
    url: str
    pg: int


class ExtensionGenerator:
    """Generates Hugo content for PostgreSQL extensions"""

    def __init__(self, pgurl: str, output_dir: str):
        self.pgurl = pgurl
        self.output_dir = Path(output_dir)
        self.conn = None
        self.pg_versions = []
        self.os_systems = []

    def connect(self):
        """Establish database connection and load configuration"""
        try:
            self.conn = psycopg2.connect(self.pgurl)
            logger.info("Connected to database")
            self._load_config()
        except Exception as e:
            logger.error(f"Failed to connect to database: {e}")
            raise

    def _load_config(self):
        """Load active PG versions and OS systems from database"""
        with self.conn.cursor() as cur:
            # Load active PostgreSQL versions
            cur.execute("SELECT pg FROM pgext.active_pg ORDER BY pg DESC")
            self.pg_versions = [row[0] for row in cur.fetchall()]
            logger.info(f"Loaded {len(self.pg_versions)} active PG versions: {self.pg_versions}")

            # Load active OS systems
            cur.execute("""SELECT os, os_full
                           FROM pgext.os
                           WHERE active
                           ORDER BY os_major, os_arch DESC""")
            self.os_systems = cur.fetchall()
            logger.info(f"Loaded {len(self.os_systems)} active OS systems")

    def disconnect(self):
        """Close database connection"""
        if self.conn:
            self.conn.close()
            logger.info("Disconnected from database")

    def parse_array(self, value: Any) -> List[str]:
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

    def parse_pg_versions(self, values: List[str]) -> List[int]:
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
        # Remove duplicates while preserving order
        seen = set()
        unique_versions = []
        for version in pg_versions:
            if version not in seen:
                unique_versions.append(version)
                seen.add(version)
        return unique_versions

    def get_extension(self, ext_name: str) -> Optional[ExtensionData]:
        """Fetch extension data from database"""
        with self.conn.cursor() as cur:
            cur.execute("""
                        SELECT id,
                               name,
                               pkg,
                               lead_ext,
                               category,
                               state,
                               url,
                               license,
                               tags,
                               version,
                               repo,
                               lang,
                               contrib,
                               lead,
                               has_bin,
                               has_lib,
                               need_ddl,
                               need_load,
                               trusted,
                               relocatable,
                               schemas,
                               pg_ver,
                               requires,
                               require_by,
                               see_also,
                               rpm_ver,
                               rpm_repo,
                               rpm_pkg,
                               rpm_pg,
                               rpm_deps,
                               deb_ver,
                               deb_repo,
                               deb_pkg,
                               deb_deps,
                               deb_pg,
                               source,
                               en_desc,
                               zh_desc,
                               comment
                        FROM pgext.extension
                        WHERE name = %s
                        """, (ext_name,))
            row = cur.fetchone()

            if not row:
                return None

            (ext_id, name, pkg, lead_ext, category, state, url, license, tags_value, version,
             repo, lang, contrib, lead, has_bin, has_lib, need_ddl, need_load, trusted,
             relocatable, schemas_value, pg_ver_value, requires_value, require_by_value,
             see_also_value, rpm_ver, rpm_repo, rpm_pkg, rpm_pg_value, rpm_deps_value,
             deb_ver, deb_repo, deb_pkg, deb_deps_value, deb_pg_value, source, en_desc,
             zh_desc, comment) = row

            # Parse arrays
            tags = self.parse_array(tags_value)
            schemas = self.parse_array(schemas_value)
            pg_ver = self.parse_pg_versions(self.parse_array(pg_ver_value))
            requires = self.parse_array(requires_value)
            require_by = self.parse_array(require_by_value)
            see_also = self.parse_array(see_also_value)
            rpm_pg = self.parse_pg_versions(self.parse_array(rpm_pg_value))
            rpm_deps = self.parse_array(rpm_deps_value)
            deb_pg = self.parse_pg_versions(self.parse_array(deb_pg_value))
            deb_deps = self.parse_array(deb_deps_value)

            return ExtensionData(
                id=ext_id,
                name=name,
                pkg=pkg,
                lead_ext=lead_ext,
                category=category,
                state=state,
                url=url,
                license=license,
                tags=tags,
                version=version,
                repo=repo,
                lang=lang,
                contrib=contrib,
                lead=lead,
                has_bin=has_bin,
                has_lib=has_lib,
                need_ddl=need_ddl,
                need_load=need_load,
                trusted=trusted,
                relocatable=relocatable,
                schemas=schemas,
                pg_ver=pg_ver,
                requires=requires,
                require_by=require_by,
                see_also=see_also,
                rpm_ver=rpm_ver,
                rpm_repo=rpm_repo,
                rpm_pkg=rpm_pkg,
                rpm_pg=rpm_pg,
                rpm_deps=rpm_deps,
                deb_ver=deb_ver,
                deb_repo=deb_repo,
                deb_pkg=deb_pkg,
                deb_deps=deb_deps,
                deb_pg=deb_pg,
                source=source,
                en_desc=en_desc,
                zh_desc=zh_desc,
                comment=comment
            )

    def get_packages(self, pkg_name: str) -> List[PackageInfo]:
        """Get package availability matrix from pgext.pkg"""
        with self.conn.cursor() as cur:
            cur.execute("""
                        SELECT pg,
                               os,
                               name,
                               pkg,
                               ext,
                               state,
                               org,
                               version
                        FROM pgext.pkg
                        WHERE pkg = %s
                        ORDER BY pg DESC, os
                        """, (pkg_name,))

            return [PackageInfo(*row) for row in cur.fetchall()]

    def get_binaries(self, pkg_name: str) -> List[BinaryPackage]:
        """Get binary package details from pgext.bin + pgext.repository"""
        with self.conn.cursor() as cur:
            cur.execute("""
                        SELECT DISTINCT b.pg,
                                        b.os,
                                        b.name,
                                        b.repo,
                                        r.org,
                                        r.type,
                                        b.version,
                                        b.file,
                                        b.sha256,
                                        b.href,
                                        r.default_url,
                                        r.mirror_url,
                                        b.size,
                                        b.size_full
                        FROM pgext.bin b
                                 JOIN pgext.repository r ON b.repo = r.id
                                 JOIN pgext.pkg p ON p.pg = b.pg AND p.os = b.os AND p.name = b.name
                        WHERE p.pkg = %s
                        ORDER BY b.pg DESC, b.os, b.name
                        """, (pkg_name,))

            return [BinaryPackage(*row) for row in cur.fetchall()]

    def get_package_details(self, pkg_name: str) -> List[PackageDetail]:
        """Get detailed package information with URLs for all PG versions"""
        with self.conn.cursor() as cur:
            cur.execute("""
                        SELECT b.name,
                               b.version,
                               p.os,
                               r.org,
                               b.size,
                               b.file,
                               format('%%s/%%s', r.default_url, b.href) AS url,
                               b.pg
                        FROM pgext.pkg p
                                 JOIN pgext.bin b USING (pg, os, name)
                                 JOIN pgext.repository r ON b.repo = r.id
                                 JOIN pgext.os o ON p.os = o.os
                        WHERE p.pkg = %s
                           OR p.ext = %s
                        ORDER BY b.pg DESC, o.os_major, b.version DESC
                        """, (pkg_name, pkg_name))

            return [PackageDetail(
                name=row[0], version=row[1], os=row[2], org=row[3],
                size=row[4], file=row[5], url=row[6], pg=row[7]
            ) for row in cur.fetchall()]

    def get_all_extensions(self) -> List[str]:
        """Get all extension names"""
        with self.conn.cursor() as cur:
            cur.execute("SELECT name FROM pgext.extension ORDER BY name")
            return [row[0] for row in cur.fetchall()]

    def get_sibling_extensions(self, pkg_name: str, current: str) -> List[str]:
        """Fetch other extensions that share the same package"""
        with self.conn.cursor() as cur:
            cur.execute("""
                        SELECT name
                        FROM pgext.extension
                        WHERE pkg = %s
                        ORDER BY name
                        """, (pkg_name,))
            names = [row[0] for row in cur.fetchall()]
        return [name for name in names if name != current]

    # Markdown generation helpers
    def badge(self, content: str, color: str = "default", alt: str = "", link: str = "") -> str:
        """Generate badge shortcode"""
        parts = [f'{{{{< badge content="{content}"']
        if color != "default":
            parts.append(f'color="{color}"')
        if alt:
            parts.append(f'alt="{alt}"')
        if link:
            parts.append(f'link="{link}"')
        return ' '.join(parts) + ' >}}'

    def pkg_shortcode(self, name: str, version: str = "", org: str = "", url: str = "") -> str:
        """Generate package shortcode"""
        parts = [f'{{{{< pkg "{name}"']
        if version:
            parts.append(f'"{version}"')
        if org:
            parts.append(f'"{org}"')
        if url:
            parts.append(f'"{url}"')
        return ' '.join(parts) + ' >}}'

    def ext_link(self, ext_name: str) -> str:
        """Generate extension link shortcode"""
        return f'{{{{< ext "{ext_name}" >}}}}'

    def ext_shortcode(self, target: str, label: Optional[str] = None) -> str:
        """Render ext shortcode with optional display label"""
        if not target:
            return '{{< ext "" >}}'
        safe_target = target.replace('"', '\\"')
        if label:
            safe_label = label.replace('"', '\\"')
            return f'{{{{< ext "{safe_target}" "{safe_label}" >}}}}'
        return f'{{{{< ext "{safe_target}" >}}}}'

    def format_size(self, size_bytes: int) -> str:
        """Format byte size to human-readable format"""
        if size_bytes < 1024:
            return f"{size_bytes} B"
        elif size_bytes < 1024 * 1024:
            return f"{size_bytes / 1024:.1f} KiB"
        elif size_bytes < 1024 * 1024 * 1024:
            return f"{size_bytes / (1024 * 1024):.1f} MiB"
        else:
            return f"{size_bytes / (1024 * 1024 * 1024):.2f} GiB"

    def sanitize_table_text(self, value: Optional[str]) -> str:
        """Normalize text for safe markdown table output"""
        if not value:
            return ""
        sanitized = value.replace('|', '\\|').replace('\n', ' ')
        sanitized = ' '.join(sanitized.split())
        return sanitized.strip()

    # Content generation methods
    def generate_frontmatter(self, ext: ExtensionData) -> str:
        """Generate YAML frontmatter"""
        category_title = ext.category.title() if ext.category else 'Unknown'
        return f"""---
title: "{ext.name}"
linkTitle: "{ext.name}"
description: "{ext.en_desc}"
weight: {ext.id}
categories: ["{category_title}"]
width: full
---

{ext.en_desc}
"""

    def generate_overview(self, ext: ExtensionData) -> str:
        """Generate overview section"""
        category = f'{{{{< category "{ext.category.upper()}" >}}}}'
        license = f'{{{{< license "{ext.license}" >}}}}'
        lang = f'{{{{< language "{ext.lang}" >}}}}'
        ext_badge = self.badge(ext.name, link=ext.url or "")
        pkg_shortcode = self.ext_shortcode(ext.name, ext.pkg)

        return f"""
## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **{ext.id}** | {ext_badge} | {pkg_shortcode} | `{ext.version}` | {category} | {license} | {lang} |

"""

    def generate_attributes(self, ext: ExtensionData) -> str:
        """Generate attributes section"""
        # Build attribute string
        attrs = ""
        attrs += "c" if ext.contrib else "-"
        attrs += "l" if ext.lead else "-"
        attrs += "b" if ext.has_bin else "-"
        attrs += "s" if ext.has_lib else "-"
        attrs += "L" if ext.need_load else "-"
        attrs += "d" if ext.need_ddl else "-"
        attrs += "t" if ext.trusted else "-"
        attrs += "r" if ext.relocatable else "-"

        # Generate badges
        has_bin = self.badge("Yes" if ext.has_bin else "No", "green")
        has_lib = self.badge("Yes" if ext.has_lib else "No", "green")
        need_load = self.badge("Yes" if ext.need_load else "No", "red" if ext.need_load else "green")
        need_ddl = self.badge("Yes" if ext.need_ddl else "No", "green")
        relocatable = self.badge("yes" if ext.relocatable else "no", "green" if ext.relocatable else "red")
        trusted = self.badge("yes" if ext.trusted else "no", "green" if ext.trusted else "red")

        attr_badge = self.badge(f"--{attrs[2:]}", "blue")

        return f"""
|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {attr_badge} | {has_bin} | {has_lib} | {need_load} | {need_ddl} | {relocatable} | {trusted} |

"""

    def generate_relationships(self, ext: ExtensionData, siblings: Optional[List[str]] = None) -> str:
        """Generate relationships section"""
        siblings = siblings or []

        if not ext.requires and not ext.require_by and not ext.see_also and not siblings:
            return ""

        lines = ["\n| **Relationships** |   |",
                 "|:-----------------:|:----|"]

        if ext.requires:
            requires = ' '.join([self.ext_link(e) for e in ext.requires])
            lines.append(f"|   **Requires**    | {requires} |")

        if ext.require_by:
            require_by = ' '.join([self.ext_link(e) for e in ext.require_by])
            lines.append(f"|    **Need By**    | {require_by} |")

        if ext.see_also:
            see_also = ' '.join([self.ext_link(e) for e in ext.see_also])
            lines.append(f"|   **See Also**    | {see_also} |")

        if siblings:
            sibling_links = ' '.join([self.ext_link(e) for e in siblings])
            lines.append(f"|    **Siblings**   | {sibling_links} |")

        return '\n'.join(lines) + '\n\n'

    def generate_contrib_tip(self, ext: ExtensionData) -> str:
        """Generate contrib extension tip"""
        if not ext.contrib:
            return ""

        return """> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel

"""

    def generate_comment_note(self, ext: ExtensionData) -> str:
        """Generate comment note if exists"""
        if not ext.comment or not ext.comment.strip():
            return ""

        return f"""> [!Note] {ext.comment}

"""

    def generate_packages_table(self, ext: ExtensionData, packages: List[PackageInfo]) -> str:
        """Generate packages overview table"""
        # Special handling for contrib extensions
        if ext.contrib:
            return self.generate_contrib_packages_table(ext)

        rows = []

        rpm_row = self._build_package_summary_row(
            label="EL",
            repo=ext.rpm_repo,
            version=ext.rpm_ver,
            pattern=ext.rpm_pkg,
            supported_pg=ext.rpm_pg,
            deps=ext.rpm_deps,
            ext=ext
        )
        if rpm_row:
            rows.append(rpm_row)

        deb_row = self._build_package_summary_row(
            label="Debian",
            repo=ext.deb_repo,
            version=ext.deb_ver,
            pattern=ext.deb_pkg,
            supported_pg=ext.deb_pg,
            deps=ext.deb_deps,
            ext=ext
        )
        if deb_row:
            rows.append(deb_row)

        if not rows:
            return ""

        # Build table
        lines = ["\n## Packages\n",
                 "| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |",
                 "|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|"]

        lines.extend(rows)

        return '\n'.join(lines) + '\n\n'

    def _build_package_summary_row(
            self,
            label: str,
            repo: Optional[str],
            version: Optional[str],
            pattern: Optional[str],
            supported_pg: List[int],
            deps: List[str],
            ext: ExtensionData
    ) -> Optional[str]:
        """Build a single summary row for the package overview table"""
        if not any([repo, version, pattern, deps, supported_pg]):
            return None

        repo_label = repo or "N/A"
        repo_badge = self.badge(repo_label, link=f"/e/{ext.lead_ext or ext.name}")

        version_display = f"`{version}`" if version else "-"
        pattern_display = f"`{pattern}`" if pattern else "-"

        supported = set(supported_pg)
        pg_badges = []
        for pg in self.pg_versions:
            if pg in supported:
                pg_badges.append(self.badge(str(pg), "green"))
            else:
                alt = ""
                if pattern:
                    alt = pattern.replace("$v", str(pg))
                pg_badges.append(self.badge(str(pg), "red", alt))
        pg_cell = ' '.join(pg_badges)

        deps_str = ', '.join([f"`{dep}`" for dep in deps]) if deps else "-"

        return f"| **{label}** | {repo_badge} | {version_display} | {pg_cell} | {pattern_display} | {deps_str} |"

    def generate_contrib_packages_table(self, ext: ExtensionData) -> str:
        """Generate simple packages table for contrib extensions"""
        # Get available PG versions from pg_ver field
        available_pgs = set(ext.pg_ver)

        # Build table header with PG versions
        header_cells = ' | '.join([f"**PG{pg}**" for pg in self.pg_versions])
        header_line = f"| {header_cells} |"

        # Build separator line
        separator = "|" + ":--------:|" * len(self.pg_versions)

        # Build availability row with badges
        availability_cells = []
        for pg in self.pg_versions:
            if pg in available_pgs:
                # Available - green badge with version
                availability_cells.append(self.badge(ext.version, "green"))
            else:
                # Not available - red badge with N/A
                availability_cells.append(self.badge("N/A", "red"))

        availability_line = f"| {' | '.join(availability_cells)} |"

        # Combine all lines
        return f"""
## Packages

{header_line}
{separator}
{availability_line}

"""

    def generate_availability_matrix(self, packages: List[PackageInfo], binaries: List[BinaryPackage]) -> str:
        """Generate detailed availability matrix"""
        # Build binary lookup map
        binary_map = {}
        for binary in binaries:
            key = (binary.os, binary.pg)
            # Prefer PGDG over PIGSTY
            if key not in binary_map or binary.org.upper() == 'PGDG':
                binary_map[key] = binary

        # Build package lookup for expected names
        pkg_map = {}
        for pkg in packages:
            pkg_map[(pkg.os, pkg.pg)] = pkg

        # Generate table header
        lines = ["\n| **Linux** / **PG** |"]
        for pg in self.pg_versions:
            lines[0] += f"                  **PG{pg}**                   |"
        lines.append(
            "|:------------------:|" + ":-------------------------------------------:|" * len(self.pg_versions))

        # Generate rows for each OS
        for os_code, os_name in self.os_systems:
            row = [f"|    `{os_code}`    |"]

            for pg in self.pg_versions:
                key = (os_code, pg)

                if key in binary_map:
                    binary = binary_map[key]
                    # Construct full URL
                    url = binary.default_url.rstrip('/') + '/' + binary.href.lstrip('/')

                    if binary.org.upper() == 'HIDE':
                        row.append(f'  {self.pkg_shortcode(binary.name, binary.version, "HIDE")}   |')
                    elif url:
                        row.append(f' {self.pkg_shortcode(binary.name, binary.version, binary.org, url)} |')
                    else:
                        row.append(f'      {self.pkg_shortcode(binary.name, binary.version, binary.org)}      |')
                elif key in pkg_map:
                    # Package expected but not available
                    pkg = pkg_map[key]
                    row.append(f'    {self.pkg_shortcode(pkg.name)}     |')
                else:
                    # Not available at all
                    row.append(f'    {self.pkg_shortcode("N/A")}     |')

            lines.append(''.join(row))

        return '\n'.join(lines) + '\n\n'

    def generate_package_details_tabs(self, pkg_name: str) -> str:
        """Generate detailed package information organized by PG version tabs"""
        details = self.get_package_details(pkg_name)

        if not details:
            return ""

        # Group packages by PG version
        pg_groups: Dict[int, List[PackageDetail]] = {}
        for detail in details:
            if detail.pg not in pg_groups:
                pg_groups[detail.pg] = []
            pg_groups[detail.pg].append(detail)

        # Generate tab items string
        tab_items = ','.join([f"PG{pg}" for pg in sorted(pg_groups.keys(), reverse=True)])

        # Start tabs shortcode
        lines = [f"\n{{{{< tabs items=\"{tab_items}\" >}}}}\n"]

        # Generate content for each tab
        for pg in sorted(pg_groups.keys(), reverse=True):
            pkg_list = pg_groups[pg]

            lines.append("\n{{< tab >}}\n")

            if pkg_list:
                # Table header
                lines.append("| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |")
                lines.append("|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|")

                # Table rows
                for pkg in pkg_list:
                    size_str = self.format_size(pkg.size)
                    lines.append(
                        f"| `{pkg.name}` | {pkg.version} | `{pkg.os}` | {pkg.org} | {size_str} | "
                        f"[{pkg.file}]({pkg.url}) |"
                    )
            else:
                lines.append("*No packages available for this PostgreSQL version.*\n")

            lines.append("\n{{< /tab >}}")

        # Close tabs shortcode
        lines.append("\n{{< /tabs >}}\n")

        return '\n'.join(lines)

    def generate_source_section(self, ext: ExtensionData) -> str:
        """Generate source section"""
        # Skip for contrib extensions - they're built-in
        if ext.contrib:
            return ""

        cards = []

        if ext.url:
            if 'github.com' in ext.url:
                icon = "github"
                subtitle = ext.url.replace('https://', '')
            else:
                icon = "link"
                subtitle = ext.url.replace('https://', '').replace('http://', '')
            cards.append(f'{{{{< card link="{ext.url}" title="Repository" icon="{icon}" subtitle="{subtitle}" >}}}}')

        if ext.source:
            cards.append(
                f'{{{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="{ext.source}" >}}}}')

        if not cards:
            return ""

        section = "\n## Source\n\n{{< cards cols=3 >}}\n"
        section += '\n'.join(cards) + '\n'
        section += "{{< /cards >}}\n\n"

        if ext.source:
            section += f"""
```bash
pig build get {ext.name}; # get {ext.name} source code
pig build dep {ext.name}; # install build dependencies
pig build pkg {ext.name}; # build extension rpm or deb
pig build ext {ext.name}; # build extension rpms
```

"""
        return section

    def generate_install_section(self, ext: ExtensionData) -> str:
        """Generate install section"""
        # Special handling for contrib extensions
        if ext.contrib:
            return self.generate_contrib_install_section(ext)

        # Generate install commands for supported PG versions
        install_cmds = []
        for pg in sorted(ext.pg_ver, reverse=True):
            install_cmds.append(f"pig ext install {ext.name} -v {pg};   # install for PG {pg}")

        # Determine CASCADE SCHEMA clause
        cascade_schema = ""
        if ext.schemas and not ext.relocatable:
            cascade_schema = f" CASCADE SCHEMA {ext.schemas[0]}"

        return f"""
## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install {ext.name}; # install by extension name, for the current active PG version
pig ext install {ext.pkg}; # install via package alias, for the active PG version
{chr(10).join(install_cmds)}

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION {ext.name}{cascade_schema};
```

"""

    def generate_contrib_install_section(self, ext: ExtensionData) -> str:
        """Generate install section for contrib extensions"""
        sections = []
        sections.append("\n## Install\n")

        # Add shared_preload_libraries section if needed
        if ext.need_load:
            sections.append(f"""
Add this extension to [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = '{ext.name}';  -- comma-separated list
```

""")

        # Determine CASCADE SCHEMA clause
        cascade_schema = ""
        if ext.schemas and not ext.relocatable:
            cascade_schema = f" CASCADE SCHEMA {ext.schemas[0]}"

        # Add CREATE section
        sections.append(f"""
[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION {ext.name}{cascade_schema};
```
""")

        return ''.join(sections)


    def generate_content(self, ext_name: str) -> Optional[str]:
        """Generate complete Markdown content for an extension"""
        # Get extension data
        ext = self.get_extension(ext_name)
        if not ext:
            logger.warning(f"Extension {ext_name} not found")
            return None

        # Get package and binary data
        packages = self.get_packages(ext.pkg)
        binaries = self.get_binaries(ext.pkg)
        siblings = self.get_sibling_extensions(ext.pkg, ext.name)

        # Generate all sections
        content = self.generate_frontmatter(ext)
        content += self.generate_overview(ext)
        content += self.generate_attributes(ext)
        content += self.generate_relationships(ext, siblings)

        # Add contrib tip and comment note after relationships
        content += self.generate_comment_note(ext)

        # Continue with package information
        content += self.generate_packages_table(ext, packages)

        # For non-contrib extensions, add detailed availability info
        if not ext.contrib:
            content += self.generate_availability_matrix(packages, binaries)
            content += self.generate_package_details_tabs(ext.pkg)
        else:
            content += self.generate_contrib_tip(ext)

        content += self.generate_source_section(ext)
        content += self.generate_install_section(ext)

        # Append stub content if exists
        stub_path = Path(__file__).parent.parent / 'stub' / f'{ext_name}.md'
        if stub_path.exists():
            try:
                with open(stub_path, 'r', encoding='utf-8') as f:
                    stub_content = f.read()
                if stub_content.strip():
                    content += '\n' + stub_content
                    logger.debug(f"Appended stub content for {ext_name}")
            except Exception as e:
                logger.warning(f"Failed to read stub file for {ext_name}: {e}")

        return content

    def generate_all(self):
        """Generate content for all extensions"""
        self.output_dir.mkdir(parents=True, exist_ok=True)

        extensions = self.get_all_extensions()
        logger.info(f"Found {len(extensions)} extensions to generate")

        success_count = 0
        for ext_name in extensions:
            try:
                content = self.generate_content(ext_name)
                if content:
                    output_file = self.output_dir / f"{ext_name}.md"
                    with open(output_file, 'w', encoding='utf-8') as f:
                        f.write(content)
                    success_count += 1
                    logger.debug(f"Generated: {output_file}")
            except Exception as e:
                logger.error(f"Error generating {ext_name}: {e}", exc_info=True)

        logger.info(f"Successfully generated {success_count}/{len(extensions)} extension pages")
        logger.info(f"Note: Run bin/gen-ext-index.py to generate index pages")


def main():
    """Main entry point"""
    parser = argparse.ArgumentParser(description='Generate Hugo content for PostgreSQL extensions')
    parser.add_argument('pgurl', nargs='?', default='postgres:///',
                        help='PostgreSQL connection URL (default: postgres:///)')
    parser.add_argument('output_dir', nargs='?', default=None,
                        help='Output directory for generated files')
    parser.add_argument('--verbose', '-v', action='store_true',
                        help='Enable verbose logging')

    args = parser.parse_args()

    # Set output directory
    if args.output_dir is None:
        script_dir = Path(__file__).parent
        output_dir = script_dir.parent / 'content' / 'e'
    else:
        output_dir = Path(args.output_dir)

    # Configure logging
    if args.verbose:
        logging.getLogger().setLevel(logging.DEBUG)

    # Create generator and run
    generator = ExtensionGenerator(args.pgurl, output_dir)

    try:
        generator.connect()
        generator.generate_all()
    finally:
        generator.disconnect()


if __name__ == "__main__":
    main()
