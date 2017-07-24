create table t_bing_image (
    id       integer auto_increment primary key,
    url_path varchar(256),
    image_name varchar(64),
    belong_date datetime,
    description    varchar(256),
    unique(image_name)
)