# 01 Filename Ext

## Требуется написать программу, которая по введенному полному пути выводит имя файла и расширение

## Пример
```shell
./filename-ext /home/robotomize/1.txt
```
```shell
filename: 1
extension: txt
```

## Сборка
```shell
go build -o filename-ext
```

## Сценарии

```shell
./filename-ext /home/robotomize/1.txt

filename: 1
extension: txt
```

```shell
./filename-ext /home/robotomize/1.txt.txt

filename: 1.txt
extension: txt
```

```shell
./filename-ext /home/robotomize/1

filename: 1
extension:
```