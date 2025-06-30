# Guia de Uso do CRUD MongoDB com MGM

Este guia mostra como usar o sistema CRUD extensível para MongoDB com a biblioteca MGM.

## 🚀 Configuração Inicial

### 1. Definir Variáveis de Ambiente

```bash
export MONGO_URI="mongodb://localhost:27017"
export MONGO_DATABASE="training_db"
export APP_NAME="training-api"
```

### 2. Inicializar a Conexão

```go
package main

import (
    "github.com/cvpose/repository"
)

func main() {
    // Inicializar a conexão com MongoDB
    repository.InitDatabase()
    
    // Criar o repository
    repo := repository.NewTrainingRepository()
    
    // Usar o repository...
}
```

## 📋 Operações CRUD Básicas

### ✨ CREATE - Criar um Novo Training

```go
import (
    "context"
    "fmt"
)

func createTraining() {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    training := &repository.Training{
        Name:        "Go Programming Avançado",
        Description: "Aprenda conceitos avançados de Go",
        ImageURL:    "https://example.com/go-training.jpg",
    }
    
    err := repo.CreateTraining(ctx, training)
    if err != nil {
        fmt.Printf("Erro ao criar training: %v\n", err)
        return
    }
    
    fmt.Printf("Training criado com ID: %s\n", training.ID.Hex())
}
```

### 🔍 READ - Buscar Trainings

#### Buscar por ID

```go
func getTrainingByID(id string) {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    training, err := repo.GetTrainingByID(ctx, id)
    if err != nil {
        fmt.Printf("Erro ao buscar training: %v\n", err)
        return
    }
    
    fmt.Printf("Training encontrado: %s - %s\n", training.Name, training.Description)
}
```

#### Buscar Todos

```go
func getAllTrainings() {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    trainings, err := repo.GetAllTrainings(ctx, nil)
    if err != nil {
        fmt.Printf("Erro ao buscar trainings: %v\n", err)
        return
    }
    
    fmt.Printf("Encontrados %d trainings:\n", len(trainings))
    for _, training := range trainings {
        fmt.Printf("  - %s: %s\n", training.Name, training.Description)
    }
}
```

#### Buscar com Filtros

```go
import "go.mongodb.org/mongo-driver/bson"

func searchTrainings() {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    // Buscar trainings que contenham "Go" no nome (case insensitive)
    filter := bson.M{
        "name": bson.M{
            "$regex":   "Go",
            "$options": "i",
        },
    }
    
    trainings, err := repo.GetAllTrainings(ctx, filter)
    if err != nil {
        fmt.Printf("Erro na busca: %v\n", err)
        return
    }
    
    fmt.Printf("Trainings com 'Go' no nome: %d\n", len(trainings))
}
```

### ✏️ UPDATE - Atualizar Training

```go
func updateTraining(id string) {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    // Atualizar apenas a descrição
    updatedTraining, err := repo.UpdateTrainingDescription(
        ctx,
        id,
        "Nova descrição atualizada com mais detalhes",
    )
    if err != nil {
        fmt.Printf("Erro ao atualizar: %v\n", err)
        return
    }
    
    fmt.Printf("Training atualizado: %s\n", updatedTraining.Description)
}
```

### 🗑️ DELETE - Deletar Training

```go
func deleteTraining(id string) {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    err := repo.DeleteTraining(ctx, id)
    if err != nil {
        fmt.Printf("Erro ao deletar: %v\n", err)
        return
    }
    
    fmt.Println("Training deletado com sucesso!")
}
```

## 🔧 Funcionalidades Avançadas

### 📄 Paginação

```go
func getTrainingsWithPagination() {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    page := int64(1)
    limit := int64(5)
    
    trainings, total, err := repo.GetTrainingsWithPagination(ctx, nil, page, limit)
    if err != nil {
        fmt.Printf("Erro na paginação: %v\n", err)
        return
    }
    
    fmt.Printf("Página %d: %d trainings de %d total\n", page, len(trainings), total)
    
    // Calcular número de páginas
    totalPages := (total + limit - 1) / limit
    fmt.Printf("Total de páginas: %d\n", totalPages)
}
```

### 🔢 Contagem

```go
func countTrainings() {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    // Contar todos os trainings
    total, err := repo.Count(ctx, nil)
    if err != nil {
        fmt.Printf("Erro ao contar: %v\n", err)
        return
    }
    
    fmt.Printf("Total de trainings: %d\n", total)
    
    // Contar com filtro
    filter := bson.M{"name": bson.M{"$regex": "Go", "$options": "i"}}
    goTrainings, err := repo.Count(ctx, filter)
    if err != nil {
        fmt.Printf("Erro ao contar com filtro: %v\n", err)
        return
    }
    
    fmt.Printf("Trainings com 'Go': %d\n", goTrainings)
}
```

### 🗂️ Ordenação

```go
import "go.mongodb.org/mongo-driver/mongo/options"

func getTrainingsSorted() {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    // Ordenar por nome (ascendente)
    opts := options.Find().SetSort(bson.M{"name": 1})
    trainings, err := repo.GetAllTrainings(ctx, nil, opts)
    if err != nil {
        fmt.Printf("Erro ao buscar ordenado: %v\n", err)
        return
    }
    
    fmt.Println("Trainings ordenados por nome:")
    for _, training := range trainings {
        fmt.Printf("  - %s\n", training.Name)
    }
}
```

### 🔍 Busca por Nome Específico

```go
func searchByName() {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    // Buscar trainings que contenham "Docker" no nome
    trainings, err := repo.GetTrainingsByName(ctx, "Docker")
    if err != nil {
        fmt.Printf("Erro na busca por nome: %v\n", err)
        return
    }
    
    fmt.Printf("Trainings com 'Docker': %d\n", len(trainings))
    for _, training := range trainings {
        fmt.Printf("  - %s\n", training.Name)
    }
}
```

## 🎯 Exemplo Completo de Uso

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/cvpose/repository"
    "go.mongodb.org/mongo-driver/bson"
)

func main() {
    // 1. Inicializar conexão
    repository.InitDatabase()
    
    // 2. Criar repository
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    // 3. Criar alguns trainings
    trainings := []*repository.Training{
        {
            Name:        "Go Fundamentals",
            Description: "Aprenda os fundamentos do Go",
            ImageURL:    "https://example.com/go.jpg",
        },
        {
            Name:        "Docker Essentials",
            Description: "Containerização com Docker",
            ImageURL:    "https://example.com/docker.jpg",
        },
        {
            Name:        "Kubernetes Basics",
            Description: "Orquestração com Kubernetes",
            ImageURL:    "https://example.com/k8s.jpg",
        },
    }
    
    fmt.Println("Criando trainings...")
    for _, training := range trainings {
        err := repo.CreateTraining(ctx, training)
        if err != nil {
            log.Printf("Erro ao criar %s: %v", training.Name, err)
        } else {
            fmt.Printf("✓ Criado: %s\n", training.Name)
        }
    }
    
    // 4. Buscar todos
    fmt.Println("\nBuscando todos os trainings...")
    allTrainings, err := repo.GetAllTrainings(ctx, nil)
    if err != nil {
        log.Printf("Erro ao buscar: %v", err)
        return
    }
    
    fmt.Printf("Total: %d trainings\n", len(allTrainings))
    
    // 5. Buscar com filtro
    fmt.Println("\nBuscando trainings com 'Go'...")
    filter := bson.M{"name": bson.M{"$regex": "Go", "$options": "i"}}
    goTrainings, err := repo.GetAllTrainings(ctx, filter)
    if err != nil {
        log.Printf("Erro na busca: %v", err)
        return
    }
    
    for _, training := range goTrainings {
        fmt.Printf("  - %s\n", training.Name)
    }
    
    // 6. Atualizar um training
    if len(allTrainings) > 0 {
        fmt.Println("\nAtualizando primeiro training...")
        updated, err := repo.UpdateTrainingDescription(
            ctx,
            allTrainings[0].ID.Hex(),
            "Descrição atualizada!",
        )
        if err != nil {
            log.Printf("Erro ao atualizar: %v", err)
        } else {
            fmt.Printf("✓ Atualizado: %s\n", updated.Description)
        }
    }
    
    // 7. Paginação
    fmt.Println("\nTestando paginação...")
    pagedTrainings, total, err := repo.GetTrainingsWithPagination(ctx, nil, 1, 2)
    if err != nil {
        log.Printf("Erro na paginação: %v", err)
    } else {
        fmt.Printf("Página 1 (2 por página): %d de %d total\n", 
            len(pagedTrainings), total)
    }
    
    fmt.Println("\n🎉 Exemplo completo executado com sucesso!")
}
```

## 🔧 Extensão para Outros Modelos

### Criar um Novo Modelo (User)

```go
// user.go
package repository

import "github.com/kamva/mgm/v3"

type User struct {
    *mgm.DefaultModel `bson:",inline"`
    Username          string `json:"username" bson:"username"`
    Email             string `json:"email" bson:"email"`
    Age               int    `json:"age" bson:"age"`
}

func (u *User) CollectionName() string {
    return "users"
}
```

### Criar Repository Específico

```go
// user_repository.go
package repository

import (
    "context"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
    *Repository
}

func NewUserRepository() *UserRepository {
    user := &User{}
    return &UserRepository{
        Repository: NewRepository(user),
    }
}

// Métodos específicos para User
func (ur *UserRepository) CreateUser(ctx context.Context, user *User) error {
    if user.Username == "" {
        return errors.New("username é obrigatório")
    }
    if user.Email == "" {
        return errors.New("email é obrigatório")
    }
    return ur.Create(ctx, user)
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
    user := &User{}
    filter := bson.M{"email": email}
    err := ur.FindOne(ctx, filter, user)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (ur *UserRepository) GetUsersByAgeRange(ctx context.Context, minAge, maxAge int) ([]*User, error) {
    var users []*User
    filter := bson.M{
        "age": bson.M{
            "$gte": minAge,
            "$lte": maxAge,
        },
    }
    err := ur.Find(ctx, filter, &users)
    if err != nil {
        return nil, err
    }
    return users, nil
}
```

## 🚨 Tratamento de Erros

```go
func handleErrors() {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    // Erro de ID inválido
    _, err := repo.GetTrainingByID(ctx, "invalid-id")
    if err != nil {
        if err.Error() == "invalid object ID format" {
            fmt.Println("ID fornecido não é válido")
        }
    }
    
    // Erro de documento não encontrado
    _, err = repo.GetTrainingByID(ctx, "64f5e7b8c9a8d1234567890a")
    if err != nil {
        if err.Error() == "document not found" {
            fmt.Println("Training não encontrado")
        }
    }
    
    // Erro de conexão
    err = repo.CreateTraining(ctx, &repository.Training{})
    if err != nil {
        fmt.Printf("Erro na operação: %v\n", err)
    }
}
```

## 💡 Dicas e Boas Práticas

### 1. Validação antes de criar

```go
func createWithValidation(training *repository.Training) error {
    if training.Name == "" {
        return errors.New("nome é obrigatório")
    }
    if training.Description == "" {
        return errors.New("descrição é obrigatória")
    }
    if training.ImageURL == "" {
        return errors.New("URL da imagem é obrigatória")
    }
    
    repo := repository.NewTrainingRepository()
    return repo.CreateTraining(context.Background(), training)
}
```

### 2. Uso de contexto com timeout

```go
import "time"

func operationWithTimeout() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    repo := repository.NewTrainingRepository()
    trainings, err := repo.GetAllTrainings(ctx, nil)
    if err != nil {
        fmt.Printf("Operação cancelada ou erro: %v\n", err)
        return
    }
    
    fmt.Printf("Encontrados %d trainings\n", len(trainings))
}
```

### 3. Filtros complexos

```go
func complexFilters() {
    repo := repository.NewTrainingRepository()
    ctx := context.Background()
    
    // Múltiplas condições
    filter := bson.M{
        "$and": []bson.M{
            {"name": bson.M{"$regex": "Go", "$options": "i"}},
            {"description": bson.M{"$regex": "avançado", "$options": "i"}},
        },
    }
    
    trainings, err := repo.GetAllTrainings(ctx, filter)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }
    
    fmt.Printf("Trainings avançados de Go: %d\n", len(trainings))
}
```

Este guia cobre todas as funcionalidades principais do sistema CRUD. Para mais exemplos específicos, consulte o arquivo `example_usage.go` no projeto.
