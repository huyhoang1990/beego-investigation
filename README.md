Change user, pass, dbname in infras/mysql.go

Create a table users

```
CREATE TABLE `users` (
  `id` varchar(40) NOT NULL,
  `username` varchar(30) NOT NULL,
  `password` varchar(155) NOT NULL,
  `created_at` bigint NOT NULL,
  `expired_at` bigint NOT NULL,
  `status` varchar(15) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `users_username_idx` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

Create an account

```
go test -v ./service -run TestAddNewUser
```

Start service with watchmedo to autoreload -> if not just need to run go run ./main.go
```
watchmedo shell-command --patterns="*.html;*.go" --recursive --command="lsof -ti tcp:8080 | xargs kill -9;go run ./main.go" .
```

Reproduce bug
```
1: open chrome access to http://127.0.0.1:8080/login
2: login with account you created above
3: after logging in -> it will be redirected to http://127.0.0.1:8080/home
4: Open an incognito browser of chrome then access to http://127.0.0.1:8080/login again -> it will redirect to home immediately instead of login page :(((( -> bug here
```