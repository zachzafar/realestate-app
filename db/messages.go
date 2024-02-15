package db

import "openlettings.com/types"

func (d *Database) CreateMessage(message *types.Message) (int, error) {
	query, variables := types.GeneratInsertQuery(message)
	query = query + " RETURNING message_id"
	var id int
	err := d.db.QueryRow(query, variables...).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *Database) GetPropertyMessageCount(propertyId int) (int, error) {
	query := `SELECT COUNT(*) AS row_count FROM messages WHERE property_id=$1`
	var count int

	err := d.db.QueryRow(query, propertyId).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (d *Database) GetUserMessageCount(userId int) (int, error) {
	query := `SELECT COUNT(*) AS row_count FROM messages WHERE user_id=$1`
	var count int

	err := d.db.QueryRow(query, userId).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (d *Database) GetMessagesByUser(user_id int) ([]*types.Message, error) {
	query := `SELECT * FROM messages WHERE user_id=$1`
	var messages []*types.Message

	rows, err := d.db.Query(query, user_id)

	if err != nil {
		return messages, err
	}

	for rows.Next() {
		var message *types.Message
		err = rows.Scan(&message)

		if err != nil {
			continue
		}

		messages = append(messages, message)
	}

	return messages, err

}

func (d *Database) GetMessageDetails(message_id int) (*types.Message, error) {
	message := &types.Message{}

	query, pointers := types.GenerateQueryByIDString(message)

	err := d.db.QueryRow(query, message_id).Scan(pointers...)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (d *Database) DeleteMessage(messageId int) error {
	query := `DELETE FROM messages WHERE message_id=$1`

	_, err := d.db.Exec(query, messageId)

	if err != nil {
		return err
	}

	return nil
}
