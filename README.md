# Рейтинг компаний в moikrug.ru

## Шаги по запуску сервисов:

1. Запуск сервиса parser

```
make parser
```


2. Запуск сервиса calculator

```
make calculator
```


3. Запуск сервиса web

```
make web
```


4. Перейти по адресу http://localhost:7000/


![rating result](https://imgur.com/V69iy7J.png)



Как вычисляется рейтинг:
-----------------------

- За пришедшего сотрудника прибавляется 1 бал
- За ушедшего сотрудника вычитается 1 бал
- За указанный сайт в профиле компании прибавляется 3 балла
- За указанный адрес в профиле компании прибавляется 3 балла
- Дополнительно прибавляется moikrug рейтинг, указанный в профиле компании 

