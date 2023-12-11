Authentication Methods
======================

PuzzleDB adds authenticators to the auth manager based on the auth configuration with auth plugins. The query plugins can query the authentication from the auth manager.

![authenticator](img/authenticator.png)

An authenticator is generated from one auth plugin and a configuration. Multiple authenticators are generated and added to the auth manager when PuzzleDB starts.

Authentication Plugins
----------------------

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

Supported Authentication Methods
--------------------------------

PuzzleDB supports the following authentication methods for the query plugins.

<table style="width:100%;"><colgroup><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /></colgroup><thead><tr class="header"><th>Method</th><th>Parameter</th><th>PostgreSQL</th><th>MySQL</th><th>MongoDB</th><th>Redis</th></tr></thead><tbody><tr class="odd"><td><p>password</p></td><td><p>user</p></td><td><p>O</p></td><td><p>-</p></td><td><p>-</p></td><td><p>X</p></td></tr><tr class="even"><td></td><td><p>password</p></td><td><p>O</p></td><td><p>-</p></td><td><p>-</p></td><td><p>O</p></td></tr><tr class="odd"><td></td><td><p>database</p></td><td><p>O</p></td><td><p>-</p></td><td><p>-</p></td><td><p>X</p></td></tr><tr class="even"><td></td><td><p>address</p></td><td><p>O</p></td><td><p>-</p></td><td><p>-</p></td><td><p>X</p></td></tr><tr class="odd"><td><p>md5</p></td><td><p>user</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr><tr class="even"><td></td><td><p>password</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr><tr class="odd"><td></td><td><p>database</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr><tr class="even"><td></td><td><p>address</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr><tr class="odd"><td><p>crypt</p></td><td><p>user</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr><tr class="even"><td></td><td><p>password</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr><tr class="odd"><td></td><td><p>database</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr><tr class="even"><td></td><td><p>address</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr></tbody></table>

O:Supported, X:Unsupported, -:Not yes supported

References
----------

### PostgreSQL

-   [PostgreSQL: Documentation: Authentication Methods](https://www.postgresql.org/docs/current/auth-methods.html)

    -   [PostgreSQL: Documentation: The pg\_hba.conf File](https://www.postgresql.org/docs/current/auth-pg-hba-conf.html)

MySQL
-----

-   [MySQL: Connection Phase](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase.html)

-   [MySQL: Authentication Methods](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_authentication_methods.html)

    -   [MySQL: Old Password Authentication](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_authentication_methods.html#page_protocol_connection_phase_authentication_methods_old_password_authentication)

    -   [MySQL: Native Password Authentication](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_authentication_methods_native_password_authentication.html)
