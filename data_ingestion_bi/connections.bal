import ballerinax/mysql;
import ballerinax/mysql.driver as _;

final mysql:Client mysqlClient = check new ("localhost", "appuser", "supersecret", "bookdb", 3306);
