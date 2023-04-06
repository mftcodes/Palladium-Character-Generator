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
  `HP`  SMALLINT NOT NULL DEFAULT 0,
  `PPE` SMALLINT NOT NULL DEFAULT 0,
  `HF`  SMALLINT NOT NULL DEFAULT 0,
  `SpdDig` SMALLINT NOT NULL DEFAULT 0,
  `OccId` INT NOT NULL DEFAULT 1,
  PRIMARY KEY (Id),
  CONSTRAINT FK__Character__Race FOREIGN KEY (RaceId)
  REFERENCES Race(Id),
  CONSTRAINT FK__Character__OCC FOREIGN KEY (OccId)
  REFERENCES OCC(Id)
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
VALUES ('Human'),('Elf'),('Dwarf'),('Gnome'),('Troglodyte'),('Kobold'),('Goblin'), 
	('Hob-Goblin'),('Orc'),('Ogre'),('Troll'),('Changeling'),('Wolfen'),('Coyle');

/* Hold on to these if needed to add to other queries. */
SET @Human = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human');
SET @Elf = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf');
SET @Dwarf = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf');
SET @Gnome = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome');
SET @Troglodyte = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte');
SET @Kobold = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold');
SET @Goblin = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Goblin');
SET @Hob-Goblin = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Hob-Goblin');
SET @Orc = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Orc');
SET @Ogre = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Ogre');
SET @Troll = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troll');
SET @Changeling = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Changeling');
SET @Wolfen = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Wolfen');
SET @Coyle = (SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Coyle');

INSERT INTO palladium.RaceAttributes (`RaceId`, `IQ`, `IQBonus`, `ME`, `MEBonus`, `MA`, `MABonus`, `PS`, `PSBonus`, `PP`, `PPBonus`, `PE`, `PEBonus`,
	`PB`, `PBBonus`, `Spd`, `SpdBonus`, `PPE`, `PPEBonus`,`HF`,  `Alignment`, `SpdDig`, `SpdDigBonus`)
VALUES 	((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human'),3,0,3,0,3,0,3,0,3,0,3,0,3,0,3,0,2,0,0,'Any, usually lean toward good and selfish',0,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf'),3,1,3,0,2,0,3,0,4,0,3,0,5,0,3,0,2,0,0,'Any, usually lean toward good and selfish',0,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf'),3,0,3,0,2,0,4,6,3,0,4,0,2,2,2,0,2,0,0,'Any, usually lean toward good and selfish',1,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome'),3,0,1,6,3,4,1,4,4,0,3,6,4,0,2,0,2,0,0,'Any, but most tend to be good or selfish; an evil gnome is a rarity',1,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte'),2,0,2,0,3,0,4,4,3,6,3,0,2,0,6,0,2,0,0,'Typically anarchist or evil; but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold'),3,0,2,0,3,0,3,3,3,0,3,0,1,6,3,0,4,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Goblin'),2,0,3,0,2,0,3,0,3,6,3,0,2,0,3,0,6,0,0,'Typically anarchist or evil; but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Hob-Goblin'),2,0,3,6,2,0,3,0,3,0,3,0,2,0,3,0,4,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Orc'),2,0,2,0,3,0,3,8,3,0,3,2,2,0,3,0,2,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',0,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Ogre'),3,0,3,0,2,0,4,4,3,0,3,6,2,0,3,0,3,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',0,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troll'),3,0,2,0,2,0,4,10,4,0,3,6,1,4,2,0,3,0,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',0,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Changeling'),3,0,4,6,4,0,3,0,3,0,2,0,2,0,2,0,5,0,0,'Any',0,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Wolfen'),3,0,3,0,2,0,4,1,3,0,3,0,3,0,4,0,3,0,0,'Any, but tend toward principled and aberrant, both alignments with a strong personal code of honor',0,0),
        ((SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Coyle'),3,0,3,0,2,0,3,1,4,1,3,0,3,0,3,0,3,0,0,'Any, but tend toward anarchist and miscreant; the antithesis of the noble Wolfen',0,0);

  --OCC
INSERT INTO palladium.OCCType (`Desc`)
VALUES ('NONE'),('Clergy'),
('Men of Arms'),
('Optional'),
('Practitioners of Magic'),
('Psychics');

INSERT INTO palladium.OCC (`Desc`, `OCCTypeId`)
VALUES  ('NONE SELECTED', (SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'NONE')),
        ('Druid',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Clergy')),
        ('Monk',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Clergy')),
        ('Priest of Light',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Clergy')),
        ('Priest of Darkness',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Clergy')),
        ('Assassin',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Men of Arms')),
        ('Knight',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Men of Arms')),
        ('Long Bowman',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Men of Arms')),
        ('Mercenary Warrior',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Men of Arms')),
        ('Palladin',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Men of Arms')),
        ('Ranger',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Men of Arms')),
        ('Soldier',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Men of Arms')),
        ('Thief',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Men of Arms')),
        ('Merchant',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Optional')),
        ('Noble',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Optional')),
        ('Scholar',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Optional')),
        ('Squire',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Optional')),
        ('Vagabond',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Optional')),
        ('Peasant',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Optional')),
        ('Farmer',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Optional')),
        ('Diabolist',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Practitioners of Magic')),
        ('Summoner',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Practitioners of Magic')),
        ('Warlock',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Practitioners of Magic')),
        ('Witch',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Practitioners of Magic')),
        ('Wizard',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Practitioners of Magic')),
        ('Mind Mage',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Psychics')),
        ('Psi-Healer',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Psychics')),
        ('Psi-Mystic',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Psychics')),
        ('Psychic Sensitive',(SELECT ot.Id FROM palladium.OCCType ot WHERE ot.`Desc` = 'Psychics'));

  --SKILLS
INSERT INTO palladium.SkillCategory (`Desc`)
VALUES  ('Communications & Performing Arts'),('Domestic'),('Espionage'),('Horsemanship'),('Medical'),('Military'),
        ('Physical'),('Rogue/Thief'),('Science'),('Scholar, Noble & Technical'),('Weapon Proficiencies'),('Wilderness');

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
VALUES  (1,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf'),0,0,60,'Feet',NULL),
        (1,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf'),0,0,90,'Feet',NULL),(2,3,40,5,0,'%',NULL),(3,3,30,5,0,'%','detection and deactivation of traps is done at half normal architecture skill level'),
        (4,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf'),40,5,0,'%',NULL),(5,3,30,5,0,'%','-25% if in unfamiliar area'),(6,3,10,0,0,'%','Equal to Field Armorer'),(7,3,10,0,0,'%','Same as Gemology'),
        (1,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome'),0,0,90,'Feet',NULL),(2,4,30,5,0,'%',NULL),(3,4,20,5,0,'%','detection and deactivation of traps is done at half normal architecture skill level'),
        (4,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome'),30,5,0,'%',NULL),(5,4,20,5,0,'%','-20% if in unfamiliar area'),
        (1,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte'),600,0,0,'Feet','day vision 30ft'),(2,5,30,5,0,'%',NULL),(3,5,20,5,0,'%','detection and deactivation of traps is done at half normal architecture skill level'),
        (4,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte'),40,5,0,'%',NULL),(5,5,15,5,0,'%','-20% if in unfamiliar area'),
        (1,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold'),0,0,400,'Feet','day vision 40ft'),(2,6,40,5,0,'%',NULL),(3,6,30,5,0,'%','detection and deactivation of traps is done at half normal architecture skill level'),
        (4,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold'),40,5,0,'%',NULL),(5,6,30,5,0,'%','-25% if in unfamiliar area'),(6,6,10,0,0,'%','Equal to Field Armorer, +10 recognize weapon quality'),
        (7,(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold'),10,0,0,'%','Same as Gemology, art (limted to jewelry) and gems');
        -- stopped after Kobold. 


INSERT INTO palladium.Race_OCC (OccId, RaceId)
VALUES  ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Druid'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Monk'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Light'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Darkness'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Assassin'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Knight'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Long Bowman'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mercenary Warrior'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Palladin'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Ranger'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Soldier'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Thief'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Merchant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Noble'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Scholar'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Squire'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Vagabond'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Peasant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Farmer'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Diabolist'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Summoner'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Warlock'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Witch'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Wizard'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mind Mage'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psi-Healer'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psi-Mystic'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psychic Sensitive'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Human')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Druid'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Monk'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Light'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Darkness'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Assassin'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Knight'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Long Bowman'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mercenary Warrior'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Palladin'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Ranger'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Soldier'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Thief'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Merchant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Noble'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Scholar'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Squire'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Vagabond'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Peasant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Farmer'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Diabolist'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Summoner'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Warlock'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Witch'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Wizard'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mind Mage'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psi-Healer'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psi-Mystic'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psychic Sensitive'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Elf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Druid'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Monk'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Light'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Darkness'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Assassin'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Knight'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Long Bowman'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mercenary Warrior'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Palladin'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Ranger'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Soldier'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Thief'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Merchant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Noble'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Scholar'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Squire'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Vagabond'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Peasant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Farmer'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mind Mage'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psi-Healer'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psi-Mystic'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psychic Sensitive'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Dwarf')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Druid'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Monk'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Light'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Darkness'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Assassin'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mercenary Warrior'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Ranger'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Soldier'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Thief'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Merchant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Noble'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Scholar'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Squire'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Vagabond'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Peasant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Farmer'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Diabolist'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Summoner'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Warlock'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Witch'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Wizard'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Gnome')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Monk'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Assassin'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mercenary Warrior'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Soldier'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Thief'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Vagabond'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Troglodyte')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Druid'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Monk'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Light'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Priest of Darkness'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Assassin'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mercenary Warrior'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Ranger'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Soldier'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Thief'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Merchant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Noble'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Scholar'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Squire'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Vagabond'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Peasant'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Farmer'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Diabolist'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Summoner'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Warlock'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Witch'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Wizard'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Mind Mage'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psi-Healer'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psi-Mystic'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold')),
        ((SELECT o.Id FROM palladium.OCC o WHERE o.Desc = 'Psychic Sensitive'),(SELECT r.Id FROM palladium.Race r WHERE r.`Desc` = 'Kobold'));
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