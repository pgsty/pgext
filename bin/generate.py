#!/usr/bin/env python3
"""
PostgreSQL Extension Hugo Content Generator

This script generates Hugo/Hextra-compatible Markdown files for PostgreSQL extensions
by reading directly from the database and applying templates.

Usage:
    python gen-ext2.py [PGURL] [OUTPUT_DIR]
    
    PGURL: PostgreSQL connection URL (default: postgres:///vonng)
    OUTPUT_DIR: Output directory for generated files (default: ../hugo/content/e/)
"""

import os
import sys
import argparse
import psycopg2
from typing import Dict, List, Optional, Any, Tuple
from dataclasses import dataclass, field
from pathlib import Path
import logging
from datetime import datetime

# Configure logging
logging.basicConfig(level=logging.INFO, format='%(levelname)s: %(message)s')
logger = logging.getLogger(__name__)

# Configuration
@dataclass
class Config:
    """Configuration for extension generation"""
    pg_versions: List[int] = field(default_factory=lambda: [18, 17, 16, 15, 14, 13])
    os_systems: List[Tuple[str, str]] = field(default_factory=lambda: [
        ('el8.x86_64', 'EL8 x86_64'),
        ('el8.aarch64', 'EL8 ARM64'),
        ('el9.x86_64', 'EL9 x86_64'),
        ('el9.aarch64', 'EL9 ARM64'),
        ('d12.x86_64', 'Debian 12 x86_64'),
        ('d12.aarch64', 'Debian 12 ARM64'),
        ('u22.x86_64', 'Ubuntu 22.04 x86_64'),
        ('u22.aarch64', 'Ubuntu 22.04 ARM64'),
        ('u24.x86_64', 'Ubuntu 24.04 x86_64'),
        ('u24.aarch64', 'Ubuntu 24.04 ARM64'),
    ])
    template_path: str = 'src/extension.md'
    default_pgurl: str = 'postgres:///vonng'
    default_output_dir: str = 'content/e/'

class ExtensionGenerator:
    """Generates Hugo content for PostgreSQL extensions"""
    
    def __init__(self, pgurl: str, output_dir: str, config: Config = None):
        self.pgurl = pgurl
        self.output_dir = Path(output_dir)
        self.config = config or Config()
        self.conn = None
        self.template = None
        
    def connect(self):
        """Establish database connection"""
        try:
            self.conn = psycopg2.connect(self.pgurl)
            logger.info(f"Connected to database")
        except Exception as e:
            logger.error(f"Failed to connect to database: {e}")
            raise
            
    def disconnect(self):
        """Close database connection"""
        if self.conn:
            self.conn.close()
            logger.info("Disconnected from database")
            
    def load_template(self):
        """Load the Markdown template"""
        template_path = Path(__file__).parent.parent / self.config.template_path
        try:
            with open(template_path, 'r', encoding='utf-8') as f:
                self.template = f.read()
            logger.info(f"Loaded template from {template_path}")
        except Exception as e:
            logger.error(f"Failed to load template: {e}")
            raise
    
    def parse_array(self, value: str) -> List[str]:
        """Parse PostgreSQL array string to Python list"""
        if isinstance(value, list):
            return value
        if not value or not value.startswith('{') or not value.endswith('}'):
            return []
        return [item.strip() for item in value[1:-1].split(',') if item.strip()]
    
    def get_extension_data(self, ext_name: str) -> Optional[Dict[str, Any]]:
        """Fetch extension data from database"""
        with self.conn.cursor() as cur:
            cur.execute("""
                SELECT id, name, pkg, lead_ext, category, state, url, license, tags, version, repo, lang, 
                       contrib, lead, has_bin, has_lib, need_ddl, need_load, trusted, relocatable, schemas, 
                       pg_ver, requires, require_by, see_also, rpm_ver, rpm_repo, rpm_pkg, rpm_pg, rpm_deps, 
                       deb_ver, deb_repo, deb_pkg, deb_deps, deb_pg, source, extra, en_desc, zh_desc, comment, mtime
                FROM pgext.extension WHERE name = %s
            """, (ext_name,))
            row = cur.fetchone()
            
            if not row:
                return None
                
            columns = [
                'id', 'name', 'pkg', 'lead_ext', 'category', 'state', 'url', 'license', 'tags', 'version', 'repo', 'lang',
                'contrib', 'lead', 'has_bin', 'has_lib', 'need_ddl', 'need_load', 'trusted', 'relocatable', 'schemas',
                'pg_ver', 'requires', 'require_by', 'see_also', 'rpm_ver', 'rpm_repo', 'rpm_pkg', 'rpm_pg', 'rpm_deps',
                'deb_ver', 'deb_repo', 'deb_pkg', 'deb_deps', 'deb_pg', 'source', 'extra', 'en_desc', 'zh_desc', 'comment', 'mtime'
            ]
            
            ext_data = dict(zip(columns, row))
            
            # Parse array fields
            array_fields = ['tags', 'schemas', 'pg_ver', 'requires', 'require_by', 'see_also', 'rpm_pg', 'rpm_deps', 'deb_pg', 'deb_deps']
            for field in array_fields:
                if ext_data[field]:
                    ext_data[field] = self.parse_array(ext_data[field])
                else:
                    ext_data[field] = []
                    
            return ext_data
    
    def get_siblings(self, pkg_name: str) -> List[str]:
        """Get all extensions in the same package"""
        with self.conn.cursor() as cur:
            cur.execute("SELECT name FROM pgext.extension WHERE pkg = %s ORDER BY name", (pkg_name,))
            return [row[0] for row in cur.fetchall()]
    
    def get_matrix_data(self, pkg_name: str) -> List[Dict[str, Any]]:
        """Get availability matrix data"""
        with self.conn.cursor() as cur:
            cur.execute("""
                SELECT pg, os, type, os_code, os_arch, pkg, ext, pname, miss, hide, pkg_repo, pkg_ver, count
                FROM pgext.matrix WHERE pkg = %s ORDER BY pg DESC, os
            """, (pkg_name,))
            
            columns = ['pg', 'os', 'type', 'os_code', 'os_arch', 'pkg', 'ext', 'pname', 'miss', 'hide', 'pkg_repo', 'pkg_ver', 'count']
            return [dict(zip(columns, row)) for row in cur.fetchall()]
    
    def get_availability_data(self, pkg_name: str) -> List[Dict[str, Any]]:
        """Get package availability data"""
        with self.conn.cursor() as cur:
            cur.execute("""
                SELECT pkg, ext, pname, os, pg, name, ver, org, type, os_code, os_arch, repo, 
                       version, release, file, sha256, url, mirror_url, size, size_full
                FROM pgext.availability WHERE pkg = %s ORDER BY pkg, pname, os, pg
            """, (pkg_name,))
            
            columns = ['pkg', 'ext', 'pname', 'os', 'pg', 'name', 'ver', 'org', 'type', 'os_code', 'os_arch', 'repo',
                       'version', 'release', 'file', 'sha256', 'url', 'mirror_url', 'size', 'size_full']
            return [dict(zip(columns, row)) for row in cur.fetchall()]
    
    def format_badge(self, content: str, color: str = "default", alt: str = "", link: str = "") -> str:
        """Format a badge shortcode"""
        parts = ['{{< badge']
        parts.append(f'content="{content}"')
        parts.append(f'color="{color}"')
        if alt:
            parts.append(f'alt="{alt}"')
        if link:
            parts.append(f'link="{link}"')
        return ' '.join(parts) + ' >}}'
        
    def format_pkg(self, pname: str, version: str = "", repo: str = "", url: str = "") -> str:
        """Format a package shortcode"""
        parts = ['{{< pkg', f'"{pname}"']
        if version:
            parts.append(f'"{version}"')
        if repo:
            parts.append(f'"{repo}"')
        if url:
            parts.append(f'"{url}"')
        return ' '.join(parts) + ' >}}'
            
    def format_ext(self, ext_name: str) -> str:
        """Format an extension link shortcode"""
        return '{{< ext "' + ext_name + '" >}}'
    
    def generate_attributes_section(self, ext_data: Dict[str, Any]) -> str:
        """Generate the attributes section"""
        # Create attribute badges
        attrs = []
        
        # Build attribute string like in vector.md
        attr_string = ""
        if ext_data.get('contrib'):
            attr_string += "c"
        else:
            attr_string += "-"
            
        if ext_data.get('lead'):
            attr_string += "l"
        else:
            attr_string += "-"
            
        if ext_data.get('has_bin'):
            attr_string += "b"
        else:
            attr_string += "-"
            
        if ext_data.get('has_lib'):
            attr_string += "s"
        else:
            attr_string += "-"
            
        if ext_data.get('need_load'):
            attr_string += "L"
        else:
            attr_string += "-"
            
        if ext_data.get('need_ddl'):
            attr_string += "d"
        else:
            attr_string += "-"
            
        if ext_data.get('trusted'):
            attr_string += "t"
        else:
            attr_string += "-"
            
        if ext_data.get('relocatable'):
            attr_string += "r"
        else:
            attr_string += "-"
            
        # Create the table
        section = "|            Attribute                                    |              Has Binary              |                Has Library                |                 Loading                  |               Has DDL                |                Relocatable                |                Trusted                 |\n"
        section += "|:------------------------------------------------------:|:--------------------------------------:|:----------------------------------------:|:----------------------------------------:|:----------------------------------------:|:----------------------------------------:|:-----------------------------------------:|:-----------------------------------------:|\n"
        section += '|      {{< badge content="--' + attr_string[2:] + '" color="blue" >}} | '
        
        # Has Binary
        if ext_data.get('has_bin'):
            section += '{{< badge content="Yes" color="green" >}}'
        else:
            section += '{{< badge content="No" color="green" >}}'
        section += '     | '
        
        # Has Library
        if ext_data.get('has_lib'):
            section += '{{< badge content="Yes" color="green" >}}'
        else:
            section += '{{< badge content="No" color="green" >}}'
        section += ' | '
        
        # Loading
        if ext_data.get('need_load'):
            section += '{{< badge content="Yes" color="red" >}}'
        else:
            section += '{{< badge content="No" color="green" >}}'
        section += ' | '
        
        # Has DDL
        if ext_data.get('need_ddl'):
            section += '{{< badge content="Yes" color="green" >}}'
        else:
            section += '{{< badge content="No" color="green" >}}'
        section += ' | '
        
        # Relocatable
        if ext_data.get('relocatable'):
            section += '{{< badge content="yes" color="green" >}}'
        else:
            section += '{{< badge content="no" color="red" >}}'
        section += ' | '
        
        # Trusted
        if ext_data.get('trusted'):
            section += '{{< badge content="yes" color="green" >}}'
        else:
            section += '{{< badge content="no" color="red" >}}'
        section += '     |'
        
        return section
    
    def generate_relationships_section(self, ext_data: Dict[str, Any]) -> str:
        """Generate the relationships section"""
        if not ext_data.get('requires') and not ext_data.get('require_by') and not ext_data.get('see_also'):
            return ""
            
        section = "\n| **Relationships** |                                                                   |\n"
        section += "|:-----------------:|:------------------------------------------------------------------|\n"
        
        if ext_data.get('requires'):
            requires = ' '.join([self.format_ext(ext) for ext in ext_data['requires']])
            section += f"|   **Requires**    | {requires}                 |\n"
            
        if ext_data.get('require_by'):
            require_by = ' '.join([self.format_ext(ext) for ext in ext_data['require_by']])
            section += f"|    **Need By**    | {require_by}                 |\n"
            
        if ext_data.get('see_also'):
            see_also = ' '.join([self.format_ext(ext) for ext in ext_data['see_also']])
            section += f"|   **See Also**    | {see_also} |\n"
            
        return section
    
    def generate_packages_table(self, ext_data: Dict[str, Any], matrix_data: List[Dict[str, Any]]) -> str:
        """Generate the packages overview table"""
        # Group by distro
        distros = {}
        for item in matrix_data:
            if item['type'] == 'rpm':
                distro = 'EL'
            else:
                distro = 'Debian'
                
            if distro not in distros:
                distros[distro] = {
                    'repo': item['pkg_repo'] or 'PIGSTY',
                    'version': item['pkg_ver'] or ext_data['version'],
                    'pg_versions': set(),
                    'pattern': item['pname'].replace(str(item['pg']), '$v') if item['pname'] else ''
                }
            if not item['miss']:
                distros[distro]['pg_versions'].add(item['pg'])
                
        # Build table
        table = "|   Distro   |                         Repo                         | Version |                                                                                                                         PG Major Availability                                                                                                                          | Package Pattern          | Dependencies |\n"
        table += "|:----------:|:----------------------------------------------------:|:-------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:-------------------------|:------------:|\n"
        
        for distro, info in distros.items():
            # Format PG versions with badges
            pg_badges = []
            for pg in sorted(self.config.pg_versions, reverse=True):
                if pg in info['pg_versions']:
                    if pg == 18:
                        alt = info["pattern"].replace("$v", str(pg))
                        pg_badges.append('{{< badge content="' + str(pg) + '" color="red" alt="' + alt + '" >}}')
                    else:
                        pg_badges.append('{{< badge content="' + str(pg) + '" color="green" >}}')
                        
            pg_str = ' '.join(pg_badges)
            
            link = "/e/" + (ext_data["lead_ext"] or ext_data["name"])
            repo_badge = '{{< badge content="' + info["repo"] + '" link="' + link + '" >}}'
            table += f"|   **{distro}**   |  {repo_badge}  | `{info['version']}` | {pg_str} | `{info['pattern']}` |      -       |\n"
            
        return table
    
    def generate_availability_table(self, matrix_data: List[Dict[str, Any]], availability_data: List[Dict[str, Any]]) -> str:
        """Generate the detailed availability matrix table"""
        # Organize data by OS and PG version
        availability_map = {}
        for item in availability_data:
            key = (item['os'], item['pg'])
            if key not in availability_map or item['org'] == 'PGDG':
                availability_map[key] = item
                
        # Build table header
        table = "| **Linux** / **PG** |"
        for pg in sorted(self.config.pg_versions, reverse=True):
            table += f"                  **PG{pg}**                   |"
        table += "\n|:------------------:|"
        for _ in self.config.pg_versions:
            table += ":-------------------------------------------:|"
        table += "\n"
        
        # Build rows for each OS
        for os_code, os_name in self.config.os_systems:
            table += f"|    `{os_code}`    |"
            
            for pg in sorted(self.config.pg_versions, reverse=True):
                # Get expected package name pattern for this OS/PG combination
                matrix_key = next((m for m in matrix_data if m['os'] == os_code and m['pg'] == pg), None)
                expected_pname = matrix_key['pname'] if matrix_key else None
                
                key = (os_code, pg)
                if key in availability_map:
                    item = availability_map[key]
                    if item['org'] == 'HIDE':
                        table += f'  {self.format_pkg(item["pname"], item["version"], "HIDE")}   |'
                    elif item['url']:
                        table += f' {self.format_pkg(item["pname"], item["version"], item["org"], item["url"])} |'
                    else:
                        table += f'      {self.format_pkg(item["pname"], item["version"], item["org"])}      |'
                else:
                    # If package doesn't exist but we know the expected name, show it
                    if expected_pname:
                        table += f'    {self.format_pkg(expected_pname)}     |'
                    else:
                        table += f'    {self.format_pkg("N/A")}     |'
            table += "\n"
            
        return table
    
    def generate_source_section(self, ext_data: Dict[str, Any]) -> str:
        """Generate the source section"""
        section = "{{< cards cols=3 >}}\n"
        
        if ext_data.get('url'):
            # Extract domain from URL
            if 'github.com' in ext_data['url']:
                icon = "github"
                subtitle = ext_data['url'].replace('https://', '')
            else:
                icon = "link"
                subtitle = ext_data['url'].replace('https://', '').replace('http://', '')
            section += '{{< card link="' + ext_data["url"] + '" title="Repository" icon="' + icon + '" subtitle="' + subtitle + '" >}}\n'
            
        if ext_data.get('source'):
            section += '{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="' + ext_data["source"] + '" >}}\n'
            
        section += "{{< /cards >}}\n\n"
        
        # Add build commands
        section += "\n```bash\n"
        section += f"pig build get {ext_data['name']}; # get {ext_data['name']} source code\n"
        section += f"pig build dep {ext_data['name']}: # install build dependencies\n"
        section += f"pig build pkg {ext_data['name']}; # build extension rpm or deb\n"
        section += f"pig build ext {ext_data['name']}; # build extension rpms\n"
        section += "```\n\n"
        
        return section
        
    def generate_install_commands(self, ext_data: Dict[str, Any]) -> str:
        """Generate install commands for different PG versions"""
        commands = ""
        for pg in sorted(ext_data.get('pg_ver', []), reverse=True):
            commands += f"pig ext install {ext_data['name']} -v {pg};   # install for PG {pg}\n"
        return commands
        
    def generate_comment_section(self, ext_data: Dict[str, Any]) -> str:
        """Generate comment section if exists"""
        if not ext_data.get('comment'):
            return ""
        return f"\n> [!Note] {ext_data['comment']}\n\n"
    
    def render_extension(self, ext_name: str) -> Optional[str]:
        """Render a single extension to Markdown"""
        # Get extension data
        ext_data = self.get_extension_data(ext_name)
        if not ext_data:
            logger.warning(f"Extension {ext_name} not found")
            return None
            
        # Get related data
        siblings = self.get_siblings(ext_data['pkg'])
        matrix_data = self.get_matrix_data(ext_data['pkg'])
        availability_data = self.get_availability_data(ext_data['pkg'])
        
        # Get category title
        category_title = ext_data['category'].title() if ext_data['category'] else 'Unknown'
        
        # Prepare template variables
        variables = {
            'id': ext_data['id'],
            'name': ext_data['name'],
            'pkg': ext_data['pkg'],
            'version': ext_data['version'] or '0.0.0',
            'category': ext_data['category'] or 'unknown',
            'category_title': category_title,
            'license': ext_data['license'] or 'unknown',
            'lang': ext_data['lang'] or 'unknown',
            'en_desc': ext_data['en_desc'] or ext_data['name'],
            'attributes_section': self.generate_attributes_section(ext_data),
            'relationships_section': self.generate_relationships_section(ext_data),
            'comment_section': self.generate_comment_section(ext_data),
            'packages_table': self.generate_packages_table(ext_data, matrix_data),
            'availability_table': self.generate_availability_table(matrix_data, availability_data),
            'source_section': self.generate_source_section(ext_data),
            'install_commands': self.generate_install_commands(ext_data),
            'create_schema': f" CASCADE SCHEMA {ext_data['schemas'][0]}" if ext_data.get('schemas') and not ext_data.get('relocatable') else ""
        }
        
        # Render template
        content = self.template
        for key, value in variables.items():
            content = content.replace(f'{{{key}}}', str(value))
            
        # Check if stub file exists and append its content
        stub_path = Path(__file__).parent.parent / 'stub' / f'{ext_name}.md'
        if stub_path.exists():
            try:
                with open(stub_path, 'r', encoding='utf-8') as f:
                    stub_content = f.read()
                if stub_content.strip():
                    content += '\n' + stub_content
                    logger.info(f"Appended stub content for {ext_name}")
            except Exception as e:
                logger.warning(f"Failed to read stub file for {ext_name}: {e}")
            
        return content
    
    def get_all_extensions(self) -> List[str]:
        """Get all extension names from database"""
        with self.conn.cursor() as cur:
            cur.execute("SELECT name FROM pgext.extension ORDER BY name")
            return [row[0] for row in cur.fetchall()]
            
    def generate_all(self):
        """Generate content for all extensions"""
        # Ensure output directory exists
        self.output_dir.mkdir(parents=True, exist_ok=True)
        
        # Get all extensions
        extensions = self.get_all_extensions()
        logger.info(f"Found {len(extensions)} extensions to generate")
        
        # Generate each extension
        success_count = 0
        for ext_name in extensions:
            try:
                content = self.render_extension(ext_name)
                if content:
                    output_file = self.output_dir / f"{ext_name}.md"
                    with open(output_file, 'w', encoding='utf-8') as f:
                        f.write(content)
                    success_count += 1
                    logger.debug(f"Generated: {output_file}")
                else:
                    logger.warning(f"No content generated for {ext_name}")
            except Exception as e:
                logger.error(f"Error generating {ext_name}: {e}", exc_info=True)
                
        logger.info(f"Successfully generated {success_count}/{len(extensions)} extension pages")
    


def main():
    """Main entry point"""
    parser = argparse.ArgumentParser(description='Generate Hugo content for PostgreSQL extensions')
    parser.add_argument('pgurl', nargs='?', default='postgres:///vonng',
                        help='PostgreSQL connection URL (default: postgres:///vonng)')
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
        generator.load_template()
        generator.generate_all()
    finally:
        generator.disconnect()


if __name__ == "__main__":
    main()