# Golang CRUD with PostgreSQL

## Table of contents 👀
* [General info](#general-info)
* [Technologies](#technologies)
* [Blog](#blog)
* [Setup](#setup)


### General info
GOPOST or Go-Post is a Golang REST API made to show some indetity for person. 

#### The GOPOST Object 🍵
| Properties | Description | Type  |
|:----------- |:---------------|:--------|
|name| the identity name | String| 
|location| location from the identity | String |
|age| the age of identity | Int64 | 


#### Routes ⚡
| Routes | HTTP Methods| Description
|:------- |:---------------|:--------------
| /api/user/     | GET                  | Displays all identity
| /api/newuser/      | POST               | Creates a new identity
| /api/deleteuser/      | DELETE            | Deletes all tea
|/api/user/{id}| GET     | Displays a specific tea, given its name
|/api/user/{id}| PUT  | Update identitiy Value
|/api/user/{id}}| DELETE | Deletes a specific identity, given its id
	
### Technologies
Project is created with:

* Golang 
* gorilla/mux 
* lib/pq  
* joho/godotenv 
* ElephantSQL



### How I built it
👉 [Check out the series here!](https://berkaryasemampunya.medium.com/build-crud-apps-in-golang-with-postgresql-3be08d31a1f1)


### Setup
To run this project locally, clone repo and add an `.env` file in the root:
```
POSTGRES_URL="Postgres connection string"
```

Then execute in command prompt:
```
$ cd go-post
$ go mod tidy
$ go run main.go
```

## API Reference

These are the endpoints available from the app

### `GET /api/user/`

Returns result identity

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "data": [
    {
     "id":40,
     "name":"hafizh",
     "location": "widodo",
     "age":22,
    }
  ]
}
```

</p>
</details>

---


### `POST /api/newuser/`

Creates a new identity

#### Request 

This request requires body payload, you can find the example below.

<details><summary>Show example payload</summary>
<p>

```json
{
 "name":"hafizh",
 "location":"kansas",
 "age":12
}
```
</p>
</details>


### `GET /players/:id`

Returns a player by id

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": {
    "id": "5f6a5c31d7c451c369802c02",
    "name": "John Doe 1",
    "nickname": "Lolo",
    "position": "forward",
    "created_at": "2020-09-22T20:18:57.957Z"
  }
}
```

</p>
</details>


### `GET /players/:id`

Returns a player by id

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": {
    "id": "5f6a5c31d7c451c369802c02",
    "name": "John Doe 1",
    "nickname": "Lolo",
    "position": "forward",
    "created_at": "2020-09-22T20:18:57.957Z"
  }
}
```

</p>
</details>

---

### `UPDATE /api/user/:ID`

Update value of identity
	
#### Response

<details><summary>Show example response</summary>
<p>

```json
{
    "name": "hafizah",
    "location": "toronto",
    "age": 25,  
}
```

</p>
</details>


---

### `DELETE /api/user/`

Delete all team
	
### `DELETE /api/user/:ID`

Delete team by id