package db

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"openlettings.com/types"
	"openlettings.com/utils"
)

func (d *Database) CreateProperty(property *types.Property) (int64, error) {
	query := `INSERT INTO properties (title,description,property_type,address,city,size,bedrooms,bathrooms,year_built,price,user_id) VALUES ($1, $2, $3, $4, $5, $6,$7,$8, $9,$10,$11)`
	result, err := d.db.Exec(query, property.Title, property.Description, property.PropertyType, property.Address, property.City, property.Size, property.Bedrooms, property.Bathrooms, property.YearBuilt, property.Price, property.UserID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

func (d *Database) UploadPropertyPhotos(files []*multipart.FileHeader, id int64) error {
	query := `INSERT INTO media (property_id,url) VALUES ($1,$2)`
	id_string := strconv.FormatInt(id, 10)
	fmt.Println("processing images...")
	if err := os.MkdirAll("./media/properties/"+id_string, 0777); err != nil {
		return utils.CustomError{Message: "failed to create directory"}
	}

	for _, file := range files {
		fmt.Println("uploading")
		uploadedFile, err := file.Open()
		if err != nil {
			return utils.CustomError{Message: "problem opening file open"}
		}
		defer uploadedFile.Close()
		f, err := os.OpenFile("./media/properties/"+id_string+"/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err.Error())
			return utils.CustomError{Message: "problem saving file"}
		}
		defer f.Close()

		_, err = io.Copy(f, uploadedFile)
		if err != nil {
			return utils.CustomError{Message: "problem copying file"}
		}

		_, err = d.db.Exec(query, id, id_string+"/"+file.Filename)
		if err != nil {
			return err
		}
		fmt.Println("entered image successfully")
	}
	return nil
}

func (d *Database) GetImages(id string) ([]string, error) {
	query := `SELECT url FROM media where property_id=?`
	rows, err := d.db.Query(query, id)

	if err != nil {
		return nil, err
	}
	urls := make([]string, 0)

	for rows.Next() {
		var image string
		rows.Scan(&image)
		urls = append(urls, image)
	}

	return urls, nil
}

func (d *Database) GetProperties(filter *types.PropertyFilter, page int, pageSize int) ([]types.PropertySummary, error) {

	query := `
		SELECT p.property_id,p.title,p.description,p.price,p.address, COALESCE(m.url, '') AS url 
		FROM properties p
		LEFT JOIN media m ON p.property_id = m.property_id
		`
	filterString, parameters := filter.GenerateQueryString()
	parameters = append(parameters, pageSize, (pageSize * (page - 1)))

	if filterString != "" {
		query = query + ` WHERE ` + filterString + " LIMIT ? OFFSET ?"
	} else {
		query = query + " LIMIT ? OFFSET ?"
	}

	rows, err := d.db.Query(query, parameters...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []types.PropertySummary

	for rows.Next() {
		var property types.PropertySummary

		err := rows.Scan(&property.PropertyId, &property.Title, &property.Description, &property.Price, &property.Address, &property.Url)

		if err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}

	return properties, nil
}

func (d *Database) GetPropertyCount(queryFilter *types.PropertyFilter) (int, error) {

	filterString, params := queryFilter.GenerateQueryString()
	query := `SELECT COUNT(*) AS row_count FROM properties`
	if filterString != "" {
		query = query + " WHERE " + filterString
	}

	row, err := d.db.Query(query, params...)
	if err != nil {
		fmt.Println(err.Error())
	}
	var count int
	if row.Next() {
		err = row.Scan(&count)
	}

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (d *Database) GetPropertyDetails(propertyId int) (*types.Property, error) {
	query := `SELECT title,description,property_type,address,city,size,bedrooms,bathrooms,year_built,price,user_id FROM properties WHERE property_id=?`
	var property types.Property

	err := d.db.QueryRow(query, propertyId).Scan(&property.Title, &property.Description, &property.PropertyType, &property.Address, &property.City, &property.Size, &property.Bedrooms, &property.Bathrooms, &property.YearBuilt, &property.Price, &property.UserID)

	if err != nil {
		return nil, err
	}

	return &property, nil
}
