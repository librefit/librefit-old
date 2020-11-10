BEGIN TRANSACTION;

INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now') || ' ' || '09:10:00+00:00',800,1);
INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now') || ' ' || '12:34:00+00:00',310,1);
INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now') || ' ' || '13:54:00+00:00',210,1);
INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now') || ' ' || '16:01:00+00:00',270,1);
INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now') || ' ' || '20:10:00+00:00',230,1);



-- Adding some weights
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-64 days') || ' ' || '00:00:00+00:00', 89+6.9,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-55 days') || ' ' || '00:00:00+00:00', 89+7,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-48 days') || ' ' || '00:00:00+00:00', 89+7.8,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-43 days') || ' ' || '00:00:00+00:00', 89+7.1,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-37 days') || ' ' || '00:00:00+00:00', 89+6.2,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-33 days') || ' ' || '00:00:00+00:00', 89+6.2,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-26 days') || ' ' || '00:00:00+00:00', 89+5.3,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-20 days') || ' ' || '00:00:00+00:00', 89+3.5,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-14 days') || ' ' || '00:00:00+00:00', 89+2.5,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now', '-6 days') || ' ' || '00:00:00+00:00', 89+1,1);
INSERT INTO "measurements" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'weight',date('now') || ' ' || '00:00:00+00:00', 89,1);

COMMIT;
