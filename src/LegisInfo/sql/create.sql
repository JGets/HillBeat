CREATE TABLE t_languages
(
c_id_lang INT NOT NULL,
c_lang_name_short VARCHAR(2),
c_lang_name_long VARCHAR(256),
PRIMARY KEY (c_id_lang)
);


CREATE TABLE t_bill_progress
(
c_id_progress INT NOT NULL,
c_progress_value VARCHAR(8) NOT NULL,
PRIMARY KEY (c_id_progress)
);


CREATE TABLE t_parliaments
(
c_id_parliament INT NOT NULL,
c_parliament_num INT NOT NULL,
c_session_num INT NOT NULL,
PRIMARY KEY (c_id_parliament)
);


CREATE TABLE t_bill_types
(
c_id_type INT NOT NULL,
PRIMARY KEY (c_id_type)
);


CREATE TABLE t_bill_type_names
(
c_id_type INT NOT NULL,
c_id_lang INT NOT NULL,
c_type_name VARCHAR(256) NOT NULL,
PRIMARY KEY (c_id_type),
FOREIGN KEY (c_id_lang) REFERENCES t_languages(c_id_lang)
);


CREATE TABLE t_persons
(
c_id_person INT NOT NULL,
c_person_name_full VARCHAR(256) NOT NULL,
c_person_name_first VARCHAR(256),
c_person_name_middle VARCHAR(256),
c_person_name_last VARCHAR(256),
PRIMARY KEY (c_id_person)
);


CREATE TABLE t_parties
(
c_id_party INT NOT NULL,
PRIMARY KEY (c_id_party)
);


CREATE TABLE t_party_names
(
c_id_party INT NOT NULL,
c_id_lang INT NOT NULL,
c_party_name_long VARCHAR(256),
c_party_name_short VARCHAR(256),
FOREIGN KEY (c_id_party) REFERENCES t_parties(c_id_party),
FOREIGN KEY (c_id_lang) REFERENCES t_languages(c_id_lang)
);


CREATE TABLE t_sponsors
(
c_id_sponsor INT NOT NULL,
c_id_person INT NOT NULL,
c_id_party INT NOT NULL,
PRIMARY KEY (c_id_sponsor),
FOREIGN KEY (c_id_person) REFERENCES t_persons(c_id_person),
FOREIGN KEY (c_id_party) REFERENCES t_parties(c_id_party)
);


CREATE TABLE t_sponsor_titles
(
c_id_sponsor INT NOT NULL,
c_id_lang INT NOT NULL,
c_sponsor_title VARCHAR(256) NOT NULL,
FOREIGN KEY (c_id_sponsor) REFERENCES t_sponsors(c_id_sponsor),
FOREIGN KEY (c_id_lang) REFERENCES t_languages(c_id_lang)
);


CREATE TABLE t_bill_stages
(
c_id_stage INT NOT NULL,
c_stage_name VARCHAR(256) NOT NULL,
PRIMARY KEY (c_id_stage)
);


CREATE TABLE t_bills
(
c_id_bill INT NOT NULL,
c_last_update DATE,
c_introduced DATE,
c_id_parliament INT NOT NULL,
c_id_type INT NOT NULL,
c_id_sponsor INT NOT NULL,
c_id_prime_minister INT NOT NULL,
c_id_stage INT NOT NULL,
c_id_progress INT NOT NULL,
PRIMARY KEY (c_id_bill),
FOREIGN KEY (c_id_parliament) REFERENCES t_parliaments(c_id_parliament),
FOREIGN KEY (c_id_type) REFERENCES t_bill_types(c_id_type),
FOREIGN KEY (c_id_sponsor) REFERENCES t_sponsors(c_id_sponsor),
FOREIGN KEY (c_id_prime_minister) REFERENCES t_sponsors(c_id_sponsor),
FOREIGN KEY (c_id_stage) REFERENCES t_bill_stages(c_id_stage),
FOREIGN KEY (c_id_progress) REFERENCES t_bill_progress(c_id_progress)
);


CREATE TABLE t_bill_statutes
(
c_id_bill INT NOT NULL,
c_statute_year INT,
c_statute_chapter INT,
FOREIGN KEY (c_id_bill) REFERENCES t_bills(c_id_bill)
);


CREATE TABLE t_bill_numbers
(
c_id_bill INT NOT NULL,
c_bill_prefix VARCHAR(4) NOT NULL,
c_bill_number INT NOT NULL,
FOREIGN KEY (c_id_bill) REFERENCES t_bills(c_id_bill)
);


CREATE TABLE t_bill_titles
(
c_id_bill INT NOT NULL,
c_id_lang INT NOT NULL,
c_bill_title_long VARCHAR(256),
c_bill_title_short VARCHAR(256),
FOREIGN KEY (c_id_bill) REFERENCES t_bills(c_id_bill),
FOREIGN KEY (c_id_lang) REFERENCES t_languages(c_id_lang)
);


CREATE TABLE t_publications
(
c_id_pub INT NOT NULL,
c_id_bill INT NOT NULL,
PRIMARY KEY (c_id_pub),
FOREIGN KEY (c_id_bill) REFERENCES t_bills(c_id_bill)
);


CREATE TABLE t_publication_files
(
c_id_pub INT NOT NULL,
c_id_lang INT NOT NULL,
c_pub_path VARCHAR(256) NOT NULL,
FOREIGN KEY (c_id_pub) REFERENCES t_publications(c_id_pub),
FOREIGN KEY (c_id_lang) REFERENCES t_languages(c_id_lang)
);


CREATE TABLE t_publication_titles
(
c_id_pub INT NOT NULL,
c_id_lang INT NOT NULL,
c_pub_title VARCHAR(256) NOT NULL,
FOREIGN KEY (c_id_pub) REFERENCES t_publications(c_id_pub),
FOREIGN KEY (c_id_lang) REFERENCES t_languages(c_id_lang)
);


CREATE TABLE t_chambers
(
c_id_chamber INT NOT NULL,
c_chamber_name VARCHAR(256) NOT NULL,
PRIMARY KEY (c_id_chamber)
);


CREATE TABLE t_bill_events
(
c_id_event INT NOT NULL,
c_id_bill INT NOT NULL,
c_id_chamber INT NOT NULL,
c_event_date DATE,
c_event_meeting_num INT,
PRIMARY KEY (c_id_event),
FOREIGN KEY (c_id_bill) REFERENCES t_bills(c_id_bill),
FOREIGN KEY (c_id_chamber) REFERENCES t_chambers(c_id_chamber)
);

CREATE TABLE t_bill_event_titles
(
c_id_event INT NOT NULL,
c_id_lang INT NOT NULL,
c_event_title VARCHAR(256) NOT NULL,
FOREIGN KEY (c_id_event) REFERENCES t_bill_events(c_id_event),
FOREIGN KEY (c_id_lang) REFERENCES t_languages(c_id_lang)
);

