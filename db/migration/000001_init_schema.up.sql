CREATE TABLE "pokemons" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "type1" varchar NOT NULL,
  "type2" varchar NOT NULL,
  "hp" integer NOT NULL,
  "attack" integer NOT NULL,
  "defense" integer NOT NULL,
  "sp_atk" integer NOT NULL,
  "sp_def" integer NOT NULL,
  "speed" integer NOT NULL,
  "generation" integer NOT NULL,
  "legendary" boolean NOT NULL, 
  "created_at" TIMESTAMPtz NOT NULL DEFAULT 'now()');