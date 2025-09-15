type BookData record {|
    string name;
    string author;
    string synopsis;
    string isbn;
    string publisher;
    json metadata;
|};

type Book record {|
    int id;
    *BookData;
    string created_at;
|};
