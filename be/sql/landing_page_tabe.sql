USE Land;
CREATE TABLE LANDING_PAGES (
	id varchar(100) not null primary key,
    title varchar(100),
    sub_title varchar(1000),
    body_text varchar(10000)
);

ALTER TABLE LANDING_PAGES ADD COLUMN GOOGLE_ANALYTICS_ID varchar(100);

INSERT INTO LANDING_PAGES VALUES (
	-2, 
    "Greenlight Financial", 
    "We give your kids the green light!", 
    "Financial literacy is essential in today's fast** paced** world**", 
    "UA-156318179-1");

SELECT * FROM LANDING_PAGES;

SELECT ID, TITLE, SUB_TITLE, BODY_TEXT FROM LANDING_PAGES WHERE ID = -2;