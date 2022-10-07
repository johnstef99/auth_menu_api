# AUTH Weakly Menu API

The purpose of this API is to weekly parse&cache the AUTH's menu from this
[webpage](https://www.auth.gr/weekly-menu/) and have endpoints available to the
public so users can create **Google Assistant and Siri shortcuts** or **Bots
for apps like Discord**.

| Endpoint                 | Method | Description                                                                | Params Available                  |
| ------------------------ | :----: | -------------------------------------------------------------------------- | --------------------------------- |
| `/api/menu`              | `GET`  | Get menu for the whole week                                                |
| `/api/menu/today`        | `GET`  | Get menu for today                                                         | `description`[^param:description] |
| `/api/menu/weekday/:day` | `GET`  | Get menu for specific weekday `day` is an `int`, _1->Monday ... 7->Sunday_ | `description`[^param:description] |
| `/api/menu/fetch`        | `GET`  | Try to fetch menu from [webpage](https://www.auth.gr/weekly-menu/)         |

[^param:description]:
    If `description` is `true` the api will return just one string for lunch and one for dinner

    ```json
    {
      "date": "2022-02-18T00:00:00Z",
      "dinner": "- Ριζότο μανιταριών\n- Αγγουροντομάτα\n  Κουνουπίδι ατμού\n- Θρ..",
      "lunch": "- Κολοκύθια τηγανητά με τζατζίκι\n- Αγγουροντομάτα\n  Κουνουπίδι.."
    }
    ```

Take a look at the [example](./example.json) to see the structure of the response

An instance of this API is hosted at [lesxi.johnstef.com](http://lesxi.johnstef.com/api/menu)

Siri shortcut: [download](https://www.icloud.com/shortcuts/c9a7fbd873504f8596893332b239c93a) (You need to enable untrusted shortcuts from settings)
