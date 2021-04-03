-- 
-- Name Table : employe_accounts
-- 

CREATE TABLE  employe_accounts (
    id integer NOT NULL PRIMARY KEY,
    username VARCHAR(12),
    password CHAR(60),
    email VARCHAR(60),
    photo_profile VARCHAR(255),
    refresh_token TEXT,
    is_active TEXT,
    date_create timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    date_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

-- 
-- Name Table : employe_addresses
-- 

CREATE TABLE employe_addresses (
    country_name VARCHAR(50),
    province_name VARCHAR(50),
    district_name VARCHAR(50),
    address_1 VARCHAR(100),
    address_2 VARCHAR(100),
    postal_code bigint,
    employe_id integer NOT NULL,
    FOREIGN KEY (employe_id)
        REFERENCES employe_accounts (id)
    ON DELETE CASCADE
);

-- 
-- Name Table : employe_attachments
-- 

CREATE TABLE employe_attachments (
    portofolio_file VARCHAR(255),
    resume_file VARCHAR(255),
    resume_object VARCHAR(50),
    portofolio_object VARCHAR(50),
    employe_id integer NOT NULL,
    FOREIGN KEY (employe_id)
        REFERENCES employe_accounts (id)
    ON DELETE CASCADE
);

-- 
-- Name Table : employe_data
-- 

CREATE TABLE employe_data (
    first_name VARCHAR(20),
    last_name VARCHAR(20),
    birth date,
    birt_place CHAR(20),
    is_male CHAR(5),
    phone bigint,
    about VARCHAR(200),
    employe_id integer NOT NULL,
    FOREIGN KEY (employe_id)
        REFERENCES employe_accounts (id)
    ON DELETE CASCADE
);

-- 
-- Name Table : employe_educations
-- 

CREATE TABLE employe_educations (
    institution_name VARCHAR(20),
    degree VARCHAR(20),
    is_active CHAR(5),
    start_education date,
    end_education date,
    employe_id integer NOT NULL,
    FOREIGN KEY (employe_id)
        REFERENCES employe_accounts (id)
    ON DELETE CASCADE
);

-- 
-- Name Table : employe_experiences
-- 

CREATE TABLE employe_experiences (
    company_name VARCHAR(20),
    job_title VARCHAR(20),
    job_desc VARCHAR(200),
    is_active CHAR(5),
    start_work date,
    end_work date,
    employe_id integer NOT NULL,
    FOREIGN KEY (employe_id)
        REFERENCES employe_accounts (id)
    ON DELETE CASCADE
);

-- 
-- Name Table : employe_socials
-- 

CREATE TABLE employe_socials (
    portofolio_link CHAR(100),
    github_link CHAR(100),
    linkedin_link CHAR(100),
    blog_link CHAR(100),
    twitter_link CHAR(100),
    employe_id integer NOT NULL,
    FOREIGN KEY (employe_id)
        REFERENCES employe_accounts (id)
    ON DELETE CASCADE
);
