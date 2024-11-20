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
  "AppUserName": "your new app user name",
  "residance": "your new app residance"
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

<h4> for the admin </h4>
- Register as an admin method post

```bash
http://localhost:<your port>/api/v1/admin/register
```

```json
{
  "name": "yourUniqueAdminName",
  "passwords": "yourPasswords",
  "RootPass": "mandresy"
}
```

<h5 style="color: red">notice:  the actual root pass is static with the value of "mandresy" but some future update will be make it configurable </h5>

- Login as an admin method post

```bash
http://localhost:<your port>/api/v1/admin/login
```

```json
{
  "name": "yourUniqueAdminName",
  "passwords": "yourPasswords"
}
```

- Creating a bank

##### notes a bank is just a place for doing some retrait and depot

```bash
http://localhost:<your port>/api/v1/admin/createBank
```

```json
{
  "place": "place of the bank ou want",
  "passwords": "yourPasswords",
  "money": "the money you want to allocate to this bank"
}
```

##### these methodes above needs an admin token

```bash
http://localhost:<your port>/api/v1/admin/getBank
```

- Get admin info

```bash
http://localhost:<your port>/api/v1/admin/getAdminInfo
```

### for transactions (money)

##### notes retrait needs an place that should be valid but in our test case money is virtual and place too

- Depot methods put

```bash
http://localhost:<your port>/api/v1/transaction/depot
```

```json
{
  "value": value of the money you want to depose,
  "place":"place where you do an depot",
  "password":"your password"
}
```

- retrait money methods put

```bash
http://localhost:<your port>/api/v1/transaction/retrait
```

- send money methods post

```bash
# uuid user uuid
http://localhost:<your port>/api/v1/transaction/:uuid
```

```json
{
  "value": value of the money you want to send,
  "password":"your passwords"
}
```

- get all historics

```bash
http://localhost:<your port>/api/v1/transaction/hitstoric
```

- get single historic

```bash
# uuid is transctions uuid
http://localhost:<your port>/api/v1/transaction/:uuid
```

#### epargne functionnality

##### notes all epargnes functionnality needs an authentification

- create epargne methodes post

```json
{
  "name": "names",
  "type": "type",
  "message": "all your messages",
  // sent_to need to be empty if your epargne is an economie
  "sent_to": "sent to uuid",
  "value": "amount of money you want to epargnes need to be inferior to the userConnected's money ",
  "date": "day of the epargne bettween 1 -> 31",
  "auto_send": "boolean if you want to auto_send it ",
  "is_economie": "boolean if the epargne is an economie"
}
```

```bash
http://localhost:<your port>/api/v1/epargne/createEpargne
```

- get all my epargnes

```bash
http://localhost:<your port>/api/v1/epargne/
```

- get single epargne

```bash
http://localhost:<your port>/api/v1/epargne/:epargneUUID
```

- delete an epargnes

##### notes : if an epargnes is deleted it won't be auto epargned anymore

```bash
http://localhost:<your port>/api/v1/epargne/:epargneUUID
```
