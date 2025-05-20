# HOW TO SET UP THE PROJECT

- TODOLIST - Add , update , delete , mark complete task

## Installation

- Instructions on how to get a copy of the project and running on your local machine.

### Step1: Clone Github Repository

```bash
git clone
cd Todolist
```

## Step 2: Install Go dependencies

```bash
go mod tidy
```

## Step 3: Build the Go binary

```bash
env GOOS=linux CGO_ENABLED=0 go build -o Todolist-app main.go
```

## Step 4: Build Docker image

```bash
docker build -f ./Dockerfile -t todolist-server:latest .
```

## Step 5: Run Kubernetes (Docker Desktop) and apply deployment.yaml

```bash
kubectl apply -f k8s/deployment.yaml
```

## Step 6: Deploy PostgreSQL to the cluster

```bash
kubectl apply -f k8s/postgres-deployment.yaml
```

## Step 7: Connect to PostgreSQL and create tables

```bash
kubectl run psql-client --rm -i -t --image=postgres:15 --restart=Never -- bash
```

- Inside the pod, connect to the database:

```bash
psql -h postgres-service -U postgres -d todolist
# Enter password: password
```

- Create the required table:

```bash
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
```

## Step 8: Port-forward the service to access locally

```bash
kubectl port-forward service/todolist-service 4000:80
```

- You can now access your API locally at:

```bash
http://localhost:4000
```

- Test API using Postman or curl.

## Endpoints

```bash
-  todo/create
   method - post
   body - {
    "title":"task4",
    "description":"I have to do this task4",
    "duedate":"2025-05-25T18:00:00Z"
    }
```

```bash
-  todo/get
   method - get
```

```bash
-  todo/get/{id}
   method - get
```

```bash
-  todo/update/{id}
   method - patch
   body - {
    "title":"task4",
    "description":"I have to do this task4",
    "duedate":"2025-05-25T18:00:00Z"
    }
```

```bash
-  todo/complete/{id}
   method - patch
```

```bash
- todo/delete/{id}
  method - Delete
```