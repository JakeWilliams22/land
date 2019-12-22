USE Land;
CREATE TABLE JOIN_EMAIL_LIST (
	id INTEGER NOT NULL PRIMARY KEY,
    landing_page_id varchar(100),
    prompt varchar(10000),
    button_text varchar(100)
);

INSERT INTO JOIN_EMAIL_LIST VALUES (-1, -1, "Keep up to date", "Join");

SELECT * FROM JOIN_EMAIL_LIST;