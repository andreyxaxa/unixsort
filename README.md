# Утилита sort

Программа читает строки (из файла или STDIN) и выводит их отсортированными.

Флаги (как в GNU sort):

    -k N — сортировать по столбцу (колонке) (разделитель — табуляция по умолчанию).
    Например, «sort -k 2» отсортирует строки по второму столбцу каждой строки.

    -n — сортировать по числовому значению (строки интерпретируются как числа).

    -r — сортировать в обратном порядке (reverse).

    -u — не выводить повторяющиеся строки (только уникальные).

    -M — сортировать по названию месяца (Jan, Feb, ... Dec), т.е. распознавать специфический формат дат.

    -b — игнорировать хвостовые пробелы (trailing blanks).

    -h — сортировать по числовому значению с учётом суффиксов (например, К = килобайт, М = мегабайт — человекочитаемые размеры).

Программа корректно обрабатывает комбинации флагов (например, -nr — числовая сортировка в обратном порядке, и т.д.).

Тесты - `go test -v ./...`

Примеры запусков: (из корневой директории "unixsort")
```
go run cmd/app/main.go pkg/unixsort/testdata/basicinput.txt > pkg/unixsort/testdata/basicexpected.txt
```

```
go run cmd/app/main.go -n pkg/unixsort/testdata/numericinput.txt > pkg/unixsort/testdata/numericexpected.txt
```

```
go run cmd/app/main.go -nruk 2 pkg/unixsort/testdata/nrukinput.txt > pkg/unixsort/testdata/nrukexpected.txt
```
