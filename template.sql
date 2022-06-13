DROP TABLE IF EXISTS category CASCADE;
DROP TABLE IF EXISTS item CASCADE;


CREATE TABLE public.category (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE public.item (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    category_id UUID NOT NULL,
    CONSTRAINT category_fk FOREIGN KEY (category_id) REFERENCES public.category(id)
);

INSERT INTO category (id, name) VALUES ('707f69e0-edac-4c3e-bb7f-918d3f190e19', 'Народ');
INSERT INTO category (id, name) VALUES ('1ad0596d-e43d-4093-a7fe-a6c1074f6219', 'Джоан Роулинг');
INSERT INTO category (id, name) VALUES ('62af3986-0963-465c-8a86-dd23ac240523', 'Джек Лондон');

INSERT INTO item (name, category_id) VALUES ('колобок', '707f69e0-edac-4c3e-bb7f-918d3f190e19');
INSERT INTO item (name, category_id) VALUES ('гарри поттер', '1ad0596d-e43d-4093-a7fe-a6c1074f6219');
INSERT INTO item (name, category_id) VALUES ('брилианты', '62af3986-0963-465c-8a86-dd23ac240523');