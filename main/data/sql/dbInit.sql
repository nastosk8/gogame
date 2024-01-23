USE gogame;

DROP TABLE IF EXISTS heroClasses;
CREATE TABLE heroClasses(
    id INT PRIMARY KEY AUTO_INCREMENT, 
    name VARCHAR(255), 
    damageType VARCHAR(255), 
    health INT, 
    mana INT
);
INSERT INTO heroClasses(name, damageType, health, mana) VALUES('Mage', 'Magical', 80, 100);
INSERT INTO heroClasses(name, damageType, health, mana) VALUES('Warrior', 'Psychical', 100, 0);

DROP TABLE IF EXISTS players;
CREATE TABLE players(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    class VARCHAR(255)
)
