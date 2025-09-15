import ballerina/http;
import ballerina/sql;

listener http:Listener httpDefaultListener = http:getDefaultListener();

service /booklisting on httpDefaultListener {
    resource function get books() returns error|Book[] {
        do {
            Book[] books = [];
            stream<Book, sql:Error?> streamRowtypeSqlError = mysqlClient->query(`SELECT * FROM books`);
            check from Book book in streamRowtypeSqlError
                do {
                    // Can perform operations using the record 'book' of type `Book`.
                    books.push(book);
                };
            return books;
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    resource function post book(@http:Payload BookData book) returns error|json {
        do {
            string metadataStr = book.metadata.toJsonString();
            sql:ExecutionResult result = check mysqlClient->execute(`INSERT INTO books (name, author, synopsis, isbn, publisher, metadata) VALUES (${book.name}, ${book.author}, ${book.synopsis}, ${book.isbn}, ${book.publisher}, ${metadataStr})`);
            return {message: "Book added successfully", result: result.toString()};
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }
}
