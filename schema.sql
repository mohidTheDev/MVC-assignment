CREATE TYPE BuildingType AS ENUM ('Defenses', 'Resource_Generators', 'Resource_Storage', 'Special'); 
CREATE TYPE ResourceType AS ENUM ('Gold', 'Elixir'); 
CREATE TYPE MovementType AS ENUM ('Ground', 'Air'); 
CREATE TYPE AttackType AS ENUM ('Melee', 'Ranged'); 
CREATE TYPE ProjectileType AS ENUM ('Arrow', 'Fireball'); 
CREATE TYPE BattleResult AS ENUM ('Victory', 'Defeat'); 
CREATE TYPE BattleAction AS ENUM ('Killed', 'Destroyed'); 
CREATE TYPE TargetingType AS ENUM ('Single', 'Area'); 
CREATE TYPE SpecialType AS ENUM ('Town Hall', 'Barracks', "Army Camp"); 

-- BASE TABLES

CREATE TABLE Players (
    Username VARCHAR(255) PRIMARY KEY,
    PWD_Hash VARCHAR(255) NOT NULL,
    Account_Creation_Date DATE NOT NULL DEFAULT CURRENT_DATE,
    Last_Login_Date DATE
);

CREATE TABLE Buildings (
    Name VARCHAR(50) PRIMARY KEY,
    Type BuildingType NOT NULL, 
    Build_Resource ResourceType NOT NULL, 
    SizeX INT NOT NULL,
    SizeY INT NOT NULL
);

CREATE TABLE Characters (
    Name VARCHAR(50) PRIMARY KEY,
    Housing_Space INT NOT NULL,
    Walkspeed FLOAT NOT NULL,
    Movement_Type MovementType NOT NULL, 
    Attack_Type AttackType NOT NULL, 
    Projectile_Type ProjectileType 
);

-- PLAYER-DEPENDENT TABLES

CREATE TABLE Stats (
    Username VARCHAR(255) PRIMARY KEY REFERENCES Players(Username) ON DELETE CASCADE,
    Attacks_Won INT DEFAULT 0,
    Attacks_Defended INT DEFAULT 0,
    Gold_Looted BIGINT DEFAULT 0,
    Elixir_Looted BIGINT DEFAULT 0,
    Trophies INT DEFAULT 0
);

CREATE TABLE Village (
    Player_Username VARCHAR(255) PRIMARY KEY REFERENCES Players(Username) ON DELETE CASCADE,
    Town_Hall_Level INT DEFAULT 1,
    Gold BIGINT DEFAULT 0,
    Elixir BIGINT DEFAULT 0
);

CREATE TABLE Army (
    Player_Username VARCHAR(255) REFERENCES Players(Username) ON DELETE CASCADE,
    Name VARCHAR(50) REFERENCES Characters(Name),
    Level INT NOT NULL,
    Quantity INT DEFAULT 0,
    PRIMARY KEY (Player_Username, Name, Level)
);

-- BATTLE TABLES

CREATE TABLE Battle_Log (
    Battle_ID BIGSERIAL PRIMARY KEY,
    Attacker VARCHAR(255) REFERENCES Players(Username) ON DELETE SET NULL,
    Defender VARCHAR(255) REFERENCES Players(Username) ON DELETE SET NULL,
    Result BattleResult 
);

CREATE TABLE Battle_Events (
    Event_ID BIGSERIAL PRIMARY KEY,
    Battle_ID BIGINT REFERENCES Battle_Log(Battle_ID) ON DELETE CASCADE,
    Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Subject VARCHAR(255),
    Object VARCHAR(255),
    Action BattleAction 
);

-- LEVEL-SPECIFIC TABLES

CREATE TABLE Buildings_Level_Specific (
    Name VARCHAR(50) REFERENCES Buildings(Name) ON DELETE CASCADE,
    Level INT NOT NULL,
    Min_Town_Hall_Level INT NOT NULL,
    Build_Cost BIGINT NOT NULL,
    HP INT NOT NULL,
    Upgrade_Time BIGINT NOT NULL, 
    PRIMARY KEY (Name, Level)
);

CREATE TABLE Character_Level_Dependent (
    Name VARCHAR(50) REFERENCES Characters(Name) ON DELETE CASCADE,
    Level INT NOT NULL,
    Upgrade_Cost BIGINT NOT NULL,
    Min_Barracks_Level INT NOT NULL,
    HP INT NOT NULL,
    Damage INT NOT NULL,
    Projectile_Range INT,
    PRIMARY KEY (Name, Level)
);

-- BUILDING SPECIFIC TABLES

CREATE TABLE Defenses (
    Name VARCHAR(50),
    Level INT,
    Projectile_Type ProjectileType, 
    Range INT NOT NULL,
    Targeting_Type TargetingType, 
    Projectile_AOE_Range INT,
    PRIMARY KEY (Name, Level),
    FOREIGN KEY (Name, Level) REFERENCES Buildings_Level_Specific(Name, Level) ON DELETE CASCADE
);

CREATE TABLE Resource_Generators (
    Name VARCHAR(50),
    Level INT,
    Generation_Rate INT NOT NULL, 
    Resource_Type ResourceType NOT NULL, 
    PRIMARY KEY (Name, Level),
    FOREIGN KEY (Name, Level) REFERENCES Buildings_Level_Specific(Name, Level) ON DELETE CASCADE
);

CREATE TABLE Resource_Storage (
    Name VARCHAR(50),
    Level INT,
    Storage_Capacity BIGINT NOT NULL,
    Resource_Type ResourceType NOT NULL, 
    PRIMARY KEY (Name, Level),
    FOREIGN KEY (Name, Level) REFERENCES Buildings_Level_Specific(Name, Level) ON DELETE CASCADE
);

CREATE TABLE Special (
    Name VARCHAR(50),
    Level INT,
    Type SpecialType, 
    PRIMARY KEY (Name, Level),
    FOREIGN KEY (Name, Level) REFERENCES Buildings_Level_Specific(Name, Level) ON DELETE CASCADE
);

-- Instances

CREATE TABLE Placed_Buildings (
    Placement_ID BIGSERIAL PRIMARY KEY,
    Player_Username VARCHAR(255) REFERENCES Players(Username) ON DELETE CASCADE,
    PositionX INT NOT NULL,
    PositionY INT NOT NULL,
    Name VARCHAR(50) NOT NULL,
    Level INT NOT NULL,
    Upgrade_Start_Time TIMESTAMP,
    FOREIGN KEY (Name, Level) REFERENCES Buildings_Level_Specific(Name, Level)
);