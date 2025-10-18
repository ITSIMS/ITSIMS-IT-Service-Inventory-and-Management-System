#!/bin/bash

# Скрипт для конвертации всех Markdown файлов в PDF с сохранением структуры проекта
# Исключает папку ТПКС согласно .gitignore и папку печати

PROJECT_ROOT="/Users/gregorykogan/Desktop/ITSIMS-IT-Service-Inventory-and-Management-System"
OUTPUT_DIR="$PROJECT_ROOT/печать"

# Функция для создания директории
create_directory() {
    local dir_path="$1"
    if [ ! -d "$dir_path" ]; then
        mkdir -p "$dir_path"
        echo "📁 Создана директория: $dir_path"
    fi
}

# Функция для конвертации markdown файла
convert_markdown_to_pdf() {
    local input_file="$1"
    local output_file="$2"
    
    echo "🔄 Конвертация $input_file в $output_file..."
    
    pandoc "$input_file" -o "$output_file" \
        --pdf-engine=xelatex \
        --variable mainfont="Arial" \
        --variable monofont="Courier New" \
        --variable fontsize=12pt \
        --variable geometry:margin=2cm \
        --variable lang=ru
    
    if [ $? -eq 0 ]; then
        echo "✅ Конвертация завершена успешно: $output_file"
    else
        echo "❌ Ошибка при конвертации $input_file"
        return 1
    fi
}

# Функция для копирования PDF файла
copy_pdf_file() {
    local source_file="$1"
    local dest_file="$2"
    
    echo "📋 Копирование $source_file в $dest_file..."
    
    if cp "$source_file" "$dest_file"; then
        echo "✅ Копирование завершено успешно: $dest_file"
    else
        echo "❌ Ошибка при копировании $source_file"
        return 1
    fi
}

# Функция для обработки файлов и директорий
process_directory() {
    local source_dir="$1"
    local dest_dir="$2"
    
    # Создаем целевую директорию
    create_directory "$dest_dir"
    
    # Обрабатываем все элементы в директории
    for item in "$source_dir"/*; do
        # Получаем имя файла/папки
        local item_name=$(basename "$item")
        
        # Пропускаем скрытые файлы и системные файлы
        if [[ "$item_name" == .* ]] || [[ "$item_name" == "DS_Store" ]]; then
            continue
        fi
        
        # Пропускаем папку ТПКС согласно .gitignore
        if [[ "$item_name" == "ТПКС" ]]; then
            echo "⏭️  Пропускаем папку ТПКС (исключена в .gitignore)"
            continue
        fi
        
        # Пропускаем папку печати, чтобы избежать рекурсии
        if [[ "$item_name" == "печать" ]]; then
            echo "⏭️  Пропускаем папку печати (избегаем рекурсии)"
            continue
        fi
        
        # Если это директория, рекурсивно обрабатываем
        if [ -d "$item" ]; then
            echo "📂 Обработка директории: $item_name"
            process_directory "$item" "$dest_dir/$item_name"
        else
            # Если это markdown файл, конвертируем в PDF
            if [[ "$item" == *.md ]]; then
                local pdf_name="${item_name%.md}.pdf"
                local pdf_path="$dest_dir/$pdf_name"
                convert_markdown_to_pdf "$item" "$pdf_path"
            # Если это уже PDF файл, копируем его
            elif [[ "$item" == *.pdf ]]; then
                local pdf_path="$dest_dir/$item_name"
                copy_pdf_file "$item" "$pdf_path"
            else
                echo "⏭️  Пропускаем файл: $item_name (не markdown и не PDF)"
            fi
        fi
    done
}

# Основная функция
main() {
    echo "🚀 Начинаем конвертацию всех Markdown файлов в PDF..."
    echo "📁 Исходная директория: $PROJECT_ROOT"
    echo "📁 Целевая директория: $OUTPUT_DIR"
    echo ""
    
    # Создаем корневую директорию для печати
    create_directory "$OUTPUT_DIR"
    
    # Обрабатываем все файлы и директории
    process_directory "$PROJECT_ROOT" "$OUTPUT_DIR"
    
    echo ""
    echo "🎉 Конвертация завершена!"
    echo "📁 Все PDF файлы сохранены в: $OUTPUT_DIR"
}

# Запускаем основную функцию
main
