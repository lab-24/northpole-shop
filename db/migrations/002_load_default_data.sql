-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO device_types (id, name, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130003', 'Tree Baubles', 1734481935);
INSERT INTO device_types (id, name, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130004', 'Tinsels', 1734481945);
INSERT INTO device_types (id, name, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130005', 'Treetop Stars', 1734481955);
INSERT INTO device_types (id, name, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130006', 'Sleigh', 1734481965);

INSERT INTO locations (id, name, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130007', 'Shop', 1734481800);
INSERT INTO locations (id, name, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130008', 'Garage', 1734481805);

INSERT INTO devices (id, device_type_id, location_id, name, serial_number, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130008', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130003', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130007', 'Santa Bauble', 'B-001', 1734481000);
INSERT INTO devices (id, device_type_id, location_id, name, serial_number, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130009', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130004', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130007', 'Red Tinsel', 'T-001', 1734481005);
INSERT INTO devices (id, device_type_id, location_id, name, serial_number, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac13000a', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130005', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130007', 'Golden Star', 'S-001', 1734481010);
INSERT INTO devices (id, device_type_id, location_id, name, serial_number, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac13000b', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130005', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130007', 'Iridescent White Treetop Star', 'S-002', 1734481010);
INSERT INTO devices (id, device_type_id, location_id, name, serial_number, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130011', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130006', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130008', 'Skids', 'SK-001', 1734481015);
INSERT INTO devices (id, device_type_id, location_id, name, serial_number, created_time)
VALUES ('e7f1f3c0-0b6b-11ec-82a8-0242ac130012', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130006', 'e7f1f3c0-0b6b-11ec-82a8-0242ac130008', 'BELL', 'SK-002', 1734481015);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM devices;
DELETE FROM locations;
DELETE FROM device_types;
