CREATE DATABASE IF NOT EXISTS truck_management;

USE truck_management;

CREATE TABLE IF NOT EXISTS trucks (
    id INT NOT NULL AUTO_INCREMENT,
    license_plate VARCHAR(7) NOT NULL,
    eld_id VARCHAR(20) NOT NULL,
    carrier_id VARCHAR(20) NOT NULL,
    type VARCHAR(20) NOT NULL,
    size INT NULL,
    color VARCHAR(20) NULL,
    make VARCHAR(20) NULL,
    model VARCHAR(20) NULL,
    year INT NULL,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE INDEX license_plate_UNIQUE (license_plate),
    UNIQUE INDEX eld_id_UNIQUE (eld_id)
);
