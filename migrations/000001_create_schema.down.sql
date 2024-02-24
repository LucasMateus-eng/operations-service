BEGIN;

DROP TABLE IF EXISTS "drivers_vehicles";
DROP TABLE IF EXISTS "drivers";
DROP TABLE IF EXISTS "vehicles";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "adresses";

DROP TYPE IF EXISTS "vehicle_licensing_status";
DROP TYPE IF EXISTS "user_role";

COMMIT;