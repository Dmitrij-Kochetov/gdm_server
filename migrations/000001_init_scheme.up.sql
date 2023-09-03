CREATE TABLE IF NOT EXISTS "projects" (
    ID INTEGER NOT NULl,
    link VARCHAR NOT NULL,
    container_name VARCHAR NOT NULL UNIQUE,
    PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS "project_info" (
    ID INTEGER NOT NULL,
    description VARCHAR,
    up_to_date BOOLEAN NOT NULL,
    running BOOLEAN NOT NULL,
    deleted BOOLEAN NOT NULL,
    PRIMARY KEY(ID),
    FOREIGN KEY (ID) REFERENCES "projects"(ID)
);

CREATE TABLE IF NOT EXISTS "versions" (
    ID INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    description VARCHAR,
    version VARCHAR NOT NULL,
    up_to_date BOOLEAN NOT NULL,
    config BLOB NOT NULL,
    PRIMARY KEY(ID),
    FOREIGN KEY (project_id) REFERENCES "projects"(ID)
);