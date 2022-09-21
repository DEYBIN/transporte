CREATE TABLE Servicios(
id_serv VARCHAR (36) NOT NULL PRIMARY KEY,
c_year 	CHAR (4) NOT NULL,
c_mes CHAR (2) NOT NULL,
n_docu VARCHAR (11) NOT NULL,
f_fact VARCHAR (10) NOT NULL,
s_impo NUMERIC (20,7) DEFAULT 0,
c_plac VARCHAR (6) NOT NULL,
k_stad INT DEFAULT 0,
f_digi DATETIME DEFAULT getdate(),
FOREIGN KEY (n_docu) REFERENCES Clientes(n_docu),
FOREIGN KEY (c_plac) REFERENCES ClientesCars(c_plac)
)

