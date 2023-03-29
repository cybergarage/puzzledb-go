![](img/logo.png)

# Redis

PuzzleDB supports MongoDB API based on [go-redis](https://github.com/cybergarage/go-redis), a database framework that makes it easy to implement Redis compatible servers using Go.

## Supported commands

PuzzleDB supports the following supported Redis commands.

### System commands

|Supported|Command|Redis Version|Note                             |
|---------|-------|-------------|---------------------------------|
|O        |ECHO   |1.0.0        |                                 |
|O        |PING   |1.0.0        |                                 |
|O        |QUIT   |1.0.0        |                                 |
|O        |SELECT |1.0.0        |                                 |

### Generic commands

|Supported|Command      |Redis Version|Note                             |
|---------|------------------|-------------|---------------------------------|
|-        |COPY              |6.2.0        |                                 |
|O        |DEL               |1.0.0        |                                 |
|-        |DUMP              |2.6.0        |                                 |
|O        |EXISTS            |1.0.0        |                                 |
|O        |EXPIRE            |1.0.0        |                                 |
|O        |EXPIREAT          |1.2.0        |                                 |
|-        |EXPIRETIME        |7.0.0        |                                 |
|O        |KEYS              |1.0.0        |                                 |
|-        |MOVE              |1.0.0        |                                 |
|-        |MIGRATE           |2.6.0        |                                 |
|-        |RANDOMKEY         |1.0.0        |                                 |
|O        |RENAME            |1.0.0        |                                 |
|O        |RENAMENX          |1.0.0        |                                 |
|-        |RESTORE           |2.8.0        |                                 |
|-        |SCAN              |2.8.0        |                                 |
|-        |SORT              |1.0.0        |                                 |
|-        |SORT_RO           |7.0.0        |                                 |
|-        |TOUCH             |3.2.1        |                                 |
|O        |TTL               |1.0.0        |                                 |
|O        |TYPE              |1.0.0        |                                 |
|-        |UNLINK            |4.0.0        |                                 |
|-        |WAIT              |3.0.0        |                                 |

### String commands

|Supported|Command      |Redis Version|Note                             |
|---------|------------------|-------------|---------------------------------|
|O        |APPEND            |2.0.0        |                                 |
|O        |DECR              |1.0.0        |                                 |
|O        |DECRBY            |1.0.0        |                                 |
|O        |GET               |1.0.0        |                                 |
|-        |GETDEL            |6.2.0        |                                 |
|-        |GETEX             |6.2.0        |                                 |
|O        |GETRANGE          |2.4.0        |                                 |
|O        |GETSET            |1.0.0        |                                 |
|O        |INCR              |1.0.0        |                                 |
|O        |INCRBY            |1.0.0        |                                 |
|-        |INCRBYFLOAT       |2.6.0        |                                 |
|-        |LCS               |7.0.0        |                                 |
|O        |MGET              |1.0.0        |                                 |
|O        |MSET              |1.0.1        |                                 |
|O        |MSETNX            |1.0.1        |                                 |
|-        |PSETNX            |2.6.0        |                                 |
|O        |SET               |1.0.0        |Any options are not supported yet|
|-        |SETEX             |2.0.0        |                                 |
|O        |SETNX             |2.0.0        |                                 |
|-        |SERANGE           |2.2.0        |                                 |
|O        |STRLEN            |2.2.0        |                                 |
|O        |SUBSTR            |1.0.0        |                                 |
