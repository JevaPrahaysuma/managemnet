## Documentation
# Application For Get Menu In API NYPL

1. Request And Response API

- **Method** : GET

- **Header** : ```http://localhost:8080/user/```

- **Params** 
| Property | Mandatory | Description |
| ----------- | ----------- | ----------- |
| q | No | Using to search name dish and name menu |
| vanues |  No | Using to search vanues |
| events | No | Using to search events |
| startPrice | No | Using to implement start price |
| endPrice | No | Using to implement end price |

- **Response**

```
{
    "status": 200,
    "message": "success",
    "data": {
        "data": [
            {
                "menu_item_id": 976059,
                "name_dish": "Mug of Famous U. S. Grant Hotel Coffee",
                "price": 0.05,
                "sponsor": "U.S. Grant Hotel",
                "physical_description": "22.5x14cm",
                "call_number": "1921-0045_wotm",
                "date": "1921-05-17",
                "location": "U.S. Grant Hotel",
                "currency": "Dollars",
                "currency_symbol": "$",
                "status": "complete",
                "created_at": "2012-06-20T21:08:00Z",
                "updated_at": "2012-06-20T21:08:00Z"
            }
        ]
    }
}
```

2. Specifications Used

- **programming language** : golang 1.19
- **Database** : Postgres
- **Container** : Docker
- **Library** :
    - gorm
    - gin gonic

3. Running Application

- Step 1:
```
docker compose up --build
```
- Step 2:
```
go run main.go
```

# Design Architecture

![design architecture](https://github.com/JevaPrahaysuma/management/src/desaign_pattern.jpg)

# Link Collection API
```
https://www.getpostman.com/collections/e744cf537ab197bed852
```

