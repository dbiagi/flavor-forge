CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

GRANT ALL PRIVILEGES ON DATABASE gororoba TO gororoba;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO gororoba;

CREATE TABLE IF NOT EXISTS recipe
(
    id          uuid      default gen_random_uuid(),
    title       TEXT    NOT NULL,
    description TEXT    NOT NULL,
    image       TEXT    NOT NULL,
    servings    INTEGER NOT NULL,
    prep_time   INTEGER NOT NULL,
    slug        TEXT    NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (id)
);
CREATE INDEX idx_recipe_slug ON recipe (slug);

CREATE TABLE IF NOT EXISTS ingredient
(
    id        SERIAL,
    name      TEXT        NOT NULL,
    image     TEXT,
    quantity  INTEGER     NOT NULL,
    unit      VARCHAR(20) NOT NULL,
    recipe_id uuid        NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (recipe_id) REFERENCES recipe (id)
);
CREATE INDEX idx_ingredient_recipe_id ON ingredient (recipe_id);

CREATE TABLE IF NOT EXISTS rating
(
    id        SERIAL,
    rating    INTEGER NOT NULL,
    recipe_id uuid    NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (recipe_id) REFERENCES recipe (id)
);
CREATE INDEX idx_rating_recipe_id ON rating (recipe_id);

INSERT INTO recipe (title, description, image, servings, prep_time, slug)
VALUES ('Chocolate Cake', 'Rich and creamy chocolate cake', 'https://example.com/chocolate-cake.jpg', 8, 70,
        'chocolate-cake'),
       ('Vanilla Ice Cream', 'Smooth and creamy vanilla ice cream', 'https://example.com/vanilla-ice-cream.jpg', 4, 30,
        'vanilla-ice-cream'),
       ('Strawberry Shortcake', 'Sweet and tangy strawberry dessert', 'https://example.com/strawberry-shortcake.jpg', 6,
        45, 'strawberry-shortcake'),
       ('Lemon Tart', 'A zesty lemon tart with a buttery crust', 'https://example.com/lemon-tart.jpg', 8, 50,
        'lemon-tart'),
       ('Apple Pie', 'Classic apple pie with cinnamon', 'https://example.com/apple-pie.jpg', 8, 60, 'apple-pie'),
       ('Cheesecake', 'Creamy cheesecake with a graham cracker crust', 'https://example.com/cheesecake.jpg', 10, 80,
        'cheesecake'),
       ('Brownies', 'Rich and fudgy chocolate brownies', 'https://example.com/brownies.jpg', 12, 40, 'brownies'),
       ('Banana Bread', 'Moist and flavorful banana bread', 'https://example.com/banana-bread.jpg', 10, 50,
        'banana-bread'),
       ('Panna Cotta', 'Italian panna cotta with a berry sauce', 'https://example.com/panna-cotta.jpg', 6, 30,
        'panna-cotta'),
       ('Tiramisu', 'Classic Italian tiramisu with espresso', 'https://example.com/tiramisu.jpg', 8, 30, 'tiramisu');
