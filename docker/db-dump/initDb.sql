ALTER TABLE RaceAttributes
DROP FOREIGN KEY FK__RaceAttributes__Race;

ALTER TABLE RaceAttributes
DROP FOREIGN KEY RaceAttributes_ibfk_1;

ALTER TABLE Character
DROP FOREIGN KEY FK__Character__race;

ALTER TABLE Character
DROP FOREIGN KEY Character_ibfk_1;

DROP TABLE IF EXISTS `Race`;
CREATE TABLE IF NOT EXISTS `Race` (
  `Id` int not null AUTO_INCREMENT,
  `Name` varchar(25) not null default 'sadness',
  Primary key (id)
);

DROP TABLE IF EXISTS `RaceAttributes`;
CREATE TABLE IF NOT EXISTS `RaceAttributes` (
  `Id` int not null AUTO_INCREMENT,
  `RaceId` int not null,
  `IQ` smallint not null default 0,
  `IQBonus` smallint not null default 0,
  `ME` smallint not null default 0,
  `MEBonus` smallint not null default 0,
  `MA` smallint not null default 0,
  `MABonus` smallint not null default 0,
  `PS` smallint not null default 0,
  `PSBonus` smallint not null default 0,
  `PP` smallint not null default 0,
  `PPBonus` smallint not null default 0,
  `PE` smallint not null default 0,
  `PEBonus` smallint not null default 0,
  `PB` smallint not null default 0,
  `PBBonus` smallint not null default 0,
  `Spd` smallint not null default 0,
  `SpdBonus` smallint not null default 0,
  `PPE` smallint not null default 0,
  `PPEBonus` smallint not null default 0,
  `Alignment` varchar(255) default "any",
  `SpdDig` smallint not null default 0,
  `SpdDigBonus` smallint not null default 0,
  constraint PK_raceAttributes Primary key (Id, RaceId),
  CONSTRAINT FK__RaceAttributes__Race FOREIGN KEY (RaceId)
  REFERENCES Race(Id)
);

ALTER TABLE `RaceAttributes` ADD FOREIGN KEY (`RaceId`) REFERENCES `Race` (`Id`);

DROP TABLE IF EXISTS `Character`;
CREATE TABLE IF NOT EXISTS `Character` (
  `Id` int not null AUTO_INCREMENT,
  `Name` varchar(25) not null,
  `RaceId` int not null,
  `Lvl` smallint not null default 1,
  `IQ` smallint not null default 0,
  `ME` smallint not null default 0,
  `MA` smallint not null default 0,
  `PS` smallint not null default 0,
  `PP` smallint not null default 0,
  `PE` smallint not null default 0,
  `PB` smallint not null default 0,
  `Spd` smallint not null default 0,
  `PPE` smallint not null default 0,
  `SpdDig` smallint not null default 0,
  Primary key (Id),
  CONSTRAINT FK__Character__Race FOREIGN KEY (RaceId)
  REFERENCES Race(Id)
);

ALTER TABLE `Character` ADD FOREIGN KEY (`RaceId`) REFERENCES `Race` (`Id`);

insert into Race (Name)
values ('Human'), ('Elf'), ('Dwarf'), ('Gnome'), ('Troglodyte'), ('Kobold'), ('Goblin'), 
	('Hob-Goblin'), ('Orc'), ('Ogre'), ('Troll'), ('Changeling'), ('Wolfen'), ('Coyle');
    

Insert into RaceAttributes (RaceId, IQ, IQBonus, ME, MEBonus, MA, MABonus, PS, PSBonus, PP, PPBonus, PE, PEBonus,
	PB, PBBonus, Spd, SpdBonus, PPE, PPEBonus, Alignment, SpdDig, SpdDigBonus)
values 	(1,3,0,3,0,3,0,3,0,3,0,3,0,3,0,3,0,2,0,'Any, usually lean toward good and selfish',0,0),
        (2,3,1,3,0,2,0,3,0,4,0,3,0,5,0,3,0,2,0,'Any, usually lean toward good and selfish',0,0),
        (3,3,0,3,0,2,0,4,6,3,0,4,0,2,2,2,0,2,0,'Any, usually lean toward good and selfish',1,0),
        (4,3,0,1,6,3,4,1,4,4,0,3,6,4,0,2,0,2,0,'Any, but most tend to be good or selfish; an evil gnome is a rarity',1,0),
        (5,2,0,2,0,3,0,4,4,3,6,3,0,2,0,6,0,2,0,'Typically anarchist or evil; but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        (6,3,0,2,0,3,0,3,3,3,0,3,0,1,6,3,0,4,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        (7,2,0,3,0,2,0,3,0,3,6,3,0,2,0,3,0,6,0,'Typically anarchist or evil; but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        (8,2,0,3,6,2,0,3,0,3,0,3,0,2,0,3,0,4,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',1,0),
        (9,2,0,2,0,3,0,3,8,3,0,3,2,2,0,3,0,2,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',0,0),
        (10,3,0,3,0,2,0,4,4,3,0,3,6,2,0,3,0,3,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',0,0),
        (11,3,0,2,0,2,0,4,10,4,0,3,6,1,4,2,0,3,0,'Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good',0,0),
        (12,3,0,4,6,4,0,3,0,3,0,2,0,2,0,2,0,5,0,'Any',0,0),
        (13,3,0,3,0,2,0,4,1,3,0,3,0,3,0,4,0,3,0,'Any, but tend toward principled and aberrant, both alignments with a strong personal code of honor',0,0),
        (14,3,0,3,0,2,0,3,1,4,1,3,0,3,0,3,0,3,0,'Any, but tend toward anarchist and miscreant; the antithesis of the noble Wolfen',0,0);