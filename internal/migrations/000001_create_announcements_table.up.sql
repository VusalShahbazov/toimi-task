CREATE TABLE announcements
(
    id          serial PRIMARY KEY,
    title       varchar(255) NOT NULL,
    description text         NOT NULL,
    price       numeric      NOT NULL,
    created_at  timestamp    NOT NULL DEFAULT now()
);