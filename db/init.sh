#!/bin/bash

DB_NAME="test_app_db"
DB_USER="test_app_user"
DB_PASS="test_app_pass"

# Check if PostgreSQL is installed
if ! command -v psql &> /dev/null; then
    echo "PostgreSQL is not installed. Please install it first."
    exit 1
fi

# Create database user if it doesn't exist
psql postgres -tc "SELECT 1 FROM pg_roles WHERE rolname='$DB_USER'" | grep -q 1 || \
    psql postgres -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASS';"

# Create database if it doesn't exist
psql postgres -tc "SELECT 1 FROM pg_database WHERE datname='$DB_NAME'" | grep -q 1 || \
    psql postgres -c "CREATE DATABASE $DB_NAME OWNER $DB_USER;"

# Populate initial data
psql -U $DB_USER -d $DB_NAME <<EOF
-- Example table creation
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Example data insertion
INSERT INTO users (username, email)
VALUES ('admin', 'admin@testapp.local')
ON CONFLICT (username) DO NOTHING;
EOF

echo "Database '$DB_NAME' initialized and populated."