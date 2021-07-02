db.createCollection('trucks');
db.createCollection('locations');
db.createCollection('trip-report');

db.trucks.createIndex({ licensePlate: 1 }, { unique: true });
db.trucks.createIndex({ eldId: 1 }, { unique: true });
