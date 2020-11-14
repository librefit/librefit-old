#!/bin/bash
echo "BEGIN TRANSACTION;"

echo 
for i in $(seq 365 -1 1); do 
  echo "INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now', '-$i days') || ' ' || '09:10:00+00:00',ABS(RANDOM()) % (850 - 700) + 700,1);"
  echo "INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now', '-$i days') || ' ' || '12:34:00+00:00',ABS(RANDOM()) % (300 - 180) + 180,1);"
  echo "INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now', '-$i days') || ' ' || '13:54:00+00:00',ABS(RANDOM()) % (300 - 180) + 180,1);"
  echo "INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now', '-$i days') || ' ' || '16:01:00+00:00',ABS(RANDOM()) % (300 - 180) + 180,1);"
  echo "INSERT INTO "fluids" ("created_at","updated_at","deleted_at","type","date","value","user_id") VALUES (datetime('now'),datetime('now'),NULL,'water',date('now', '-$i days') || ' ' || '20:10:00+00:00',ABS(RANDOM()) % (300 - 180) + 180,1);"
done

echo "COMMIT;"
