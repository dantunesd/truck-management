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

INSERT INTO trucks (id, license_plate, eld_id, carrier_id, type, size, color, make, model, year, created_at, updated_at) VALUES ('1', 'INSERTT', 'INSERTED ELD', 'INSERTT MY CARRIER', 'REEFERR', '23', 'blue', 'Maker', 'Model', '1900', '2021-07-04 21:34:26', '2021-07-04 21:34:26');
INSERT INTO trucks (id, license_plate, eld_id, carrier_id, type, size, color, make, model, year, created_at, updated_at) VALUES ('2', 'TODELET', 'TO DELETE ELD', 'TO DELETE MY CARRIER', 'REEFERR', '23', 'blue', 'Maker', 'Model', '1900', '2021-07-04 21:34:26', '2021-07-04 21:34:26');

CREATE TABLE IF NOT EXISTS `truck_management`.`locations` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `truck_id` INT NOT NULL,
  `eld_id` VARCHAR(20) NOT NULL,
  `engine_state` VARCHAR(3) NOT NULL,
  `current_speed` INT NOT NULL,
  `latitude` INT NOT NULL,
  `longitude` INT NOT NULL,
  `engine_hours` INT NOT NULL,
  `odometer` INT NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_locations_1_idx` (`truck_id`),
  INDEX `fk_locations_2_idx` (`eld_id`),
  CONSTRAINT `fk_locations_1`
    FOREIGN KEY (`truck_id`)
    REFERENCES `truck_management`.`trucks` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_locations_2`
    FOREIGN KEY (`eld_id`)
    REFERENCES `truck_management`.`trucks` (`eld_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);