package database

var GenerationTable = `
	CREATE TABLE IF NOT EXISTS generation_tasks (
    id SERIAL PRIMARY KEY,
    generation_Id VARCHAR(8) UNIQUE NOT NULL,
    status VARCHAR(20) NOT NULL,
    file_content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`