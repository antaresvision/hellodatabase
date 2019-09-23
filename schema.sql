CREATE TABLE public.items (
	id serial NOT NULL,
	ntinid integer NOT NULL,
	"serial" varchar NOT NULL,
	status int NOT NULL,
	CONSTRAINT items_pk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX items_serial_idx ON public.items ("serial",ntinid);
