CREATE TABLE IF NOT EXISTS Races (
       ID integer PRIMARY KEY,
       CreatedAT datetime default current_timestamp,
       Name varchar(255) NOT NULL,
       Location varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Clubs (
       ID integer PRIMARY KEY,
       CreatedAT datetime default current_timestamp,
       Name varchar(255) NOT NULL,
       CONSTRAINT uc_clubName UNIQUE (Name)
);

CREATE TABLE IF NOT EXISTS RaceClubAssociations (
       ID integer PRIMARY KEY,
       RaceID integer,
       ClubID integer
);
