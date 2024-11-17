CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

GRANT ALL PRIVILEGES ON DATABASE gororoba TO gororoba;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO gororoba;

CREATE TABLE IF NOT EXISTS recipe(
    id uuid default gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    image TEXT NOT NULL,
    servings INTEGER NOT NULL,
    prep_time INTEGER NOT NULL,
    slug TEXT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS ingredient(
    id SERIAL,
    name TEXT NOT NULL,
    image TEXT,
    quantity INTEGER NOT NULL,
    unit VARCHAR(20) NOT NULL,
    recipe_id uuid NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY(recipe_id) REFERENCES recipe(id)
);
CREATE INDEX idx_ingredient_recipe_id ON ingredient (recipe_id);

CREATE TABLE IF NOT EXISTS rating(
    id SERIAL,
    rating INTEGER NOT NULL,
    recipe_id uuid NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY(recipe_id) REFERENCES recipe(id)
);
CREATE INDEX idx_rating_recipe_id ON rating (recipe_id);
