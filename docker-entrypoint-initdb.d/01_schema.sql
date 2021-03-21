CREATE TABLE clients
(
    id          BIGSERIAL   PRIMARY KEY,
    login       TEXT        NOT NULL UNIQUE,
    password    TEXT        NOT NULL,
    full_name   TEXT        NOT NULL,
    passport    TEXT        NOT NULL,
    birthday    DATE        NOT NULL,
    status      TEXT        NOT Null DEFAULT 'INACTIVE' CHECK (status IN ('INACTIVE', 'ACTIVE')),
    created     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cards
(
    id          BIGSERIAL   PRIMARY KEY,
    number      TEXT        NOT NULL,
    balance     BIGINT      NOT NULL DEFAULT 0,
    issuer      TEXT        NOT NULL CHECK (issuer IN ('Visa', 'MasterCard', 'MIR')),
    holder      TEXT        NOT NULL,
    owner_id    BIGINT      NOT NULL REFERENCES clients,
    status      TEXT        NOT NULL DEFAULT 'INACTIVE' CHECK (status IN ('INACTIVE', 'ACTIVE')),
    created     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE icons
(
    id  BIGSERIAL    PRIMARY KEY,
    url TEXT         NOT NULL
);

CREATE TABLE mcc
(
    id              BIGSERIAL   PRIMARY KEY,
    mcc_code        TEXT    NOT NULL,
    description     TEXT    NOT NULL
);

CREATE TABLE transactions
(
    id          BIGSERIAL   PRIMARY KEY,
    card_id     BIGINT      NOT NULL REFERENCES cards,
    sum         BIGINT      NOT NULL,
    mcc_id      BIGINT      NOT NULL REFERENCES mcc,
    description TEXT        NOT NULL,
    icon_id     BIGINT      NOT NULL REFERENCES icons,
    status      TEXT        NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'DONE')),
    created     TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);



