USE station_db;

LOAD DATA INFILE
    '/docker-entrypoint-initdb.d/station20230327free.csv'
INTO TABLE
    station
FIELDS
TERMINATED BY ','
IGNORE 1 ROWS
(`id`, @dummy, `name`, @dummy, @dummy, @dummy, @dummy, @dummy, @dummy, @dummy, @dummy, @dummy, @dummy, @dummy, @dummy)
;
