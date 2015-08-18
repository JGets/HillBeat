CREATE TABLE t_languages
(
c_id_lang INT,
c_lang_name_short VARCHAR(256),
c_lang_name_long VARCHAR(256)
);


CREATE TABLE t_bill_progress
(
c_id_progress INT,
c_progress_value VARCHAR(8)
);


CREATE TABLE t_parliaments
(
c_id_parliament INT,
c_parliament_num INT,
c_session_num INT
);


CREATE TABLE t_bill_types
(
c_id_type INT
);


CREATE TABLE t_bill_type_names
(
c_id_type INT,
c_id_lang INT,
c_type_name VARCHAR(256)
);


CREATE TABLE t_persons
(
c_id_person INT,
c_person_name_full VARCHAR(256),
c_person_name_first VARCHAR(256),
c_person_name_middle VARCHAR(256),
c_person_name_last VARCHAR(256)
);


CREATE TABLE t_parties
(
c_id_party INT
);


CREATE TABLE t_party_names
(
c_id_party INT,
c_id_lang INT,
c_party_name_long VARCHAR(256),
c_party_name_short VARCHAR(256)
);


CREATE TABLE t_sponsors
(
c_id_sponsor INT,
c_id_person INT,
c_id_party INT
);


CREATE TABLE t_sponsor_titles
(
c_id_sponsor INT,
c_id_lang INT,
c_sponsor_title VARCHAR(256)
);


CREATE TABLE t_bill_stage
(
c_id_stage INT,
c_stage_name VARCHAR(256)
);


CREATE TABLE t_bills
(
c_id_bill INT,
c_last_update DATE,
c_introduced DATE,
c_id_parliament INT,
c_id_type INT,
c_id_sponsor INT,
c_id_prime_minister INT,
c_id_stage INT,
c_id_progress INT
);


CREATE TABLE t_bill_statutes
(
c_id_bill INT,
c_statute_year INT,
c_statute_chapter INT
);


CREATE TABLE t_bill_numbers
(
c_id_bill INT,
c_bill_prefix VARCHAR(4),
c_bill_number INT
);


CREATE TABLE t_bill_titles
(
c_id_bill INT,
c_id_lang INT,
c_bill_title_long VARCHAR(256),
c_bill_title_short VARCHAR(256)
);


CREATE TABLE t_publications
(
c_id_pub INT,
c_id_bill INT
);


CREATE TABLE t_publication_files
(
c_id_pub INT,
c_id_lang INT,
c_pub_path VARCHAR(256)
);


CREATE TABLE t_publication_titles
(
c_id_pub INT,
c_id_lang INT,
c_pub_title VARCHAR(256)
);


CREATE TABLE t_bill_events
(
c_id_event INT,
c_id_bill INT,
c_id_chamber INT,
c_event_date DATE,
c_event_meeting_num INT,
c_id_event_title INT
);

CREATE TABLE t_bill_event_titles
(
c_id_event_title INT,
c_id_lang INT,
c_event_title VARCHAR(256)
);

