ALTER TABLE todoapp.users 
ADD password_hash VARCHAR(128) DEFAULT 'unset' NOT NULL,
ADD role VARCHAR(20) DEFAULT 'user' NOT NULL;




