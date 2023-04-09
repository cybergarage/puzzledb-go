# Quick Start

This chapter shows you how to get started with PuzzleDB quickly: you can start a standalone PuzzleDB server with Docker and use Redis, MongoDB or MySQL CLI commands to insert and read sample data.

## Starting PuzzleDB Server

### Building from Source

To start the latest PuzzleDB, refer to [Go Stared](https://go.dev/learn/) to set up your Go development environment and run the following command:

```
git clone https://github.com/cybergarage/puzzledb-go.git
cd puzzledb
make run
```

### Using Docker image

PuzzleDB Docker image is the easiest way; if you do not have Docker installed, go there and install it first. To start the standalone server, run the following command:

```
docker run -it --rm \
 -p 6379:6379 \
 -p 27017:27017 \
 -p 3306:3306 \
 cybergarage/puzzledb
```

## Using database clients

The started PuzzleDB listens on the standard ports of the supported Redis, MongoDB, and MySQL database protocols, and you can connect with PuzzleDB using the standard CLI commands.

## Redis

To operate PuzzleDB with the Redis protocol, use the standard Redis command [redis-cli](https://redis.io/docs/ui/cli/) as follows:

```
% redis-cli 
127.0.0.1:6379> SET mykey "Hello"
OK
127.0.0.1:6379> GET mykey
"Hello"
```

PuzzleDB currently supports the Redis commands in stages. See [Redis](doc/redis.md) for current support status.

## MongoDB

To operate PuzzleDB with the MongoDB protocol, use the standard MongoDB shell [mongosh](https://www.mongodb.com/docs/mongodb-shell/#mongodb-binary-bin.mongosh) as follows:

```
% mongosh  
test> db.trainers.insertOne({name: "Ash", age: 10, city: "Pallet Town"})
test> db.trainers.findOne({name: "Ash"})
test> db.trainers.findOne({age: 10})
```

PuzzleDB currently supports the MongoDB commands in stages. See [MongoDB](doc/mongodb.md) for current support status.

## MySQL

To operate PuzzleDB with the MySQL protocol, use the standard MySQL shell [mysql](https://dev.mysql.com/doc/refman/8.0/en/mysql.html) as follows:

```
% mysql -h 127.0.0.1
mysql> CREATE DATABASE test;
mysql> USE test;
mysql> CREATE TABLE test (k VARCHAR(255) PRIMARY KEY, v int);
mysql> INSERT INTO test (k, v) VALUES ('foo', 0);
mysql> SELECT * FROM test WHERE k = 'foo';
+------+------+
| k  | v  |
+------+------+
| foo |  0 |
+------+------+
1 row in set (0.00 sec)
```

PuzzleDB currently supports the MySQL commands in stages. See [MySQL](doc/mysql.md) for current support status.

