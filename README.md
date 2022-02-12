# AUTH Weakly Menu API

The purpose of this API is to weekly parse&cache the AUTH's menu from this
[webpage](https://www.auth.gr/weekly-menu/) and have endpoints available to the
public so users can create **Google Assistant and Siri shortcuts** or **Bots
for apps like Discord**.

| Endpoint                 | Method | Description                                                                |
| ------------------------ | :----: | -------------------------------------------------------------------------- |
| `/api/menu`              | `GET`  | Get menu for the whole week                                                |
| `/api/menu/today`        | `GET`  | Get menu for today                                                         |
| `/api/menu/weekday/:day` | `GET`  | Get menu for specific weekday `day` is an `int`, _1->Monday ... 7->Sunday_ |
| `/api/menu/fetch`        | `GET`  | Try to fetch menu from [webpage](https://www.auth.gr/weekly-menu/)         |
