
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';



CREATE SCHEMA IF NOT EXISTS `proyectoSA` DEFAULT CHARACTER SET utf8 ;
USE `proyectoSA` ;


CREATE TABLE IF NOT EXISTS `proyectoSA`.`Region` (
  `idRegion` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idRegion`));


CREATE TABLE IF NOT EXISTS `proyectoSA`.`User` (
  `idUser` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NULL,
  `lastName` VARCHAR(100) NULL,
  `username` VARCHAR(100) NULL,
  `birthDate` VARCHAR(45) NULL,
  `email` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  `imageURL` VARCHAR(200) NULL,
  `isActive` TINYINT NULL DEFAULT 0,
  `verifyToken` VARCHAR(60) NULL,
  `idRegion` INT NOT NULL,
  `timeOut` BIGINT NULL,
  PRIMARY KEY (`idUser`),
  INDEX `FK_Region_UserRegion_idx` (`idRegion` ASC) VISIBLE,
  CONSTRAINT `FK_Region_UserRegion`
    FOREIGN KEY (`idRegion`)
    REFERENCES `proyectoSA`.`Region` (`idRegion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


CREATE TABLE IF NOT EXISTS `proyectoSA`.`Game` (
  `idGame` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `imageURL` VARCHAR(200) NULL,
  `releaseDate` VARCHAR(45) NULL,
  `restrictionAge` VARCHAR(5) NULL,
  `group` INT NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  `description` TEXT NULL,
  `isGlobal` TINYINT NULL DEFAULT 0,
  `globalPrice` DECIMAL(15,2) NULL DEFAULT 0,
  `globalDiscount` DECIMAL(15,2) NULL DEFAULT 0,
  PRIMARY KEY (`idGame`));


CREATE TABLE IF NOT EXISTS `proyectoSA`.`Category` (
  `idCategory` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idCategory`));


CREATE TABLE IF NOT EXISTS `proyectoSA`.`GameCategory` (
  `idGameCategory` INT NOT NULL AUTO_INCREMENT,
  `idGame` INT NOT NULL,
  `idCategory` INT NOT NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idGameCategory`),
  INDEX `FK_idGame_idx` (`idGame` ASC) VISIBLE,
  INDEX `FK_idCategory_idx` (`idCategory` ASC) VISIBLE,
  CONSTRAINT `FK_idGame_GameCategory`
    FOREIGN KEY (`idGame`)
    REFERENCES `proyectoSA`.`Game` (`idGame`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `FK_idCategory_GameCategory`
    FOREIGN KEY (`idCategory`)
    REFERENCES `proyectoSA`.`Category` (`idCategory`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


CREATE TABLE IF NOT EXISTS `proyectoSA`.`DLC` (
  `idDLC` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NULL,
  `idGame` INT NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  `imageURL` VARCHAR(200) NULL,
  `description` TEXT NULL,
  `releaseDate` VARCHAR(100) NULL,
  `isGlobal` TINYINT NULL DEFAULT 0,
  `globalPrice` DECIMAL(15,2) NULL DEFAULT 0,
  `globalDiscount` DECIMAL(15,2) NULL DEFAULT 0,
  PRIMARY KEY (`idDLC`),
  CONSTRAINT `FK_idGame`
    FOREIGN KEY (`idGame`)
    REFERENCES `proyectoSA`.`Game` (`idGame`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


CREATE TABLE IF NOT EXISTS `proyectoSA`.`RegionPrice` (
  `idRegionPrice` INT NOT NULL AUTO_INCREMENT,
  `idGame` INT NULL,
  `idRegion` INT NOT NULL,
  `idDLC` INT NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  `price` DECIMAL(15,2) NULL,
  `discount` DECIMAL(15,2) NULL,
  `isDLC` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idRegionPrice`),
  INDEX `FK_idGame_idx` (`idGame` ASC) VISIBLE,
  INDEX `FK_idRegion_idx` (`idRegion` ASC) VISIBLE,
  INDEX `FK_idDLC_RegionPrice_idx` (`idDLC` ASC) VISIBLE,
  CONSTRAINT `FK_idGame_RegionPrice`
    FOREIGN KEY (`idGame`)
    REFERENCES `proyectoSA`.`Game` (`idGame`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `FK_idRegion_RegionPrice`
    FOREIGN KEY (`idRegion`)
    REFERENCES `proyectoSA`.`Region` (`idRegion`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `FK_idDLC_RegionPrice`
    FOREIGN KEY (`idDLC`)
    REFERENCES `proyectoSA`.`DLC` (`idDLC`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


CREATE TABLE IF NOT EXISTS `proyectoSA`.`GameDiscount` (
  `idGameDiscount` INT NOT NULL AUTO_INCREMENT,
  `idGame` INT NULL,
  `idDLC` INT NULL,
  `discount` DECIMAL(15,2) NULL,
  `startDateTime` VARCHAR(100) NULL,
  `finishDateTime` VARCHAR(100) NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  `isDLC` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idGameDiscount`),
  INDEX `FK_idGame_GameDiscount_idx` (`idGame` ASC) VISIBLE,
  INDEX `FK_idDLC_GameDiscount_idx` (`idDLC` ASC) VISIBLE,
  CONSTRAINT `FK_idGame_GameDiscount`
    FOREIGN KEY (`idGame`)
    REFERENCES `proyectoSA`.`Game` (`idGame`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `FK_idDLC_GameDiscount`
    FOREIGN KEY (`idDLC`)
    REFERENCES `proyectoSA`.`DLC` (`idDLC`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


CREATE TABLE IF NOT EXISTS `proyectoSA`.`Developer` (
  `idDeveloper` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NULL,
  `country` VARCHAR(45) NULL,
  `imageURL` VARCHAR(200) NULL,
  `email` VARCHAR(100) NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idDeveloper`));


CREATE TABLE IF NOT EXISTS `proyectoSA`.`GameDeveloper` (
  `idGameDeveloper` INT NOT NULL AUTO_INCREMENT,
  `idGame` INT NOT NULL,
  `idDeveloper` INT NOT NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idGameDeveloper`),
  INDEX `FK_idGame_GameDeveloper_idx` (`idGame` ASC) VISIBLE,
  INDEX `FK_idDeveloper_GameDeveloper_idx` (`idDeveloper` ASC) VISIBLE,
  CONSTRAINT `FK_idGame_GameDeveloper`
    FOREIGN KEY (`idGame`)
    REFERENCES `proyectoSA`.`Game` (`idGame`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `FK_idDeveloper_GameDeveloper`
    FOREIGN KEY (`idDeveloper`)
    REFERENCES `proyectoSA`.`Developer` (`idDeveloper`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


CREATE TABLE IF NOT EXISTS `proyectoSA`.`MyGames` (
  `idMyGame` INT NOT NULL AUTO_INCREMENT,
  `idUser` INT NOT NULL,
  `idGame` INT NOT NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  `isWishlist` TINYINT NULL DEFAULT 0,
  `isLibrary` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idMyGame`),
  INDEX `FK_idUser_Wishlist_idx` (`idUser` ASC) VISIBLE,
  INDEX `FK_idGame_Wishlist_idx` (`idGame` ASC) VISIBLE,
  CONSTRAINT `FK_idUser_Wishlist`
    FOREIGN KEY (`idUser`)
    REFERENCES `proyectoSA`.`User` (`idUser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `FK_idGame_Wishlist`
    FOREIGN KEY (`idGame`)
    REFERENCES `proyectoSA`.`Game` (`idGame`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


CREATE TABLE IF NOT EXISTS `proyectoSA`.`Sale` (
  `idSale` INT NOT NULL AUTO_INCREMENT,
  `idUser` INT NOT NULL,
  `saleDate` VARCHAR(45) NULL,
  `total` DECIMAL(15,2) NULL,
  `metododePago` VARCHAR(45) NULL DEFAULT 'paypal',
  `isDeleted` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idSale`),
  INDEX `FK_idUser_Sale_idx` (`idUser` ASC) VISIBLE,
  CONSTRAINT `FK_idUser_Sale`
    FOREIGN KEY (`idUser`)
    REFERENCES `proyectoSA`.`User` (`idUser`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


CREATE TABLE IF NOT EXISTS `proyectoSA`.`SaleDetail` (
  `idSaleDetail` INT NOT NULL AUTO_INCREMENT,
  `idSale` INT NOT NULL,
  `idGame` INT NULL,
  `idDLC` INT NULL,
  `subTotal` DECIMAL(15,2) NULL,
  `isDeleted` TINYINT NULL DEFAULT 0,
  `isDLC` TINYINT NULL DEFAULT 0,
  PRIMARY KEY (`idSaleDetail`),
  INDEX `FK_idSale_SaleDatail_idx` (`idSale` ASC) VISIBLE,
  INDEX `FK_idGame_SaleDetail_idx` (`idGame` ASC) VISIBLE,
  INDEX `FK_idDLC_SaleDetail_idx` (`idDLC` ASC) VISIBLE,
  CONSTRAINT `FK_idSale_SaleDatail`
    FOREIGN KEY (`idSale`)
    REFERENCES `proyectoSA`.`Sale` (`idSale`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `FK_idGame_SaleDetail`
    FOREIGN KEY (`idGame`)
    REFERENCES `proyectoSA`.`Game` (`idGame`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `FK_idDLC_SaleDetail`
    FOREIGN KEY (`idDLC`)
    REFERENCES `proyectoSA`.`DLC` (`idDLC`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
