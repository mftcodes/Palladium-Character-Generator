/* REMOVE CONSTRAINTS */
ALTER TABLE palladium.RaceAttributes
DROP FOREIGN KEY FK__RaceAttributes__Race;

ALTER TABLE palladium.RaceAttributes 
DROP FOREIGN KEY RaceAttributes_ibfk_1;

ALTER TABLE palladium.Character
DROP FOREIGN KEY FK__Character__race;

ALTER TABLE palladium.Character
DROP FOREIGN KEY Character_ibfk_1;

DROP TABLE IF EXISTS `Race`;
CREATE TABLE IF NOT EXISTS `Race` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Desc` VARCHAR(25) NOT NULL DEFAULT 'oops',
  PRIMARY KEY (id)
);


/* BUILD TABLES */
DROP TABLE IF EXISTS `RaceAttributes`;
CREATE TABLE IF NOT EXISTS `RaceAttributes` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `RaceId` INT NOT NULL,
  `IQ` SMALLINT NOT NULL DEFAULT 0,
  `IQBonus` SMALLINT NOT NULL DEFAULT 0,
  `ME` SMALLINT NOT NULL DEFAULT 0,
  `MEBonus` SMALLINT NOT NULL DEFAULT 0,
  `MA` SMALLINT NOT NULL DEFAULT 0,
  `MABonus` SMALLINT NOT NULL DEFAULT 0,
  `PS` SMALLINT NOT NULL DEFAULT 0,
  `PSBonus` SMALLINT NOT NULL DEFAULT 0,
  `PP` SMALLINT NOT NULL DEFAULT 0,
  `PPBonus` SMALLINT NOT NULL DEFAULT 0,
  `PE` SMALLINT NOT NULL DEFAULT 0,
  `PEBonus` SMALLINT NOT NULL DEFAULT 0,
  `PB` SMALLINT NOT NULL DEFAULT 0,
  `PBBonus` SMALLINT NOT NULL DEFAULT 0,
  `Spd` SMALLINT NOT NULL DEFAULT 0,
  `SpdBonus` SMALLINT NOT NULL DEFAULT 0,
  `PPE` SMALLINT NOT NULL DEFAULT 0,
  `PPEBonus` SMALLINT NOT NULL DEFAULT 0,
  `HF`  SMALLINT NOT NULL DEFAULT 0,
  `Alignment` VARCHAR(255) DEFAULT "any",
  `SpdDig` SMALLINT NOT NULL DEFAULT 0,
  `SpdDigBonus` SMALLINT NOT NULL DEFAULT 0,
  CONSTRAINT PK_raceAttributes PRIMARY KEY (Id, RaceId),
  CONSTRAINT FK__RaceAttributes__Race FOREIGN KEY (RaceId)
  REFERENCES Race(Id)
);

DROP TABLE IF EXISTS `Character`;
CREATE TABLE IF NOT EXISTS `Character` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(25) NOT NULL,
  `RaceId` INT NOT NULL,
  `Lvl` SMALLINT NOT NULL DEFAULT 1,
  `IQ` SMALLINT NOT NULL DEFAULT 0,
  `ME` SMALLINT NOT NULL DEFAULT 0,
  `MA` SMALLINT NOT NULL DEFAULT 0,
  `PS` SMALLINT NOT NULL DEFAULT 0,
  `PP` SMALLINT NOT NULL DEFAULT 0,
  `PE` SMALLINT NOT NULL DEFAULT 0,
  `PB` SMALLINT NOT NULL DEFAULT 0,
  `Spd` SMALLINT NOT NULL DEFAULT 0,
  `PPE` SMALLINT NOT NULL DEFAULT 0,
  `HF`  SMALLINT NOT NULL DEFAULT 0,
  `SpdDig` SMALLINT NOT NULL DEFAULT 0,
  PRIMARY KEY (Id),
  CONSTRAINT FK__Character__Race FOREIGN KEY (RaceId)
  REFERENCES Race(Id)
);

DROP TABLE IF EXISTS `OCCType`;
CREATE TABLE IF NOT EXISTS `OCCType` (
  `Id`    INT NOT NULL AUTO_INCREMENT,
  `Desc`  VARCHAR(50) NOT NULL,
  PRIMARY KEY (Id)
);

DROP TABLE IF EXISTS `OCC`;
CREATE TABLE IF NOT EXISTS `OCC` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Desc` VARCHAR(50) NOT NULL,
  `OCCTypeId` INT NOT NULL,
  PRIMARY KEY (Id),
  CONSTRAINT FK__OCC__OCCType FOREIGN KEY (OCCTypeId)
  REFERENCES OCCType(Id)
);

DROP TABLE IF EXISTS `SkillCategory`;
CREATE TABLE IF NOT EXISTS `SkillCategory` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Desc` VARCHAR(50) NOT NULL,
  PRIMARY KEY (Id)
);

DROP TABLE IF EXISTS `Skill`;
CREATE TABLE IF NOT EXISTS `Skill` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Desc` VARCHAR(50) NOT NULL,
  `SkillCategoryId` INT NOT NULL,
  PRIMARY KEY (Id),
  CONSTRAINT FK__Skill__SkillCategory FOREIGN KEY (SkillCategoryId)
  REFERENCES SkillCategory(Id)
);

DROP TABLE IF EXISTS `NaturalAbility`;
CREATE TABLE IF NOT EXISTS `NaturalAbility` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Desc` VARCHAR(50) NOT NULL,
  PRIMARY KEY (Id)
);

DROP TABLE IF EXISTS `Race_NaturalAbility`;
CREATE TABLE IF NOT EXISTS `Race_NaturalAbility` (
  `NaturalAbilityId`  INT NOT NULL,
  `RaceId`            INT NOT NULL,
  `BonusInitial`      INT NOT NULL DEFAULT 0,
  `BonusPerLevel`     INT NOT NULL DEFAULT 0,
  `Value`             INT NOT NULL DEFAULT 0,
  `Measurement`       VARCHAR(10) NOT NULL DEFAULT 'NA',
  `Note`              VARCHAR(100),
  CONSTRAINT FK__Race_NaturalAbility__NaturalAbility FOREIGN KEY (NaturalAbilityId)
  REFERENCES NaturalAbility(Id),
  CONSTRAINT FK__Race_NaturalAbility__Race FOREIGN KEY (RaceId)
  REFERENCES Race(Id)
);

DROP TABLE IF EXISTS `Race_OCC`;
CREATE TABLE IF NOT EXISTS `Race_OCC` (
  `OccId`	INT NOT NULL,
  `RaceId`	INT NOT NULL,
  CONSTRAINT FK__Race_OCC__OCC FOREIGN KEY (OccId)
  REFERENCES OCC(Id),
  CONSTRAINT FK__Race_OCC__Race FOREIGN KEY (RaceId)
  REFERENCES Race(Id)
)


/* FILL TABLES */
  -- RACE

INSERT INTO palladium.Race (`Desc`)
VALUES ('Human'), ('Elf'), ('Dwarf'), ('Gnome'), ('Troglodyte'), ('Kobold'), ('Goblin'), 
	('Hob-Goblin'), ('Orc'), ('Ogre'), ('Troll'), ('Changeling'), ('Wolfen'), ('Coyle');

INSERT INTO palladium.RaceAttributes (`RaceId`, `IQ`, `IQBonus`, `ME`, `MEBonus`, `MA`, `MABonus`, `PS`, `PSBonus`, `PP`, `PPBonus`, `PE`, `PEBonus`,
	`PB`, `PBBonus`, `Spd`, `SpdBonus`, `PPE`, `PPEBonus`,`HF`,  `Alignment`, `SpdDig`, `SpdDigBonus`)
VALUES 	(1,3,0,3,0,3,0,3,0,3,0,3,0,3,0,3,0,2,0,0,'Any, usually lean toward good and selfish',0,0),
        (2,3,1,3,0,2,0,3,0,4,0,3,0,5,0,3,0,2,0,0,'Any, usually lean toward good and selfish',0,0),
        (3,3,0,3,0,2,0,4,6,3,0,4,0,2,2,2,0,2,0,0,'Any, usually lean toward good and selfish',1,0),
        (4,3,0,1,6,3,4,1,4,4,0,3,6,4,0,2,0,2,0,0,'Any, but most tend to be good or selfish; an evil gnome is a rarity',1,0),
        (5,2,0,2,0,3,0,4,4,3,6,3,0,2,0,6,0,2,0,0,'Typically anarchist or evil; but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        (6,3,0,2,0,3,0,3,3,3,0,3,0,1,6,3,0,4,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        (7,2,0,3,0,2,0,3,0,3,6,3,0,2,0,3,0,6,0,0,'Typically anarchist or evil; but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        (8,2,0,3,6,2,0,3,0,3,0,3,0,2,0,3,0,4,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        (9,2,0,2,0,3,0,3,8,3,0,3,2,2,0,3,0,2,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',0,0),
        (10,3,0,3,0,2,0,4,4,3,0,3,6,2,0,3,0,3,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',0,0),
        (11,3,0,2,0,2,0,4,10,4,0,3,6,1,4,2,0,3,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',0,0),
        (12,3,0,4,6,4,0,3,0,3,0,2,0,2,0,2,0,5,0,0,'Any',0,0),
        (13,3,0,3,0,2,0,4,1,3,0,3,0,3,0,4,0,3,0,0,'Any, but tend toward principled and aberrant, both alignments with a strong personal code of honor',0,0),
        (14,3,0,3,0,2,0,3,1,4,1,3,0,3,0,3,0,3,0,0,'Any, but tend toward anarchist and miscreant; the antithesis of the noble Wolfen',0,0);

  --OCC
INSERT INTO palladium.OCCType (`Desc`)
VALUES ('Clergy'), ('Men of Arms'), ('Optional'), ('Practitioners of Magic'), ('Psychics');

INSERT INTO palladium.OCC (`Desc`, `OCCTypeId`)
VALUES  ('Druid', 1), ('Monk', 1), ('Priest of Light', 1), ('Priest of Darkness', 1),
        ('Assassin', 2), ('Knight', 2), ('Long Bowman', 2), ('Mercenary Warrior', 2), ('Palladin', 2), ('Ranger', 2), ('Soldier', 2), ('Thief', 2),
        ('Merchant', 3), ('Noble', 3), ('Scholar', 3), ('Squire', 3), ('Vagabond', 3), ('Peasant', 3), ('Farmer', 3),
        ('Diabolist', 4), ('Summoner', 4), ('Warlock', 4), ('Witch', 4), ('Wizard', 4);,
        ('Mind Mage', 5), ('Psi-Healer', 5), ('Psi-Mystic', 5), ('Psychic Sensitive', 5);

  --SKILLS
INSERT INTO palladium.SkillCategory (`Desc`)
VALUES  ('Communications & Performing Arts'), ('Domestic'), ('Espionage'), ('Horsemanship'), ('Medical'), ('Military'),
        ('Physical'), ('Rogue/Thief'), ('Science'), ('Scholar, Noble & Technical'), ('Weapon Proficiencies'), ('Wilderness');

INSERT INTO palladium.Skill (`Desc`, `SkillCategoryId`)
VALUES  ('Cryptography',1),('Dance',1),('Language',1),('Literacy',1),('Mime',1),('Play Musical Instrument',1),('Public Speaking',1),
        ('Sign Language',1),('Sing',1),('Writing',1),
        ('Cook',2),('Dance',2),('Fishing',2),('Play Musical Instrument',2),('Sew',2),('Sing',2),
        ('Detect Ambush',3),('Detect Concealment & Traps',3),('Disquise',3),('Escape Artist',3),('Forgery',3),('Imitate Voices & Impersonation',3),
        ('Intelligence',3),('Pick Locks',3),('Pick Pockets',3),('Sniper',3),('Track Humanoids',3),
        ('General',4),('Knight',4),('Palladin',4),('Exotic Animals',4),
        ('Animal Husbandry',5),('Biology',5),('Brewing',5),('First Aid',5),('Holistic Medicine',5),('Surgeon/Medical Doctor',5),
        ('Camouflage',6),('Falconry',6),('Field Armorer',6),('Heraldry',6),('Interrogation Techniques',6),('Military Etiquitte',6),
        ('Recognize Weapon Quality',6),('Surveillance',6),
        ('Hand to Hand: Basic',7),('Hand to Hand: Expert',7),('Hand to Hand: Martial Arts',7),('Hand to Hand: Assassin',7),('Acrobatics',7),
        ('Athletics',7),('Body Building & Weight Lifting',7),('Boxing',7),('Climb/Scale Walls',7),('Forced March',7),('Gymnastics',7),
        ('Juggling',7),('Prowl',7),('Running',7),('Swimming',7),('Wrestling',7),
        ('Card Shark',8),('Concealment',8),('Locate Secret Compartments/Doors',8),('Palming',8),('Pick Locks',8),('Pick Pockets',8),('Prowl',8),
        ('Streetwise',8),('Use & Recognize Poison',8),('Ventriloquism',8),
        ('Anthropology',9),('Archaeology',9),('Astronomy & Navigation',9),('Biology',9),('Botany',9),('Mathematics: Basic',9),('Mathematics: Advanced',9),
        ('Art',10),('Breed Dogs',10),('Gemology',10),('General Repair',10),('History',10),('Language',10),('Literacy',10),('Lore: Demons & Monsters',10),
        ('Lore: Faerie Folk',10),('Lore: Geomancy & Ley Lines',10),('Lore: Magic',10),('Lore: Religion',10),('Masonry',10),('Rope Works',10),
        ('Sailing',10),('Sculpting & Whittling',10),('Writing',10),
        ('Archery',11),('Blunt',11),('Chain',11),('Forked Weapons/Trident',11),('Grappling Hook',11),('Knife',11),('Modern Weapons',11),
        ('Mouth Weapons/Blowguns',11),('Net',11),('Paired Weapons',11),('Shield',11),('Siege Weapons',11),('Spear',11),('Staff',11),('Sword',11),
        ('Targeting/Missle Weapons',11),('Throwing Weapons',11),('Whip',11),
        ('Boat Building',12),('Capentry',12),('Dowsing',12),('Identify Plants & Fruits',12),('Land Navigation',12),('Preseve Food',12),
        ('Skin & Prepare Animal Hides',12),('Track & Trap Animals',12),('Wilderness Survival',12);


  --NATURAL ABILITIES
INSERT INTO palladium.NaturalAbility (`Desc`)
VALUES  ('nightvision'),('Underground Tunneling'),('Underground Architecture'),('Underground Sense of Direction'),
        ('Underground Sense of Surface Structure Location'),('Metal Working'),('Recognize Precious Metals & Stones');

INSERT INTO palladium.Race_NaturalAbility (`NaturalAbilityId`, `RaceId`, `BonusInitial`, `BonusPerLevel`, `Value`, `Measurement`, `Note`)
VALUES  (1,2,0,0,60,'Feet',NULL),
        (1,3,0,0,90,'Feet',NULL),(2,3,40,5,0,'%',NULL),(3,3,30,5,0,'%','detection and deactivation of traps is done at half normal architecture skill level'),
        (4,3,40,5,0,'%',NULL),(5,3,30,5,0,'%','-25% if in unfamiliar area'),(6,3,10,0,0,'%','Equal to Field Armorer'),(7,3,10,0,0,'%','Same as Gemology'),
        (1,4,0,0,90,'Feet',NULL),(2,4,30,5,0,'%',NULL),(3,4,20,5,0,'%','detection and deactivation of traps is done at half normal architecture skill level'),
        (4,4,30,5,0,'%',NULL),(5,4,20,5,0,'%','-20% if in unfamiliar area'),
        (1,5,600,0,0,'Feet','day vision 30ft'),(2,5,30,5,0,'%',NULL),(3,5,20,5,0,'%','detection and deactivation of traps is done at half normal architecture skill level'),
        (4,5,40,5,0,'%',NULL),(5,5,15,5,0,'%','-20% if in unfamiliar area'),
        (1,6,0,0,400,'Feet','day vision 40ft'),(2,6,40,5,0,'%',NULL),(3,6,30,5,0,'%','detection and deactivation of traps is done at half normal architecture skill level'),
        (4,6,40,5,0,'%',NULL),(5,6,30,5,0,'%','-25% if in unfamiliar area'),(6,6,10,0,0,'%','Equal to Field Armorer, +10 recognize weapon quality'),
        (7,6,10,0,0,'%','Same as Gemology, art (limted to jewelry) and gems');
        -- stopped after Kobold. 


INSERT INTO palladium.Race_OCC (OccId, RaceId)
VALUES  (1,1),(2,1),(3,1),(4,1),(5,1),(6,1),(7,1),(8,1),(9,1),(10,1),(11,1),(12,1),(13,1),(14,1),(15,1),(16,1),(17,1),
        (18,1),(19,1),(20,1),(21,1),(22,1),(23,1),(24,1),(25,1),(26,1),(27,1),(28,1),
        (1,2),(2,2),(3,2),(4,2),(5,2),(6,2),(7,2),(8,2),(9,2),(10,2),(11,2),(12,2),(13,2),(14,2),(15,2),(16,2),(17,2),
        (18,2),(19,2),(20,2),(21,2),(22,2),(23,2),(24,2),(25,2),(26,2),(27,2),(28,2),
        (1,3),(2,3),(3,3),(4,3)(5,3),(6,3),(7,3),(8,3),(9,3),(10,3),(11,3),(12,3),(13,3),(14,3),(15,3),(16,3),(17,3),
        (18,3),(19,3),(25,3),(26,3),(27,3),(28,3),
        (1,4),(2,4),(3,4),(4,4),(5,4),(8,4),(10,4),(11,4),(12,4),(13,4),(14,4),(15,4),(16,4),(17,4),
        (18,4),(19,4),(20,4),(21,4),(22,4),(23,4),(24,4),
        (2,5),(5,5),(8,5),(11,5),(12,5),(17,5),
        (1,6),(2,6),(3,6),(4,6),(5,6),(8,6),(10,6),(11,6),(12,6),(13,6),(14,6),(15,6),(16,6),(17,6),
        (18,6),(19,6),(20,6),(21,6),(22,6),(23,6),(24,6),(25,6),(26,6),(27,6),(28,6);
        -- stopped after Kobold. 
        
        
/*  
  `Id`    INT NOT NULL AUTO_INCREMENT,
  `Desc`  VARCHAR(50) NOT NULL,
  `Roll`  VARCHAR(5) NOT NULL DEFAULT 'NA',
  `Bonus` INT NOT NULL DEFAULT 0,
  `Value` INT NOT NULL DEFAULT 0,
  `ValueMeasurement` VARCHAR(10) NOT NULL DEFAULT 'NA',

,('gggggg',0)
*/