<!-- @format -->

HR Backend Project
Overview
This project focuses on building the backend of an HR application using Go (Golang). The main functionalities include:

Company Management: Register, login, and retrieve companies.
Employee Management: Create, login, retrieve, edit, and delete employees.
Onboarding Workflow Management: Create and assign onboarding workflows, retrieve all workflows, and get details of a specific workflow.
This HR backend project reimagines the employee onboarding and offboarding experiences, leveraging technology, automation, and streamlined procedures to make transitions as seamless and efficient as possible. The solution aims to significantly improve the HR processes for welcoming new employees and managing the exit of departing employees.

Technologies
Golang: The programming language used to develop the backend.
MongoDB: A NoSQL database used to store data for companies, employees, and onboarding workflows.
Features
Company Management
Register a Company

Endpoint: POST /companies/register
Description: Register a new company.
Login a Company

Endpoint: POST /companies/login
Description: Authenticate a company.
Get All Companies

Endpoint: GET /companies/all
Description: Retrieve a list of all registered companies.
Get a Single Company

Endpoint: GET /companies/:id
Description: Retrieve details of a specific company by ID.
Employee Management
Create an Employee

Endpoint: POST /employees/create
Description: Create a new employee.
Login an Employee

Endpoint: POST /employees/login
Description: Authenticate an employee.
Get All Employees

Endpoint: GET /employees/all
Description: Retrieve a list of all employees.
Get Employees in a Company

Endpoint: GET /employees/all/:companyId
Description: Retrieve a list of all employees in a specific company.
Get a Single Employee

Endpoint: GET /employees/:id
Description: Retrieve details of a specific employee by ID.
Edit an Employee

Endpoint: PUT /employees/:id
Description: Edit details of a specific employee.
Delete an Employee

Endpoint: DELETE /employees/:id
Description: Delete a specific employee.
Onboarding Workflow Management
Create an Onboarding Workflow

Endpoint: POST /onboarding-workflow
Description: Create a new onboarding workflow.
Get All Onboarding Workflows

Endpoint: GET /onboarding-workflow/all
Description: Retrieve a list of all onboarding workflows.
Get a Single Onboarding Workflow

Endpoint: GET /onboarding-workflow/:id
Description: Retrieve details of a specific onboarding workflow by ID.
Assign a Workflow

Endpoint: POST /onboarding-workflow/assign
Description: Assign an onboarding workflow to an employee.
