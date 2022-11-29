package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/core"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/secondary_adapter/repositories/user"
)

const (
	userDatabase   = "user"
	userCollection = "users"
)

type Cfg struct {
	connectionString connstring.ConnString
}

func NewMongoCfg(connString string) (*Cfg, error) {
	parsedConnStr, err := connstring.Parse(connString)
	if err != nil {
		err := fmt.Errorf("parse Mongo connstring : %s, err : %w", connString, err)
		log.Print(err)
		return nil, err
	}

	return &Cfg{
		connectionString: parsedConnStr,
	}, nil
}

type UserRepository struct {
	cfg    *Cfg
	client *mongo.Client
}

func (u UserRepository) newCollection() *mongo.Collection {
	return u.client.Database(userDatabase).Collection(userCollection)
}

func (u UserRepository) Get() ([]core.User, error) {
	cursor, err := u.newCollection().Find(context.Background(), options.Find())
	if err != nil {
		if errors.Is(err, mongo.ErrNilDocument) {
			return nil, nil
		}
		err := fmt.Errorf("get users: %w", err)
		log.Print(err)
		return nil, err
	}

	var users []core.User
	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, fmt.Errorf("user: failed to get user: %w", err)
	}

	return users, nil
}

func (u UserRepository) Create(user *core.User) error {
	res, err := u.newCollection().InsertOne(context.Background(), user)
	if err != nil {
		err := fmt.Errorf("create user: %w", err)
		log.Print(err)
		return err
	}

	log.Printf("created user: %s", res.InsertedID)
	return nil
}

func (u UserRepository) Shutdown(ctx context.Context) error {
	err := u.client.Disconnect(ctx)
	return err
}

func NewUserRepository(ctx context.Context, cfg *Cfg) (user.Repository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.connectionString.String()))
	if err != nil {
		err := fmt.Errorf("connect to Mongo : %s, err : %w", cfg.connectionString.String(), err)
		log.Print(err)
		return nil, err
	}

	return UserRepository{
		cfg:    cfg,
		client: client,
	}, nil
}
