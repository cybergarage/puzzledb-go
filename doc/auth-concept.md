# Authentication Methods

PuzzleDB adds authenticators to the auth manager based on the auth configuration with auth plugins. The query plugins can query the authentication from the auth manager.

<figure>
<img src="img/authenticator.png" alt="authenticator" />
</figure>

An authenticator is generated from one auth plugin and a configuration. Multiple authenticators are generated and added to the auth manager when PuzzleDB starts.

## Authentication Plugins

PuzzleDB offers the following common authentication plugins for query plugins as default.

-   Password

-   MD5 (Not yes supported)

-   Crypt (Not yes supported)

-   SHA256 (Not yes supported)

-   SHA512 (Not yes supported)

-   GSSAPI (Not yes supported)

-   SSPI (Not yes supported)

-   LDAP (Not yes supported)

-   PAM (Not yes supported)

-   Kerberos (Not yes supported)

## Supported Authentication Methods

PuzzleDB supports the following authentication methods for the query plugins.

<table style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Method</th>
<th style="text-align: left;">Parameter</th>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">MongoDB</th>
<th style="text-align: left;">Redis</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>password</p></td>
<td style="text-align: left;"><p>user</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>X</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>password</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>O</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>database</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>X</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>address</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>X</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>md5</p></td>
<td style="text-align: left;"><p>user</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>password</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>database</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>address</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>crypt</p></td>
<td style="text-align: left;"><p>user</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>password</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>database</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>address</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
</tbody>
</table>

O:Supported, X:Unsupported, -:Not yes supported

## References

### PostgreSQL

-   [PostgreSQL: Documentation: Authentication Methods](https://www.postgresql.org/docs/current/auth-methods.html)

    -   [PostgreSQL: Documentation: The pg\_hba.conf File](https://www.postgresql.org/docs/current/auth-pg-hba-conf.html)

## MySQL

-   [MySQL: Connection Phase](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase.html)

-   [MySQL: Authentication Methods](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_authentication_methods.html)

    -   [MySQL: Old Password Authentication](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_authentication_methods.html#page_protocol_connection_phase_authentication_methods_old_password_authentication)

    -   [MySQL: Native Password Authentication](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_authentication_methods_native_password_authentication.html)

## MongoDB

-   [Security — MongoDB Manual](https://www.mongodb.com/docs/manual/security/)

    -   [Authentication — MongoDB Manual](https://www.mongodb.com/docs/manual/core/authentication/)

    -   [Configure Database User Authentication — MongoDB Atlas](https://www.mongodb.com/docs/atlas/security/config-db-auth/)

## Redis

-   [Security – Redis](https://redis.io/docs/management/security/)

    -   [AUTH | Redis](https://redis.io/commands/auth/)
