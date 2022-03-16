# ApiMarvel

Cette API permet d'obtenir les infos des heros marvel, on pourra obtenir leur Histoire, identité secrète, Personnalité, Pouvoir, première aparition...
Ce personnage est apparu dans plusieurs adaptations cintematographique, jeux videos ... 
zero ou plusieurs comédiens on interprté le rôle du personnage

#### Get all heroes

```http
  GET /heroes
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get heroe

```http
  GET /heroes/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

## Installation

```
  git clone https://github.com/RyanDervisevic/apiMarvel.git
  go run main.go
```

## Authors

- [@RyanDervisevic](https://github.com/RyanDervisevic)

