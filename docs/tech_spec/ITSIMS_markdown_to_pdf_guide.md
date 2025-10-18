# Руководство по конвертации Markdown в PDF

**Проект:** ITSIMS  
**Версия:** 1.1  
**Дата:** 18 октября 2025  
**Автор:** Gregory K  
**Ответственный:** Gregory K

---

## Лист изменений

| Версия | Дата | Автор | Описание изменений | Основание |
|--------|------|-------|--------------------|-----------|
| 1.0 | 15.10.2025 | Gregory K | Первая версия документа | [Ссылка на таск](https://www.notion.so/Markdown-PDF-28dcf70d42c9800f8b0ff1c56f59f24e?source=copy_link) |
| 1.1 | 18.10.2025 | Gregory K | Добавлен скрипт для конвертации всех файлов проекта | Обеспечение удобства работы с большим проектом |

**Подпись таблицы:** Таблица 1. Лист изменений.

---

## 1. Установка необходимых инструментов

### 1.1 Pandoc

```bash
# macOS (via Homebrew)
brew install pandoc

# Ubuntu/Debian
sudo apt-get install pandoc

# Windows (via Chocolatey)
choco install pandoc
```

### 1.2 LaTeX (для PDF-движка)

```bash
# macOS
brew install --cask mactex

# Ubuntu/Debian
sudo apt-get install texlive-full

# Windows
# Скачать с https://miktex.org/
```

---

## 2. Базовые команды конвертации

### 2.1 Простая конвертация

```bash
pandoc input.md -o output.pdf
```

### 2.2 Конвертация с настройками для ITSIMS

```bash
pandoc ITSIMS_spec_technical.md -o ITSIMS_spec_technical.pdf \
  --pdf-engine=xelatex \
  --variable mainfont="Arial" \
  --variable monofont="Courier New" \
  --variable fontsize=12pt \
  --variable geometry:margin=2cm \
  --variable lang=ru
```

---

## 3. Скрипты автоматической конвертации

### 3.1 Конвертация одного файла

```bash
# Auto naming
./convert_to_pdf.sh ITSIMS_spec_technical.md

# Manual naming
./convert_to_pdf.sh ITSIMS_spec_technical.md technical_spec.pdf
```

### 3.2 Конвертация всех файлов проекта

```bash
./convert_all_to_pdf.sh
```

**Описание скрипта `convert_all_to_pdf.sh`:**

- Автоматически находит все `.md` файлы в проекте
- Конвертирует их в PDF с настройками ITSIMS
- Сохраняет структуру папок в директории `/печать`
- Копирует существующие PDF файлы
- Избегает рекурсии с папкой печати

---

## 4. Контакты для экстренных случаев

- **Антон** (@Anton Zatsepilin) - основной полиграфист
- **Илья** (@Ilya Grechko) - резервный полиграфист
- **Gregory K** - техническая поддержка по конвертации
