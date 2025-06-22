FlowBoard

FlowBoard is a lightweight CRM / mini-ERP frontend for small businesses, built with React-admin.

It provides a simple and extensible UI for managing clients, deals, statuses, tags, users, and comments — all with secure JWT-based authentication.

✨ Features

Full CRUD for:

Clients

Deals & statuses

Tags

Users

Comments

Search and filtering by key fields

Role-based access control (delete actions restricted to admins)

JWT login/logout authentication

Clean and efficient forms for creating and editing records

🚀 Getting Started

Install dependencies:

npm install

Start the development server:

npm start

Open http://localhost:3000 in your browser.

🔧 Backend Configuration

By default, the frontend expects the backend API to be available at:

http://localhost:8080

If your backend is running elsewhere, update the apiUrl constant in:

src/App.js

🔐 Authentication

Login using an email and password of a registered user.

To access admin-only features (e.g., delete), the user must have the role: admin.

🛠 Tech Stack

React
React-admin
Material UI
Axios
ra-data-simple-rest
