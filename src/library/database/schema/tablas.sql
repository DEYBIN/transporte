
CREATE TABLE transporte.dbo.Clientes (
	c_docu char(2) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	n_docu varchar(11) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NOT NULL,
	l_clie varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	k_gene int DEFAULT 0 NULL,
	f_naci varchar(10) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_dire varchar(400) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_refe varchar(400) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	c_ubig char(6) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	n_tele varchar(30) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	n_celu varchar(30) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_obse varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	CONSTRAINT PK__Clientes__9F49F5FE8309AAB9 PRIMARY KEY (n_docu)
);


CREATE TABLE transporte.dbo.Seguridad (
	id varchar(36) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	users varchar(25) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	n_docu varchar(11) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_nomb varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_apl1 varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_apl2 varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_emai varchar(150) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	n_celu varchar(11) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	k_carg int DEFAULT 0 NULL,
	l_pass varchar(200) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	CONSTRAINT PK__Segurida__3213E83F4741BF87 PRIMARY KEY (id),
	CONSTRAINT UQ__Segurida__2B8C777D20A884B0 UNIQUE (users)
);


CREATE TABLE transporte.dbo.ClientesCars (
	c_plac varchar(6) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NOT NULL,
	n_docu varchar(11) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NOT NULL,
	l_marc varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_mode varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_color varchar(70) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	c_year char(4) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	c_mode char(4) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	n_seri varchar(17) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	n_pasa varchar(20) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	l_obse varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS DEFAULT '' NULL,
	k_stad int DEFAULT 0 NULL,
	CONSTRAINT PK__Clientes__4BC8D1A7725D2855 PRIMARY KEY (c_plac),
	CONSTRAINT FK__ClientesC__n_doc__59063A47 
    FOREIGN KEY (n_docu) REFERENCES transporte.dbo.Clientes(n_docu)
);


CREATE TABLE Servicios (
	id_serv varchar(36)  NOT NULL,
	c_year char(4)  NOT NULL,
	c_mes char(2)  NOT NULL,
	n_docu varchar(11)  NOT NULL,
	f_fact varchar(10)  NOT NULL,
	s_impo numeric(20,7) DEFAULT 0,
	c_plac varchar(6)  NOT NULL,
	k_stad int DEFAULT 0,
	f_digi datetime DEFAULT getdate(),
	CONSTRAINT Servicios_n_docu_c_plac UNIQUE (n_docu,c_plac),
    FOREIGN KEY (c_plac) REFERENCES ClientesCars(c_plac),
    FOREIGN KEY (n_docu) REFERENCES Clientes(n_docu)
);