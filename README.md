<h1 align="center"> go gin banks is an rest api banking written in go with gin </h1>
<h3 align="center"> requirement for the app </h3>
first you need to have go installed on your device to check that type 

``` bash
$ go version
```
if you don't have it installed check the documentation of golang
secondly you need to have postgreSQL.

or you can just use the docker by doing 
``` bash
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
```
http://localhost:<your port>/api/v1/user/register
```
body content 
``` json
  {
  "AppUserName":"appusername",
  "Name":"name ",
  "FirstName":"firsname",
  "Residance":"place where you liv",
  "Email":"mail@gmail.com",
  "Password":"your password",
  "Date_de_naissance":"when you are born",
  "BirthDate":"date when you born"
}
```
