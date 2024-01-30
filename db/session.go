package db

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"

	"openlettings.com/types"
)

func generateID() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	sessionID := hex.EncodeToString(randomBytes)
	return sessionID, nil
}

func (d *Database) CreateSession(user *types.User) (string, error) {
	query := `INSERT INTO sessions (session_id,data) VALUES ($1,$2)`
	sessionData := &types.SessionData{
		UserId: user.UserId,
	}

	sessionDataBytes, err := json.Marshal(sessionData) //

	if err != nil {

		return "", err
	}

	sessionId, err := generateID()

	_, err = d.db.Exec(query, sessionId, sessionDataBytes)

	if err != nil {

		return "", err
	}

	return sessionId, nil
}

func (d *Database) GetSessionData(sessionID string) (*types.SessionData, error) {
	query := `SELECT data FROM sessions WHERE session_id=$1`
	var sessionByteData []byte
	err := d.db.QueryRow(query, sessionID).Scan(&sessionByteData)

	if err != nil {
		return nil, err
	}

	var sessionData types.SessionData

	err = json.Unmarshal(sessionByteData, &sessionData)

	if err != nil {
		return nil, err
	}

	return &sessionData, nil

}

func (d *Database) DeleteSession(sessionID string) error {
	query := `DELETE FROM sessions WHERE session_id=$1`
	_, err := d.db.Exec(query, sessionID)

	if err != nil {
		return err
	}

	return nil
}
