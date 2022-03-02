# GO CRUD application

## Made by:

| Members  |
| ------------- | 
| Henry-Gerret Grüning  | 
| Anette Aguraiuja  | 

## Made with:

- GoLang
- MySql

## Setting up the application

- MySql Server using the Mysql Installation Guide [Mysql Installation Guide](https://dev.mysql.com/doc/mysql-getting-started/en/#mysql-getting-started-installing)

``` 
mkdir crud
``` 
``` 
cd crud
``` 
``` 
git clone https://github.com/AnetteAgura/CRUDGo
``` 
``` 
sudo apt-get update
``` 
``` 
sudo apt-get install mysql-server
``` 
``` 
cd CRUDGo
``` 
``` 
mysql -h localhost -u root < RUN.sql
``` 
``` 
sudo apt install golang-go
``` 

## How to run:

```
go run main.go
```
## How to run:
Create a .env file and add the following into it
```
DBDRIVER=mysql
DBUSER= [YOUR DATABASE USERNAME]
DBPASS= [YOUR DATABASE PASSWORD]
DBNAME=games
```
## Possibilities:

- You can see all games in database
- You can add new games to the database
- You can delete or update the game
- You can see game detailed view
  - In detailed view you can click on the company to see other games made by same company
  - In detailed view you can click on the year to see other games made on same year

## PS!

**Shows ERRORS but they do not matter**
