USE Land;
CREATE TABLE LANDING_PAGES (
	id varchar(100) not null primary key,
    title varchar(100),
    sub_title varchar(1000),
    body_text varchar(10000)
);

INSERT INTO LANDING_PAGES VALUES (-1, "Moolah", "Get some money for your kids", "We like to do good. You like to do good. Let's all do good", -1);

SELECT * FROM LANDING_PAGES;

