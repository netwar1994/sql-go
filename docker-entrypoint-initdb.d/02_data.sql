INSERT INTO clients (login, password, full_name, passport, birthday, status)
VALUES ('john', 'secret', 'John Doe', 'AN000001', '1998-02-03', 'ACTIVE'),
       ('rick', 'secret', 'Rick Sanchez', 'AN000002', '1958-01-08', 'INACTIVE'),
       ('morty', 'secret', 'Morty Smith', 'AN000003', '2005-10-20', 'ACTIVE');

INSERT INTO cards (number, balance, issuer, holder, owner_id, status)
VALUES ('4444444444444444', 1000000, 'MasterCard', 'john', 1, 'ACTIVE'),
       ('4444444444444445', 1000000, 'Visa', 'john', 1, 'ACTIVE'),
       ('7777777777777777', 1000000, 'Visa', 'rick', 2, 'INACTIVE'),
       ('1111111111111111', 1000000, 'MIR', 'morty', 3, 'ACTIVE');

INSERT INTO icons (url)
VALUES ('https://google.com'),
       ('https://google.com'),
       ('https://google.com');

INSERT INTO mcc (mcc_code, description)
VALUES ('5411', 'Grocery Stores, Supermarkets'),
       ('5812', 'Eating Places and Restaurants'),
       ('2741', 'Miscellaneous Publishing and Printing'),
       ('4119', 'Ambulance Services');

INSERT INTO transactions (card_id, sum, mcc_id, description, icon_id, status)
VALUES (3, 350, 1, 'Food', 1, 'DONE'),
       (1, 240, 2, 'Lunch', 2, 'DONE'),
       (2, 350, 3, 'Printing', 3, 'DONE'),
       (3, 350, 1, 'Food', 3, 'PENDING');
