# Portfolio Server Homework

This is a simple server that serves a portfolio website. It is built using Go, HTMX, RippleUI and TailwindCSS.

## Features
- Toggle between dark and light mode
- Admin panel for changing informations, email and password
- Login and register system
- Session management

## Installation
1. Clone the repository
```bash
git clone https://github.com/mt190502/portfolio.git
```

2. Change directory to the project folder
```bash
cd portfolio
```

3. Install the dependencies
```bash
go mod tidy
```

4. Copy config-example.yaml to config.yaml and fill in the required fields
```bash
cp config-example.yaml config.yaml
vim config.yaml
```

5. Run the server
```bash
go run src/main.go
```


# Screenshots

<details>
  <summary>Main Page</summary>

| Dark | Light |
| ---- | ----- |
| ![image](https://github.com/mt190502/portfolio/assets/62564400/77ef4472-3af3-4a3f-a2c2-5c06c14a182f) | ![image](https://github.com/mt190502/portfolio/assets/62564400/6ca0a4d4-7a13-4260-8bba-9a48fad91e89) |
  
</details>


<details>
  <summary>Login Page (/login)</summary>

| Dark | Light |
| ---- | ----- |
| ![image](https://github.com/mt190502/portfolio/assets/62564400/7e49709b-cec8-4a69-8359-378d192c0052) | ![image](https://github.com/mt190502/portfolio/assets/62564400/e8b19410-d711-4040-92c0-76bf1962957f) |
| ![image](https://github.com/mt190502/portfolio/assets/62564400/015fb296-f86d-49f0-865f-6a6116b0bb43) | ![image](https://github.com/mt190502/portfolio/assets/62564400/bc7b990c-f493-4f2b-84b5-04fdcfacf781) |

</details>


<details>
  <summary>Register Page (/register)</summary>

| Dark | Light |
| ---- | ----- |
| ![image](https://github.com/mt190502/portfolio/assets/62564400/04e07dd1-8e43-4014-abb6-500dbc8caf36) | ![image](https://github.com/mt190502/portfolio/assets/62564400/cb09e7e9-c374-4670-899a-9e60a6df40c9) |
| ![image](https://github.com/mt190502/portfolio/assets/62564400/23cb80e4-912f-45e5-b461-15161ab8ae5d) | ![image](https://github.com/mt190502/portfolio/assets/62564400/22b62656-a057-4377-b9f4-c6f7b5ad7f94) |
| ![image](https://github.com/mt190502/portfolio/assets/62564400/a5b5c983-ca1e-4263-990c-eaf19a31347d) | ![image](https://github.com/mt190502/portfolio/assets/62564400/c52f7e5f-54fe-4f84-85eb-10e0f13d0768) |
  
</details>


<details>
  <summary>Admin Page (/admin)</summary>

| Dark | Light |
| ---- | ----- |
| ![image](https://github.com/mt190502/portfolio/assets/62564400/e82f2016-efb6-435a-a754-e38732073080)  | ![image](https://github.com/mt190502/portfolio/assets/62564400/24a3b80a-56ba-4b76-a593-eb1a4abc0543) |
| ![image](https://github.com/mt190502/portfolio/assets/62564400/ca528811-5765-46fd-9cf4-a4ea464ed09c)  | ![image](https://github.com/mt190502/portfolio/assets/62564400/d3dbf78d-29b9-42ce-bb70-aecd0a600876) |  
| ![image](https://github.com/mt190502/portfolio/assets/62564400/16dcf682-e89d-4df4-9762-4b346522e28d)  | ![image](https://github.com/mt190502/portfolio/assets/62564400/728714fe-c1dc-43ca-9420-4f655d39fbf2) |
  
</details>
