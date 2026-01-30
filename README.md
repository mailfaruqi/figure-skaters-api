Figure Skating Elements API
API for figure skating elements with scoring and category

Database Schema
Table: categories
CREATE TABLE categories (
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
description TEXT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
Table: skating_elements
CREATE TABLE skating_elements (
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
code VARCHAR(20) NOT NULL,
category_id INTEGER REFERENCES categories(id),
base_value DECIMAL(4,2),
difficulty_level VARCHAR(20),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
Sample Data
Categories:
INSERT INTO categories (name, description) VALUES
('Jumps', 'Jumping elements in figure skating'),
('Spins', 'Spinning elements'),
('Steps', 'Footwork and step sequences'),
('Lifts', 'Partner lifting elements');
Elements:
INSERT INTO skating_elements (name, code, category_id, base_value, difficulty_level) VALUES
('Triple Axel', '3A', 1, 8.00, 'Advanced'),
('Quadruple Toe Loop', '4T', 1, 9.50, 'Advanced'),
('Double Lutz', '2Lz', 1, 2.10, 'Intermediate'),
('Flying Camel Spin', 'FCSp', 2, 2.00, 'Intermediate'),
('Sit Spin Level 4', 'SSp4', 2, 2.50, 'Intermediate'),
('Step Sequence Level 3', 'StSq3', 3, 3.30, 'Intermediate');
API Endpoints
Categories
GET /api/categories
Get all categories

curl http://localhost:8080/api/categories
GET /api/categories/{id}
Get category by ID

curl http://localhost:8080/api/categories/1
POST /api/categories
Create new category

curl -X POST http://localhost:8080/api/categories \
 -H "Content-Type: application/json" \
 -d '{
"name": "Jumps",
"description": "Jumping elements in figure skating"
}'
PUT /api/categories/{id}
Update category

curl -X PUT http://localhost:8080/api/categories/1 \
 -H "Content-Type: application/json" \
 -d '{
"name": "Jumps Updated",
"description": "Updated description"
}'
DELETE /api/categories/{id}
Delete category

curl -X DELETE http://localhost:8080/api/categories/1
Elements
GET /api/elements
Get all skating elements

curl http://localhost:8080/api/elements
GET /api/elements/{id}
Get element by ID (with category name via JOIN)

curl http://localhost:8080/api/elements/1
Response:

{
"id": 1,
"name": "Triple Axel",
"code": "3A",
"category_id": 1,
"category_name": "Jumps",
"base_value": 8.00,
"difficulty_level": "Advanced",
"created_at": "2025-01-30T12:00:00Z"
}
POST /api/elements
Create new element

curl -X POST http://localhost:8080/api/elements \
 -H "Content-Type: application/json" \
 -d '{
"name": "Triple Axel",
"code": "3A",
"category_id": 1,
"base_value": 8.00,
"difficulty_level": "Advanced"
}'
PUT /api/elements/{id}
Update element

curl -X PUT http://localhost:8080/api/elements/1 \
 -H "Content-Type: application/json" \
 -d '{
"name": "Triple Axel Updated",
"code": "3A",
"category_id": 1,
"base_value": 8.50,
"difficulty_level": "Advanced"
}'
DELETE /api/elements/{id}
Delete element

curl -X DELETE http://localhost:8080/api/elements/1
Health Check
GET /health
curl http://localhost:8080/health
Response:

{
"status": "OK",
"message": "Figure Skating Elements API Running"
}
Environment Variables
Create .env file:

PORT=8080
DB_CONN=postgresql://user:password@host:port/database?sslmode=require
Project Structure
figure-skaters-api/
├── main.go
├── database/
│ └── database.go
├── models/
│ └── models.go
├── repositories/
│ ├── category_repository.go
│ └── element_repository.go
├── services/
│ ├── category_service.go
│ └── element_service.go
└── handlers/
├── category_handler.go
└── element_handler.go
Run Locally
go mod init figure-skaters-api
go mod tidy
go run main.go
Deploy to Zeabur
Push code to GitHub
Connect repository to Zeabur
Add environment variables (PORT, DB_CONN)
Deploy!
