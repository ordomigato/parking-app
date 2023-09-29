CREATE TABLE IF NOT EXISTS client (
	client_uuid TEXT NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    last_login TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS workspace (
	workspace_uuid TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS client_workspace (
    client_uuid TEXT NOT NULL REFERENCES client (client_uuid),
    workspace_uuid TEXT NOT NULL REFERENCES workspace (workspace_uuid),
    role_id TEXT NOT NULL,
    PRIMARY KEY (client_uuid, workspace_uuid)
);

CREATE TABLE IF NOT EXISTS role (
    id TEXT NOT NULL PRIMARY KEY,
  	name TEXT NOT NULL
);

INSERT INTO role(id, name)
VALUES
    ('superadmin', 'Superadmin'),
    ('admin', 'Admin');

CREATE TABLE IF NOT EXISTS permit_form (
    permit_form_uuid TEXT NOT NULL PRIMARY KEY,
    workspace_uuid TEXT NOT NULL REFERENCES workspace (workspace_uuid),
    name TEXT NOT NULL,
    submission_constraint_type TEXT DEFAULT NULL,
    submission_constraint_limit SMALLINT DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS submission_constraint_type (
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS form_notice (
    notice_uuid TEXT NOT NULL PRIMARY KEY,
    notice_type TEXT NOT NULL,
    message TEXT DEFAULT NULL,
    permit_form_uuid TEXT NOT NULL REFERENCES permit_form (permit_form_uuid)
);

CREATE TABLE IF NOT EXISTS consumer (
	consumer_uuid TEXT NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    last_login TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS permit (
    permit_uuid TEXT NOT NULL PRIMARY KEY,
    permit_form_uuid TEXT NOT NULL REFERENCES permit_form (permit_form_uuid),
    consumer_uuid TEXT DEFAULT NULL REFERENCES consumer (consumer_uuid),
    response_data JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS field_type (
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO field_type(id, name)
VALUES
    ('text-input', 'Text'),
    ('phone-input', 'Phone'),
    ('email-input', 'Email'),
    ('first-name-input', 'First Name'),
    ('last-name-input', 'Last Name'),
    ('select-input', 'Dropdown');

CREATE TABLE IF NOT EXISTS question (
    question_uuid TEXT NOT NULL PRIMARY KEY,
    question TEXT NOT NULL,
    permit_form_uuid TEXT NOT NULL REFERENCES permit_form (permit_form_uuid),
    field_type TEXT NOT NULL REFERENCES field_type (id)
);


