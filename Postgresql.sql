-- Database: perbukuan

-- DROP DATABASE IF EXISTS perbukuan;

CREATE DATABASE perbukuan
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;


-- SEQUENCE: public.perpustakaans_id_seq

-- DROP SEQUENCE IF EXISTS public.perpustakaans_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.perpustakaans_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.perpustakaans_id_seq
    OWNER TO postgres;


-- Table: public.perpustakaans

-- DROP TABLE IF EXISTS public.perpustakaans;

CREATE TABLE IF NOT EXISTS public.perpustakaans
(
    id integer NOT NULL DEFAULT nextval('perpustakaans_id_seq'::regclass),
    judulbuku character varying(30) COLLATE pg_catalog."default" NOT NULL,
    deskripsibuku character varying(1000) COLLATE pg_catalog."default",
    isbn character varying(100) COLLATE pg_catalog."default",
    issn character varying(100) COLLATE pg_catalog."default",
    bahasabuku character varying(100) COLLATE pg_catalog."default",
    CONSTRAINT perpustakaans_pk PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.perpustakaans
    OWNER to postgres;

GRANT ALL ON TABLE public.perpustakaans TO postgres;