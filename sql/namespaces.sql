CREATE TABLE IF NOT EXISTS namespaces (
	etag VARCHAR(512),
    id VARCHAR PRIMARY KEY,
    created TIMESTAMP,
    updated TIMESTAMP,

    name VARCHAR(512),

    CONSTRAINT namespace_id UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS namespaces_settings (
    etag VARCHAR(512),
    id VARCHAR PRIMARY KEY,
    created TIMESTAMP,
    updated TIMESTAMP,

    namespace VARCHAR REFERENCES namespaces (id),

    key VARCHAR(512),
    value VARCHAR(512),

    CONSTRAINT setting_in_namespace UNIQUE (namespace, key)
);