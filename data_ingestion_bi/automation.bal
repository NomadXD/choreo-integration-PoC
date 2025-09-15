import ballerina/log;
import ballerina/sql;

public function main() returns error? {
    do {
        sql:ExecutionResult sqlExecutionresult = check mysqlClient->execute(`
        CREATE TABLE IF NOT EXISTS books (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            author VARCHAR(255) NOT NULL,
            synopsis TEXT,
            isbn VARCHAR(20) NOT NULL,
            publisher VARCHAR(255) NOT NULL,
            metadata JSON,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            INDEX idx_name (name)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`);
        log:printInfo("Table created successfully", 'data = sqlExecutionresult);
    } on fail error e {
        log:printError("Error occurred", 'error = e);
        return e;
    }
}
