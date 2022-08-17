CREATE TABLE announcement_photos
(
    id              serial PRIMARY KEY,
    announcement_id bigint       NOT NULL,
    url             varchar(255) NOT NULL,


    CONSTRAINT fk_an_pt
        FOREIGN KEY (announcement_id)
            REFERENCES announcements (id)
);