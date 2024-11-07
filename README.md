<h1 align="center"> go gin banks is an rest api banking written in go with gin </h1>
<h3 align="center"> requirement for the app </h3>
first you need to have go installed on your device to check that type

```bash
$ go version
```

if you don't have it installed check the documentation of golang
secondly you need to have postgreSQL.

or you can just use the docker by doing

```bash
$ docker compose build
$ docker compose up
```

<h3 align="center" > about the app </h3>
as i've said above this app is an bank rest api , this means it's include all the functionnality of an typic banking app :
- depot , retreat money
- sending money
- chat in realtime 
- an admin for managing 
- an epargne (in building)
and many more ...

<h3 align="center"> Routes and body </h3>
<h4> for the user </h4>

- Register user method post :

```bash
http://localhost:<your port>/api/v1/user/register
```

body content :

```json
{
  "AppUserName": "appusername",
  "Name": "name ",
  "FirstName": "firsname",
  "Residance": "place where you liv",
  "Email": "mail@gmail.com",
  "Password": "your password",
  "Date_de_naissance": "when you are born",
  "BirthDate": "date when you born"
}
```

- Login user method post :

```bash
http://localhost:<your port>/api/v1/user/login
```

```json
{
  "email": "your mail",
  "password": "your password"
}
```

<h5 style="color : red;" > notes these routes use bearer token authorization , token is the  token given by the register or login </h5>

- Get all users

```bash
http://localhost:<your port>/api/v1/user/
```

- Get single user 

```bash
http://localhost:<your port>/api/v1/user/
```

- Get connected user 

```bash
http://localhost:<your port>/api/v1/user/logedUser
```

- Upload profile picture user : method post

```bash
http://localhost:<your port>/api/v1/user/pp
```
you need an multipart-form data as your content-type then "filePP" is the name given for the file.

- Update user info : method patch

```bash
http://localhost:<your port>/api/v1/user/
```
<br> the avalaible body update yall is type of string<br>

```json
{
  "AppUserName":"your new app user name",
  "residance":"your new app residance"
}
```
- Update user : method patch 

```bash
http://localhost:<your port>/api/v1/user/
```

<br> the avalaible body setting rmEpargne and DeleteMyAccount are type bool , and block and unblock are type of string the uuid of the user you want to block<br>

```json
{
  "rmEpargne":bool,
  "rmAccount":bool,
  "blockAcc":"user you want to block uuid",
  "unblockAcc":"user you want to unblock"
}
```

