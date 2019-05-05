# gibkv
Gibkv is a simple key value database storage. It stores data sequentially and relies on an base index file that points to the particular file/data entriy to retrieve the information.

## Index file
The index file consists of table metadata and indecies pointing to the offset inside the file of that particular table.
+---------+------------+----------------+
|Magic(4) | Version(2) | Table Count(4) |
+---------+-----------------------------+
|TableName(128)        | TableId(2)     |
+----------------------+----------------+
|              ...                      |
+---------+------------+----------------+
|Count(2) | Hash(4)    | Index(4)       |
+---------+------------+----------------+

## Data file
+---------+------------+----------------+
|Magic(4) | Version(2) | Id(2)          |
+---------+-----------------------------+
|Length(4)             | Data(len)      |
+----------------------+----------------+
|              ...                      |
+---------+------------+----------------+
|Length(4)             | Data(len)      |
+---------+------------+----------------+





