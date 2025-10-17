#!/bin/bash

# Скрипт конвертации Markdown в PDF для ITSIMS
# Использование: ./convert_to_pdf.sh input.md [output.pdf]

INPUT_FILE="$1"
OUTPUT_FILE="${2:-${INPUT_FILE%.*}.pdf}"

if [ ! -f "$INPUT_FILE" ]; then
    echo "Ошибка: Файл $INPUT_FILE не найден"
    exit 1
fi

echo "Конвертация $INPUT_FILE в $OUTPUT_FILE..."

pandoc "$INPUT_FILE" -o "$OUTPUT_FILE" \
  --pdf-engine=xelatex \
  --variable mainfont="Arial" \
  --variable monofont="Courier New" \
  --variable fontsize=12pt \
  --variable geometry:margin=2cm \
  --variable lang=ru

if [ $? -eq 0 ]; then
    echo "✅ Конвертация завершена успешно: $OUTPUT_FILE"
else
    echo "❌ Ошибка при конвертации"
    exit 1
fi
