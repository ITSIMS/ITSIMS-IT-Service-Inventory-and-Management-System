#!/bin/bash

# Скрипт конвертации Markdown в PDF для ITSIMS
# Использование: ./convert_to_pdf.sh input.md [output.pdf]

INPUT_FILE="$1"

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
PRINT_DIR="$PROJECT_ROOT/print"

# Если выходной файл не указан, сохраняем PDF в корневую папку print
if [ -z "$2" ]; then
    mkdir -p "$PRINT_DIR"
    OUTPUT_FILE="$PRINT_DIR/$(basename "${INPUT_FILE%.*}").pdf"
else
    OUTPUT_FILE="$2"
fi

if [ ! -f "$INPUT_FILE" ]; then
    echo "Ошибка: Файл $INPUT_FILE не найден"
    exit 1
fi
echo "Конвертация $INPUT_FILE в $OUTPUT_FILE..."

pandoc "$INPUT_FILE" -o "$OUTPUT_FILE" \
  --pdf-engine=xelatex \
  -H "$SCRIPT_DIR/resources/listings-setup.tex" \
  --lua-filter="$SCRIPT_DIR/resources/default_table_width.lua" \
  --wrap=auto \
  --variable mainfont="Arial" \
  --variable monofont="Courier New" \
  --variable fontsize=14pt \
  --variable geometry:margin=2cm \
  --variable lang=ru

if [ $? -eq 0 ]; then
    echo "✅ Конвертация завершена успешно: $OUTPUT_FILE"
else
    echo "❌ Ошибка при конвертации"
    exit 1
fi
