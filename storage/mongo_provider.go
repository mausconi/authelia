package storage

import (
	"context"
	"time"

	"github.com/clems4ever/authelia/configuration/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/clems4ever/authelia/models"
)

// MongoProvider is a storage provider persisting data in a SQLite database.
type MongoProvider struct {
	configuration schema.MongoStorageConfiguration
}

// NewMongoProvider construct a mongo provider.
func NewMongoProvider(configuration schema.MongoStorageConfiguration) *MongoProvider {
	return &MongoProvider{configuration}
}

func (p *MongoProvider) connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credentials := options.Credential{
		Username: p.configuration.Auth.Username,
		Password: p.configuration.Auth.Password,
	}
	clientOptions := options.Client().ApplyURI(p.configuration.URL)
	clientOptions.SetAuth(credentials)
	return mongo.Connect(ctx, clientOptions)
}

type prefered2FAMethod struct {
	UserID string `json:"userId"`
	Method string `json:"method"`
}

// LoadPrefered2FAMethod load the prefered method for 2FA from sqlite db.
func (p *MongoProvider) LoadPrefered2FAMethod(username string) (string, error) {
	client, err := p.connect()
	if err != nil {
		return "", nil
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(p.configuration.Database).Collection("prefered_2fa_method")

	res := prefered2FAMethod{}
	err = collection.FindOne(context.Background(), bson.M{"userId": username}).Decode(&res)

	if err != nil {
		return "", err
	}

	return res.Method, nil
}

// SavePrefered2FAMethod save the prefered method for 2FA in sqlite db.
func (p *MongoProvider) SavePrefered2FAMethod(username string, method string) error {
	client, err := p.connect()
	if err != nil {
		return nil
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(p.configuration.Database).Collection("prefered_2fa_method")

	updateOptions := options.ReplaceOptions{}
	updateOptions.SetUpsert(true)
	_, err = collection.ReplaceOne(context.Background(), bson.M{"userId": username},
		bson.M{"userId": username, "method": method}, &updateOptions)

	if err != nil {
		return err
	}

	return nil
}

// FindIdentityVerificationToken look for an identity verification token in DB.
func (p *MongoProvider) FindIdentityVerificationToken(token string) (bool, error) {
	return false, nil
}

// SaveIdentityVerificationToken save an identity verification token in DB.
func (p *MongoProvider) SaveIdentityVerificationToken(token string) error {
	return nil
}

// RemoveIdentityVerificationToken remove an identity verification token from the DB.
func (p *MongoProvider) RemoveIdentityVerificationToken(token string) error {
	return nil
}

// SaveTOTPSecret save a TOTP secret of a given user.
func (p *MongoProvider) SaveTOTPSecret(username string, secret string) error {
	return nil
}

// LoadTOTPSecret load a TOTP secret given a username.
func (p *MongoProvider) LoadTOTPSecret(username string) (string, error) {
	return "", nil
}

// SaveU2FRegistration save a registered U2F device registration blob.
func (p *MongoProvider) SaveU2FRegistration(username string, keyHandle []byte) error {
	return nil
}

// LoadU2FRegistration load a U2F device registration blob for a given username.
func (p *MongoProvider) LoadU2FRegistration(username string) ([]byte, error) {
	return nil, nil
}

// AppendAuthenticationLog append a mark to the authentication log.
func (p *MongoProvider) AppendAuthenticationLog(attempt models.AuthenticationAttempt) error {
	return nil
}

// LoadLatestAuthenticationLogs retrieve the latest marks from the authentication log.
func (p *MongoProvider) LoadLatestAuthenticationLogs(username string, fromDate time.Time) ([]models.AuthenticationAttempt, error) {
	return nil, nil
}
