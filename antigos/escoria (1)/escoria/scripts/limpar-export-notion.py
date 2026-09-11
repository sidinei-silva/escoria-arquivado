#!/usr/bin/env python3
"""
Limpa uma exportação Markdown do Notion.

O export do Notion produz nomes como:
    Tom & Estilo 3840b63b76f281a1b3ecd19fc752beef.md
    30 · Mecânicas 3840b63b76f281bba0b5cb76b9efc5b1/

e links internos apontando para esses nomes, URL-encoded.

Este script:
  - remove o sufixo UUID de arquivos e pastas
  - transforma os nomes em slugs sem acento (a-z, 0-9, hífen)
  - reescreve os links internos para os novos caminhos
  - opcionalmente converte os links para [[wikilinks]] do Obsidian

Uso:
    python3 limpar-export-notion.py ORIGEM DESTINO [--wikilinks] [--dry-run]

Exemplo:
    unzip ~/Downloads/Export-abc123.zip -d /tmp/gdd-export
    python3 scripts/limpar-export-notion.py /tmp/gdd-export gdd/
"""

import argparse
import re
import shutil
import sys
import unicodedata
from pathlib import Path
from urllib.parse import unquote

UUID_RE = re.compile(r"[ _-]?[0-9a-f]{32}$", re.IGNORECASE)
LINK_RE = re.compile(r"\[([^\]]*)\]\(([^)]+)\)")


def slug(texto: str) -> str:
    """Nome de arquivo previsível: sem acento, minúsculo, hifenizado."""
    texto = UUID_RE.sub("", texto).strip()
    texto = unicodedata.normalize("NFKD", texto)
    texto = texto.encode("ascii", "ignore").decode("ascii")
    texto = texto.replace("·", "-").replace("&", "e")
    texto = re.sub(r"[^\w\s-]", "", texto)
    texto = re.sub(r"[\s_]+", "-", texto).strip("-").lower()
    texto = re.sub(r"-{2,}", "-", texto)
    return texto or "sem-nome"


def planejar(origem: Path, destino: Path):
    """Monta o mapa caminho-antigo -> caminho-novo, preservando a hierarquia."""
    mapa = {}
    for item in sorted(origem.rglob("*")):
        rel = item.relative_to(origem)
        partes = []
        for i, parte in enumerate(rel.parts):
            ultimo = i == len(rel.parts) - 1
            if ultimo and item.is_file():
                base = Path(parte).stem
                ext = Path(parte).suffix or ".md"
                partes.append(slug(base) + ext)
            else:
                partes.append(slug(parte))
        mapa[rel] = Path(*partes)
    return mapa


def indice_de_links(mapa):
    """
    Índice para resolver links: o Notion linka pelo nome de arquivo original
    (URL-encoded). Mapeia tanto o nome cru quanto o decodificado.
    """
    idx = {}
    for antigo, novo in mapa.items():
        idx[antigo.name] = novo
        idx[unquote(antigo.name)] = novo
        idx[str(antigo)] = novo
        idx[unquote(str(antigo))] = novo
    return idx


def reescrever_links(texto: str, idx, novo_rel: Path, wikilinks: bool) -> str:
    def troca(m):
        rotulo, alvo = m.group(1), m.group(2)

        if alvo.startswith(("http://", "https://", "#", "mailto:")):
            return m.group(0)

        alvo_limpo = unquote(alvo.split("#")[0])
        destino = idx.get(alvo_limpo) or idx.get(Path(alvo_limpo).name)
        if destino is None:
            return m.group(0)

        if wikilinks:
            return f"[[{destino.stem}|{rotulo}]]" if rotulo else f"[[{destino.stem}]]"

        try:
            rel = Path(*([".."] * (len(novo_rel.parts) - 1))) / destino
        except ValueError:
            rel = destino
        return f"[{rotulo}]({rel.as_posix()})"

    return LINK_RE.sub(troca, texto)


def main():
    ap = argparse.ArgumentParser(description="Limpa export Markdown do Notion.")
    ap.add_argument("origem", type=Path, help="pasta do export descompactado")
    ap.add_argument("destino", type=Path, help="pasta de destino (ex.: gdd/)")
    ap.add_argument("--wikilinks", action="store_true",
                    help="converter links para [[wikilinks]] do Obsidian")
    ap.add_argument("--dry-run", action="store_true",
                    help="mostrar o que faria, sem escrever")
    args = ap.parse_args()

    if not args.origem.is_dir():
        sys.exit(f"erro: origem não encontrada: {args.origem}")

    mapa = planejar(args.origem, args.destino)
    idx = indice_de_links(mapa)

    arquivos = md = 0
    for antigo_rel, novo_rel in mapa.items():
        antigo = args.origem / antigo_rel
        novo = args.destino / novo_rel

        if antigo.is_dir():
            if not args.dry_run:
                novo.mkdir(parents=True, exist_ok=True)
            continue

        arquivos += 1
        if args.dry_run:
            print(f"  {antigo_rel}  ->  {novo_rel}")
            continue

        novo.parent.mkdir(parents=True, exist_ok=True)

        if antigo.suffix.lower() == ".md":
            texto = antigo.read_text(encoding="utf-8", errors="replace")
            texto = reescrever_links(texto, idx, novo_rel, args.wikilinks)
            # o Notion repete o título como H1 na primeira linha; mantém, mas limpa o UUID
            texto = UUID_RE.sub("", texto.split("\n")[0]) + "\n" + "\n".join(texto.split("\n")[1:])
            novo.write_text(texto, encoding="utf-8")
            md += 1
        else:
            shutil.copy2(antigo, novo)

    if args.dry_run:
        print(f"\n[dry-run] {arquivos} arquivos seriam processados")
    else:
        print(f"pronto: {arquivos} arquivos ({md} markdown) em {args.destino}")
        print("\nConfira o drift restante:")
        print(f'  grep -rin "destiny board\\|adventurer\\|woodcutter\\|o farol\\|a cova" {args.destino}')


if __name__ == "__main__":
    main()
