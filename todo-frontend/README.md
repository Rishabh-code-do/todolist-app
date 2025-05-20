# HOW TO SET UP THE PROJECT

- TODOLIST - Frontend

## Installation

- Instructions on how to get a copy of the project and running on your local machine.

### Step1: Clone Github Repository

```bash
git clone https://github.com/Rishabh-code-do/todolist-app.git
cd todo-frontend
```

## Step 2: Build Docker image

```bash
docker build -t react-todo-app:latest .
```

## Step 3: Run Kubernetes (Docker Desktop) and apply deployment.yaml

```bash
kubectl apply -f k8s/deployment.yaml
```

## Step 4: Access the React app in your browser

```bash
http://localhost:30080
```

## Step 5: If you want to port-forward

```bash
kubectl port-forward deployment/react-app 8080:80
```

- Then open http://localhost:8080

