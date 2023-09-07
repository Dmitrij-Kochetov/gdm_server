CREATE TABLE IF NOT EXISTS "projects" (
    project_id INTEGER NOT NULl,
    link VARCHAR NOT NULL,
    container_name VARCHAR NOT NULL UNIQUE,
    description VARCHAR,
    up_to_date BOOLEAN NOT NULL DEFAULT FALSE,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY(project_id)
);

CREATE TABLE IF NOT EXISTS "versions" (
    version_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    version VARCHAR NOT NULL,
    config BLOB,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY(version_id),
    FOREIGN KEY (project_id) REFERENCES "projects"(project_id)
);