#!/usr/bin/env python3

import os
import psycopg2
import re
from typing import Dict, List, Optional, Any, Tuple
from collections import defaultdict, Counter
from dataclasses import dataclass
from pathlib import Path


# =============================================================================
# CONFIGURATION
# =============================================================================

@dataclass
class Config:
    """Configuration for the list generator."""
    
    # Database connection
    DB_CONNECTION: str = 'postgres:///vonng'
    
    # Supported operating systems
    OS_VERSIONS: List[str] = None
    
    # Supported PostgreSQL major versions
    PG_VERSIONS: List[int] = None
    
    # Output directory
    OUTPUT_DIR: str = None
    
    def __post_init__(self):
        if self.OS_VERSIONS is None:
            self.OS_VERSIONS = [
                'el8.x86_64', 'el8.aarch64', 
                'el9.x86_64', 'el9.aarch64',
                'd12.x86_64', 'd12.aarch64', 
                'u22.x86_64', 'u22.aarch64', 
                'u24.x86_64', 'u24.aarch64'
            ]
        
        if self.PG_VERSIONS is None:
            self.PG_VERSIONS = [18, 17, 16, 15, 14, 13]
        
        if self.OUTPUT_DIR is None:
            script_dir = os.path.dirname(os.path.abspath(__file__))
            self.OUTPUT_DIR = os.path.abspath(os.path.join(script_dir, '..', 'content', 'list'))


# OS and Language mappings
OS_DESCRIPTIONS = {
    'el8': 'Enterprise Linux 8 (RHEL 8, CentOS 8, Rocky 8, Alma 8)',
    'el9': 'Enterprise Linux 9 (RHEL 9, CentOS 9, Rocky 9, Alma 9)',
    'd12': 'Debian 12 (Bookworm)',
    'u22': 'Ubuntu 22.04 LTS (Jammy)',
    'u24': 'Ubuntu 24.04 LTS (Noble)'
}

LANGUAGE_DESCRIPTIONS = {
    'C': 'The traditional PostgreSQL extension language',
    'C++': 'Extensions leveraging C++ features and libraries',
    'Rust': 'Extensions written in Rust with the pgrx framework',
    'Python': 'Extensions written in Python',
    'SQL': 'Pure SQL extensions and functions',
    'Java': 'Extensions running on JVM',
    'Data': 'Data-only extensions',
}

# Category metadata - loaded from category.csv
CATEGORY_META = {}

LICENSE_INFO = {
    'PostgreSQL': {'url': 'https://opensource.org/licenses/postgresql', 'description': 'Very liberal license based on the BSD license, allowing almost unlimited freedom.', 'anchor': 'postgresql'},
    'MIT': {'url': 'https://opensource.org/licenses/MIT', 'description': 'A permissive license that allows commercial use, modification, and private use.', 'anchor': 'mit'},
    'ISC': {'url': 'https://opensource.org/licenses/ISC', 'description': 'A permissive license similar to MIT, allowing commercial use and modification.', 'anchor': 'isc'},
    'BSD 0-Clause': {'url': 'https://opensource.org/license/0bsd', 'description': 'Public domain equivalent license with no restrictions on use.', 'anchor': 'bsd-0-clause'},
    'BSD 2-Clause': {'url': 'https://opensource.org/license/bsd-2-clause', 'description': 'Permissive license requiring attribution but allowing commercial use.', 'anchor': 'bsd-2-clause'},
    'BSD 3-Clause': {'url': 'https://opensource.org/license/bsd-3-clause', 'description': 'Permissive license with attribution and endorsement restriction clauses.', 'anchor': 'bsd-3-clause'},
    'Artistic': {'url': 'https://opensource.org/license/artistic-2-0', 'description': 'Copyleft license allowing modification with certain distribution requirements.', 'anchor': 'artistic'},
    'Apache-2.0': {'url': 'https://opensource.org/licenses/Apache-2.0', 'description': 'Permissive license with patent protection and attribution requirements.', 'anchor': 'apache-20'},
    'MPL-2.0': {'url': 'https://opensource.org/licenses/MPL-2.0', 'description': 'Weak copyleft license allowing proprietary combinations with file-level copyleft.', 'anchor': 'mpl-20'},
    'GPL-2.0': {'url': 'https://opensource.org/licenses/GPL-2.0', 'description': 'Strong copyleft license requiring derivative works to be open source.', 'anchor': 'gpl-20'},
    'GPL-3.0': {'url': 'https://opensource.org/licenses/GPL-3.0', 'description': 'Strong copyleft license with additional patent and hardware restrictions.', 'anchor': 'gpl-30'},
    'LGPL-2.1': {'url': 'https://opensource.org/licenses/LGPL-2.1', 'description': 'Weak copyleft license allowing proprietary applications to link dynamically.', 'anchor': 'lgpl-21'},
    'LGPL-3.0': {'url': 'https://opensource.org/licenses/LGPL-3.0', 'description': 'Weak copyleft license with additional patent and hardware provisions.', 'anchor': 'lgpl-30'},
    'AGPL-3.0': {'url': 'https://opensource.org/licenses/AGPL-3.0', 'description': 'Network copyleft license extending GPL to cover network-distributed software.', 'anchor': 'agpl-30'},
    'Timescale': {'url': 'https://www.timescale.com/legal/licenses', 'description': 'Proprietary license with restrictions on commercial use and distribution.', 'anchor': 'timescale'}
}

LICENSE_NORMALIZATION = {
    'BSD-0': 'BSD 0-Clause', 'BSD-2': 'BSD 2-Clause', 'BSD-3': 'BSD 3-Clause',
    'GPLv2': 'GPL-2.0', 'GPLv3': 'GPL-3.0', 'LGPLv2': 'LGPL-2.1', 
    'LGPLv3': 'LGPL-3.0', 'AGPLv3': 'AGPL-3.0', 'MPLv2': 'MPL-2.0'
}


# =============================================================================
# HUGO SHORTCODE HELPERS
# =============================================================================

class HugoShortcode:
    """Helper class for generating Hugo shortcodes."""
    
    @staticmethod
    def badge(content: str, color: str = "default", alt: str = "", link: str = "") -> str:
        """Generate a badge shortcode."""
        parts = ['{{< badge']
        parts.append(f'content="{content}"')
        if color != "default":
            parts.append(f'color="{color}"')
        if alt:
            parts.append(f'alt="{alt}"')
        if link:
            parts.append(f'link="{link}"')
        return ' '.join(parts) + ' >}}'
    
    @staticmethod
    def category(name: str) -> str:
        """Generate a category shortcode."""
        return '{{< category "' + name + '" >}}'
    
    @staticmethod
    def license(name: str) -> str:
        """Generate a license shortcode."""
        return '{{< license "' + name + '" >}}'
    
    @staticmethod
    def language(name: str) -> str:
        """Generate a language shortcode."""
        return '{{< language "' + name + '" >}}'
    
    @staticmethod
    def ext(name: str) -> str:
        """Generate an extension shortcode."""
        return '{{< ext "' + name + '" >}}'
    
    @staticmethod
    def pkg(name: str, version: str = "", repo: str = "", url: str = "") -> str:
        """Generate a package shortcode."""
        parts = ['{{< pkg', f'"{name}"']
        if version:
            parts.append(f'"{version}"')
        if repo:
            parts.append(f'"{repo}"')
        if url:
            parts.append(f'"{url}"')
        return ' '.join(parts) + ' >}}'


# =============================================================================
# UTILITY FUNCTIONS
# =============================================================================

def parse_array(value: str) -> List[str]:
    """Parse PostgreSQL array string to Python list."""
    if isinstance(value, list):
        return value
    if not value or not value.startswith('{') or not value.endswith('}'):
        return []
    return [item.strip() for item in value[1:-1].split(',') if item.strip()]


def normalize_os_arch(arch_in: str) -> str:
    """Normalize architecture names."""
    arch_map = {
        'amd': 'x86_64', 'x86_64': 'x86_64', 'amd64': 'x86_64',
        'arm': 'aarch64', 'arm64': 'aarch64', 'aarch64': 'aarch64', 'armv8': 'aarch64'
    }
    return arch_map.get(arch_in.lower(), arch_in.lower())


def normalize_license_name(license_name: str) -> str:
    """Normalize license names to standard format."""
    return LICENSE_NORMALIZATION.get(license_name, license_name)


def extract_semantic_version(version: str) -> str:
    """Extract semantic version from package version string."""
    semantic_version = re.split(r'[-_]\d*[A-Z]', version)[0]
    semantic_version = re.split(r'[-_]\d*\w*\.', semantic_version)[0]
    semantic_version = re.split(r'[-+]', semantic_version)[0]
    return semantic_version


# =============================================================================
# DATA CLASSES
# =============================================================================

@dataclass
class Package:
    """Represents a PostgreSQL extension package."""
    pg: int
    os: str
    pname: str
    org: str
    type: str
    os_code: str
    os_arch: str
    repo: str
    name: str
    ver: str
    version: str
    release: str
    file: str
    sha256: str
    url: str
    mirror_url: str
    size: int
    size_full: int
    
    @classmethod
    def from_row(cls, row: Tuple) -> 'Package':
        """Create Package from database row."""
        return cls(*row)


@dataclass
class Extension:
    """Represents a PostgreSQL extension with all its metadata."""
    id: int
    name: str
    pkg: str
    lead_ext: Optional[str]
    category: str
    state: str
    url: Optional[str]
    license: Optional[str]
    tags: List[str]
    version: Optional[str]
    repo: str
    lang: Optional[str]
    contrib: bool
    lead: bool
    has_bin: bool
    has_lib: bool
    need_ddl: bool
    need_load: bool
    trusted: bool
    relocatable: bool
    schemas: List[str]
    pg_ver: List[str]
    requires: List[str]
    require_by: List[str]
    see_also: List[str]
    rpm_ver: Optional[str]
    rpm_repo: Optional[str]
    rpm_pkg: Optional[str]
    rpm_pg: List[str]
    rpm_deps: List[str]
    deb_ver: Optional[str]
    deb_repo: Optional[str]
    deb_pkg: Optional[str]
    deb_deps: List[str]
    deb_pg: List[str]
    source: Optional[str]
    extra: Optional[str]
    en_desc: Optional[str]
    zh_desc: Optional[str]
    comment: Optional[str]
    mtime: str
    packages: List[Package] = None
    
    @classmethod
    def from_row(cls, row: Tuple) -> 'Extension':
        """Create Extension from database row."""
        (id, name, pkg, lead_ext, category, state, url, license, tags, version, repo, lang, 
         contrib, lead, has_bin, has_lib, need_ddl, need_load, trusted, relocatable, schemas, 
         pg_ver, requires, require_by, see_also, rpm_ver, rpm_repo, rpm_pkg, rpm_pg, rpm_deps, 
         deb_ver, deb_repo, deb_pkg, deb_deps, deb_pg, source, extra, en_desc, zh_desc, comment, mtime) = row
        
        # Parse array fields
        tags = parse_array(tags) if tags else []
        pg_ver = parse_array(pg_ver) if pg_ver else []
        requires = parse_array(requires) if requires else []
        require_by = parse_array(require_by) if require_by else []
        see_also = parse_array(see_also) if see_also else []
        schemas = parse_array(schemas) if schemas else []
        rpm_deps = parse_array(rpm_deps) if rpm_deps else []
        deb_deps = parse_array(deb_deps) if deb_deps else []
        rpm_pg = parse_array(rpm_pg) if rpm_pg else []
        deb_pg = parse_array(deb_pg) if deb_pg else []
        
        return cls(
            id=id, name=name, pkg=pkg, lead_ext=lead_ext, category=category, state=state, url=url,
            license=license, tags=tags, version=version, repo=repo, lang=lang, contrib=contrib,
            lead=lead, has_bin=has_bin, has_lib=has_lib, need_ddl=need_ddl, need_load=need_load,
            trusted=trusted, relocatable=relocatable, schemas=schemas, pg_ver=pg_ver,
            requires=requires, require_by=require_by, see_also=see_also, rpm_ver=rpm_ver, 
            rpm_repo=rpm_repo, rpm_pkg=rpm_pkg, rpm_pg=rpm_pg, rpm_deps=rpm_deps, deb_ver=deb_ver, 
            deb_repo=deb_repo, deb_pkg=deb_pkg, deb_deps=deb_deps, deb_pg=deb_pg, source=source,
            extra=extra, en_desc=en_desc, zh_desc=zh_desc, comment=comment, mtime=mtime, packages=[]
        )
    
    @property
    def has_rpm(self) -> bool:
        return bool(self.rpm_repo)
    
    @property
    def has_deb(self) -> bool:
        return bool(self.deb_repo)
    
    def load_packages(self, all_packages: List[Package]):
        """Load package data for this extension from the provided list."""
        if self.packages is None:
            self.packages = []
        
        # Find packages that match this extension's package name
        for pkg in all_packages:
            # Handle different package naming patterns:
            # RPM: pkg_name, pkg_name_XX (where XX is PG version)
            # DEB: postgresql-XX-pkg-name (where XX is PG version)
            if (pkg.pname == self.pkg or 
                pkg.pname.startswith(f'{self.pkg}_') or
                f'-{self.pkg.replace("_", "-")}' in pkg.pname):
                self.packages.append(pkg)


# =============================================================================
# TABLE GENERATOR CLASSES
# =============================================================================

class TableGenerator:
    """Centralized table generation functionality."""
    
    def __init__(self, leading_map: Dict[str, str]):
        self.leading_map = leading_map
    
    def generate_simple_table(self, extensions: List[Extension]) -> str:
        """Generate simple extension table (ID, Name, Package, Description)."""
        if not extensions:
            return "No extensions found."
        
        headers = ['ID', 'Extension', 'Package', 'Description']
        rows = [self._format_table_header(headers, [':---:',':---',':---',':---'])]
        
        for ext in extensions:
            package_cell = f'[`{ext.pkg}`](/e/{self.leading_map.get(ext.pkg, ext.name)})'
            row_data = [
                str(ext.id),
                f'[`{ext.name}`](/e/{ext.name})',
                package_cell,
                ext.en_desc or 'No description'
            ]
            rows.append('| ' + ' | '.join(row_data) + ' |')
        
        return '\n'.join(rows)
    
    def generate_category_table(self, extensions: List[Extension]) -> str:
        """Generate extension table for category lists (ID, Name, Package, Version, Description)."""
        if not extensions:
            return "No extensions found."
        
        headers = ['ID', 'Extension', 'Package', 'Version', 'Description']
        rows = [self._format_table_header(headers, [':---:',':---',':---',':---',':---'])]
        
        for ext in extensions:
            package_cell = f'[`{ext.pkg}`](/e/{self.leading_map.get(ext.pkg, ext.name)})'
            row_data = [
                str(ext.id),
                f'[`{ext.name}`](/e/{ext.name})',
                package_cell,
                ext.version or 'N/A',
                ext.en_desc or 'No description'
            ]
            rows.append('| ' + ' | '.join(row_data) + ' |')
        
        return '\n'.join(rows)
    
    def generate_category_table_zh(self, extensions: List[Extension]) -> str:
        """Generate Chinese extension table for category lists."""
        if not extensions:
            return "未找到扩展。"
        
        headers = ['ID', '扩展', '包', '版本', '描述']
        rows = [self._format_table_header(headers, [':---:',':---',':---',':---',':---'])]
        
        for ext in extensions:
            package_cell = f'[`{ext.pkg}`](/zh/e/{self.leading_map.get(ext.pkg, ext.name)})'
            row_data = [
                str(ext.id),
                f'[`{ext.name}`](/zh/e/{ext.name})',
                package_cell,
                ext.version or 'N/A',
                ext.zh_desc or ext.en_desc or '暂无描述'
            ]
            rows.append('| ' + ' | '.join(row_data) + ' |')
        
        return '\n'.join(rows)
    
    def generate_repo_table(self, extensions: List[Extension]) -> str:
        """Generate extension table for repo lists (ID, Name, Category, RPM, DEB, Description)."""
        if not extensions:
            return "No extensions found."
        
        headers = ['ID', 'Name', 'Category', 'RPM', 'DEB', 'Description']
        rows = [self._format_table_header(headers, [':---:',':---',':---',':---:',':---:',':---'])]
        
        for ext in extensions:
            # Use Hugo shortcodes for badges
            rpm_badge = HugoShortcode.badge(ext.rpm_repo, self._get_repo_color(ext.rpm_repo)) if ext.rpm_repo else '-'
            deb_badge = HugoShortcode.badge(ext.deb_repo, self._get_repo_color(ext.deb_repo)) if ext.deb_repo else '-'
            category_badge = HugoShortcode.category(ext.category) if ext.category else '-'
            
            row_data = [
                str(ext.id),
                f'[`{ext.name}`](/e/{ext.name})',
                category_badge,
                rpm_badge,
                deb_badge,
                ext.en_desc or 'No description'
            ]
            rows.append('| ' + ' | '.join(row_data) + ' |')
        
        return '\n'.join(rows)
    
    def generate_linux_table(self, extensions: List[Extension], platform: str) -> str:
        """Generate extension table showing availability for a specific Linux platform."""
        if not extensions:
            return "No extensions found."
        
        headers = ['ID', 'Name', 'Category', 'PG Versions', 'Description']
        rows = [self._format_table_header(headers, [':---:',':---',':---',':---:',':---'])]
        
        for ext in extensions:
            # Determine available PG versions for this platform
            available_pg = []
            if platform.startswith('el'):
                available_pg = ext.rpm_pg
            else:
                available_pg = ext.deb_pg
            
            pg_badges = ' '.join([HugoShortcode.badge(pg, 'green') for pg in sorted(available_pg)])
            category_badge = HugoShortcode.category(ext.category) if ext.category else '-'
            
            row_data = [
                str(ext.id),
                f'[`{ext.name}`](/e/{ext.name})',
                category_badge,
                pg_badges or '-',
                ext.en_desc or 'No description'
            ]
            rows.append('| ' + ' | '.join(row_data) + ' |')
        
        return '\n'.join(rows)
    
    def _get_repo_color(self, repo: str) -> str:
        """Get color for repository badge."""
        repo_colors = {
            'PIGSTY': 'amber',
            'PGDG': 'blue',
            'CONTRIB': 'purple',
            'OTHER': 'gray'
        }
        return repo_colors.get(repo, 'gray')
    
    def _format_table_header(self, headers: List[str], alignments: List[str]) -> str:
        """Format table header with alignment."""
        header_row = '| ' + ' | '.join(headers) + ' |'
        separator_row = '|' + '|'.join(alignments) + '|'
        return header_row + '\n' + separator_row


# =============================================================================
# DATABASE OPERATIONS
# =============================================================================

class DatabaseManager:
    """Handles all database operations."""
    
    def __init__(self, connection_string: str):
        self.connection_string = connection_string
        self._conn = None
    
    def get_connection(self):
        if self._conn is None:
            self._conn = psycopg2.connect(self.connection_string)
        return self._conn
    
    def load_category_metadata(self) -> Dict[str, Dict[str, str]]:
        """Load category metadata from CSV file."""
        print("Loading category metadata from CSV...")
        
        import csv
        category_data = {}
        csv_path = os.path.join(os.path.dirname(__file__), '..', 'src', 'category.csv')
        
        if not os.path.exists(csv_path):
            # Try alternative path
            csv_path = os.path.join(os.path.dirname(__file__), '..', 'data', 'category.csv')
        
        with open(csv_path, 'r', encoding='utf-8') as f:
            reader = csv.DictReader(f)
            for row in reader:
                category_data[row['name']] = {
                    'zh_desc': row['zh_desc'],
                    'en_desc': row['en_desc'],
                    'icon': row.get('icon', 'Blocks'),
                    'color': 'blue'  # Default color, can be customized
                }
        
        print(f"Loaded {len(category_data)} categories.")
        return category_data
    
    def load_extensions(self) -> List[Extension]:
        """Load all extensions from database."""
        print("Loading extensions from database...")
        
        with self.get_connection().cursor() as cur:
            cur.execute('SELECT * FROM pgext.extension ORDER BY id')
            rows = cur.fetchall()
        
        extensions = [Extension.from_row(row) for row in rows]
        print(f"Found {len(extensions)} extensions.")
        return extensions
    
    def load_packages(self) -> List[Package]:
        """Load all packages from database."""
        print("Loading packages from database...")
        
        with self.get_connection().cursor() as cur:
            cur.execute('SELECT pg, os, pname, org, type, os_code, os_arch, repo, name, ver, version, release, file, sha256, url, mirror_url, size, size_full FROM pgext.package ORDER BY os, pg DESC')
            rows = cur.fetchall()
        
        packages = [Package.from_row(row) for row in rows]
        print(f"Found {len(packages)} packages.")
        return packages


# =============================================================================
# CONTENT GENERATORS
# =============================================================================

class ContentGenerator:
    """Base class for content generators."""
    
    def __init__(self, config: Config, extensions: List[Extension], table_gen: TableGenerator):
        self.config = config
        self.extensions = extensions
        self.table_gen = table_gen
    
    def write_content(self, filename: str, content: str):
        """Write content to file."""
        output_path = os.path.join(self.config.OUTPUT_DIR, filename)
        with open(output_path, 'w', encoding='utf-8') as f:
            f.write(content)
        print(f"Generated: {output_path}")


class MainIndexGenerator(ContentGenerator):
    """Generate main extension index."""
    
    def generate(self):
        print("Generating main extension index...")
        
        # Generate statistics
        total = len(self.extensions)
        contrib_count = sum(1 for ext in self.extensions if ext.contrib)
        pigsty_count = sum(1 for ext in self.extensions if ext.rpm_repo == 'PIGSTY' or ext.deb_repo == 'PIGSTY')
        pgdg_count = sum(1 for ext in self.extensions if ext.rpm_repo == 'PGDG' or ext.deb_repo == 'PGDG')
        
        # Group by category
        category_groups = defaultdict(list)
        for ext in self.extensions:
            category_groups[ext.category].append(ext)
        
        # Generate category sections
        category_sections = []
        for category in sorted(CATEGORY_META.keys()):
            if category not in category_groups:
                continue
            
            cat_extensions = sorted(category_groups[category], key=lambda e: e.name)
            count = len(cat_extensions)
            meta = CATEGORY_META[category]
            
            # Generate extension list
            ext_links = []
            for ext in cat_extensions[:20]:  # Show first 20
                ext_links.append(HugoShortcode.ext(ext.name))
            
            if count > 20:
                ext_links.append(f'... and {count - 20} more')
            
            section = f'''
### {HugoShortcode.category(category)} {meta["en_desc"]}

{HugoShortcode.badge(f"{count} Extensions", "gray")}

{' '.join(ext_links)}
'''
            category_sections.append(section)
        
        content = f'''---
title: Extension Index
description: Complete index of all PostgreSQL extensions
---

## Overview

There are **{total}** PostgreSQL extensions available in the catalog:

- **{contrib_count}** contrib extensions (included in PostgreSQL)
- **{pigsty_count}** extensions in PIGSTY repository
- **{pgdg_count}** extensions in PGDG repository

## Extensions by Category

{''.join(category_sections)}

## Quick Links

- [By Category](/list/cate) - Browse extensions organized by functionality
- [By Language](/list/lang) - Find extensions by implementation language
- [By License](/list/license) - Filter extensions by open source license
- [By Repository](/list/repo) - View extensions by distribution repository
- [By Attribute](/list/attr) - Search by extension attributes
- [By Linux](/list/linux) - Check availability on Linux distributions
- [By PostgreSQL](/list/pgsql) - Check compatibility with PostgreSQL versions
'''
        
        self.write_content('index.md', content)
        self._generate_chinese_version(total, contrib_count, pigsty_count, pgdg_count, category_groups)
    
    def _generate_chinese_version(self, total, contrib_count, pigsty_count, pgdg_count, category_groups):
        """Generate Chinese version of main index."""
        print("Generating Chinese main index...")
        
        # Generate category sections
        category_sections = []
        for category in sorted(CATEGORY_META.keys()):
            if category not in category_groups:
                continue
            
            cat_extensions = sorted(category_groups[category], key=lambda e: e.name)
            count = len(cat_extensions)
            meta = CATEGORY_META[category]
            
            # Generate extension list
            ext_links = []
            for ext in cat_extensions[:20]:  # Show first 20
                ext_links.append(HugoShortcode.ext(ext.name))
            
            if count > 20:
                ext_links.append(f'... 以及另外 {count - 20} 个')
            
            section = f'''
### {HugoShortcode.category(category)} {meta["zh_desc"]}

{HugoShortcode.badge(f"{count} 个扩展", "gray")}

{' '.join(ext_links)}
'''
            category_sections.append(section)
        
        content = f'''---
title: 扩展索引
description: PostgreSQL 扩展完整索引
---

## 概览

目录中共有 **{total}** 个 PostgreSQL 扩展：

- **{contrib_count}** 个 contrib 扩展（PostgreSQL 自带）
- **{pigsty_count}** 个扩展在 PIGSTY 仓库中
- **{pgdg_count}** 个扩展在 PGDG 仓库中

## 按分类浏览扩展

{''.join(category_sections)}

## 快速链接

- [按分类](/zh/list/cate) - 按功能分类浏览扩展
- [按语言](/zh/list/lang) - 按实现语言查找扩展
- [按许可证](/zh/list/license) - 按开源许可证过滤扩展
- [按仓库](/zh/list/repo) - 按发行仓库查看扩展
- [按属性](/zh/list/attr) - 按扩展属性搜索
- [按 Linux](/zh/list/linux) - 检查 Linux 发行版上的可用性
- [按 PostgreSQL](/zh/list/pgsql) - 检查与 PostgreSQL 版本的兼容性
'''
        
        self.write_content('index.zh.md', content)


class CategoryListGenerator(ContentGenerator):
    """Generate category-based extension list."""
    
    def generate(self):
        print("Generating category list...")
        
        # Group extensions by category
        category_groups = defaultdict(list)
        for ext in self.extensions:
            category_groups[ext.category].append(ext)
        
        # Generate category overview table with Hugo shortcodes
        category_table_rows = []
        for category in CATEGORY_META.keys():
            count = len(category_groups.get(category, []))
            meta = CATEGORY_META[category]
            category_badge = HugoShortcode.category(category)
            category_table_rows.append(f'| {category_badge} | {count} | {meta["en_desc"]} |')
        
        category_overview_table = f'''| Category | Count | Description |
|:---------|:-----:|:------------|
{chr(10).join(category_table_rows)}'''
        
        # Generate category sections
        category_sections = []
        for category in CATEGORY_META.keys():
            if category not in category_groups:
                continue
            
            cat_extensions = sorted(category_groups[category], key=lambda e: e.name)
            count = len(cat_extensions)
            meta = CATEGORY_META[category]
            
            section = f'''
## {category}

{meta["en_desc"]}

{HugoShortcode.category(category)} {HugoShortcode.badge(f"{count} Extensions", "gray")}

{self.table_gen.generate_category_table(cat_extensions)}
'''
            category_sections.append(section)
        
        # Generate the shortcode badges for navigation
        nav_badges = []
        for category in ['TIME', 'GIS', 'RAG', 'FTS', 'OLAP', 'FEAT', 'LANG', 'TYPE', 
                        'UTIL', 'FUNC', 'ADMIN', 'STAT', 'SEC', 'FDW', 'SIM', 'ETL']:
            if category in CATEGORY_META:
                nav_badges.append(HugoShortcode.category(category))
        
        content = f'''---
title: By Category
description: PostgreSQL extensions organized by functionality categories
---

{' '.join(nav_badges[:8])}

{' '.join(nav_badges[8:])}

## Summary

{category_overview_table}

{''.join(category_sections)}
'''
        
        self.write_content('cate.md', content)
        self._generate_chinese_version(category_groups, category_overview_table, category_sections)
    
    def _generate_chinese_version(self, category_groups, category_overview_table, category_sections):
        """Generate Chinese version of category list."""
        print("Generating Chinese category list...")
        
        # Generate Chinese category overview table
        zh_category_table_rows = []
        for category in CATEGORY_META.keys():
            count = len(category_groups.get(category, []))
            meta = CATEGORY_META[category]
            category_badge = HugoShortcode.category(category)
            zh_category_table_rows.append(f'| {category_badge} | {count} | {meta["zh_desc"]} |')
        
        zh_category_overview_table = f'''| 分类 | 数量 | 描述 |
|:---------|:-----:|:------------|
{chr(10).join(zh_category_table_rows)}'''
        
        # Generate Chinese category sections
        zh_category_sections = []
        for category in CATEGORY_META.keys():
            if category not in category_groups:
                continue
            
            cat_extensions = sorted(category_groups[category], key=lambda e: e.name)
            count = len(cat_extensions)
            meta = CATEGORY_META[category]
            
            section = f'''
## {category}

{meta["zh_desc"]}

{HugoShortcode.category(category)} {HugoShortcode.badge(f"{count} 个扩展", "gray")}

{self.table_gen.generate_category_table_zh(cat_extensions)}
'''
            zh_category_sections.append(section)
        
        # Generate the shortcode badges for navigation
        nav_badges = []
        for category in ['TIME', 'GIS', 'RAG', 'FTS', 'OLAP', 'FEAT', 'LANG', 'TYPE', 
                        'UTIL', 'FUNC', 'ADMIN', 'STAT', 'SEC', 'FDW', 'SIM', 'ETL']:
            if category in CATEGORY_META:
                nav_badges.append(HugoShortcode.category(category))
        
        zh_content = f'''---
title: 按分类
description: 按功能分类组织的 PostgreSQL 扩展
---

{' '.join(nav_badges[:8])}

{' '.join(nav_badges[8:])}

## 概览

{zh_category_overview_table}

{''.join(zh_category_sections)}
'''
        
        self.write_content('cate.zh.md', zh_content)


class LicenseListGenerator(ContentGenerator):
    """Generate license-based extension list."""
    
    def generate(self):
        print("Generating license list...")
        
        # Count extensions by license and normalize names
        license_counts = Counter()
        license_extensions_map = defaultdict(list)
        
        for ext in self.extensions:
            if not ext.license:
                continue
            normalized_name = normalize_license_name(ext.license)
            license_counts[normalized_name] += 1
            license_extensions_map[normalized_name].append(ext)
        
        # Generate summary table with Hugo shortcodes
        summary_rows = []
        for license_name, count in license_counts.most_common():
            info = LICENSE_INFO.get(license_name, {'url': '#', 'description': 'Open source license.'})
            license_badge = HugoShortcode.license(license_name)
            summary_rows.append(f'| {license_badge} | {count} | [License Text]({info["url"]}) | {info["description"]} |')
        
        summary_table = f'''| License | Count | Reference |  Description |
|:--------|:-----:|:-------:|:----------|
{chr(10).join(summary_rows)}'''
        
        # Generate license sections
        license_sections = []
        for license_name, count in license_counts.most_common():
            license_extensions = sorted(license_extensions_map[license_name], key=lambda e: e.name)
            info = LICENSE_INFO.get(license_name, {'url': '#', 'description': 'Open source license.'})
            
            section = f'''
## {license_name}

{HugoShortcode.license(license_name)} {HugoShortcode.badge(f"{count} Extensions", "gray")}

[{license_name} License Text]({info["url"]}) : {info["description"]}

{self.table_gen.generate_simple_table(license_extensions)}
'''
            license_sections.append(section)
        
        # Generate navigation badges with Hugo shortcodes
        nav_badges = []
        for license_name in ['MIT', 'ISC', 'PostgreSQL', 'BSD 0-Clause', 'BSD 2-Clause', 'BSD 3-Clause',
                             'Artistic', 'Apache-2.0', 'MPL-2.0', 'GPL-2.0', 'GPL-3.0', 
                             'LGPL-2.1', 'LGPL-3.0', 'AGPL-3.0', 'Timescale']:
            if license_name in license_counts:
                nav_badges.append(HugoShortcode.license(license_name))
        
        content = f'''---
title: By License
description: PostgreSQL extensions organized by open source license
---

{' '.join(nav_badges[:9])}

{' '.join(nav_badges[9:])}

## Summary

{summary_table}

---------

{''.join(license_sections)}
'''
        
        self.write_content('license.md', content)
        self._generate_chinese_version(license_counts, license_extensions_map)
    
    def _generate_chinese_version(self, license_counts, license_extensions_map):
        """Generate Chinese version of license list."""
        print("Generating Chinese license list...")
        
        # Generate Chinese summary table
        zh_summary_rows = []
        for license_name, count in license_counts.most_common():
            info = LICENSE_INFO.get(license_name, {'url': '#', 'description': '开源许可证。'})
            license_badge = HugoShortcode.license(license_name)
            zh_summary_rows.append(f'| {license_badge} | {count} | [许可证文本]({info["url"]}) | {info["description"]} |')
        
        zh_summary_table = f'''| 许可证 | 数量 | 参考 | 描述 |
|:--------|:-----:|:-------:|:----------|
{chr(10).join(zh_summary_rows)}'''
        
        # Generate Chinese license sections
        zh_license_sections = []
        for license_name, count in license_counts.most_common():
            license_extensions = sorted(license_extensions_map[license_name], key=lambda e: e.name)
            info = LICENSE_INFO.get(license_name, {'url': '#', 'description': '开源许可证。'})
            
            # Generate Chinese table
            headers = ['ID', '扩展', '包', '描述']
            rows = [self.table_gen._format_table_header(headers, [':---:',':---',':---',':---'])]
            
            for ext in license_extensions:
                package_cell = f'[`{ext.pkg}`](/zh/e/{self.table_gen.leading_map.get(ext.pkg, ext.name)})'
                row_data = [
                    str(ext.id),
                    f'[`{ext.name}`](/zh/e/{ext.name})',
                    package_cell,
                    ext.zh_desc or ext.en_desc or '暂无描述'
                ]
                rows.append('| ' + ' | '.join(row_data) + ' |')
            
            zh_table = '\n'.join(rows)
            
            section = f'''
## {license_name}

{HugoShortcode.license(license_name)} {HugoShortcode.badge(f"{count} 个扩展", "gray")}

[{license_name} 许可证文本]({info["url"]}) : {info["description"]}

{zh_table}
'''
            zh_license_sections.append(section)
        
        # Generate navigation badges with Hugo shortcodes
        nav_badges = []
        for license_name in ['MIT', 'ISC', 'PostgreSQL', 'BSD 0-Clause', 'BSD 2-Clause', 'BSD 3-Clause',
                             'Artistic', 'Apache-2.0', 'MPL-2.0', 'GPL-2.0', 'GPL-3.0', 
                             'LGPL-2.1', 'LGPL-3.0', 'AGPL-3.0', 'Timescale']:
            if license_name in license_counts:
                nav_badges.append(HugoShortcode.license(license_name))
        
        zh_content = f'''---
title: 按许可证
description: 按开源许可证组织的 PostgreSQL 扩展
---

{' '.join(nav_badges[:9])}

{' '.join(nav_badges[9:])}

## 概览

{zh_summary_table}

---------

{''.join(zh_license_sections)}
'''
        
        self.write_content('license.zh.md', zh_content)


class LanguageListGenerator(ContentGenerator):
    """Generate language-based extension list."""
    
    def generate(self):
        print("Generating language list...")
        
        # Count extensions by language
        language_counts = Counter(ext.lang for ext in self.extensions if ext.lang)
        
        # Generate summary table with Hugo shortcodes
        summary_rows = []
        for lang, count in language_counts.most_common():
            desc = LANGUAGE_DESCRIPTIONS.get(lang, f'Extensions written in {lang}')
            lang_badge = HugoShortcode.language(lang)
            summary_rows.append(f'| {lang_badge} | {count} | {desc} |')
        
        summary_table = f'''| Language | Count | Description |
|:-------:|:-----:|:------------|
{chr(10).join(summary_rows)}'''
        
        # Generate language sections
        language_sections = []
        for lang, count in language_counts.most_common():
            lang_extensions = [ext for ext in self.extensions if ext.lang == lang]
            lang_extensions.sort(key=lambda e: e.name)
            
            desc = LANGUAGE_DESCRIPTIONS.get(lang, f'Extensions written in {lang}')
            
            section = f'''
## {lang}

{HugoShortcode.language(lang)} {HugoShortcode.badge(f"{count} Extensions", "gray")}

{desc}

{self.table_gen.generate_simple_table(lang_extensions)}
'''
            language_sections.append(section)
        
        # Generate navigation badges
        nav_badges = []
        for lang in ['C', 'C++', 'Rust', 'Java', 'Python', 'SQL', 'Data']:
            if lang in language_counts:
                nav_badges.append(HugoShortcode.language(lang))
        
        content = f'''---
title: By Language
description: PostgreSQL extensions organized by implementation language
---

{' '.join(nav_badges)}

## Summary

{summary_table}

{''.join(language_sections)}
'''
        
        self.write_content('lang.md', content)
        self._generate_chinese_version(language_counts)
    
    def _generate_chinese_version(self, language_counts):
        """Generate Chinese version of language list."""
        print("Generating Chinese language list...")
        
        zh_lang_descriptions = {
            'C': '传统的 PostgreSQL 扩展开发语言',
            'C++': '使用 C++ 特性和库的扩展',
            'Rust': '使用 pgrx 框架用 Rust 编写的扩展',
            'Python': '使用 Python 编写的扩展',
            'SQL': '纯 SQL 扩展和函数',
            'Java': '在 JVM 上运行的扩展',
            'Data': '仅包含数据的扩展',
        }
        
        # Generate Chinese summary table
        zh_summary_rows = []
        for lang, count in language_counts.most_common():
            desc = zh_lang_descriptions.get(lang, f'使用 {lang} 编写的扩展')
            lang_badge = HugoShortcode.language(lang)
            zh_summary_rows.append(f'| {lang_badge} | {count} | {desc} |')
        
        zh_summary_table = f'''| 语言 | 数量 | 描述 |
|:-------:|:-----:|:------------|
{chr(10).join(zh_summary_rows)}'''
        
        # Generate Chinese language sections
        zh_language_sections = []
        for lang, count in language_counts.most_common():
            lang_extensions = [ext for ext in self.extensions if ext.lang == lang]
            lang_extensions.sort(key=lambda e: e.name)
            
            desc = zh_lang_descriptions.get(lang, f'使用 {lang} 编写的扩展')
            
            # Generate Chinese table
            headers = ['ID', '扩展', '包', '描述']
            rows = [self.table_gen._format_table_header(headers, [':---:',':---',':---',':---'])]
            
            for ext in lang_extensions:
                package_cell = f'[`{ext.pkg}`](/zh/e/{self.table_gen.leading_map.get(ext.pkg, ext.name)})'
                row_data = [
                    str(ext.id),
                    f'[`{ext.name}`](/zh/e/{ext.name})',
                    package_cell,
                    ext.zh_desc or ext.en_desc or '暂无描述'
                ]
                rows.append('| ' + ' | '.join(row_data) + ' |')
            
            zh_table = '\n'.join(rows)
            
            section = f'''
## {lang}

{HugoShortcode.language(lang)} {HugoShortcode.badge(f"{count} 个扩展", "gray")}

{desc}

{zh_table}
'''
            zh_language_sections.append(section)
        
        # Generate navigation badges
        nav_badges = []
        for lang in ['C', 'C++', 'Rust', 'Java', 'Python', 'SQL', 'Data']:
            if lang in language_counts:
                nav_badges.append(HugoShortcode.language(lang))
        
        zh_content = f'''---
title: 按语言
description: 按实现语言组织的 PostgreSQL 扩展
---

{' '.join(nav_badges)}

## 概览

{zh_summary_table}

{''.join(zh_language_sections)}
'''
        
        self.write_content('lang.zh.md', zh_content)


class RepoListGenerator(ContentGenerator):
    """Generate repository-based extension list."""
    
    def generate(self):
        print("Generating repository list...")
        
        # Categorize extensions
        contrib_extensions = [ext for ext in self.extensions if ext.contrib]
        both_extensions = []
        el_only_extensions = []
        debian_only_extensions = []
        
        for ext in self.extensions:
            if ext.contrib:
                continue
                
            has_rpm = ext.has_rpm
            has_deb = ext.has_deb
            
            if has_rpm and has_deb:
                both_extensions.append(ext)
            elif has_rpm and not has_deb:
                el_only_extensions.append(ext)
            elif not has_rpm and has_deb:
                debian_only_extensions.append(ext)
        
        # Sort all extension lists by ID
        for ext_list in [contrib_extensions, both_extensions, el_only_extensions, debian_only_extensions]:
            ext_list.sort(key=lambda e: e.id)
        
        content = f'''---
title: By Repository
description: PostgreSQL extensions organized by distribution repository
---

## CONTRIB Extensions

Extensions included in standard PostgreSQL installation.

{HugoShortcode.badge(f"{len(contrib_extensions)} Extensions", "purple")}

{self.table_gen.generate_simple_table(contrib_extensions)}

## Both EL and Debian

Extensions available in both Enterprise Linux and Debian/Ubuntu repositories.

{HugoShortcode.badge(f"{len(both_extensions)} Extensions", "green")}

{self.table_gen.generate_repo_table(both_extensions)}

## EL Only

Extensions only available in Enterprise Linux (RHEL/CentOS/Rocky/Alma) repositories.

{HugoShortcode.badge(f"{len(el_only_extensions)} Extensions", "amber")}

{self.table_gen.generate_repo_table(el_only_extensions)}

## Debian Only

Extensions only available in Debian/Ubuntu repositories.

{HugoShortcode.badge(f"{len(debian_only_extensions)} Extensions", "blue")}

{self.table_gen.generate_repo_table(debian_only_extensions)}
'''
        
        self.write_content('repo.md', content)
        self._generate_chinese_version(contrib_extensions, both_extensions, el_only_extensions, debian_only_extensions)
    
    def _generate_chinese_version(self, contrib_extensions, both_extensions, el_only_extensions, debian_only_extensions):
        """Generate Chinese version of repository list."""
        print("Generating Chinese repository list...")
        
        content = f'''---
title: 按仓库
description: 按发行仓库组织的 PostgreSQL 扩展
---

## CONTRIB 扩展

PostgreSQL 标准安装中包含的扩展。

{HugoShortcode.badge(f"{len(contrib_extensions)} 个扩展", "purple")}

{self.table_gen.generate_simple_table(contrib_extensions)}

## EL 和 Debian 都有

在 Enterprise Linux 和 Debian/Ubuntu 仓库中都可用的扩展。

{HugoShortcode.badge(f"{len(both_extensions)} 个扩展", "green")}

{self.table_gen.generate_repo_table(both_extensions)}

## 仅 EL

仅在 Enterprise Linux (RHEL/CentOS/Rocky/Alma) 仓库中可用的扩展。

{HugoShortcode.badge(f"{len(el_only_extensions)} 个扩展", "amber")}

{self.table_gen.generate_repo_table(el_only_extensions)}

## 仅 Debian

仅在 Debian/Ubuntu 仓库中可用的扩展。

{HugoShortcode.badge(f"{len(debian_only_extensions)} 个扩展", "blue")}

{self.table_gen.generate_repo_table(debian_only_extensions)}
'''
        
        self.write_content('repo.zh.md', content)


class AttributeListGenerator(ContentGenerator):
    """Generate attribute-based extension list."""
    
    def generate(self):
        print("Generating attribute list...")
        
        # Find extensions with special attributes
        need_loading = [ext for ext in self.extensions if ext.need_load]
        without_ddl = [ext for ext in self.extensions if not ext.need_ddl]
        has_dependency = [ext for ext in self.extensions if ext.requires]
        not_relocatable = [ext for ext in self.extensions if not ext.relocatable]
        not_trusted = [ext for ext in self.extensions if not ext.trusted]
        
        # Sort all lists
        for ext_list in [need_loading, without_ddl, has_dependency, not_relocatable, not_trusted]:
            ext_list.sort(key=lambda e: e.name)
        
        content = f'''---
title: By Attribute
description: PostgreSQL extensions organized by special attributes
---

## Extensions Requiring Loading

Extensions that must be loaded via `shared_preload_libraries` or `session_preload_libraries`.

{HugoShortcode.badge(f"{len(need_loading)} Extensions", "red")}

{self.table_gen.generate_simple_table(need_loading)}

## Extensions Without DDL

Extensions that don't require `CREATE EXTENSION` statement.

{HugoShortcode.badge(f"{len(without_ddl)} Extensions", "amber")}

{self.table_gen.generate_simple_table(without_ddl)}

## Extensions With Dependencies

Extensions that depend on other extensions.

{HugoShortcode.badge(f"{len(has_dependency)} Extensions", "blue")}

{self.table_gen.generate_simple_table(has_dependency)}

## Non-Relocatable Extensions

Extensions that cannot be moved to a different schema after creation.

{HugoShortcode.badge(f"{len(not_relocatable)} Extensions", "purple")}

{self.table_gen.generate_simple_table(not_relocatable)}

## Non-Trusted Extensions

Extensions that require superuser privileges to create.

{HugoShortcode.badge(f"{len(not_trusted)} Extensions", "gray")}

{self.table_gen.generate_simple_table(not_trusted)}
'''
        
        self.write_content('attr.md', content)
        self._generate_chinese_version(need_loading, without_ddl, has_dependency, not_relocatable, not_trusted)
    
    def _generate_chinese_version(self, need_loading, without_ddl, has_dependency, not_relocatable, not_trusted):
        """Generate Chinese version of attribute list."""
        print("Generating Chinese attribute list...")
        
        content = f'''---
title: 按属性
description: 按特殊属性组织的 PostgreSQL 扩展
---

## 需要加载的扩展

必须通过 `shared_preload_libraries` 或 `session_preload_libraries` 加载的扩展。

{HugoShortcode.badge(f"{len(need_loading)} 个扩展", "red")}

{self.table_gen.generate_simple_table(need_loading)}

## 无需 DDL 的扩展

不需要 `CREATE EXTENSION` 语句的扩展。

{HugoShortcode.badge(f"{len(without_ddl)} 个扩展", "amber")}

{self.table_gen.generate_simple_table(without_ddl)}

## 有依赖的扩展

依赖于其他扩展的扩展。

{HugoShortcode.badge(f"{len(has_dependency)} 个扩展", "blue")}

{self.table_gen.generate_simple_table(has_dependency)}

## 不可重定位的扩展

创建后无法移动到不同模式的扩展。

{HugoShortcode.badge(f"{len(not_relocatable)} 个扩展", "purple")}

{self.table_gen.generate_simple_table(not_relocatable)}

## 非受信任的扩展

需要超级用户权限才能创建的扩展。

{HugoShortcode.badge(f"{len(not_trusted)} 个扩展", "gray")}

{self.table_gen.generate_simple_table(not_trusted)}
'''
        
        self.write_content('attr.zh.md', content)


class LinuxListGenerator(ContentGenerator):
    """Generate Linux distribution-based extension list."""
    
    def generate(self):
        print("Generating Linux distribution list...")
        
        # Platform groups
        platform_groups = {
            'el8.x86_64': [], 'el8.aarch64': [],
            'el9.x86_64': [], 'el9.aarch64': [],
            'd12.x86_64': [], 'd12.aarch64': [],
            'u22.x86_64': [], 'u22.aarch64': [],
            'u24.x86_64': [], 'u24.aarch64': []
        }
        
        # Categorize extensions by platform availability
        for ext in self.extensions:
            for platform in platform_groups.keys():
                if platform.startswith('el'):
                    if ext.has_rpm:
                        platform_groups[platform].append(ext)
                else:
                    if ext.has_deb:
                        platform_groups[platform].append(ext)
        
        # Generate platform sections
        platform_sections = []
        for platform, extensions in platform_groups.items():
            extensions.sort(key=lambda e: e.name)
            os_desc = OS_DESCRIPTIONS.get(platform.split('.')[0], platform)
            arch = platform.split('.')[1]
            
            section = f'''
## {os_desc} ({arch})

{HugoShortcode.badge(f"{len(extensions)} Extensions", "blue")}

{self.table_gen.generate_linux_table(extensions, platform)}
'''
            platform_sections.append(section)
        
        content = f'''---
title: By Linux Distribution
description: PostgreSQL extension availability across Linux distributions
---

## Overview

Extension availability varies across different Linux distributions and architectures.

{''.join(platform_sections)}
'''
        
        self.write_content('linux.md', content)
        self._generate_chinese_version(platform_groups)
    
    def _generate_chinese_version(self, platform_groups):
        """Generate Chinese version of Linux list."""
        print("Generating Chinese Linux distribution list...")
        
        zh_os_descriptions = {
            'el8': 'Enterprise Linux 8 (RHEL 8, CentOS 8, Rocky 8, Alma 8)',
            'el9': 'Enterprise Linux 9 (RHEL 9, CentOS 9, Rocky 9, Alma 9)',
            'd12': 'Debian 12 (Bookworm)',
            'u22': 'Ubuntu 22.04 LTS (Jammy)',
            'u24': 'Ubuntu 24.04 LTS (Noble)'
        }
        
        # Generate platform sections
        platform_sections = []
        for platform, extensions in platform_groups.items():
            extensions.sort(key=lambda e: e.name)
            os_desc = zh_os_descriptions.get(platform.split('.')[0], platform)
            arch = platform.split('.')[1]
            
            section = f'''
## {os_desc} ({arch})

{HugoShortcode.badge(f"{len(extensions)} 个扩展", "blue")}

{self.table_gen.generate_linux_table(extensions, platform)}
'''
            platform_sections.append(section)
        
        content = f'''---
title: 按 Linux 发行版
description: PostgreSQL 扩展在各 Linux 发行版上的可用性
---

## 概览

扩展的可用性因不同的 Linux 发行版和架构而异。

{''.join(platform_sections)}
'''
        
        self.write_content('linux.zh.md', content)


class PGSQLListGenerator(ContentGenerator):
    """Generate PostgreSQL version-based extension list."""
    
    def generate(self):
        print("Generating PostgreSQL version list...")
        
        pg_sections = []
        for pg in self.config.PG_VERSIONS:
            # Find extensions NOT available for this PG version
            unavailable = [ext for ext in self.extensions if str(pg) not in ext.pg_ver]
            available = [ext for ext in self.extensions if str(pg) in ext.pg_ver]
            unavailable.sort(key=lambda e: e.name)
            
            section = f'''
## PostgreSQL {pg}

{HugoShortcode.badge(f"{len(available)} Available", "green")} {HugoShortcode.badge(f"{len(unavailable)} Unavailable", "red")}

### Extensions NOT available for PostgreSQL {pg}

{self.table_gen.generate_simple_table(unavailable) if unavailable else "All extensions are available for this PostgreSQL version."}
'''
            pg_sections.append(section)
        
        content = f'''---
title: By PostgreSQL Version
description: PostgreSQL extension compatibility by major version
---

## Overview

Check which extensions are available for each PostgreSQL major version.

{''.join(pg_sections)}
'''
        
        self.write_content('pgsql.md', content)
        self._generate_chinese_version()
    
    def _generate_chinese_version(self):
        """Generate Chinese version of PostgreSQL version list."""
        print("Generating Chinese PostgreSQL version list...")
        
        pg_sections = []
        for pg in self.config.PG_VERSIONS:
            # Find extensions NOT available for this PG version
            unavailable = [ext for ext in self.extensions if str(pg) not in ext.pg_ver]
            available = [ext for ext in self.extensions if str(pg) in ext.pg_ver]
            unavailable.sort(key=lambda e: e.name)
            
            section = f'''
## PostgreSQL {pg}

{HugoShortcode.badge(f"{len(available)} 个可用", "green")} {HugoShortcode.badge(f"{len(unavailable)} 个不可用", "red")}

### PostgreSQL {pg} 不可用的扩展

{self.table_gen.generate_simple_table(unavailable) if unavailable else "所有扩展都可用于此 PostgreSQL 版本。"}
'''
            pg_sections.append(section)
        
        content = f'''---
title: 按 PostgreSQL 版本
description: 按主版本的 PostgreSQL 扩展兼容性
---

## 概览

检查每个 PostgreSQL 主版本可用的扩展。

{''.join(pg_sections)}
'''
        
        self.write_content('pgsql.zh.md', content)


class CategoryIndexGenerator(ContentGenerator):
    """Generate individual category index pages."""
    
    def generate(self):
        print("Generating individual category pages...")
        
        # Create cate directory if it doesn't exist
        cate_dir = os.path.join(self.config.OUTPUT_DIR, 'cate')
        os.makedirs(cate_dir, exist_ok=True)
        
        # Group extensions by category
        category_groups = defaultdict(list)
        for ext in self.extensions:
            category_groups[ext.category].append(ext)
        
        # Generate a page for each category
        for category_code, cat_extensions in category_groups.items():
            if category_code not in CATEGORY_META:
                continue
            
            cat_extensions.sort(key=lambda e: e.id)
            meta = CATEGORY_META[category_code]
            
            # Generate English version
            en_content = f'''---
title: {category_code} Extensions
description: {meta["en_desc"]}
---

{HugoShortcode.category(category_code)} {HugoShortcode.badge(f"{len(cat_extensions)} Extensions", "gray")}

{meta["en_desc"]}

## Extensions

{self.table_gen.generate_category_table(cat_extensions)}

## Details

'''
            
            # Add detailed descriptions for each extension
            for ext in cat_extensions[:10]:  # Show details for first 10
                en_content += f'''
### {ext.name}

{HugoShortcode.ext(ext.name)} - {ext.en_desc or "No description"}

- **Version**: {ext.version or "N/A"}
- **License**: {HugoShortcode.license(ext.license) if ext.license else "N/A"}
- **Language**: {HugoShortcode.language(ext.lang) if ext.lang else "N/A"}
- **Package**: [`{ext.pkg}`](/e/{self.table_gen.leading_map.get(ext.pkg, ext.name)})
'''
            
            # Write English version
            en_path = os.path.join(cate_dir, f'{category_code.lower()}.md')
            with open(en_path, 'w', encoding='utf-8') as f:
                f.write(en_content)
            print(f"Generated: {en_path}")
            
            # Generate Chinese version
            zh_content = f'''---
title: {category_code} 扩展
description: {meta["zh_desc"]}
---

{HugoShortcode.category(category_code)} {HugoShortcode.badge(f"{len(cat_extensions)} 个扩展", "gray")}

{meta["zh_desc"]}

## 扩展列表

{self.table_gen.generate_category_table_zh(cat_extensions)}

## 详细信息

'''
            
            # Add detailed descriptions for each extension
            for ext in cat_extensions[:10]:  # Show details for first 10
                zh_content += f'''
### {ext.name}

{HugoShortcode.ext(ext.name)} - {ext.zh_desc or ext.en_desc or "暂无描述"}

- **版本**: {ext.version or "N/A"}
- **许可证**: {HugoShortcode.license(ext.license) if ext.license else "N/A"}
- **语言**: {HugoShortcode.language(ext.lang) if ext.lang else "N/A"}
- **包**: [`{ext.pkg}`](/zh/e/{self.table_gen.leading_map.get(ext.pkg, ext.name)})
'''
            
            # Write Chinese version
            zh_path = os.path.join(cate_dir, f'{category_code.lower()}.zh.md')
            with open(zh_path, 'w', encoding='utf-8') as f:
                f.write(zh_content)
            print(f"Generated: {zh_path}")


# =============================================================================
# MAIN APPLICATION
# =============================================================================

class ExtensionListGenerator:
    """Main application class."""
    
    def __init__(self, config: Config = None):
        self.config = config or Config()
        self.db_manager = DatabaseManager(self.config.DB_CONNECTION)
        self.extensions = []
        self.packages = []
        self.leading_map = {}
        self.table_gen = None
    
    def run(self):
        """Run the complete generation process."""
        self._setup()
        self._generate_all_lists()
        print("List generation complete!")
    
    def _setup(self):
        """Setup data and dependencies."""
        # Load category metadata from CSV
        global CATEGORY_META
        CATEGORY_META = self.db_manager.load_category_metadata()
        
        # Load data from database
        self.extensions = self.db_manager.load_extensions()
        self.packages = self.db_manager.load_packages()
        
        # Load packages for all extensions
        for ext in self.extensions:
            ext.load_packages(self.packages)
        
        # Build leading extension map
        self.leading_map = self._build_leading_map()
        
        # Initialize table generator
        self.table_gen = TableGenerator(self.leading_map)
        
        # Ensure output directory exists
        os.makedirs(self.config.OUTPUT_DIR, exist_ok=True)
    
    def _build_leading_map(self) -> Dict[str, str]:
        """Build a mapping from package names to their leading extensions."""
        leading_map = {}
        for ext in self.extensions:
            if ext.lead and ext.pkg:
                leading_map[ext.pkg] = ext.name
        return leading_map
    
    def _generate_all_lists(self):
        """Generate all list pages."""
        generators = [
            MainIndexGenerator(self.config, self.extensions, self.table_gen),
            CategoryListGenerator(self.config, self.extensions, self.table_gen),
            CategoryIndexGenerator(self.config, self.extensions, self.table_gen),
            LicenseListGenerator(self.config, self.extensions, self.table_gen),
            LanguageListGenerator(self.config, self.extensions, self.table_gen),
            RepoListGenerator(self.config, self.extensions, self.table_gen),
            AttributeListGenerator(self.config, self.extensions, self.table_gen),
            LinuxListGenerator(self.config, self.extensions, self.table_gen),
            PGSQLListGenerator(self.config, self.extensions, self.table_gen),
        ]
        
        for generator in generators:
            generator.generate()


def main():
    """Main entry point."""
    config = Config()
    app = ExtensionListGenerator(config)
    app.run()


if __name__ == "__main__":
    main()