-- Create users table
CREATE TABLE IF NOT EXISTS users (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE
);

-- Create countries table
CREATE TABLE IF NOT EXISTS countries (
  country_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE
);

-- Create parishes table

CREATE TABLE IF NOT EXISTS parishes (
  country_id INT,
  parish_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  UNIQUE (country_id, name),
  CONSTRAINT unique_parish_country UNIQUE (parish_id, country_id),
  FOREIGN KEY (country_id) REFERENCES countries(country_id) ON DELETE CASCADE
);


-- Create property_types table
CREATE TABLE IF NOT EXISTS property_types (
  property_type_id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE
);

-- Create properties table
CREATE TABLE IF NOT EXISTS properties (
  property_id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  property_type_id INT NOT NULL,
  address VARCHAR(255) NOT NULL,
  country_id INT NOT NULL,
  parish_id INT NOT NULL,
  neighborhood VARCHAR(255),
  size INTEGER,
  bedrooms INTEGER,
  bathrooms INTEGER,
  year_built INTEGER,
  flooring_type VARCHAR(50),
  price DECIMAL(10, 2),
  currency VARCHAR(3),
  payment_terms VARCHAR(255),
  contact_name VARCHAR(255),   
  contact_email VARCHAR(255),  
  contact_phone VARCHAR(20), 
  availability BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
  FOREIGN KEY (country_id) REFERENCES countries(country_id) ON DELETE CASCADE,
  FOREIGN KEY (parish_id, country_id) REFERENCES parishes(parish_id, country_id) ON DELETE CASCADE,
  FOREIGN KEY (property_type_id) REFERENCES property_types(property_type_id) ON DELETE CASCADE
);

-- Create media table
CREATE TABLE IF NOT EXISTS media (
  media_id SERIAL PRIMARY KEY,
  property_id INT REFERENCES properties(property_id) ON DELETE CASCADE,
  url VARCHAR(255) NOT NULL
);

-- Create amenities table
CREATE TABLE IF NOT EXISTS amenities (
  amenity_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- Create features table
CREATE TABLE IF NOT EXISTS features (
  feature_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- Create property_amenities junction table
CREATE TABLE IF NOT EXISTS property_amenities (
  property_id INT REFERENCES properties(property_id) ON DELETE CASCADE,
  amenity_id INT REFERENCES amenities(amenity_id) ON DELETE CASCADE,
  PRIMARY KEY (property_id, amenity_id)
);

-- Create property_features junction table
CREATE TABLE IF NOT EXISTS property_features (
  property_id INT REFERENCES properties(property_id) ON DELETE CASCADE,
  feature_id INT REFERENCES features(feature_id) ON DELETE CASCADE,
  PRIMARY KEY (property_id, feature_id)
);

CREATE TABLE IF NOT EXISTS sessions (
  id SERIAL PRIMARY KEY,
  session_id VARCHAR(64) NOT NULL,
  data BYTEA,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);



INSERT INTO countries (name) VALUES ('Barbados'),('St.Lucia'),('Grenada');

INSERT INTO parishes (name,country_id) VALUES 
  ('St. James',1),
  ('St.Micahel',1),
  ('St.Peter',1),
  ('St.Joseph',1),
  ('St.George',1),
  ('St.Phillip',1),
  ('St.Thomas',1),
  ('St.Lucy',1),
  ('Christ Church',1),
  ('St.Andrew',1),
  ('St.John',1)
  ;

INSERT INTO parishes (name, country_id) VALUES 
  ('Anse la Raye', 2),
  ('Canaries', 2),
  ('Castries', 2),
  ('Choiseul', 2),
  ('Dennery', 2),
  ('Gros Islet', 2),
  ('Laborie', 2),
  ('Micoud', 2),
  ('Praslin', 2),
  ('Soufriere', 2),
  ('Vieux Fort', 2);

-- Insert parishes for Grenada (country_id = 3)
INSERT INTO parishes (name, country_id) VALUES 
  ('Saint Andrew', 3),
  ('Saint David', 3),
  ('Saint George', 3),
  ('Saint John', 3),
  ('Saint Mark', 3),
  ('Saint Patrick', 3);


INSERT INTO property_types (name) VALUES
  ('Townhouse'),
  ('Beachfront Villas'),
  ('Condominiums'),
  ('Apartment'),
  ('House')
;

