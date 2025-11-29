#!/bin/bash

# Скрипт подготовки документа к печати для ITSIMS
# Конвертирует Markdown файл в PDF с настройками для печати
# Использование: ./prepare_for_print.sh input.md [output.pdf]

INPUT_FILE="$1"

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
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
echo "🖨️  Подготовка к печати: $INPUT_FILE -> $OUTPUT_FILE..."

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
    echo "✅ Документ подготовлен к печати: $OUTPUT_FILE"
else
    echo "❌ Ошибка при подготовке документа к печати"
    exit 1
fi

