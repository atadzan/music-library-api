### Simple Music Library API written in GoLang
Service has the following REST API methods:
* Get library info with filter of all fields
* Get song lyrics with verse pagination
* Delete song
* Update song details
* Add new song 

### Technical task for this service
You should implement a REST API service which can do the following actions:
* Get library info with filter of all fields
* Get song lyrics with verse pagination
* Delete song
* Update song details
* Add new song 
```
{
   "group": "Muse",
   "song":  "Supermassive Black Hole"
}
```
```
Song details in JSON. 
Path: /info
Query params: 
            * group string
            * song  string   
Response:
200: SongInfo 
400: Bad Request
500: Internal server error
type SongInfo struct 
{
    ReleaseDate string `example:"16.07.2024"` 
    Text        string `example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"` 
    Link        string `example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"` 
}
```
1. Service should be provided with Swagger documentation
2. Prepare database schema migrations and apply them in PostgreSQL database after db connection established
3. Cover with debug and info logs
4. All configuration environments should be in .env file