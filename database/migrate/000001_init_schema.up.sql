CREATE TABLE tbl_employe_account (
    id smallint PRIMARY KEY,
    username VARCHAR(12),
    password CHAR(16),
    refreshToken VARCHAR(20),
    dateCreate TIMESTAMP,
    dateUpdate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tbl_employe_data (
    firstName VARCHAR(20) NOT NULL,
    lastName VARCHAR(20) NOT NULL,
    birth DATE,
    birthPlace CHAR(20),
    isMale CHAR(5) NOT NULL,
    phone INT NOT NULL,
    about VARCHAR(200) NOT NULL,
    employe_id smallint NOT NULL,
    PRIMARY KEY (employe_id),
    CONSTRAINT fk_employe_data_id 
        FOREIGN KEY (employe_id) 
    REFERENCES tbl_employe_account (id)
);

CREATE TABLE tbl_country (
    idCountry smallint NOT NULL,
    countryName VARCHAR(50) NOT NULL,
    PRIMARY KEY(idCountry)
);

CREATE TABLE tbl_province (
    idProvince smallint NOT NULL,
    provinceName VARCHAR(50) NOT NULL,
    PRIMARY KEY(idProvince)
);

CREATE TABLE tbl_district (
    idDistrict smallint NOT NULL,
    districtName VARCHAR(50) NOT NULL,
    PRIMARY KEY(idDistrict)
);


CREATE TABLE tbl_employe_address (
    idCountry  smallint NOT NULL,
    idProvince smallint NOT NULL,
    idDistrict smallint NOT NULL,
    address_1  VARCHAR(100),
    address_2  VARCHAR(100),
    postalCode smallint,
    employe_id smallint NOT NULL,
    PRIMARY KEY (idCountry,idProvince,idDistrict,employe_id),
    CONSTRAINT fk_address_employe_id 
        FOREIGN KEY (employe_id) 
    REFERENCES tbl_employe_account (id),
    CONSTRAINT fk_country_employe_idCountry 
        FOREIGN KEY (idCountry)
    REFERENCES tbl_country (idCountry),
    CONSTRAINT fk_province_employe_idProvince 
        FOREIGN KEY (idProvince)
    REFERENCES tbl_province (idProvince),
    CONSTRAINT fk_district_employe_idDistrict 
        FOREIGN KEY (idDistrict)
    REFERENCES tbl_district (idDistrict)
);

CREATE TABLE tbl_employe_attachment (
    portofolioFile CHAR(100),
    resumeFile CHAR(100),
    employe_id smallint NOT NULL,
    CONSTRAINT fk_employe_attach_id
        FOREIGN KEY (employe_id)
    REFERENCES tbl_employe_account(id)
);

CREATE TABLE tbl_employe_social_media (
    portofolioLink CHAR(100),
    githubLink CHAR(100),
    linkedinLink CHAR(100),
    blogLink CHAR(100),
    twitterLink CHAR(100),
    employe_id smallint NOT NULL,
    PRIMARY KEY(employe_id),
    CONSTRAINT fk_employe_socmed_id 
        FOREIGN KEY (employe_id)
    REFERENCES tbl_employe_account(id)
);

CREATE TABLE tbl_employe_experience (
    companyName VARCHAR(20),
    jobTitle    VARCHAR(20),
    jobDesc     VARCHAR(200),
    isActive    CHAR(5),
    startWork   DATE,
    endWork     DATE,
    employe_id smallint NOT NULL,
    PRIMARY KEY(employe_id),
    CONSTRAINT fl_employe_work_id
        FOREIGN KEY (employe_id)
    REFERENCES tbl_employe_account(id)
);

CREATE TABLE tbl_employe_education (
    institution VARCHAR(20),
    degree      VARCHAR(20),
    jobDesc     VARCHAR(200),
    certificare CHAR(100),
    startEducation DATE,
    endEducation   DATE,
    employe_id smallint NOT NULL,
    PRIMARY KEY(employe_id),
    CONSTRAINT fl_employe_edu_id
        FOREIGN KEY (employe_id)
    REFERENCES tbl_employe_account(id)
);


