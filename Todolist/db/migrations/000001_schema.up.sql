CREATE TYPE task_status AS ENUM ('PENDING', 'COMPLETED');

CREATE TABLE todo(
    id SERIAL PRIMARY KEY,
    title text NOT NULL,
    description text,
    dueDate date,
    status task_status DEFAULT 'PENDING',
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);